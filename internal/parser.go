package internal

import (
	"regexp"
	"strings"
)

var simpleCommentPattern = regexp.MustCompile(`//`)
var hasStringLiterals = regexp.MustCompile(`["'` + "`" + `]`)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) removeCommentFast(line string) string {
	if idx := strings.Index(line, "//"); idx != -1 {
		return strings.TrimRight(line[:idx], " \t")
	}

	return line
}

func (p *Parser) removeCommentSafe(line string) string {
	inString := false
	inRune := false
	inRawString := false

	for i := 0; i < len(line)-1; i++ {
		char := line[i]

		switch {
		case char == '`' && !inString && !inRune:
			inRawString = !inRawString
		case inRawString:
			continue
		case char == '\\' && (inString || inRune):
			i++
		case char == '"' && !inRune:
			inString = !inString
		case char == '\'' && !inString:
			inRune = !inRune
		case !inString && !inRune && char == '/' && line[i+1] == '/':
			return strings.TrimRight(line[:i], " \t")
		}
	}

	return line
}

func (p *Parser) removeComment(line string) string {
	if !simpleCommentPattern.MatchString(line) {
		return line
	}

	if !hasStringLiterals.MatchString(line) {
		return p.removeCommentFast(line)
	}

	return p.removeCommentSafe(line)
}

func (p *Parser) consolidateBlankLines(content string, lineEnding string) string {
	lines := strings.Split(content, lineEnding)

	var result []string

	prevWasBlank := false

	for _, line := range lines {
		isBlank := strings.TrimSpace(line) == ""

		if isBlank {
			if !prevWasBlank {
				result = append(result, line)
			}

			prevWasBlank = true
		} else {
			result = append(result, line)

			prevWasBlank = false
		}
	}

	return strings.Join(result, lineEnding)
}

func (p *Parser) removeCommentsFromContent(content string) (string, int) {
	inString := false
	inRune := false
	inRawString := false

	var result strings.Builder

	result.Grow(len(content))

	removalCount := 0

	lineHasContent := false

	i := 0

	for i < len(content) {
		if i >= len(content)-1 {
			if i < len(content) {
				result.WriteByte(content[i])
			}

			break
		}

		char := content[i]

		switch {
		case char == '`' && !inString && !inRune:
			inRawString = !inRawString

			result.WriteByte(char)

			lineHasContent = true
		case inRawString:
			result.WriteByte(char)

			if char != '\n' && char != '\r' {
				lineHasContent = true
			}
		case char == '\\' && (inString || inRune):
			result.WriteByte(char)

			lineHasContent = true

			if i+1 < len(content) {
				i++

				result.WriteByte(content[i])
			}
		case char == '"' && !inRune:
			inString = !inString

			result.WriteByte(char)

			lineHasContent = true
		case char == '\'' && !inString:
			inRune = !inRune

			result.WriteByte(char)

			lineHasContent = true
		case char == '\n' || char == '\r':
			result.WriteByte(char)

			lineHasContent = false
		case char == ' ' || char == '\t':
			result.WriteByte(char)
		case !inString && !inRune && char == '/' && i+1 < len(content) && content[i+1] == '/':
			trimToPos := result.Len()

			if trimToPos > 0 {
				resultStr := result.String()

				for trimToPos > 0 {
					if resultStr[trimToPos-1] == ' ' || resultStr[trimToPos-1] == '\t' {
						trimToPos--
					} else {
						break
					}
				}

				if trimToPos < result.Len() {
					trimmed := resultStr[:trimToPos]

					result.Reset()
					result.WriteString(trimmed)
				}
			}

			for i < len(content) && content[i] != '\n' && content[i] != '\r' {
				i++
			}

			if i < len(content) {
				if content[i] == '\r' && i+1 < len(content) && content[i+1] == '\n' {
					if lineHasContent {
						result.WriteByte('\r')
						result.WriteByte('\n')
					}

					i++
				} else if content[i] == '\n' || content[i] == '\r' {
					if lineHasContent {
						result.WriteByte(content[i])
					}
				}
			}

			lineHasContent = false

			removalCount++
		default:
			result.WriteByte(char)

			lineHasContent = true
		}

		i++
	}

	return result.String(), removalCount
}

func (p *Parser) postProcessRemoveTrailingNewlines(content, originalContent, lineEnding string) string {
	originalLines := strings.Split(originalContent, lineEnding)

	lastNonEmptyIndex := -1

	for i := len(originalLines) - 1; i >= 0; i-- {
		line := strings.TrimSpace(originalLines[i])

		if line != "" {
			lastNonEmptyIndex = i
			break
		}
	}

	if lastNonEmptyIndex >= 0 {
		lastLine := strings.TrimSpace(originalLines[lastNonEmptyIndex])
		if strings.HasPrefix(lastLine, "//") {
			content = strings.TrimSuffix(content, lineEnding)
		}
	}

	return content
}

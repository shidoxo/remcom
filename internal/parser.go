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
		case inRawString:
			result.WriteByte(char)
		case char == '\\' && (inString || inRune):
			result.WriteByte(char)

			if i+1 < len(content) {
				i++
				result.WriteByte(content[i])
			}
		case char == '"' && !inRune:
			inString = !inString

			result.WriteByte(char)
		case char == '\'' && !inString:
			inRune = !inRune
			result.WriteByte(char)
		case !inString && !inRune && char == '/' && i+1 < len(content) && content[i+1] == '/':
			lineStart := result.Len()

			for lineStart > 0 && result.String()[lineStart-1] != '\n' && result.String()[lineStart-1] != '\r' {
				lineStart--
			}

			lineContent := strings.TrimSpace(result.String()[lineStart:])

			hasContentBeforeComment := lineContent != ""

			for result.Len() > 0 {
				lastChar := result.String()[result.Len()-1]

				if lastChar == ' ' || lastChar == '\t' {
					str := result.String()

					result.Reset()
					result.WriteString(str[:len(str)-1])
				} else {
					break
				}
			}

			for i < len(content) && content[i] != '\n' && content[i] != '\r' {
				i++
			}

			if i < len(content) {
				if content[i] == '\r' && i+1 < len(content) && content[i+1] == '\n' {
					if hasContentBeforeComment {
						result.WriteByte('\r')
						result.WriteByte('\n')
					}

					i++
				} else if content[i] == '\n' || content[i] == '\r' {
					if hasContentBeforeComment {
						result.WriteByte(content[i])
					}
				}
			}

			removalCount++
		default:
			result.WriteByte(char)
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

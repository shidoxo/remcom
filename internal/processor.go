package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Processor struct {
	parser *Parser
}

type ProcessingResult struct {
	Content string
	Removed int
}

func NewProcessor() *Processor {
	return &Processor{
		parser: NewParser(),
	}
}

func (p *Processor) Process(content string, mode string) (*ProcessingResult, error) {
	lineEnding := p.detectLineEnding(content)

	switch mode {
	case "auto":
		return p.processAuto(content, lineEnding)
	case "manual":
		return p.processManual(content, lineEnding)
	default:
		return nil, fmt.Errorf("invalid mode: %s", mode)
	}
}

func (p *Processor) detectLineEnding(content string) string {
	if strings.Contains(content, "\r\n") {
		return "\r\n"
	} else if strings.Contains(content, "\r") {
		return "\r"
	}

	return "\n"
}

func (p *Processor) processAuto(content string, lineEnding string) (*ProcessingResult, error) {
	processed, count := p.parser.removeCommentsFromContent(content)
	processed = p.parser.postProcessRemoveTrailingNewlines(processed, content, lineEnding)

	consolidated := p.parser.consolidateBlankLines(processed, lineEnding)

	return &ProcessingResult{
		Content: consolidated,
		Removed: count,
	}, nil
}

func (p *Processor) processManual(content string, lineEnding string) (*ProcessingResult, error) {
	lines := strings.Split(content, lineEnding)

	var result strings.Builder

	result.Grow(len(content))
	removalCount := 0

	for i, line := range lines {
		modifiedLine := p.parser.removeComment(line)
		commentRemoved := line != modifiedLine

		if commentRemoved {
			printRemovalPreview(lines, i)

			remove, err := getUserInput(bufio.NewReader(os.Stdin))
			if err != nil {
				return nil, fmt.Errorf("failed to get user input: %w", err)
			}

			if remove {
				if strings.TrimSpace(modifiedLine) != "" {
					result.WriteString(modifiedLine)
					if i < len(lines)-1 || strings.HasSuffix(content, lineEnding) {
						result.WriteString(lineEnding)
					}
				}
				removalCount++
			} else {
				result.WriteString(line)
				if i < len(lines)-1 || strings.HasSuffix(content, lineEnding) {
					result.WriteString(lineEnding)
				}
			}
		} else {
			result.WriteString(line)
			if i < len(lines)-1 || strings.HasSuffix(content, lineEnding) {
				result.WriteString(lineEnding)
			}
		}
	}

	consolidated := p.parser.consolidateBlankLines(result.String(), lineEnding)

	return &ProcessingResult{
		Content: consolidated,
		Removed: removalCount,
	}, nil
}

func getUserInput(reader *bufio.Reader) (bool, error) {
	for {
		fmt.Print("Remove this line? (y/n): ")

		response, err := reader.ReadString('\n')

		if err != nil {
			return false, fmt.Errorf("error reading input: %w", err)
		}

		response = strings.TrimSpace(strings.ToLower(response))

		switch response {
		case "y", "yes":
			return true, nil
		case "n", "no":
			return false, nil
		default:
			fmt.Println("Please enter 'y' or 'n'")
			continue
		}
	}
}

func printRemovalPreview(lines []string, lineNum int) {
	isFirstLine := lineNum == 0
	isLastLine := lineNum == len(lines)-1

	var prevLineNum int
	var nextLineNum int

	if isFirstLine {
		prevLineNum = lineNum
	} else {
		prevLineNum = lineNum - 1
	}

	if isLastLine {
		nextLineNum = lineNum
	} else {
		nextLineNum = lineNum + 1
	}

	prevLine := lines[prevLineNum]
	line := lines[lineNum]
	nextLine := lines[nextLineNum]

	if prevLineNum != lineNum {
		fmt.Printf("%d: %s\n", prevLineNum+1, prevLine)
	}

	fmt.Printf("\033[41m%d: %s\033[0m\n", lineNum+1, line)

	if nextLineNum != lineNum {
		fmt.Printf("%d: %s\n", nextLineNum+1, nextLine)
	}
}

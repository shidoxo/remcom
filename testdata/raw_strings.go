package main

import "fmt"

func main() {
	// This comment should be removed

	// Raw string with comment-like content (should NOT be removed from string)
	rawStr := `This is a raw string with // comment syntax inside`

	// Multi-line raw string
	multiLine := `
        Line 1 with // comment syntax
        Line 2 with // more comment syntax
        // This looks like a comment but it's inside raw string
    ` // This actual comment should be removed

	// Complex raw string with various content
	complexRaw := `
package main

import "fmt"

// This looks like a comment but is inside raw string
func main() {
    fmt.Println("Hello") // Another fake comment
}
    ` // Real comment to be removed

	// Raw string with backticks and quotes
	withQuotes := `Raw string with "quotes" and 'single quotes' and // slashes`

	fmt.Println(rawStr)     // Print statement comment
	fmt.Println(multiLine)  // Multi-line comment
	fmt.Println(complexRaw) // Complex comment
	fmt.Println(withQuotes) // Quotes comment
}

package main

import "fmt"

func main() {
	// Basic escaped characters
	escaped1 := "Quote: \"Hello World\""             // Comment after escaped quotes
	escaped2 := "Backslash: \\ and forward slash: /" // Escaped backslash
	escaped3 := "Tab: \t and newline: \n"            // Escape sequences

	// Tricky escape scenarios
	trickyEscape := "End quote: \""        // This comment should be removed
	anotherTricky := "Escaped quote: \\\"" // And this comment too

	// Multiple escapes
	multiEscape := "Multiple: \\\\\\\"" // Lots of escapes, remove this comment

	// Raw strings with escapes (escapes don't work in raw strings)
	rawWithEscapes := `Raw string with \n and \t and \"` // Comment here

	// Single character literals with escapes
	escapedChar1 := '\'' // Single quote character
	escapedChar2 := '\\' // Backslash character
	escapedChar3 := '\n' // Newline character

	// Complex mixing
	complex := "Start \"middle // not a comment\" end" // Actual comment

	fmt.Println(escaped1, escaped2, escaped3)             // Print escaped
	fmt.Println(trickyEscape, anotherTricky, multiEscape) // Print tricky
	fmt.Println(rawWithEscapes)                           // Print raw
	fmt.Println(escapedChar1, escapedChar2, escapedChar3) // Print chars
	fmt.Println(complex)                                  // Print complex
}

package main

import "fmt"

func main() {
	// This comment should be removed
	str1 := "This // is not a comment inside string"
	str2 := "URL: https://example.com" // This comment should be removed

	// More complex string cases
	url := "https://api.example.com/v1/users"   // API endpoint comment
	comment := "Use // for line comments in Go" // Explaining comment syntax

	// Escaped quotes
	escaped := "He said \"// is for comments\" to me" // Quote explanation

	// Single quotes (runes)
	char := '/'         // Single slash character
	doubleSlash := '\\' // Backslash character, this comment should be removed

	// Mixed quotes
	mixed := "Single quote: ' and // inside" // This comment removed

	fmt.Println(str1, str2, url, comment, escaped, char, doubleSlash, mixed)
}

package main

import "fmt"

// Single slash (not a comment): /
// This is a real comment

func main() {
	// Empty string with comment
	empty := "" // Empty string comment

	// String with just slashes
	slashes := "//////" // Not comments, just slashes

	// URL-like strings
	url1 := "http://example.com"           // HTTP URL comment
	url2 := "https://api.github.com/repos" // HTTPS URL comment
	url3 := "ftp://files.example.com"      // FTP URL comment

	// File paths
	unixPath := "/usr/local/bin"        // Unix path comment
	winPath := "C:\\Program Files\\App" // Windows path comment

	// Comment-like content in strings
	notComment1 := "Use // for comments"        // Teaching comment
	notComment2 := "The URL is http://test.com" // URL in string comment
	notComment3 := "Path: /home/user"           // Path in string comment

	// Single characters
	slash := '/'      // Single slash character
	backslash := '\\' // Backslash character

	// Very long line with comment at end
	veryLongVariableName := "This is a very long string that goes on and on and on and takes up a lot of space before the comment" // Long line comment

	// Comments with special characters
	// Comment with symbols: !@#$%^&*()
	// Comment with numbers: 123456789
	// Comment with Unicode: ñáéíóú

	// Malformed-looking quotes (but valid)
	weird1 := "String with \"escaped quote at end\"" // Valid string comment
	weird2 := `Raw string with "unescaped quotes"`   // Raw string comment

	fmt.Println(empty)                                 // Print empty
	fmt.Println(slashes)                               // Print slashes
	fmt.Println(url1, url2, url3)                      // Print URLs
	fmt.Println(unixPath, winPath)                     // Print paths
	fmt.Println(notComment1, notComment2, notComment3) // Print not-comments
	fmt.Println(slash, backslash)                      // Print characters
	fmt.Println(veryLongVariableName)                  // Print long
	fmt.Println(weird1, weird2)                        // Print weird
}

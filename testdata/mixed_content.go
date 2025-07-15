package main

import (
	"fmt"     // Import comment
	"strings" // Another import comment
)

// Package level comment

// Function comment with blank line above
func processString(input string) string {

	// Local variable comment
	result := strings.TrimSpace(input) // Inline processing comment

	// Multiple blank lines above and below

	if len(result) == 0 { // Length check comment
		return "empty" // Return comment
	}

	// Processing logic comment
	processed := strings.ToLower(result) // Convert to lowercase

	return processed // Final return comment
}

// Main function comment
func main() {
	// Test cases comment
	testCases := []string{
		"Hello World", // First case comment
		"  SPACES  ",  // Second case comment
		"",            // Empty case comment
	}

	// Process each test case
	for i, test := range testCases { // Loop comment
		result := processString(test)                            // Process comment
		fmt.Printf("Case %d: '%s' -> '%s'\n", i+1, test, result) // Output comment
	}

	// End of main comment
}

// End of file comment

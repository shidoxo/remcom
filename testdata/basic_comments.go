package main

import "fmt"

// This is a package comment
func main() {
	// Simple line comment
	fmt.Println("Hello World") // End of line comment

	var x int // Variable declaration comment
	x = 42    // Assignment comment

	// Multiple consecutive comments
	// should be handled properly
	// and consolidated

	if x > 0 { // Inline condition comment
		fmt.Println("Positive") // Inside block comment
	}

	// Final comment at end
}

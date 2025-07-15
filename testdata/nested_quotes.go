package main

import "fmt"

func main() {
	// Complex nesting scenarios
	nested1 := "Outer \"inner // not a comment\" outer" // Real comment
	nested2 := `Raw "with quotes" and // fake comment`  // Another real comment

	// Multiple quote types
	mixed1 := "Double quote with 'single' inside"     // Mixed quotes comment
	mixed2 := `Raw with "double" and 'single' quotes` // Raw mixed comment

	// Tricky escape combinations
	tricky1 := "End with quote: \""      // Comment after quote
	tricky2 := "Escaped backslash: \\\\" // Comment after escapes
	tricky3 := "Escaped quote: \\\""     // Comment after escaped quote

	// Very complex nesting
	complex := "Start \"middle 'inner // fake' middle\" end" // Real comment

	// JSON-like strings
	jsonLike := `{"url": "https://api.com/endpoint", "comment": "// not real"}` // JSON comment

	// Code-like strings
	codeLike := "func main() { // this is fake code }" // Code comment

	// Edge case: quote at very end
	endQuote := "String ending with quote \"" // End quote comment

	fmt.Println(nested1, nested2)          // Print nested
	fmt.Println(mixed1, mixed2)            // Print mixed
	fmt.Println(tricky1, tricky2, tricky3) // Print tricky
	fmt.Println(complex)                   // Print complex
	fmt.Println(jsonLike)                  // Print JSON-like
	fmt.Println(codeLike)                  // Print code-like
	fmt.Println(endQuote)                  // Print end quote
}

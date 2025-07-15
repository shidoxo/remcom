package main

import "fmt"

func main() {
	// Unicode strings with comments
	greeting := "Hello 世界" // World in Chinese
	emoji := "😀🎉🚀"         // Emoji string

	// Mathematical symbols
	formula := "E = mc²" // Einstein's formula
	greek := "α β γ δ ε" // Greek letters

	// Different languages
	japanese := "こんにちは" // Hello in Japanese
	russian := "Привет" // Hello in Russian
	arabic := "مرحبا"   // Hello in Arabic

	// Unicode in raw strings
	rawUnicode := `
        Symbols: ∑ ∏ ∫ ∂
        Arrows: → ← ↑ ↓
        Currency: € £ ¥ ¢
    ` // Raw string with symbols

	// Unicode variable names (valid in Go)
	π := 3.14159 // Pi constant
	θ := 45.0    // Theta angle

	fmt.Println(greeting, emoji)           // Print greeting
	fmt.Println(formula, greek)            // Print formula
	fmt.Println(japanese, russian, arabic) // Print languages
	fmt.Println(rawUnicode)                // Print raw unicode
	fmt.Println(π, θ)                      // Print mathematical variables
}

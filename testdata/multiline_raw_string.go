package main

import "fmt"

func main() {
	// This comment should be removed

	// Multiline raw string with fake comments throughout
	sqlQuery := `
        SELECT * 
        FROM users 
        WHERE active = true  -- SQL comment style
        AND created_date > '2023-01-01'  -- Another SQL comment
        // This looks like Go comment but it's in string
        /* This looks like block comment */
    ` // This real comment should be removed

	// HTML/XML content with comments
	htmlContent := `
        <!DOCTYPE html>
        <html>
        <head>
            <!-- HTML comment inside raw string -->
            <title>Test Page</title>
            // This is not a real comment
        </head>
        <body>
            <h1>Hello World</h1> <!-- Another HTML comment -->
            <script>
                // JavaScript comment inside HTML inside raw string
                console.log("Hello");
            </script>
        </body>
        </html>
    ` // Real Go comment to be removed

	// Configuration file content
	configContent := `
        # This is a config comment (hash style)
        server.port = 8080
        server.host = localhost  // Go-style comment in config
        
        # Database settings
        db.url = postgres://localhost:5432/mydb
        db.user = admin  // Username comment
        
        // Another Go-style comment in config
        log.level = info
    ` // Configuration comment

	// Code snippet as string
	codeSnippet := `
        package example
        
        import "fmt"
        
        // This function prints hello
        func hello() {
            fmt.Println("Hello") // Print statement
        }
        
        func main() {
            hello() // Call function
        }
    ` // Code snippet comment

	fmt.Println("SQL Query:")        // Print label
	fmt.Println(sqlQuery)            // Print SQL
	fmt.Println("\nHTML Content:")   // Print label
	fmt.Println(htmlContent)         // Print HTML
	fmt.Println("\nConfig Content:") // Print label
	fmt.Println(configContent)       // Print config
	fmt.Println("\nCode Snippet:")   // Print label
	fmt.Println(codeSnippet)         // Print code
}

# remcom

A simple command-line tool for removing comments from source code files while preserving code integrity.

## 🚀 Features

- **Smart Comment Detection**: Safely removes comments while preserving string literals and raw strings
- **Multiple Modes**: Choose between automatic removal or manual confirmation for each comment
- **Atomic File Operations**: Safe file writing with temporary files to prevent data loss
- **Cross-Platform**: Works on Linux, macOS, and Windows
- **Line Ending Preservation**: Maintains original file line endings (LF, CRLF, CR)
- **Blank Line Consolidation**: Removes excessive blank lines after comment removal

## 📦 Installation

### From Source

```bash
git clone https://github.com/shidoxo/remcom.git
cd remcom
go build -o remcom cmd/remcom/main.go
```

### Using Go Install

```bash
go install github.com/shidoxo/remcom/cmd/remcom@latest
```

## 🛠️ Usage

### Basic Usage

```bash
# Remove comments automatically
remcom path/to/your/file
remcom path/to/your/file -mode auto
remcom path/to/your/file -m auto

# Remove comments with manual confirmation
remcom path/to/your/file -mode manual
remcom path/to/your/file -m manual 
```

### Command Options

```
Usage: remcom <path> [--mode MODE]

Arguments:
  <path>    Path to the file or directory

Flags:
  -h, --help           Show context-sensitive help.
  -m, --mode MODE      Mode to run in, can be 'auto' or 'manual' (default: auto)
```

## 🎯 Modes

### Auto Mode (Default)
Automatically removes all comments from the file without user intervention.

```bash
remcom path/to/your/file
```

### Manual Mode
Prompts for confirmation before removing each comment, showing context lines for better decision making.

```bash
remcom path/to/your/file -m manual
```

## 📝 Examples

### Before Processing
```go
package main

import "fmt"

// Main function starts the program
func main() {
    name := "World" // Store the name
    // Print greeting message
    fmt.Printf("Hello, %s!\n", name)
}
```

### After Processing
```go
package main

import "fmt"

func main() {
    name := "World"
    fmt.Printf("Hello, %s!\n", name)
}
```

## 🐛 Supported Comment Styles

Currently supports:
- Single-line comments (`//`)

## 🗺️ Roadmap

Planned features and improvements:
- **Hash Comments**: Support for `#` comments (Python, Shell, etc.)
- **Recursive Directory Processing**: Process all files in a directory and its subdirectories

## 📄 License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

---

**Made with ❤️ by [shidoxo](https://github.com/shidoxo)** 
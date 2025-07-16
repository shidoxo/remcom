package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/kong"
	"github.com/shidoxo/remcom/internal"
)

type Mode string

const (
	ModeAuto   Mode = "auto"
	ModeManual Mode = "manual"
)

var CLI struct {
	Path string `arg:"" name:"path" help:"Path to the file or directory" type:"path"`
	Mode Mode   `name:"mode" short:"m" help:"Mode to run in, can be 'auto' or 'manual'." default:"auto" enum:"auto,manual"`
}

func main() {
	if len(os.Args) == 1 {
		os.Args = append(os.Args, "--help")
	}

	kong.Parse(&CLI)

	fileContent, err := internal.ReadFile(CLI.Path)

	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	processor := internal.NewProcessor()

	result, err := processor.Process(fileContent.Content, string(CLI.Mode), fileContent.LineEnding)

	if err != nil {
		log.Fatalf("Failed to process file: %v", err)
	}

	err = internal.WriteFileAtomic(CLI.Path, []byte(result.Content), fileContent.Mode)

	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}

	fmt.Printf("Removed %d comments\n", result.Removed)
}

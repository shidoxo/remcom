package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type FileContent struct {
	Path       string
	Mode       os.FileMode
	Content    string
	LineEnding string
}

func ReadFile(path string) (*FileContent, error) {
	if err := validatePath(path); err != nil {
		return nil, err
	}

	fileInfo, err := os.Stat(path)

	if err != nil {
		return nil, fmt.Errorf("failed to stat file %s: %w", path, err)
	}

	if fileInfo.IsDir() {
		return nil, fmt.Errorf("path %s is a directory, not a file", path)
	}

	data, err := os.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", path, err)
	}

	content := string(data)

	lineEnding := detectLineEnding(content)

	return &FileContent{
		Path:       path,
		Mode:       fileInfo.Mode(),
		Content:    content,
		LineEnding: lineEnding,
	}, nil
}

func WriteFileAtomic(path string, content []byte, mode os.FileMode) error {
	if err := validatePath(path); err != nil {
		return err
	}

	tempPath := path + ".tmp"

	err := os.WriteFile(tempPath, content, mode)

	if err != nil {
		return fmt.Errorf("failed to write temporary file %s: %w", tempPath, err)
	}

	err = atomicRename(tempPath, path)

	if err != nil {
		os.Remove(tempPath)
		return fmt.Errorf("failed to rename %s to %s: %w", tempPath, path, err)
	}

	return nil
}

func atomicRename(tempPath, targetPath string) error {
	if runtime.GOOS == "windows" {
		if _, err := os.Stat(targetPath); err == nil {
			if err := os.Remove(targetPath); err != nil {
				return fmt.Errorf("failed to remove existing file: %w", err)
			}
		}
	}

	return os.Rename(tempPath, targetPath)
}

func validatePath(path string) error {
	if path == "" {
		return fmt.Errorf("path cannot be empty")
	}

	cleanPath := filepath.Clean(path)

	if cleanPath != path {
		return fmt.Errorf("invalid path: %s (cleaned: %s)", path, cleanPath)
	}

	if strings.Contains(path, "..") {
		return fmt.Errorf("path contains directory traversal: %s", path)
	}

	return nil
}

func detectLineEnding(content string) string {
	if strings.Contains(content, "\r\n") {
		return "\r\n"
	} else if strings.Contains(content, "\r") {
		return "\r"
	}

	return "\n"
}

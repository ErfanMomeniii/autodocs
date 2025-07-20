package project_parser

import (
	"fmt"
	"os"
	"path/filepath"
)

// File represents a file with its path.
type File struct {
	Path string
}

// Write writes newContent to the file represented by File.
func (f *File) Write(newContent []byte) error {
	return os.WriteFile(f.Path, newContent, 0644)
}

// Read reads the content of the file represented by File and returns it.
func (f *File) Read() ([]byte, error) {
	content, err := os.ReadFile(f.Path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", f.Path, err)
	}
	return content, nil
}

// Parser provides functionality to parse files in a project.
type Parser struct{}

// AllFiles returns a slice of Go source files found recursively under the given path.
func (p *Parser) AllFiles(path string) ([]File, error) {
	var files []File

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".go" {
			return nil
		}

		files = append(files, File{
			Path: path,
		})
		return nil
	})

	return files, err
}
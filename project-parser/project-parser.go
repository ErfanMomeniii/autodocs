package project_parser

import (
	"fmt"
	"os"
	"path/filepath"
)

type File struct {
	path string
}

func (f *File) Write(newContent []byte) error {
	return os.WriteFile(f.path, newContent, 0644)
}

func (f *File) Read() ([]byte, error) {
	content, err := os.ReadFile(f.path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", f.path, err)
	}
	return content, nil
}

type Parser struct{}

func (p *Parser) AllFiles(path string) ([]File, error) {
	var files []File

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		files = append(files, File{
			path: path,
		})
		return nil
	})

	return files, err
}

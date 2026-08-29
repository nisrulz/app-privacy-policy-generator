package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var generatedDirs = map[string]bool{
	"de":  true,
	"js":  true,
	"tmp": true,
}

var generatedFiles = map[string]bool{
	"index.html": true,
}

var cssGeneratedFiles = map[string]bool{
	"style.min.css":   true,
	"reviews.min.css": true,
}

func cleanPublic() error {
	entries, err := os.ReadDir("public")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("read public/ directory: %w", err)
	}

	for _, entry := range entries {
		name := entry.Name()
		path := filepath.Join("public", name)

		if generatedDirs[name] || generatedFiles[name] {
			if err := os.RemoveAll(path); err != nil {
				return fmt.Errorf("remove %s: %w", path, err)
			}
			continue
		}

		if entry.IsDir() && name == "css" {
			subEntries, err := os.ReadDir(path)
			if err != nil {
				return fmt.Errorf("read css directory: %w", err)
			}
			for _, sub := range subEntries {
				if cssGeneratedFiles[sub.Name()] {
					subPath := filepath.Join(path, sub.Name())
					if err := os.RemoveAll(subPath); err != nil {
						return fmt.Errorf("remove %s: %w", subPath, err)
					}
				}
			}
		}
	}
	return nil
}

func ensureDir(path string) error {
	return os.MkdirAll(path, 0755)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func copyFile(src, dst string) error {
	if err := ensureDir(filepath.Dir(dst)); err != nil {
		return err
	}

	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("open %s: %w", src, err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("create %s: %w", dst, err)
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return fmt.Errorf("copy %s to %s: %w", src, dst, err)
	}
	return nil
}

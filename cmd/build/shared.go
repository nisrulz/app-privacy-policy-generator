package main

import (
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
		return err
	}

	for _, entry := range entries {
		name := entry.Name()
		path := filepath.Join("public", name)

		if generatedDirs[name] || generatedFiles[name] {
			if err := os.RemoveAll(path); err != nil {
				return err
			}
			continue
		}

		if entry.IsDir() && name == "css" {
			subEntries, err := os.ReadDir(path)
			if err != nil {
				return err
			}
			for _, sub := range subEntries {
				if cssGeneratedFiles[sub.Name()] {
					if err := os.RemoveAll(filepath.Join(path, sub.Name())); err != nil {
						return err
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
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

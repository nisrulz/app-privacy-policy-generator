package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/js"
)

func buildMinifyJS() error {
	if err := ensureDir("public/js"); err != nil {
		return fmt.Errorf("create js dir: %w", err)
	}

	m := minify.New()
	m.AddFunc("text/javascript", js.Minify)

	sourcePatterns := []string{
		"src/js/*.js",
		"public/tmp/*.js",
	}
	for _, pattern := range sourcePatterns {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			return fmt.Errorf("glob %s: %w", pattern, err)
		}
		for _, src := range matches {
			name := strings.TrimSuffix(filepath.Base(src), ".js")
			dst := filepath.Join("public/js", name+".min.js")
			if err := minifySingleJS(m, src, dst); err != nil {
				return err
			}
		}
	}

	return nil
}

func minifySingleJS(m *minify.M, src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return fmt.Errorf("read %s: %w", src, err)
	}

	out, err := m.Bytes("text/javascript", data)
	if err != nil {
		return fmt.Errorf("minify %s: %w", src, err)
	}

	if err := os.WriteFile(dst, out, 0644); err != nil {
		return fmt.Errorf("write %s: %w", dst, err)
	}

	return nil
}

func minifyJS(data []byte) ([]byte, error) {
	m := minify.New()
	m.AddFunc("text/javascript", js.Minify)
	return m.Bytes("text/javascript", data)
}

package main

import (
	"fmt"
	"os"

	less "github.com/toakleaf/less.go/less"
)

func buildCSS() error {
	if err := ensureDir("public/css"); err != nil {
		return fmt.Errorf("create css dir: %w", err)
	}

	result, err := less.CompileFile("src/less/style.less", &less.CompileOptions{
		Compress: true,
	})
	if err != nil {
		return fmt.Errorf("less compile: %w", err)
	}

	if err := os.WriteFile("public/css/style.min.css", []byte(result.CSS), 0644); err != nil {
		return fmt.Errorf("write css: %w", err)
	}

	return nil
}

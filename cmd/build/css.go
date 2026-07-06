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

	lessFiles := []struct {
		src string
		dst string
	}{
		{"src/less/style.less", "public/css/style.min.css"},
		{"src/less/reviews.less", "public/css/reviews.min.css"},
	}

	for _, f := range lessFiles {
		result, err := less.CompileFile(f.src, &less.CompileOptions{
			Compress: true,
		})
		if err != nil {
			return fmt.Errorf("less compile %s: %w", f.src, err)
		}
		if err := os.WriteFile(f.dst, []byte(result.CSS), 0644); err != nil {
			return fmt.Errorf("write %s: %w", f.dst, err)
		}
	}

	return nil
}

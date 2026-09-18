package main

import (
	"fmt"
	"path/filepath"
)

func copyVendorAssets() error {
	if err := ensureDir("public/js/vendor"); err != nil {
		return fmt.Errorf("create vendor js directory: %w", err)
	}

	matches, err := filepath.Glob("src/includes/vendor/*")
	if err != nil {
		return fmt.Errorf("glob vendor assets: %w", err)
	}
	for _, src := range matches {
		dst := filepath.Join("public/js/vendor", filepath.Base(src))
		if err := copyFile(src, dst); err != nil {
			return fmt.Errorf("copy vendor asset %s: %w", src, err)
		}
	}

	return nil
}

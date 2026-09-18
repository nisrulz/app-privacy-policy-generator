package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func copyVendorAssets() error {
	types, err := os.ReadDir("src/includes/vendor")
	if err != nil {
		return fmt.Errorf("read vendor dir: %w", err)
	}

	for _, t := range types {
		if !t.IsDir() {
			continue
		}

		dstDir := filepath.Join("public", t.Name(), "vendor")
		if err := ensureDir(dstDir); err != nil {
			return fmt.Errorf("create %s: %w", dstDir, err)
		}

		matches, err := filepath.Glob(filepath.Join("src/includes/vendor", t.Name(), "*"))
		if err != nil {
			return fmt.Errorf("glob vendor %s: %w", t.Name(), err)
		}
		for _, src := range matches {
			dst := filepath.Join(dstDir, filepath.Base(src))
			if err := copyFile(src, dst); err != nil {
				return fmt.Errorf("copy vendor asset %s: %w", src, err)
			}
		}
	}

	return nil
}

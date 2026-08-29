package main

import "fmt"

func copyVendorAssets() error {
	if err := ensureDir("public/js/vendor"); err != nil {
		return fmt.Errorf("create vendor js directory: %w", err)
	}
	if err := ensureDir("public/images/vendor"); err != nil {
		return fmt.Errorf("create vendor images directory: %w", err)
	}

	vendorJS := map[string]string{
		"src/includes/vendor/vue.global.prod.js": "public/js/vendor/vue.global.prod.js",
		"src/includes/vendor/to-markdown.min.js": "public/js/vendor/to-markdown.min.js",
	}
	for src, dst := range vendorJS {
		if fileExists(src) {
			if err := copyFile(src, dst); err != nil {
				return fmt.Errorf("copy vendor asset %s: %w", src, err)
			}
		}
	}

	return nil
}

package main

func copyVendorAssets() error {
	if err := ensureDir("public/js/vendor"); err != nil {
		return err
	}
	if err := ensureDir("public/images/vendor"); err != nil {
		return err
	}

	vendorJS := map[string]string{
		"src/includes/vendor/vue.global.prod.js": "public/js/vendor/vue.global.prod.js",
		"src/includes/vendor/to-markdown.min.js": "public/js/vendor/to-markdown.min.js",
	}
	for src, dst := range vendorJS {
		if fileExists(src) {
			if err := copyFile(src, dst); err != nil {
				return err
			}
		}
	}

	return nil
}

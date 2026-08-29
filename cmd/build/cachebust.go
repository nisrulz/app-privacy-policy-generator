package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func cacheBust() error {
	sharedFiles := []string{
		"public/css/style.min.css",
		"public/css/reviews.min.css",
		"public/css/vendor/normalize.min.css",
		"public/css/vendor/bulma.min.css",
		"public/js/main.min.js",
		"public/js/locales.min.js",
		"public/js/utils.min.js",
		"public/js/thirdpartyservices.min.js",
		"public/js/flycricket.min.js",
		"public/js/reviewThemeToggle.min.js",
		"public/js/reviewsMain.min.js",
		"public/js/vendor/vue.global.prod.js",
		"public/js/vendor/to-markdown.min.js",
	}

	hashMap := make(map[string]string)
	for _, f := range sharedFiles {
		if !fileExists(f) {
			continue
		}
		hash, err := computeHash(f)
		if err != nil {
			continue
		}
		base := filepath.Base(f)
		hashMap[base] = fmt.Sprintf("%s?v=%s", base, hash)
	}

	err := filepath.Walk("public", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("walk %s: %w", path, err)
		}
		if info.IsDir() || !strings.HasSuffix(path, ".html") {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read %s: %w", path, err)
		}

		html := string(content)
		for old, new := range hashMap {
			html = strings.ReplaceAll(html, old+"\"", new+"\"")
		}

		if err := os.WriteFile(path, []byte(html), 0644); err != nil {
			return fmt.Errorf("write %s: %w", path, err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("cache-bust HTML files: %w", err)
	}

	langs, err := getLocales()
	if err != nil {
		return fmt.Errorf("determine locales: %w", err)
	}

	for _, lang := range langs {
		outDir := "public"
		if lang != "en" {
			outDir = fmt.Sprintf("public/%s", lang)
		}
		f := filepath.Join(outDir, "js", "locale.min.js")
		if !fileExists(f) {
			continue
		}
		hash, err := computeHash(f)
		if err != nil {
			return fmt.Errorf("hash locale js for %s: %w", lang, err)
		}
		htmlPath := filepath.Join(outDir, "index.html")
		if !fileExists(htmlPath) {
			continue
		}
		content, err := os.ReadFile(htmlPath)
		if err != nil {
			return fmt.Errorf("read %s: %w", htmlPath, err)
		}
		html := strings.ReplaceAll(string(content), "locale.min.js\"", "locale.min.js?v="+hash+"\"")
		if err := os.WriteFile(htmlPath, []byte(html), 0644); err != nil {
			return fmt.Errorf("write %s: %w", htmlPath, err)
		}
	}

	return nil
}

func computeHash(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read %s: %w", path, err)
	}

	hash := md5.Sum(data)
	return fmt.Sprintf("%x", hash[:4]), nil
}

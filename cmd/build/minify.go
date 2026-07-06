package main

import (
	"fmt"
	"os"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/js"
)

func buildMinifyJS() error {
	if err := ensureDir("public/js"); err != nil {
		return fmt.Errorf("create js dir: %w", err)
	}

	m := minify.New()
	m.AddFunc("text/javascript", js.Minify)

	jsFiles := []struct {
		src string
		dst string
	}{
		{"src/js/utils.js", "public/js/utils.min.js"},
		{"public/tmp/thirdpartyservices.js", "public/js/thirdpartyservices.min.js"},
		{"src/js/flycricket.js", "public/js/flycricket.min.js"},
	}

	for _, f := range jsFiles {
		if err := minifySingleJS(m, f.src, f.dst); err != nil {
			return err
		}
	}

	if err := buildMainJS(m); err != nil {
		return err
	}

	return nil
}

func buildMainJS(m *minify.M) error {
	files := []string{
		"src/js/wizardMixin.js",
		"src/js/platformMixin.js",
		"src/js/localeMixin.js",
		"src/js/formDataMixin.js",
		"src/js/generatorMixin.js",
		"src/js/modalMixin.js",
		"src/js/thirdPartyMixin.js",
		"src/js/contentMixin.js",
		"src/js/main.js",
	}

	var combined []byte
	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			return fmt.Errorf("read %s: %w", f, err)
		}
		combined = append(combined, data...)
		combined = append(combined, '\n')
	}

	out, err := m.Bytes("text/javascript", combined)
	if err != nil {
		return fmt.Errorf("minify main.js: %w", err)
	}

	if err := os.WriteFile("public/js/main.min.js", out, 0644); err != nil {
		return fmt.Errorf("write main.min.js: %w", err)
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

package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type pageData struct {
	Lang        string
	Flycricket  bool
	NoTracking  bool
	Gdpr        bool
	ThemeToggle bool
}

func dict(pairs ...any) (map[string]any, error) {
	if len(pairs)%2 != 0 {
		return nil, fmt.Errorf("dict expects key/value pairs")
	}
	m := make(map[string]any, len(pairs)/2)
	for i := 0; i < len(pairs); i += 2 {
		k, ok := pairs[i].(string)
		if !ok {
			return nil, fmt.Errorf("dict keys must be strings")
		}
		m[k] = pairs[i+1]
	}
	return m, nil
}

func parsePageTemplates() (*template.Template, error) {
	fsys := os.DirFS("src/tpl")
	var files []string
	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(path, ".html") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("walk src/tpl: %w", err)
	}

	tmpl := template.New("page.html").Funcs(template.FuncMap{"dict": dict})
	tmpl, err = tmpl.ParseFS(fsys, files...)
	if err != nil {
		return nil, fmt.Errorf("parse templates: %w", err)
	}
	return tmpl, nil
}

func renderHTML(lang, outDir string) error {
	tmpl, err := parsePageTemplates()
	if err != nil {
		return err
	}

	data := pageData{
		Lang:        lang,
		Flycricket:  true,
		NoTracking:  true,
		Gdpr:        true,
		ThemeToggle: true,
	}

	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "page.html", data); err != nil {
		return fmt.Errorf("execute templates for %s: %w", lang, err)
	}

	outPath := filepath.Join(outDir, "index.html")
	if err := os.WriteFile(outPath, buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("write %s: %w", outPath, err)
	}
	return nil
}

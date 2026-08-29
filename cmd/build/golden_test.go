package main

import (
	"bytes"
	"flag"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

var whitespaceRe = regexp.MustCompile(`\s+`)
var interTagRe = regexp.MustCompile(`>\s+<`)
var cacheBustRe = regexp.MustCompile(`\?v=[a-f0-9]+`)

func normalizeHTML(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	return normalize(data)
}

func normalize(data []byte) string {
	s := cacheBustRe.ReplaceAll(data, nil)
	s = interTagRe.ReplaceAll(s, []byte("><"))
	s = whitespaceRe.ReplaceAll(s, []byte(" "))
	return strings.TrimSpace(string(s))
}

var updateGolden = flag.Bool("update", false, "regenerate golden test fixtures")

func TestGoldenParity(t *testing.T) {
	goldenDir := "testdata/golden"
	if _, err := os.Stat("src/tpl"); err != nil {
		if err := os.Chdir("../.."); err != nil {
			t.Fatalf("chdir to repo root: %v", err)
		}
		goldenDir = "cmd/build/testdata/golden"
	}

	for _, lang := range []string{"en", "de"} {
		lang := lang
		t.Run(lang, func(t *testing.T) {
			tmpl, err := parsePageTemplates()
			if err != nil {
				t.Fatalf("parse templates: %v", err)
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
				t.Fatalf("execute: %v", err)
			}

			if *updateGolden {
				if err := os.WriteFile(filepath.Join(goldenDir, lang+".html"), buf.Bytes(), 0644); err != nil {
					t.Fatalf("write golden %s: %v", lang, err)
				}
				return
			}

			got := normalize(buf.Bytes())
			want := normalizeHTML(t, filepath.Join(goldenDir, lang+".html"))

			if got != want {
				i := 0
				for i < len(got) && i < len(want) && got[i] == want[i] {
					i++
				}
				start := max(0, i-120)
				endGot, endWant := min(len(got), i+180), min(len(want), i+180)
				t.Errorf("output mismatch for %s at byte %d\nGOT:  ...%s\nWANT: ...%s", lang, i,
					got[start:endGot], want[start:endWant])
			}
		})
	}
}

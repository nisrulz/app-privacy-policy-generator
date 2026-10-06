package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestComputeHashIsMD5TruncatedToFourBytes(t *testing.T) {
	path := filepath.Join(t.TempDir(), "asset.js")
	if err := os.WriteFile(path, []byte("hello"), 0644); err != nil {
		t.Fatal(err)
	}

	hash, err := computeHash(path)
	if err != nil {
		t.Fatalf("computeHash: %v", err)
	}

	want := "5d41402a"
	if hash != want {
		t.Errorf("computeHash = %q, want %q (first 4 bytes of md5(hello))", hash, want)
	}
}

func TestComputeHashMissingFileReturnsError(t *testing.T) {
	if _, err := computeHash(filepath.Join(t.TempDir(), "missing.js")); err == nil {
		t.Error("computeHash on a missing file should return an error")
	}
}

func TestCacheBust(t *testing.T) {
	t.Chdir(t.TempDir())

	write := func(path, content string) {
		t.Helper()
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}

	write("public/css/style.min.css", "body{}")
	write("public/js/main.min.js", "console.log(1)")
	write("public/js/locale.min.js", "window.__locale = {greeting:\"Hello\"}")
	write("public/de/js/locale.min.js", "window.__locale = {greeting:\"Hallo\"}")
	write("src/locales/en.json", "{}")
	write("src/locales/de.json", "{}")
	write("public/index.html", `<link href="css/style.min.css"><script src="js/main.min.js"></script><script src="js/locale.min.js"></script>`)
	write("public/de/index.html", `<script src="js/locale.min.js"></script>`)

	if err := cacheBust(); err != nil {
		t.Fatalf("cacheBust: %v", err)
	}

	read := func(path string) string {
		t.Helper()
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read %s: %v", path, err)
		}
		return string(data)
	}

	cssHash, err := computeHash("public/css/style.min.css")
	if err != nil {
		t.Fatal(err)
	}
	jsHash, err := computeHash("public/js/main.min.js")
	if err != nil {
		t.Fatal(err)
	}
	enHash, err := computeHash("public/js/locale.min.js")
	if err != nil {
		t.Fatal(err)
	}
	deHash, err := computeHash("public/de/js/locale.min.js")
	if err != nil {
		t.Fatal(err)
	}

	enHTML := read("public/index.html")
	if !strings.Contains(enHTML, "css/style.min.css?v="+cssHash) {
		t.Errorf("index.html should reference the hashed stylesheet, got %s", enHTML)
	}
	if !strings.Contains(enHTML, "js/main.min.js?v="+jsHash) {
		t.Errorf("index.html should reference the hashed main script, got %s", enHTML)
	}
	if !strings.Contains(enHTML, "js/locale.min.js?v="+enHash) {
		t.Errorf("index.html should reference the en locale hash %q, got %s", enHash, enHTML)
	}
	if strings.Contains(enHTML, "js/locale.min.js?v="+jsHash) {
		t.Error("locale.min.js must be skipped by the general pass and not receive the main script hash")
	}

	deHTML := read("public/de/index.html")
	if !strings.Contains(deHTML, "js/locale.min.js?v="+deHash) {
		t.Errorf("de/index.html should reference the de locale hash %q, got %s", deHash, deHTML)
	}
	if strings.Contains(deHTML, "js/locale.min.js?v="+enHash) {
		t.Error("de/index.html must not receive the en locale hash")
	}
}

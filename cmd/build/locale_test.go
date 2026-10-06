package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeFixture(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
}

func TestGetLocalesListsOnlyTopLevelJSONFiles(t *testing.T) {
	t.Chdir(t.TempDir())

	writeFixture(t, "src/locales/en.json", "{}")
	writeFixture(t, "src/locales/de.json", "{}")
	writeFixture(t, "src/locales/notes.txt", "not a locale")
	if err := os.MkdirAll("src/locales/partial", 0755); err != nil {
		t.Fatal(err)
	}

	langs, err := getLocales()
	if err != nil {
		t.Fatalf("getLocales: %v", err)
	}

	want := []string{"de", "en"}
	if len(langs) != len(want) {
		t.Fatalf("getLocales = %v, want %v", langs, want)
	}
	for i, lang := range langs {
		if lang != want[i] {
			t.Errorf("getLocales[%d] = %q, want %q", i, lang, want[i])
		}
	}
}

func TestLoadLocale(t *testing.T) {
	t.Chdir(t.TempDir())

	writeFixture(t, "src/locales/en.json", `{"_flag":"x","greeting":"Hello"}`)

	locale, err := loadLocale("en")
	if err != nil {
		t.Fatalf("loadLocale: %v", err)
	}
	if locale["greeting"] != "Hello" || locale["_flag"] != "x" {
		t.Errorf("loadLocale = %v, want the parsed key/value pairs", locale)
	}

	if _, err := loadLocale("missing"); err == nil {
		t.Error("loadLocale on a missing file should return an error")
	}

	writeFixture(t, "src/locales/broken.json", "{not json")
	if _, err := loadLocale("broken"); err == nil {
		t.Error("loadLocale on invalid JSON should return an error")
	}
}

func TestBuildLocalesRegistry(t *testing.T) {
	t.Chdir(t.TempDir())

	writeFixture(t, "src/locales/en.json", `{"_flag":"🇬🇧","greeting":"Hello"}`)
	writeFixture(t, "src/locales/de.json", `{"_flag":"🇩🇪","greeting":"Hallo"}`)
	writeFixture(t, "src/locales/broken.json", "{not json")
	if err := os.MkdirAll("public/tmp", 0755); err != nil {
		t.Fatal(err)
	}

	if err := buildLocalesRegistry(); err != nil {
		t.Fatalf("buildLocalesRegistry: %v", err)
	}

	data, err := os.ReadFile("public/tmp/locales.js")
	if err != nil {
		t.Fatalf("read registry output: %v", err)
	}

	js := string(data)
	const prefix = "var availableLocalesJsonArray = "
	if !strings.HasPrefix(js, prefix) {
		t.Fatalf("registry output should start with %q, got %s", prefix, js)
	}

	var entries []struct {
		Code string `json:"code"`
		Flag string `json:"flag"`
	}
	if err := json.Unmarshal([]byte(strings.TrimPrefix(js, prefix)), &entries); err != nil {
		t.Fatalf("parse registry JSON: %v", err)
	}

	want := map[string]string{"en": "🇬🇧", "de": "🇩🇪"}
	if len(entries) != len(want) {
		t.Fatalf("registry has %d entries, want %d (broken locales must be skipped): %v", len(entries), len(want), entries)
	}
	for _, entry := range entries {
		if want[entry.Code] != entry.Flag {
			t.Errorf("registry entry %q has flag %q, want %q", entry.Code, entry.Flag, want[entry.Code])
		}
	}
}

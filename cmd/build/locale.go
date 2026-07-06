package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type LocaleData map[string]string

func loadLocale(lang string) (LocaleData, error) {
	path := fmt.Sprintf("src/locales/%s.json", lang)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read locale %s: %w", lang, err)
	}

	var locale LocaleData
	if err := json.Unmarshal(data, &locale); err != nil {
		return nil, fmt.Errorf("parse locale %s: %w", lang, err)
	}

	return locale, nil
}

func getLocales() ([]string, error) {
	entries, err := os.ReadDir("src/locales")
	if err != nil {
		return nil, err
	}

	var langs []string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".json") {
			lang := strings.TrimSuffix(entry.Name(), ".json")
			langs = append(langs, lang)
		}
	}
	return langs, nil
}

func buildLocalesRegistry() error {
	if err := ensureDir("public/js"); err != nil {
		return err
	}

	entries, err := os.ReadDir("src/locales")
	if err != nil {
		return fmt.Errorf("read locales dir: %w", err)
	}

	type LocaleEntry struct {
		Code string `json:"code"`
		Flag string `json:"flag"`
	}

	var localeEntries []LocaleEntry
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".json") {
			code := strings.TrimSuffix(entry.Name(), ".json")
			data, err := os.ReadFile(filepath.Join("src/locales", entry.Name()))
			if err != nil {
				continue
			}
			var locale map[string]interface{}
			if err := json.Unmarshal(data, &locale); err != nil {
				continue
			}
			flag := ""
			if f, ok := locale["_flag"].(string); ok {
				flag = f
			}
			localeEntries = append(localeEntries, LocaleEntry{Code: code, Flag: flag})
		}
	}

	jsonData, err := json.Marshal(localeEntries)
	if err != nil {
		return fmt.Errorf("marshal locales: %w", err)
	}

	js := fmt.Sprintf("var availableLocalesJsonArray = %s", jsonData)
	if err := os.WriteFile("public/tmp/locales.js", []byte(js), 0644); err != nil {
		return fmt.Errorf("write locales js: %w", err)
	}

	minified, err := minifyJS([]byte(js))
	if err != nil {
		return fmt.Errorf("minify locales: %w", err)
	}

	if err := os.WriteFile("public/js/locales.min.js", minified, 0644); err != nil {
		return fmt.Errorf("write minified locales: %w", err)
	}

	return nil
}

func buildLocales(langs []string) error {
	for _, lang := range langs {
		lang = strings.TrimSpace(lang)
		if err := buildSingleLocale(lang); err != nil {
			return fmt.Errorf("build locale %s: %w", lang, err)
		}
	}
	return nil
}

func buildLocaleJS(lang string) error {
	outDir := "public"
	if lang != "en" {
		outDir = fmt.Sprintf("public/%s", lang)
	}

	if err := ensureDir(outDir); err != nil {
		return err
	}

	locale, err := loadLocale(lang)
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(locale)
	if err != nil {
		return fmt.Errorf("marshal locale %s: %w", lang, err)
	}

	localeJS := fmt.Sprintf("window.__locale = %s", jsonData)

	if err := ensureDir(filepath.Join(outDir, "js")); err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(outDir, "js", "locale.min.js"), []byte(localeJS), 0644)
}

func buildSingleLocale(lang string) error {
	outDir := "public"
	if lang != "en" {
		outDir = fmt.Sprintf("public/%s", lang)
	}

	if err := ensureDir(outDir); err != nil {
		return err
	}

	locale, err := loadLocale(lang)
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(locale)
	if err != nil {
		return fmt.Errorf("marshal locale %s: %w", lang, err)
	}

	localeJS := fmt.Sprintf("window.__locale = %s", jsonData)

	if err := ensureDir(filepath.Join(outDir, "js")); err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(outDir, "js", "locale.min.js"), []byte(localeJS), 0644); err != nil {
		return fmt.Errorf("write locale js: %w", err)
	}

	if err := renderPug(lang, outDir); err != nil {
		return err
	}

	return nil
}

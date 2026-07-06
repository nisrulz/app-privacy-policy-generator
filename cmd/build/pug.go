package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func renderPug(lang string, outDir string) error {
	if err := ensureDir(outDir); err != nil {
		return err
	}

	if lang == "en" {
		cmd := exec.Command("npx", "-q", "pug3", "src/index.pug", "--out", outDir, "--silent")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("pug3 render failed: %w", err)
		}
		return nil
	}

	return renderPugWithLang(lang, outDir)
}

func renderPugWithLang(lang string, outDir string) error {
	pugPath := "src/index.pug"

	content, err := os.ReadFile(pugPath)
	if err != nil {
		return fmt.Errorf("read pug file: %w", err)
	}

	original := string(content)
	modified := strings.Replace(original, "- var lang = 'en'", "- var lang = '"+lang+"'", 1)

	if err := os.WriteFile(pugPath, []byte(modified), 0644); err != nil {
		return fmt.Errorf("write temp pug file: %w", err)
	}

	defer func() {
		os.WriteFile(pugPath, []byte(original), 0644)
	}()

	cmd := exec.Command("npx", "-q", "pug3", pugPath, "--out", outDir, "--silent")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("pug3 render failed for %s: %w", lang, err)
	}

	return nil
}

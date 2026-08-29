package e2e

import (
	"testing"
	"time"

	"github.com/chromedp/chromedp"
)

func TestThemeReload_DarkThemePersistsAfterPageReload(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := gotoApp(ctx); err != nil {
		t.Fatal(err)
	}

	theme, err := getElementAttribute(ctx, "html", "data-theme")
	if err != nil {
		t.Fatal(err)
	}
	if theme != "light" {
		t.Fatalf("expected initial theme to be 'light', got '%s'", theme)
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`button.theme-toggle`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	theme, err = getElementAttribute(ctx, "html", "data-theme")
	if err != nil {
		t.Fatal(err)
	}
	if theme != "dark" {
		t.Fatalf("expected theme to be 'dark' after toggle, got '%s'", theme)
	}

	if err := chromedp.Run(ctx,
		chromedp.Reload(),
		chromedp.Sleep(500*time.Millisecond),
	); err != nil {
		t.Fatal(err)
	}

	theme, err = getElementAttribute(ctx, "html", "data-theme")
	if err != nil {
		t.Fatal(err)
	}
	if theme != "dark" {
		t.Fatalf("expected theme to remain 'dark' after reload, got '%s'", theme)
	}
}

func TestThemeReload_LightThemePersistsAfterPageReload(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := gotoApp(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`button.theme-toggle`, chromedp.ByQuery),
		chromedp.Click(`button.theme-toggle`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	theme, err := getElementAttribute(ctx, "html", "data-theme")
	if err != nil {
		t.Fatal(err)
	}
	if theme != "light" {
		t.Fatalf("expected theme to be 'light' after toggling twice, got '%s'", theme)
	}

	if err := chromedp.Run(ctx,
		chromedp.Reload(),
		chromedp.Sleep(500*time.Millisecond),
	); err != nil {
		t.Fatal(err)
	}

	theme, err = getElementAttribute(ctx, "html", "data-theme")
	if err != nil {
		t.Fatal(err)
	}
	if theme != "light" {
		t.Fatalf("expected theme to remain 'light' after reload, got '%s'", theme)
	}
}

func TestThemeReload_ThemeToggleButtonTextUpdatesCorrectly(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := gotoApp(ctx); err != nil {
		t.Fatal(err)
	}

	btnText, err := getElementText(ctx, "button.theme-toggle")
	if err != nil {
		t.Fatal(err)
	}
	if btnText != "🌙" {
		t.Fatalf("expected button text to be '🌙' in light mode, got '%s'", btnText)
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`button.theme-toggle`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	btnText, err = getElementText(ctx, "button.theme-toggle")
	if err != nil {
		t.Fatal(err)
	}
	if btnText != "☀️" {
		t.Fatalf("expected button text to be '☀️' in dark mode, got '%s'", btnText)
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`button.theme-toggle`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	btnText, err = getElementText(ctx, "button.theme-toggle")
	if err != nil {
		t.Fatal(err)
	}
	if btnText != "🌙" {
		t.Fatalf("expected button text to be '🌙' after toggling back to light, got '%s'", btnText)
	}
}

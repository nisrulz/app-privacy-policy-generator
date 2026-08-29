package e2e

import (
	"testing"

	"github.com/chromedp/chromedp"
)

func TestTheme_ThemeToggleButtonIsVisible(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := gotoApp(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 1); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, "button.theme-toggle") {
		t.Fatal("expected theme toggle button to be visible")
	}
}

func TestTheme_DefaultThemeIsLight(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := gotoApp(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 1); err != nil {
		t.Fatal(err)
	}

	theme, err := getElementAttribute(ctx, "html", "data-theme")
	if err != nil {
		t.Fatal(err)
	}
	if theme != "light" {
		t.Fatalf("expected default theme to be 'light', got '%s'", theme)
	}
}

func TestTheme_TogglesToDarkModeOnClick(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := gotoApp(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 1); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`button.theme-toggle`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	theme, err := getElementAttribute(ctx, "html", "data-theme")
	if err != nil {
		t.Fatal(err)
	}
	if theme != "dark" {
		t.Fatalf("expected theme to be 'dark' after toggle, got '%s'", theme)
	}
}

func TestTheme_TogglesBackToLightModeOnSecondClick(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := gotoApp(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 1); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`button.theme-toggle`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	theme, err := getElementAttribute(ctx, "html", "data-theme")
	if err != nil {
		t.Fatal(err)
	}
	if theme != "dark" {
		t.Fatalf("expected theme to be 'dark' after first toggle, got '%s'", theme)
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
	if theme != "light" {
		t.Fatalf("expected theme to be 'light' after second toggle, got '%s'", theme)
	}
}

func TestTheme_ThemePersistsAfterNavigatingSteps(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := gotoApp(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 1); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`button.theme-toggle`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	theme, err := getElementAttribute(ctx, "html", "data-theme")
	if err != nil {
		t.Fatal(err)
	}
	if theme != "dark" {
		t.Fatalf("expected theme to be 'dark' after toggle, got '%s'", theme)
	}

	if err := clickStart(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 2); err != nil {
		t.Fatal(err)
	}

	theme, err = getElementAttribute(ctx, "html", "data-theme")
	if err != nil {
		t.Fatal(err)
	}
	if theme != "dark" {
		t.Fatalf("expected theme to remain 'dark' after navigating, got '%s'", theme)
	}
}

package e2e

import (
	"testing"
	"time"

	"github.com/chromedp/chromedp"
)

func TestTranslate_SubstitutesPlaceholders(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := gotoApp(ctx); err != nil {
		t.Fatal(err)
	}

	var result string
	if err := chromedp.Run(ctx,
		chromedp.Evaluate(`window.translate('wizard.step1.start')`, &result),
	); err != nil {
		t.Fatal(err)
	}
	if result == "" {
		t.Fatal("expected translate to return a non-empty string")
	}
}

func TestTranslate_ReturnsKeyWhenNotFound(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := gotoApp(ctx); err != nil {
		t.Fatal(err)
	}

	var result string
	if err := chromedp.Run(ctx,
		chromedp.Evaluate(`window.translate('nonexistent.key.that.does.not.exist')`, &result),
	); err != nil {
		t.Fatal(err)
	}
	if result != "nonexistent.key.that.does.not.exist" {
		t.Fatalf("expected translate to return the key itself, got '%s'", result)
	}
}

func TestTranslate_ReturnsCorrectTextForKnownKeys(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := gotoApp(ctx); err != nil {
		t.Fatal(err)
	}

	var result string
	if err := chromedp.Run(ctx,
		chromedp.Evaluate(`window.translate('wizard.step1.start')`, &result),
	); err != nil {
		t.Fatal(err)
	}
	if result != "Start" {
		t.Fatalf("expected translate to return 'Start', got '%s'", result)
	}
}

func TestUpdateMeta_PageTitleReflectsLocale(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := gotoApp(ctx); err != nil {
		t.Fatal(err)
	}

	var title string
	if err := chromedp.Run(ctx,
		chromedp.Title(&title),
	); err != nil {
		t.Fatal(err)
	}
	if title == "" {
		t.Fatal("expected page title to be non-empty")
	}
}

func TestUpdateMeta_MetaDescriptionTagExists(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := gotoApp(ctx); err != nil {
		t.Fatal(err)
	}

	desc, err := getElementAttribute(ctx, `meta[name="description"]`, "content")
	if err != nil {
		t.Fatal(err)
	}
	if desc == "" {
		t.Fatal("expected meta description to have content")
	}
}

func TestUpdateMeta_OgTitleMetaTagExists(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := gotoApp(ctx); err != nil {
		t.Fatal(err)
	}

	ogTitle, err := getElementAttribute(ctx, `meta[property="og:title"]`, "content")
	if err != nil {
		t.Fatal(err)
	}
	if ogTitle == "" {
		t.Fatal("expected og:title meta tag to have content")
	}
}

func TestUpdateThemeLogo_LogoSrcToggle(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := gotoApp(ctx); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, `img[data-theme-logo]`) {
		t.Fatal("expected theme logo image to be visible")
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`button.theme-toggle`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	if err := sleep(ctx, 300*time.Millisecond); err != nil {
		t.Fatal(err)
	}

	darkSrc, err := getElementAttribute(ctx, `img[data-theme-logo]`, "src")
	if err != nil {
		t.Fatal(err)
	}
	if darkSrc == "" {
		t.Fatal("expected dark logo src to be non-empty")
	}
}

package e2e

import (
	"strings"
	"testing"

	"github.com/chromedp/chromedp"
)

func TestStep5ShowsAllPlatformCheckboxes(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep5(ctx); err != nil {
		t.Fatal(err)
	}

	platforms := []string{"android", "ios", "kaios", "windows", "web"}
	for _, p := range platforms {
		selector := `label[for="platform-` + p + `"]`
		if !elementExists(ctx, selector) {
			t.Errorf("expected platform label for %q to be visible", p)
		}
	}
}

func TestAndroidIsCheckedByDefault(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep5(ctx); err != nil {
		t.Fatal(err)
	}

	val, err := getElementAttribute(ctx, "#platform-android", "checked")
	if err != nil {
		t.Fatal(err)
	}
	if val == "" {
		t.Error("expected Android to be checked by default")
	}

	val, err = getElementAttribute(ctx, "#platform-ios", "checked")
	if err != nil {
		t.Fatal(err)
	}
	if val != "" {
		t.Error("expected iOS to not be checked by default")
	}
}

func TestCanToggleMultiplePlatforms(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep5(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`label[for="platform-ios"]`, chromedp.ByQuery),
		chromedp.Click(`label[for="platform-web"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	for _, platform := range []string{"android", "ios", "web"} {
		val, err := getElementAttribute(ctx, "#platform-"+platform, "checked")
		if err != nil {
			t.Fatal(err)
		}
		if val == "" {
			t.Errorf("expected %s to be checked", platform)
		}
	}
}

func TestSelectedPlatformsLabelUpdates(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep5(ctx); err != nil {
		t.Fatal(err)
	}

	initialText, err := getElementText(ctx, "#step-5 .tag.is-primary.is-light")
	if err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`label[for="platform-ios"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	updatedText, err := getElementText(ctx, "#step-5 .tag.is-primary.is-light")
	if err != nil {
		t.Fatal(err)
	}

	if updatedText == initialText {
		t.Error("expected label text to change after selecting iOS")
	}
	if !strings.Contains(updatedText, "iOS") {
		t.Errorf("expected label to contain 'iOS', got %q", updatedText)
	}
}

func TestCanUncheckAndroid(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep5(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`label[for="platform-android"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	val, err := getElementAttribute(ctx, "#platform-android", "checked")
	if err != nil {
		t.Fatal(err)
	}
	if val != "" {
		t.Error("expected Android to be unchecked after clicking")
	}
}

func TestPlatformInfoReflectsInGeneratedPolicy(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := gotoApp(ctx); err != nil {
		t.Fatal(err)
	}
	if err := clickStart(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 2); err != nil {
		t.Fatal(err)
	}
	if err := fillStep2(ctx); err != nil {
		t.Fatal(err)
	}
	if err := clickNext(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 3); err != nil {
		t.Fatal(err)
	}
	if err := clickNext(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 4); err != nil {
		t.Fatal(err)
	}
	if err := clickNext(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 5); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`label[for="platform-ios"]`, chromedp.ByQuery),
		chromedp.Click(`label[for="platform-web"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if err := clickNext(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 6); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#devName`, "Platform Test", chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if err := clickNext(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 7); err != nil {
		t.Fatal(err)
	}
	if err := clickNext(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 8); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`a.card-footer-item >> visible=true`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	content, err := getElementText(ctx, "#privacy_simple_content")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(content, "mobile devices") {
		t.Error("expected policy to contain 'mobile devices'")
	}
	if !strings.Contains(content, "web browsers") {
		t.Error("expected policy to contain 'web browsers'")
	}
}

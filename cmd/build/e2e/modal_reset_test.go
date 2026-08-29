package e2e

import (
	"testing"
	"time"

	"github.com/chromedp/chromedp"
)

func TestModalReset_ReopeningPrivacyModalResetsFromHTMLView(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizard(ctx); err != nil {
		t.Fatal(err)
	}

	if err := clickStep8Button(ctx, "Privacy Policy"); err != nil {
		t.Fatal(err)
	}
	if err := sleep(ctx, 500*time.Millisecond); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, ".modal.is-active") {
		t.Fatal("expected modal to be active")
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`.modal.is-active button[onclick*="getHtml"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if err := sleep(ctx, 300*time.Millisecond); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, "#privacy_simple_txtarea") {
		t.Fatal("expected textarea to be visible in HTML view")
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`.modal.is-active button.delete`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if err := sleep(ctx, 300*time.Millisecond); err != nil {
		t.Fatal(err)
	}

	if elementExists(ctx, ".modal.is-active") {
		t.Fatal("expected modal to be closed")
	}

	if err := clickStep8Button(ctx, "Privacy Policy"); err != nil {
		t.Fatal(err)
	}
	if err := sleep(ctx, 500*time.Millisecond); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, ".modal.is-active") {
		t.Fatal("expected modal to be active after reopen")
	}
	if !elementExists(ctx, "#privacy_simple_content") {
		t.Fatal("expected content to be in Preview mode after reopen")
	}
}

func TestModalReset_ReopeningTCModalResetsFromMarkdownView(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizard(ctx); err != nil {
		t.Fatal(err)
	}

	if err := clickStep8Button(ctx, "Terms & Conditions"); err != nil {
		t.Fatal(err)
	}
	if err := sleep(ctx, 500*time.Millisecond); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, ".modal.is-active") {
		t.Fatal("expected modal to be active")
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`.modal.is-active button[onclick*="getMarkdown"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if err := sleep(ctx, 300*time.Millisecond); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, "#tandc_txtarea") {
		t.Fatal("expected textarea to be visible in Markdown view")
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`.modal.is-active button.delete`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if err := sleep(ctx, 300*time.Millisecond); err != nil {
		t.Fatal(err)
	}

	if err := clickStep8Button(ctx, "Terms & Conditions"); err != nil {
		t.Fatal(err)
	}
	if err := sleep(ctx, 500*time.Millisecond); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, "#tandc_content") {
		t.Fatal("expected content to be in Preview mode after reopen")
	}
}

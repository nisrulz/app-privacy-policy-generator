package e2e

import (
	"context"
	"strings"
	"testing"

	"github.com/chromedp/chromedp"
)

// clickModalButtonByText clicks a button by its exact text content within the active modal.
func clickModalButtonByText(ctx context.Context, text string) error {
	return chromedp.Run(ctx,
		chromedp.Evaluate(
			`Array.from(document.querySelectorAll('.modal.is-active button')).find(el => el.textContent.trim() === '`+text+`').click()`,
			nil,
		),
	)
}

// clickStep8LinkByText clicks a link by its text content within step 8.
func clickStep8LinkByText(ctx context.Context, text string) error {
	return chromedp.Run(ctx,
		chromedp.Evaluate(
			`Array.from(document.querySelectorAll('#step-8 a.card-footer-item')).find(el => el.textContent.trim() === '`+text+`').click()`,
			nil,
		),
	)
}

func TestExport_TCModalHTMLExportShowsRawHTML(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizard(ctx); err != nil {
		t.Fatal(err)
	}

	if err := clickStep8LinkByText(ctx, "Terms & Conditions"); err != nil {
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

	if !elementExists(ctx, "#tandc_txtarea") {
		t.Fatal("expected textarea to be visible")
	}

	var html string
	if err := chromedp.Run(ctx,
		chromedp.Value(`#tandc_txtarea`, &html, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(html, "<") {
		t.Fatal("expected HTML content to contain angle brackets")
	}
	if len(html) <= 50 {
		t.Fatal("expected HTML content length to be greater than 50")
	}
}

func TestExport_TCModalMarkdownExportShowsMarkdown(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizard(ctx); err != nil {
		t.Fatal(err)
	}

	if err := clickStep8LinkByText(ctx, "Terms & Conditions"); err != nil {
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

	if !elementExists(ctx, "#tandc_txtarea") {
		t.Fatal("expected textarea to be visible")
	}

	var md string
	if err := chromedp.Run(ctx,
		chromedp.Value(`#tandc_txtarea`, &md, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if len(md) <= 50 {
		t.Fatal("expected Markdown content length to be greater than 50")
	}
}

func TestExport_TCPreviewButtonSwitchesBackFromHTMLView(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizard(ctx); err != nil {
		t.Fatal(err)
	}

	if err := clickStep8LinkByText(ctx, "Terms & Conditions"); err != nil {
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

	if !elementExists(ctx, "#tandc_txtarea") {
		t.Fatal("expected textarea to be visible in HTML mode")
	}

	if err := clickModalButtonByText(ctx, "Preview"); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, "#tandc_content") {
		t.Fatal("expected content to be visible after switching to Preview")
	}
	if elementExists(ctx, "#tandc_txtarea") {
		t.Fatal("expected textarea to be hidden after switching to Preview")
	}
}

func TestExport_SimplePolicyPreviewButtonSwitchesBackFromHTMLView(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizard(ctx); err != nil {
		t.Fatal(err)
	}

	if err := clickStep8LinkByText(ctx, "Privacy Policy"); err != nil {
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

	if !elementExists(ctx, "#privacy_simple_txtarea") {
		t.Fatal("expected textarea to be visible in HTML mode")
	}

	if err := clickModalButtonByText(ctx, "Preview"); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, "#privacy_simple_content") {
		t.Fatal("expected content to be visible after switching to Preview")
	}
}

func TestExport_NoTrackingPolicyHasHTMLAndMarkdownExport(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizardNoTracking(ctx); err != nil {
		t.Fatal(err)
	}

	if err := clickStep8LinkByText(ctx, "Privacy Policy"); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, ".modal.is-active") {
		t.Fatal("expected modal to be active")
	}

	if !elementExists(ctx, `.modal.is-active button[onclick*="getHtml"]`) {
		t.Fatal("expected HTML button in No Tracking policy modal")
	}
	if !elementExists(ctx, `.modal.is-active button[onclick*="getMarkdown"]`) {
		t.Fatal("expected Markdown button in No Tracking policy modal")
	}
}

func TestExport_GDPRPolicyHasHTMLAndMarkdownExport(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizardGDPR(ctx, ""); err != nil {
		t.Fatal(err)
	}

	if err := clickStep8LinkByText(ctx, "Privacy Policy"); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, ".modal.is-active") {
		t.Fatal("expected modal to be active")
	}

	if !elementExists(ctx, `.modal.is-active button[onclick*="getHtml"]`) {
		t.Fatal("expected HTML button in GDPR policy modal")
	}
	if !elementExists(ctx, `.modal.is-active button[onclick*="getMarkdown"]`) {
		t.Fatal("expected Markdown button in GDPR policy modal")
	}
}

func TestExport_CanOpenAndCloseFAQModalFromStep1(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := gotoApp(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 1); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`a.has-info`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, ".modal.is-active") {
		t.Fatal("expected modal to be active")
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`.modal.is-active button.delete`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	if elementExists(ctx, ".modal.is-active") {
		t.Fatal("expected modal to be closed after clicking delete button")
	}
}

func TestExport_CanOpenAndCloseDisclaimerModalFromStep1(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := gotoApp(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 1); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Evaluate(
			`Array.from(document.querySelectorAll('#step-1 a')).find(el => /Disclaimer|Haftungsausschluss/.test(el.textContent.trim())).click()`,
			nil,
		),
	); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, ".modal.is-active") {
		t.Fatal("expected modal to be active")
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`.modal.is-active button.delete`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	if elementExists(ctx, ".modal.is-active") {
		t.Fatal("expected modal to be closed after clicking delete button")
	}
}

func TestExport_FAQModalCanBeClosedWithOKButton(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := gotoApp(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 1); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`a.has-info`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, ".modal.is-active") {
		t.Fatal("expected modal to be active")
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`.modal.is-active button.is-info`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	if elementExists(ctx, ".modal.is-active") {
		t.Fatal("expected modal to be closed after clicking OK button")
	}
}

package e2e

import (
	"context"
	"strings"
	"testing"

	"github.com/chromedp/chromedp"
)

// clickStep8Button clicks a button in step 8 footer by its visible text.
// Step 8 has multiple card-footer-item links (Previous, Privacy Policy, Terms & Conditions),
// so we need to target them by text content.
func clickStep8Button(ctx context.Context, text string) error {
	return chromedp.Run(ctx,
		chromedp.Evaluate(
			`Array.from(document.querySelectorAll('#step-8 a.card-footer-item')).find(el => el.textContent.trim() === '`+text+`').click()`,
			nil,
		),
	)
}

func TestOutput_GeneratesAndDisplaysSimplePrivacyPolicy(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizard(ctx); err != nil {
		t.Fatal(err)
	}

	if err := clickStep8Button(ctx, "Privacy Policy"); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, `.modal.is-active`) {
		t.Fatal("expected modal to be active")
	}
	if !elementExists(ctx, `#privacy_simple_content`) {
		t.Fatal("expected privacy_simple_content to be visible")
	}

	text, err := getElementText(ctx, `#privacy_simple_content`)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text, "Test App") {
		t.Fatal("expected policy to contain app name 'Test App'")
	}
	if !strings.Contains(text, "test@example.com") {
		t.Fatal("expected policy to contain contact 'test@example.com'")
	}
}

func TestOutput_GeneratesAndDisplaysTermsAndConditions(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizard(ctx); err != nil {
		t.Fatal(err)
	}

	if err := clickStep8Button(ctx, "Terms & Conditions"); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, `.modal.is-active`) {
		t.Fatal("expected modal to be active")
	}
	if !elementExists(ctx, `#tandc_content`) {
		t.Fatal("expected tandc_content to be visible")
	}

	text, err := getElementText(ctx, `#tandc_content`)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text, "Test App") {
		t.Fatal("expected T&C to contain app name 'Test App'")
	}
}

func TestOutput_CanSwitchToHTMLViewInPrivacyModal(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizard(ctx); err != nil {
		t.Fatal(err)
	}

	if err := clickStep8Button(ctx, "Privacy Policy"); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, `.modal.is-active`) {
		t.Fatal("expected modal to be active")
	}

	// Click the HTML button
	if err := chromedp.Run(ctx,
		chromedp.Click(`.modal.is-active button[onclick*="getHtml"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, `#privacy_simple_txtarea`) {
		t.Fatal("expected textarea to be visible in HTML mode")
	}

	var html string
	if err := chromedp.Run(ctx,
		chromedp.Value(`#privacy_simple_txtarea`, &html, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(html, "<") {
		t.Fatal("expected HTML content to contain angle brackets")
	}
}

func TestOutput_CanSwitchToMarkdownViewInPrivacyModal(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizard(ctx); err != nil {
		t.Fatal(err)
	}

	if err := clickStep8Button(ctx, "Privacy Policy"); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, `.modal.is-active`) {
		t.Fatal("expected modal to be active")
	}

	// Click the Markdown button
	if err := chromedp.Run(ctx,
		chromedp.Click(`.modal.is-active button[onclick*="getMarkdown"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, `#privacy_simple_txtarea`) {
		t.Fatal("expected textarea to be visible in Markdown mode")
	}

	var md string
	if err := chromedp.Run(ctx,
		chromedp.Value(`#privacy_simple_txtarea`, &md, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if len(md) == 0 {
		t.Fatal("expected Markdown content to be non-empty")
	}
}

func TestOutput_CanCloseModalWithCloseButton(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizard(ctx); err != nil {
		t.Fatal(err)
	}

	if err := clickStep8Button(ctx, "Privacy Policy"); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, `.modal.is-active`) {
		t.Fatal("expected modal to be active")
	}

	// Click the close button
	if err := chromedp.Run(ctx,
		chromedp.Click(`.modal.is-active button[aria-label="close"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	if elementExists(ctx, `.modal.is-active`) {
		t.Fatal("expected modal to be closed after clicking close button")
	}
}

func TestOutput_GDPRPolicyIncludesBusinessAddress(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizardGDPR(ctx, "123 Test Street, Berlin, Germany"); err != nil {
		t.Fatal(err)
	}

	if err := clickStep8Button(ctx, "Privacy Policy"); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, `.modal.is-active`) {
		t.Fatal("expected modal to be active")
	}

	text, err := getElementText(ctx, `#privacy_gdpr_content`)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text, "123 Test Street, Berlin, Germany") {
		t.Fatal("expected GDPR policy to contain business address")
	}
}

func TestOutput_PolicyIncludesAppNameAndDeveloperName(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizard(ctx); err != nil {
		t.Fatal(err)
	}

	if err := clickStep8Button(ctx, "Privacy Policy"); err != nil {
		t.Fatal(err)
	}

	text, err := getElementText(ctx, `#privacy_simple_content`)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text, "Test App") {
		t.Fatal("expected policy to contain app name 'Test App'")
	}
	if !strings.Contains(text, "John Doe") {
		t.Fatal("expected policy to contain developer name 'John Doe'")
	}
}

func TestOutput_TCModalHasHTMLAndMarkdownButtons(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizard(ctx); err != nil {
		t.Fatal(err)
	}

	if err := clickStep8Button(ctx, "Terms & Conditions"); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, `.modal.is-active`) {
		t.Fatal("expected modal to be active")
	}
	if !elementExists(ctx, `.modal.is-active button[onclick*="getHtml"]`) {
		t.Fatal("expected HTML button in T&C modal")
	}
	if !elementExists(ctx, `.modal.is-active button[onclick*="getMarkdown"]`) {
		t.Fatal("expected Markdown button in T&C modal")
	}
}

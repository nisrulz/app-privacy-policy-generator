package e2e

import (
	"testing"
	"time"

	"github.com/chromedp/chromedp"
)

func TestGenerateBlocking_ModalDoesNotOpenWhenFieldsCleared(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizard(ctx); err != nil {
		t.Fatal(err)
	}

	// Clear required fields by setting input values to empty and triggering blur
	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#appName`, "", chromedp.ByQuery),
		chromedp.Blur(`#appName`, chromedp.ByQuery),
		chromedp.SendKeys(`#appContact`, "", chromedp.ByQuery),
		chromedp.Blur(`#appContact`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	time.Sleep(200 * time.Millisecond)

	// Click the Privacy Policy button (has onclick handler for togglePrivacyModalVisibility)
	if err := chromedp.Run(ctx,
		chromedp.Click(`#step-8 a[onclick*="togglePrivacyModalVisibility"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	time.Sleep(500 * time.Millisecond)

	if elementExists(ctx, `.modal.is-active`) {
		t.Fatal("expected modal NOT to open when required fields are cleared")
	}
}

func TestGenerateBlocking_TermsModalDoesNotOpenWhenFieldsCleared(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizard(ctx); err != nil {
		t.Fatal(err)
	}

	// Clear required fields by setting input values to empty and triggering blur
	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#appName`, "", chromedp.ByQuery),
		chromedp.Blur(`#appName`, chromedp.ByQuery),
		chromedp.SendKeys(`#appContact`, "", chromedp.ByQuery),
		chromedp.Blur(`#appContact`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	time.Sleep(200 * time.Millisecond)

	// Click the Terms & Conditions button (has onclick handler for toggleTermsModalVisibility)
	if err := chromedp.Run(ctx,
		chromedp.Click(`#step-8 a[onclick*="toggleTermsModalVisibility"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	time.Sleep(500 * time.Millisecond)

	if elementExists(ctx, `.modal.is-active`) {
		t.Fatal("expected T&C modal NOT to open when required fields are cleared")
	}
}

func TestGenerateBlocking_ModalOpensWhenAllFieldsFilled(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizard(ctx); err != nil {
		t.Fatal(err)
	}

	// Click the Privacy Policy button
	if err := chromedp.Run(ctx,
		chromedp.Click(`#step-8 a[onclick*="togglePrivacyModalVisibility"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, `.modal.is-active`) {
		t.Fatal("expected modal to open when all fields are filled")
	}
}

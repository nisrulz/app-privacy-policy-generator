package e2e

import (
	"testing"

	"github.com/chromedp/chromedp"
)

func TestFormValidation_ShowsErrorWhenAppNameEmpty(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#appName`, "", chromedp.ByQuery),
		chromedp.Blur(`#appName`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, `#step-2 .help.is-danger`) {
		t.Fatal("expected error message for empty appName")
	}
}

func TestFormValidation_ShowsErrorWhenContactEmpty(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#appContact`, "", chromedp.ByQuery),
		chromedp.Blur(`#appContact`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, `#step-2 .help.is-danger`) {
		t.Fatal("expected error message for empty contact")
	}
}

func TestFormValidation_ClearsErrorWhenFieldFilled(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#appName`, "", chromedp.ByQuery),
		chromedp.Blur(`#appName`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if !elementExists(ctx, `#step-2 .help.is-danger`) {
		t.Fatal("expected error message for empty appName")
	}

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#appName`, "Test App", chromedp.ByQuery),
		chromedp.SendKeys(`#appContact`, "test@example.com", chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if elementExists(ctx, `#step-2 .help.is-danger`) {
		t.Fatal("expected no error messages when fields are filled")
	}
}

func TestFormValidation_Step6ShowsErrorForEmptyDevName(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep6(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#devName`, "", chromedp.ByQuery),
		chromedp.Blur(`#devName`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, `#step-6 .help.is-danger`) {
		t.Fatal("expected error message for empty dev name")
	}
}

func TestFormValidation_Step6ShowsErrorForEmptyCompanyName(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep6(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`input[type="radio"][value="Company"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#companyName`, "", chromedp.ByQuery),
		chromedp.Blur(`#companyName`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, `#step-6 .help.is-danger`) {
		t.Fatal("expected error message for empty company name")
	}
}

func TestFormValidation_GDPRShowsBusinessAddressField(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#appName`, "Test App", chromedp.ByQuery),
		chromedp.SendKeys(`#appContact`, "test@example.com", chromedp.ByQuery),
		chromedp.Click(`input[type="radio"][value="3"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if err := clickNext(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 3); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, `#businessAddress`) {
		t.Fatal("expected business address field to be visible for GDPR policy")
	}
}

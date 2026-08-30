package e2e

import (
	"strings"
	"testing"

	"github.com/chromedp/chromedp"
)

func TestCanAdvance_ClickingNextOnStep2EmptyFieldsDoesNotAdvance(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}

	// Force-click the Next button even though it is disabled
	if err := chromedp.Run(ctx,
		chromedp.Click(`#step-2 a.card-footer-item >> visible=true`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	if err := expectStep(ctx, 2); err != nil {
		t.Fatal("should still be on step 2 after clicking disabled Next")
	}
}

func TestCanAdvance_ClickingNextOnStep2OnlyAppNameDoesNotAdvance(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#appName`, "Test App", chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	// Force-click the Next button even though it is disabled
	if err := chromedp.Run(ctx,
		chromedp.Click(`#step-2 a.card-footer-item >> visible=true`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	if err := expectStep(ctx, 2); err != nil {
		t.Fatal("should still be on step 2 with only appName filled")
	}
}

func TestCanAdvance_ClickingNextOnStep2OnlyContactDoesNotAdvance(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#appContact`, "test@example.com", chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	// Force-click the Next button even though it is disabled
	if err := chromedp.Run(ctx,
		chromedp.Click(`#step-2 a.card-footer-item >> visible=true`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	if err := expectStep(ctx, 2); err != nil {
		t.Fatal("should still be on step 2 with only contact filled")
	}
}

func TestCanAdvance_ClickingNextOnStep6IndividualEmptyDevNameDoesNotAdvance(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}
	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#appName`, "Test App", chromedp.ByQuery),
		chromedp.SendKeys(`#appContact`, "test@example.com", chromedp.ByQuery),
	); err != nil {
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
	if err := clickNext(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 6); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`input[type="radio"][value="Individual"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	// Force-click the Next button even though it is disabled
	if err := chromedp.Run(ctx,
		chromedp.Click(`#step-6 a.card-footer-item >> visible=true`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	if err := expectStep(ctx, 6); err != nil {
		t.Fatal("should still be on step 6 with empty devName")
	}
}

func TestCanAdvance_ClickingNextOnStep6CompanyEmptyCompanyNameDoesNotAdvance(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}
	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#appName`, "Test App", chromedp.ByQuery),
		chromedp.SendKeys(`#appContact`, "test@example.com", chromedp.ByQuery),
	); err != nil {
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
	if err := clickNext(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 6); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`input[type="radio"][value="Company"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	// Force-click the Next button even though it is disabled
	if err := chromedp.Run(ctx,
		chromedp.Click(`#step-6 a.card-footer-item >> visible=true`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	if err := expectStep(ctx, 6); err != nil {
		t.Fatal("should still be on step 6 with empty companyName")
	}
}

func TestCanAdvance_NextButtonHasIsDisabledClassWhenCannotAdvance(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}

	class, err := getElementAttribute(ctx, `#step-2 a.card-footer-item >> visible=true`, "class")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(class, "is-disabled") {
		t.Fatal("expected Next button to have is-disabled class when canAdvance is false")
	}
}

func TestCanAdvance_NextButtonLosesIsDisabledClassWhenFieldsFilled(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#appName`, "Test App", chromedp.ByQuery),
		chromedp.SendKeys(`#appContact`, "test@example.com", chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	class, err := getElementAttribute(ctx, `#step-2 a.card-footer-item >> visible=true`, "class")
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(class, "is-disabled") {
		t.Fatal("expected Next button to NOT have is-disabled class when fields are filled")
	}
}

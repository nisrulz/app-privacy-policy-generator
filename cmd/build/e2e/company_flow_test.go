package e2e

import (
	"strings"
	"testing"

	"github.com/chromedp/chromedp"
)

func TestStep6ShowsCompanyNameFieldWhenCompanySelected(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep6(ctx); err != nil {
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

	if !elementExists(ctx, "#companyName") {
		t.Error("expected company name field to be visible")
	}
	if elementExists(ctx, "#devName") {
		t.Error("expected dev name field to be hidden")
	}
}

func TestStep6ShowsDevNameFieldWhenIndividualSelected(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep6(ctx); err != nil {
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

	if !elementExists(ctx, "#devName") {
		t.Error("expected dev name field to be visible")
	}
	if elementExists(ctx, "#companyName") {
		t.Error("expected company name field to be hidden")
	}
}

func TestStep6BlocksAdvancementWithoutCompanyName(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep6(ctx); err != nil {
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

	class, err := getElementAttribute(ctx, `a.card-footer-item >> visible=true`, "class")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(class, "is-disabled") {
		t.Error("expected Next button to be disabled when company name is empty")
	}

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#companyName`, "Acme Corp", chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	class, err = getElementAttribute(ctx, `a.card-footer-item >> visible=true`, "class")
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(class, "is-disabled") {
		t.Error("expected Next button to be enabled after entering company name")
	}
}

func TestCompanyNameAppearsInGeneratedPrivacyPolicy(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizardCompany(ctx); err != nil {
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
	if !strings.Contains(content, "Acme Corp") {
		t.Error("expected policy to contain 'Acme Corp'")
	}
	if strings.Contains(content, "[Developer/Company name]") {
		t.Error("expected policy to not contain placeholder text")
	}
}

func TestCompanyNameAppearsInTermsAndConditions(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizardCompany(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`a.card-footer-item >> visible=true`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	content, err := getElementText(ctx, "#tandc_content")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(content, "Acme Corp") {
		t.Error("expected T&C to contain 'Acme Corp'")
	}
}

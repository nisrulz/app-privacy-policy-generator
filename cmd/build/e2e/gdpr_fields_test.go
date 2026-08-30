package e2e

import (
	"strings"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
)

func TestGDPRFields_EURepresentativeAppearsWhenGDPRSelected(t *testing.T) {
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

	if !elementExists(ctx, `#euRepresentative`) {
		t.Fatal("expected EU representative field to be visible for GDPR policy")
	}
}

func TestGDPRFields_EURepresentativeHiddenForNonGDPR(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#appName`, "Test App", chromedp.ByQuery),
		chromedp.SendKeys(`#appContact`, "test@example.com", chromedp.ByQuery),
		chromedp.Click(`input[type="radio"][value="1"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	if elementExists(ctx, `#euRepresentative`) {
		t.Fatal("expected EU representative field to be hidden for non-GDPR policy")
	}
}

func TestGDPRFields_EURepresentativeAppearsInOutput(t *testing.T) {
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

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#euRepresentative`, "EU Rep GmbH", chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	if err := clickNext(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 3); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#businessAddress`, "123 EU Street, Berlin", chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	for step := 4; step <= 6; step++ {
		if err := clickNext(ctx); err != nil {
			t.Fatal(err)
		}
		if err := expectStep(ctx, step); err != nil {
			t.Fatal(err)
		}
	}

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#devName`, "John Doe", chromedp.ByQuery),
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
	if !strings.Contains(text, "EU Rep GmbH") {
		t.Fatal("expected GDPR policy to contain EU representative name")
	}
}

func TestGDPRFields_BusinessAddressRequiredForGDPR(t *testing.T) {
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

	class, err := getElementAttribute(ctx, `#businessAddress`, "class")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(class, "is-danger") {
		t.Fatal("expected business address to have is-danger class for GDPR policy")
	}
}

func TestGDPRFields_BusinessAddressNotRequiredForSimple(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#appName`, "Test App", chromedp.ByQuery),
		chromedp.SendKeys(`#appContact`, "test@example.com", chromedp.ByQuery),
		chromedp.Click(`input[type="radio"][value="1"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if err := clickNext(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 3); err != nil {
		t.Fatal(err)
	}

	class, err := getElementAttribute(ctx, `#businessAddress`, "class")
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(class, "is-danger") {
		t.Fatal("expected business address to not have is-danger class for Simple policy")
	}
}

func TestGDPRFields_BusinessAddressAppearsInOutput(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizardGDPR(ctx, "456 GDPR Lane, Munich"); err != nil {
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
	if !strings.Contains(text, "456 GDPR Lane, Munich") {
		t.Fatal("expected GDPR policy to contain business address")
	}
}

func TestGDPRFields_EffectiveDateHasTodayDefault(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
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

	if !elementExists(ctx, `#effectiveFromDate`) {
		t.Fatal("expected effective date field to be visible")
	}

	today := time.Now().Format("2006-01-02")
	value, err := getElementAttribute(ctx, `#effectiveFromDate`, "value")
	if err != nil {
		t.Fatal(err)
	}
	if value != today {
		t.Fatalf("expected effective date to be %q, got %q", today, value)
	}
}

func TestGDPRFields_CanChangeEffectiveDate(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
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

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#effectiveFromDate`, "2025-01-15", chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	value, err := getElementAttribute(ctx, `#effectiveFromDate`, "value")
	if err != nil {
		t.Fatal(err)
	}
	if value != "2025-01-15" {
		t.Fatalf("expected effective date to be %q, got %q", "2025-01-15", value)
	}
}

func TestGDPRFields_PIDInfoVisibleForSimple(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}
	if err := fillStep2(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`input[type="radio"][value="1"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if err := clickNext(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 3); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, `#pidInfoIn`) {
		t.Fatal("expected PID info field to be visible for Simple policy")
	}
}

func TestGDPRFields_AgeOfDigitalConsentDefault16(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
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

	if !elementExists(ctx, `#ageOfDigitalConsent`) {
		t.Fatal("expected age of digital consent field to be visible")
	}

	value, err := getElementAttribute(ctx, `#ageOfDigitalConsent`, "value")
	if err != nil {
		t.Fatal(err)
	}
	if value != "16" {
		t.Fatalf("expected age of digital consent to be %q, got %q", "16", value)
	}
}

func TestGDPRFields_CanChangeAgeOfDigitalConsent(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
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

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#ageOfDigitalConsent`, "13", chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	value, err := getElementAttribute(ctx, `#ageOfDigitalConsent`, "value")
	if err != nil {
		t.Fatal(err)
	}
	if value != "13" {
		t.Fatalf("expected age of digital consent to be %q, got %q", "13", value)
	}
}

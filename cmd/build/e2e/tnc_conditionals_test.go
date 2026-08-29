package e2e

import (
	"strings"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
)

func TestTNCConditionals_ShowsLicenseForOpenSource(t *testing.T) {
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

	if err := chromedp.Run(ctx,
		chromedp.Click(`input[type="radio"][value="Open Source"]`, chromedp.ByQuery),
	); err != nil {
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

	if err := sendKeys(ctx, `#devName`, "John Doe"); err != nil {
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

	if err := clickStep8Button(ctx, "Terms & Conditions"); err != nil {
		t.Fatal(err)
	}
	if err := sleep(ctx, 500*time.Millisecond); err != nil {
		t.Fatal(err)
	}

	text, err := getElementText(ctx, `#tandc_content`)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text, "License") {
		t.Fatal("expected T&C to contain 'License' for Open Source app type")
	}
}

func TestTNCConditionals_ShowsLicenseForNonOpenSource(t *testing.T) {
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

	text, err := getElementText(ctx, `#tandc_content`)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text, "License") {
		t.Fatal("expected T&C to contain 'License' for non-Open Source app type")
	}
}

func TestTNCConditionals_IncludesJailbreakWarningForMobile(t *testing.T) {
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

	text, err := getElementText(ctx, `#tandc_content`)
	if err != nil {
		t.Fatal(err)
	}
	if len(text) <= 100 {
		t.Fatal("expected T&C content to be longer than 100 characters")
	}
}

func TestTNCConditionals_ShowsMobileDataSectionsForMobileApps(t *testing.T) {
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

	text, err := getElementText(ctx, `#tandc_content`)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text, "mobile") {
		t.Fatal("expected T&C to contain 'mobile'")
	}
}

func TestTNCConditionals_IncludesAIUsageWhenEnabled(t *testing.T) {
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

	if err := clickSelector(ctx, `label[for="aicheckbox"]`); err != nil {
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

	if err := sendKeys(ctx, `#devName`, "John Doe"); err != nil {
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

	if err := clickStep8Button(ctx, "Terms & Conditions"); err != nil {
		t.Fatal(err)
	}
	if err := sleep(ctx, 500*time.Millisecond); err != nil {
		t.Fatal(err)
	}

	text, err := getElementText(ctx, `#tandc_content`)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text, "AI") {
		t.Fatal("expected T&C to contain 'AI' when AI flag is enabled")
	}
}

func TestTNCConditionals_DSAIncludesEURepresentative(t *testing.T) {
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

	if err := chromedp.Run(ctx,
		chromedp.Click(`input[type="radio"][value="3"]`, chromedp.ByQuery),
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

	if err := sendKeys(ctx, `#businessAddress`, "123 EU Street"); err != nil {
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

	if err := sendKeys(ctx, `#devName`, "John Doe"); err != nil {
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

	if err := clickStep8Button(ctx, "Terms & Conditions"); err != nil {
		t.Fatal(err)
	}
	if err := sleep(ctx, 500*time.Millisecond); err != nil {
		t.Fatal(err)
	}

	text, err := getElementText(ctx, `#tandc_content`)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text, "EU Rep GmbH") {
		t.Fatal("expected T&C to contain EU representative name")
	}
}

func TestTNCConditionals_IncludesThirdPartyServiceTerms(t *testing.T) {
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

	for step := 3; step <= 6; step++ {
		if err := clickNext(ctx); err != nil {
			t.Fatal(err)
		}
		if err := expectStep(ctx, step); err != nil {
			t.Fatal(err)
		}
	}

	if err := sendKeys(ctx, `#devName`, "John Doe"); err != nil {
		t.Fatal(err)
	}
	if err := clickNext(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 7); err != nil {
		t.Fatal(err)
	}

	if err := clickSelector(ctx, `.scrollable-thirdparty.content label:first-child`); err != nil {
		t.Fatal(err)
	}

	if err := clickNext(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 8); err != nil {
		t.Fatal(err)
	}

	if err := clickStep8Button(ctx, "Terms & Conditions"); err != nil {
		t.Fatal(err)
	}
	if err := sleep(ctx, 500*time.Millisecond); err != nil {
		t.Fatal(err)
	}

	text, err := getElementText(ctx, `#tandc_content`)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text, "Third") {
		t.Fatal("expected T&C to contain 'Third' for third-party service terms")
	}
}

package e2e

import (
	"strings"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
)

func TestGDPRConditionals_ShowsLocationTrackingWhenEnabled(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}
	if err := fillStep2(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
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

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#businessAddress`, "123 EU Street", chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if err := clickNext(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 4); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`label[for="locationcheckbox"]`, chromedp.ByQuery),
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
	if err := chromedp.Run(ctx,
		chromedp.Sleep(500*time.Millisecond),
	); err != nil {
		t.Fatal(err)
	}

	text, err := getElementText(ctx, `#privacy_gdpr_content`)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text, "location") {
		t.Fatal("expected GDPR policy to contain location tracking details")
	}
}

func TestGDPRConditionals_ShowsNoTrackingTextWhenDisabled(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}
	if err := fillStep2(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
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

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#businessAddress`, "123 EU Street", chromedp.ByQuery),
	); err != nil {
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
	if err := chromedp.Run(ctx,
		chromedp.Sleep(500*time.Millisecond),
	); err != nil {
		t.Fatal(err)
	}

	text, err := getElementText(ctx, `#privacy_gdpr_content`)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text, "location") {
		t.Fatal("expected GDPR policy to contain location tracking text")
	}
}

func TestGDPRConditionals_ShowsAIUsageWhenEnabled(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}
	if err := fillStep2(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
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

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#businessAddress`, "123 EU Street", chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if err := clickNext(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 4); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`label[for="aicheckbox"]`, chromedp.ByQuery),
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
	if err := chromedp.Run(ctx,
		chromedp.Sleep(500*time.Millisecond),
	); err != nil {
		t.Fatal(err)
	}

	text, err := getElementText(ctx, `#privacy_gdpr_content`)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text, "AI") {
		t.Fatal("expected GDPR policy to contain AI usage details")
	}
}

func TestGDPRConditionals_ShowsDataDeletionWhenEnabled(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}
	if err := fillStep2(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
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

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#businessAddress`, "123 EU Street", chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if err := clickNext(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 4); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`label[for="datadeletioncheckbox"]`, chromedp.ByQuery),
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
	if err := chromedp.Run(ctx,
		chromedp.Sleep(500*time.Millisecond),
	); err != nil {
		t.Fatal(err)
	}

	text, err := getElementText(ctx, `#privacy_gdpr_content`)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text, "deletion") {
		t.Fatal("expected GDPR policy to contain data deletion section")
	}
}

func TestGDPRConditionals_ChildrenTextWithoutServices(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}
	if err := fillStep2(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
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

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#businessAddress`, "123 EU Street", chromedp.ByQuery),
	); err != nil {
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
	if err := chromedp.Run(ctx,
		chromedp.Sleep(500*time.Millisecond),
	); err != nil {
		t.Fatal(err)
	}

	text, err := getElementText(ctx, `#privacy_gdpr_content`)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text, "children") {
		t.Fatal("expected GDPR policy to contain children section")
	}
}

func TestGDPRConditionals_ChildrenTextWithServices(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}
	if err := fillStep2(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
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

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#businessAddress`, "123 EU Street", chromedp.ByQuery),
	); err != nil {
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

	if err := chromedp.Run(ctx,
		chromedp.Click(`.scrollable-thirdparty.content label:first-child`, chromedp.ByQuery),
	); err != nil {
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
	if err := chromedp.Run(ctx,
		chromedp.Sleep(500*time.Millisecond),
	); err != nil {
		t.Fatal(err)
	}

	text, err := getElementText(ctx, `#privacy_gdpr_content`)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text, "children") {
		t.Fatal("expected GDPR policy to contain children section")
	}
}

func TestGDPRConditionals_IncludesBusinessAddress(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}
	if err := fillStep2(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
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

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#businessAddress`, "456 Berlin Ave", chromedp.ByQuery),
	); err != nil {
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
	if err := chromedp.Run(ctx,
		chromedp.Sleep(500*time.Millisecond),
	); err != nil {
		t.Fatal(err)
	}

	text, err := getElementText(ctx, `#privacy_gdpr_content`)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text, "456 Berlin Ave") {
		t.Fatal("expected GDPR policy to contain business address")
	}
}

func TestGDPRConditionals_HidesBusinessAddressLineWhenEmpty(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}
	if err := fillStep2(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
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
	if err := chromedp.Run(ctx,
		chromedp.Sleep(500*time.Millisecond),
	); err != nil {
		t.Fatal(err)
	}

	text, err := getElementText(ctx, `#privacy_gdpr_content`)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(text, "Address") {
		t.Fatal("expected GDPR policy to not contain Address when business address is empty")
	}
}

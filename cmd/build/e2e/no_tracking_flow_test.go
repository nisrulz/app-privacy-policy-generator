package e2e

import (
	"strings"
	"testing"

	"github.com/chromedp/chromedp"
)

func TestNoTrackingFlow_NavigatesFullWizard(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#appName`, "NoTrack App", chromedp.ByQuery),
		chromedp.SendKeys(`#appContact`, "notrack@example.com", chromedp.ByQuery),
		chromedp.Click(`input[type="radio"][value="2"]`, chromedp.ByQuery),
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
		chromedp.SendKeys(`#devName`, "NoTrack Dev", chromedp.ByQuery),
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
}

func TestNoTrackingFlow_Step3HidesPIDInfo(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#appName`, "Test App", chromedp.ByQuery),
		chromedp.SendKeys(`#appContact`, "test@example.com", chromedp.ByQuery),
		chromedp.Click(`input[type="radio"][value="2"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if err := clickNext(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 3); err != nil {
		t.Fatal(err)
	}

	if elementExists(ctx, `#pidInfoIn`) {
		t.Fatal("expected PID info field to be hidden for No Tracking policy")
	}
}

func TestNoTrackingFlow_Step4HidesFeatureCheckboxes(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#appName`, "Test App", chromedp.ByQuery),
		chromedp.SendKeys(`#appContact`, "test@example.com", chromedp.ByQuery),
		chromedp.Click(`input[type="radio"][value="2"]`, chromedp.ByQuery),
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

	if elementExists(ctx, `#locationcheckbox`) {
		t.Fatal("expected location checkbox to be hidden for No Tracking policy")
	}
	if elementExists(ctx, `#aicheckbox`) {
		t.Fatal("expected AI checkbox to be hidden for No Tracking policy")
	}
	if elementExists(ctx, `#datadeletioncheckbox`) {
		t.Fatal("expected data deletion checkbox to be hidden for No Tracking policy")
	}
}

func TestNoTrackingFlow_OpensPrivacyPolicyModal(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizardNoTracking(ctx); err != nil {
		t.Fatal(err)
	}

	if err := clickStep8Button(ctx, "Privacy Policy"); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, `.modal.is-active`) {
		t.Fatal("expected modal to be active")
	}
	if !elementExists(ctx, `#privacy_notrack_content`) {
		t.Fatal("expected privacy_notrack_content to be visible")
	}

	text, err := getElementText(ctx, `#privacy_notrack_content`)
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

func TestNoTrackingFlow_PrivacyPolicyHasHTMLAndMarkdownExport(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizardNoTracking(ctx); err != nil {
		t.Fatal(err)
	}

	if err := clickStep8Button(ctx, "Privacy Policy"); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, `.modal.is-active`) {
		t.Fatal("expected modal to be active")
	}
	if !elementExists(ctx, `.modal.is-active button[onclick*="getHtml"]`) {
		t.Fatal("expected HTML button in No Tracking privacy modal")
	}
	if !elementExists(ctx, `.modal.is-active button[onclick*="getMarkdown"]`) {
		t.Fatal("expected Markdown button in No Tracking privacy modal")
	}
}

func TestNoTrackingFlow_ShowsNoServicesOnStep7(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#appName`, "Test App", chromedp.ByQuery),
		chromedp.SendKeys(`#appContact`, "test@example.com", chromedp.ByQuery),
		chromedp.Click(`input[type="radio"][value="2"]`, chromedp.ByQuery),
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
		chromedp.SendKeys(`#devName`, "Test Dev", chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if err := clickNext(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 7); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, `.thirdparty-section`) {
		t.Fatal("expected third-party section to be visible on step 7")
	}
}

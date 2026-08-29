package e2e

import (
	"strings"
	"testing"

	"github.com/chromedp/chromedp"
)

func TestWizardFlow_NavigatesThroughAllStepsSimple(t *testing.T) {
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
		chromedp.SendKeys(`#devName`, "Test Developer", chromedp.ByQuery),
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

func TestWizardFlow_NavigatesThroughAllStepsGDPR(t *testing.T) {
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
		chromedp.SendKeys(`#businessAddress`, "123 Test Street, Berlin, Germany", chromedp.ByQuery),
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
		chromedp.SendKeys(`#devName`, "Test Developer", chromedp.ByQuery),
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

func TestWizardFlow_CanNavigateBackwards(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}

	if err := clickPrevious(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 1); err != nil {
		t.Fatal(err)
	}

	if err := clickStart(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 2); err != nil {
		t.Fatal(err)
	}
}

func TestWizardFlow_Step2BlocksWithoutRequiredFields(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}

	nextBtnSelector := `#step-2 a.card-footer-item >> visible=true`
	class, err := getElementAttribute(ctx, nextBtnSelector, "class")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(class, "is-disabled") {
		t.Fatal("expected Next button to be disabled initially")
	}

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#appName`, "My App", chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	class, err = getElementAttribute(ctx, nextBtnSelector, "class")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(class, "is-disabled") {
		t.Fatal("expected Next button to be disabled with only appName")
	}

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#appContact`, "test@example.com", chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	class, err = getElementAttribute(ctx, nextBtnSelector, "class")
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(class, "is-disabled") {
		t.Fatal("expected Next button to be enabled when both fields filled")
	}
}

func TestWizardFlow_Step6BlocksWithoutDevName(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep6(ctx); err != nil {
		t.Fatal(err)
	}

	nextBtnSelector := `#step-6 a.card-footer-item >> visible=true`
	class, err := getElementAttribute(ctx, nextBtnSelector, "class")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(class, "is-disabled") {
		t.Fatal("expected Next button to be disabled initially")
	}

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#devName`, "Test Dev", chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	class, err = getElementAttribute(ctx, nextBtnSelector, "class")
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(class, "is-disabled") {
		t.Fatal("expected Next button to be enabled when devName filled")
	}
}

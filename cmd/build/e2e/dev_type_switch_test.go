package e2e

import (
	"testing"

	"github.com/chromedp/chromedp"
)

func TestSwitchingFromCompanyToIndividualShowsDevName(t *testing.T) {
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

	if err := chromedp.Run(ctx,
		chromedp.Click(`input[type="radio"][value="Individual"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if !elementExists(ctx, "#devName") {
		t.Error("expected dev name field to be visible after switching to Individual")
	}
	if elementExists(ctx, "#companyName") {
		t.Error("expected company name field to be hidden after switching to Individual")
	}
}

func TestSwitchingFromIndividualToCompanyShowsCompanyName(t *testing.T) {
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

	if err := chromedp.Run(ctx,
		chromedp.Click(`input[type="radio"][value="Company"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if !elementExists(ctx, "#companyName") {
		t.Error("expected company name field to be visible after switching to Company")
	}
	if elementExists(ctx, "#devName") {
		t.Error("expected dev name field to be hidden after switching to Company")
	}
}

package e2e

import (
	"testing"

	"github.com/chromedp/chromedp"
)

func TestLocales_DropdownIsVisibleOnStep1(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := gotoApp(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 1); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, "select.locale-dropdown") {
		t.Fatal("expected locale dropdown to be visible on step 1")
	}
}

func TestLocales_LocalePersistsAfterNavigatingToStep2(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := gotoApp(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 1); err != nil {
		t.Fatal(err)
	}

	secondValue, err := getElementAttribute(ctx, "select.locale-dropdown option:nth-child(2)", "value")
	if err != nil {
		t.Fatal(err)
	}
	if secondValue == "" {
		t.Fatal("expected a second locale option to exist")
	}

	if err := chromedp.Run(ctx,
		chromedp.SetValue("select.locale-dropdown", secondValue, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	if err := clickStart(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 2); err != nil {
		t.Fatal(err)
	}

	if err := clickPrevious(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 1); err != nil {
		t.Fatal(err)
	}

	currentValue, err := getElementAttribute(ctx, "select.locale-dropdown", "value")
	if err != nil {
		t.Fatal(err)
	}
	if currentValue != secondValue {
		t.Fatalf("expected locale to persist as '%s', got '%s'", secondValue, currentValue)
	}
}

package e2e

import (
	"strconv"
	"strings"
	"testing"

	"github.com/chromedp/chromedp"
)

func TestProgressBar_Step2ShowsProgressBar(t *testing.T) {
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

	if !elementExists(ctx, `#step-2 progress.progress`) {
		t.Fatal("expected progress bar to be visible on step 2")
	}
}

func TestProgressBar_ValueIncreasesAsWizardAdvances(t *testing.T) {
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

	value2Str, err := getElementAttribute(ctx, `#step-2 progress.progress`, "value")
	if err != nil {
		t.Fatal(err)
	}
	value2, err := strconv.ParseFloat(value2Str, 64)
	if err != nil {
		t.Fatalf("failed to parse progress value on step 2: %v", err)
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

	value3Str, err := getElementAttribute(ctx, `#step-3 progress.progress`, "value")
	if err != nil {
		t.Fatal(err)
	}
	value3, err := strconv.ParseFloat(value3Str, 64)
	if err != nil {
		t.Fatalf("failed to parse progress value on step 3: %v", err)
	}

	if value3 <= value2 {
		t.Fatalf("expected progress to increase from step 2 (%v) to step 3 (%v)", value2, value3)
	}
}

func TestProgressBar_Reaches100OnStep8(t *testing.T) {
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

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#devName`, "Test", chromedp.ByQuery),
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

	if !elementExists(ctx, `#step-8 progress.progress`) {
		t.Fatal("expected progress bar to be visible on step 8")
	}

	valueStr, err := getElementAttribute(ctx, `#step-8 progress.progress`, "value")
	if err != nil {
		t.Fatal(err)
	}
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		t.Fatalf("failed to parse progress value on step 8: %v", err)
	}
	if value != 100 {
		t.Fatalf("expected progress to be 100 on step 8, got %v", value)
	}
}

func TestProgressBar_UsesPrimaryClassOnStep8(t *testing.T) {
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

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#devName`, "Test", chromedp.ByQuery),
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

	class, err := getElementAttribute(ctx, `#step-8 progress.progress`, "class")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(class, "is-primary") {
		t.Fatalf("expected progress bar to have 'is-primary' class, got '%s'", class)
	}
}

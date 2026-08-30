package e2e

import (
	"strings"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
)

func TestTNCWebSections_ShowsWebDataSectionsForWebApps(t *testing.T) {
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

	if err := sendKeys(ctx, `#appName`, "Web App"); err != nil {
		t.Fatal(err)
	}
	if err := sendKeys(ctx, `#appContact`, "test@example.com"); err != nil {
		t.Fatal(err)
	}

	for step := 3; step <= 5; step++ {
		if err := clickNext(ctx); err != nil {
			t.Fatal(err)
		}
		if err := expectStep(ctx, step); err != nil {
			t.Fatal(err)
		}
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`label[for="platform-android"]`, chromedp.ByQuery),
		chromedp.Click(`label[for="platform-web"]`, chromedp.ByQuery),
	); err != nil {
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
	if !strings.Contains(text, "internet service provider") {
		t.Fatal("expected T&C to contain 'internet service provider' for web apps")
	}
}

func TestTNCWebSections_DoesNotShowMobileDataForWebOnlyApps(t *testing.T) {
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

	if err := sendKeys(ctx, `#appName`, "Web App"); err != nil {
		t.Fatal(err)
	}
	if err := sendKeys(ctx, `#appContact`, "test@example.com"); err != nil {
		t.Fatal(err)
	}

	for step := 3; step <= 5; step++ {
		if err := clickNext(ctx); err != nil {
			t.Fatal(err)
		}
		if err := expectStep(ctx, step); err != nil {
			t.Fatal(err)
		}
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`label[for="platform-android"]`, chromedp.ByQuery),
		chromedp.Click(`label[for="platform-web"]`, chromedp.ByQuery),
	); err != nil {
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
	if strings.Contains(text, "mobile network") {
		t.Fatal("expected T&C to not contain 'mobile network' for web-only apps")
	}
}

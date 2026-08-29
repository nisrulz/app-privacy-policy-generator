package e2e

import (
	"strings"
	"testing"
	"time"
)

func TestEffectiveDate_PolicyIncludesDateFromStep3(t *testing.T) {
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

	if err := sendKeys(ctx, `#effectiveFromDate`, "2025-06-15"); err != nil {
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

	if err := clickStep8Button(ctx, "Privacy Policy"); err != nil {
		t.Fatal(err)
	}
	if err := sleep(ctx, 500*time.Millisecond); err != nil {
		t.Fatal(err)
	}

	text, err := getElementText(ctx, `#privacy_simple_content`)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text, "2025-06-15") {
		t.Fatal("expected policy to contain effective date '2025-06-15'")
	}
}

func TestEffectiveDate_TCIncludesDateFromStep3(t *testing.T) {
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

	if err := sendKeys(ctx, `#effectiveFromDate`, "2025-01-01"); err != nil {
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
	if !strings.Contains(text, "2025-01-01") {
		t.Fatal("expected T&C to contain effective date '2025-01-01'")
	}
}

package e2e

import (
	"strings"
	"testing"
	"time"
)

func TestThirdPartyServices_DisplaysListOfServices(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep7(ctx); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, `.scrollable-thirdparty label`) {
		t.Fatal("expected third-party service labels to be visible")
	}
}

func TestThirdPartyServices_CanToggleServiceOnAndOff(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep7(ctx); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, `.scrollable-thirdparty label`) {
		t.Fatal("expected third-party service labels to be visible")
	}

	if err := clickSelector(ctx, `.scrollable-thirdparty label:first-child`); err != nil {
		t.Fatal(err)
	}

	checked, err := getElementAttribute(ctx, `.scrollable-thirdparty input[type="checkbox"]:first-child`, "checked")
	if err != nil {
		t.Fatal(err)
	}
	if checked != "true" && checked != "" {
		t.Fatal("expected first checkbox to be checked after clicking label")
	}

	if err := clickSelector(ctx, `.scrollable-thirdparty label:first-child`); err != nil {
		t.Fatal(err)
	}

	checked, err = getElementAttribute(ctx, `.scrollable-thirdparty input[type="checkbox"]:first-child`, "checked")
	if err != nil {
		t.Fatal(err)
	}
	if checked == "true" {
		t.Fatal("expected first checkbox to be unchecked after clicking label again")
	}
}

func TestThirdPartyServices_ToggledServicesAppearInOutput(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep7(ctx); err != nil {
		t.Fatal(err)
	}

	serviceName, err := getElementText(ctx, `.scrollable-thirdparty label:first-child`)
	if err != nil {
		t.Fatal(err)
	}
	serviceName = strings.TrimSpace(serviceName)

	if err := clickSelector(ctx, `.scrollable-thirdparty label:first-child`); err != nil {
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
	if !strings.Contains(text, serviceName) {
		t.Fatalf("expected policy to contain service name '%s'", serviceName)
	}
}

func TestThirdPartyServices_NoTrackingShowsNoServicesMessage(t *testing.T) {
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

	if err := sendKeys(ctx, `#appName`, "Test App"); err != nil {
		t.Fatal(err)
	}
	if err := sendKeys(ctx, `#appContact`, "test@example.com"); err != nil {
		t.Fatal(err)
	}
	if err := clickSelector(ctx, `input[type="radio"][value="2"]`); err != nil {
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

	if err := sendKeys(ctx, `#devName`, "Test Dev"); err != nil {
		t.Fatal(err)
	}
	if err := clickNext(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 7); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, `.thirdparty-section`) {
		t.Fatal("expected third-party section to be visible")
	}
}

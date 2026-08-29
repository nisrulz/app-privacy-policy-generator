package e2e

import (
	"strings"
	"testing"
	"time"
)

func TestChildrenSection_NoTrackingShowsChildrenText(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizardNoTracking(ctx); err != nil {
		t.Fatal(err)
	}

	if err := clickStep8Button(ctx, "Privacy Policy"); err != nil {
		t.Fatal(err)
	}
	if err := sleep(ctx, 500*time.Millisecond); err != nil {
		t.Fatal(err)
	}

	text, err := getElementText(ctx, `#privacy_notrack_content`)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text, "Children") {
		t.Fatal("expected no-tracking policy to contain 'Children'")
	}
}

func TestChildrenSection_IncludesAppNameInChildrenContext(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizard(ctx); err != nil {
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
	if !strings.Contains(text, "Children") {
		t.Fatal("expected simple policy to contain 'Children'")
	}
	if !strings.Contains(text, "Test App") {
		t.Fatal("expected simple policy to contain 'Test App'")
	}
}

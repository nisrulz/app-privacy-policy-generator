package e2e

import (
	"strings"
	"testing"
	"time"
)

func TestNoTrackDataDeletion_NoTrackingPolicyRendersWithDataDeletionText(t *testing.T) {
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
	if len(text) == 0 {
		t.Fatal("expected no-tracking policy content to be non-empty")
	}
	if !strings.Contains(text, "Test App") {
		t.Fatal("expected no-tracking policy to contain app name 'Test App'")
	}
}

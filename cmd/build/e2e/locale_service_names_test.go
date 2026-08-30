package e2e

import (
	"strings"
	"testing"
)

func TestLocaleServiceNames_TpsNameReturnsLocaleSpecificName(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep7(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 7); err != nil {
		t.Fatal(err)
	}

	text, err := getElementText(ctx, `label[for="list-switch-Google Play Services"]`)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text, "Google Play Services") {
		t.Fatalf("expected label to contain 'Google Play Services', got '%s'", text)
	}
}

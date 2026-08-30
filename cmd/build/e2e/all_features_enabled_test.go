package e2e

import (
	"strings"
	"testing"
	"time"
)

func TestAllFeaturesEnabled_IncludesLocationAIAndDataDeletion(t *testing.T) {
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

	if err := clickSelector(ctx, `label[for="locationcheckbox"]`); err != nil {
		t.Fatal(err)
	}
	if err := clickSelector(ctx, `label[for="aicheckbox"]`); err != nil {
		t.Fatal(err)
	}
	if err := clickSelector(ctx, `label[for="datadeletioncheckbox"]`); err != nil {
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
	if !strings.Contains(text, "Location") {
		t.Fatal("expected policy to contain 'Location'")
	}
	if !strings.Contains(text, "Artificial Intelligence") {
		t.Fatal("expected policy to contain 'Artificial Intelligence'")
	}
	if !strings.Contains(text, "Data Deletion") {
		t.Fatal("expected policy to contain 'Data Deletion'")
	}
}

package e2e

import (
	"strings"
	"testing"
	"time"
)

func TestSpecialCharacters_HandlesQuotesAndAmpersandsInAppName(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}

	if err := sendKeys(ctx, `#appName`, `Tom's "App" & More`); err != nil {
		t.Fatal(err)
	}
	if err := sendKeys(ctx, `#appContact`, "test@example.com"); err != nil {
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
	if !strings.Contains(text, `Tom's "App" & More`) {
		t.Fatal("expected policy to contain app name with quotes and ampersand")
	}
}

func TestSpecialCharacters_HandlesQuotesAndAmpersandInDevName(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}

	if err := sendKeys(ctx, `#appName`, "Test App"); err != nil {
		t.Fatal(err)
	}
	if err := sendKeys(ctx, `#appContact`, "test@example.com"); err != nil {
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

	if err := sendKeys(ctx, `#devName`, `Dev "Script" & Co`); err != nil {
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
	if !strings.Contains(text, `Dev "Script" & Co`) {
		t.Fatal("expected policy to contain dev name with quotes and ampersand")
	}
}

func TestSpecialCharacters_HandlesEmojiInAppName(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}

	if err := sendKeys(ctx, `#appName`, "My App 🚀"); err != nil {
		t.Fatal(err)
	}
	if err := sendKeys(ctx, `#appContact`, "test@example.com"); err != nil {
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
	if !strings.Contains(text, "My App 🚀") {
		t.Fatal("expected policy to contain app name with emoji")
	}
}

func TestSpecialCharacters_HandlesLongStrings(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	longName := strings.Repeat("A", 200)

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}

	if err := sendKeys(ctx, `#appName`, longName); err != nil {
		t.Fatal(err)
	}
	if err := sendKeys(ctx, `#appContact`, "test@example.com"); err != nil {
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
	if !strings.Contains(text, longName) {
		t.Fatal("expected policy to contain long app name")
	}
}

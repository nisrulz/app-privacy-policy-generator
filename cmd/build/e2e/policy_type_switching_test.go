package e2e

import (
	"testing"
	"time"

	"github.com/chromedp/chromedp"
)

func TestPolicyTypeSwitching_SimpleToGDPRChangesModal(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}
	if err := fillStep2(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`input[type="radio"][value="1"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if err := clickNext(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 3); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#businessAddress`, "123 EU Street", chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if err := clickNext(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 4); err != nil {
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

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#devName`, "John Doe", chromedp.ByQuery),
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

	// Verify Simple Privacy Policy modal is available
	if err := clickStep8Button(ctx, "Privacy Policy"); err != nil {
		t.Fatal(err)
	}
	if err := chromedp.Run(ctx,
		chromedp.Sleep(500*time.Millisecond),
	); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, `.modal.is-active`) {
		t.Fatal("expected modal to be active")
	}
	if !elementExists(ctx, `#privacy_simple_content`) {
		t.Fatal("expected privacy_simple_content to be visible")
	}

	// Close modal
	if err := chromedp.Run(ctx,
		chromedp.Click(`.modal.is-active button[aria-label="close"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if err := chromedp.Run(ctx,
		chromedp.Sleep(300*time.Millisecond),
	); err != nil {
		t.Fatal(err)
	}

	// Go back to step 2
	if err := clickStep8Button(ctx, "Previous"); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 7); err != nil {
		t.Fatal(err)
	}
	if err := clickStep8Button(ctx, "Previous"); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 6); err != nil {
		t.Fatal(err)
	}
	if err := clickStep8Button(ctx, "Previous"); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 5); err != nil {
		t.Fatal(err)
	}
	if err := clickStep8Button(ctx, "Previous"); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 4); err != nil {
		t.Fatal(err)
	}
	if err := clickStep8Button(ctx, "Previous"); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 3); err != nil {
		t.Fatal(err)
	}
	if err := clickStep8Button(ctx, "Previous"); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 2); err != nil {
		t.Fatal(err)
	}

	// Switch to GDPR
	if err := chromedp.Run(ctx,
		chromedp.Click(`input[type="radio"][value="3"]`, chromedp.ByQuery),
	); err != nil {
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

	// Now GDPR Privacy Policy modal should be available
	if err := clickStep8Button(ctx, "Privacy Policy"); err != nil {
		t.Fatal(err)
	}
	if err := chromedp.Run(ctx,
		chromedp.Sleep(500*time.Millisecond),
	); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, `.modal.is-active`) {
		t.Fatal("expected modal to be active")
	}
	if !elementExists(ctx, `#privacy_gdpr_content`) {
		t.Fatal("expected privacy_gdpr_content to be visible")
	}
}

func TestPolicyTypeSwitching_SimpleToNoTrackingChangesModal(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep2(ctx); err != nil {
		t.Fatal(err)
	}
	if err := fillStep2(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`input[type="radio"][value="1"]`, chromedp.ByQuery),
	); err != nil {
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

	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#devName`, "John Doe", chromedp.ByQuery),
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

	// Go back to step 2
	for i := 0; i < 6; i++ {
		if err := clickStep8Button(ctx, "Previous"); err != nil {
			t.Fatal(err)
		}
		if err := chromedp.Run(ctx,
			chromedp.Sleep(100*time.Millisecond),
		); err != nil {
			t.Fatal(err)
		}
	}
	if err := expectStep(ctx, 2); err != nil {
		t.Fatal(err)
	}

	// Switch to No Tracking
	if err := chromedp.Run(ctx,
		chromedp.Click(`input[type="radio"][value="2"]`, chromedp.ByQuery),
	); err != nil {
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

	// No Tracking Privacy Policy modal should be available
	if err := clickStep8Button(ctx, "Privacy Policy"); err != nil {
		t.Fatal(err)
	}
	if err := chromedp.Run(ctx,
		chromedp.Sleep(500*time.Millisecond),
	); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, `.modal.is-active`) {
		t.Fatal("expected modal to be active")
	}
	if !elementExists(ctx, `#privacy_notrack_content`) {
		t.Fatal("expected privacy_notrack_content to be visible")
	}
}

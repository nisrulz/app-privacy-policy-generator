package e2e

import (
	"strings"
	"testing"

	"github.com/chromedp/chromedp"
)

func TestStep4ShowsAllAppTypeRadioOptions(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep4(ctx); err != nil {
		t.Fatal(err)
	}

	options := []string{"Free", "Open Source", "Freemium", "Ad Supported", "Commercial"}
	for _, opt := range options {
		selector := `input[type="radio"][value="` + opt + `"]`
		if !elementExists(ctx, selector) {
			t.Errorf("expected radio option %q to be visible", opt)
		}
	}
}

func TestFreeIsSelectedByDefault(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep4(ctx); err != nil {
		t.Fatal(err)
	}

	val, err := getElementAttribute(ctx, `input[type="radio"][value="Free"]`, "checked")
	if err != nil {
		t.Fatal(err)
	}
	if val == "" {
		t.Error("expected Free to be selected by default")
	}
}

func TestCanSelectDifferentAppTypesAndTagUpdates(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep4(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`input[type="radio"][value="Commercial"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	val, err := getElementAttribute(ctx, `input[type="radio"][value="Commercial"]`, "checked")
	if err != nil {
		t.Fatal(err)
	}
	if val == "" {
		t.Error("expected Commercial to be checked")
	}

	val, err = getElementAttribute(ctx, `input[type="radio"][value="Free"]`, "checked")
	if err != nil {
		t.Fatal(err)
	}
	if val != "" {
		t.Error("expected Free to not be checked")
	}

	tagText, err := getElementText(ctx, "#step-4 .tag.is-primary.is-light")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(tagText, "Commercial") {
		t.Errorf("expected tag to contain 'Commercial', got %q", tagText)
	}
}

func TestAppTypeAffectsPlatformDescriptionInPolicy(t *testing.T) {
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

	if err := chromedp.Run(ctx,
		chromedp.Click(`input[type="radio"][value="Commercial"]`, chromedp.ByQuery),
	); err != nil {
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
		chromedp.SendKeys(`#devName`, "Type Test", chromedp.ByQuery),
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

	if err := chromedp.Run(ctx,
		chromedp.Click(`a.card-footer-item >> visible=true`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	content, err := getElementText(ctx, "#privacy_simple_content")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(content, "Type Test") {
		t.Error("expected policy to contain 'Type Test'")
	}
}

func TestFeatureCheckboxLabelsAreVisibleOnStep4(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep4(ctx); err != nil {
		t.Fatal(err)
	}

	features := []string{"locationcheckbox", "aicheckbox", "datadeletioncheckbox"}
	for _, f := range features {
		selector := `label[for="` + f + `"]`
		if !elementExists(ctx, selector) {
			t.Errorf("expected label for %q to be visible", f)
		}
	}
}

func TestFeatureCheckboxesAreUncheckedByDefault(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep4(ctx); err != nil {
		t.Fatal(err)
	}

	features := []string{"locationcheckbox", "aicheckbox", "datadeletioncheckbox"}
	for _, f := range features {
		val, err := getElementAttribute(ctx, "#"+f, "checked")
		if err != nil {
			t.Fatal(err)
		}
		if val != "" {
			t.Errorf("expected %s to be unchecked by default", f)
		}
	}
}

func TestCanToggleFeatureCheckboxesViaLabels(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep4(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`label[for="locationcheckbox"]`, chromedp.ByQuery),
		chromedp.Click(`label[for="aicheckbox"]`, chromedp.ByQuery),
		chromedp.Click(`label[for="datadeletioncheckbox"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	features := []string{"locationcheckbox", "aicheckbox", "datadeletioncheckbox"}
	for _, f := range features {
		val, err := getElementAttribute(ctx, "#"+f, "checked")
		if err != nil {
			t.Fatal(err)
		}
		if val == "" {
			t.Errorf("expected %s to be checked after clicking label", f)
		}
	}
}

func TestLocationTrackingAddsLocationSectionInPolicy(t *testing.T) {
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

	if err := chromedp.Run(ctx,
		chromedp.Click(`label[for="locationcheckbox"]`, chromedp.ByQuery),
	); err != nil {
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
		chromedp.SendKeys(`#devName`, "Feature Test", chromedp.ByQuery),
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

	if err := chromedp.Run(ctx,
		chromedp.Click(`a.card-footer-item >> visible=true`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	content, err := getElementText(ctx, "#privacy_simple_content")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(content, "Location Information") {
		t.Error("expected policy to contain 'Location Information'")
	}
}

func TestAIUsageAddsAISectionInPolicy(t *testing.T) {
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

	if err := chromedp.Run(ctx,
		chromedp.Click(`label[for="aicheckbox"]`, chromedp.ByQuery),
	); err != nil {
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
		chromedp.SendKeys(`#devName`, "Feature Test", chromedp.ByQuery),
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

	if err := chromedp.Run(ctx,
		chromedp.Click(`a.card-footer-item >> visible=true`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	content, err := getElementText(ctx, "#privacy_simple_content")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(content, "Artificial Intelligence") {
		t.Error("expected policy to contain 'Artificial Intelligence'")
	}
}

func TestDataDeletionAddsDeletionSectionInPolicy(t *testing.T) {
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

	if err := chromedp.Run(ctx,
		chromedp.Click(`label[for="datadeletioncheckbox"]`, chromedp.ByQuery),
	); err != nil {
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
		chromedp.SendKeys(`#devName`, "Feature Test", chromedp.ByQuery),
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

	if err := chromedp.Run(ctx,
		chromedp.Click(`a.card-footer-item >> visible=true`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	content, err := getElementText(ctx, "#privacy_simple_content")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(content, "Data Deletion") {
		t.Error("expected policy to contain 'Data Deletion'")
	}
}

func TestFeatureCheckboxesHiddenForNoTrackingPolicy(t *testing.T) {
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
	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#appName`, "Test App", chromedp.ByQuery),
		chromedp.SendKeys(`#appContact`, "test@example.com", chromedp.ByQuery),
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

	features := []string{"locationcheckbox", "aicheckbox", "datadeletioncheckbox"}
	for _, f := range features {
		selector := `label[for="` + f + `"]`
		if elementExists(ctx, selector) {
			t.Errorf("expected label for %q to be hidden for No Tracking policy", f)
		}
	}
}

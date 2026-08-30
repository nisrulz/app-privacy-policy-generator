package e2e

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
)

func selectPlatforms(ctx context.Context, platforms map[string]bool) error {
	defaults := map[string]bool{"android": true, "ios": false, "kaios": false, "windows": false, "web": false}
	for k, v := range platforms {
		defaults[k] = v
	}
	for key, shouldBeChecked := range defaults {
		selector := "#platform-" + key
		val, err := getElementAttribute(ctx, selector, "checked")
		if err != nil {
			return err
		}
		isChecked := val != ""
		if shouldBeChecked != isChecked {
			if err := chromedp.Run(ctx,
				chromedp.Click(`label[for="platform-`+key+`"]`, chromedp.ByQuery),
			); err != nil {
				return err
			}
		}
	}
	return nil
}

func completeToStep8(ctx context.Context) error {
	if err := clickNext(ctx); err != nil {
		return err
	}
	if err := expectStep(ctx, 6); err != nil {
		return err
	}
	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#devName`, "Test Dev", chromedp.ByQuery),
	); err != nil {
		return err
	}
	if err := clickNext(ctx); err != nil {
		return err
	}
	if err := expectStep(ctx, 7); err != nil {
		return err
	}
	if err := clickNext(ctx); err != nil {
		return err
	}
	return expectStep(ctx, 8)
}

func openPrivacyModal(ctx context.Context) error {
	return chromedp.Run(ctx,
		chromedp.Click(`a.card-footer-item >> visible=true`, chromedp.ByQuery),
		chromedp.Sleep(500*time.Millisecond),
	)
}

func getSimplePolicyText(ctx context.Context) (string, error) {
	if err := openPrivacyModal(ctx); err != nil {
		return "", err
	}
	return getElementText(ctx, "#privacy_simple_content")
}

func getGDPRPolicyText(ctx context.Context) (string, error) {
	if err := openPrivacyModal(ctx); err != nil {
		return "", err
	}
	return getElementText(ctx, "#privacy_gdpr_content")
}

func getTandCText(ctx context.Context) (string, error) {
	if err := chromedp.Run(ctx,
		chromedp.Click(`a.card-footer-item >> visible=true`, chromedp.ByQuery),
		chromedp.Sleep(500*time.Millisecond),
	); err != nil {
		return "", err
	}
	return getElementText(ctx, "#tandc_content")
}

func TestWebOnlyShowsWebBrowsersTextInPolicy(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep5(ctx); err != nil {
		t.Fatal(err)
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

	if err := openPrivacyModal(ctx); err != nil {
		t.Fatal(err)
	}
	content, err := getElementText(ctx, "#privacy_simple_content")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(content, "web browsers") {
		t.Error("expected policy to contain 'web browsers'")
	}
}

func TestMultiplePlatformsJoinedWithCommasAndConjunction(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep5(ctx); err != nil {
		t.Fatal(err)
	}
	if err := chromedp.Run(ctx,
		chromedp.Click(`label[for="platform-ios"]`, chromedp.ByQuery),
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

	if err := openPrivacyModal(ctx); err != nil {
		t.Fatal(err)
	}
	content, err := getElementText(ctx, "#privacy_simple_content")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(content, "mobile devices") {
		t.Error("expected policy to contain 'mobile devices'")
	}
	if !strings.Contains(content, "web browsers") {
		t.Error("expected policy to contain 'web browsers'")
	}
}

func TestMobileOnlyAndroidIOSShowsForMobileDevicesWithoutLeadingComma(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep5(ctx); err != nil {
		t.Fatal(err)
	}
	if err := chromedp.Run(ctx,
		chromedp.Click(`label[for="platform-ios"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if err := clickNext(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 6); err != nil {
		t.Fatal(err)
	}
	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#devName`, "Dev Name", chromedp.ByQuery),
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

	if err := openPrivacyModal(ctx); err != nil {
		t.Fatal(err)
	}
	content, err := getElementText(ctx, "#privacy_simple_content")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(content, "for mobile devices, together") {
		t.Error("expected policy to contain 'for mobile devices, together'")
	}
	if strings.Contains(content, ", and mobile devices") {
		t.Error("policy should not contain ', and mobile devices'")
	}
	if strings.Contains(content, ", or mobile devices") {
		t.Error("policy should not contain ', or mobile devices'")
	}
}

func TestNoTrackingMobileOnlyUsesMobileDeviceWithoutLeadingComma(t *testing.T) {
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
		chromedp.SendKeys(`#appName`, "NoTrack App", chromedp.ByQuery),
		chromedp.SendKeys(`#appContact`, "notrack@example.com", chromedp.ByQuery),
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

	if err := chromedp.Run(ctx,
		chromedp.Click(`label[for="platform-ios"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if err := clickNext(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 6); err != nil {
		t.Fatal(err)
	}
	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#devName`, "NoTrack Dev", chromedp.ByQuery),
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

	if err := openPrivacyModal(ctx); err != nil {
		t.Fatal(err)
	}
	content, err := getElementText(ctx, "#privacy_notrack_content")
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(content, ", or mobile device") {
		t.Error("policy should not contain ', or mobile device'")
	}
	if !strings.Contains(content, "your mobile device") {
		t.Error("expected policy to contain 'your mobile device'")
	}
}

func TestWindowsOnlyShowsWindowsDevicesWithoutLeadingComma(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep5(ctx); err != nil {
		t.Fatal(err)
	}
	if err := chromedp.Run(ctx,
		chromedp.Click(`label[for="platform-android"]`, chromedp.ByQuery),
		chromedp.Click(`label[for="platform-windows"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}
	if err := clickNext(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 6); err != nil {
		t.Fatal(err)
	}
	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#devName`, "Dev Name", chromedp.ByQuery),
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

	if err := openPrivacyModal(ctx); err != nil {
		t.Fatal(err)
	}
	content, err := getElementText(ctx, "#privacy_simple_content")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(content, "for Windows devices, together") {
		t.Error("expected policy to contain 'for Windows devices, together'")
	}
	if strings.Contains(content, ", and Windows devices") {
		t.Error("policy should not contain ', and Windows devices'")
	}
}

func TestSinglePlatformTextHasNoStrayCommasOrConjunctions(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizard(ctx); err != nil {
		t.Fatal(err)
	}

	if err := openPrivacyModal(ctx); err != nil {
		t.Fatal(err)
	}
	content, err := getElementText(ctx, "#privacy_simple_content")
	if err != nil {
		t.Fatal(err)
	}

	// Should not start with comma+conjunction like ", and mobile" or ", or mobile"
	if strings.Contains(content, ", and mobile") || strings.Contains(content, ", or mobile") {
		t.Error("policy should not contain stray comma+conjunction at start")
	}
}

func TestWebOnlyIntroSaysForWebBrowsers(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep5(ctx); err != nil {
		t.Fatal(err)
	}
	if err := selectPlatforms(ctx, map[string]bool{"android": false, "web": true}); err != nil {
		t.Fatal(err)
	}
	if err := completeToStep8(ctx); err != nil {
		t.Fatal(err)
	}

	text, err := getSimplePolicyText(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text, "for web browsers, together") {
		t.Error("expected text to contain 'for web browsers, together'")
	}
	if strings.Contains(text, ", and web") {
		t.Error("text should not contain ', and web'")
	}
}

func TestMobileWindowsUsesAndInIntro(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep5(ctx); err != nil {
		t.Fatal(err)
	}
	if err := selectPlatforms(ctx, map[string]bool{"ios": true, "windows": true}); err != nil {
		t.Fatal(err)
	}
	if err := completeToStep8(ctx); err != nil {
		t.Fatal(err)
	}

	text, err := getSimplePolicyText(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text, "for mobile devices and Windows devices") {
		t.Error("expected text to contain 'for mobile devices and Windows devices'")
	}
}

func TestMobileWebUsesAndInIntro(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep5(ctx); err != nil {
		t.Fatal(err)
	}
	if err := selectPlatforms(ctx, map[string]bool{"ios": true, "web": true}); err != nil {
		t.Fatal(err)
	}
	if err := completeToStep8(ctx); err != nil {
		t.Fatal(err)
	}

	text, err := getSimplePolicyText(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text, "for mobile devices and web browsers") {
		t.Error("expected text to contain 'for mobile devices and web browsers'")
	}
}

func TestWindowsWebUsesAndInIntro(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep5(ctx); err != nil {
		t.Fatal(err)
	}
	if err := selectPlatforms(ctx, map[string]bool{"android": false, "windows": true, "web": true}); err != nil {
		t.Fatal(err)
	}
	if err := completeToStep8(ctx); err != nil {
		t.Fatal(err)
	}

	text, err := getSimplePolicyText(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text, "for Windows devices and web browsers") {
		t.Error("expected text to contain 'for Windows devices and web browsers'")
	}
}

func TestAllThreePlatformsUsesCommasAndFinalAnd(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep5(ctx); err != nil {
		t.Fatal(err)
	}
	if err := selectPlatforms(ctx, map[string]bool{"ios": true, "windows": true, "web": true}); err != nil {
		t.Fatal(err)
	}
	if err := completeToStep8(ctx); err != nil {
		t.Fatal(err)
	}

	text, err := getSimplePolicyText(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text, "mobile devices, Windows devices, and web browsers") {
		t.Error("expected text to contain 'mobile devices, Windows devices, and web browsers'")
	}
}

func TestWebOnlyNoTrackingUsesComputerForLocation(t *testing.T) {
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
		chromedp.SendKeys(`#appName`, "NoTrack Web", chromedp.ByQuery),
		chromedp.SendKeys(`#appContact`, "web@example.com", chromedp.ByQuery),
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

	if err := selectPlatforms(ctx, map[string]bool{"android": false, "web": true}); err != nil {
		t.Fatal(err)
	}
	if err := clickNext(ctx); err != nil {
		t.Fatal(err)
	}
	if err := expectStep(ctx, 6); err != nil {
		t.Fatal(err)
	}
	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#devName`, "Web Dev", chromedp.ByQuery),
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

	if err := openPrivacyModal(ctx); err != nil {
		t.Fatal(err)
	}
	content, err := getElementText(ctx, "#privacy_notrack_content")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(content, "your computer") {
		t.Error("expected no-tracking web policy to contain 'your computer'")
	}
	if strings.Contains(content, ", and computer") || strings.Contains(content, ", or computer") ||
		strings.Contains(content, ", and mobile") || strings.Contains(content, ", or mobile") {
		t.Error("no-tracking web policy should not contain stray comma+conjunction")
	}
}

func TestGDPRMobileOnlyIntroUsesForMobileDevices(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizardGDPR(ctx, ""); err != nil {
		t.Fatal(err)
	}
	if err := selectPlatforms(ctx, map[string]bool{"ios": true}); err != nil {
		t.Fatal(err)
	}
	if err := completeToStep8(ctx); err != nil {
		t.Fatal(err)
	}

	text, err := getGDPRPolicyText(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text, "for mobile devices, together") {
		t.Error("expected text to contain 'for mobile devices, together'")
	}
	if strings.Contains(text, ", and mobile devices") {
		t.Error("text should not contain ', and mobile devices'")
	}
}

func TestGDPRAutoCollectUsesCorrectDeviceType(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizardGDPR(ctx, ""); err != nil {
		t.Fatal(err)
	}
	if err := selectPlatforms(ctx, map[string]bool{"ios": true, "windows": true}); err != nil {
		t.Fatal(err)
	}
	if err := completeToStep8(ctx); err != nil {
		t.Fatal(err)
	}

	text, err := getGDPRPolicyText(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text, "mobile device or Windows device") {
		t.Error("expected text to contain 'mobile device or Windows device'")
	}
}

func TestTandCMobileOnlyIntroUsesForMobileDevices(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep5(ctx); err != nil {
		t.Fatal(err)
	}
	if err := selectPlatforms(ctx, map[string]bool{"ios": true}); err != nil {
		t.Fatal(err)
	}
	if err := completeToStep8(ctx); err != nil {
		t.Fatal(err)
	}

	text, err := getTandCText(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text, "for mobile devices, together") {
		t.Error("expected text to contain 'for mobile devices, together'")
	}
	if strings.Contains(text, ", and mobile devices") {
		t.Error("text should not contain ', and mobile devices'")
	}
}

func TestTandCWebOnlyIntroUsesForWebBrowsers(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := navigateToStep5(ctx); err != nil {
		t.Fatal(err)
	}
	if err := selectPlatforms(ctx, map[string]bool{"android": false, "web": true}); err != nil {
		t.Fatal(err)
	}
	if err := completeToStep8(ctx); err != nil {
		t.Fatal(err)
	}

	text, err := getTandCText(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(text, "for web browsers, together") {
		t.Error("expected text to contain 'for web browsers, together'")
	}
	if strings.Contains(text, ", and web") {
		t.Error("text should not contain ', and web'")
	}
}

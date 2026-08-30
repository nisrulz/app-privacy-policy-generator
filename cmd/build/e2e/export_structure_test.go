package e2e

import (
	"strings"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
)

func TestExportStructure_HTMLExportContainsCompleteDocumentStructure(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizard(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Evaluate(
			`Array.from(document.querySelectorAll('#step-8 a.card-footer-item')).find(el => el.textContent.trim() === 'Privacy Policy').click()`,
			nil,
		),
	); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Sleep(500*time.Millisecond),
	); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, ".modal.is-active") {
		t.Fatal("expected modal to be active")
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`.modal.is-active button[onclick*="getHtml"]`, chromedp.ByQuery),
		chromedp.Sleep(300*time.Millisecond),
	); err != nil {
		t.Fatal(err)
	}

	var html string
	if err := chromedp.Run(ctx,
		chromedp.Value(`#privacy_simple_txtarea`, &html, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(html, "<!DOCTYPE html>") {
		t.Fatal("expected HTML to contain <!DOCTYPE html>")
	}
	if !strings.Contains(html, "<meta charset='utf-8'>") {
		t.Fatal("expected HTML to contain <meta charset='utf-8'>")
	}
	if !strings.Contains(html, "<meta name='viewport'") {
		t.Fatal("expected HTML to contain <meta name='viewport'")
	}
	if !strings.Contains(html, "<title>") {
		t.Fatal("expected HTML to contain <title>")
	}
	if !strings.Contains(html, "<style>") {
		t.Fatal("expected HTML to contain <style>")
	}
	if !strings.Contains(html, "<body>") {
		t.Fatal("expected HTML to contain <body>")
	}
	if !strings.Contains(html, "</body>") {
		t.Fatal("expected HTML to contain </body>")
	}
	if !strings.Contains(html, "</html>") {
		t.Fatal("expected HTML to contain </html>")
	}
}

func TestExportStructure_MarkdownExportContainsValidMarkdownFormatting(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizard(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Evaluate(
			`Array.from(document.querySelectorAll('#step-8 a.card-footer-item')).find(el => el.textContent.trim() === 'Privacy Policy').click()`,
			nil,
		),
	); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Sleep(500*time.Millisecond),
	); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, ".modal.is-active") {
		t.Fatal("expected modal to be active")
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`.modal.is-active button[onclick*="getMarkdown"]`, chromedp.ByQuery),
		chromedp.Sleep(300*time.Millisecond),
	); err != nil {
		t.Fatal(err)
	}

	var md string
	if err := chromedp.Run(ctx,
		chromedp.Value(`#privacy_simple_txtarea`, &md, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	if len(md) <= 50 {
		t.Fatal("expected Markdown content length to be greater than 50")
	}

	hasMarkdown := strings.Contains(md, "**") || strings.Contains(md, "- ") || strings.Contains(md, "#")
	if !hasMarkdown {
		t.Fatal("expected Markdown to contain bold markers, list markers, or headers")
	}
}

func TestExportStructure_TCHTMLExportContainsCompleteDocumentStructure(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := completeWizard(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Evaluate(
			`Array.from(document.querySelectorAll('#step-8 a.card-footer-item')).find(el => el.textContent.trim() === 'Terms & Conditions').click()`,
			nil,
		),
	); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Sleep(500*time.Millisecond),
	); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, ".modal.is-active") {
		t.Fatal("expected modal to be active")
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`.modal.is-active button[onclick*="getHtml"]`, chromedp.ByQuery),
		chromedp.Sleep(300*time.Millisecond),
	); err != nil {
		t.Fatal(err)
	}

	var html string
	if err := chromedp.Run(ctx,
		chromedp.Value(`#tandc_txtarea`, &html, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(html, "<!DOCTYPE html>") {
		t.Fatal("expected HTML to contain <!DOCTYPE html>")
	}
	if !strings.Contains(html, "<meta charset='utf-8'>") {
		t.Fatal("expected HTML to contain <meta charset='utf-8'>")
	}
	if !strings.Contains(html, "<title>") {
		t.Fatal("expected HTML to contain <title>")
	}
	if !strings.Contains(html, "<body>") {
		t.Fatal("expected HTML to contain <body>")
	}
}

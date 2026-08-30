package e2e

import (
	"testing"

	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

func TestAppMount_NoConsoleErrorsOnLoad(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	consoleErrors := []string{}
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		if e, ok := ev.(*runtime.EventConsoleAPICalled); ok {
			if e.Type == "error" {
				for _, arg := range e.Args {
					consoleErrors = append(consoleErrors, string(arg.Value))
				}
			}
		}
	})

	if err := gotoApp(ctx); err != nil {
		t.Fatal(err)
	}

	if len(consoleErrors) > 0 {
		t.Fatalf("console errors on page load: %v", consoleErrors)
	}
}

func TestAppMount_VueIsMounted(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := gotoApp(ctx); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, "#app") {
		t.Fatal("#app element not found")
	}

	if !elementExists(ctx, "#step-1") {
		t.Fatal("wizard step 1 not rendered — Vue app may not be mounted")
	}
}

func TestAppMount_ClickingStartAdvancesWizard(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := gotoApp(ctx); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, "#step-1") {
		t.Fatal("expected step 1 to be visible initially")
	}

	if err := clickStart(ctx); err != nil {
		t.Fatal(err)
	}

	if err := expectStep(ctx, 2); err != nil {
		t.Fatal("clicking Start did not advance to step 2 — reactivity is broken")
	}
}

func TestAppMount_ModalOpensOnLinkClick(t *testing.T) {
	ctx, cancel := newBrowser(t)
	defer cancel()

	if err := gotoApp(ctx); err != nil {
		t.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Click(`a[href="#"][onclick*="toggleFaqModalVisibility"]`, chromedp.ByQuery),
	); err != nil {
		t.Fatal(err)
	}

	if !elementExists(ctx, ".modal.is-active") {
		t.Fatal("clicking FAQ link did not open a modal")
	}
}

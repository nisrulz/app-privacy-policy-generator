package e2e

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

const baseURL = "http://localhost:8000"

func newBrowser(t *testing.T) (context.Context, context.CancelFunc) {
	t.Helper()
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-gpu", true),
	)
	allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(allocCtx)
	return ctx, func() {
		cancel()
		allocCancel()
	}
}

func gotoApp(ctx context.Context) error {
	return chromedp.Run(ctx,
		chromedp.Navigate(baseURL+"/"),
		chromedp.WaitReady("body", chromedp.ByQuery),
	)
}

func clickNext(ctx context.Context) error {
	return chromedp.Run(ctx,
		chromedp.Click(`a.card-footer-item >> visible=true`, chromedp.ByQuery),
		chromedp.Sleep(200*time.Millisecond),
	)
}

func clickPrevious(ctx context.Context) error {
	return chromedp.Run(ctx,
		chromedp.Click(`a.card-footer-item >> visible=true`, chromedp.ByQuery),
		chromedp.Sleep(200*time.Millisecond),
	)
}

func clickStart(ctx context.Context) error {
	return chromedp.Run(ctx,
		chromedp.Click(`button.start-btn`, chromedp.ByQuery),
		chromedp.Sleep(200*time.Millisecond),
	)
}

func expectStep(ctx context.Context, step int) error {
	selector := fmt.Sprintf(`#step-%d`, step)
	return chromedp.Run(ctx,
		chromedp.WaitVisible(selector, chromedp.ByQuery),
	)
}

func navigateToStep2(ctx context.Context) error {
	if err := gotoApp(ctx); err != nil {
		return err
	}
	if err := clickStart(ctx); err != nil {
		return err
	}
	return expectStep(ctx, 2)
}

func fillStep2(ctx context.Context) error {
	return chromedp.Run(ctx,
		chromedp.SendKeys(`#appName`, "Test App", chromedp.ByQuery),
		chromedp.SendKeys(`#appContact`, "test@example.com", chromedp.ByQuery),
	)
}

func navigateToStep4(ctx context.Context) error {
	if err := navigateToStep2(ctx); err != nil {
		return err
	}
	if err := fillStep2(ctx); err != nil {
		return err
	}
	if err := clickNext(ctx); err != nil {
		return err
	}
	if err := expectStep(ctx, 3); err != nil {
		return err
	}
	if err := clickNext(ctx); err != nil {
		return err
	}
	return expectStep(ctx, 4)
}

func navigateToStep5(ctx context.Context) error {
	if err := navigateToStep4(ctx); err != nil {
		return err
	}
	if err := clickNext(ctx); err != nil {
		return err
	}
	return expectStep(ctx, 5)
}

func navigateToStep6(ctx context.Context) error {
	if err := navigateToStep2(ctx); err != nil {
		return err
	}
	if err := fillStep2(ctx); err != nil {
		return err
	}
	for step := 3; step <= 6; step++ {
		if err := clickNext(ctx); err != nil {
			return err
		}
		if err := expectStep(ctx, step); err != nil {
			return err
		}
	}
	return nil
}

func navigateToStep7(ctx context.Context) error {
	if err := navigateToStep6(ctx); err != nil {
		return err
	}
	if err := chromedp.Run(ctx,
		chromedp.Click(`input[type="radio"][value="Individual"]`, chromedp.ByQuery),
		chromedp.SendKeys(`#devName`, "John Doe", chromedp.ByQuery),
	); err != nil {
		return err
	}
	if err := clickNext(ctx); err != nil {
		return err
	}
	return expectStep(ctx, 7)
}

func completeWizard(ctx context.Context) error {
	if err := navigateToStep7(ctx); err != nil {
		return err
	}
	if err := clickNext(ctx); err != nil {
		return err
	}
	return expectStep(ctx, 8)
}

func completeWizardNoTracking(ctx context.Context) error {
	if err := gotoApp(ctx); err != nil {
		return err
	}
	if err := clickStart(ctx); err != nil {
		return err
	}
	if err := expectStep(ctx, 2); err != nil {
		return err
	}
	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#appName`, "Test App", chromedp.ByQuery),
		chromedp.SendKeys(`#appContact`, "test@example.com", chromedp.ByQuery),
		chromedp.Click(`input[type="radio"][value="2"]`, chromedp.ByQuery),
	); err != nil {
		return err
	}
	for step := 3; step <= 6; step++ {
		if err := clickNext(ctx); err != nil {
			return err
		}
		if err := expectStep(ctx, step); err != nil {
			return err
		}
	}
	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#devName`, "John Doe", chromedp.ByQuery),
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

func completeWizardCompany(ctx context.Context) error {
	if err := gotoApp(ctx); err != nil {
		return err
	}
	if err := clickStart(ctx); err != nil {
		return err
	}
	if err := expectStep(ctx, 2); err != nil {
		return err
	}
	if err := fillStep2(ctx); err != nil {
		return err
	}
	for step := 3; step <= 6; step++ {
		if err := clickNext(ctx); err != nil {
			return err
		}
		if err := expectStep(ctx, step); err != nil {
			return err
		}
	}
	if err := chromedp.Run(ctx,
		chromedp.Click(`input[type="radio"][value="Company"]`, chromedp.ByQuery),
		chromedp.SendKeys(`#companyName`, "Acme Corp", chromedp.ByQuery),
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

func completeWizardGDPR(ctx context.Context, businessAddress string) error {
	if err := gotoApp(ctx); err != nil {
		return err
	}
	if err := clickStart(ctx); err != nil {
		return err
	}
	if err := expectStep(ctx, 2); err != nil {
		return err
	}
	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#appName`, "GDPR App", chromedp.ByQuery),
		chromedp.SendKeys(`#appContact`, "gdpr@example.com", chromedp.ByQuery),
		chromedp.Click(`input[type="radio"][value="3"]`, chromedp.ByQuery),
	); err != nil {
		return err
	}
	if err := clickNext(ctx); err != nil {
		return err
	}
	if err := expectStep(ctx, 3); err != nil {
		return err
	}
	if businessAddress == "" {
		businessAddress = "123 EU Street, Berlin"
	}
	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#businessAddress`, businessAddress, chromedp.ByQuery),
	); err != nil {
		return err
	}
	for step := 4; step <= 6; step++ {
		if err := clickNext(ctx); err != nil {
			return err
		}
		if err := expectStep(ctx, step); err != nil {
			return err
		}
	}
	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`#devName`, "GDPR Dev", chromedp.ByQuery),
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

func getElementText(ctx context.Context, selector string) (string, error) {
	var text string
	err := chromedp.Run(ctx,
		chromedp.Text(selector, &text, chromedp.ByQuery),
	)
	return text, err
}

func getElementAttribute(ctx context.Context, selector, attr string) (string, error) {
	var value string
	err := chromedp.Run(ctx,
		chromedp.AttributeValue(selector, attr, &value, nil, chromedp.ByQuery),
	)
	return value, err
}

func elementExists(ctx context.Context, selector string) bool {
	var nodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Nodes(selector, &nodes, chromedp.ByQuery),
	)
	return err == nil && len(nodes) > 0
}

func clickSelector(ctx context.Context, selector string) error {
	return chromedp.Run(ctx,
		chromedp.Click(selector, chromedp.ByQuery),
	)
}

func sendKeys(ctx context.Context, selector, value string) error {
	return chromedp.Run(ctx,
		chromedp.SendKeys(selector, value, chromedp.ByQuery),
	)
}

func sleep(ctx context.Context, d time.Duration) error {
	return chromedp.Run(ctx,
		chromedp.Sleep(d),
	)
}

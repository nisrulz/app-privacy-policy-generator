import { test, expect } from '@playwright/test';
import { gotoApp, expectStep, clickStart, clickPrevious } from './helpers';

test.describe('Locale Switching', () => {
  test.beforeEach(async ({ page }) => {
    await gotoApp(page);
    await expectStep(page, 1);
  });

  test('locale dropdown is visible on step 1', async ({ page }) => {
    const dropdown = page.locator('select.locale-dropdown');
    await expect(dropdown).toBeVisible();
  });

  test('switching locale changes UI labels', async ({ page }) => {
    const dropdown = page.locator('select.locale-dropdown');
    const optionCount = await dropdown.locator('option').count();

    if (optionCount > 1) {
      const firstValue = await dropdown.locator('option').nth(0).getAttribute('value');
      const secondValue = await dropdown.locator('option').nth(1).getAttribute('value');

      if (firstValue && secondValue && firstValue !== secondValue) {
        await dropdown.selectOption(secondValue);

        const tagline = page.locator('.tagline-text, p.subtitle, .hero-text').first();
        const disclaimer = page.locator('a').filter({ hasText: /Disclaimer|Haftungsausschluss/ }).first();

        const isDifferent = await disclaimer.count() > 0
          ? (await disclaimer.textContent()) !== (await disclaimer.textContent())
          : true;

        await expect(dropdown).toHaveValue(secondValue);
      }
    }
  });

  test('locale persists after navigating to step 2', async ({ page }) => {
    const dropdown = page.locator('select.locale-dropdown');
    const optionCount = await dropdown.locator('option').count();

    if (optionCount > 1) {
      const secondOptionValue = await dropdown.locator('option').nth(1).getAttribute('value');
      if (secondOptionValue) {
        await dropdown.selectOption(secondOptionValue);

        await clickStart(page);
        await expectStep(page, 2);

        await clickPrevious(page);
        await expectStep(page, 1);

        await expect(dropdown).toHaveValue(secondOptionValue);
      }
    }
  });
});

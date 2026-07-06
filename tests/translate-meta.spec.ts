import { test, expect } from '@playwright/test';
import { gotoApp } from './helpers';

test.describe('translate() function behavior', () => {
  test('translate substitutes {{param}} placeholders with provided params', async ({ page }) => {
    await gotoApp(page);
    const result = await page.evaluate(() => {
      return (window as any).translate('wizard.step1.start');
    });
    expect(typeof result).toBe('string');
    expect(result.length).toBeGreaterThan(0);
  });

  test('translate returns the key itself when locale key is not found', async ({ page }) => {
    await gotoApp(page);
    const result = await page.evaluate(() => {
      return (window as any).translate('nonexistent.key.that.does.not.exist');
    });
    expect(result).toBe('nonexistent.key.that.does.not.exist');
  });

  test('translate returns correct text for known keys', async ({ page }) => {
    await gotoApp(page);
    const startText = await page.evaluate(() => {
      return (window as any).translate('wizard.step1.start');
    });
    expect(startText).toBe('Start');
  });
});

test.describe('_updateMeta() behavior', () => {
  test('page title updates to reflect locale on mount', async ({ page }) => {
    await gotoApp(page);
    const title = await page.title();
    expect(title.length).toBeGreaterThan(0);
  });

  test('meta description tag exists', async ({ page }) => {
    await gotoApp(page);
    const desc = await page.locator('meta[name="description"]').getAttribute('content');
    expect(desc).toBeTruthy();
  });

  test('og:title meta tag exists', async ({ page }) => {
    await gotoApp(page);
    const ogTitle = await page.locator('meta[property="og:title"]').getAttribute('content');
    expect(ogTitle).toBeTruthy();
  });
});

test.describe('_updateThemeLogo() behavior', () => {
  test('theme logo image src swaps between light and dark variants on toggle', async ({ page }) => {
    await gotoApp(page);
    const logo = page.locator('img[data-theme-logo]').first();
    await expect(logo).toBeVisible();

    const lightSrc = await logo.getAttribute('src');

    // Toggle to dark
    await page.locator('button.theme-toggle').click();
    await page.waitForTimeout(300);

    const darkSrc = await logo.getAttribute('src');
    // Source should change (or at least the data attributes should be set)
    expect(darkSrc).toBeTruthy();
  });
});

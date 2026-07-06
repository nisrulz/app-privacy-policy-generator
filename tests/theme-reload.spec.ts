import { test, expect } from '@playwright/test';
import { gotoApp } from './helpers';

test.describe('Theme persistence across page reload', () => {
  test('dark theme persists after page reload', async ({ page }) => {
    await gotoApp(page);
    await expect(page.locator('html')).toHaveAttribute('data-theme', 'light');

    // Toggle to dark
    await page.locator('button.theme-toggle').click();
    await expect(page.locator('html')).toHaveAttribute('data-theme', 'dark');

    // Reload and verify
    await page.reload();
    await page.waitForLoadState('networkidle');
    await expect(page.locator('html')).toHaveAttribute('data-theme', 'dark');
  });

  test('light theme persists after page reload', async ({ page }) => {
    await gotoApp(page);
    // Toggle to dark then back to light
    await page.locator('button.theme-toggle').click();
    await page.locator('button.theme-toggle').click();
    await expect(page.locator('html')).toHaveAttribute('data-theme', 'light');

    // Reload and verify
    await page.reload();
    await page.waitForLoadState('networkidle');
    await expect(page.locator('html')).toHaveAttribute('data-theme', 'light');
  });

  test('theme toggle button text updates correctly', async ({ page }) => {
    await gotoApp(page);
    const btn = page.locator('button.theme-toggle');
    // Default is light, button should show moon emoji
    await expect(btn).toHaveText('🌙');
    await btn.click();
    await expect(btn).toHaveText('☀️');
    await btn.click();
    await expect(btn).toHaveText('🌙');
  });
});

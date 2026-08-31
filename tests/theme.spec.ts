import { test, expect } from '@playwright/test';
import { gotoApp, expectStep, clickStart } from './helpers';

test.describe('Theme Toggle', () => {
  test.beforeEach(async ({ page }) => {
    await gotoApp(page);
    await expectStep(page, 1);
  });

  test('theme toggle button is visible', async ({ page }) => {
    await expect(page.locator('button.theme-toggle')).toBeVisible();
  });

  test('default theme is light', async ({ page }) => {
    await expect(page.locator('html')).toHaveAttribute('data-theme', 'light');
  });

  test('toggles to dark mode on click', async ({ page }) => {
    await page.locator('button.theme-toggle').click();
    await expect(page.locator('html')).toHaveAttribute('data-theme', 'dark');
  });

  test('toggles back to light mode on second click', async ({ page }) => {
    await page.locator('button.theme-toggle').click();
    await expect(page.locator('html')).toHaveAttribute('data-theme', 'dark');

    await page.locator('button.theme-toggle').click();
    await expect(page.locator('html')).toHaveAttribute('data-theme', 'light');
  });

  test('theme persists after navigating steps', async ({ page }) => {
    await page.locator('button.theme-toggle').click();
    await expect(page.locator('html')).toHaveAttribute('data-theme', 'dark');

    await clickStart(page);
    await expectStep(page, 2);
    await expect(page.locator('html')).toHaveAttribute('data-theme', 'dark');
  });
});

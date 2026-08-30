import { test, expect } from '@playwright/test';
import { completeWizard, clickNext, expectStep, navigateToStep2, fillStep2 } from './helpers';

test.describe('Modal reset to Preview mode on reopen', () => {
  test('reopening privacy modal resets to Preview from HTML view', async ({ page }) => {
    await completeWizard(page);

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await page.waitForTimeout(500);
    await expect(page.locator('.modal.is-active')).toBeVisible();

    await page.locator('.modal.is-active button').filter({ hasText: 'HTML' }).click();
    await page.waitForTimeout(300);
    await expect(page.locator('#privacy_simple_txtarea')).toBeVisible();

    await page.locator('.modal.is-active button.delete').click();
    await page.waitForTimeout(300);
    await expect(page.locator('.modal.is-active')).toHaveCount(0);

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await page.waitForTimeout(500);
    await expect(page.locator('.modal.is-active')).toBeVisible();
    await expect(page.locator('#privacy_simple_content')).toBeVisible();
  });

  test('reopening T&C modal resets to Preview from Markdown view', async ({ page }) => {
    await completeWizard(page);

    await page.locator('a.card-footer-item').filter({ hasText: 'Terms' }).click();
    await page.waitForTimeout(500);
    await expect(page.locator('.modal.is-active')).toBeVisible();

    await page.locator('.modal.is-active button').filter({ hasText: 'Markdown' }).click();
    await page.waitForTimeout(300);
    await expect(page.locator('#tandc_txtarea')).toBeVisible();

    await page.locator('.modal.is-active button.delete').click();
    await page.waitForTimeout(300);

    await page.locator('a.card-footer-item').filter({ hasText: 'Terms' }).click();
    await page.waitForTimeout(500);
    await expect(page.locator('#tandc_content')).toBeVisible();
  });
});

import { test, expect } from '@playwright/test';
import { completeWizard, completeWizardNoTracking } from './helpers';

test.describe('Children section varies by third-party selection', () => {
  test('simple policy shows different children text with vs without services', async ({ page }) => {
    await completeWizardNoTracking(page);

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await page.waitForTimeout(500);
    await expect(page.locator('.modal.is-active')).toBeVisible();
    const content = await page.locator('#privacy_notrack_content').innerText();
    expect(content).toContain('Children');
  });

  test('policy includes app name in children section context', async ({ page }) => {
    await completeWizard(page);

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await page.waitForTimeout(500);
    await expect(page.locator('.modal.is-active')).toBeVisible();
    const content = await page.locator('#privacy_simple_content').innerText();
    expect(content).toContain('Children');
    expect(content).toContain('Test App');
  });
});

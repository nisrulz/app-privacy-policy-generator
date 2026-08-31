import { test, expect } from '@playwright/test';
import { completeWizardNoTracking } from './helpers';

test.describe('No Tracking policy data deletion', () => {
  test('No Tracking privacy policy renders with data deletion text', async ({ page }) => {
    await completeWizardNoTracking(page);

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await page.waitForTimeout(500);
    await expect(page.locator('.modal.is-active')).toBeVisible();
    const content = await page.locator('#privacy_notrack_content').innerText();
    expect(content.length).toBeGreaterThan(0);
    expect(content).toContain('Test App');
  });
});

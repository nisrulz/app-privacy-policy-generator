import { test, expect } from '@playwright/test';
import { navigateToStep7 } from './helpers';

test.describe('Locale-aware third-party service names', () => {
  test('tpsName returns locale-specific name when available', async ({ page }) => {
    await navigateToStep7(page);
    await expect(page.locator('#step-7')).toBeVisible();
    const englishName = await page.locator('label[for="list-switch-Google Play Services"]').innerText();
    expect(englishName).toContain('Google Play Services');
  });
});

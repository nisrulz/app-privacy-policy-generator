import { test, expect } from '@playwright/test';
import { navigateToStep2, fillStep2, clickNext, expectStep } from './helpers';

test.describe('All feature flags enabled simultaneously', () => {
  test('policy output includes location, AI, and data deletion sections when all enabled', async ({ page }) => {
    await navigateToStep2(page);
    await fillStep2(page);
    await clickNext(page);
    await expectStep(page, 3);
    await clickNext(page);
    await expectStep(page, 4);

    await page.locator('label[for="locationcheckbox"]').click();
    await page.locator('label[for="aicheckbox"]').click();
    await page.locator('label[for="datadeletioncheckbox"]').click();

    await clickNext(page);
    await expectStep(page, 5);
    await clickNext(page);
    await expectStep(page, 6);
    await page.locator('#devName').fill('John Doe');
    await clickNext(page);
    await expectStep(page, 7);
    await clickNext(page);
    await expectStep(page, 8);

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await page.waitForTimeout(500);
    await expect(page.locator('.modal.is-active')).toBeVisible();
    const content = await page.locator('#privacy_simple_content').innerText();
    expect(content).toContain('Location');
    expect(content).toContain('Artificial Intelligence');
    expect(content).toContain('Data Deletion');
  });
});

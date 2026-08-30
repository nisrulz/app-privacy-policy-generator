import { test, expect } from '@playwright/test';
import { gotoApp, clickStart, clickNext, expectStep, navigateToStep2, fillStep2 } from './helpers';

test.describe('Policy type switching mid-wizard', () => {
  test('switching from Simple to GDPR changes which modal appears at step 8', async ({ page }) => {
    await navigateToStep2(page);
    await fillStep2(page);
    await page.locator('input[type="radio"][value="1"]').check();
    await clickNext(page);
    await expectStep(page, 3);
    await page.locator('#businessAddress').fill('123 EU Street');
    await clickNext(page);
    await expectStep(page, 4);
    await clickNext(page);
    await expectStep(page, 5);
    await clickNext(page);
    await expectStep(page, 6);
    await page.locator('#devName').fill('John Doe');
    await clickNext(page);
    await expectStep(page, 7);
    await clickNext(page);
    await expectStep(page, 8);

    // Verify Simple Privacy Policy modal is available
    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await page.waitForTimeout(500);
    await expect(page.locator('.modal.is-active')).toBeVisible();
    await expect(page.locator('#privacy_simple_content')).toBeVisible();

    // Close modal
    await page.locator('.modal.is-active button.delete[aria-label="close"]').click();
    await page.waitForTimeout(300);

    // Go back to step 2 and switch to GDPR
    await page.locator('a.card-footer-item').filter({ hasText: 'Previous' }).click();
    await expectStep(page, 7);
    await page.locator('a.card-footer-item').filter({ hasText: 'Previous' }).click();
    await expectStep(page, 6);
    await page.locator('a.card-footer-item').filter({ hasText: 'Previous' }).click();
    await expectStep(page, 5);
    await page.locator('a.card-footer-item').filter({ hasText: 'Previous' }).click();
    await expectStep(page, 4);
    await page.locator('a.card-footer-item').filter({ hasText: 'Previous' }).click();
    await expectStep(page, 3);
    await page.locator('a.card-footer-item').filter({ hasText: 'Previous' }).click();
    await expectStep(page, 2);

    // Switch to GDPR
    await page.locator('input[type="radio"][value="3"]').check();
    await clickNext(page);
    await expectStep(page, 3);
    await clickNext(page);
    await expectStep(page, 4);
    await clickNext(page);
    await expectStep(page, 5);
    await clickNext(page);
    await expectStep(page, 6);
    await clickNext(page);
    await expectStep(page, 7);
    await clickNext(page);
    await expectStep(page, 8);

    // Now GDPR Privacy Policy modal should be available
    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await page.waitForTimeout(500);
    await expect(page.locator('.modal.is-active')).toBeVisible();
    await expect(page.locator('#privacy_gdpr_content')).toBeVisible();
  });

  test('switching from Simple to No Tracking changes modal at step 8', async ({ page }) => {
    await navigateToStep2(page);
    await fillStep2(page);
    await page.locator('input[type="radio"][value="1"]').check();
    await clickNext(page);
    await expectStep(page, 3);
    await clickNext(page);
    await expectStep(page, 4);
    await clickNext(page);
    await expectStep(page, 5);
    await clickNext(page);
    await expectStep(page, 6);
    await page.locator('#devName').fill('John Doe');
    await clickNext(page);
    await expectStep(page, 7);
    await clickNext(page);
    await expectStep(page, 8);

    // Go back to step 2
    for (let i = 0; i < 6; i++) {
      await page.locator('a.card-footer-item').filter({ hasText: 'Previous' }).click();
      await page.waitForTimeout(100);
    }
    await expectStep(page, 2);

    // Switch to No Tracking
    await page.locator('input[type="radio"][value="2"]').check();
    await clickNext(page);
    await expectStep(page, 3);
    await clickNext(page);
    await expectStep(page, 4);
    await clickNext(page);
    await expectStep(page, 5);
    await clickNext(page);
    await expectStep(page, 6);
    await clickNext(page);
    await expectStep(page, 7);
    await clickNext(page);
    await expectStep(page, 8);

    // No Tracking Privacy Policy modal should be available
    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await page.waitForTimeout(500);
    await expect(page.locator('.modal.is-active')).toBeVisible();
    await expect(page.locator('#privacy_notrack_content')).toBeVisible();
  });
});

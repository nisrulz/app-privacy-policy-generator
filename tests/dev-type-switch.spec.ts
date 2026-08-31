import { test, expect } from '@playwright/test';
import { navigateToStep2, fillStep2, clickNext, expectStep, navigateToStep7 } from './helpers';

test.describe('Company/Individual switch mid-step 6', () => {
  test('switching from Company to Individual shows dev name and hides company name', async ({ page }) => {
    await navigateToStep2(page);
    await fillStep2(page);
    await clickNext(page);
    await expectStep(page, 3);
    await clickNext(page);
    await expectStep(page, 4);
    await clickNext(page);
    await expectStep(page, 5);
    await clickNext(page);
    await expectStep(page, 6);

    await page.locator('input[type="radio"][value="Company"]').check();
    await expect(page.locator('#companyName')).toBeVisible();
    await expect(page.locator('#devName')).not.toBeVisible();

    await page.locator('input[type="radio"][value="Individual"]').check();
    await expect(page.locator('#devName')).toBeVisible();
    await expect(page.locator('#companyName')).not.toBeVisible();
  });

  test('switching from Individual to Company shows company name and hides dev name', async ({ page }) => {
    await navigateToStep2(page);
    await fillStep2(page);
    await clickNext(page);
    await expectStep(page, 3);
    await clickNext(page);
    await expectStep(page, 4);
    await clickNext(page);
    await expectStep(page, 5);
    await clickNext(page);
    await expectStep(page, 6);

    await page.locator('input[type="radio"][value="Individual"]').check();
    await expect(page.locator('#devName')).toBeVisible();
    await expect(page.locator('#companyName')).not.toBeVisible();

    await page.locator('input[type="radio"][value="Company"]').check();
    await expect(page.locator('#companyName')).toBeVisible();
    await expect(page.locator('#devName')).not.toBeVisible();
  });
});

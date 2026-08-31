import { test, expect, type Page } from '@playwright/test';
import { gotoApp, clickStart, clickNext, expectStep, navigateToStep2 } from './helpers';

const disabledNext = (page: Page) =>
  page.locator('a.card-footer-item').filter({ hasText: 'Next' });

test.describe('canAdvance blocking behavior', () => {
  test('clicking Next on step 2 with empty fields does not advance', async ({ page }) => {
    await navigateToStep2(page);
    await disabledNext(page).click({ force: true });
    await expectStep(page, 2);
  });

  test('clicking Next on step 2 with only appName filled does not advance', async ({ page }) => {
    await navigateToStep2(page);
    await page.locator('#appName').fill('Test App');
    await disabledNext(page).click({ force: true });
    await expectStep(page, 2);
  });

  test('clicking Next on step 2 with only contact filled does not advance', async ({ page }) => {
    await navigateToStep2(page);
    await page.locator('#appContact').fill('test@example.com');
    await disabledNext(page).click({ force: true });
    await expectStep(page, 2);
  });

  test('clicking Next on step 6 with Individual and empty devName does not advance', async ({ page }) => {
    await navigateToStep2(page);
    await page.locator('#appName').fill('Test App');
    await page.locator('#appContact').fill('test@example.com');
    await clickNext(page);
    await expectStep(page, 3);
    await clickNext(page);
    await expectStep(page, 4);
    await clickNext(page);
    await expectStep(page, 5);
    await clickNext(page);
    await expectStep(page, 6);
    await page.locator('input[type="radio"][value="Individual"]').check();
    await disabledNext(page).click({ force: true });
    await expectStep(page, 6);
  });

  test('clicking Next on step 6 with Company and empty companyName does not advance', async ({ page }) => {
    await navigateToStep2(page);
    await page.locator('#appName').fill('Test App');
    await page.locator('#appContact').fill('test@example.com');
    await clickNext(page);
    await expectStep(page, 3);
    await clickNext(page);
    await expectStep(page, 4);
    await clickNext(page);
    await expectStep(page, 5);
    await clickNext(page);
    await expectStep(page, 6);
    await page.locator('input[type="radio"][value="Company"]').check();
    await disabledNext(page).click({ force: true });
    await expectStep(page, 6);
  });

  test('Next button has is-disabled class when canAdvance is false', async ({ page }) => {
    await navigateToStep2(page);
    await expect(disabledNext(page)).toHaveClass(/is-disabled/);
  });

  test('Next button loses is-disabled class when fields are filled', async ({ page }) => {
    await navigateToStep2(page);
    await page.locator('#appName').fill('Test App');
    await page.locator('#appContact').fill('test@example.com');
    await expect(disabledNext(page)).not.toHaveClass(/is-disabled/);
  });
});

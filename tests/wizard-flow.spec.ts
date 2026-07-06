import { test, expect } from '@playwright/test';
import { gotoApp, clickStart, clickNext, clickPrevious, expectStep, fillStep2 } from './helpers';

test.describe('Wizard Flow', () => {
  test.beforeEach(async ({ page }) => {
    await gotoApp(page);
    await expectStep(page, 1);
  });

  test('navigates through all steps with Simple policy', async ({ page }) => {
    await clickStart(page);
    await expectStep(page, 2);

    await fillStep2(page);
    await clickNext(page);
    await expectStep(page, 3);

    await clickNext(page);
    await expectStep(page, 4);

    await clickNext(page);
    await expectStep(page, 5);

    await clickNext(page);
    await expectStep(page, 6);

    await page.locator('#devName').fill('Test Developer');
    await clickNext(page);
    await expectStep(page, 7);

    await clickNext(page);
    await expectStep(page, 8);
  });

  test('navigates through all steps with GDPR policy', async ({ page }) => {
    await clickStart(page);
    await expectStep(page, 2);

    await fillStep2(page);
    await page.locator('input[type="radio"][value="3"]').check();
    await clickNext(page);
    await expectStep(page, 3);

    await page.locator('#businessAddress').fill('123 Test Street, Berlin, Germany');
    await clickNext(page);
    await expectStep(page, 4);

    await clickNext(page);
    await expectStep(page, 5);

    await clickNext(page);
    await expectStep(page, 6);

    await page.locator('#devName').fill('Test Developer');
    await clickNext(page);
    await expectStep(page, 7);

    await clickNext(page);
    await expectStep(page, 8);
  });

  test('can navigate backwards between steps', async ({ page }) => {
    await clickStart(page);
    await expectStep(page, 2);

    await clickPrevious(page);
    await expectStep(page, 1);

    await clickStart(page);
    await expectStep(page, 2);
  });

  test('step 2 blocks advancement without required fields', async ({ page }) => {
    await clickStart(page);

    const nextBtn = page.locator('a.card-footer-item').filter({ hasText: 'Next' });
    await expect(nextBtn).toHaveClass(/is-disabled/);

    await page.locator('#appName').fill('My App');
    await expect(nextBtn).toHaveClass(/is-disabled/);

    await page.locator('#appContact').fill('test@example.com');
    await expect(nextBtn).not.toHaveClass(/is-disabled/);
  });

  test('step 6 blocks advancement without dev name', async ({ page }) => {
    await clickStart(page);
    await fillStep2(page);
    await clickNext(page);
    await expectStep(page, 3);

    await clickNext(page);
    await expectStep(page, 4);

    await clickNext(page);
    await expectStep(page, 5);

    await clickNext(page);
    await expectStep(page, 6);

    const nextBtn = page.locator('a.card-footer-item').filter({ hasText: 'Next' });
    await expect(nextBtn).toHaveClass(/is-disabled/);

    await page.locator('#devName').fill('Test Dev');
    await expect(nextBtn).not.toHaveClass(/is-disabled/);
  });
});

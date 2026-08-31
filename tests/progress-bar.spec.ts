import { test, expect } from '@playwright/test';
import { gotoApp, clickStart, clickNext, expectStep, fillStep2, navigateToStep6 } from './helpers';

test.describe('Progress Bar', () => {
  test('step 2 shows progress bar', async ({ page }) => {
    await gotoApp(page);
    await clickStart(page);
    await expectStep(page, 2);

    const progress = page.locator('#step-2 progress.progress');
    await expect(progress).toBeVisible();
  });

  test('progress bar value increases as wizard advances', async ({ page }) => {
    await gotoApp(page);
    await clickStart(page);
    await expectStep(page, 2);

    const progress2 = page.locator('#step-2 progress.progress');
    const value2 = Number(await progress2.getAttribute('value'));

    await fillStep2(page);
    await clickNext(page);
    await expectStep(page, 3);

    const progress3 = page.locator('#step-3 progress.progress');
    const value3 = Number(await progress3.getAttribute('value'));

    expect(value3).toBeGreaterThan(value2);
  });

  test('progress bar reaches 100 on step 8', async ({ page }) => {
    await gotoApp(page);
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
    await page.locator('#devName').fill('Test');
    await clickNext(page);
    await expectStep(page, 7);
    await clickNext(page);
    await expectStep(page, 8);

    const progress = page.locator('#step-8 progress.progress');
    await expect(progress).toBeVisible();
    expect(Number(await progress.getAttribute('value'))).toBe(100);
  });

  test('progress bar uses primary class on step 8', async ({ page }) => {
    await gotoApp(page);
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
    await page.locator('#devName').fill('Test');
    await clickNext(page);
    await expectStep(page, 7);
    await clickNext(page);
    await expectStep(page, 8);

    const progress = page.locator('#step-8 progress.progress');
    await expect(progress).toHaveClass(/is-primary/);
  });
});

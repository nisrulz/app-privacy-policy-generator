import { test, expect } from '@playwright/test';
import { gotoApp, clickStart, clickNext, expectStep, fillStep2, navigateToStep2 } from './helpers';

test.describe('Effective date in output', () => {
  test('generated policy includes the effective date from step 3', async ({ page }) => {
    await navigateToStep2(page);
    await fillStep2(page);
    await clickNext(page);
    await expectStep(page, 3);
    await page.locator('#effectiveFromDate').fill('2025-06-15');
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

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await page.waitForTimeout(500);
    await expect(page.locator('.modal.is-active')).toBeVisible();
    const content = await page.locator('#privacy_simple_content').innerText();
    expect(content).toContain('2025-06-15');
  });

  test('generated T&C includes the effective date from step 3', async ({ page }) => {
    await navigateToStep2(page);
    await fillStep2(page);
    await clickNext(page);
    await expectStep(page, 3);
    await page.locator('#effectiveFromDate').fill('2025-01-01');
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

    await page.locator('a.card-footer-item').filter({ hasText: 'Terms' }).click();
    await page.waitForTimeout(500);
    await expect(page.locator('.modal.is-active')).toBeVisible();
    const content = await page.locator('#tandc_content').innerText();
    expect(content).toContain('2025-01-01');
  });
});

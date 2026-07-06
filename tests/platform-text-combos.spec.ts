import { test, expect } from '@playwright/test';
import { gotoApp, clickStart, clickNext, expectStep, fillStep2, navigateToStep5 } from './helpers';

test.describe('Platform text combinations', () => {
  test('web-only shows web browsers text in policy', async ({ page }) => {
    await navigateToStep5(page);
    await page.locator('label[for="platform-android"]').click();
    await page.locator('label[for="platform-web"]').click();
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
    expect(content).toContain('web browsers');
  });

  test('multiple platforms joined with commas and conjunction', async ({ page }) => {
    await navigateToStep5(page);
    await page.locator('label[for="platform-android"]').click();
    await page.locator('label[for="platform-ios"]').click();
    await page.locator('label[for="platform-web"]').click();
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
    expect(content).toContain('mobile devices');
    expect(content).toContain('web browsers');
  });
});

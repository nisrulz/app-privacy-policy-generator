import { test, expect } from '@playwright/test';
import { completeWizard, gotoApp, clickStart, clickNext, expectStep, navigateToStep2 } from './helpers';

test.describe('Special characters in form inputs', () => {
  test('policy output handles quotes and ampersands in app name', async ({ page }) => {
    await navigateToStep2(page);
    await page.locator('#appName').fill("Tom's \"App\" & More");
    await page.locator('#appContact').fill('test@example.com');
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

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await page.waitForTimeout(500);
    await expect(page.locator('.modal.is-active')).toBeVisible();
    const content = await page.locator('#privacy_simple_content').innerText();
    expect(content).toContain("Tom's \"App\" & More");
  });

  test('policy output handles quotes and ampersand in developer name', async ({ page }) => {
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
    await page.locator('#devName').fill('Dev "Script" & Co');
    await clickNext(page);
    await expectStep(page, 7);
    await clickNext(page);
    await expectStep(page, 8);

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await page.waitForTimeout(500);
    await expect(page.locator('.modal.is-active')).toBeVisible();
    const content = await page.locator('#privacy_simple_content').innerText();
    expect(content).toContain('Dev "Script" & Co');
  });

  test('policy output handles emoji in app name', async ({ page }) => {
    await navigateToStep2(page);
    await page.locator('#appName').fill('My App 🚀');
    await page.locator('#appContact').fill('test@example.com');
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

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await page.waitForTimeout(500);
    await expect(page.locator('.modal.is-active')).toBeVisible();
    const content = await page.locator('#privacy_simple_content').innerText();
    expect(content).toContain('My App 🚀');
  });

  test('policy output handles long strings', async ({ page }) => {
    const longName = 'A'.repeat(200);
    await navigateToStep2(page);
    await page.locator('#appName').fill(longName);
    await page.locator('#appContact').fill('test@example.com');
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

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await page.waitForTimeout(500);
    await expect(page.locator('.modal.is-active')).toBeVisible();
    const content = await page.locator('#privacy_simple_content').innerText();
    expect(content).toContain(longName);
  });
});

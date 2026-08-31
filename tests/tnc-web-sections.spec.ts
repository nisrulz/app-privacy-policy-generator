import { test, expect } from '@playwright/test';
import { navigateToStep2, fillStep2, clickNext, expectStep, gotoApp } from './helpers';

test.describe('T&C web-specific sections', () => {
  test('T&C shows web data and charges sections for web apps', async ({ page }) => {
    await gotoApp(page);
    await page.locator('button.start-btn').click();
    await page.waitForTimeout(200);

    await page.locator('#appName').fill('Web App');
    await page.locator('#appContact').fill('test@example.com');
    await clickNext(page);
    await expectStep(page, 3);
    await clickNext(page);
    await expectStep(page, 4);
    await clickNext(page);
    await expectStep(page, 5);

    await page.locator('label[for="platform-android"]').click();
    await page.locator('label[for="platform-web"]').click();

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
    expect(content).toContain('internet service provider');
  });

  test('T&C does not show mobile data sections for web-only apps', async ({ page }) => {
    await gotoApp(page);
    await page.locator('button.start-btn').click();
    await page.waitForTimeout(200);

    await page.locator('#appName').fill('Web App');
    await page.locator('#appContact').fill('test@example.com');
    await clickNext(page);
    await expectStep(page, 3);
    await clickNext(page);
    await expectStep(page, 4);
    await clickNext(page);
    await expectStep(page, 5);

    await page.locator('label[for="platform-android"]').click();
    await page.locator('label[for="platform-web"]').click();

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
    expect(content).not.toContain('mobile network');
  });
});

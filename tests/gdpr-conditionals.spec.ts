import { test, expect } from '@playwright/test';
import { gotoApp, clickStart, clickNext, expectStep, fillStep2, navigateToStep2 } from './helpers';

test.describe('GDPR policy conditional sections', () => {
  test('GDPR policy shows location tracking details when enabled', async ({ page }) => {
    await gotoApp(page);
    await clickStart(page);
    await expectStep(page, 2);
    await fillStep2(page);
    await page.locator('input[type="radio"][value="3"]').check();
    await clickNext(page);
    await expectStep(page, 3);
    await page.locator('#businessAddress').fill('123 EU Street');
    await clickNext(page);
    await expectStep(page, 4);
    // Enable location tracking
    await page.locator('label[for="locationcheckbox"]').click();
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
    const content = await page.locator('#privacy_gdpr_content').innerText();
    expect(content).toContain('location');
  });

  test('GDPR policy shows no-tracking text when location tracking disabled', async ({ page }) => {
    await gotoApp(page);
    await clickStart(page);
    await expectStep(page, 2);
    await fillStep2(page);
    await page.locator('input[type="radio"][value="3"]').check();
    await clickNext(page);
    await expectStep(page, 3);
    await page.locator('#businessAddress').fill('123 EU Street');
    await clickNext(page);
    await expectStep(page, 4);
    // Leave location tracking unchecked
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
    const content = await page.locator('#privacy_gdpr_content').innerText();
    expect(content).toContain('location');
  });

  test('GDPR policy shows AI usage details when enabled', async ({ page }) => {
    await gotoApp(page);
    await clickStart(page);
    await expectStep(page, 2);
    await fillStep2(page);
    await page.locator('input[type="radio"][value="3"]').check();
    await clickNext(page);
    await expectStep(page, 3);
    await page.locator('#businessAddress').fill('123 EU Street');
    await clickNext(page);
    await expectStep(page, 4);
    // Enable AI
    await page.locator('label[for="aicheckbox"]').click();
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
    const content = await page.locator('#privacy_gdpr_content').innerText();
    expect(content).toContain('AI');
  });

  test('GDPR policy shows data deletion section when flag is enabled', async ({ page }) => {
    await gotoApp(page);
    await clickStart(page);
    await expectStep(page, 2);
    await fillStep2(page);
    await page.locator('input[type="radio"][value="3"]').check();
    await clickNext(page);
    await expectStep(page, 3);
    await page.locator('#businessAddress').fill('123 EU Street');
    await clickNext(page);
    await expectStep(page, 4);
    // Enable data deletion
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
    const content = await page.locator('#privacy_gdpr_content').innerText();
    expect(content).toContain('deletion');
  });

  test('GDPR policy shows different children text based on third-party service selection', async ({ page }) => {
    await gotoApp(page);
    await clickStart(page);
    await expectStep(page, 2);
    await fillStep2(page);
    await page.locator('input[type="radio"][value="3"]').check();
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
    // Don't enable any services
    await clickNext(page);
    await expectStep(page, 8);

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await page.waitForTimeout(500);
    const content = await page.locator('#privacy_gdpr_content').innerText();
    // Children section should still be present
    expect(content).toContain('children');
  });

  test('GDPR policy shows different children text when services enabled', async ({ page }) => {
    await gotoApp(page);
    await clickStart(page);
    await expectStep(page, 2);
    await fillStep2(page);
    await page.locator('input[type="radio"][value="3"]').check();
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
    // Enable a service
    const firstLabel = page.locator('.scrollable-thirdparty.content label').first();
    await firstLabel.click();
    await clickNext(page);
    await expectStep(page, 8);

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await page.waitForTimeout(500);
    const content = await page.locator('#privacy_gdpr_content').innerText();
    expect(content).toContain('children');
  });

  test('GDPR policy includes business address when provided', async ({ page }) => {
    await gotoApp(page);
    await clickStart(page);
    await expectStep(page, 2);
    await fillStep2(page);
    await page.locator('input[type="radio"][value="3"]').check();
    await clickNext(page);
    await expectStep(page, 3);
    await page.locator('#businessAddress').fill('456 Berlin Ave');
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
    const content = await page.locator('#privacy_gdpr_content').innerText();
    expect(content).toContain('456 Berlin Ave');
  });

  test('GDPR policy hides business address line when address is empty', async ({ page }) => {
    await gotoApp(page);
    await clickStart(page);
    await expectStep(page, 2);
    await fillStep2(page);
    await page.locator('input[type="radio"][value="3"]').check();
    await clickNext(page);
    await expectStep(page, 3);
    // Leave business address empty
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
    const content = await page.locator('#privacy_gdpr_content').innerText();
    // Address should not appear in output
    expect(content).not.toContain('Address');
  });
});

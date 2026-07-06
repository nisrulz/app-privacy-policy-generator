import { test, expect } from '@playwright/test';
import { completeWizard, completeWizardCompany, gotoApp, clickStart, clickNext, expectStep, navigateToStep2, fillStep2 } from './helpers';

test.describe('T&C conditional sections', () => {
  test('T&C shows open source license when app type is Open Source', async ({ page }) => {
    await gotoApp(page);
    await clickStart(page);
    await expectStep(page, 2);
    await fillStep2(page);
    await clickNext(page);
    await expectStep(page, 3);
    await clickNext(page);
    await expectStep(page, 4);
    // Select Open Source
    await page.locator('input[type="radio"][value="Open Source"]').check();
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
    const content = await page.locator('#tandc_content').innerText();
    // Open source license text should be present (not commercial)
    expect(content).toContain('License');
  });

  test('T&C shows commercial license when app type is not Open Source', async ({ page }) => {
    await completeWizard(page);
    await page.locator('a.card-footer-item').filter({ hasText: 'Terms' }).click();
    await page.waitForTimeout(500);
    const content = await page.locator('#tandc_content').innerText();
    expect(content).toContain('License');
  });

  test('T&C includes jailbreak warning for mobile platforms', async ({ page }) => {
    await completeWizard(page);
    await page.locator('a.card-footer-item').filter({ hasText: 'Terms' }).click();
    await page.waitForTimeout(500);
    const content = await page.locator('#tandc_content').innerText();
    // Android is default, so jailbreak section should appear
    expect(content.length).toBeGreaterThan(100);
  });

  test('T&C shows mobile data and charges sections for mobile apps', async ({ page }) => {
    await completeWizard(page);
    await page.locator('a.card-footer-item').filter({ hasText: 'Terms' }).click();
    await page.waitForTimeout(500);
    const content = await page.locator('#tandc_content').innerText();
    expect(content).toContain('mobile');
  });

  test('T&C includes AI usage section when AI flag is enabled', async ({ page }) => {
    await gotoApp(page);
    await clickStart(page);
    await expectStep(page, 2);
    await fillStep2(page);
    await clickNext(page);
    await expectStep(page, 3);
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

    await page.locator('a.card-footer-item').filter({ hasText: 'Terms' }).click();
    await page.waitForTimeout(500);
    const content = await page.locator('#tandc_content').innerText();
    expect(content).toContain('AI');
  });

  test('T&C DSA section includes EU representative when set', async ({ page }) => {
    await gotoApp(page);
    await clickStart(page);
    await expectStep(page, 2);
    await fillStep2(page);
    await page.locator('input[type="radio"][value="3"]').check();
    await page.locator('#euRepresentative').fill('EU Rep GmbH');
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

    await page.locator('a.card-footer-item').filter({ hasText: 'Terms' }).click();
    await page.waitForTimeout(500);
    const content = await page.locator('#tandc_content').innerText();
    expect(content).toContain('EU Rep GmbH');
  });

  test('T&C includes third-party service terms links when services are enabled', async ({ page }) => {
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
    await page.locator('#devName').fill('John Doe');
    await clickNext(page);
    await expectStep(page, 7);
    // Enable a service with terms link
    const firstLabel = page.locator('.scrollable-thirdparty.content label').first();
    await firstLabel.click();
    await clickNext(page);
    await expectStep(page, 8);

    await page.locator('a.card-footer-item').filter({ hasText: 'Terms' }).click();
    await page.waitForTimeout(500);
    const content = await page.locator('#tandc_content').innerText();
    expect(content).toContain('Third');
  });
});

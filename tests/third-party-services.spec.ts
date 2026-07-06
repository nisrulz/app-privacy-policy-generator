import { test, expect } from '@playwright/test';
import { gotoApp, clickStart, clickNext, expectStep, navigateToStep7 } from './helpers';

test.describe('Third-Party Services', () => {
  test('displays list of third-party services', async ({ page }) => {
    await navigateToStep7(page);

    const firstLabel = page.locator('.scrollable-thirdparty label').first();
    await expect(firstLabel).toBeVisible({ timeout: 10000 });

    const count = await page.locator('.scrollable-thirdparty label').count();
    expect(count).toBeGreaterThan(0);
  });

  test('can toggle a service on and off', async ({ page }) => {
    await navigateToStep7(page);

    const firstLabel = page.locator('.scrollable-thirdparty label').first();
    await expect(firstLabel).toBeVisible({ timeout: 10000 });
    const firstCheckbox = page.locator('.scrollable-thirdparty input[type="checkbox"]').first();

    await firstLabel.click();
    await expect(firstCheckbox).toBeChecked();

    await firstLabel.click();
    await expect(firstCheckbox).not.toBeChecked();
  });

  test('toggled services appear in generated output', async ({ page }) => {
    await navigateToStep7(page);

    const firstServiceLabel = page.locator('.scrollable-thirdparty label').first();
    await expect(firstServiceLabel).toBeVisible({ timeout: 10000 });
    const serviceName = await firstServiceLabel.textContent();

    await firstServiceLabel.click();

    await clickNext(page);
    await expectStep(page, 8);

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await expect(page.locator('.modal.is-active')).toBeVisible();

    const content = await page.locator('#privacy_simple_content').textContent();
    expect(content).toContain(serviceName?.trim() || '');
  });

  test('No Tracking policy shows no third-party services message', async ({ page }) => {
    await gotoApp(page);
    await clickStart(page);
    await expectStep(page, 2);

    await page.locator('#appName').fill('Test App');
    await page.locator('#appContact').fill('test@example.com');
    await page.locator('input[type="radio"][value="2"]').check();
    await clickNext(page);
    await expectStep(page, 3);

    await clickNext(page);
    await expectStep(page, 4);

    await clickNext(page);
    await expectStep(page, 5);

    await clickNext(page);
    await expectStep(page, 6);

    await page.locator('#devName').fill('Test Dev');
    await clickNext(page);
    await expectStep(page, 7);

    const noServicesSection = page.locator('.thirdparty-section').last();
    await expect(noServicesSection).toBeVisible();
  });
});

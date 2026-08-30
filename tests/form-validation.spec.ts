import { test, expect } from '@playwright/test';
import { gotoApp, clickStart, clickNext, expectStep, navigateToStep6 } from './helpers';

test.describe('Form Validation', () => {
  test('shows error when app name is empty', async ({ page }) => {
    await gotoApp(page);
    await clickStart(page);
    await expectStep(page, 2);

    const appNameInput = page.locator('#appName');
    await appNameInput.fill('');
    await appNameInput.blur();

    await expect(page.locator('#step-2 .help.is-danger').first()).toBeVisible();
  });

  test('shows error when contact info is empty', async ({ page }) => {
    await gotoApp(page);
    await clickStart(page);
    await expectStep(page, 2);

    const contactInput = page.locator('#appContact');
    await contactInput.fill('');
    await contactInput.blur();

    await expect(page.locator('#step-2 .help.is-danger').first()).toBeVisible();
  });

  test('clears error when field is filled', async ({ page }) => {
    await gotoApp(page);
    await clickStart(page);
    await expectStep(page, 2);

    const appNameInput = page.locator('#appName');
    await appNameInput.fill('');
    await appNameInput.blur();
    await expect(page.locator('#step-2 .help.is-danger').first()).toBeVisible();

    await appNameInput.fill('Test App');
    await page.locator('#appContact').fill('test@example.com');
    await expect(page.locator('#step-2 .help.is-danger')).toHaveCount(0);
  });

  test('step 6 shows error for empty dev name (Individual)', async ({ page }) => {
    await navigateToStep6(page);
    await expectStep(page, 6);

    const devNameInput = page.locator('#devName');
    await devNameInput.fill('');
    await devNameInput.blur();

    await expect(page.locator('#step-6 .help.is-danger').first()).toBeVisible();
  });

  test('step 6 shows error for empty company name (Company)', async ({ page }) => {
    await navigateToStep6(page);
    await expectStep(page, 6);

    await page.locator('input[type="radio"][value="Company"]').check();

    const companyNameInput = page.locator('#companyName');
    await companyNameInput.fill('');
    await companyNameInput.blur();

    await expect(page.locator('#step-6 .help.is-danger').first()).toBeVisible();
  });

  test('GDPR policy shows business address field as required', async ({ page }) => {
    await gotoApp(page);
    await clickStart(page);
    await expectStep(page, 2);

    await page.locator('#appName').fill('Test App');
    await page.locator('#appContact').fill('test@example.com');
    await page.locator('input[type="radio"][value="3"]').check();
    await clickNext(page);
    await expectStep(page, 3);

    const businessAddress = page.locator('#businessAddress');
    await expect(businessAddress).toBeVisible();
  });
});

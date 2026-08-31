import { test, expect } from '@playwright/test';
import { gotoApp, clickStart, clickNext, expectStep, fillStep2, completeWizardGDPR } from './helpers';

test.describe('GDPR Fields', () => {
  test.describe('EU Representative', () => {
    test('EU representative field appears when GDPR is selected', async ({ page }) => {
      await gotoApp(page);
      await clickStart(page);
      await expectStep(page, 2);

      await page.locator('#appName').fill('Test App');
      await page.locator('#appContact').fill('test@example.com');
      await page.locator('input[type="radio"][value="3"]').check();

      await expect(page.locator('#euRepresentative')).toBeVisible();
    });

    test('EU representative field is hidden for non-GDPR policies', async ({ page }) => {
      await gotoApp(page);
      await clickStart(page);
      await expectStep(page, 2);

      await page.locator('#appName').fill('Test App');
      await page.locator('#appContact').fill('test@example.com');
      await page.locator('input[type="radio"][value="1"]').check();

      await expect(page.locator('#euRepresentative')).not.toBeVisible();
    });

    test('EU representative appears in GDPR policy output', async ({ page }) => {
      await completeWizardGDPR(page, { euRepresentative: 'EU Rep GmbH' });

      await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
      await expect(page.locator('.modal.is-active')).toBeVisible();

      const content = await page.locator('#privacy_gdpr_content').textContent();
      expect(content).toContain('EU Rep GmbH');
    });
  });

  test.describe('Business Address', () => {
    test('GDPR policy requires business address on step 3', async ({ page }) => {
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

      await expect(businessAddress).toHaveClass(/is-danger/);
    });

    test('business address not required for Simple policy', async ({ page }) => {
      await gotoApp(page);
      await clickStart(page);
      await expectStep(page, 2);

      await page.locator('#appName').fill('Test App');
      await page.locator('#appContact').fill('test@example.com');
      await page.locator('input[type="radio"][value="1"]').check();
      await clickNext(page);
      await expectStep(page, 3);

      await expect(page.locator('#businessAddress')).not.toHaveClass(/is-danger/);
    });

    test('business address appears in GDPR policy output', async ({ page }) => {
      await completeWizardGDPR(page, { businessAddress: '456 GDPR Lane, Munich' });

      await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
      await expect(page.locator('.modal.is-active')).toBeVisible();

      const content = await page.locator('#privacy_gdpr_content').textContent();
      expect(content).toContain('456 GDPR Lane, Munich');
    });
  });

  test.describe('Step 3 Fields', () => {
    test('effective date field has today as default', async ({ page }) => {
      await gotoApp(page);
      await clickStart(page);
      await expectStep(page, 2);
      await fillStep2(page);
      await clickNext(page);
      await expectStep(page, 3);

      const dateInput = page.locator('#effectiveFromDate');
      await expect(dateInput).toBeVisible();

      const today = new Date().toISOString().slice(0, 10);
      const value = await dateInput.inputValue();
      expect(value).toBe(today);
    });

    test('can change effective date', async ({ page }) => {
      await gotoApp(page);
      await clickStart(page);
      await expectStep(page, 2);
      await fillStep2(page);
      await clickNext(page);
      await expectStep(page, 3);

      const dateInput = page.locator('#effectiveFromDate');
      await dateInput.fill('2025-01-15');
      expect(await dateInput.inputValue()).toBe('2025-01-15');
    });

    test('PID info field is visible for Simple policy', async ({ page }) => {
      await gotoApp(page);
      await clickStart(page);
      await expectStep(page, 2);
      await fillStep2(page);
      await page.locator('input[type="radio"][value="1"]').check();
      await clickNext(page);
      await expectStep(page, 3);

      await expect(page.locator('#pidInfoIn')).toBeVisible();
    });

    test('age of digital consent field has default value 16', async ({ page }) => {
      await gotoApp(page);
      await clickStart(page);
      await expectStep(page, 2);
      await fillStep2(page);
      await clickNext(page);
      await expectStep(page, 3);

      const ageInput = page.locator('#ageOfDigitalConsent');
      await expect(ageInput).toBeVisible();
      expect(await ageInput.inputValue()).toBe('16');
    });

    test('can change age of digital consent', async ({ page }) => {
      await gotoApp(page);
      await clickStart(page);
      await expectStep(page, 2);
      await fillStep2(page);
      await clickNext(page);
      await expectStep(page, 3);

      const ageInput = page.locator('#ageOfDigitalConsent');
      await ageInput.fill('13');
      expect(await ageInput.inputValue()).toBe('13');
    });
  });
});

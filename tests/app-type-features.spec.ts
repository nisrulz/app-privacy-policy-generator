import { test, expect } from '@playwright/test';
import { gotoApp, clickStart, clickNext, expectStep, fillStep2, navigateToStep4, completeWizard } from './helpers';

test.describe('App Type and Feature Flags', () => {
  test.describe('App Type', () => {
    test('step 4 shows all app type radio options', async ({ page }) => {
      await navigateToStep4(page);

      await expect(page.locator('input[type="radio"][value="Free"]')).toBeVisible();
      await expect(page.locator('input[type="radio"][value="Open Source"]')).toBeVisible();
      await expect(page.locator('input[type="radio"][value="Freemium"]')).toBeVisible();
      await expect(page.locator('input[type="radio"][value="Ad Supported"]')).toBeVisible();
      await expect(page.locator('input[type="radio"][value="Commercial"]')).toBeVisible();
    });

    test('Free is selected by default', async ({ page }) => {
      await navigateToStep4(page);

      await expect(page.locator('input[type="radio"][value="Free"]')).toBeChecked();
    });

    test('can select different app types and tag updates', async ({ page }) => {
      await navigateToStep4(page);

      await page.locator('input[type="radio"][value="Commercial"]').check();
      await expect(page.locator('input[type="radio"][value="Commercial"]')).toBeChecked();
      await expect(page.locator('input[type="radio"][value="Free"]')).not.toBeChecked();

      const tag = page.locator('#step-4 .tag.is-primary.is-light');
      await expect(tag).toContainText('Commercial');
    });

    test('app type affects platform description in policy', async ({ page }) => {
      await gotoApp(page);
      await clickStart(page);
      await expectStep(page, 2);
      await fillStep2(page);
      await clickNext(page);
      await expectStep(page, 3);
      await clickNext(page);
      await expectStep(page, 4);

      await page.locator('input[type="radio"][value="Commercial"]').check();
      await clickNext(page);
      await expectStep(page, 5);
      await clickNext(page);
      await expectStep(page, 6);
      await page.locator('#devName').fill('Type Test');
      await clickNext(page);
      await expectStep(page, 7);
      await clickNext(page);
      await expectStep(page, 8);

      await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
      await expect(page.locator('.modal.is-active')).toBeVisible();

      const content = await page.locator('#privacy_simple_content').textContent();
      expect(content).toContain('Type Test');
    });
  });

  test.describe('Feature Flags', () => {
    test('feature checkboxes labels are visible on step 4', async ({ page }) => {
      await navigateToStep4(page);

      await expect(page.locator('label[for="locationcheckbox"]')).toBeVisible();
      await expect(page.locator('label[for="aicheckbox"]')).toBeVisible();
      await expect(page.locator('label[for="datadeletioncheckbox"]')).toBeVisible();
    });

    test('feature checkboxes are unchecked by default', async ({ page }) => {
      await navigateToStep4(page);

      await expect(page.locator('#locationcheckbox')).not.toBeChecked();
      await expect(page.locator('#aicheckbox')).not.toBeChecked();
      await expect(page.locator('#datadeletioncheckbox')).not.toBeChecked();
    });

    test('can toggle feature checkboxes via labels', async ({ page }) => {
      await navigateToStep4(page);

      await page.locator('label[for="locationcheckbox"]').click();
      await page.locator('label[for="aicheckbox"]').click();
      await page.locator('label[for="datadeletioncheckbox"]').click();

      await expect(page.locator('#locationcheckbox')).toBeChecked();
      await expect(page.locator('#aicheckbox')).toBeChecked();
      await expect(page.locator('#datadeletioncheckbox')).toBeChecked();
    });

    test('location tracking adds location section in policy', async ({ page }) => {
      await gotoApp(page);
      await clickStart(page);
      await expectStep(page, 2);
      await fillStep2(page);
      await clickNext(page);
      await expectStep(page, 3);
      await clickNext(page);
      await expectStep(page, 4);

      await page.locator('label[for="locationcheckbox"]').click();
      await clickNext(page);
      await expectStep(page, 5);
      await clickNext(page);
      await expectStep(page, 6);
      await page.locator('#devName').fill('Feature Test');
      await clickNext(page);
      await expectStep(page, 7);
      await clickNext(page);
      await expectStep(page, 8);

      await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
      await expect(page.locator('.modal.is-active')).toBeVisible();

      const content = await page.locator('#privacy_simple_content').textContent();
      expect(content).toContain('Location Information');
    });

    test('AI usage adds AI section in policy', async ({ page }) => {
      await gotoApp(page);
      await clickStart(page);
      await expectStep(page, 2);
      await fillStep2(page);
      await clickNext(page);
      await expectStep(page, 3);
      await clickNext(page);
      await expectStep(page, 4);

      await page.locator('label[for="aicheckbox"]').click();
      await clickNext(page);
      await expectStep(page, 5);
      await clickNext(page);
      await expectStep(page, 6);
      await page.locator('#devName').fill('Feature Test');
      await clickNext(page);
      await expectStep(page, 7);
      await clickNext(page);
      await expectStep(page, 8);

      await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
      await expect(page.locator('.modal.is-active')).toBeVisible();

      const content = await page.locator('#privacy_simple_content').textContent();
      expect(content).toContain('Artificial Intelligence');
    });

    test('data deletion adds deletion section in policy', async ({ page }) => {
      await gotoApp(page);
      await clickStart(page);
      await expectStep(page, 2);
      await fillStep2(page);
      await clickNext(page);
      await expectStep(page, 3);
      await clickNext(page);
      await expectStep(page, 4);

      await page.locator('label[for="datadeletioncheckbox"]').click();
      await clickNext(page);
      await expectStep(page, 5);
      await clickNext(page);
      await expectStep(page, 6);
      await page.locator('#devName').fill('Feature Test');
      await clickNext(page);
      await expectStep(page, 7);
      await clickNext(page);
      await expectStep(page, 8);

      await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
      await expect(page.locator('.modal.is-active')).toBeVisible();

      const content = await page.locator('#privacy_simple_content').textContent();
      expect(content).toContain('Data Deletion');
    });

    test('feature checkboxes hidden for No Tracking policy', async ({ page }) => {
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

      await expect(page.locator('label[for="locationcheckbox"]')).not.toBeVisible();
      await expect(page.locator('label[for="aicheckbox"]')).not.toBeVisible();
      await expect(page.locator('label[for="datadeletioncheckbox"]')).not.toBeVisible();
    });
  });
});

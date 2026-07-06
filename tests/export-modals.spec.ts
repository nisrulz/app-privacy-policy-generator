import { test, expect } from '@playwright/test';
import { completeWizard, completeWizardNoTracking, completeWizardGDPR } from './helpers';

test.describe('Export and Modals', () => {
  test.describe('T&C Export', () => {
    test('T&C modal HTML export shows raw HTML', async ({ page }) => {
      await completeWizard(page);

      await page.locator('a.card-footer-item').filter({ hasText: 'Terms & Conditions' }).click();
      await expect(page.locator('.modal.is-active')).toBeVisible();

      await page.locator('.modal.is-active').getByRole('button', { name: 'HTML' }).click();
      const textarea = page.locator('#tandc_txtarea');
      await expect(textarea).toBeVisible();
      const html = await textarea.inputValue();
      expect(html).toContain('<');
      expect(html.length).toBeGreaterThan(50);
    });

    test('T&C modal Markdown export shows markdown', async ({ page }) => {
      await completeWizard(page);

      await page.locator('a.card-footer-item').filter({ hasText: 'Terms & Conditions' }).click();
      await expect(page.locator('.modal.is-active')).toBeVisible();

      await page.locator('.modal.is-active').getByRole('button', { name: 'Markdown' }).click();
      const textarea = page.locator('#tandc_txtarea');
      await expect(textarea).toBeVisible();
      const md = await textarea.inputValue();
      expect(md.length).toBeGreaterThan(50);
    });

    test('T&C preview button switches back from HTML/Markdown view', async ({ page }) => {
      await completeWizard(page);

      await page.locator('a.card-footer-item').filter({ hasText: 'Terms & Conditions' }).click();
      await expect(page.locator('.modal.is-active')).toBeVisible();

      await page.locator('.modal.is-active').getByRole('button', { name: 'HTML' }).click();
      await expect(page.locator('#tandc_txtarea')).toBeVisible();

      await page.locator('.modal.is-active').getByRole('button', { name: 'Preview' }).click();
      await expect(page.locator('#tandc_content')).toBeVisible();
      await expect(page.locator('#tandc_txtarea')).not.toBeVisible();
    });
  });

  test.describe('Privacy Policy Export', () => {
    test('Simple policy preview button switches back from HTML view', async ({ page }) => {
      await completeWizard(page);

      await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
      await expect(page.locator('.modal.is-active')).toBeVisible();

      await page.locator('.modal.is-active').getByRole('button', { name: 'HTML' }).click();
      await expect(page.locator('#privacy_simple_txtarea')).toBeVisible();

      await page.locator('.modal.is-active').getByRole('button', { name: 'Preview' }).click();
      await expect(page.locator('#privacy_simple_content')).toBeVisible();
    });

    test('No Tracking policy has HTML and Markdown export', async ({ page }) => {
      await completeWizardNoTracking(page);

      await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
      await expect(page.locator('.modal.is-active')).toBeVisible();

      await expect(page.locator('.modal.is-active').getByRole('button', { name: 'HTML' })).toBeVisible();
      await expect(page.locator('.modal.is-active').getByRole('button', { name: 'Markdown' })).toBeVisible();
    });

    test('GDPR policy has HTML and Markdown export', async ({ page }) => {
      await completeWizardGDPR(page);

      await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
      await expect(page.locator('.modal.is-active')).toBeVisible();

      await expect(page.locator('.modal.is-active').getByRole('button', { name: 'HTML' })).toBeVisible();
      await expect(page.locator('.modal.is-active').getByRole('button', { name: 'Markdown' })).toBeVisible();
    });
  });

  test.describe('FAQ and Disclaimer Modals', () => {
    test('can open and close FAQ modal from step 1', async ({ page }) => {
      await page.goto('/');
      await page.waitForLoadState('networkidle');

      await page.locator('a.has-info').first().click();
      await expect(page.locator('.modal.is-active')).toBeVisible();

      await page.locator('.modal.is-active button.delete').click();
      await expect(page.locator('.modal.is-active')).not.toBeVisible();
    });

    test('can open and close Disclaimer modal from step 1', async ({ page }) => {
      await page.goto('/');
      await page.waitForLoadState('networkidle');

      const disclaimerLink = page.locator('#step-1 a').filter({ hasText: /Disclaimer|Haftungsausschluss/ }).first();
      await disclaimerLink.click();
      await expect(page.locator('.modal.is-active')).toBeVisible();

      await page.locator('.modal.is-active button.delete').click();
      await expect(page.locator('.modal.is-active')).not.toBeVisible();
    });

    test('FAQ modal can be closed with OK button', async ({ page }) => {
      await page.goto('/');
      await page.waitForLoadState('networkidle');

      await page.locator('a.has-info').first().click();
      await expect(page.locator('.modal.is-active')).toBeVisible();

      await page.locator('.modal.is-active button.is-info').click();
      await expect(page.locator('.modal.is-active')).not.toBeVisible();
    });
  });
});

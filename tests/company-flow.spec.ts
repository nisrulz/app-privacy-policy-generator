import { test, expect } from '@playwright/test';
import { gotoApp, clickStart, clickNext, expectStep, fillStep2, navigateToStep6, completeWizardCompany } from './helpers';

test.describe('Company Developer Flow', () => {
  test('step 6 shows company name field when Company is selected', async ({ page }) => {
    await navigateToStep6(page);
    await expectStep(page, 6);

    await page.locator('input[type="radio"][value="Company"]').check();

    await expect(page.locator('#companyName')).toBeVisible();
    await expect(page.locator('#devName')).not.toBeVisible();
  });

  test('step 6 shows dev name field when Individual is selected', async ({ page }) => {
    await navigateToStep6(page);
    await expectStep(page, 6);

    await page.locator('input[type="radio"][value="Individual"]').check();

    await expect(page.locator('#devName')).toBeVisible();
    await expect(page.locator('#companyName')).not.toBeVisible();
  });

  test('step 6 blocks advancement without company name when Company selected', async ({ page }) => {
    await navigateToStep6(page);
    await expectStep(page, 6);

    await page.locator('input[type="radio"][value="Company"]').check();

    const nextBtn = page.locator('a.card-footer-item').filter({ hasText: 'Next' });
    await expect(nextBtn).toHaveClass(/is-disabled/);

    await page.locator('#companyName').fill('Acme Corp');
    await expect(nextBtn).not.toHaveClass(/is-disabled/);
  });

  test('company name appears in generated privacy policy', async ({ page }) => {
    await completeWizardCompany(page);

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await expect(page.locator('.modal.is-active')).toBeVisible();

    const content = await page.locator('#privacy_simple_content').textContent();
    expect(content).toContain('Acme Corp');
    expect(content).not.toContain('[Developer/Company name]');
  });

  test('company name appears in Terms & Conditions', async ({ page }) => {
    await completeWizardCompany(page);

    await page.locator('a.card-footer-item').filter({ hasText: 'Terms & Conditions' }).click();
    await expect(page.locator('.modal.is-active')).toBeVisible();

    const content = await page.locator('#tandc_content').textContent();
    expect(content).toContain('Acme Corp');
  });
});

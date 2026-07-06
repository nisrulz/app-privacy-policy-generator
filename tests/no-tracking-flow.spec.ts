import { test, expect } from '@playwright/test';
import { gotoApp, clickStart, clickNext, clickPrevious, expectStep, fillStep2, completeWizardNoTracking } from './helpers';

test.describe('No Tracking Flow', () => {
  test('navigates full wizard with No Tracking policy', async ({ page }) => {
    await gotoApp(page);
    await clickStart(page);
    await expectStep(page, 2);

    await page.locator('#appName').fill('NoTrack App');
    await page.locator('#appContact').fill('notrack@example.com');
    await page.locator('input[type="radio"][value="2"]').check();
    await clickNext(page);
    await expectStep(page, 3);

    await clickNext(page);
    await expectStep(page, 4);

    await clickNext(page);
    await expectStep(page, 5);

    await clickNext(page);
    await expectStep(page, 6);

    await page.locator('#devName').fill('NoTrack Dev');
    await clickNext(page);
    await expectStep(page, 7);

    await clickNext(page);
    await expectStep(page, 8);
  });

  test('step 3 hides PID info field for No Tracking', async ({ page }) => {
    await gotoApp(page);
    await clickStart(page);
    await expectStep(page, 2);

    await page.locator('#appName').fill('Test App');
    await page.locator('#appContact').fill('test@example.com');
    await page.locator('input[type="radio"][value="2"]').check();
    await clickNext(page);
    await expectStep(page, 3);

    await expect(page.locator('#pidInfoIn')).not.toBeVisible();
  });

  test('step 4 hides feature checkboxes for No Tracking', async ({ page }) => {
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

    await expect(page.locator('#locationcheckbox')).not.toBeVisible();
    await expect(page.locator('#aicheckbox')).not.toBeVisible();
    await expect(page.locator('#datadeletioncheckbox')).not.toBeVisible();
  });

  test('opens No Tracking privacy policy modal', async ({ page }) => {
    await completeWizardNoTracking(page);

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await expect(page.locator('.modal.is-active')).toBeVisible();

    const content = page.locator('#privacy_notrack_content');
    await expect(content).toBeVisible();
    const text = await content.textContent();
    expect(text).toContain('Test App');
    expect(text).toContain('test@example.com');
  });

  test('No Tracking privacy policy has HTML and Markdown export', async ({ page }) => {
    await completeWizardNoTracking(page);

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await expect(page.locator('.modal.is-active')).toBeVisible();

    await expect(page.locator('.modal.is-active').getByRole('button', { name: 'HTML' })).toBeVisible();
    await expect(page.locator('.modal.is-active').getByRole('button', { name: 'Markdown' })).toBeVisible();
  });

  test('No Tracking policy shows no third-party services on step 7', async ({ page }) => {
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

import { test, expect } from '@playwright/test';
import { gotoApp, clickStart, clickNext, expectStep, fillStep2, navigateToStep5, completeWizard } from './helpers';

test.describe('Platform Selection', () => {
  test('step 5 shows all platform checkboxes', async ({ page }) => {
    await navigateToStep5(page);

    await expect(page.locator('label[for="platform-android"]')).toBeVisible();
    await expect(page.locator('label[for="platform-ios"]')).toBeVisible();
    await expect(page.locator('label[for="platform-kaios"]')).toBeVisible();
    await expect(page.locator('label[for="platform-windows"]')).toBeVisible();
    await expect(page.locator('label[for="platform-web"]')).toBeVisible();
  });

  test('Android is checked by default', async ({ page }) => {
    await navigateToStep5(page);

    await expect(page.locator('#platform-android')).toBeChecked();
    await expect(page.locator('#platform-ios')).not.toBeChecked();
  });

  test('can toggle multiple platforms', async ({ page }) => {
    await navigateToStep5(page);

    await page.locator('label[for="platform-ios"]').click();
    await page.locator('label[for="platform-web"]').click();

    await expect(page.locator('#platform-android')).toBeChecked();
    await expect(page.locator('#platform-ios')).toBeChecked();
    await expect(page.locator('#platform-web')).toBeChecked();
  });

  test('selected platforms label updates', async ({ page }) => {
    await navigateToStep5(page);

    const label = page.locator('#step-5 .tag.is-primary.is-light');
    const initialText = await label.textContent();

    await page.locator('label[for="platform-ios"]').click();
    const updatedText = await label.textContent();
    expect(updatedText).not.toBe(initialText);
    expect(updatedText).toContain('iOS');
  });

  test('can uncheck Android', async ({ page }) => {
    await navigateToStep5(page);

    await page.locator('label[for="platform-android"]').click();
    await expect(page.locator('#platform-android')).not.toBeChecked();
  });

  test('platform info reflects in generated policy', async ({ page }) => {
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

    await page.locator('label[for="platform-ios"]').click();
    await page.locator('label[for="platform-web"]').click();
    await clickNext(page);
    await expectStep(page, 6);

    await page.locator('#devName').fill('Platform Test');
    await clickNext(page);
    await expectStep(page, 7);
    await clickNext(page);
    await expectStep(page, 8);

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await expect(page.locator('.modal.is-active')).toBeVisible();

    const content = await page.locator('#privacy_simple_content').textContent();
    expect(content).toContain('mobile devices');
    expect(content).toContain('web browsers');
  });
});

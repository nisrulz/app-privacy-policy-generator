import { test, expect } from '@playwright/test';
import { gotoApp, clickStart, clickNext, expectStep, completeWizard } from './helpers';

test.describe('Policy Output', () => {
  test('generates and displays Simple privacy policy', async ({ page }) => {
    await completeWizard(page);

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await expect(page.locator('.modal.is-active')).toBeVisible();

    const content = page.locator('#privacy_simple_content');
    await expect(content).toBeVisible();
    const text = await content.textContent();
    expect(text).toContain('Test App');
    expect(text).toContain('test@example.com');
  });

  test('generates and displays Terms & Conditions', async ({ page }) => {
    await completeWizard(page);

    await page.locator('a.card-footer-item').filter({ hasText: 'Terms & Conditions' }).click();
    await expect(page.locator('.modal.is-active')).toBeVisible();

    const content = page.locator('#tandc_content');
    await expect(content).toBeVisible();
    const text = await content.textContent();
    expect(text).toContain('Test App');
  });

  test('can switch to HTML view in privacy policy modal', async ({ page }) => {
    await completeWizard(page);

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await expect(page.locator('.modal.is-active')).toBeVisible();

    await page.locator('.modal.is-active').getByRole('button', { name: 'HTML' }).click();
    const textarea = page.locator('#privacy_simple_txtarea');
    await expect(textarea).toBeVisible();
    const html = await textarea.inputValue();
    expect(html).toContain('<');
  });

  test('can switch to Markdown view in privacy policy modal', async ({ page }) => {
    await completeWizard(page);

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await expect(page.locator('.modal.is-active')).toBeVisible();

    await page.locator('.modal.is-active').getByRole('button', { name: 'Markdown' }).click();
    const textarea = page.locator('#privacy_simple_txtarea');
    await expect(textarea).toBeVisible();
    const md = await textarea.inputValue();
    expect(md.length).toBeGreaterThan(0);
  });

  test('can close modal with close button', async ({ page }) => {
    await completeWizard(page);

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await expect(page.locator('.modal.is-active')).toBeVisible();

    await page.locator('.modal.is-active button[aria-label="close"]').click();
    await expect(page.locator('.modal.is-active')).not.toBeVisible();
  });

  test('GDPR policy includes business address', async ({ page }) => {
    await gotoApp(page);
    await clickStart(page);
    await expectStep(page, 2);

    await page.locator('#appName').fill('Test App');
    await page.locator('#appContact').fill('test@example.com');
    await page.locator('input[type="radio"][value="3"]').check();
    await clickNext(page);
    await expectStep(page, 3);

    await page.locator('#businessAddress').fill('123 Test Street, Berlin, Germany');
    await clickNext(page);
    await expectStep(page, 4);

    await clickNext(page);
    await expectStep(page, 5);

    await clickNext(page);
    await expectStep(page, 6);

    await page.locator('#devName').fill('Test Dev');
    await clickNext(page);
    await expectStep(page, 7);

    await clickNext(page);
    await expectStep(page, 8);

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await expect(page.locator('.modal.is-active')).toBeVisible();

    const content = await page.locator('#privacy_gdpr_content').textContent();
    expect(content).toContain('123 Test Street, Berlin, Germany');
  });

  test('policy includes app name and developer name', async ({ page }) => {
    await completeWizard(page);

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    const text = await page.locator('#privacy_simple_content').textContent();
    expect(text).toContain('Test App');
    expect(text).toContain('John Doe');
  });

  test('T&C modal has HTML and Markdown export buttons', async ({ page }) => {
    await completeWizard(page);

    await page.locator('a.card-footer-item').filter({ hasText: 'Terms & Conditions' }).click();
    await expect(page.locator('.modal.is-active')).toBeVisible();

    await expect(page.locator('.modal.is-active').getByRole('button', { name: 'HTML' })).toBeVisible();
    await expect(page.locator('.modal.is-active').getByRole('button', { name: 'Markdown' })).toBeVisible();
  });
});

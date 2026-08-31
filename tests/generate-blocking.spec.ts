import { test, expect } from '@playwright/test';
import { gotoApp, clickStart, clickNext, expectStep, fillStep2, completeWizard } from './helpers';

test.describe('Generate Blocking', () => {
  test('modal does not open when required fields are cleared', async ({ page }) => {
    await completeWizard(page);

    await page.evaluate(() => {
      const app = (document.getElementById('app') as any).__vue_app__;
      const proxy = app._container._vnode.component.proxy;
      proxy.appName = '';
      proxy.appContact = '';
    });
    await page.waitForTimeout(200);

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click({ force: true });
    await page.waitForTimeout(500);
    await expect(page.locator('.modal.is-active')).not.toBeVisible();
  });

  test('T&C modal does not open when required fields are cleared', async ({ page }) => {
    await completeWizard(page);

    await page.evaluate(() => {
      const app = (document.getElementById('app') as any).__vue_app__;
      const proxy = app._container._vnode.component.proxy;
      proxy.appName = '';
      proxy.appContact = '';
    });
    await page.waitForTimeout(200);

    await page.locator('a.card-footer-item').filter({ hasText: 'Terms & Conditions' }).click({ force: true });
    await page.waitForTimeout(500);
    await expect(page.locator('.modal.is-active')).not.toBeVisible();
  });

  test('modal opens when all required fields are filled', async ({ page }) => {
    await completeWizard(page);

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await expect(page.locator('.modal.is-active')).toBeVisible();
  });
});

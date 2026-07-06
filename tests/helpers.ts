import { type Page, expect } from '@playwright/test';

export async function gotoApp(page: Page) {
  await page.goto('/');
  await page.waitForLoadState('networkidle');
}

export async function clickNext(page: Page) {
  await page.locator('a.card-footer-item').filter({ hasText: 'Next' }).click();
  await page.waitForTimeout(200);
}

export async function clickPrevious(page: Page) {
  await page.locator('a.card-footer-item').filter({ hasText: 'Previous' }).click();
  await page.waitForTimeout(200);
}

export async function clickStart(page: Page) {
  await page.locator('button.start-btn').click();
  await page.waitForTimeout(200);
}

export async function expectStep(page: Page, step: number) {
  await expect(page.locator(`#step-${step}`)).toBeVisible();
}

export async function navigateToStep2(page: Page) {
  await gotoApp(page);
  await clickStart(page);
  await expectStep(page, 2);
}

export async function fillStep2(page: Page) {
  await page.locator('#appName').fill('Test App');
  await page.locator('#appContact').fill('test@example.com');
}

export async function navigateToStep6(page: Page) {
  await navigateToStep2(page);
  await fillStep2(page);
  await clickNext(page);
  await expectStep(page, 3);
  await clickNext(page);
  await expectStep(page, 4);
  await clickNext(page);
  await expectStep(page, 5);
  await clickNext(page);
  await expectStep(page, 6);
}

export async function navigateToStep7(page: Page) {
  await navigateToStep6(page);
  await page.locator('input[type="radio"][value="Individual"]').check();
  await page.locator('#devName').fill('John Doe');
  await clickNext(page);
  await expectStep(page, 7);
}

export async function completeWizard(page: Page) {
  await navigateToStep7(page);
  await clickNext(page);
  await expectStep(page, 8);
}

export async function navigateToStep4(page: Page) {
  await navigateToStep2(page);
  await fillStep2(page);
  await clickNext(page);
  await expectStep(page, 3);
  await clickNext(page);
  await expectStep(page, 4);
}

export async function navigateToStep5(page: Page) {
  await navigateToStep4(page);
  await clickNext(page);
  await expectStep(page, 5);
}

export async function completeWizardNoTracking(page: Page) {
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
  await page.locator('#devName').fill('John Doe');
  await clickNext(page);
  await expectStep(page, 7);
  await clickNext(page);
  await expectStep(page, 8);
}

export async function completeWizardCompany(page: Page) {
  await gotoApp(page);
  await clickStart(page);
  await expectStep(page, 2);
  await page.locator('#appName').fill('Test App');
  await page.locator('#appContact').fill('test@example.com');
  await clickNext(page);
  await expectStep(page, 3);
  await clickNext(page);
  await expectStep(page, 4);
  await clickNext(page);
  await expectStep(page, 5);
  await clickNext(page);
  await expectStep(page, 6);
  await page.locator('input[type="radio"][value="Company"]').check();
  await page.locator('#companyName').fill('Acme Corp');
  await clickNext(page);
  await expectStep(page, 7);
  await clickNext(page);
  await expectStep(page, 8);
}

export async function completeWizardGDPR(page: Page, opts?: { businessAddress?: string; euRepresentative?: string }) {
  await gotoApp(page);
  await clickStart(page);
  await expectStep(page, 2);
  await page.locator('#appName').fill('GDPR App');
  await page.locator('#appContact').fill('gdpr@example.com');
  await page.locator('input[type="radio"][value="3"]').check();
  if (opts?.euRepresentative) {
    await page.locator('#euRepresentative').fill(opts.euRepresentative);
  }
  await clickNext(page);
  await expectStep(page, 3);
  await page.locator('#businessAddress').fill(opts?.businessAddress || '123 EU Street, Berlin');
  await clickNext(page);
  await expectStep(page, 4);
  await clickNext(page);
  await expectStep(page, 5);
  await clickNext(page);
  await expectStep(page, 6);
  await page.locator('#devName').fill('GDPR Dev');
  await clickNext(page);
  await expectStep(page, 7);
  await clickNext(page);
  await expectStep(page, 8);
}

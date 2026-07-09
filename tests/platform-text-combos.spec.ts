import { test, expect } from '@playwright/test';
import { gotoApp, clickStart, clickNext, expectStep, fillStep2, navigateToStep5, completeWizard, completeWizardNoTracking, navigateToStep6 } from './helpers';

async function selectPlatforms(page, toSelect) {
  var defaults = { android: true, ios: false, kaios: false, windows: false, web: false };
  var platforms = Object.assign({}, defaults, toSelect);
  for (var key of Object.keys(platforms)) {
    var el = page.locator('#platform-' + key);
    var isChecked = await el.isChecked();
    if (platforms[key] !== isChecked) {
      await page.locator('label[for="platform-' + key + '"]').click();
    }
  }
}

async function completeToStep8(page) {
  await clickNext(page);
  await expectStep(page, 6);
  await page.locator('#devName').fill('Test Dev');
  await clickNext(page);
  await expectStep(page, 7);
  await clickNext(page);
  await expectStep(page, 8);
}

async function openPrivacyModal(page) {
  await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
  await page.waitForTimeout(500);
  await expect(page.locator('.modal.is-active')).toBeVisible();
}

async function getSimplePolicyText(page) {
  await openPrivacyModal(page);
  return await page.locator('#privacy_simple_content').innerText();
}

async function getGDPRPolicyText(page) {
  await openPrivacyModal(page);
  return await page.locator('#privacy_gdpr_content').innerText();
}

async function getTandCText(page) {
  await page.locator('a.card-footer-item').filter({ hasText: 'Terms & Conditions' }).click();
  await page.waitForTimeout(500);
  await expect(page.locator('.modal.is-active')).toBeVisible();
  return await page.locator('#tandc_content').innerText();
}

async function completeWizardGDPR(page) {
  await gotoApp(page);
  await clickStart(page);
  await expectStep(page, 2);
  await page.locator('#appName').fill('GDPR App');
  await page.locator('#appContact').fill('gdpr@example.com');
  await page.locator('input[type="radio"][value="3"]').check();
  await clickNext(page);
  await expectStep(page, 3);
  await page.locator('#businessAddress').fill('123 EU Street, Berlin');
  await clickNext(page);
  await expectStep(page, 4);
  await clickNext(page);
  await expectStep(page, 5);
}

test.describe('Platform text combinations', () => {
  test('web-only shows web browsers text in policy', async ({ page }) => {
    await navigateToStep5(page);
    await page.locator('label[for="platform-android"]').click();
    await page.locator('label[for="platform-web"]').click();
    await clickNext(page);
    await expectStep(page, 6);
    await page.locator('#devName').fill('John Doe');
    await clickNext(page);
    await expectStep(page, 7);
    await clickNext(page);
    await expectStep(page, 8);

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await page.waitForTimeout(500);
    await expect(page.locator('.modal.is-active')).toBeVisible();
    const content = await page.locator('#privacy_simple_content').innerText();
    expect(content).toContain('web browsers');
  });

  test('multiple platforms joined with commas and conjunction', async ({ page }) => {
    await navigateToStep5(page);
    await page.locator('label[for="platform-android"]').click();
    await page.locator('label[for="platform-ios"]').click();
    await page.locator('label[for="platform-web"]').click();
    await clickNext(page);
    await expectStep(page, 6);
    await page.locator('#devName').fill('John Doe');
    await clickNext(page);
    await expectStep(page, 7);
    await clickNext(page);
    await expectStep(page, 8);

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await page.waitForTimeout(500);
    await expect(page.locator('.modal.is-active')).toBeVisible();
    const content = await page.locator('#privacy_simple_content').innerText();
    expect(content).toContain('mobile devices');
    expect(content).toContain('web browsers');
  });

  test('mobile-only (android+ios) shows "for mobile devices" without leading comma', async ({ page }) => {
    await navigateToStep5(page);
    await page.locator('label[for="platform-ios"]').click();
    await clickNext(page);
    await expectStep(page, 6);
    await page.locator('#devName').fill('Dev Name');
    await clickNext(page);
    await expectStep(page, 7);
    await clickNext(page);
    await expectStep(page, 8);

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await page.waitForTimeout(500);
    await expect(page.locator('.modal.is-active')).toBeVisible();
    const content = await page.locator('#privacy_simple_content').innerText();
    expect(content).toContain('for mobile devices, together');
    expect(content).not.toContain(', and mobile devices');
    expect(content).not.toContain(', or mobile devices');
  });

  test('no-tracking mobile-only uses "mobile device" without leading comma', async ({ page }) => {
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

    await page.locator('label[for="platform-ios"]').click();
    await clickNext(page);
    await expectStep(page, 6);
    await page.locator('#devName').fill('NoTrack Dev');
    await clickNext(page);
    await expectStep(page, 7);
    await clickNext(page);
    await expectStep(page, 8);

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await page.waitForTimeout(500);
    await expect(page.locator('.modal.is-active')).toBeVisible();
    const content = await page.locator('#privacy_notrack_content').innerText();
    expect(content).not.toContain(', or mobile device');
    expect(content).toContain('your mobile device');
  });

  test('windows-only shows "Windows devices" without leading comma', async ({ page }) => {
    await navigateToStep5(page);
    await page.locator('label[for="platform-android"]').click();
    await page.locator('label[for="platform-windows"]').click();
    await clickNext(page);
    await expectStep(page, 6);
    await page.locator('#devName').fill('Dev Name');
    await clickNext(page);
    await expectStep(page, 7);
    await clickNext(page);
    await expectStep(page, 8);

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await page.waitForTimeout(500);
    await expect(page.locator('.modal.is-active')).toBeVisible();
    const content = await page.locator('#privacy_simple_content').innerText();
    expect(content).toContain('for Windows devices, together');
    expect(content).not.toContain(', and Windows devices');
  });

  test('single platform text has no stray commas or conjunctions at start', async ({ page }) => {
    await completeWizard(page);

    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await page.waitForTimeout(500);
    await expect(page.locator('.modal.is-active')).toBeVisible();
    const content = await page.locator('#privacy_simple_content').innerText();
    expect(content).not.toMatch(/, (and|or) mobile/);
  });

  test.describe('all platform category combos produce correct text', () => {
    test('web-only intro says "for web browsers"', async ({ page }) => {
      await navigateToStep5(page);
      await selectPlatforms(page, { android: false, web: true });
      await completeToStep8(page);
      var text = await getSimplePolicyText(page);
      expect(text).toContain('for web browsers, together');
      expect(text).not.toContain(', and web');
    });

    test('mobile+windows uses "and" in intro', async ({ page }) => {
      await navigateToStep5(page);
      await selectPlatforms(page, { ios: true, windows: true });
      await completeToStep8(page);
      var text = await getSimplePolicyText(page);
      expect(text).toContain('for mobile devices and Windows devices');
    });

    test('mobile+web uses "and" in intro', async ({ page }) => {
      await navigateToStep5(page);
      await selectPlatforms(page, { ios: true, web: true });
      await completeToStep8(page);
      var text = await getSimplePolicyText(page);
      expect(text).toContain('for mobile devices and web browsers');
    });

    test('windows+web uses "and" in intro', async ({ page }) => {
      await navigateToStep5(page);
      await selectPlatforms(page, { android: false, windows: true, web: true });
      await completeToStep8(page);
      var text = await getSimplePolicyText(page);
      expect(text).toContain('for Windows devices and web browsers');
    });

    test('all three uses commas and final "and" in intro', async ({ page }) => {
      await navigateToStep5(page);
      await selectPlatforms(page, { ios: true, windows: true, web: true });
      await completeToStep8(page);
      var text = await getSimplePolicyText(page);
      expect(text).toContain('mobile devices, Windows devices, and web browsers');
    });

    test('web-only no-tracking uses "computer" for location', async ({ page }) => {
      await gotoApp(page);
      await clickStart(page);
      await expectStep(page, 2);
      await page.locator('#appName').fill('NoTrack Web');
      await page.locator('#appContact').fill('web@example.com');
      await page.locator('input[type="radio"][value="2"]').check();
      await clickNext(page);
      await expectStep(page, 3);
      await clickNext(page);
      await expectStep(page, 4);
      await clickNext(page);
      await expectStep(page, 5);
      await selectPlatforms(page, { android: false, web: true });
      await clickNext(page);
      await expectStep(page, 6);
      await page.locator('#devName').fill('Web Dev');
      await clickNext(page);
      await expectStep(page, 7);
      await clickNext(page);
      await expectStep(page, 8);
      await openPrivacyModal(page);
      var content = await page.locator('#privacy_notrack_content').innerText();
      expect(content).toContain('your computer');
      expect(content).not.toMatch(/, (and|or) (computer|mobile)/);
    });
  });

  test.describe('platform text correct in all policy types', () => {
    test('GDPR mobile-only intro uses "for mobile devices"', async ({ page }) => {
      await completeWizardGDPR(page);
      await selectPlatforms(page, { ios: true });
      await completeToStep8(page);
      var text = await getGDPRPolicyText(page);
      expect(text).toContain('for mobile devices, together');
      expect(text).not.toContain(', and mobile devices');
    });

    test('GDPR autoCollect uses correct device type', async ({ page }) => {
      await completeWizardGDPR(page);
      await selectPlatforms(page, { ios: true, windows: true });
      await completeToStep8(page);
      var text = await getGDPRPolicyText(page);
      expect(text).toContain('mobile device or Windows device');
    });

    test('T&C mobile-only intro uses "for mobile devices"', async ({ page }) => {
      await navigateToStep5(page);
      await selectPlatforms(page, { ios: true });
      await completeToStep8(page);
      var text = await getTandCText(page);
      expect(text).toContain('for mobile devices, together');
      expect(text).not.toContain(', and mobile devices');
    });

    test('T&C web-only intro uses "for web browsers"', async ({ page }) => {
      await navigateToStep5(page);
      await selectPlatforms(page, { android: false, web: true });
      await completeToStep8(page);
      var text = await getTandCText(page);
      expect(text).toContain('for web browsers, together');
      expect(text).not.toContain(', and web');
    });
  });
});

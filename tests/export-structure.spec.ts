import { test, expect } from '@playwright/test';
import { completeWizard } from './helpers';

test.describe('HTML/Markdown export structure', () => {
  test('HTML export contains complete document structure', async ({ page }) => {
    await completeWizard(page);
    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await page.waitForTimeout(500);
    await expect(page.locator('.modal.is-active')).toBeVisible();

    // Switch to HTML view
    await page.locator('.modal.is-active button.button.is-primary').filter({ hasText: 'HTML' }).click();
    await page.waitForTimeout(300);

    const html = await page.locator('#privacy_simple_txtarea').inputValue();
    expect(html).toContain('<!DOCTYPE html>');
    expect(html).toContain("<meta charset='utf-8'>");
    expect(html).toContain("<meta name='viewport'");
    expect(html).toContain('<title>');
    expect(html).toContain('<style>');
    expect(html).toContain('<body>');
    expect(html).toContain('</body>');
    expect(html).toContain('</html>');
  });

  test('Markdown export contains valid markdown formatting', async ({ page }) => {
    await completeWizard(page);
    await page.locator('a.card-footer-item').filter({ hasText: 'Privacy Policy' }).click();
    await page.waitForTimeout(500);
    await expect(page.locator('.modal.is-active')).toBeVisible();

    // Switch to Markdown view
    await page.locator('.modal.is-active button.button.is-primary').filter({ hasText: 'Markdown' }).click();
    await page.waitForTimeout(300);

    const md = await page.locator('#privacy_simple_txtarea').inputValue();
    // Markdown should contain bold markers or list markers
    expect(md.length).toBeGreaterThan(50);
    // Should have some markdown structure (bold, lists, or headers)
    const hasMarkdown = md.includes('**') || md.includes('- ') || md.includes('#');
    expect(hasMarkdown).toBeTruthy();
  });

  test('T&C HTML export contains complete document structure', async ({ page }) => {
    await completeWizard(page);
    await page.locator('a.card-footer-item').filter({ hasText: 'Terms' }).click();
    await page.waitForTimeout(500);
    await expect(page.locator('.modal.is-active')).toBeVisible();

    // Switch to HTML view
    await page.locator('.modal.is-active button.button.is-primary').filter({ hasText: 'HTML' }).click();
    await page.waitForTimeout(300);

    const html = await page.locator('#tandc_txtarea').inputValue();
    expect(html).toContain('<!DOCTYPE html>');
    expect(html).toContain("<meta charset='utf-8'>");
    expect(html).toContain('<title>');
    expect(html).toContain('<body>');
  });
});

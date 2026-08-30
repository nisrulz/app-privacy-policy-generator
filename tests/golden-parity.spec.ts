import { expect, test } from '@playwright/test';
import { readFile } from 'node:fs/promises';
import { resolve } from 'node:path';

const goldenDir = resolve('cmd/build/testdata/golden');

function normalizeHtml(html: string): string {
  return html
    .replace(/\?v=[a-f0-9]+/g, '')
    .replace(/>\s+</g, '><')
    .replace(/\s+/g, ' ')
    .trim();
}

for (const locale of ['en', 'de']) {
  test(`generated ${locale} HTML matches its golden fixture`, async ({ request }) => {
    const path = locale === 'en' ? '/' : `/${locale}/`;
    const response = await request.get(path);
    const actual = await response.text();
    const expected = await readFile(resolve(goldenDir, `${locale}.html`), 'utf8');

    expect(response.ok()).toBeTruthy();
    expect(normalizeHtml(actual)).toBe(normalizeHtml(expected));
  });
}

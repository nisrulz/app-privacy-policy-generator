import { expect, test } from '@playwright/test';
import { readFile } from 'node:fs/promises';
import { resolve } from 'node:path';

let precacheUrls: string[] = [];

test.beforeAll(async () => {
  const source = await readFile(resolve('public/sw.js'), 'utf8');
  const block = source.match(/const PRECACHE_URLS = \[([\s\S]*?)\]/);
  if (block) {
    precacheUrls = Array.from(block[1].matchAll(/"([^"]+)"/g), (m) => m[1]);
  }
});

test('precache list is well-formed', () => {
  expect(precacheUrls.length, 'PRECACHE_URLS parsed from public/sw.js').toBeGreaterThan(0);
  expect(new Set(precacheUrls).size, 'no duplicate entries').toBe(precacheUrls.length);
  for (const url of precacheUrls) {
    expect(url.startsWith('/') && !url.startsWith('//'), `${url} is site-relative`).toBe(true);
  }
});

test('every precached URL resolves in the built output', async ({ request }) => {
  for (const url of precacheUrls) {
    const response = await request.get(url);
    expect(response.ok(), `${url} should be served`).toBeTruthy();
  }
});

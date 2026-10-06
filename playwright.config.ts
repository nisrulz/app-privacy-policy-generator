import { defineConfig, devices } from '@playwright/test';

export default defineConfig({
  testDir: './tests',
  fullyParallel: true,
  forbidOnly: !!process.env.CI,
  retries: process.env.CI ? 2 : 0,
  workers: process.env.CI ? Number(process.env.PLAYWRIGHT_WORKERS || 2) : undefined,
  reporter: [['list'], ['html', { open: 'never' }]],
  use: {
    baseURL: 'http://localhost:8000',
    trace: 'on-first-retry',
  },
  projects: [
    {
      name: 'chromium',
      use: { ...devices['Desktop Chrome'] },
    },
  ],
  webServer: {
    // exec the server binary directly rather than going through `make serve`.
    // Playwright tears down only the process it spawned; behind make, the
    // server binary is a grandchild and survives as an orphan holding port
    // 8000, which then breaks the next `make serve`. `exec` replaces the
    // shell so the process Playwright owns is the server itself.
    command:
      'go build -o .bin/serve ./cmd/build && exec .bin/serve -serve -port 8000',
    url: 'http://localhost:8000',
    reuseExistingServer: !process.env.CI,
    timeout: 120000,
  },
});

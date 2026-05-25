# Reviews Page Generator

Generates a static Guest Book page from [GitHub Issue #65](https://github.com/nisrulz/app-privacy-policy-generator/issues/65) comments. Requires [`uv`](https://docs.astral.sh/uv/) (single Rust binary, no Python needed).

## How it works

1. Fetches all comment pages from the GitHub Issues API (cached as JSON for 1 week)
2. Downloads & resizes profile pictures (48×48)
3. Downloads embedded images from comment bodies
4. Converts comment markdown → HTML
5. Renders `template.mustache` → `reviews.html`
6. Copies assets into `../../public/`

## Usage

```bash
# Normal run (uses cached data if less than 1 week old)
./gen_reviews_page.sh

# Force re-fetch from GitHub
./gen_reviews_page.sh -f
```

## Requirements

- [`uv`](https://docs.astral.sh/uv/#installation) — install once: `curl -LsSf https://astral.sh/uv/install.sh | sh`
- Everything else (Python, venv, deps) is managed automatically by `uv`

## Files

| File | Purpose |
|------|---------|
| `fetch_and_generate.py` | Main script (uv inline deps, runs with `uv run`) |
| `gen_reviews_page.sh` | Entry point — `-f` forces re-fetch, runs script, copies output |
| `template.mustache` | Mustache template for the reviews page |
| `comments_json/` | Cached API responses (gitignored after generation) |
| `profile_pictures/` | Downloaded & resized author avatars |
| `downloaded_images/` | Images embedded in comments |

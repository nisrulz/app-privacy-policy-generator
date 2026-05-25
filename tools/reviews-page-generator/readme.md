# Reviews Page Generator

Generates a static Guest Book page from [GitHub Issue #65](https://github.com/nisrulz/app-privacy-policy-generator/issues/65) comments.

## Quick start

```bash
# From repo root — uses cached data (if < 1 week old)
./gen_reviews_page.sh

# Force re-fetch from GitHub API
./gen_reviews_page.sh -f
```

## How it works

| Step | Action |
|------|--------|
| 1 | Fetch comment pages from the GitHub Issues API (paginated, cached as JSON for 1 week) |
| 2 | Download & resize profile pictures (48×48) in parallel (8 workers) |
| 3 | Download embedded images in parallel |
| 4 | Convert markdown bodies → HTML, format timestamps, remap reactions |
| 5 | Render `template.mustache` → `reviews.html` |
| 6 | Copy generated assets into `../../public/` |

## Requirements

- [`uv`](https://docs.astral.sh/uv/#installation) — install once:
  ```bash
  curl -LsSf https://astral.sh/uv/install.sh | sh
  ```
- Python, venv, and all dependencies are managed automatically by `uv`

## Files

| File | Purpose |
|------|---------|
| `fetch_and_generate.py` | Main script — fetch, process, render, deploy |
| `gen_reviews_page.sh` | Entry point (at repo root) |
| `template.mustache` | Mustache template |
| `comments_json/` | Cached API responses |
| `profile_pictures/` | Resized author avatars |
| `downloaded_images/` | Images embedded in comments |

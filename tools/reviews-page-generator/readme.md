# Reviews Page Generator

Generates a static Guest Book page from [GitHub Issue #65](https://github.com/nisrulz/app-privacy-policy-generator/issues/65) comments.

## Quick start

```bash
# From repo root — uses cached data (if < 1 week old)
./scripts/gen_reviews_page.sh

# Force re-fetch from GitHub API
./scripts/gen_reviews_page.sh -f
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

- [Go](https://go.dev/dl/) 1.25+
- All dependencies managed via `go.mod`

## Files

| File | Purpose |
|------|---------|
| `main.go` | Main program — fetch, process, render, deploy |
| `go.mod` / `go.sum` | Go module definition |
| `scripts/gen_reviews_page.sh` | Entry point (at repo root) |
| `template.mustache` | Mustache template (rendered via string substitution) |
| `comments_json/` | Cached API responses |

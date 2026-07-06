# Reviews Page Generator

Pulls comments from [GitHub Issue #65](https://github.com/nisrulz/app-privacy-policy-generator/issues/65) and generates a static Guestbook page.

## Quick start

From the repo root:

```sh
❯ make reviews          # uses cached data (if < 1 week old)
❯ make reviews-force    # re-fetches everything from GitHub
```

Or run the script directly:

```sh
❯ ./scripts/gen_reviews_page.sh
❯ ./scripts/gen_reviews_page.sh -f    # force re-fetch
```

## How it works

1. Fetches comment pages from the GitHub Issues API (paginated, cached as JSON for 1 week)
2. Downloads and resizes profile pictures (48x48) in parallel
3. Downloads embedded images in parallel
4. Converts markdown bodies to HTML, formats timestamps, remaps reactions
5. Renders `template.mustache` into `public/reviews.html` and `public/reviews-data.json`
6. Copies downloaded images into `public/profile_pictures/`

## Requirements

- [Go](https://go.dev/dl/) 1.25+
- All dependencies are in `go.mod`

## Files

| File | What it does |
|------|-------------|
| `main.go` | Fetches, processes, renders, and writes output |
| `go.mod` / `go.sum` | Go module definition |
| `template.mustache` | Mustache template with `{{ total_comments }}` placeholder |
| `comments_json/` | Cached API responses |
| `downloaded_images/` | Raw images pulled from GitHub |

.PHONY: build serve watch firebase-local-preview firebase-deploy reviews reviews-force test test-ui test-debug help

help: ## Show available commands
	@grep -E '^[a-zA-Z_-]+:.*##' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*## "}; {printf "  %-15s %s\n", $$1, $$2}'

build: ## Build the project
	@go run ./cmd/build/

serve: ## Build and serve locally
	@go run ./cmd/build/ -serve

watch: ## Watch for changes and rebuild
	@go run ./cmd/build/ -watch

firebase-local-preview: ## Build and preview via Firebase local server
	@./scripts/firebase_local_preview.sh

firebase-deploy: ## Build and deploy to Firebase Hosting
	@./scripts/firebase_deploy.sh

reviews: ## Generate reviews page from cached data
	@./scripts/gen_reviews_page.sh

reviews-force: ## Generate reviews page (force re-fetch from GitHub)
	@./scripts/gen_reviews_page.sh -f

test: ## Run Playwright E2E tests
	@npx playwright test

test-ui: ## Run Playwright tests in UI mode
	@npx playwright test --ui

test-debug: ## Run Playwright tests in debug mode
	@npx playwright test --debug


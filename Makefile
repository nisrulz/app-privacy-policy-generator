.PHONY: build clean compress-images format check update-deps serve watch deploy firebase-preview reviews test-ui test-debug help

help: ## Show available commands
	@./scripts/help.sh

build: ## Build the project
	@echo "→ Building project..."
	@go run ./cmd/build/
	@echo "✓ Build complete"

clean: ## Clean public/ directory
	@echo "→ Cleaning public/..."
	@go run ./cmd/build/ -clean-only
	@echo "✓ Cleaned"

format: ## Format Go source, templates, and tidy modules
	@echo "→ Formatting..."
	@gofmt -w cmd/build/*.go tools/reviews-page-generator/*.go
	@gotmplfmt -w src/tpl/*.html
	@go mod tidy
	@echo "✓ Format complete"

check: node_modules ## Run Go, golden, and Playwright checks
	@echo "→ Running checks..."
	@go vet ./...
	@go build ./...
	@npx playwright install chromium
	@npx playwright test
	@echo "✓ All checks passed"

compress-images: ## Compress images in public/images
	@echo "→ Compressing images..."
	@./scripts/compress_images.sh
	@echo "✓ Images compressed"

update-deps: ## Update Go dependencies
	@echo "→ Updating dependencies..."
	@./scripts/update_deps.sh
	@echo "✓ Dependencies updated"

serve: ## Build and serve locally || Example: make serve PORT="9090"
	@echo "→ Starting server..."
	@./scripts/serve.sh $(if $(PORT),$(PORT))

watch: ## Watch for changes and rebuild
	@echo "→ Watching for changes..."
	@go run ./cmd/build/ -watch

firebase-preview: ## Build and preview via Firebase local server
	@echo "→ Starting Firebase preview..."
	@./scripts/firebase_local_preview.sh
	@echo "✓ Preview running"

deploy: ## Deploy to Firebase Hosting || Example: make deploy VERSION="3.0.9"
	@echo "→ Deploying to Firebase..."
	@./scripts/firebase_deploy.sh $(VERSION)
	@echo "✓ Deployed"

reviews: ## Generate reviews page || Example: make reviews FORCE="true"
	@echo "→ Generating reviews page..."
	@if [ "$(FORCE)" = "true" ]; then ./scripts/gen_reviews_page.sh -f; else ./scripts/gen_reviews_page.sh; fi
	@echo "✓ Reviews page generated"

test-ui: node_modules ## Run Playwright tests in UI mode
	@npx playwright test --ui

test-debug: node_modules ## Run Playwright tests in debug mode
	@npx playwright test --debug

node_modules: package-lock.json package.json
	@echo "→ Installing Playwright dependencies..."
	@npm ci --no-audit --no-fund

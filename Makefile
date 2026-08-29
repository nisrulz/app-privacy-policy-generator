.PHONY: build clean compress-images format check serve watch deploy reviews test help

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
	@gofmt -w cmd/build/*.go cmd/build/e2e/*.go tools/reviews-page-generator/*.go
	@gotmplfmt -w src/tpl/*.html
	@go mod tidy
	@echo "✓ Format complete"

check: ## Run tests, vet, and build checks
	@echo "→ Running checks..."
	@go test ./cmd/build/ -run TestGolden -v
	@go vet ./...
	@go build ./...
	@echo "✓ All checks passed"

compress-images: ## Compress images in public/images
	@echo "→ Compressing images..."
	@./scripts/compress_images.sh
	@echo "✓ Images compressed"

serve: ## Build and serve locally || Example: make serve PORT="9090"
	@echo "→ Starting server..."
	@go run ./cmd/build/ -serve -port $(PORT)

watch: ## Watch for changes and rebuild
	@echo "→ Watching for changes..."
	@go run ./cmd/build/ -watch

deploy: ## Deploy to Firebase Hosting || Example: make deploy VERSION="3.0.9"
	@echo "→ Deploying to Firebase..."
	@./scripts/firebase_deploy.sh $(VERSION)
	@echo "✓ Deployed"

reviews: ## Generate reviews page || Example: make reviews FORCE="true"
	@echo "→ Generating reviews page..."
	@if [ "$(FORCE)" = "true" ]; then ./scripts/gen_reviews_page.sh -f; else ./scripts/gen_reviews_page.sh; fi
	@echo "✓ Reviews page generated"

test: ## Run E2E tests (requires make serve running)
	@echo "→ Running E2E tests..."
	@go test ./cmd/build/e2e/... -v -count=1 -timeout 5m
	@echo "✓ Tests complete"

TAG:=$(shell git describe --tags)

tag: ## Describes current tag
	@echo ${TAG}

.PHONY: init-dev-environment
init-dev-environment: ## Initializes local development environment after first clone
	@./scripts/install-buf.sh

.PHONY: install-linters
install-linters: ## Install linters and setup environment
	@mkdir -p outputs
	@./scripts/ci/install-linters.sh
	@./scripts/install-buf.sh

.PHONY: format
format: ## Format code
	@./scripts/ci/go-format.sh

.PHONY: lint
lint: ## Lint code
	@./scripts/ci/lint.sh

.PHONY: server-binary
server-binary: ## Builds server binary
	@go build -o build/release-registry cmd/server/main.go

.PHONY: server-image
server-image: ## Builds server image
	@docker build . -f image/Dockerfile -t quay.io/rhacs-eng/release-registry:${TAG}

.PHONY: server-image-push
server-image-push: ## Pushes server image to registry
	@docker push quay.io/rhacs-eng/release-registry:${TAG}

.PHONY: server-helm-deploy
server-helm-deploy: ## Deploys the server with Helm
	envsubst
	@helm upgrade release-registry --install deploy/chart/release-registry --set image.tag=${TAG}

.PHONY: tests-unit
tests-unit: ## Runs all unit tests without cache
	@go test ./pkg/... -count=1

.PHONY: tests-integration
tests-integration: ## Runs all integration tests without cache
	@go test -v ./tests/integration/... -count=1

.PHONY: tests-e2e
tests-e2e: ## Runs all e2e tests without cache
	@go test -v ./tests/e2e/... -count=1

.PHONE: help
help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
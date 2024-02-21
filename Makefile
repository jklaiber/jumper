GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: all build

all: go-gen

install-deps: ## Install dependencies
	go mod tidy
	go install go.uber.org/mock/mockgen@latest
	go install github.com/goreleaser/goreleaser@latest

build: ## Build the binary with goreleaser
	goreleaer build --snapshot --clean

release: ## Make a local release
	goreleaser release --snapshot --clean

go-gen: ## Run go generate
	go generate ./...

test: ## Run tests
	go clean -testcache
	go test ./...

test-coverage: ## Run tests with coverage
	go clean -testcache
	go test ./... -coverprofile=coverage.out

help: ## Show this help message
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_0-9-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)
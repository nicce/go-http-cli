SHELL := /bin/bash

OS := $(shell uname)
ARCHITECTURE := $(shell uname -m | sed "s/x86_64/amd64/g")
GIT_REPO := $(shell git config --get remote.origin.url)
REPO_NAME := $(shell basename ${GIT_REPO} .git)
SHORT_SHA := $(shell git rev-parse --short HEAD)

GOLANGCI_LINT_VERSION := v1.51.0
GOLANGCI_LINT := bin/golangci-lint_$(GOLANGCI_LINT_VERSION)/golangci-lint
GOLINTCI_LINT_URL := https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh

GOTESTSUM_VERSION := 1.9.0
GOTESTSUM := bin/gotestsum_v$(GOTESTSUM_VERSION)/gotestsum
GOTESTSUM_URL := https://github.com/gotestyourself/gotestsum/releases/download/v$(GOTESTSUM_VERSION)/gotestsum_$(GOTESTSUM_VERSION)_$(OS)_$(ARCHITECTURE).tar.gz

ACTIONLINT_VERSION := 1.6.23
ACTIONLINT := bin/actionlint_v$(ACTIONLINT_VERSION)/actionlint
ACTIONLINT_URL :=  https://raw.githubusercontent.com/rhysd/actionlint/main/scripts/download-actionlint.bash

all: help

clean:
	@echo "🚀 Cleaning up old artifacts"
	@rm -f ${REPO_NAME}

## build: Build the application artifacts.
build: clean
	@echo "🚀 Building artifacts"
	@go build -ldflags="-s -w -X '${VERSION_PATH}.Version=${VERSION}' -X '${VERSION_PATH}.Commit=${SHORT_SHA}'" -o bin/${REPO_NAME} .

## run: Run the application
run: build
	@echo "🚀 Running binary"
	@./bin/${REPO_NAME}

## lint: Lint the go code
lint: ${GOLANGCI_LINT}
	@echo "🚀 Linting go code"
	@$(GOLANGCI_LINT) run

## lint-info: Returns information about the current go linter being used
lint-info:
	@echo ${GOLANGCI_LINT}

## gha-lint: Lint the github actions code
gha-lint: ${ACTIONLINT}
	@echo "🚀 Linting github actions code"
	@$(ACTIONLINT)

## gha-linter-info: Returns information about the current github actions linter being used
gha-linter-info:
	@echo ${ACTIONLINT}

## test: Run Go tests
test: ${GOTESTSUM}
	@echo "🚀 Running tests"
	@set -o pipefail; ${GOTESTSUM} --format testname --no-color=false -- -race ./... | grep -v 'EMPTY'; exit $$?

## test-benchmark: Run Go benchmark tests
test-benchmark:
	@echo "🚀 Running benchmark tests"
	@go test -bench=. -benchmem ./...

${GOLANGCI_LINT}:
	@echo "📦 Installing golangci-lint ${GOLANGCI_LINT_VERSION}"
	@mkdir -p $(dir ${GOLANGCI_LINT})
	@curl -sfL ${GOLINTCI_LINT_URL} | sh -s -- -b ./$(patsubst %/,%,$(dir ${GOLANGCI_LINT})) ${GOLANGCI_LINT_VERSION} > /dev/null 2>&1

${GOTESTSUM}:
	@echo "📦 Installing GoTestSum ${GOTESTSUM_VERSION}"
	@mkdir -p $(dir ${GOTESTSUM})
	@curl -sSL ${GOTESTSUM_URL} > bin/gotestsum.tar.gz
	@tar -xzf bin/gotestsum.tar.gz -C $(patsubst %/,%,$(dir ${GOTESTSUM}))
	@rm -f bin/gotestsum.tar.gz

${GOSEC}:
	@echo "📦 Installing gosec v${GOSEC_VERSION}"
	@mkdir -p $(dir ${GOSEC})
	@curl -sSL ${GOSEC_URL} > bin/gosec.tar.gz
	@tar -xzf bin/gosec.tar.gz -C $(patsubst %/,%,$(dir ${GOSEC}))
	@rm -f bin/gosec.tar.gz

${ACTIONLINT}:
	$(call check_shellcheck_installation)
	@echo "📦 Installing actionlint v${ACTIONLINT_VERSION}"
	@mkdir -p $(dir ${ACTIONLINT})
	@bash <(curl -sSL https://raw.githubusercontent.com/rhysd/actionlint/main/scripts/download-actionlint.bash) ${ACTIONLINT_VERSION} $(shell dirname ${ACTIONLINT})  > /dev/null 2>&1

## check_shellcheck_installation: checks if shellcheck is installed. In general it is installed on every ubuntu gha runner.
check_shellcheck_installation:
	@if [[ $(shell which shellcheck) == "" ]]; then \
		echo "the programm 'shellcheck' has to be installed for 'actionlint', because actionlint uses it to lint shell scripts in gha workflows."; \
		exit 1; \
	fi

help: Makefile
	@echo
	@echo "📗 Choose a command run in "${REPO_NAME}":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

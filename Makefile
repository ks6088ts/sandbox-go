GOGET ?= go get "-u"
GOBUILD ?= go build
PKGS ?= $(shell go list ./...)

# https://stackoverflow.com/questions/18136918/how-to-get-current-relative-directory-of-your-makefile
MAKEFILE_PATH ?= $(abspath $(lastword $(MAKEFILE_LIST)))
REPO_NAME ?= $(notdir $(patsubst %/,%,$(dir $(MAKEFILE_PATH))))
LDFLAGS ?= -ldflags="-s -w"
OUT_DIR ?= outputs
PKG_DIR ?= ./cmd/$(REPO_NAME)
BIN_PATH ?= $(OUT_DIR)/$(REPO_NAME)

# https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.DEFAULT_GOAL := help

.PHONY: build
build: ## build applications
	$(GOBUILD) $(LDFLAGS) -o $(BIN_PATH) $(PKG_DIR)

.PHONY: lint-install
lint-install: ## install lint
	@hash golint > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GOGET) golang.org/x/lint/golint; \
	fi

.PHONY: lint
lint: lint-install ## lint codes
	for PKG in $(PKGS); do golint -set_exit_status $$PKG || exit 1; done;

.PHONY: ci
ci: lint build ## run ci tests
	$(BIN_PATH) --help
	$(BIN_PATH) hello --help

# ---
# Cobra: https://github.com/spf13/cobra
# ---

COBRA_CONFIG ?= .cobra.yml
COBRA_CMD_DIR ?= cmd/$(REPO_NAME)
COBRA_PKG_NAME ?= github.com/ks6088ts/$(REPO_NAME)/$(COBRA_CMD_DIR)
COBRA_CMD ?= hello
COBRA_PARENT_CMD ?= rootCmd

.PHONY: cobra-install
cobra-install: ## install cobra
	@hash cobra > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GOGET) github.com/spf13/cobra/cobra; \
	fi

.PHONY: cobra-init
cobra-init: cobra-install ## initialize cobra cli
	mkdir -p $(COBRA_CMD_DIR) && \
	cd $(COBRA_CMD_DIR) && \
	cobra init \
		--pkg-name $(COBRA_PKG_NAME) \
		--config ../../$(COBRA_CONFIG)

.PHONY: cobra-add
cobra-add: cobra-install ## add cobra command
	cd $(COBRA_CMD_DIR) && \
	cobra add $(COBRA_CMD) \
		--config ../../$(COBRA_CONFIG) \
		--parent $(COBRA_PARENT_CMD)

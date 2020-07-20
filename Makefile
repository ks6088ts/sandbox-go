GOGET ?= go get "-u"

# https://stackoverflow.com/questions/18136918/how-to-get-current-relative-directory-of-your-makefile
MAKEFILE_PATH ?= $(abspath $(lastword $(MAKEFILE_LIST)))
REPO_NAME ?= $(notdir $(patsubst %/,%,$(dir $(MAKEFILE_PATH))))

# https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.DEFAULT_GOAL := help


# ---
# Cobra: https://github.com/spf13/cobra
# ---

COBRA_CONFIG ?= .cobra.yml
COBRA_CMD_DIR ?= cmd/$(REPO_NAME)
COBRA_PKG_NAME ?= github.com/ks6088ts/$(REPO_NAME)/$(COBRA_CMD_DIR)

.PHONY: cobra-install
cobra-install: ## install cobra
	@hash cobra > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GOGET) github.com/spf13/cobra/cobra; \
	fi

.PHONY: cobra-init
cobra-init: ## initialize cobra cli
	mkdir -p $(COBRA_CMD_DIR) && \
	cd $(COBRA_CMD_DIR) && \
	cobra init \
		--pkg-name $(COBRA_PKG_NAME) \
		--config ../../$(COBRA_CONFIG)

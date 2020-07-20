GOGET ?= go get "-u"

# https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.DEFAULT_GOAL := help


# ---
# Cobra: https://github.com/spf13/cobra
# ---

.PHONY: cobra-install
cobra-install: ## install cobra
	@hash cobra > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GOGET) github.com/spf13/cobra/cobra; \
	fi

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Always (re)generate version.json via a phony "version" step so builds
# get fresh metadata even if version.json already exists.
.PHONY: version
version: ## Generate/refresh version.json metadata
	@scripts/build-version

.PHONY: docker-build
docker-build: version ## Build service container
	@docker compose build web

.PHONY: mozcloud-postbuild
mozcloud-postbuild: ## Tasks to be run after building service container(s) and before pushing them
	@scripts/verify-container

.PHONY: mozcloud-prebuild
mozcloud-prebuild: ## Tasks to be run before building service container(s)
	@scripts/build-version


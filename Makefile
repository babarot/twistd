TARGETS_NOVENDOR := $(shell glide novendor)

.DEFAULT_GOAL := help

.PHONY: all serve build bundle fmt help

all:

serve: bundle build ## Start to serve twistd
	./twist

build: ## Build go binary
	go build -o twist ./cmd/twistd

bundle: ## Install packages via glide
	glide install

fmt: ## Run go fmt
	@echo $(TARGETS_NOVENDOR) | xargs go fmt

help: ## Self-documented Makefile
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

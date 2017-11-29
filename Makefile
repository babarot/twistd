TARGETS_NOVENDOR := $(shell glide novendor)

.DEFAULT_GOAL := help

.PHONY: all serve build bundle fmt help install cross

all:

serve: bundle build ## Start to serve twistd
	./twist

build: ## Build go binary
	go build -o twist.$(GOOS)-$(GOARCH) ./cmd/twistd

bundle: ## Install packages via glide
	glide install

fmt: ## Run go fmt
	@echo $(TARGETS_NOVENDOR) | xargs go fmt

install: bundle build ## Install command and config file
	install -m 0755 ./twist /usr/bin
	install -m 0644 ./config.toml /etc/twistd.conf

cross: ## Cross-compile
	@$(MAKE) build GOOS=windows GOARCH=amd64
	@$(MAKE) build GOOS=windows GOARCH=386
	@$(MAKE) build GOOS=linux   GOARCH=amd64
	@$(MAKE) build GOOS=linux   GOARCH=386
	@$(MAKE) build GOOS=darwin  GOARCH=amd64
	@$(MAKE) build GOOS=darwin  GOARCH=386

help: ## Self-documented Makefile
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

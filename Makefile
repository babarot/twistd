.DEFAULT_GOAL := help

.PHONY: all build install uninstall help

all:

build:
	mkdir -p /var/log
	mkdir -p /var/run
	cd cmd/twistd && go build

install: build ## Install twistd and init script
	@echo "Installing Components"
	-@cp run.sh /etc/init.d
	-@cp config.toml /etc/twistd.conf
	-@cp cmd/twistd/twistd /etc/rc.d/init.d

uninstall: ## Uninstall twistd and init script
	@echo "Uninstalling complete"

help: ## Self-documented Makefile
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

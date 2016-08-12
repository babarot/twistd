.DEFAULT_GOAL := help

.PHONY: all build install uninstall help

all:

build:
	mkdir -p /var/log
	mkdir -p /var/run
	L="b4b4r07/twistd" bash -c "`curl -L git.io/releases`" -s "os"

install: build ## Install twistd and init script
	@echo "Installing Components"
	-@cp run.sh /etc/init.d/twistd
	-@cp config.toml /etc/twistd.conf
	-@cp cmd/twistd/twistd /etc/rc.d/init.d/twistd

uninstall: ## Uninstall twistd and init script
	@echo "Uninstalling complete"

help: ## Self-documented Makefile
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

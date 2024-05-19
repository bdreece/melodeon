# melodeon v0.1.0
# Copyright (C) 2024 Brian Reece

include scripts/definitions.mk

all: build

# ================ #
# SCRIPTS
# ================ #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage: make [SCRIPT]'
	@echo 'Scripts:'
	@sed -n 's/^## //p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/  /'

## build: build the application
.PHONY: build
build: $(src) $(dist) | $(targetdir)
	$(GO) build $(GOFLAGS) -o $(abspath $(targetdir)) ./...

## run: run the application
.PHONY: run
run:
	$(GO) run $(GOFLAGS) $(abspath $(cmddir))

## test: run all tests
.PHONY: test
test:
	$(GO) test $(GOFLAGS) ./...
	$(NPM) run -ws test

## watch: run the application with live-reload
.PHONY: watch
watch:
	@$(AIR) $(AIRFLAGS)

# ================ #
# UTILITIES
# ================ #
$(dist) &: $(assets)
	$(NPM) run -w @melodeon/app build

$(targetdir):
	@mkdir -p $@


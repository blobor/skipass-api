M = $(shell printf "\033[34;1m▶\033[0m")

setup: ; $(info $(M) Fetching github.com/golang/dep...)
	npm ci

vet: setup ; $(info $(M) Running code analysis... )
	npm audit

clean: ; $(info $(M) Removing node modules... )
	rm -rf ./node_modules

build: setup

.PHONY: setup clean build

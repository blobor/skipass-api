M = $(shell printf "\033[34;1m▶\033[0m")

vet: ; $(info $(M) Running code analysis... )
	$(MAKE) -C functions/graphql vet
	$(MAKE) -C functions/playground vet

test: ; $(info $(M) Running tests... )
	$(MAKE) -C functions/graphql test

clean: ; $(info $(M) Removing generated files... )
	$(MAKE) -C functions/graphql clean
	$(MAKE) -C functions/playground clean

build: ; $(info $(M) Building All project...)
	$(MAKE) -C functions/graphql build
	$(MAKE) -C functions/playground build

deploy: clean build ; $(info $(M) Deploying project... )
	npx serverless deploy

.PHONY: vet test clean build deploy
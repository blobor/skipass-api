M = $(shell printf "\033[34;1m▶\033[0m")

setup: ; $(info $(M) Fetching github.com/golang/dep...)
	go get -u github.com/jteeuwen/go-bindata/...

clean: ; $(info $(M) Removing generated files... )
	rm -rf ./bin
	rm -rf ./schema/schemaBinData.go

schema: setup ; $(info $(M) Embedding schema files into binary...)
	go generate ./schema

vet: schema ; $(info $(M) Running code analysis... )
	env GO111MODULE=on go vet ./*.go

test: ; $(info $(M) Running tests... )
	env GO111MODULE=on go test ./*.go

build: schema vet test
	env GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o bin/graphql

.PHONY: setup schema vet test clean build deploy

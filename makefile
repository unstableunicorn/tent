OUT := tent
PKG := github.com/unstableunicorn/tent
VERSION ?= $(shell git describe --always --long --dirty)
VERSION_PATH := github.com/unstableunicorn/tent/cmd.version
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/)


all: run

buildlocal: 
	go build -v -o bin/${OUT} -ldflags="-X ${VERSION_PATH}=${VERSION}" ${PKG}

build: static
	GOOS=darwin GOARCH=amd64 go build -v -o bin/${OUT}-darwin-amd64 -ldflags="-X ${VERSION_PATH}=${VERSION}" ${PKG}
	GOOS=linux GOARCH=amd64 go build -v -o bin/${OUT}-linux-amd64 -ldflags="-X ${VERSION_PATH}=${VERSION}" ${PKG}
	GOOS=windows GOARCH=amd64 go build -v -o bin/${OUT}-windows-amd64 -ldflags="-X ${VERSION_PATH}=${VERSION}" ${PKG}

test:
	@go test -short ${PKG_LIST}

getmodules:
	go mod download

vet:
	@go vet ${PKG_LIST}

lint:
	staticcheck ./...

static: vet lint
	go build -v -o bin/${OUT}-${VERSION} -tags netgo -ldflags="-extldflags \"-static\" -w -s -X ${VERSION_PATH}=${VERSION}" ${PKG}

run: buildlocal
	./bin/${OUT}

clean:
	-@rm ./bin/${OUT}*

.PHONY: run server static vet lint
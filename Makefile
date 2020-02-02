#!/bin/make
GOROOT:=$(shell PATH="/pkg/main/dev-lang.go/bin:$$PATH" go env GOROOT)
GO=$(GOROOT)/bin/go
GOPATH:=$(shell $(GO) env GOPATH)

# disable cgo
export CGO_ENABLED=0

.PHONY: test deps

all:
	$(GOPATH)/bin/goimports -w -l .
	$(GO) build -v ./...

deps:
	$(GO) get -v -t ./...

test:
	$(GO) test -v ./...

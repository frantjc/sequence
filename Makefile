BIN ?= /usr/local/bin

GO ?= go
GO_LINUX_AMD64 ?= GOOS=linux GOARCH=amd64 $(GO)
GIT ?= git
DOCKER ?= docker
PROTOC ?= protoc
GOLANGCI-LINT ?= golangci-lint

VERSION ?= 0.0.0
PRERELEASE ?= alpha0

BRANCH ?= $(shell $(GIT) rev-parse --abbrev-ref HEAD 2>/dev/null)
COMMIT ?= $(shell $(GIT) rev-parse HEAD 2>/dev/null)
SHORT_SHA ?= $(shell $(GIT) rev-parse --short $(COMMIT))

REGISTRY ?= ghcr.io
REPOSITORY ?= frantjc/sequence
MODULE ?= github.com/$(REPOSITORY)
TAG ?= latest
IMAGE ?= $(REGISTRY)/$(REPOSITORY):$(TAG)

BUILD_ARGS ?= --build-arg version=$(VERSION) --build-arg prerelease=$(PRERELEASE)

INSTALL ?= sudo install

.DEFAULT: install

install: binaries
	@$(INSTALL) $(CURDIR)/bin/sqncd $(CURDIR)/bin/sqnc $(BIN)

bins binaries: sqnc sqncd

sqnc sqncd: shims
	@$(GO) build -ldflags "-s -w -X $(MODULE).Version=$(VERSION) -X $(MODULE).Prerelease=$(PRERELEASE)" -o $(CURDIR)/bin $(CURDIR)/cmd/$@

shims: shimuses shimsource

shimuses: uses

shimsource: source

uses source:
	@$(GO_LINUX_AMD64) build -ldflags "-s -w" -o $(CURDIR)/bin $(CURDIR)/internal/cmd/shim/$@
	@cp $(CURDIR)/bin/$@ $(CURDIR)/shim/$@

placeholders:
	@cp $(CURDIR)/shim/shim.sh $(CURDIR)/shim/shim-source
	@cp $(CURDIR)/shim/shim.sh $(CURDIR)/shim/shim-uses

image img: 
	@$(DOCKER) build -t $(IMAGE) $(BUILD_ARGS) .

generate fmt vet test:
	@$(GO) $@ ./...

tidy vendor verify:
	@$(GO) mod $@

clean: tidy placeholders
	@rm -rf bin/* vendor
	@$(DOCKER) system prune --volumes -a --filter label=sequence=true

lint:
	@$(GOLANGCI-LINT) run

tools:
	@$(GO) install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	@$(GO) install ./internal/cmd/protoc-gen-sqnc
	@echo 'Update your PATH so that the protoc compiler can find the plugins:'
	@echo '$$ export PATH=$$PATH:$(shell $(GO) env GOPATH)/bin"'

.PHONY: \
	install bins binaries sqnc sqncd \
	shims shimuses shimsource uses source placeholders \
	image img \
	format fmt vet test \
	tidy vendor verify \
	clean \
	protos \
	lint tools

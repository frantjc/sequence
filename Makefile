BIN ?= /usr/local/bin

GO ?= go
GIT ?= git
DOCKER ?= docker
PROTOC ?= protoc

BRANCH := $(shell git rev-parse --abbrev-ref HEAD 2>/dev/null)
COMMIT ?= $(shell git rev-parse HEAD 2>/dev/null | cut -c1-7)

REGISTRY ?= docker.io
REPOSITORY ?= frantjc/sequence
MODULE ?= github.com/$(REPOSITORY)
TAG ?= latest
IMAGE ?= $(REGISTRY)/$(REPOSITORY):$(TAG)

BUILD_ARGS ?= --build-arg repository=$(REPOSITORY) --build-arg tag=$(TAG) --build-arg commit=$(COMMIT)

PROTOS ?= $(shell find . -type f -name *.proto)
PROTOC_ARGS ?= --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative

INSTALL ?= sudo install

.DEFAULT: bin

bin bins binaries: sqnc sqncd sqncshim

sqnc sqncd sqncshim:
	$(GO) build -ldflags "-s -w -X github.com/frantjc/sequence/meta.Build=$(COMMIT) -X github.com/frantjc/sequence/meta.Tag=$(TAG)" -o $(CURDIR)/bin $(CURDIR)/cmd/$@
	$(INSTALL) $(CURDIR)/bin/$@ $(BIN)

image img: 
	$(DOCKER) build -t $(IMAGE) $(BUILD_ARGS) .

test: image
	$(DOCKER) build -t $(REGISTRY)/$(REPOSITORY):test $(BUILD_ARGS) --target=test .

fmt lint pretty:
	$(GO) fmt ./...

vet: fmt
	$(GO) vet ./...

tidy: vet
	$(GO) mod tidy

vendor: tidy
	$(GO) mod vendor
	$(GO) mod verify

clean: tidy
	rm -rf bin/* vendor
	$(DOCKER) system prune --volumes -a --filter label=sequence=true

protos:
	$(PROTOC) $(PROTOC_ARGS) $(PROTOS)

.PHONY: bin bins binaries sequence sqnc sqncshim image img test fmt lint pretty vet tidy vendor clean protos

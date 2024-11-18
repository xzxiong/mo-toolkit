
# SET BUILD Env
GOPROXY ?= "https://proxy.golang.com.cn,direct"
GOPATH ?= $(shell go env GOPATH)
GO_MODULE=$(shell go list -m)

# VERSION INFO
GO_VERSION=$(shell go version)
GIT_COMMIT ?= $(shell git rev-parse --short HEAD)
GIT_BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)
BUILD_TIME=$(shell date '+%F %T')
VERSION ?= "0.1.0"
GOLDFLAGS=-ldflags="-X '$(GO_MODULE)/pkg/version.GoVersion=$(GO_VERSION)' -X '$(GO_MODULE)/pkg/version.BranchName=$(GIT_BRANCH)' -X '$(GO_MODULE)/pkg/version.CommitID=$(GIT_COMMIT)' -X '$(GO_MODULE)/pkg/version.BuildTime=$(BUILD_TIME)' -X '$(GO_MODULE)/pkg/version.Version=$(VERSION)'"


###############################################################################
# Build
###############################################################################

all: build

.PHONY: build
build:
	go build $(GOLDFLAGS) -o build/

###############################################################################
# Clean
###############################################################################

.PHONY: clean
clean:
	rm -f build/*

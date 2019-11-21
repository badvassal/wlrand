PKG := github.com/badvassal/wlrand/version
VERSION:="0.0.2"
DATE := $(shell date +%F,%R)
COMMIT := $(shell git rev-parse --short HEAD)
ifneq ($(shell git status --porcelain),)
    GIT_STATE := "dirty"
endif

GOOS?=linux
GOARCH?=amd64

GWARCH=${GOOS}-${GOARCH}

all:
	@echo "Usage:"
	@echo "    make build GOOS=linux   # Linux"
	@echo "    make build GOOS=darwin  # MacOS"
	@echo "    make build GOOS=windows # Windows"

.PHONY: version

version:
	@echo ${VERSION}

build:
	@GOOS=${GOOS} GOARCH=${GOARCH} GO111MODULE=on go build -ldflags \
	    "-X ${PKG}.Version=${VERSION} -X ${PKG}.BuildDate=${DATE} -X ${PKG}.CommitHash=${COMMIT} -X ${PKG}.GitState=${GIT_STATE}"



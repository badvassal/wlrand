PKG := github.com/badvassal/wlrand/version
VERSION:="0.0.12"
DATE := $(shell date --utc +%F,%R)
COMMIT := $(shell git rev-parse --short HEAD)
ifneq ($(shell git status --porcelain),)
    COMMIT_SUFFIX := "-dirty"
endif
GIT_STATE := ${COMMIT}${COMMIT_SUFFIX}

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

.PHONY: gitstate
gitstate:
	@echo ${GIT_STATE}

build:
	@GOOS=${GOOS} GOARCH=${GOARCH} GO111MODULE=on go build -ldflags \
	    "-X ${PKG}.Version=${VERSION} -X ${PKG}.BuildDate=${DATE} -X ${PKG}.GitState=${GIT_STATE}"

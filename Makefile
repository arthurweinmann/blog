BUILDDIR                 := $(CURDIR)/build


# It's necessary to set this because some environments don't link sh -> bash.
SHELL                             = /usr/bin/env bash

GOPATH                            = $(shell go env GOPATH)
GOBIN                             = $(shell which go)
ARCH                              = $(shell uname -p)

GIT_COMMIT                        = $(shell git rev-parse HEAD)
GIT_SHA                           = $(shell git rev-parse --short HEAD)
GIT_TAG                           = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
GIT_DIRTY                         = $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")

# +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
# +++ Check bash installed

DEPTEST=$(shell command -v $(SHELL) 2> /dev/null)
ifeq ($(DEPTEST),)
$(error "We could not find bash")
endif


# +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++


# --------------------------------------------------------------------------------
# --------------------------------------------------------------------------------
# --------------------------------------------------------------------------------

.PHONY: all
all: build

.PHONY: build
build: build-web build-blog

.PHONY: build-blog
build-blog:
	rm -f $(BUILDDIR)/blog && \
	cd $(CURDIR)/cmd/blog && \
	go mod tidy && GOPROXY=direct go get -u github.com/arthurweinmann/go-https-hug@latest && \
	go mod tidy && \
	go build && mv blog $(BUILDDIR)/ && cd $(CURDIR)

.PHONY: build-web
build-web: $(BUILDDIR) check-generic-dep
	rm -rf $(BUILDDIR)/web && \
	cd $(CURDIR)/cmd/build && go mod tidy && go build && mv build $(BUILDDIR)/ && cd $(CURDIR) && \
	$(BUILDDIR)/build && cp -r $(CURDIR)/web/css $(BUILDDIR)/web && cp -r $(CURDIR)/web/assets $(BUILDDIR)/web

PHONY:check-generic-dep
check-generic-dep:
	@command -v git >/dev/null 2>&1 || { echo >&2 "git is not installed or not in path"; exit 1; }
	@command -v which >/dev/null 2>&1 || { echo >&2 "which is not installed or not in path"; exit 1; }
	@command -v make >/dev/null 2>&1 || { echo >&2 "make is not installed or not in path"; exit 1; }
	@command -v tar >/dev/null 2>&1 || { echo >&2 "tar is not installed or not in path"; exit 1; }
	@command -v go >/dev/null 2>&1 || { echo >&2 "go is not installed or not in path"; exit 1; }

$(BUILDDIR):
	@mkdir -p $(BUILDDIR)
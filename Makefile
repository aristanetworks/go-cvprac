#!/usr/bin/make
# Copyright (c) 2017 Arista Networks, Inc.  All rights reserved.
########################################################
# Makefile for go-cvprac
#
# useful targets:
#       make coverage -- go coverage tool
#       make fmtcheck -- formate check
#       make vet -- go vet
#       make lint -- go lint
#       make test -- run tests
#       make clean -- clean
#
########################################################

# Supply defaults if not provided
GOOS ?= linux
GOARCH ?= 386
GOTEST_FLAGS ?= -v -cover -timeout=240s
RACE_FLAGS ?= -race -timeout=60s
GOLDFLAGS := -ldflags="-s -w"

DEFAULT_GOPATH := $${GOPATH%%:*}
GO := go
GOFMT := gofmt
LINT := golangci-lint run
LINTFLAGS ?= --deadline=10m --exclude-use-default=false --print-issued-lines --print-linter-name --out-format=colored-line-number --disable-all --max-same-issues=0 --max-issues-per-linter=0
LINTCONFIG := --config golangci.yml
GOFILES := find . -name '*.go' ! -path './Godeps/*' ! -path './vendor/*'
GOFOLDERS := $(GO) list ./... | sed 's:^github.com/aristanetworks/go-cvprac/v3:.:' | grep -vw -e './vendor'

VERSION_FILE = version.go
GOPKGVERSION := $(shell git describe --tags --always --match "v[0-9]*" --abbrev=7 HEAD)
ifndef GOPKGVERSION
$(error unable to determine git version)
endif

# External Tools
EXTERNAL_TOOLS=\
   golang.org/x/tools/cmd/godoc \
   golang.org/x/tools/cmd/goimports

check: fmtcheck vet lint unittest

lint:
	$(GOFOLDERS) | xargs $(LINT) $(LINTFLAGS) --disable-all --enable=deadcode --tests=false
	$(GOFOLDERS) | xargs $(LINT) $(LINTCONFIG) $(LINTFLAGS)

fmtcheck:
	@if ! which $(GOFMT) >/dev/null; then echo Please install $(GOFMT); exit 1; fi
	goimports=`$(GOFILES) | xargs $(GOFMT) -d 2>&1`; \
	if test -n "$$goimports"; then \
		echo Check the following files for coding style AND USE goimports; \
		echo "$$goimports"; \
		exit 1; \
	fi

fmt:
	 $(GOFOLDERS) | xargs $(GO) fmt

vet:
	 $(GOFOLDERS) | xargs $(GO) vet

test:
	$(GO) test $(GOTEST_FLAGS)

systest:
	$(GO) test $(GOTEST_FLAGS) -tags=systest -run SystemTest ./...

unittest:
	$(GO) test $(GOTEST_FLAGS) -run UnitTest ./...

doc:
	godoc -http="localhost:6060" -play=true

COVER_PKGS := `find . -name '*_test.go' ! -path "./.git/*" ! -path "./Godeps/*" ! -path "./vendor/*" | xargs -I{} dirname {} | sort -u`
COVER_MODE := count
COVER_TMPFILE := coverage.out
coverdata:
	echo 'mode: $(COVER_MODE)' > $(COVER_TMPFILE) ;\
	for dir in $(COVER_PKGS); do \
	  $(GO) test -covermode=$(COVER_MODE) -coverprofile=$(COVER_TMPFILE).tmp $$dir || exit; \
	  tail -n +2 $(COVER_TMPFILE).tmp >> $(COVER_TMPFILE) && \
	  rm $(COVER_TMPFILE).tmp; \
	done;

coverage: coverdata
	$(GO) tool cover -html=$(COVER_TMPFILE)
	rm -f $(COVER_TMPFILE)

bootstrap:
	@for tool in  $(EXTERNAL_TOOLS) ; do \
		echo "Installing $$tool" ; \
			go get $$tool; \
	done

version: $(VERSION_FILE)

$(VERSION_FILE): $(VERSION_FILE).in .git/HEAD .git/index
	sed -e 's/@VERSION@/$(GOPKGVERSION)/' $(VERSION_FILE).in >$(VERSION_FILE)-t
	mv $(VERSION_FILE)-t $(VERSION_FILE)

clean:
	rm -rf $(COVER_TMPFILE).tmp $(COVER_TMPFILE) $(VERSION_FILE){,-t}
	$(GO) clean ./...

.PHONY: all fmtcheck test vet check doc lint
.PHONY: clean coverage coverdata version

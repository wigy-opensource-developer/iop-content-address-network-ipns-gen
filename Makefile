# Minimum version numbers for software required to build IPFS
MIN_GO_VERSION = 1.5.2

dist_root=/ipfs/QmXZQzBAFuoELw3NtjQZHkWSdA332PyQUj6pQjuhEukvg8
gx_bin=bin/gx-v0.7.0
gx-go_bin=bin/gx-go-v1.2.0

# use things in our bin before any other system binaries
export PATH := bin:$(PATH)
export IPFS_API ?= v04x.ipfs.io

all: help

go_check:
	@bin/check_go_version $(MIN_GO_VERSION)

bin/gx-v%:
	@echo "installing gx $(@:bin/gx-%=%)"
	@bin/dist_get ${dist_root} gx $@ $(@:bin/gx-%=%)
	rm -f bin/gx
	ln -s $(@:bin/%=%) bin/gx

bin/gx-go-v%:
	@echo "installing gx-go $(@:bin/gx-go-%=%)"
	@bin/dist_get ${dist_root} gx-go $@ $(@:bin/gx-go-%=%)
	rm -f bin/gx-go
	ln -s $(@:bin/%=%) bin/gx-go

gx_check: ${gx_bin} ${gx-go_bin}

path_check:
	@bin/check_go_path $(realpath $(shell pwd)) $(realpath $(GOPATH)/src/github.com/DeCentral-Budapest/ipns-gen)

deps: go_check gx_check path_check
	${gx_bin} --verbose install --global

install: deps
	go install

build: deps
	go build -o ipns-gen.a

clean:
	rm -rf ./ipns-gen.a ./ipns-gen.exe

uninstall:
	go clean github.com/DeCentral-Budapest/ipns-gen

PHONY += all help gx_check go_check deps
PHONY += install build windows_build clean uninstall

test: test_go_fmt build test_short

test_go_fmt:
	bin/test-go-fmt

test_short:
	go test -v ./...

windows_build: deps
	GOOS=windows GOARCH=amd64 go build -o ipns-gen.exe

##############################################################
# A semi-helpful help message

help:
	@echo 'DEPENDENCY TARGETS:'
	@echo ''
	@echo '  gx_check      - Installs or upgrades gx and gx-go'
	@echo '  deps          - Download dependencies using gx'
	@echo ''
	@echo 'BUILD TARGETS:'
	@echo ''
	@echo '  all           - print this help message'
	@echo '  build         - Build binary'
	@echo '  windows_build - Build Windows x64 binary'
	@echo '  install       - Build binary and install into $$GOPATH/bin'
	@echo ''
	@echo 'CLEANING TARGETS:'
	@echo ''
	@echo '  clean         - Remove binary from build directory'
	@echo '  uninstall     - Remove binary from $$GOPATH/bin'
	@echo ''
	@echo 'TESTING TARGETS:'
	@echo ''
	@echo '  test          - Run all tests'
	@echo ''

PHONY += help

.PHONY: $(PHONY)

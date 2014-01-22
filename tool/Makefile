# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(PATH)

# path of Dart and utilities
DART_DIR=$(realpath ../../dart)
DART_SDK=$(DART_DIR)/dart-sdk
DART_SDK_BIN=$(DART_SDK)/bin
DARTVM=$(DART_SDK_BIN)/dart
DART2JS=$(DART_SDK_BIN)/dart2js
DARTPUB=$(DART_SDK_BIN)/pub
DARTIUM=$(DART_DIR)/chromium/chrome


all: server

server:
	go run server.go

dartium:
	DART_FLAGS='--checked' $(DARTIUM) --user-data-dir=./data &

help:
	go help

env:
	go env


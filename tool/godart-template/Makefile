# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../../../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(PATH)

# path of Dart and utilities
DART_DIR=$(realpath ../../../dart)
DART_SDK=$(DART_DIR)/dart-sdk
DART_SDK_BIN=$(DART_SDK)/bin
DARTVM=$(DART_SDK_BIN)/dart
DART2JS=$(DART_SDK_BIN)/dart2js
DARTPUB=$(DART_SDK_BIN)/pub
DARTIUM=$(DART_DIR)/chromium/chrome


# Dart app file
APP_DART=static/app.dart
APP_JS=static/app.js

all: server

server:
	go run server.go

dartium:
	DART_FLAGS='--checked' $(DARTIUM) --user-data-dir=./data &

js: $(APP_DART)
	$(DART2JS) --minify --out=$(APP_JS) $(APP_DART)

clean:
	rm static/app.js.*
	rm static/app.precompiled.js

help:
	go help

env:
	go env


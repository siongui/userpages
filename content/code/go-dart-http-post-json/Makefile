# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../../../../../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(PATH)

# path of Dart and utilities
DART_DIR=../../../../../dart
DART_SDK=$(DART_DIR)/dart-sdk
DART_SDK_BIN=$(DART_SDK)/bin
DARTVM=$(DART_SDK_BIN)/dart
DART2JS=$(DART_SDK_BIN)/dart2js
DARTPUB=$(DART_SDK_BIN)/pub
DARTIUM=$(DART_DIR)/chromium/chrome


devserver:
	go run server.go

# Fix Dartium startup error:
# http://askubuntu.com/questions/369310/how-to-fix-missing-libudev-so-0-for-chrome-to-start-again
dartium:
	DART_FLAGS='--checked' $(DARTIUM) --user-data-dir=/tmp/data http://localhost:8000/ &

js: app.dart
	$(DART2JS) --minify --out=app.js app.dart

clean:
	[ -e app.js.deps ] && rm app.js.deps
	[ -e app.js.map  ] && rm app.js.map
	[ -e app.js      ] && rm app.js

help:
	$(DARTVM) --print-flags
	go help

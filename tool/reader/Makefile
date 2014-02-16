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

all: devserver

devserver:
	go run server.go

production:
	go run server.go -production=true

dartium:
	DART_FLAGS='--checked' $(DARTIUM) --user-data-dir=./data &

js: $(APP_DART)
	$(DART2JS) --minify --out=$(APP_JS) $(APP_DART)

test:
	go test -test.v ./mylib

polling_test:
	go test -test.v mylib/polling_test.go mylib/polling.go \
		mylib/opml.go mylib/seed.go mylib/db.go

db_test:
	go test -test.v mylib/db_test.go mylib/db.go mylib/opml.go \
		mylib/seed.go

seed_test:
	go test -test.v mylib/seed_test.go mylib/seed.go

tmp_test:
	go test -test.v tmp.go tmp_test.go

opml_test:
	go test -test.v mylib/opml_test.go mylib/opml.go

help:
	go help

env:
	go env

# An autocompletion daemon for the Go programming language
# https://github.com/nsf/gocode
gocode:
	go get -u github.com/nsf/gocode

# sqlite3 driver for go that using database/sql
# https://github.com/mattn/go-sqlite3
sqlite3:
	go get -u github.com/mattn/go-sqlite3


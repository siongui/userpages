DART_DIR=../../dart
DART_SDK=$(DART_DIR)/dart-sdk
DART_SDK_BIN=$(DART_SDK)/bin
DARTIUM=$(DART_DIR)/chromium/chrome

PATH=$PATH:$(DART_SDK_BIN)

client: test.html test.dart
	DART_FLAGS='--checked' $(DARTIUM) --user-data-dir=./data test.html

standalone: test.dart
	dart --checked test.dart

help:
	dart --print-flags

all: client

.PHONY: client standalone help

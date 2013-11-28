DART_DIR=../../../dart
DART_SDK=$(DART_DIR)/dart-sdk
DART_SDK_BIN=$(DART_SDK)/bin
DARTIUM=$(DART_DIR)/chromium/chrome

PATH=$PATH:$(DART_SDK_BIN)

START_HTML=toggle.html
DART_APP=toggle.dart

# Fix Dartium startup error:
# http://askubuntu.com/questions/369310/how-to-fix-missing-libudev-so-0-for-chrome-to-start-again
web: $(START_HTML) $(DART_APP)
	DART_FLAGS='--checked' $(DARTIUM) --user-data-dir=./data $(START_HTML)

cmd: $(DART_APP)
	dart --checked $(DART_APP)

help:
	dart --print-flags

all: web

.PHONY: web cmd help

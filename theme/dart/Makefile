DART_DIR=../../../dart
DART_SDK=$(DART_DIR)/dart-sdk
DART_SDK_BIN=$(DART_SDK)/bin
DARTVM=$(DART_SDK_BIN)/dart
DART2JS=$(DART_SDK_BIN)/dart2js
DARTPUB=$(DART_SDK_BIN)/pub
DARTIUM=$(DART_DIR)/chromium/chrome

START_HTML=example/toggle.html
DART_APP=example/toggle.dart
DART_APP_JS=example/toggle.js
DART_SERVER=example/server.dart

all: cmd

# Fix Dartium startup error:
# http://askubuntu.com/questions/369310/how-to-fix-missing-libudev-so-0-for-chrome-to-start-again
web: $(START_HTML) $(DART_APP)
	DART_FLAGS='--checked' $(DARTIUM) --user-data-dir=./data $(START_HTML)

cmd: $(DART_APP)
	$(DARTVM) --checked $(DART_SERVER)

install: pubspec.yaml
	$(DARTPUB) get

upgrade: pubspec.yaml
	$(DARTPUB) upgrade

drypub: pubspec.yaml
	$(DARTPUB) publish --dry-run

js: $(DART_APP)
	$(DART2JS) --minify --out=$(DART_APP_JS) $(DART_APP)

help:
	$(DARTVM) --print-flags

# http://stackoverflow.com/questions/2989465/rm-rf-versus-rm-rf
clean:
	-rm *.js
	-rm $(DART_APP_JS).deps
	-rm $(DART_APP_JS).map

.PHONY: web cmd install upgrade js clean help

DART_DIR=../../dart
DART_SDK=$(DART_DIR)/dart-sdk
DART_SDK_BIN=$(DART_SDK)/bin
DARTVM=$(DART_SDK_BIN)/dart
DART2JS=$(DART_SDK_BIN)/dart2js
DARTIUM=$(DART_DIR)/chromium/chrome

START_HTML=toggle.html
DART_APP=toggle.dart
DART_APP_JS=toggle.js
DART_APP_CSS=toggle.css

# Fix Dartium startup error:
# http://askubuntu.com/questions/369310/how-to-fix-missing-libudev-so-0-for-chrome-to-start-again
web: $(START_HTML) $(DART_APP) $(DART_APP_CSS)
	DART_FLAGS='--checked' $(DARTIUM) --user-data-dir=./data $(START_HTML)

js: $(DART_APP)
	$(DART2JS) --minify --out=$(DART_APP_JS) $(DART_APP)

help:
	$(DARTVM) --print-flags

# http://stackoverflow.com/questions/2989465/rm-rf-versus-rm-rf
clean:
	-rm *.js
	-rm $(DART_APP_JS).deps
	-rm $(DART_APP_JS).map

all: web

.PHONY: web js clean help

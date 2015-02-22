# path of Dart and utilities
DART_DIR=../../../../../dart
DART_SDK=$(DART_DIR)/dart-sdk
DART_SDK_BIN=$(DART_SDK)/bin
DARTVM=$(DART_SDK_BIN)/dart
DART2JS=$(DART_SDK_BIN)/dart2js
DARTPUB=$(DART_SDK_BIN)/pub
DARTIUM=$(DART_DIR)/chromium/chrome

HTML=pali.html
DART_APP=pali.dart
DART_APP_JS=pali.js

all: test

# Fix Dartium startup error:
# http://askubuntu.com/questions/369310/how-to-fix-missing-libudev-so-0-for-chrome-to-start-again
test:
	DART_FLAGS='--checked' $(DARTIUM) --user-data-dir=/tmp/data $(HTML) &

demo: js
	chromium-browser $(HTML)

js: $(DART_APP)
	$(DART2JS) --minify --out=$(DART_APP_JS) $(DART_APP)

# http://stackoverflow.com/questions/2989465/rm-rf-versus-rm-rf
clean:
	-rm *.js
	-rm *.deps
	-rm *.map

help:
	$(DARTVM) --print-flags


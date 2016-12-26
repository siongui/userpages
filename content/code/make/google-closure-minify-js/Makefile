SOURCE_JS_FILES=$(filter-out %.min.js, $(wildcard *.js))
CONCAT_JS=/tmp/tmp.js
MINIFIED_JS=app.min.js

minjs: concat
	curl \
	-d compilation_level=SIMPLE_OPTIMIZATIONS \
	-d language=ECMASCRIPT5 \
	-d output_format=text \
	-d output_info=compiled_code \
	--data-urlencode "js_code@$(CONCAT_JS)" \
	https://closure-compiler.appspot.com/compile \
	> $(MINIFIED_JS)

concat:
	@cat $(SOURCE_JS_FILES) > $(CONCAT_JS)

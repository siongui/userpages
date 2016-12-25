#!/bin/bash

TMPJS=/tmp/tmp.js

# $1 is the directory in which js files to be processed
# $2 is the minified js from online Google Closure Compiler
cat $(find $1 -type f -name "*.js") > ${TMPJS}
curl \
  -d compilation_level=SIMPLE_OPTIMIZATIONS \
  -d language=ECMASCRIPT5 \
  -d output_format=text \
  -d output_info=compiled_code \
  --data-urlencode "js_code@${TMPJS}" \
  https://closure-compiler.appspot.com/compile \
  > $2
rm ${TMPJS}

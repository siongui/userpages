#!/bin/sh
# Convert all files in directory from Big5 to UTF-8

# list all files recursively
for file in $(find dhpstory/ -type f)
do
  output=${file}-utf8.html
  echo "\033[92mConverting ${file} (BIG5) to ${output} (UTF8) ...\033[0m"
  iconv -f BIG5 -t UTF8 ${file} > ${output}
done

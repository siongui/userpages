#!/bin/bash

# $1 is the directory in which files to be processed
IFS=$'\n'
for path in $(find $1 -type f)
do
  filename=$(basename "$path")
  filenameWithoutExt=${filename%.*}
  if [[ ${filenameWithoutExt} =~ \ \([0-9]+\)$ ]]; then
    echo ${path}
    #rm ${path}
  fi
done

#!/bin/bash

langPrefix=bash-
targetDir=../../
rstdir=../../../articles/

for srcdir in $(ls -d ${targetDir}*/)
do
  dirprefix=$(dirname ${srcdir})
  # remove extension
  rmext=${srcdir%%/}
  # remove prefix
  rawdir=${rmext##*/}
  if [[ ${rawdir} == ${langPrefix}* ]]; then
    newdir=${rawdir##${langPrefix}}
    dstdir=${dirprefix}/${langPrefix%%-}/${newdir}
    mv ${srcdir} ${dstdir}
    find ${rstdir} -type f -name "*.rst" | xargs sed -i "s/${rawdir}/${rawdir/-/\\\/}/g"
  fi
done

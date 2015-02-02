#!/bin/bash
# list-files-recursively-rename.sh
# List All Files in Directory Recursively and Rename

count=0 # count the number of processed files

# list all files recursively
for file in $(find content/articles -type f)
do
  # rename the file by replacing # with %
  new=`echo $file | sed 's/#/%/g'`
  if [ $file != $new ]; then
    count=$((count+=1))
    echo "$count . $file => $new"
    # files are version controlled by Git, so use git mv
    git mv $file $new
  fi
done

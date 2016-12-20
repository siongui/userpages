#!/bin/bash

# $1 is the directory in which files to be copied
# $2 is the destination directory

tar cvf /tmp/foo.tar $1
tar xvf /tmp/foo.tar -C $2
rm /tmp/foo.tar

#!/usr/bin/env python
# -*- coding:utf-8 -*-

import os
import re
import filecmp

redundant = re.compile('.+ \(\d{1}\)')

def processDir(rootDir):
  # http://www.tutorialspoint.com/python/os_walk.htm
  for root, dirs, files in os.walk(rootDir):
    for filename in files:
      # http://stackoverflow.com/questions/678236/how-to-get-the-filename-without-the-extension-from-a-path-in-python
      result = redundant.findall(os.path.splitext(filename)[0])
      if len(result) == 1:
        ext = os.path.splitext(filename)[1]
        orig = result[0][:-4] + ext
        orig_path = os.path.join(root, orig)
        # http://stackoverflow.com/questions/82831/how-to-check-whether-a-file-exists-using-python
        if os.path.isfile(orig_path):
          path = os.path.join(root, filename)
          # http://stackoverflow.com/questions/1072569/see-if-two-files-have-the-same-content-in-python
          if filecmp.cmp(orig_path, path):
            # this is a redundant file
            print("redundant: " + path)


if __name__ == '__main__':
  dstDir = "YOUR_DIR_PATH"
  # http://stackoverflow.com/questions/50499/how-do-i-get-the-path-and-name-of-the-file-that-is-currently-executing
  processDir(os.path.join(os.path.dirname(__file__), dstDir))

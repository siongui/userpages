#!/usr/bin/env python
# -*- coding:utf-8 -*-

import os
import json

if __name__ == '__main__':
  for path in os.listdir('.'):
    if path.startswith("http"):
      with open(path, 'r') as f:
        j = json.loads(f.read())
	print(j)

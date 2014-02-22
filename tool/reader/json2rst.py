#!/usr/bin/env python
# -*- coding:utf-8 -*-

import os
import json

if __name__ == '__main__':
  langTagLink = {}

  for path in os.listdir('.'):
    if path.startswith("http"):
      with open(path, 'r') as f:
        j = json.loads(f.read())
	#print(j)

	if j["Language"] not in langTagLink:
          langTagLink[j["Language"]] = {}

	if j['Tag'] not in langTagLink[j["Language"]]:
          langTagLink[j["Language"]][j['Tag']] = []

	langTagLink[j["Language"]][j['Tag']].append(j)

  print(langTagLink)

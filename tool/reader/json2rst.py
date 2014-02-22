#!/usr/bin/env python
# -*- coding:utf-8 -*-

import os
import json

if __name__ == '__main__':
  langTagLink = {}

  os.chdir("./links/")
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
  with open("links.rst", 'w') as f:
    for lang in langTagLink:
      space = 0
      if lang == u"zh_CN":
        f.write(u"Simplified Chinese:\n\n")
	space += 2

      tags = u""
      for tag in langTagLink[lang]:
	tags = tags + tag + u", "
	f.write(" " * space)
	f.write(tag)
	f.write(":\n\n")

	for j in langTagLink[lang][tag]:
	  space2 = space + 2
          f.write(" " * space2)
	  f.write("`")
	  f.write(j[u"Title"].encode("utf8"))
	  f.write(" <")
	  f.write(j[u"Link"])
	  f.write(">`_\n")
	  if j[u"HNComments"].startswith(u"https://news.ycombinator.com/item?id="):
            f.write(" " * space2)
            f.write("(`HN discuss <")
	    f.write(j[u"HNComments"])
	    f.write(">`__)\n\n")
	  else:
	    f.write("\n")

      f.write(tags[:-2])
      f.write("\n\n")

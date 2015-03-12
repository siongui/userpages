#!/usr/bin/env python
# -*- coding:utf-8 -*-

from jinja2 import Environment
import sys

tmpl = """
<span>hello {{name | gettext}}</span>
"""

def gettext(s):
  if s == "world":
    return u"世界"
  return s

if __name__ == '__main__':
  env = Environment()
  env.filters['gettext'] = gettext
  t = env.from_string(tmpl)

  sys.stdout.write(t.render(name="world"))

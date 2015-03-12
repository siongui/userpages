#!/usr/bin/env python
# -*- coding:utf-8 -*-

from jinja2 import Template
import sys

tmpl = """
<span>hello {{gettext(name)}}</span>
"""

def gettext(s):
  if s == "world":
    return u"世界"
  return s

if __name__ == '__main__':
  t = Template(tmpl)
  t.globals['gettext'] = gettext

  sys.stdout.write(t.render(name="world"))

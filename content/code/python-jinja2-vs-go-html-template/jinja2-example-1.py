#!/usr/bin/env python
# -*- coding:utf-8 -*-

from jinja2 import Template
import sys

if __name__ == '__main__':
  tmpl = Template("Hello {{ name }}!")
  sys.stdout.write(tmpl.render(name="World"))

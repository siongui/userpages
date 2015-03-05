#!/usr/bin/env python
# -*- coding:utf-8 -*-

from jinja2 import Template
import sys

tmpl = """
{% for link in links %}
<a href="{{link.href}}">{{link.name}}</a>
{% endfor %}
"""

if __name__ == '__main__':
  links = [
    {'name': 'Google', 'href': 'https://www.google.com'},
    {'name': 'Facebook', 'href': 'https://www.facebook.com'}
  ]
  t = Template(tmpl)
  sys.stdout.write(t.render(links))

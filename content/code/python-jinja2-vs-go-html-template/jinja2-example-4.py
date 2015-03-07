#!/usr/bin/env python
# -*- coding:utf-8 -*-

from jinja2 import Template
import sys

tmpl = """
{% for name, href in links.iteritems() %}
<a href="{{ href }}">{{ name }}</a>
{% endfor %}

{% for name, href in links.iteritems() %}
{{ loop.index }}: <a href="{{ href }}">{{ name }}</a>
{% endfor %}

{% for name, href in links.iteritems() %}
{{ loop.index0 }}: <a href="{{ href }}">{{ name }}</a>
{% endfor %}
"""

if __name__ == '__main__':
  links = {
    'Google': 'https://www.google.com',
    'Facebook': 'https://www.facebook.com',
  }
  t = Template(tmpl)
  sys.stdout.write(t.render(links=links))

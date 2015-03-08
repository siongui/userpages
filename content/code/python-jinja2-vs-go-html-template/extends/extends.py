#!/usr/bin/env python
# -*- coding:utf-8 -*-

import jinja2
import os


JINJA_ENVIRONMENT = jinja2.Environment(
    loader=jinja2.FileSystemLoader(os.path.dirname(__file__)),
    extensions=['jinja2.ext.autoescape'],
    autoescape=True)


if __name__ == '__main__':
  template_values = {
    'name': 'world',
  }
  template = JINJA_ENVIRONMENT.get_template('index.html')
  print(template.render(template_values))

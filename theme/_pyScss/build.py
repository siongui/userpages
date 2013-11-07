#!/usr/bin/env python
# -*- coding: utf-8 -*- #

from os.path import join, dirname
import scss

scss.config.PROJECT_ROOT = dirname(__file__)

_scss = scss.Scss(
  #scss_opts={
  #  'compress': True,
  #  'debug_info': True,
  #}
)

compiled_css_from_file = _scss.compile(
  scss_file=join(dirname(__file__), 'style.scss')
)
print(compiled_css_from_file)

output_path = join(dirname(__file__), '../static/css/style.css')
with open(output_path, 'w') as f:
  f.write(compiled_css_from_file)

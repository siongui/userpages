#!/usr/bin/env python
# -*- coding: utf-8 -*- #

from os.path import join, dirname, exists
import scss

scss.config.PROJECT_ROOT = dirname(__file__)
scss.config.LOAD_PATHS = dirname(__file__)

_scss = scss.Scss(
  #scss_opts={
  #  'compress': True,
  #  'debug_info': True,
  #}
)

compiled_css_from_file = _scss.compile(
  scss_file=join(dirname(__file__), 'grid.scss')
)
print(compiled_css_from_file)

output_path = join(dirname(__file__), 'all.css')
with open(output_path, 'w') as f:
  f.write(compiled_css_from_file)

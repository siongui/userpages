#!/usr/bin/env python
# -*- coding: utf-8 -*- #

from os.path import join, dirname
import scss

scss.config.PROJECT_ROOT = join(dirname(__file__))

_scss = scss.Scss(
  #scss_opts={
  #  'compress': True,
  #  'debug_info': True,
  #}
)

compiled_css_from_file = _scss.compile(scss_file='style.scss')
print(compiled_css_from_file)

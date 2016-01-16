#!/usr/bin/env python
# -*- coding:utf-8 -*-

import os
import jinja2
import i18n

env = jinja2.Environment(
    loader=jinja2.FileSystemLoader(
      [os.path.join(os.path.dirname(__file__), 'view')]),
    extensions=['jinja2.ext.i18n'])

env.install_gettext_translations(i18n)


if __name__ == '__main__':
  template = env.get_template('demo.html')
  i18n.setLocale("zh_TW")
  print(template.render())

  print("\n-----\n")

  i18n.setLocale("vi_VN")
  print(template.render())

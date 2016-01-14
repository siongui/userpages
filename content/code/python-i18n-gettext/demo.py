#!/usr/bin/env python
# -*- coding:utf-8 -*-

import i18n


if __name__ == '__main__':
  i18n.setLocale("zh_TW")
  print(i18n.gettext("Home"))
  print(i18n.gettext("Canon"))
  print(i18n.gettext("About"))
  print(i18n.gettext("Setting"))
  print(i18n.gettext("Translation"))
  i18n.setLocale("vi_VN")
  print(i18n.gettext("Home"))
  print(i18n.gettext("Canon"))
  print(i18n.gettext("About"))
  print(i18n.gettext("Setting"))
  print(i18n.gettext("Translation"))

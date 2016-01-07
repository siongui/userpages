#!/usr/bin/env python
# -*- coding:utf-8 -*-

import re
import pyopencc
tw2cn = pyopencc.OpenCC('zht2zhs.ini').convert


if __name__ == '__main__':
  with open("locale/zh_TW/LC_MESSAGES/messages.po", 'r') as ftw:
    with open("locale/zh_CN/LC_MESSAGES/messages.po", "w") as fcn:
      for line in ftw.readlines():
        if 'zh_TW' in line:
          fcn.write(line.replace('zh_TW', 'zh_CN'))
        elif line.startswith('msgstr'):
          try:
            fcn.write(re.sub('msgstr "(.+)"', lambda m: 'msgstr "%s"' % tw2cn(m.group(1)), line))
          except UnicodeEncodeError:
            fcn.write(re.sub('msgstr "(.+)"', lambda m: 'msgstr "%s"' % tw2cn(m.group(1)), line).encode('utf-8'))
        else:
          fcn.write(line)

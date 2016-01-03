#!/usr/bin/env python
# -*- coding:utf-8 -*-

import pyopencc
CN2TW = pyopencc.OpenCC('zhs2zhtw_vp.ini').convert

if __name__ == '__main__':
  print(CN2TW("中国鼠标软件打印机"))

#!/usr/bin/env python
# -*- coding:utf-8 -*-

import xml.dom.minidom

def decodeItem(item):
  dict = item.getElementsByTagName("dict")[0]
  word = item.getElementsByTagName("word")[0]
  explain = item.getElementsByTagName("explain")[0]

  dictstr = dict.childNodes[0].data
  wordstr = word.childNodes[0].data
  explainstr = explain.childNodes[0].data

  print("dict: %s" % dictstr)
  print("word: %s" % wordstr)
  print("explain: %s" % explainstr)


def main():
  dom = xml.dom.minidom.parse("example.xml")

  items = dom.getElementsByTagName("item")
  for item in items:
    decodeItem(item)


if __name__ == '__main__':
  main()

#!/usr/bin/env python
# -*- coding:utf-8 -*-

import xml.dom.minidom

def main():
  impl = xml.dom.minidom.getDOMImplementation()
  dom = impl.createDocument(None, u'html', None)

  root = dom.documentElement
  demoNode = dom.createElement(u'demoTag')
  root.appendChild(demoNode)

  print(dom.toxml())

if __name__ == '__main__':
  main()

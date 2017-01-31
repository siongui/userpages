#!/usr/bin/env python
# -*- coding:utf-8 -*-

import xml.dom.minidom

def main():
  impl = xml.dom.minidom.getDOMImplementation()
  dom = impl.createDocument(None, u'html', None)

  demoNode = dom.createElement(u'demoTag')
  demoNode.setAttribute(u'integer', u'1')
  demoTextNode = dom.createTextNode(u'Hello World!')
  demoNode.appendChild(demoTextNode)

  root = dom.documentElement
  root.appendChild(demoNode)

  print(dom.toxml())

if __name__ == '__main__':
  main()

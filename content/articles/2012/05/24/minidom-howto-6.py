#!/usr/bin/env python
# -*- coding:utf-8 -*-

import xml.dom.minidom
import codecs

def main():
  impl = xml.dom.minidom.getDOMImplementation()
  dom = impl.createDocument(None, u'html', None)

  demoNode = dom.createElement(u'demoTag')
  demoNode.setAttribute(u'integer', unicode(1))
  demoTextNode = dom.createTextNode(unicode(u'Hello World!'))
  demoNode.appendChild(demoTextNode)

  root = dom.documentElement
  root.appendChild(demoNode)

  f = file('example.xml', 'wb')
  writer = codecs.lookup('utf-8')[3](f)
  dom.writexml(writer, encoding = 'utf-8')
  f.close()

if __name__ == '__main__':
  main()

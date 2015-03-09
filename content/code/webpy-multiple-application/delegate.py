#!/usr/bin/env python
# -*- coding:utf-8 -*-

# Simulate app.yaml of Google App Engine

import web
import staticweb
import mainweb

mapping = (
  r'/robots.txt', staticweb.app,
  r'/', mainweb.app,
)

class customApp(web.application):
  def _delegate_sub_application(self, dir, app):
    return app.handle_with_processors()

if __name__ == '__main__':
  app = customApp(mapping)
  app.run()

#!/usr/bin/env python
# -*- coding:utf-8 -*-

import web

urls = (
  r"/", "MainPage"
)

class MainPage:
  def GET(self):
    return "Hello World"


app = web.application(urls, globals())
try:
  from google.appengine.ext import ndb
  # runs on Google App Engine
  app = app.gaerun()
except ImportError:
  application = app.wsgifunc()

if __name__ == '__main__':
  app.run()

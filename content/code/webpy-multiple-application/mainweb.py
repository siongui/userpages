#!/usr/bin/env python
# -*- coding:utf-8 -*-

import web

urls = (
  r"/about", "AboutPage",
  r"/", "MainPage",
)

class AboutPage:
  def GET(self):
    return "Hello About"

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

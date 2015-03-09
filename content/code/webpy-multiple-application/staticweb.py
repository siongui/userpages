#!/usr/bin/env python
# -*- coding:utf-8 -*-

import web

urls = (
  r"/robots.txt", "StaticPage",
)


class StaticPage:
  def GET(self):
    return 'User-agent: *\nDisallow: /'


app = web.application(urls, globals())
try:
  from google.appengine.ext import ndb
  # runs on Google App Engine
  app = app.gaerun()
except ImportError:
  application = app.wsgifunc()

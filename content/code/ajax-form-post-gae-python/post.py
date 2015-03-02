#!/usr/bin/env python
# -*- coding:utf-8 -*-

import webapp2
import urllib2


class PostPage(webapp2.RequestHandler):
  def post(self):
    input1 = urllib2.unquote(self.request.get('input1'))
    input2 = urllib2.unquote(self.request.get('input2'))

    self.response.out.write("input1: %s\ninput2: %s" % (input1, input2))


application = webapp2.WSGIApplication([
    ('/post', PostPage),
], debug=True)

#!/usr/bin/env python
# -*- coding:utf-8 -*-

import jinja2
import os
import webapp2


JINJA_ENVIRONMENT = jinja2.Environment(
    loader=jinja2.FileSystemLoader(os.path.dirname(__file__)),
    extensions=['jinja2.ext.autoescape'],
    autoescape=True)


class MainPage(webapp2.RequestHandler):
  def get(self):
    template_values = {
      'name': 'World',
    }
    template = JINJA_ENVIRONMENT.get_template('index-jinja2.html')
    self.response.write(template.render(template_values))


application = webapp2.WSGIApplication([
    ('/', MainPage),
], debug=True)

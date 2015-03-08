#!/usr/bin/env python
# -*- coding:utf-8 -*-

import web
import jinja2
import os


JINJA_ENVIRONMENT = jinja2.Environment(
    loader=jinja2.FileSystemLoader(os.path.dirname(__file__)),
    extensions=['jinja2.ext.autoescape'],
    autoescape=True)

urls = (
  r"/", "MainPage"
)

class MainPage:
  def GET(self):
    template_values = {
      'name': 'World',
    }
    template = JINJA_ENVIRONMENT.get_template('index-jinja2.html')
    return template.render(template_values)


if __name__ == '__main__':
  app = web.application(urls, globals())
  app.run()

#!/usr/bin/env python
# -*- coding:utf-8 -*-

# for webapp2 to import babel and pytz
import sys
import os
sys.path.insert(0, os.path.join(os.path.dirname(__file__), 'babel.zip'))
sys.path.insert(0, os.path.join(os.path.dirname(__file__), 'pytz.zip'))

import webapp2
import jinja2
from webapp2_extras import i18n


JINJA_ENVIRONMENT = jinja2.Environment(
    loader=jinja2.FileSystemLoader(os.path.dirname(__file__)),
    extensions=['jinja2.ext.i18n', 'jinja2.ext.autoescape'],
    autoescape=True)

JINJA_ENVIRONMENT.install_gettext_translations(i18n)

class MainPage(webapp2.RequestHandler):
    def get(self):
        locale = self.request.GET.get('locale', 'en_US')
        i18n.get_i18n().set_locale(locale)

        template_values = {}
        template = JINJA_ENVIRONMENT.get_template('index.html')
        self.response.write(template.render(template_values))

app = webapp2.WSGIApplication([
    ('/', MainPage),
], debug=True)

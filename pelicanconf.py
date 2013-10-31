#!/usr/bin/env python
# -*- coding: utf-8 -*- #
from __future__ import unicode_literals

# use LC_ALL="en_US.UTF-8" with make command to build English datetime
AUTHOR = u'Siong-Ui Te'
SITENAME = u'Theory and Practice'
SITEURL = ''

TIMEZONE = 'Asia/Taipei'

DEFAULT_LANG = u'en'

# Take advantage of the following defaults
# STATIC_SAVE_AS = '{path}'
# STATIC_URL = '{path}'
STATIC_PATHS = ['extra/robots.txt',
                'extra/yezi.png',]
EXTRA_PATH_METADATA = {'extra/robots.txt': {'path': 'robots.txt'},
                       'extra/yezi.png': {'path': 'favicon.ico'},}

# Feed generation is usually not desired when developing
FEED_ALL_ATOM = None
CATEGORY_FEED_ATOM = None
TRANSLATION_FEED_ATOM = None

# Blogroll
LINKS =  (('Pelican', 'http://getpelican.com/'),
          ('Python.org', 'http://python.org/'),
          ('Jinja2', 'http://jinja.pocoo.org/'),
          ('You can modify those links in your config file', '#'),)

# Social widget
SOCIAL = (('github', 'https://github.com/siongui'),
          ('facebook', 'https://www.facebook.com/siongui.te'),
          ('google-plus', 'https://plus.google.com/u/0/+SiongUiTe/posts'),)

DEFAULT_PAGINATION = 10

# Uncomment following line if you want document-relative URLs when developing
#RELATIVE_URLS = True

# Specify a customized theme, via path relative to the settings file
THEME = 'theme'

# Bootstrap3, Font Awesome, and jQuery URL
# Bootstrap CDN: http://www.bootstrapcdn.com/
BOOTSTRAP3_CSS_URL = '//netdna.bootstrapcdn.com/bootstrap/3.0.0/css/bootstrap.min.css'
FONT_AWESOME_URL = '//netdna.bootstrapcdn.com/font-awesome/4.0.0/css/font-awesome.min.css'
BOOTSTRAP3_JS_URL = '//netdna.bootstrapcdn.com/bootstrap/3.0.0/js/bootstrap.min.js'
JQUERY_URL = '//ajax.googleapis.com/ajax/libs/jquery/2.0.3/jquery.min.js'

# HTML meta info
META_KEYWORDS = 'blog'
META_DESCRIPTION = 'Theory and Practice'

# Facebook Social Plugins › Comments Plugin
FB_APPID = '498013980272599'
FB_COMMENTS = True

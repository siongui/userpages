#!/usr/bin/env python
# -*- coding: utf-8 -*- #
from __future__ import unicode_literals

# use LC_ALL="en_US.UTF-8" with make command to build English datetime
AUTHOR = u'Siong-Ui Te'
SITENAME = u'Theory and Practice'
SITEURL = ''

TIMEZONE = 'Asia/Taipei'

DEFAULT_LANG = u'en'
LOCALE = 'en_US.UTF-8'

# @see http://docs.getpelican.com/en/3.3.0/settings.html#basic-settings
# @see http://docs.getpelican.com/en/3.3.0/settings.html#path-metadata
PATH_METADATA = '(?P<date>\d{4}/\d{2}/\d{2})/(?P<slug>[-a-zA-Z0-9]*)#(?P<lang>[_a-zA-Z]{2,5})\.rst'

# @see http://docs.getpelican.com/en/3.3.0/settings.html#url-settings
ARTICLE_URL = '{date:%Y}/{date:%m}/{date:%d}/{slug}/'
ARTICLE_SAVE_AS = '{date:%Y}/{date:%m}/{date:%d}/{slug}/index.html'
ARTICLE_LANG_URL = '{lang}/{date:%Y}/{date:%m}/{date:%d}/{slug}/'
ARTICLE_LANG_SAVE_AS = '{lang}/{date:%Y}/{date:%m}/{date:%d}/{slug}/index.html'

# Take advantage of the following defaults
# STATIC_SAVE_AS = '{path}'
# STATIC_URL = '{path}'
STATIC_PATHS = ['extra/robots.txt',
                'extra/Feeder.opml',
                'extra/yezi.png',]
EXTRA_PATH_METADATA = {'extra/robots.txt': {'path': 'robots.txt'},
                       'extra/Feeder.opml': {'path': 'misc/Feeder.opml'},
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
DISPLAY_CATEGORIES_ON_MENU = False

# Uncomment following line if you want document-relative URLs when developing
#RELATIVE_URLS = True

# Specify a customized theme, via path relative to the settings file
THEME = 'theme'

# For Github Ribbon
GITHUB_REPO_URL = 'https://github.com/siongui/userpages'

# Bootstrap3, Font Awesome, and jQuery URL
# Bootstrap CDN: http://www.bootstrapcdn.com/
FONT_AWESOME_URL = '//netdna.bootstrapcdn.com/font-awesome/4.0.0/css/font-awesome.min.css'

# HTML meta info
META_KEYWORDS = 'blog'
META_DESCRIPTION = 'Theory and Practice'

# Facebook Social Plugins › Comments Plugin
#FB_APPID = ''
#FB_COMMENTS = True

PLUGIN_PATH = "plugins"
PLUGINS = ["gist", "code"]


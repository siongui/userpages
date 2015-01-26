#!/usr/bin/env python
# -*- coding: utf-8 -*- #
from __future__ import unicode_literals

AUTHOR = u'Siong-Ui Te'
SITENAME = u'Theory and Practice'
SITEURL = ''

PATH = 'content'

TIMEZONE = 'Asia/Taipei'

DEFAULT_LANG = u'en'

# Feed generation is usually not desired when developing
FEED_ALL_ATOM = None
CATEGORY_FEED_ATOM = None
TRANSLATION_FEED_ATOM = None
AUTHOR_FEED_ATOM = None
AUTHOR_FEED_RSS = None

DEFAULT_PAGINATION = 10

THEME = 'theme'

PLUGIN_PATHS = ['plugins']
PLUGINS = ['i18n_subsites', 'edit_on_github', 'embed_github_repository_file']

# mapping: language_code -> settings_overrides_dict
I18N_SUBSITES = {
  'zh': {
    'SITENAME': '理論與實作',
    'AUTHOR': '戴上為',
  }
}

DIRECT_TEMPLATES = ['index']
PAGE_ORDER_BY = 'order'

# CONTENT_DIR_URL is the setting for edit_on_github plugin
CONTENT_DIR_URL = u'https://github.com/siongui/userpages/tree/master/content'

# Uncomment following line if you want document-relative URLs when developing
#RELATIVE_URLS = True

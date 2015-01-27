#!/usr/bin/env python
# -*- coding: utf-8 -*- #
from __future__ import unicode_literals

AUTHOR = u'Siong-Ui Te'
SITENAME = u'Theory and Practice'
SITEURL = ''

PATH = 'content'

TIMEZONE = 'Asia/Taipei'

DEFAULT_LANG = u'en'

# @see http://docs.getpelican.com/en/3.5.0/settings.html#basic-settings
# @see http://docs.getpelican.com/en/3.5.0/settings.html#path-metadata
PATH_METADATA = '(?P<date>\d{4}/\d{2}/\d{2})/(?P<slug>[-a-zA-Z0-9]*)\.rst'

# @see http://docs.getpelican.com/en/3.5.0/settings.html#url-settings
ARTICLE_URL = '{date:%Y}/{date:%m}/{date:%d}/{slug}/'
ARTICLE_SAVE_AS = '{date:%Y}/{date:%m}/{date:%d}/{slug}/index.html'

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

# generate only index.html and pages. (no archives, tags, categories)
DIRECT_TEMPLATES = ['index']
# use metadata attribute 'order' in page files for ordering
# @see http://docs.getpelican.com/en/3.5.0/settings.html#url-settings
PAGE_ORDER_BY = 'order'

# CONTENT_DIR_URL is the setting for edit_on_github plugin
CONTENT_DIR_URL = u'https://github.com/siongui/userpages/tree/master/content'

# my custom setting for HTML meta info
META_KEYWORDS = 'Python, SCSS, Web, blog'
META_DESCRIPTION = 'Theory and Practice - Web Development'

# Uncomment following line if you want document-relative URLs when developing
#RELATIVE_URLS = True

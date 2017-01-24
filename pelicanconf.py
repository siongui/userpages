#!/usr/bin/env python
# -*- coding: utf-8 -*- #
from __future__ import unicode_literals

AUTHOR = u'Siong-Ui Te'
SITENAME = u'Theory and Practice'
SITEURL = ''

PATH = 'content'

# avoid processing .html files
READERS = {'html': None}

# mix articles and static files in the same place
# @see https://github.com/getpelican/pelican/issues/1587
ARTICLE_PATHS = ['articles']
STATIC_PATHS = ['articles', 'extra', 'code', 'chanting']
EXTRA_PATH_METADATA = {'extra/robots.txt': {'path': 'robots.txt'},
                       'extra/yezi.png': {'path': 'favicon.ico'},}

TIMEZONE = 'Asia/Taipei'

DEFAULT_LANG = u'en'
LOCALE = 'en_US.UTF-8'

# @see http://docs.getpelican.com/en/3.5.0/settings.html#basic-settings
# @see http://docs.getpelican.com/en/3.5.0/settings.html#path-metadata
PATH_METADATA = 'articles/(?P<date>\d{4}/\d{2}/\d{2})/(?P<slug>[-a-zA-Z0-9.]*)%(?P<lang>[_a-zA-Z]{2,5})\.rst'

# @see http://docs.getpelican.com/en/3.5.0/settings.html#url-settings
ARTICLE_URL = '{date:%Y}/{date:%m}/{date:%d}/{slug}/'
ARTICLE_SAVE_AS = '{date:%Y}/{date:%m}/{date:%d}/{slug}/index.html'

# Feed generation is usually not desired when developing
FEED_ALL_ATOM = None
CATEGORY_FEED_ATOM = None
TRANSLATION_FEED_ATOM = None
AUTHOR_FEED_ATOM = None
AUTHOR_FEED_RSS = None

DEFAULT_PAGINATION = False

# https://github.com/getpelican/pelican/issues/1513
# {tag}tagName syntax not working now
# content/articles/2012/09/26/python-create-html-element-dynamically%en.rst

THEME = 'theme'

PLUGIN_PATHS = ['plugins']
PLUGINS = ['i18n_subsites', 'edit_on_github', 'embed_github_repository_file',
           'embed_picasaweb_image', 'adsense']

# my custom setting for HTML meta info
META_KEYWORDS = 'Web Development, Python, SCSS, blog'
META_DESCRIPTION = 'My blog and sharing'

# mapping: language_code -> settings_overrides_dict
I18N_SUBSITES = {
  'zh': {
    'SITENAME': '理論與實作',
    'AUTHOR': '戴上為',
    'LOCALE': 'zh_TW.UTF-8',
    'META_KEYWORDS': 'Web開發, Python, SCSS, 部落格',
    'META_DESCRIPTION': '我的部落格及分享',
  },
  'th': {
    'SITENAME': 'ทฤษฎีและการปฏิบัติ',
    'AUTHOR': 'ฉ่ง-หวี',
    'LOCALE': 'th_TH.UTF-8',
    'META_KEYWORDS': 'ภาษาบาลีสวดมนต์',
  },
}
I18N_UNTRANSLATED_ARTICLES = 'remove'

# generate only index.html and pages and articles. (no archives, tags, categories)
#DIRECT_TEMPLATES = ['index']
# use metadata attribute 'order' in page files for ordering
# @see http://docs.getpelican.com/en/3.5.0/settings.html#url-settings
PAGE_ORDER_BY = 'order'

# CONTENT_DIR_URL is the setting for edit_on_github plugin
CONTENT_DIR_URL = u'https://github.com/siongui/userpages/tree/master/content'

# Uncomment following line if you want document-relative URLs when developing
#RELATIVE_URLS = True

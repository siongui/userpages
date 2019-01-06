#!/usr/bin/env python
# -*- coding: utf-8 -*- #

# Creating reStructuredText Directives
# @see http://docutils.sourceforge.net/docs/howto/rst-directives.html
from docutils.parsers.rst import directives, Directive
from docutils import nodes

from pelican import signals
_CONTENT_PATH = None
_DEBUG = False

from os.path import basename
from os.path import join


ad1st = """
"""

ad2nd = """
"""

ad3rd = """
"""

ad4th = """
"""

ad5th = """
"""

ad6th = """
"""

class embed_adsense_code(Directive):
  required_arguments = 1
  has_content = False

  def run(self):
    sel = self.arguments[0].strip()
    html = ""
    if sel == "1":
      html = ad1st
    if sel == "2":
      html = ad2nd
    if sel == "3":
      html = ad3rd
    if sel == "4":
      html = ad4th
    if sel == "5":
      html = ad5th
    if sel == "6":
      html = ad6th

    return [nodes.raw('', html, format='html')]


def init_adsense_plugin(pelican_obj):
  global _CONTENT_PATH
  if _CONTENT_PATH is None:
    _CONTENT_PATH = pelican_obj.settings['PATH']


def register():
  signals.get_generators.connect(init_adsense_plugin)
  directives.register_directive('adsu', embed_adsense_code)

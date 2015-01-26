#!/usr/bin/env python
# -*- coding: utf-8 -*- #

from pelican import signals
from os.path import join


def add_github_url_to_content_object(content_instance):
  """
  :param content_instance: @see: https://github.com/getpelican/pelican/blob/master/pelican/contents.py
  """
  CONTENT_DIR_URL = content_instance.settings.get('CONTENT_DIR_URL', None)
  if CONTENT_DIR_URL is None:
    return

  content_instance.github_url = join(CONTENT_DIR_URL,
      content_instance.get_relative_source_path())


# @see http://docs.getpelican.com/en/3.5.0/plugins.html#list-of-signals
def register():
  signals.content_object_init.connect(add_github_url_to_content_object)

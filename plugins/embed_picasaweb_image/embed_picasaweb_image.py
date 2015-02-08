#!/usr/bin/env python
# -*- coding: utf-8 -*- #

# Creating reStructuredText Directives
# @see http://docutils.sourceforge.net/docs/howto/rst-directives.html
from docutils.parsers.rst import directives, Directive
from docutils import nodes


class embed_picasaweb_image(Directive):
  required_arguments = 1
  optional_arguments = 0
  final_argument_whitespace = True
  option_spec = { 'album_name'  : directives.unchanged,
                  'css_class'   : directives.unchanged,
                  'description' : directives.unchanged,
                  'album_url'   : directives.uri,
                  'image_url'   : directives.uri,
                }
  has_content = False

  def run(self):
    url = directives.uri(self.arguments[0])
    album_name = self.options.get('album_name', None)
    album_url = self.options.get('album_url', None)
    image_url = self.options.get('image_url', None)
    css_class = self.options.get('css_class', None)
    description = self.options.get('description', u'')

    if album_name and album_url:
      html = u'<div class="{}"><a href="{}"><image src="{}"></a><div>{}</div><div class="album">From Album: <a href="{}">{}</a></div></div>'.format(
             css_class, image_url, url, description, album_url, album_name)
    else:
      html = u'<div class="{}"><a href="{}"><image src="{}"></a><div>{}</div></div>'.format(
             css_class, image_url, url, description)

    return [nodes.raw('', html, format='html')]


def register():
  directives.register_directive('embed_picasaweb_image', embed_picasaweb_image)

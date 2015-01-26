#!/usr/bin/env python
# -*- coding: utf-8 -*- #

# Creating reStructuredText Directives
# @see http://docutils.sourceforge.net/docs/howto/rst-directives.html
from docutils.parsers.rst import directives, Directive
from docutils import nodes


class show_github_file(Directive):
  required_arguments = 3
  has_content = False

  def run(self):
    username = self.arguments[0].strip()
    repo = self.arguments[1].strip()
    relative_path = self.arguments[2].strip()

    html = '<script type="text/javascript" src="http://www.gistfy.com/github/{}/{}/{}"></script>'.format(username, repo, relative_path)
    html += "<noscript>You need to enable JavaScript to see {}!</noscript>".format(relative_path)
    return [nodes.raw('', html, format='html')]


def register():
  directives.register_directive('show_github_file', show_github_file)

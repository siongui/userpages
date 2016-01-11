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

# http://pygments.org/docs/quickstart/
from pygments import highlight
from pygments.lexers import guess_lexer_for_filename
from pygments.lexers import get_lexer_by_name
from pygments.formatters import HtmlFormatter
from pygments.util import ClassNotFound


class show_github_file(Directive):
  required_arguments = 3
  has_content = False

  def run(self):
    username = self.arguments[0].strip()
    repo = self.arguments[1].strip()
    relative_path = self.arguments[2].strip()
    filename = basename(relative_path)

    repo_url = 'https://github.com/{}/{}'.format(username, repo)
    github_url = '{}/blob/master/{}'.format(repo_url, relative_path)
    raw_url = 'https://raw.githubusercontent.com/{}/{}/master/{}'.format(username, repo, relative_path)

    # FIXME: do not assume PATH='content'
    abs_path = join(_CONTENT_PATH, relative_path[8:])
    try:
      with open(abs_path, 'r') as f:
        code = f.read()
        try:
          lexer = guess_lexer_for_filename(filename, code)
        except ClassNotFound:
          if _DEBUG: print("guess fail: {}".format(filename))
          lexer = get_lexer_by_name("text")
        html = """<figure class="github-file">
            <figcaption>
              <a href="{}">{}</a> |
              <a href="{}">repository</a> |
              <a href="{}">view raw</a>
            </figcaption><div class="code-file">
            """.format(github_url, filename, repo_url, raw_url)
        html += highlight(code, lexer, HtmlFormatter(linenos='table'))
        html += '</div></figure>'

    except IOError:
      # use Gistfy
      if _DEBUG: print("IO fail: {}".format(filename))
      html = '<script type="text/javascript" src="//www.gistfy.com/github/{}/{}/{}"></script>'.format(username, repo, relative_path)
      html += "<noscript>You need to enable JavaScript to see {}!</noscript>".format(relative_path)

    return [nodes.raw('', html, format='html')]


def init_github_file_plugin(pelican_obj):
  global _CONTENT_PATH
  if _CONTENT_PATH is None:
    _CONTENT_PATH = pelican_obj.settings['PATH']


def register():
  signals.get_generators.connect(init_github_file_plugin)
  directives.register_directive('show_github_file', show_github_file)

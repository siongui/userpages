#!/usr/bin/env python
# -*- coding: utf-8 -*- #
from __future__ import unicode_literals

import os
from docutils import nodes
from docutils import utils
from docutils import io
from docutils import statemachine
from docutils.parsers.rst import directives
from docutils.parsers.rst import Directive
from docutils.parsers.rst import states
from docutils.utils.error_reporting import SafeString
from docutils.utils.error_reporting import ErrorString
from pygments import highlight
from pygments.lexers import guess_lexer
from pygments.lexers import guess_lexer_for_filename
from pygments.formatters import HtmlFormatter


class Code(Directive):
  """Include source code file in posts.

  Usage:
    .. code-file:: {{ filename }}

  """
  required_arguments = 1
  optional_arguments = 0
  final_argument_whitespace = True
  option_spec = {'encoding': directives.encoding,
                 'tab-width': int,
                 'start-line': int,
                 'end-line': int,
                 'start-after': directives.unchanged_required,
                 'end-before': directives.unchanged_required,
                 'class': directives.class_option,
                 'name': directives.unchanged}

  standard_include_path = os.path.join(os.path.dirname(states.__file__),
                                         'include')

  def run(self):
    """Include source code file as part of the content of this reST file."""
    if not self.state.document.settings.file_insertion_enabled:
      raise self.warning('"%s" directive disabled.' % self.name)
    # determine path
    source = self.state_machine.input_lines.source(
        self.lineno - self.state_machine.input_offset - 1)
    source_dir = os.path.dirname(os.path.abspath(source))
    path = directives.path(self.arguments[0])
    if path.startswith('<') and path.endswith('>'):
        path = os.path.join(self.standard_include_path, path[1:-1])
    path = os.path.normpath(os.path.join(source_dir, path))
    path = utils.relative_path(None, path)
    path = nodes.reprunicode(path)

    encoding = self.options.get(
        'encoding', self.state.document.settings.input_encoding)
    e_handler=self.state.document.settings.input_encoding_error_handler
    tab_width = self.options.get(
        'tab-width', self.state.document.settings.tab_width)

    # read file
    try:
      self.state.document.settings.record_dependencies.add(path)
      include_file = io.FileInput(source_path=path,
                                  encoding=encoding,
                                  error_handler=e_handler)
    except UnicodeEncodeError, error:
      raise self.severe(u'Problems with "%s" directive path:\n'
                        'Cannot encode input file path "%s" '
                        '(wrong locale?).' %
                        (self.name, SafeString(path)))
    except IOError, error:
      raise self.severe(u'Problems with "%s" directive path:\n%s.' %
                (self.name, ErrorString(error)))

    startline = self.options.get('start-line', None)
    endline = self.options.get('end-line', None)
    try:
      if startline or (endline is not None):
        lines = include_file.readlines()
        rawtext = ''.join(lines[startline:endline])
      else:
        rawtext = include_file.read()
    except UnicodeError, error:
      raise self.severe(u'Problem with "%s" directive:\n%s' %
                        (self.name, ErrorString(error)))

    # start-after/end-before: no restrictions on newlines in match-text,
    # and no restrictions on matching inside lines vs. line boundaries
    after_text = self.options.get('start-after', None)
    if after_text:
      # skip content in rawtext before *and incl.* a matching text
      after_index = rawtext.find(after_text)
      if after_index < 0:
        raise self.severe('Problem with "start-after" option of "%s" '
                          'directive:\nText not found.' % self.name)
      rawtext = rawtext[after_index + len(after_text):]
    before_text = self.options.get('end-before', None)
    if before_text:
      # skip content in rawtext after *and incl.* a matching text
      before_index = rawtext.find(before_text)
      if before_index < 0:
        raise self.severe('Problem with "end-before" option of "%s" '
                          'directive:\nText not found.' % self.name)
      rawtext = rawtext[:before_index]

    #include_lines = statemachine.string2lines(rawtext, tab_width,
    #                                          convert_whitespace=True)
    #for line in include_lines:
    #  print(line)
    lexer = guess_lexer_for_filename(self.arguments[0], rawtext)
    html_code = highlight(rawtext, lexer, HtmlFormatter())
    html = ( u'<figure class="code"><figcaption>%s' +
             u'</figcaption>%s</figure>' ) % \
        (self.arguments[0], html_code)
    #print(html)
    return [nodes.raw('', html, format='html')]


def register():
  directives.register_directive('code-file', Code)

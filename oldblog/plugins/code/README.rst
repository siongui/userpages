=============================================
Pelican Plugin for Including Source Code File
=============================================

Pelican plugin to include source code file in your post by reStructuredText directive.

Installation
============

Follow the `instructions <http://docs.getpelican.com/en/latest/plugins.html>`_
in `official Pelican documentation <http://docs.getpelican.com/>`_.

Usage
=====

Use the following reStructuredText directive in your post:

.. code-block:: rst

  .. code-file:: {{ filename }}
      :encoding: (optional)
      :tab-width: (optional)
      :start-line: (optional)
      :end-line: (optional)
      :start-after: (optional)
      :end-before: (optional)
      :class: (optional)
      :name: (optional)

The source code file must be put in the same directory of your post.
The only necessary argument is {{ filename }}. For details descriptions 
of options, please see the options of `Including an External Document Fragment <http://docutils.sourceforge.net/docs/ref/rst/directives.html#including-an-external-document-fragment>`__ (This directive is derived from it.).


UNLICENSE
=========

This plugin is released in public domain. See `UNLICENSE <http://unlicense.org/>`_.

References
==========

1. `Including an External Document Fragment <http://docutils.sourceforge.net/docs/ref/rst/directives.html#including-an-external-document-fragment>`__

2. `source code <http://sourceforge.net/p/docutils/code/HEAD/tree/tags/docutils-0.11/docutils/parsers/rst/directives/misc.py>`_ of "Including an External Document Fragment"

3. `Creating reStructuredText Directives <http://docutils.sourceforge.net/docs/howto/rst-directives.html>`_

4. `Introduction and Quickstart - Pygments <http://pygments.org/docs/quickstart/>`_


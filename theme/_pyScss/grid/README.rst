=============================
Simple pyScss Responsive Grid
=============================

A simple responsive grid system written in `SASS <http://sass-lang.com/>`_
, compiled by `pyScss <https://github.com/Kronuz/pyScss>`_.

Quick Start
===========

Install `pyScss <https://github.com/Kronuz/pyScss>`_:

.. code-block:: bash

  $ pip install pyScss

`GNU make <http://www.gnu.org/software/make/>`
can be used to compile scss to css:

.. code-block:: bash

  $ make

Or you can use Python to compile scss to css:

.. code-block:: bash

  $ python build.py


Two css class are provided: ``.row`` and ``.spanX-of-12``,
where ``X`` is the number of columns in a row of total 12 columns.

Because ``display: inline-block`` is used for ``.spanX-of-12``, there is
`whitespace issue <http://designshack.net/articles/css/whats-the-deal-with-display-inline-block/>`_
in such design. Please refer to
`this article <http://css-tricks.com/fighting-the-space-between-inline-block-elements/>`_
for solution.

The following is an example to use the grid in your website:

.. code-block:: html
  <div class="row"><!--
    --><div class="span9-of-12">
      main content here
    </div><!--
    --><div class="span3-of-12">
      sidebar here
    </div><!--
  --></div>

Note that ``<!--`` and ``-->`` is used to prevent whitespace issue.

Demo
====

Please check my `blog <http://siongui.github.io/>`_ on Github Pages.

UNLICENSE
=========

Released in public domain. See `UNLICENSE <http://unlicense.org/>`_.

References
==========

  `What’s the Deal With Display: Inline-Block? <http://designshack.net/articles/css/whats-the-deal-with-display-inline-block/>`_

  `Fighting the Space Between Inline Block Elements <http://css-tricks.com/fighting-the-space-between-inline-block-elements/>`_

  `Should You Use Inline-Blocks As A Substitute For Floats? <http://www.vanseodesign.com/css/inline-blocks/>`_

  `FLUID AND RESPONSIVE GRID LAYOUTS <http://www.stephanboyer.com/post/41/fluid-and-responsive-grid-layouts>`_

  `Yet Another CSS Grid System <http://sans0r.github.io/yacgs/>`_


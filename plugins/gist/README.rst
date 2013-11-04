============
Pelican Gist
============

Pelican Plugin to embed gist_ in your pages and articles in reStructuredText syntax.

Usage
=====

.. code-block:: rst

  .. gist:: GIST_ID GITHUB_USERNAME

Will result in

.. code-block:: html

  <script src="https://gist.github.com/{{ GITHUB_USERNAME }}/{{ GIST_ID }}.js"></script>

in your generated html webpage.

.. _gist: https://gist.github.com/

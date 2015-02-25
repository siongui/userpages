[JavaScript] Width of Browser Window in Pixel
#############################################

:date: 2012-09-16 00:06
:modified: 2015-02-25 09:30
:tags: html, JavaScript, DOM
:category: JavaScript
:summary: Width of browser window in pixel.


Here is an interesting and sometimes practical question:

*How do I get the width of browser window in pixel?*

The answer is easy: add a *div* element right after the *body* tag, make it 100%
wide, and get the *offsetWidth* property of the *div*.

.. rubric:: `Demo <{filename}getWindowWidthPixel.html>`_
      :class: align-center

Source Code for Demo:

.. show_github_file:: siongui userpages content/articles/2012/09/16/getWindowWidthPixel.html

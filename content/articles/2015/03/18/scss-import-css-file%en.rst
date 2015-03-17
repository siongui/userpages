Import CSS File in SCSS with pyScss
###################################

:date: 2015-03-18 02:28
:tags: CSS, html, SASS, SCSS
:category: CSS
:summary: Include regular CSS file in SCSS file with pyScss.


pyScss_ is a Scss_ compiler for Python. Sometimes we need to include (import)
regular CSS file in SCSS file. This post shows how to do it with pyScss_:

Assume that we have a css file called ``pygments.css``, we want to include this
file in our scss file:

1. Rename *pygments.css* to **_pygments.scss**
   (Prepend an underscore and switch the extension to ``scss``.)

2. Use `@import directive`_ to include the renamed CSS file in SCSS file:

   .. code-block:: scss

     @import "pygments";


Tested on: ``pyScss 1.3.4``

----

Reference:

.. [1] `sass - Import regular css file in scss file? - Stack Overflow <http://stackoverflow.com/questions/7111610/import-regular-css-file-in-scss-file>`_


.. _Scss: http://sass-lang.com/

.. _pyScss: https://github.com/Kronuz/pyScss

.. _@import directive: http://sass-lang.com/documentation/file.SASS_REFERENCE.html#import

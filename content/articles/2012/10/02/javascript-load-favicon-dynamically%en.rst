[JavaScript] Load Favicon Dynamically
#####################################

:date: 2012-10-02 01:50
:modified: 2015-02-23 10:46
:tags: html, CORS, JavaScript
:category: JavaScript
:summary: Load website icon (favicon) dynamically.
:adsu: yes


This post is an extension of reference [1]_, which discusses how to load
external JavaScript or CSS dynamically. The same idea can be applied to load
favicon_ dynamically. Give the URL or path of favicon_ to the following
function, and then favicon_ will be loaded dynamically:

.. code-block:: javascript

  function LoadFavicon(url) {
    var ext = document.createElement('link');
    ext.setAttribute("rel", "shortcut icon");
    ext.setAttribute("href", url);
    document.getElementsByTagName("head")[0].appendChild(ext);
  }

.. adsu:: 2

Example Usage of LoadFavicon Function
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Assume that the path of favicon is located at */favicon.ico* under the domain of
the website. You can call:

.. code-block:: javascript

  LoadFavicon('/favicon.ico');

to load favicon dynamically. Even if your favicon is located at a different
domain, you can still pass the URL of favicon at the different domain to above
function and it also works!

.. adsu:: 3

Enjoy!

----

Reference:

.. [1] `Load External JavaScript or CSS file Dynamically <{filename}../../06/18/load-external-javascript-or-css-file-dynamically%en.rst>`_

.. _favicon: http://en.wikipedia.org/wiki/Favicon

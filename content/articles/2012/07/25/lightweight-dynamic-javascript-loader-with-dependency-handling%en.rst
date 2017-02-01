Lightweight Dynamic JavaScript Loader with Dependency Handling
##############################################################

:date: 2012-07-25 05:07
:modified: 2015-04-03 08:08
:tags: JavaScript, XMLHttpRequest
:category: JavaScript
:summary: Load JavaScript_ scripts dynamically with dependencies.
:adsu: yes


I write a lightweight dynamic JavaScript_ loader with dependency handling, which
means it can dynamically loads JavaScript_ files with the order you want. The
following is the source code for the dynamic JavaScript_ loader:

.. show_github_file:: siongui userpages content/code/javascript-loader/dynamic.js
.. adsu:: 2

Usage
+++++

I will show how to use the JavaScript_ loader by example.

Assume that we need to load *a.js*, *b.js*, *c.js*, *d.js*.

They are located at *http://example.com/static/js/*.

*a.js* has no dependency.

*b.js* has to be loaded after *a.js* is loaded.

*c.js* has to be loaded after *a.js* is loaded.

*d.js* has to be loaded after *b.js* and *c.js* are loaded.

The following is the usage code for above example:

.. adsu:: 3
.. show_github_file:: siongui userpages content/code/javascript-loader/usage.js

Enjoy!


.. _JavaScript: https://www.google.com/search?q=JavaScript

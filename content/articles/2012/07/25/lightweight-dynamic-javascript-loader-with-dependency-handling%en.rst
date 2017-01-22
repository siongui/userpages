Lightweight Dynamic JavaScript Loader with Dependency Handling
##############################################################

:date: 2012-07-25 05:07
:modified: 2015-04-03 08:08
:tags: JavaScript, XMLHttpRequest
:category: JavaScript
:summary: Load JavaScript scripts dynamically with dependencies.
:adsu: yes


I write a lightweight dynamic JavaScript loader with dependency handling, which
means it can dynamically loads JavaScript files with the order you want. The
following is the source code for the dynamic JavaScript loader:

.. show_github_file:: siongui userpages content/code/javascript-loader/dynamic.js

Usage
+++++

I will show how to use the JavaScript loader by example.

Assume that we need to load *a.js*, *b.js*, *c.js*, *d.js*.

They are located at *http://example.com/static/js/*.

*a.js* has no dependency.

*b.js* has to be loaded after *a.js* is loaded.

*c.js* has to be loaded after *a.js* is loaded.

*d.js* has to be loaded after *b.js* and *c.js* are loaded.

The following is the usage code for above example:

.. show_github_file:: siongui userpages content/code/javascript-loader/usage.js

Enjoy!

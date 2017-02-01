[JavaScipt] Cross-Browser HTTP GET Request
##########################################

:date: 2012-10-05 05:18
:modified: 2015-03-02 23:19
:tags: JavaScript, HTTP GET, XMLHttpRequest
:category: JavaScript
:summary: JavaScript cross-browser HTTP GET request
:adsu: yes


This post shows how to make `HTTP GET request`_ in cross-browser way. The
following *HTTPGET* function demonstrates how to make such AJAX request:

.. show_github_file:: siongui userpages content/code/javascript-http-get/get.js
.. adsu:: 2

Usage
+++++

.. show_github_file:: siongui userpages content/code/javascript-http-get/usage.js
.. adsu:: 3

Note that you can only make requests in the same domain. To make cross-domain
requests, please see reference [1]_ for more detail. To make cross-browser HTTP
POST request, please see reference [2]_. To make HTTP POST requests to GAE
Python server, see reference [3]_.

----

References:

.. [1] `JavaScript Cross-Browser Cross-Domain XMLHttpRequest (XDomainRequest in IE) <{filename}../../09/25/javascript-cors-xmlhttprequest%en.rst>`_

.. [2] `[JavaScipt] Cross-Browser HTTP POST Request <{filename}javascript-http-post-request%en.rst>`_

.. [3] `AJAX Form POST Request to Google App Engine Python <{filename}../../07/24/ajax-form-http-post-gae-python%en.rst>`_

.. [4] `XMLHttpRequest - Web API Interfaces | MDN <https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest>`_


.. _HTTP GET request: http://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods

[JavaScipt] Cross-Browser HTTP POST Request
###########################################

:date: 2012-10-05 15:40
:modified: 2015-03-02 23:15
:tags: JavaScript, HTTP POST, XMLHttpRequest
:category: JavaScript
:summary: JavaScript cross-browser HTTP POST request

This post shows how to make `HTTP POST request`_ with
`query string (key-value pairs)`_ in cross-browser way. The following *HTTPPOST*
function demonstrates how to make such AJAX request:

.. show_github_file:: siongui userpages content/code/javascript-http-post/post.js

Usage
+++++

.. show_github_file:: siongui userpages content/code/javascript-http-post/usage.js

Note that you can only make requests in the same domain. To make cross-domain
requests, please see reference [1]_ for more detail. To make cross-browser HTTP
GET request, please see reference [2]_. To make HTTP POST requests to GAE Python
server, see reference [3]_.

----

References:

.. [1] `JavaScript Cross-Browser Cross-Domain XMLHttpRequest (XDomainRequest in IE) <{filename}../../09/25/javascript-cors-xmlhttprequest%en.rst>`_

.. [2] `[JavaScipt] Cross-Browser HTTP GET Request <{filename}javascript-http-get-request%en.rst>`_

.. [3] `AJAX Form POST Request to Google App Engine Python <{filename}../../07/24/ajax-form-http-post-gae-python%en.rst>`_

.. [4] `XMLHttpRequest - Web API Interfaces | MDN <https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest>`_


.. _HTTP POST request: http://en.wikipedia.org/wiki/POST_(HTTP)

.. _query string (key-value pairs): http://en.wikipedia.org/wiki/Query_string

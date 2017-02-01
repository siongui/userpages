AJAX Form POST Request to Google App Engine Python
##################################################

:date: 2012-07-24 21:50
:modified: 2015-03-02 23:21
:tags: JavaScript, XMLHttpRequest, HTTP POST, Web application,
       Google App Engine, Python
:category: JavaScript
:summary: Send/receive data to/from `Google App Engine Python`_ servers by
          `HTTP POST`_ requests in AJAX_ way.
:adsu: yes


This post shows how to send data of `HTML form`_ in client browsers to
`GAE Python`_ server, and receive response data from the server. This technique
is common in the communication of client-server web applications.

Client side (*HTML*): Users fill the HTML form.

.. show_github_file:: siongui userpages content/code/ajax-form-post-gae-python/index.html
.. adsu:: 2

Client side (*JavaScript*): Send form data to server by `HTTP POST`_ request.
Note that the data are encoded by JavaScript_ `encodeURIComponent()`_ function
to send the data with special characters to servers robustly. ([4]_, [5]_, [6]_,
[7]_, [8]_)

.. show_github_file:: siongui userpages content/code/ajax-form-post-gae-python/post.js
.. adsu:: 3

Server side (GAE *Python*): The data from client browsers are decoded by
*urllib2.unquote()* function. ([9]_, [10]_)

.. show_github_file:: siongui userpages content/code/ajax-form-post-gae-python/post.py

Server side (GAE *Python* config):

.. show_github_file:: siongui userpages content/code/ajax-form-post-gae-python/app.yaml

----

References:

.. [1] `AJAX Tutorial - W3Schools <http://www.w3schools.com/ajax/default.asp>`_

.. [2] `AJAX Form POST Request - HTML Form POST/Submit with AJAX/Javascript Example/Tutorial <http://snipplr.com/view/3701/>`_

.. [3] `Python Runtime Environment - Python — Google Cloud Platform <https://cloud.google.com/appengine/docs/python/>`_

.. [4] `A Complete Guide to URL Escape Characters | We Rock Your Web <http://www.werockyourweb.com/url-escape-characters/>`_

.. [5] `xkr.us / javascript / escape(), encodeURI(), encodeURIComponent() <http://xkr.us/articles/javascript/encode-compare/>`_

.. [6] `javascript - Best practice: escape, or encodeURI / encodeURIComponent - Stack Overflow <http://stackoverflow.com/questions/75980/best-practice-escape-or-encodeuri-encodeuricomponent>`_

.. [7] `javascript - encodeURI or encodeURIComponent - Stack Overflow <http://stackoverflow.com/questions/4540753/encodeuri-or-encodeuricomponent>`_

.. [8] `javascript - What is the difference between decodeURIComponent and decodeURI? - Stack Overflow <http://stackoverflow.com/questions/747641/what-is-the-difference-between-decodeuricomponent-and-decodeuri>`_

.. [9] `google app engine - How to decode encodeURIComponent in GAE (python)? - Stack Overflow <http://stackoverflow.com/questions/9880173/how-to-decode-encodeuricomponent-in-gae-python>`_

.. [10] `Equivalent Javascript Functions for Python's urllib.quote() and urllib.unquote() - Stack Overflow <http://stackoverflow.com/questions/946170/equivalent-javascript-functions-for-pythons-urllib-quote-and-urllib-unquote>`_


.. _HTML form: http://www.w3schools.com/html/html_forms.asp
.. _Google App Engine Python: https://cloud.google.com/appengine/docs/python/
.. _GAE Python: https://cloud.google.com/appengine/docs/python/
.. _HTTP POST: https://www.google.com/search?q=HTTP+POST
.. _AJAX: https://www.google.com/search?q=AJAX
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _encodeURIComponent(): http://www.w3schools.com/jsref/jsref_encodeURIComponent.asp

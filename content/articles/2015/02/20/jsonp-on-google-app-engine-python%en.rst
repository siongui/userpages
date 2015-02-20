JSONP on Google App Engine Python
#################################

:date: 2015-02-20 15:11
:modified: 2015-02-20 21:14
:tags: JavaScript, JSONP, Google App Engine, Python
:category: JavaScript
:summary: Cross-domain requests by JSONP on Google App Engine for Python.


There are many explanations about JSONP_ on internet, but a complete sample code
for JSONP_ on `Google App Engine for Python`_ is very rare. After trial and
error, I finally made it work on `Google App Engine for Python`_. To put JSONP
in short, it is a mechanism to make a cross-domain request, which is usually
denied by browsers because of security concerns. The technique behind JSONP_ is
simple: use HTML *script* tag to make HTTP GET request, as if the client is
requesting an external JavaScript file. A callback function must be provided to
handle the response data from server. For more information, please refer to [1]_
and [2]_.

The following is complete sample code.

*index.html* (run on client side, i.e., browser):

.. show_github_file:: siongui userpages content/code/jsonp-gae-python/index.html

*jsonp.js* (run on client side, i.e., browser):

.. show_github_file:: siongui userpages content/code/jsonp-gae-python/jsonp.js

.. note::

  The name of callback function, *mycallback* in this case, must be consistent.
  To process the response data from server in the callback function, web
  application designer must know the structure of the response JSON data in
  advance.

*jsonp.py* (run on server side, i.e., GAE Python):

.. show_github_file:: siongui userpages content/code/jsonp-gae-python/jsonp.py

.. note::

  The following code in *JSONPPage* class:

  .. code-block:: python

    self.response.out.write(
        "%s(%s)" %
        (urllib2.unquote(self.request.get('callback')),
        json.dumps(result))
    )

  The JSON data is inside parentheses after the name of callback function.

*app.yaml* (on server side, GAE Python config file):

.. show_github_file:: siongui userpages content/code/jsonp-gae-python/app.yaml

If you want to use JSONP with anonymous callback function, see [4]_.


Tested on: ``Ubuntu Linux 14.10``, ``Google App Engine Python SDK 1.9.18``

----

References:

.. [1] `JSONP Example | Online Solutions Development <http://www.osd.net/blog/web-development/javascript/jsonp-example/>`_

.. [2] `Cross Domain Web Mashups with JQuery and Google App Engine <http://www.slideshare.net/andymckay/cross-domain-webmashups-with-jquery-and-google-app-engine>`_

.. [3] `Google App Engine for Python <https://cloud.google.com/appengine/docs/python/>`_

.. [4] `JSONP with Anonymous Callback Function <{filename}jsonp-anonymous-callback-function%en.rst>`_

.. _JSONP: http://en.wikipedia.org/wiki/JSONP

.. _Google App Engine for Python: https://cloud.google.com/appengine/docs/python/

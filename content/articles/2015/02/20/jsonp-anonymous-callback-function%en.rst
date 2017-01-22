JSONP with Anonymous Callback Function
######################################

:date: 2015-02-20 21:12
:modified: 2015-02-20 21:50
:tags: JavaScript, CORS, JSONP, Google App Engine, Python, JSON
:category: JavaScript
:summary: Cross-domain requests by JSONP with anonymous callback function on Google App Engine for Python.
:adsu: yes


In my previous post [1]_, a complete JSONP example on both client side (browser)
and server side (Google App Engine Python) is shown. The example in previous
post needs a *named, JavaScript callback function* to handle the response
returned from server. In this post, instead of using *named callback function*,
an alternative choice is shown: use an *anonymous callback function* in JSONP
request.

The main difference between *named callback function* and *anonymous callback
function* is that in named callback function, only the name of callback function
is supplied in the HTTP request, while in anonymous callback function, the whole
callback function is supplied in the HTTP request.

The following is a complete example of JSONP with anonymous callback function.

*index.html* (run on client side, i.e., browser):

.. show_github_file:: siongui userpages content/code/jsonp-anonymous-gae/index.html

*jsonp.js* (run on client side, i.e., browser):

.. show_github_file:: siongui userpages content/code/jsonp-anonymous-gae/jsonp.js

.. note::

  In :code:`'/jsonp?callback=' + encodeURIComponent(callback.toString())`, the
  whole anonymous callback function is supplied in the HTTP request.

*jsonp.py* (run on server side, i.e., GAE Python):

.. show_github_file:: siongui userpages content/code/jsonp-anonymous-gae/jsonp.py

.. note::

  In the code:

  .. code-block:: python

    self.response.out.write(
        "(%s)(%s);" %
        (urllib2.unquote(self.request.get('callback')),
        json.dumps(result))
    )

  The anonymous callback function is put inside parentheses. Without the
  parentheses, error will occur while executing the function on client. Compared
  with example in previous post [1]_, the named callback function does not need
  to be inside parentheses. This is another difference between named and
  anonymous callback function.

  Also do not forget the *semicolon* in :code:`(%s)(%s);`

*app.yaml* (on server side, GAE Python config file):

.. show_github_file:: siongui userpages content/code/jsonp-anonymous-gae/app.yaml

Caveat
~~~~~~

The anonymous function is passed as argument in the URL. If the anonymous
function is too long, this may exceed the length limit of URI, i.e., the server
will return HTTP 414 error (Requested URI too long). As a result, if a callback
function is so long such that it exceeds URI length limit, the only choice is to
use named callback function.

----

If you want to use JSONP with object instance function as callback function,
see [2]_.


Tested on: ``Ubuntu Linux 14.10``, ``Google App Engine Python SDK 1.9.18``

----

References:

.. [1] `JSONP on Google App Engine Python <{filename}jsonp-on-google-app-engine-python%en.rst>`_

.. [2] `Use Object Instance Function as JSONP Callback Function <{filename}jsonp-object-instance-callback%en.rst>`_

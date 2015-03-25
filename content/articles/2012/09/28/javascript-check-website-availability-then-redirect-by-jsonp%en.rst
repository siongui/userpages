[JavaScript] Check Website Availability and Then Redirect by JSONP
##################################################################

:date: 2012-09-28 14:17
:modified: 2015-03-25 20:36
:tags: JavaScript, JSONP, CORS
:category: JavaScript
:summary: Redirect users to another website if the website is available.


The problem to be solved:

  Assume there are two websites. One of the domain names of the two websites is
  *a.com*, and the other is *b.com*. When a visitor visits the website of
  *b.com*, the browser of the visitor checks (by JavaScript) if the website of
  *a.com* is available for the visitor. If the website of *a.com* is available
  for the visitor, the visitor will be redirected to the website of *a.com* (by
  JavaScript URL redirect).

To solve the above problem, we need:

1. Implement JSONP_ service at the website of *a.com*.

2. Insert JavaScript availability checking and redirection code at the website
   of *b.com*.

We will give details one by one.

Implement JSONP service at the website of a.com
+++++++++++++++++++++++++++++++++++++++++++++++

I wrote three posts on how to implement JSONP_ service on Google App Engine,
please refer to following posts:

  - `JSONP on Google App Engine Python`_

  - `JSONP with Anonymous Callback Function`_

  - `Use Object Instance Function as JSONP Callback Function`_

Insert JavaScript availability checking and redirection code at the website of b.com
++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

Here comes the tricky part of the solution. We will use
`JSONP with Anonymous Callback Function`_ to check if the website of *a.com* is
available for the visitor, and if available, the visitor will be redirect to
*a.com*. To achieve this, insert the following JavaScript to the webpage of the
website of *b.com* (Assume the JSONP URL of *a.com* is */jsonp*).

.. code-block:: javascript
  :linenos: table

  if (window.location.host == 'b.com') {
    var callback = function() {
      window.location = 'http://a.com' + window.location.pathname + window.location.search;
    };
    var qry = '?callback=' + encodeURIComponent('('+callback.toString()+')');
    var ext = document.createElement('script');
    ext.setAttribute('src', 'http://a.com/jsonp' + qry);
    document.getElementsByTagName("head")[0].appendChild(ext);
  }

In line 2~4 is the JSONP callback function, which does only one thing: redirect
the user to the website of *a.com*. If the website of *a.com* is down or not
available, the callback function will not be executed and hence the visitor will
still stay at the website of *b.com*. If the website of *a.com* is available,
the callback function will be executed and hence the visitor will be redirected
to *a.com*. This is exactly the solution we need.



.. _JSONP: {tag}JSONP

.. _JSONP on Google App Engine Python: {filename}../../../2015/02/20/jsonp-on-google-app-engine-python%en.rst

.. _JSONP with Anonymous Callback Function: {filename}../../../2015/02/20/jsonp-anonymous-callback-function%en.rst

.. _Use Object Instance Function as JSONP Callback Function: {filename}../../../2015/02/20/jsonp-object-instance-callback%en.rst

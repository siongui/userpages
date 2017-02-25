JavaScript Cross-Browser Cross-Domain XMLHttpRequest (XDomainRequest in IE)
###########################################################################

:date: 2012-09-25 22:40
:modified: 2015-02-24 16:20
:tags: html, JavaScript, CORS, DOM, HTTP GET, HTTP POST, XMLHttpRequest
:category: JavaScript
:summary: Cross-domain AJAX requests
:og_image: http://www.javatpoint.com/images/javascript/javascript_logo.png
:adsu: yes


This post gives a client-side sample code for very useful technique in AJAX_
programming: Cross-Domain, Cross-Browser XMLHttpRequest requests (XDomainRequest
for IE8+).

Before doing Cross-Domain AJAX requests, Cross-Origin Resource Sharing (CORS)
must be enabled on servers first. Visit `Enable CORS website`_ to see how to
enable CORS on your server.

.. rubric:: `Demo <{filename}/code/javascript-xmlhttprequest/crossDomainXHR.html>`_
      :class: align-center

Source Code for Demo (*HTML*):

.. show_github_file:: siongui userpages content/code/javascript-xmlhttprequest/crossDomainXHR.html
.. adsu:: 2

Source Code for Demo (*JavaScript*):

.. show_github_file:: siongui userpages content/code/javascript-xmlhttprequest/xhr.js
.. adsu:: 3

The *AJAXRequest* function provides the Cross-Domain, Cross-Browser XHR. Put the
function in your code and re-write the content of *callback* and *failCallback*
to fit your needs.

----

References:

.. [1] `CORS for XHR in IE10 <http://blogs.msdn.com/b/ie/archive/2012/02/09/cors-for-xhr-in-ie10.aspx>`_

.. [2] `Cross-Domain, Cross-Browser AJAX Requests <https://www.bionicspirit.com/blog/2011/03/24/cross-domain-requests.html>`_

.. [3] `XDomainRequest object <https://msdn.microsoft.com/en-us/library/ie/cc288060(v=vs.85).aspx>`_

.. [4] `enable cross-origin resource sharing <http://enable-cors.org/>`_


.. _AJAX: http://en.wikipedia.org/wiki/Ajax_(programming)

.. _Enable CORS website: http://enable-cors.org/

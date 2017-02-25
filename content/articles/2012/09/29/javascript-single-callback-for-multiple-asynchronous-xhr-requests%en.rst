[JavaScript] Single Callback For Multiple Asynchronous AJAX Requests (XMLHttpRequest)
#####################################################################################

:date: 2012-09-29 03:16
:modified: 2015-03-18 04:28
:tags: JavaScript, HTTP GET, CORS, XMLHttpRequest
:category: JavaScript
:summary: JavaScript single callback for multiple asynchronous XHR requests.
:og_image: http://www.javatpoint.com/images/javascript/javascript_logo.png
:adsu: yes

Assume that multiple Asynchronous AJAX_ Requests (XMLHttpRequest_) are made, and
we would like to have a single final callback_ function to be run after all the
requests are completed successfully. How do we do it? (Reference [1]_ is
actually the same question as the above.)

The following code is my solution to the question:

Source Code (*JavaScript*)
++++++++++++++++++++++++++

.. show_github_file:: siongui userpages content/code/single-callback-multiple-xhr/xhr.js

There are two main functions in above code snippet. One is *AjaxRequest*, which
is the function to issue single AJAX request (XMLHttpRequest). The other is
*AjaxRequestsMulti*, which is the function to issue multiple AJAX requests. I
put a lot of comments in the code in order to trace the code easily.

Usage and Demo
++++++++++++++

.. adsu:: 2

.. rubric:: `Demo <{filename}/code/single-callback-multiple-xhr/asynchronous.html>`_
      :class: align-center

Source Code for Demo (*JavaScript*):

.. show_github_file:: siongui userpages content/code/single-callback-multiple-xhr/usage.js
.. adsu:: 3

The variable *urls* contains the URLs of AJAX requests, overwrite this variable
to fit your needs. The function *callbackMulti* is the callback function to be
run after all AJAX requests are completed successfully, and *failCallbackMulti*
will be run if some requests fail.

Source Code for Demo (*HTML*):

.. show_github_file:: siongui userpages content/code/single-callback-multiple-xhr/asynchronous.html

----

References:

.. [1] `jquery - Javascript callback for multiple ajax calls - Stack Overflow <http://stackoverflow.com/questions/4368946/javascript-callback-for-multiple-ajax-calls>`_


.. _AJAX: http://en.wikipedia.org/wiki/Ajax_(programming)

.. _XMLHttpRequest: https://duckduckgo.com/?q=XMLHttpRequest

.. _callback: http://en.wikipedia.org/wiki/Callback_%28computer_programming%29

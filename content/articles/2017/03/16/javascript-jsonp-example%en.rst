[JavaScript] JSONP Example
##########################

:date: 2017-03-16T03:10+08:00
:tags: JavaScript, JSON, JSONP, CORS
:category: JavaScript
:summary: JSONP example - show your HTTP request header.
:og_image: http://www.zackgrossbart.com/hackito/wp-content/uploads/2009/04/jsonp_flow.jpg
:adsu: yes


Demo
++++

JSONP_ is a method which help you retrieve data from other domains.
The following demo will show the header of your HTTP request via JSONP.

.. raw:: html

  <div style="background-color: Azure; padding: 5px;">
    <button id="btn">Click to get HTTP header via JSONP</button>
    <pre id="result"></pre>
  </div>

.. raw:: html

  <script>
  'use strict';

  var btn = document.getElementById("btn");
  var result = document.getElementById("result");

  function myCallback(acptlang) {
    result.innerHTML = JSON.stringify(acptlang, null, 2);
  }

  function jsonp() {
    result.innerHTML = "Loading ...";
    var tag = document.createElement("script");
    tag.src = "https://ajaxhttpheaders.appspot.com/?callback=myCallback";
    document.querySelector("head").appendChild(tag);
  }

  btn.addEventListener("click", jsonp);
  </script>


Source Code and Explanation
+++++++++++++++++++++++++++

**HTML**:

.. code-block:: html

  <button id="btn">Click to get HTTP header via JSONP</button>
  <pre id="result"></pre>

.. adsu:: 2

**JavaScript**:

Someone provide online service which returns HTTP request headers via JSONP.
The URL of the service is:

::

   https://ajaxhttpheaders.appspot.com/?callback={{YOUR_CALLBACK_NAME}}

If you want to know how to implement the server which returns data via JSONP,
see [1]_. The front-end JavaScript code for retrieving HTTP request headers from
the above service via JSONP is as follows:

.. code-block:: javascript

  'use strict';

  var btn = document.getElementById("btn");
  var result = document.getElementById("result");

  function myCallback(acptlang) {
    result.innerHTML = JSON.stringify(acptlang, null, 2);
  }

  function jsonp() {
    result.innerHTML = "Loading ...";
    var tag = document.createElement("script");
    tag.src = "https://ajaxhttpheaders.appspot.com/?callback=myCallback";
    document.querySelector("head").appendChild(tag);
  }

  btn.addEventListener("click", jsonp);

You need to assign a callback function, which receives data from the server, in
the JSONP request. The name of the callback function in the demo is
*myCallback*. In the click event handler, the JSONP action is performed. A
*script* element is created and then appended to the end of *head* element. The
source of the *script* element is the URL of the online service, and the name of
the callback is assigned in the query string of the URL.

.. adsu:: 3

.. note::

   If you web page is served via HTTPS, the server that returns data via JSONP
   also needs to serve via HTTPS. Otherwise browsers will block the request and
   make the request fail.


----

Tested on:

- ``Chromium Version 56.0.2924.76 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``

----

References:

.. [1] | `localization - JavaScript for detecting browser language preference - Stack Overflow <http://stackoverflow.com/a/3335420>`_
       | `GitHub - dansingerman/app-engine-headers: Google app engine application that is the server side counterpart to https://github.com/dansingerman/jQuery-Browser-Language <https://github.com/dansingerman/app-engine-headers>`_
.. [2] `[Vue.js] Pretty Print JSON String <{filename}../15/vuejs-pretty-print-json-string%en.rst>`_

.. _JSONP: https://www.google.com/search?q=JSONP
.. _JavaScript: https://www.google.com/search?q=JavaScript

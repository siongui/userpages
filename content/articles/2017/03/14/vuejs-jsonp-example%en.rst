[Vue.js] JSONP Example
######################

:date: 2017-03-14T02:16+08:00
:tags: Vue.js, JavaScript, JSON, JSONP, CORS
:category: Vue.js
:summary: JSONP_ example by Vue.js_.
:og_image: https://avatars1.githubusercontent.com/u/6128107?v=3&s=200
:adsu: yes


Demo
++++

JSONP_ is a method which help you retrieve data from other domains.
The following demo will show the header of your HTTP request via JSONP and
Vue.js_.

.. raw:: html

  <div id="vueapp" style="background-color: Azure; padding: 5px;">
    <button v-on:click.once="jsonp">Click to get HTTP header via JSONP</button>
    <pre>{{ info }}</pre>
  </div>

  <script src="https://unpkg.com/vue@2.2.4/dist/vue.js"></script>

.. raw:: html

  <script>
  'use strict';

  var vm = new Vue({
    el: '#vueapp',
    data: {
      info: ""
    },
    methods: {
      jsonp: function() {
        this.info = "Loading ...";
        var tag = document.createElement("script");
        tag.src = "https://ajaxhttpheaders.appspot.com/?callback=myCallback";
        document.querySelector("head").appendChild(tag);
      }
    }
  })

  function myCallback(acptlang) {
    vm.info = acptlang;
  }
  </script>


Source Code and Explanation
+++++++++++++++++++++++++++

**HTML**:

.. code-block:: html

  <div id="vueapp">
    <button v-on:click.once="jsonp">Click to get HTTP header via JSONP</button>
    <pre>{{ info }}</pre>
  </div>

  <script src="https://unpkg.com/vue@2.2.4/dist/vue.js"></script>

Register an click event handler of the *button* element, and run the click event
handler only once.

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

  var vm = new Vue({
    el: '#vueapp',
    data: {
      info: ""
    },
    methods: {
      jsonp: function() {
        this.info = "Loading ...";
        var tag = document.createElement("script");
        tag.src = "https://ajaxhttpheaders.appspot.com/?callback=myCallback";
        document.querySelector("head").appendChild(tag);
      }
    }
  })

  function myCallback(acptlang) {
    vm.info = acptlang;
  }

The name of the callback function in demo is *myCallback*. In the click event
handler, a *script* element is created and then appended to the end of *head*
element. The source of the *script* element is the URL of the online service,
and the name of the callback is assigned in the query string of the URL.

.. adsu:: 3

.. note::

   If you web page is served via HTTPS, the server that returns data via JSONP
   also needs to serve via HTTPS. Otherwise browsers will block the request and
   make the request fail.


----

Tested on:

- ``Chromium Version 56.0.2924.76 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``
- ``Vue.js 2.2.4``

----

References:

.. [1] | `localization - JavaScript for detecting browser language preference - Stack Overflow <http://stackoverflow.com/a/3335420>`_
       | `GitHub - dansingerman/app-engine-headers: Google app engine application that is the server side counterpart to https://github.com/dansingerman/jQuery-Browser-Language <https://github.com/dansingerman/app-engine-headers>`_
.. [2] `JSONP on Google App Engine Python <{filename}../../../2015/02/20/jsonp-on-google-app-engine-python%en.rst>`_
.. [3] `Event Handling — Vue.js <https://vuejs.org/v2/guide/events.html>`_

.. _Vue.js: https://vuejs.org/
.. _JSONP: https://www.google.com/search?q=JSONP

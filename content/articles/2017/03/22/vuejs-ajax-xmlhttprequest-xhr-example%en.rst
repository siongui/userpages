[Vue.js] AJAX XMLHttpRequest (XHR) Example
##########################################

:date: 2017-03-22T16:30+08:00
:tags: Vue.js, JavaScript, XMLHttpRequest, CORS, HTTP GET
:category: Vue.js
:summary: XMLHttpRequest (XHR) example by Vue.js.
:og_image: https://avatars1.githubusercontent.com/u/6128107?v=3&s=200
:adsu: yes


Demo
++++

Example for XMLHttpRequest_ (XHR) by Vue.js, which helps you transfer data
between a client and a server. This is usually used to retrieve data from a URL,
and update only part of the webpage without full refresh.

The following demo make HTTP GET XHR request to the given URL.

.. raw:: html

  <div id="vueapp" style="background-color: Azure; padding: 5px;">
    URL: <input type="text" v-model.trim="url" v-on:keyup.enter="xhr">
    <button v-on:click="xhr">XHR GET</button><br>
    <pre>{{ info }}</pre>
  </div>

  <script src="https://unpkg.com/vue@2.2.4/dist/vue.js"></script>

.. raw:: html

  <script>
  'use strict';

  new Vue({
    el: '#vueapp',
    data: {
      url: "https://siongui.github.io/xemaauj9k5qn34x88m4h/sukhada.json",
      info: ""
    },
    methods: {
      xhr: function() {
        this.info = "Requesting ...";
        var rq = new XMLHttpRequest();

        rq.onreadystatechange = function(vm) {
          if (this.readyState === XMLHttpRequest.DONE) {
            if (this.status === 200) {
              vm.info = this.responseText;
            } else {
              vm.info = "Request Failed";
            }
          }
        }.bind(rq, this);

        rq.open("GET", this.url);
        rq.send();
      }
    }
  })
  </script>

This demo can also make cross-domain requests if the server supports CORS, which
has the ``Access-Control-Allow-Origin`` header set properly. You can try the
following cross-domain URL in the demo:

::

  https://golden-operator-130720.appspot.com/sukhada.json

.. adsu:: 2


Source Code and Explanation
+++++++++++++++++++++++++++

**HTML**:

.. code-block:: html

  <div id="vueapp">
    URL: <input type="text" v-model.trim="url" v-on:keyup.enter="xhr">
    <button v-on:click="xhr">XHR</button><br>
    <pre>{{ info }}</pre>
  </div>

  <script src="https://unpkg.com/vue@2.2.4/dist/vue.js"></script>

If the user press the button or enter key in the input element, use the value of
the input element as URL and make HTTP GET request by *xhr* method.

**JavaScript**:

.. code-block:: javascript

  'use strict';

  new Vue({
    el: '#vueapp',
    data: {
      url: "https://siongui.github.io/xemaauj9k5qn34x88m4h/sukhada.json",
      info: ""
    },
    methods: {
      xhr: function() {
        this.info = "Requesting ...";
        var rq = new XMLHttpRequest();

        rq.onreadystatechange = function(vm) {
          if (this.readyState === XMLHttpRequest.DONE) {
            if (this.status === 200) {
              vm.info = this.responseText;
            } else {
              vm.info = "Request Failed";
            }
          }
        }.bind(rq, this);

        rq.open("GET", this.url);
        rq.send();
      }
    }
  })

The implementation of *xhr* method is the same as the code in other tutorials,
except that we use bind_ method to assign the value of *this* and pass the
`Vue instance`_ to the *onreadystatechange* event handler.

.. adsu:: 3

.. note::

   If you web page is served via HTTPS, the server that returns data also needs
   to serve via HTTPS. Otherwise browsers will block the request and make the
   request fail.

----

Tested on:

- ``Chromium Version 56.0.2924.76 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``
- ``Vue.js 2.2.4``

----

References:

.. [1] | `XMLHttpRequest - Google search <https://www.google.com/search?q=XMLHttpRequest>`_
       | `XMLHttpRequest - DuckDuckGo search <https://duckduckgo.com/?q=XMLHttpRequest>`_
       | `XMLHttpRequest - Ecosia search <https://www.ecosia.org/search?q=XMLHttpRequest>`_
       | `XMLHttpRequest - Qwant search <https://www.qwant.com/?q=XMLHttpRequest>`_
       | `XMLHttpRequest - Bing search <https://www.bing.com/search?q=XMLHttpRequest>`_
       | `XMLHttpRequest - Yahoo search <https://search.yahoo.com/search?p=XMLHttpRequest>`_
       | `XMLHttpRequest - Baidu search <https://www.baidu.com/s?wd=XMLHttpRequest>`_
       | `XMLHttpRequest - Yandex search <https://www.yandex.com/search/?text=XMLHttpRequest>`_

.. _Vue.js: https://vuejs.org/
.. _XMLHttpRequest: https://www.google.com/search?q=XMLHttpRequest
.. _XHR: https://www.google.com/search?q=XHR
.. _bind: https://www.google.com/search?q=javascript+bind
.. _Vue instance: https://vuejs.org/guide/instance.html

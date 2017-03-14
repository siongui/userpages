[Vue.js] Pretty Print JSON String
#################################

:date: 2017-03-15T02:55+08:00
:tags: Vue.js, JavaScript, JSON, pretty print
:category: Vue.js
:summary: Pretty print JSON_ string via Vue.js_.
:og_image: https://avatars1.githubusercontent.com/u/6128107?v=3&s=200
:adsu: yes


Demo
++++

Paste any valid JSON_ string in the following textarea. You will see the result
of pretty print on the fly.

.. raw:: html

  <div id="vueapp" style="background-color: Azure; padding: 5px;">
    <textarea v-model="jsonstr" rows="8" cols="40"></textarea>
    <pre>{{ jsonstr | pretty }}</pre>
  </div>

  <script src="https://unpkg.com/vue@2.2.4/dist/vue.js"></script>

.. raw:: html

  <script>
  'use strict';

  new Vue({
    el: '#vueapp',
    data: {
      jsonstr: '{"id":1,"name":"A green door","price":12.50,"tags":["home","green"]}'
    },
    filters: {
      pretty: function(value) {
        return JSON.stringify(JSON.parse(value), null, 2);
      }
    }
  })
  </script>

.. adsu:: 2

Source Code and Explanation
+++++++++++++++++++++++++++

**HTML**:

.. code-block:: html

  <div id="vueapp">
    <textarea v-model="jsonstr" rows="8" cols="40"></textarea>
    <pre>{{ jsonstr | pretty }}</pre>
  </div>

  <script src="https://unpkg.com/vue@2.2.4/dist/vue.js"></script>

The idea is simple. We implement a Vue.js_ filter_ to pretty print the string
in the textarea. According to the answer in [2]_, the built-in JavaScript
*JSON.stringify* method can help us pretty print the JSON string, so in the
filter, *JSON.parse* the string first and then *JSON.stringify* it again.

**JavaScript**:

.. code-block:: javascript

  'use strict';

  new Vue({
    el: '#vueapp',
    data: {
      jsonstr: '{"id":1,"name":"A green door","price":12.50,"tags":["home","green"]}'
    },
    filters: {
      pretty: function(value) {
        return JSON.stringify(JSON.parse(value), null, 2);
      }
    }
  })

----

Tested on:

- ``Chromium Version 56.0.2924.76 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``
- ``Vue.js 2.2.4``

----

References:

.. [1] | `javascript pretty print - Google search <https://www.google.com/search?q=javascript+pretty+print>`_
       | `javascript pretty print - DuckDuckGo search <https://duckduckgo.com/?q=javascript+pretty+print>`_
       | `javascript pretty print - Ecosia search <https://www.ecosia.org/search?q=javascript+pretty+print>`_
       | `javascript pretty print - Qwant search <https://www.qwant.com/?q=javascript+pretty+print>`_
       | `javascript pretty print - Bing search <https://www.bing.com/search?q=javascript+pretty+print>`_
       | `javascript pretty print - Yahoo search <https://search.yahoo.com/search?p=javascript+pretty+print>`_
       | `javascript pretty print - Baidu search <https://www.baidu.com/s?wd=javascript+pretty+print>`_
       | `javascript pretty print - Yandex search <https://www.yandex.com/search/?text=javascript+pretty+print>`_
.. [2] `How can I pretty-print JSON using JavaScript? - Stack Overflow <http://stackoverflow.com/a/7220510>`_

.. _Vue.js: https://vuejs.org/
.. _JSON: https://www.google.com/search?q=JSON
.. _filter: https://vuejs.org/v2/guide/syntax.html#Filters

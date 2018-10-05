[Vue.js] 優美列印JSON字串
#########################

:date: 2018-10-05T23:44+08:00
:tags: Vue.js, JavaScript, JSON, 優美列印
:category: Vue.js
:summary: 用 *Vue.js* 優美列印 *JSON* 字串。
:og_image: https://avatars1.githubusercontent.com/u/6128107?v=3&s=200
:adsu: yes


展示(demo)
++++++++++

在下方的 textarea 貼上任何有效的 JSON_ 字串。你會立即看到優美列印的結果。

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

程式原始碼及解釋
++++++++++++++++

**HTML**:

.. code-block:: html

  <div id="vueapp">
    <textarea v-model="jsonstr" rows="8" cols="40"></textarea>
    <pre>{{ jsonstr | pretty }}</pre>
  </div>

  <script src="https://unpkg.com/vue@2.2.4/dist/vue.js"></script>

想法很簡單。我們實作一個 Vue.js_ filter_ 來優美列印 textarea 裡的字串。
根據 [2]_ 裡面的答案， JavaScript 內建的 *JSON.stringify*
方法可以幫我們優美列印 JSON 字串，所以在 filter 裡頭，先用 *JSON.parse* 該字串，
然後再 *JSON.stringify* 它。

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

.. adsu:: 3

----

測試環境：

- ``Chromium Version 69.0.3497.81 on Ubuntu 18.04 (64-bit)``
- ``Vue.js 2.2.4``

----

參考：

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

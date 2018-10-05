[Vue.js] 輸入元素文字改變事件
#############################

:date: 2018-09-29T23:20+08:00
:tags: Vue.js, JavaScript, Web開發
:category: Vue.js
:summary: *Vue.js* 如何監看輸入文字元素 (input text element)
          的改變事件(change event)。
:og_image: https://vuejs.org/images/logo.png
:adsu: yes

.. contents:: 內容

解釋
++++

首先我們把 `輸入文字元素`_ (input text element) 的值(value)透過 v-model_ 綁定
到 *userinput* 這個變數：

.. code-block:: html

  <input type="text" v-model="userinput">

接著我們想要監聽這個輸入文字元素的 `改變事件`_ 。根據
`Event Handling in Vue.js Guide`_ ，我使用 v-on_ 來監聽輸入元素的改變事件，
如下所示：

.. code-block:: html

  <form> id="app">
    <input type="text" v-model="userinput" v-on:change="changeHandler">
  </form>

|

.. code-block:: javascript

  var app = new Vue({
    el: '#app',
    data: {
      userinput: ''
    },
    methods: {
      changeHandler: function(event) {
        // change of userinput, do something
      }
    }
  })

.. adsu:: 2

但這個程式碼不能運作。 *changeHandler* 並沒有執行。接著我做了一些谷歌搜尋 [1]_
，我發現正確的方式是要使用 watchers_:

監聽輸入元素的改變事件(onchange)正確方式是：

.. code-block:: html

  <form> id="app">
    <input type="text" v-model="userinput">
  </form>

|

.. code-block:: javascript

  var app = new Vue({
    el: '#app',
    data: {
      userinput: ''
    },
    watch: {
      userinput: function(val, oldVal) {
        // change of userinput, do something
      }
    }
  })

.. adsu:: 3


示例及完整範例
++++++++++++++

我寫了一個簡單的示例來展示如何利用 Vue.js_ 監聽
`輸入文字元素`_ 的 `改變事件`_ ：

.. rubric:: `Demo <{filename}/code/vuejs/input-change-event/index.html>`_
   :class: align-center
.. `Demo on CodePen <http://codepen.io/anon/pen/OWZVRX>`__
.. show_github_file:: siongui userpages content/code/vuejs/input-change-event/index.html
.. adsu:: 4
.. show_github_file:: siongui userpages content/code/vuejs/input-change-event/app.js

----

Tested on:

- ``Chromium Version 55.0.2883.87 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``
- ``Vue.js 2.1.10``.

----

References:

.. [1] `vuejs input change event - Google search <https://www.google.com/search?q=vuejs+input+change+event>`_

       `vuejs input change event - DuckDuckGo search <https://duckduckgo.com/?q=vuejs+input+change+event>`_

       `vuejs input change event - Bing search <https://www.bing.com/search?q=vuejs+input+change+event>`_

       `vuejs input change event - Yahoo search <https://search.yahoo.com/search?p=vuejs+input+change+event>`_

       `vuejs input change event - Baidu search <https://www.baidu.com/s?wd=vuejs+input+change+event>`_

       `vuejs input change event - Yandex search <https://www.yandex.com/search/?text=vuejs+input+change+event>`_

.. [2] `javascript - Vuejs event on change of element value? - Stack Overflow <http://stackoverflow.com/questions/34723680/vuejs-event-on-change-of-element-value>`_
.. adsu:: 5
.. [3] `vuejs v-on change - Google search <https://www.google.com/search?q=vuejs+v-on+change>`_

       `vuejs v-on change - DuckDuckGo search <https://duckduckgo.com/?q=vuejs+v-on+change>`_

       `vuejs v-on change - Bing search <https://www.bing.com/search?q=vuejs+v-on+change>`_

       `vuejs v-on change - Yahoo search <https://search.yahoo.com/search?p=vuejs+v-on+change>`_

       `vuejs v-on change - Baidu search <https://www.baidu.com/s?wd=vuejs+v-on+change>`_

       `vuejs v-on change - Yandex search <https://www.yandex.com/search/?text=vuejs+v-on+change>`_

.. [4] `javascript - How to fire an event when v-model changes ? (vue js) - Stack Overflow <http://stackoverflow.com/a/33305093>`_

.. _Vue.js: https://vuejs.org/
.. _改變事件: https://www.google.com/search?q=input+text+change+event
.. _輸入文字元素: https://www.google.com/search?q=input+text+element
.. _v-model: https://vuejs.org/v2/api/#v-model
.. _Event Handling in Vue.js Guide: https://vuejs.org/v2/guide/events.html
.. _v-on: https://vuejs.org/v2/api/#v-on
.. _watchers: https://vuejs.org/v2/guide/computed.html#Watchers
.. _vm.$watch: https://vuejs.org/v2/api/#vm-watch

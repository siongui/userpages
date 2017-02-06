[Vue.js] Input Text Element Change Event
########################################

:date: 2017-02-03T06:56+08:00
:tags: Vue.js, JavaScript, DOM, html
:category: Vue.js
:summary: How to monitor the `change event`_ of `input text element`_
          in Vue.js_.
:og_image: https://vuejs.org/images/logo.png
:adsu: yes

.. contents::

Explanation
+++++++++++

First we bind the value of `input text element`_ to the variable *userinput*
via v-model_:

.. code-block:: html

  <input type="text" v-model="userinput">

Next, we want to listen to the `change event`_ of this input text element.
According to `Event Handling in Vue.js Guide`_, I use v-on_ to listen to the
change event of input event as follows:

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

But the code does not work. *changeHandler* does not run. Then I do some
googling [1]_, I found that the correct way is to use watchers_:

The correct way to listen to change event(onchange) of input element is:

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


Demo & Full Example
+++++++++++++++++++

I write a simple demo to show how to monitor the `change event`_ of
`input text element`_ in Vue.js_.

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
.. _change event: https://www.google.com/search?q=input+text+change+event
.. _input text element: https://www.google.com/search?q=input+text+element
.. _v-model: https://vuejs.org/v2/api/#v-model
.. _Event Handling in Vue.js Guide: https://vuejs.org/v2/guide/events.html
.. _v-on: https://vuejs.org/v2/api/#v-on
.. _watchers: https://vuejs.org/v2/guide/computed.html#Watchers
.. _vm.$watch: https://vuejs.org/v2/api/#vm-watch

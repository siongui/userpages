[Vue.js] HTML Select Element Example
####################################

:date: 2017-04-27T02:13+08:00
:tags: Vue.js, JavaScript, html, DOM, CSS Animation, Animate.css
:category: Vue.js
:summary: Use Vue.js to manipulate HTML *select* element. First show how to use
          *v-model* directive to create data bindings for *select* element. Then
          use the selected option to make some animation via *animate.css*.
:og_image: https://image.slidesharecdn.com/enjoythevue-161013154440/95/enjoy-the-vuejs-20-638.jpg
:adsu: yes

Use Vue.js_ to manipulate `HTML select`_ element. First show how to use v-model_
directive to create data bindings for *select* element. Then use the selected
option to make some animation via animate.css_.

The following demo shows how to use v-model_ to get selected option.

.. raw:: html

  <div id="vueapp" style="background-color: Azure; padding: 5px;">
  <br>

  <select v-model="action">
    <option value="bounce">bounce</option>
    <option value="flash">flash</option>
    <option value="pulse">pulse</option>
  </select>

  <pre>selected option: {{ action }}</pre>

  </div>

  <script src="https://unpkg.com/vue@2.2.6/dist/vue.js"></script>

.. raw:: html

  <script>
  'use strict';

  var app = new Vue({
    el: '#vueapp',
    data: {
      action: 'bounce'
    }
  });

  </script>

The following is the source code for above demo.

**HTML**:

.. code-block:: html

  <div id="vueapp">
    <select v-model="action">
      <option value="bounce">bounce</option>
      <option value="flash">flash</option>
      <option value="pulse">pulse</option>
    </select>

    <pre>selected option: {{ action }}</pre>
  </div>

  <script src="https://unpkg.com/vue@2.2.6/dist/vue.js"></script>

The user selection is saved in ``action`` variable.

**JavaScript**:

.. code-block:: javascript

  'use strict';

  var app = new Vue({
    el: '#vueapp',
    data: {
      action: 'bounce'
    }
  });

.. adsu:: 2

Now we give a more realistic example. Show you how to make animations based on
the selected options.

.. raw:: html

  <div id="vueapp2" style="background-color: Azure; padding: 5px;">
  <br>

  <select v-model="action">
    <option value="bounce">bounce</option>
    <option value="flash">flash</option>
    <option value="pulse">pulse</option>
  </select>

  <br><br><br>
  <div :class="'animated ' + action">Animate.css</div>

  </div>
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/animate.css/3.5.2/animate.css" />

.. raw:: html

  <script>
  'use strict';

  var app = new Vue({
    el: '#vueapp2',
    data: {
      action: 'bounce'
    }
  });

  </script>

To use *animate.css*, first insert the following line in the *head* section of
your HTML.

.. code-block:: html

  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/animate.css/3.5.2/animate.css" />

Replace the HTML code in the first demo as follows (JavaScript code unchanged):

**HTML**:

.. code-block:: html

  <div id="vueapp">
    <select v-model="action">
      <option value="bounce">bounce</option>
      <option value="flash">flash</option>
      <option value="pulse">pulse</option>
    </select>

    <div :class="'animated ' + action">Animate.css</div>
  </div>

To make animation via *animate.css*, we need to add two classes to the animated
element. One is ``animated``, and the other is the name of the animation. You
can get all the animation names in *animate.css* website. Here only three
animation names are used for demo purpose.

We use the `class and style bindings`_ provided by Vue.js to update the classes
of animated element, and hence achieve the animations without writing any
JavaScript code!

.. adsu:: 3

----

Tested on:

- ``Chromium Version 58.0.3029.81 Built on Ubuntu , running on Ubuntu 17.04 (64-bit)``
- ``Vue.js 2.2.6``

----

References:

.. [1] | `html select - Google search <https://www.google.com/search?q=html+select>`_
       | `html select - DuckDuckGo search <https://duckduckgo.com/?q=html+select>`_
       | `html select - Ecosia search <https://www.ecosia.org/search?q=html+select>`_
       | `html select - Qwant search <https://www.qwant.com/?q=html+select>`_
       | `html select - Bing search <https://www.bing.com/search?q=html+select>`_
       | `html select - Yahoo search <https://search.yahoo.com/search?p=html+select>`_
       | `html select - Baidu search <https://www.baidu.com/s?wd=html+select>`_
       | `html select - Yandex search <https://www.yandex.com/search/?text=html+select>`_

.. [2] `[Vue.js] Animate.css Test Demo <{filename}../../01/25/vuejs-animate.css-test-demo%en.rst>`_

.. _Vue.js: https://vuejs.org/
.. _HTML select: https://www.google.com/search?q=HTML+select
.. _v-model: https://vuejs.org/v2/guide/forms.html
.. _animate.css: https://daneden.github.io/animate.css/
.. _class and style bindings: https://vuejs.org/v2/guide/class-and-style.html

[Vue.js] Typing Text Effect
###########################

:date: 2017-03-09T03:50+08:00
:tags: Vue.js, JavaScript, DOM, html, Typing Text Effect
:category: Vue.js
:summary: Sequential and parallel typing text effect by Vue.js_.
:og_image: http://jquery-plugins.net/image/plugin/theaterjs-javascript-typing-effect-plugin.png
:adsu: yes


Demo
++++

.. raw:: html

  <div id="vueapp" style="background-color: Azure; padding: 5px;">
    <span data-typedtext="May I be happy. May I be free from stress & pain. May I be free from animosity. May I be free from oppression. May I be free from trouble. May I look after myself with ease."></span>
    <br><br>
    <span data-typedtext="願我得喜樂。願我離憂苦。願我不受敵意。願我不受壓迫。願我免遭困難。願我輕鬆照顧自己。"></span>
    <br><br>
    <span data-typedtext="ขอให้ข้าพเจ้าจงเป็นผู้ถึงสุข จงเป็นผู้ไร้ทุกข์ จงเป็นผู้ไม่มีเวร จงเป็นผู้ไม่เบียดเบียนซึ่งกันและกัน จงเป็นผู้ไม่มีทุกข์ จงรักษาตนอยู่เป็นสุขเถิด"></span>
  </div>

The three lines in above demo are being typed parallel.

.. raw:: html

  <div id="vueapp2" style="background-color: Azure; padding: 5px;">
    <span data-typedtext="May I be happy. May I be free from stress & pain. May I be free from animosity. May I be free from oppression. May I be free from trouble. May I look after myself with ease."></span>
    <br><br>
    <span data-typedtext="願我得喜樂。願我離憂苦。願我不受敵意。願我不受壓迫。願我免遭困難。願我輕鬆照顧自己。"></span>
    <br><br>
    <span data-typedtext="ขอให้ข้าพเจ้าจงเป็นผู้ถึงสุข จงเป็นผู้ไร้ทุกข์ จงเป็นผู้ไม่มีเวร จงเป็นผู้ไม่เบียดเบียนซึ่งกันและกัน จงเป็นผู้ไม่มีทุกข์ จงรักษาตนอยู่เป็นสุขเถิด"></span>
  </div>

The three lines in above demo are being typed sequentially.

.. adsu:: 2

.. raw:: html

  <script src="https://unpkg.com/vue@2.2.1/dist/vue.js"></script>
  <script>
  'use strict';

  new Vue({
    el: '#vueapp',
    data: {
      speed: 15
    },
    methods: {
      playing: function(elements, elementIndex, textIndex) {
        if (elementIndex == elements.length) {
          return;
        }
        var element = elements[elementIndex];
        if (textIndex == element.dataset.typedtext.length) {
          setTimeout(this.playing, this.speed, elements, elementIndex + 1, 0);
        } else {
          element.textContent += element.dataset.typedtext[textIndex];
          setTimeout(this.playing, this.speed, elements, elementIndex, textIndex + 1);
        }
      },
      playSequential: function(elements) {
        this.playing(elements, 0, 0);
      },
      playParallel: function(elements) {
        for (var i = 0; i < elements.length; ++i) {
          this.playing([elements[i]], 0, 0);
        }
      }
    },
    mounted: function() {
      var spans = this.$el.querySelectorAll("span[data-typedtext]");
      //this.playSequential(spans);
      this.playParallel(spans);
    }
  })
  </script>
  <script>
  'use strict';

  new Vue({
    el: '#vueapp2',
    data: {
      speed: 15
    },
    methods: {
      playing: function(elements, elementIndex, textIndex) {
        if (elementIndex == elements.length) {
          return;
        }
        var element = elements[elementIndex];
        if (textIndex == element.dataset.typedtext.length) {
          setTimeout(this.playing, this.speed, elements, elementIndex + 1, 0);
        } else {
          element.textContent += element.dataset.typedtext[textIndex];
          setTimeout(this.playing, this.speed, elements, elementIndex, textIndex + 1);
        }
      },
      playSequential: function(elements) {
        this.playing(elements, 0, 0);
      },
      playParallel: function(elements) {
        for (var i = 0; i < elements.length; ++i) {
          this.playing([elements[i]], 0, 0);
        }
      }
    },
    mounted: function() {
      var spans = this.$el.querySelectorAll("span[data-typedtext]");
      this.playSequential(spans);
      //this.playParallel(spans);
    }
  })
  </script>

Source Code
+++++++++++

**HTML**: Put the texts to be typed in *data-typedtext* attribute of *span*
element.

.. code-block:: html

  <div id="vueapp">
    <span data-typedtext="May I be happy. May I be free from stress & pain. May I be free from animosity. May I be free from oppression. May I be free from trouble. May I look after myself with ease."></span>
    <span data-typedtext="願我得喜樂。願我離憂苦。願我不受敵意。願我不受壓迫。願我免遭困難。願我輕鬆照顧自己。"></span>
    <span data-typedtext="ขอให้ข้าพเจ้าจงเป็นผู้ถึงสุข จงเป็นผู้ไร้ทุกข์ จงเป็นผู้ไม่มีเวร จงเป็นผู้ไม่เบียดเบียนซึ่งกันและกัน จงเป็นผู้ไม่มีทุกข์ จงรักษาตนอยู่เป็นสุขเถิด"></span>
  </div>

  <script src="https://unpkg.com/vue@2.2.1/dist/vue.js"></script>

**JavaScript**:

.. code-block:: javascript

  'use strict';

  new Vue({
    el: '#vueapp',
    data: {
      speed: 15
    },
    methods: {
      playing: function(elements, elementIndex, textIndex) {
        if (elementIndex == elements.length) {
          return;
        }
        var element = elements[elementIndex];
        if (textIndex == element.dataset.typedtext.length) {
          setTimeout(this.playing, this.speed, elements, elementIndex + 1, 0);
        } else {
          element.textContent += element.dataset.typedtext[textIndex];
          setTimeout(this.playing, this.speed, elements, elementIndex, textIndex + 1);
        }
      },
      playSequential: function(elements) {
        this.playing(elements, 0, 0);
      },
      playParallel: function(elements) {
        for (var i = 0; i < elements.length; ++i) {
          this.playing([elements[i]], 0, 0);
        }
      }
    },
    mounted: function() {
      var spans = this.$el.querySelectorAll("span[data-typedtext]");
      this.playSequential(spans);
      //this.playParallel(spans);
    }
  })

In mounted_ hook of `Vue instance`_ [1]_, use querySelectorAll_ to find all
*span* elements with *data-typedtext* attibute in `vm.$el`_. If you want the
texts being typed sequentially, call *this.playSequential* method. If you want
the texts being typed parallel, call *this.playParallel* method.

Another key point in the above code is that the *playing* method calls itself
repeatedly by setTimeout_, and in each function call, only one character or
letter is appended to the *textContent* of the element.

.. adsu:: 3

----

Tested on:

- ``Chromium Version 56.0.2924.76 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``
- ``Vue.js 2.2.1``

----

References:

.. [1] `Instance Lifecycle Hooks - The Vue Instance — Vue.js <https://vuejs.org/v2/guide/instance.html#Instance-Lifecycle-Hooks>`_


.. _setTimeout: https://www.w3schools.com/jsref/met_win_settimeout.asp
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _Vue.js: https://vuejs.org/
.. _ref: https://vuejs.org/v2/api/#ref
.. _Vue instance: https://vuejs.org/v2/guide/instance.html
.. _$refs: https://vuejs.org/v2/api/#vm-refs
.. _mounted: https://vuejs.org/v2/api/#mounted
.. _vm.$el: https://vuejs.org/v2/api/#vm-el
.. _querySelectorAll: https://www.google.com/search?q=querySelectorAll

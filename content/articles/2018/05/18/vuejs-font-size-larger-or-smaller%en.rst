[Vue.js] Font Size Larger or Smaller
####################################

:date: 2018-05-18T23:31+08:00
:tags: Vue.js, JavaScript, CSS, html
:category: Vue.js
:summary: Increase or decrease font size with buttons via Vue.js.
:og_image: https://s3.amazonaws.com/coursetro/posts/136-full.png
:adsu: yes


See demo first. Click the buttons and you will see font size becomes
larger/smaller.

.. raw:: html

  <div id="vueapp" style="margin: 2rem;">

  <div id="vueapp">
    <button v-on:click="fontSize += 0.25">A+</button>
    <button v-on:click="fontSize < 0.5? fontSize = 0.25: fontSize -= 0.25">A-</button>

    <br><br>
    <div v-bind:style="{ fontSize: fontSize + 'rem' }">
      Main content of website
    </div>
  </div>

  </div>

  <script src="https://cdn.jsdelivr.net/npm/vue@2.5.16/dist/vue.js"></script>

.. raw:: html

  <script>
  'use strict';

  new Vue({
    el: '#vueapp',
    data: {
      fontSize: 1.25
    }
  });
  </script>


The following is source code for the demo.

**HTML**:

.. code-block:: html

  <div id="vueapp">
    <button v-on:click="fontSize += 0.25">A+</button>
    <button v-on:click="fontSize < 0.5? fontSize = 0.25: fontSize -= 0.25">A-</button>

    <br><br>
    <div v-bind:style="{ fontSize: fontSize + 'rem' }">
      Main content of website
    </div>
  </div>

  <script src="https://cdn.jsdelivr.net/npm/vue@2.5.16/dist/vue.js"></script>

We use *fontSize* to control the font size inside the *div* which contains the
main content. When users click *A+* button, *fontSize* will increase 0.25. If
users click *A-* button, we will check if *fontSize* is smaller than 0.5. If
yes, the keep the *fontSize* at 0.25. If no, decrease the *fontSize* by 0.25.

**JavaScript**:

.. code-block:: javascript

  'use strict';

  new Vue({
    el: '#vueapp',
    data: {
      fontSize: 1.25
    }
  });

.. adsu:: 2

----

Tested on:

- ``Chromium 66.0.3359.139 on Ubuntu 18.04 (64-bit)``
- ``Vue.js 2.5.16``

----

References:

.. [1] `[JavaScript] Font Size Larger/Smaller <{filename}/articles/2016/02/27/javascript-font-size-larger-smaller%en.rst>`_

.. _Vue.js: https://vuejs.org/

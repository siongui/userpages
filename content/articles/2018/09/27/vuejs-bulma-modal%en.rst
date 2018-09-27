[Vue.js] Bulma Modal
####################

:date: 2018-09-27T22:14+08:00
:tags: Vue.js, JavaScript, html, DOM, CSS, Modal (Popup), Bulma, toggle,
       toggleable
:category: Vue.js
:summary: Toggle *Bulma* modal via *Vue.js*.
:og_image: https://www.designil.com/wp-content/uploads/2016/10/modal-box-bulma.png
:adsu: yes


Toggle `Bulma modal`_ via Vue.js_. Try demo first:

|

.. raw:: html

  <!-- you can put the following line in your html head -->
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.7.1/css/bulma.css" integrity="sha256-zKA1Bf41O96+gJSlkn/Bh2HATW/OhwkApPlYTp3B5O8=" crossorigin="anonymous" />

  <div id="vueapp">

  <!-- Button trigger modal -->
  <button class="button is-primary is-large modal-button" v-on:click="isShowModal = true">
    Launch card modal
  </button>

  <div class="modal" v-bind:class="{ 'is-active': isShowModal }">
    <div class="modal-background" v-on:click="isShowModal = false"></div>
    <div class="modal-card">
      <header class="modal-card-head">
        <p class="modal-card-title">Modal title</p>
        <button class="delete" aria-label="close" v-on:click="isShowModal = false"></button>
      </header>
      <section class="modal-card-body">
        Modal Content
      </section>
      <footer class="modal-card-foot">
        <button class="button" v-on:click="isShowModal = false">Cancel</button>
      </footer>
    </div>
  </div>

  </div>

  <!-- you can put the following line at the end of html body -->
  <script src="https://cdn.jsdelivr.net/npm/vue@2.5.17/dist/vue.js"></script>


.. raw:: html

  <script>
  'use strict';

  var app = new Vue({
    el: '#vueapp',
    data: {
      isShowModal: false
    }
  });
  </script>

|

To toggle Bulma modal, add the *is-active* class to the class attribute of the
modal to activate the modal. Remove the *is-active* class will make the modal
invisible.

The following is the source code for the above demo:

**HTML**:

.. code-block:: html

  <!-- you can put the following line in your html head -->
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.7.1/css/bulma.css" integrity="sha256-zKA1Bf41O96+gJSlkn/Bh2HATW/OhwkApPlYTp3B5O8=" crossorigin="anonymous" />

  <div id="vueapp">

  <!-- Button trigger modal -->
  <button class="button is-primary is-large modal-button" v-on:click="isShowModal = true">
    Launch card modal
  </button>

  <div class="modal" v-bind:class="{ 'is-active': isShowModal }">
    <div class="modal-background" v-on:click="isShowModal = false"></div>
    <div class="modal-card">
      <header class="modal-card-head">
        <p class="modal-card-title">Modal title</p>
        <button class="delete" aria-label="close" v-on:click="isShowModal = false"></button>
      </header>
      <section class="modal-card-body">
        Modal Content
      </section>
      <footer class="modal-card-foot">
        <button class="button" v-on:click="isShowModal = false">Cancel</button>
      </footer>
    </div>
  </div>

  </div>

  <!-- you can put the following line at the end of html body -->
  <script src="https://cdn.jsdelivr.net/npm/vue@2.5.17/dist/vue.js"></script>

The truthiness of variable *isShowModal* will determine whether *is-active*
class is present in the class attribute of the modal, which toggles the
visiblity of the modal.

.. adsu:: 2

**JavaScript**:

.. code-block:: javascript

  'use strict';

  var app = new Vue({
    el: '#vueapp',
    data: {
      isShowModal: false
    }
  });

----

**You might be interested in ...**

- `JavaScript for Bulma Modal <{filename}/articles/2018/01/17/bulma-modal-with-javascript%en.rst>`_
- `Bulma Modal with Pure CSS Toggle <{filename}/articles/2018/01/27/css-only-toggle-bulma-modal%en.rst>`_
- `Bulma Modal with Go Toggle <{filename}/articles/2017/12/04/bulma-modal-with-go-toggle%en.rst>`_

----

Tested on:

- ``Bulma 0.7.1``
- ``Vue.js 2.5.17``
- ``Chromium 69.0.3497.81 on Ubuntu 18.04 (64-bit)``

----

.. adsu:: 3

References:

.. [1] `Modal | Bulma: a modern CSS framework based on Flexbox <https://bulma.io/documentation/components/modal/>`_
.. [2] `Class and Style Bindings — Vue.js <https://vuejs.org/v2/guide/class-and-style.html>`_

.. _Vue.js: https://vuejs.org/
.. _Bulma modal: https://bulma.io/documentation/components/modal

[Vue.js] Bootstrap Modal
########################

:date: 2018-09-26T23:38+08:00
:tags: Vue.js, JavaScript, html, DOM, CSS, Modal (Popup), Bootstrap, toggle,
       toggleable
:category: Vue.js
:summary: Toggle *Bootstrap* modal via *Vue.js*.
:og_image: http://junerockwell.com/wp-content/uploads/Screen-Shot-2017-03-03-at-9.24.48-PM.png
:adsu: yes


Toggle `Bootstrap modal`_ via Vue.js_. Try demo first:

.. raw:: html

  <!-- you can put the following line in your html head -->
  <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">

  <div id="vueapp">

  <!-- Button trigger modal -->
  <button type="button" class="btn btn-primary" data-toggle="modal" v-on:click="isShowModal = 'block'">
    Launch demo modal
  </button>

  <!-- Modal -->
  <div class="modal fade show" v-bind:style="{ display: isShowModal }" tabindex="-1" role="dialog" aria-labelledby="exampleModalLabel" aria-hidden="true">
    <div class="modal-dialog" role="document">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title" id="exampleModalLabel">Modal title</h5>
          <button type="button" class="close" data-dismiss="modal" aria-label="Close">
            <span aria-hidden="true" v-on:click="isShowModal = 'none'">&times;</span>
          </button>
        </div>
        <div class="modal-body">
          ...
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-secondary"  v-on:click="isShowModal = 'none'">Close</button>
        </div>
      </div>
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
      isShowModal: 'none'
    }
  });
  </script>

|

To toggle Bootstrap modal, the class attribute of the modal is set to
*"modal fade show"*. The style of the modal is set to *display: none* in the
beginning to make the modal invisible, and set to *display: block* to make the
modal visible.

The following is the source code for the above demo:

**HTML**:

.. code-block:: html

  <!-- you can put the following line in your html head -->
  <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">

  <div id="vueapp">

  <!-- Button trigger modal -->
  <button type="button" class="btn btn-primary" data-toggle="modal" v-on:click="isShowModal = 'block'">
    Launch demo modal
  </button>

  <!-- Modal -->
  <div class="modal fade show" v-bind:style="{ display: isShowModal }" tabindex="-1" role="dialog" aria-labelledby="exampleModalLabel" aria-hidden="true">
    <div class="modal-dialog" role="document">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title" id="exampleModalLabel">Modal title</h5>
          <button type="button" class="close" data-dismiss="modal" aria-label="Close">
            <span aria-hidden="true" v-on:click="isShowModal = 'none'">&times;</span>
          </button>
        </div>
        <div class="modal-body">
          ...
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-secondary"  v-on:click="isShowModal = 'none'">Close</button>
        </div>
      </div>
    </div>
  </div>

  </div>

  <!-- you can put the following line at the end of html body -->
  <script src="https://cdn.jsdelivr.net/npm/vue@2.5.17/dist/vue.js"></script>

|

The variable *isShowModal* stores the value of **display** style of the modal.
*isShowModal* is set to *none* to make the modal invisible, and set to *block*
to make the modal visible.

.. adsu:: 2

**JavaScript**:

.. code-block:: javascript

  'use strict';

  var app = new Vue({
    el: '#vueapp',
    data: {
      isShowModal: 'none'
    }
  });

----

Tested on:

- ``Chromium 69.0.3497.81 on Ubuntu 18.04 (64-bit)``
- ``Bootstrap 4.1.3``
- ``Vue.js 2.5.17``

----

.. adsu:: 3

References:

.. [1] | `vue.js bootstrap modal - Google search <https://www.google.com/search?q=vue.js+bootstrap+modal>`_
       | `vue.js bootstrap modal - DuckDuckGo search <https://duckduckgo.com/?q=vue.js+bootstrap+modal>`_
       | `vue.js bootstrap modal - Ecosia search <https://www.ecosia.org/search?q=vue.js+bootstrap+modal>`_
       | `vue.js bootstrap modal - Qwant search <https://www.qwant.com/?q=vue.js+bootstrap+modal>`_
       | `vue.js bootstrap modal - Bing search <https://www.bing.com/search?q=vue.js+bootstrap+modal>`_
       | `vue.js bootstrap modal - Yahoo search <https://search.yahoo.com/search?p=vue.js+bootstrap+modal>`_
       | `vue.js bootstrap modal - Baidu search <https://www.baidu.com/s?wd=vue.js+bootstrap+modal>`_
       | `vue.js bootstrap modal - Yandex search <https://www.yandex.com/search/?text=vue.js+bootstrap+modal>`_
.. [2] `Modal · Bootstrap <https://getbootstrap.com/docs/4.1/components/modal/>`_
.. [3] `Class and Style Bindings — Vue.js <https://vuejs.org/v2/guide/class-and-style.html>`_

.. _Vue.js: https://vuejs.org/
.. _Bootstrap modal: https://getbootstrap.com/docs/4.1/components/modal/

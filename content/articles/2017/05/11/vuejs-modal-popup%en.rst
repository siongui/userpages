[Vue.js] Modal (Popup)
######################

:date: 2017-05-11T21:49+08:00
:tags: Vue.js, JavaScript, html, DOM, CSS, Modal (Popup)
:category: Vue.js
:summary: Modal (Popup) implementation via Vue.js and CSS. Modal is dialog
          box/popup window that is displayed on top of the current page.
:og_image: https://dab1nmslvvntp.cloudfront.net/wp-content/uploads/2014/02/1392433498bootstrap-modal-basic.jpg
:adsu: yes

Modal (Popup) implementation via Vue.js and CSS. `According to w3schools`_,
modal is dialog box/popup window that is displayed on top of the current page.
The modal here is similar to `Bootstrap modal`_.

.. raw:: html

  <style>
  .modal {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translateX(-50%) translateY(-50%);
    background: azure;
    padding: 1rem;
    border-radius: 1rem;
    border: 1px solid gray;
  }
  </style>

  <div id="vueapp">

  <!-- Button trigger modal -->
  <button type="button" v-on:click="isShowMyModal = true">
    Launch demo modal
  </button>

  <!-- Modal -->
  <div class="modal" v-show="isShowMyModal">
    <h4>Modal Title</h4>
    <p>
       Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas turpis
       felis, tempus sit amet sollicitudin quis, aliquet ut felis.
       Nam a malesuada sem.
    </p>
    <button type="button" v-on:click="isShowMyModal = false">Close</button>
  </div>

  </div>

  <script src="https://unpkg.com/vue@2.3.3/dist/vue.js"></script>

.. raw:: html

  <script>
  'use strict';

  var app = new Vue({
    el: '#vueapp',
    data: {
      isShowMyModal: false
    }
  });
  </script>

The following is the source code for above demo.

**HTML**:

.. code-block:: html

  <div id="vueapp">

  <!-- Button trigger modal -->
  <button type="button" v-on:click="isShowMyModal = true">
    Launch demo modal
  </button>

  <!-- Modal -->
  <div class="modal" v-show="isShowMyModal">
    <h4>Modal Title</h4>
    <p>
       Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas turpis
       felis, tempus sit amet sollicitudin quis, aliquet ut felis.
       Nam a malesuada sem.
    </p>
    <button type="button" v-on:click="isShowMyModal = false">Close</button>
  </div>

  </div>

  <script src="https://unpkg.com/vue@2.3.3/dist/vue.js"></script>

We use the variable *isShowMyModal* and conditional rendering [2]_ to control
the visibility of the modal.

.. adsu:: 2

**CSS**:

.. code-block:: css

  .modal {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translateX(-50%) translateY(-50%);
    background: azure;
    padding: 1rem;
    border-radius: 1rem;
    border: 1px solid gray;
  }

The CSS here is to center the modal horizontally and vertically [5]_.

**JavaScript**:

.. code-block:: javascript

  'use strict';

  var app = new Vue({
    el: '#vueapp',
    data: {
      isShowMyModal: false
    }
  });

Initially set *isShowMyModal* as *false* to make modal invisible in the
beginning.

----

Tested on:

- ``Chromium Version 58.0.3029.96 Built on Ubuntu , running on Ubuntu 17.04 (64-bit)``
- ``Vue.js 2.3.3``

----

.. adsu:: 3

References:

.. [1] | `CSS modal popup - Google search <https://www.google.com/search?q=CSS+modal+popup>`_
       | `CSS modal popup - DuckDuckGo search <https://duckduckgo.com/?q=CSS+modal+popup>`_
       | `CSS modal popup - Ecosia search <https://www.ecosia.org/search?q=CSS+modal+popup>`_
       | `CSS modal popup - Qwant search <https://www.qwant.com/?q=CSS+modal+popup>`_
       | `CSS modal popup - Bing search <https://www.bing.com/search?q=CSS+modal+popup>`_
       | `CSS modal popup - Yahoo search <https://search.yahoo.com/search?p=CSS+modal+popup>`_
       | `CSS modal popup - Baidu search <https://www.baidu.com/s?wd=CSS+modal+popup>`_
       | `CSS modal popup - Yandex search <https://www.yandex.com/search/?text=CSS+modal+popup>`_

.. [2] `Conditional Rendering — Vue.js <https://vuejs.org/v2/guide/conditional.html>`_
.. [3] `Event Handling — Vue.js <https://vuejs.org/v2/guide/events.html>`_

.. [4] | `center div horizontally and vertically - Google search <https://www.google.com/search?q=center+div+horizontally+and+vertically>`_
       | `center div horizontally and vertically - DuckDuckGo search <https://duckduckgo.com/?q=center+div+horizontally+and+vertically>`_
       | `center div horizontally and vertically - Ecosia search <https://www.ecosia.org/search?q=center+div+horizontally+and+vertically>`_
       | `center div horizontally and vertically - Qwant search <https://www.qwant.com/?q=center+div+horizontally+and+vertically>`_
       | `center div horizontally and vertically - Bing search <https://www.bing.com/search?q=center+div+horizontally+and+vertically>`_
       | `center div horizontally and vertically - Yahoo search <https://search.yahoo.com/search?p=center+div+horizontally+and+vertically>`_
       | `center div horizontally and vertically - Baidu search <https://www.baidu.com/s?wd=center+div+horizontally+and+vertically>`_
       | `center div horizontally and vertically - Yandex search <https://www.yandex.com/search/?text=center+div+horizontally+and+vertically>`_

.. [5] `html - How to center an element horizontally and vertically? - Stack Overflow <http://stackoverflow.com/questions/19461521/how-to-center-an-element-horizontally-and-vertically>`_

.. _Vue.js: https://vuejs.org/
.. _According to w3schools: https://www.w3schools.com/bootstrap/bootstrap_modal.asp
.. _Bootstrap modal: http://getbootstrap.com/javascript/#modals

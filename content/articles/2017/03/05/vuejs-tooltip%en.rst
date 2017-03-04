[Vue.js] Tooltip
################

:date: 2017-03-05T01:49+08:00
:tags: Vue.js, Tooltip, JavaScript, mouseenter, mouseleave, element offset, DOM,
       element position
:category: Vue.js
:summary: Simple tooltip_ implementation via Vue.js_.
:og_image: https://d30y9cdsu7xlg0.cloudfront.net/png/383467-200.png
:adsu: yes


See demo first. Move the cursor (mouse pointer) to hover over the blue text with
underline:

.. raw:: html

  <style>
  span[data-descr] {
      text-decoration: underline;
      color: blue;
      cursor: help;
  }
  .invisible {
      display: none;
  }
  .tooltip {
      position: absolute;
      background-color: yellow;
      border: 1px gray solid;
      border-radius: 3px;
      padding: 3px;
  }
  </style>


  <div style="text-align: center; padding: 2em;">

  <div id="vueapp">
    <div class="tooltip invisible" ref="tooltip"></div>
    This is a <span data-descr="hello, I am tooltip!">tooltip</span> example
    via <span data-descr="I am Vue.js!">Vue.js</span>.
  </div>

  </div>

  <script src="https://unpkg.com/vue@2.2.1/dist/vue.js"></script>
  <script>
  'use strict';

  new Vue({
    el: '#vueapp',
    mounted: function() {

      var tt = this.$refs.tooltip;

      function ShowTooltip(e) {
        var elm = e.target;
        tt.style.left = elm.offsetLeft + 'px';
        tt.style.top = (elm.offsetTop + elm.offsetHeight + 5) + 'px';
        tt.textContent = elm.dataset.descr;
        tt.classList.remove("invisible");
      }

      function HideTooltip(e) {
        tt.classList.add("invisible");
      }

      var spans = this.$el.querySelectorAll("span[data-descr]");
      for (var i = 0; i < spans.length; ++i) {
        var span = spans[i];
        span.addEventListener("mouseenter", ShowTooltip);
        span.addEventListener("mouseleave", HideTooltip);
      }

    }
  })
  </script>

The idea comes from the JavaScript_ implementation of my post [1]_. The
following is complete source code.

**HTML**:

.. code-block:: html

  <div id="vueapp">
    <div class="tooltip invisible" ref="tooltip"></div>
    This is a <span data-descr="hello, I am tooltip!">tooltip</span> example
    via <span data-descr="I am Vue.js!">Vue.js</span>.
  </div>

  <script src="https://unpkg.com/vue@2.2.1/dist/vue.js"></script>

The *div* element with *ref="tooltip"* [3]_ is the tooltip which shows the hint.
If the mouse hovers over the span element with *data-descr* attribute, the
tooltip will appear and the content of the tooltip is the texts in *data-descr*
attribute.

.. adsu:: 2

**CSS**:

.. code-block:: css

  span[data-descr] {
      text-decoration: underline;
      color: blue;
      cursor: help;
  }
  .invisible {
      display: none;
  }
  .tooltip {
      position: absolute;
      background-color: yellow;
      border: 1px gray solid;
      border-radius: 3px;
      padding: 3px;
  }

**JavaScript**:

.. code-block:: javascript

  'use strict';

  new Vue({
    el: '#vueapp',
    mounted: function() {

      var tt = this.$refs.tooltip;

      function ShowTooltip(e) {
        var elm = e.target;
        tt.style.left = elm.offsetLeft + 'px';
        tt.style.top = (elm.offsetTop + elm.offsetHeight + 5) + 'px';
        tt.textContent = elm.dataset.descr;
        tt.classList.remove("invisible");
      }

      function HideTooltip(e) {
        tt.classList.add("invisible");
      }

      var spans = this.$el.querySelectorAll("span[data-descr]");
      for (var i = 0; i < spans.length; ++i) {
        var span = spans[i];
        span.addEventListener("mouseenter", ShowTooltip);
        span.addEventListener("mouseleave", HideTooltip);
      }

    }
  })

.. adsu:: 3

In mounted_ hook of `Vue instance`_ [4]_, use querySelectorAll_ to find all
*span* elements with *data-descr* attibute in `vm.$el`_, and setup corresponding
mouseenter_/mouseleave_ event handlder to show/hide the tooltip.

----

Tested on:

- ``Chromium Version 56.0.2924.76 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``
- ``Vue.js 2.2.1``

----

References:

.. [1] `Pure CSS Tooltip and JavaScript Implementation <{filename}../04/css-only-tooltip-and-javascript-implementation%en.rst>`_
.. [2] `vuejs tooltip - Google search <https://www.google.com/search?q=vuejs+tooltip>`_
.. [3] | `Special Attributes - ref - Vue.js <https://vuejs.org/v2/api/#ref>`_
       | `Instance Properties - vm.$refs - Vue.js <https://vuejs.org/v2/api/#vm-refs>`_
.. adsu:: 4
.. [4] `Instance Lifecycle Hooks - The Vue Instance — Vue.js <https://vuejs.org/v2/guide/instance.html#Instance-Lifecycle-Hooks>`_

.. _tooltip: https://www.google.com/search?q=tooltip
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _Vue.js: https://vuejs.org/
.. _ref: https://vuejs.org/v2/api/#ref
.. _Vue instance: https://vuejs.org/v2/guide/instance.html
.. _$refs: https://vuejs.org/v2/api/#vm-refs
.. _mounted: https://vuejs.org/v2/api/#mounted
.. _vm.$el: https://vuejs.org/v2/api/#vm-el
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _querySelectorAll: https://www.google.com/search?q=querySelectorAll
.. _mouseenter: https://developer.mozilla.org/en/docs/Web/Events/mouseenter
.. _mouseleave: https://developer.mozilla.org/en/docs/Web/Events/mouseleave

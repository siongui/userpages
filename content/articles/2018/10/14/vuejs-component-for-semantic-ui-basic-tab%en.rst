Vue.js Component for Semantic UI Basic Tabs
###########################################

:date: 2018-10-14T22:28+08:00
:tags: Vue.js, JavaScript, html, CSS, Tab Panel, Semantic UI
:category: Vue.js
:summary: Reusable *Vue.js* component for *Semantic UI* basic tabs example.
:og_image: https://webdesignledger.com/wp-content/uploads/2015/04/08-semantic-ui-tabs-design.jpg
:adsu: yes


Reusable Vue.js_ component_ for `Semantic UI`_ basic tab_.
Click the following tabs to see the demo.

|

.. raw:: html

  <!-- you can put the following line in your html head -->
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/semantic-ui@2.4.1/dist/semantic.min.css">

  <div id="vueapp">
    <semantic-ui-basic-tab :tabpanes="mytabpanes"></semantic-ui-basic-tab>
  </div>

  <!-- you can put the following line at the end of html body -->
  <script src="https://cdn.jsdelivr.net/npm/vue@2.5.17/dist/vue.js"></script>

.. raw:: html

  <script>
  'use strict';

  Vue.component("semantic-ui-basic-tab", {
      template: `
  <div>
    <!-- Nav tabs -->
    <div class="ui top attached tabular menu">
      <a class="item"
         v-for="(tabpane, index) in tabpanes"
         v-bind:class="{ 'active': tabsel == index }"
         @click="tabsel = index">
        {{tabpane.nav}}
      </a>
    </div>
    <!-- Tab panes -->
    <div class="ui bottom attached tab segment"
         v-for="(tabpane, index) in tabpanes"
         v-bind:class="{ 'active': tabsel == index }">
       {{tabpane.pane}}
    </div>
  </div>
  `,
      props: ["tabpanes"],
      data: function() {
        return {
          tabsel: 0
        }
      }
  });

  var app = new Vue({
    el: '#vueapp',
    data: {
      mytabpanes: [
        {nav: "tab1", pane:"tab1 content"},
        {nav: "tab2", pane:"tab2 content"},
        {nav: "tab3", pane:"tab3 content"},
        {nav: "tab4", pane:"tab4 content"},
      ]
    }
  });
  </script>

|
|

The following is the source code for above demo.


**HTML**:

.. code-block:: html

  <!-- you can put the following line in your html head -->
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/semantic-ui@2.4.1/dist/semantic.min.css">

  <div id="vueapp">
    <semantic-ui-basic-tab :tabpanes="mytabpanes"></semantic-ui-basic-tab>
  </div>

  <!-- you can put the following line at the end of html body -->
  <script src="https://cdn.jsdelivr.net/npm/vue@2.5.17/dist/vue.js"></script>

|

We define a Vue.js component whose name is *semantic-ui-basic-tab*, and we pass
data of tabs and panes to it.

|

.. adsu:: 2

**JavaScript**:

.. code-block:: javascript

  'use strict';

  Vue.component("semantic-ui-basic-tab", {
      template: `
  <div>
    <!-- Nav tabs -->
    <div class="ui top attached tabular menu">
      <a class="item"
         v-for="(tabpane, index) in tabpanes"
         v-bind:class="{ 'active': tabsel == index }"
         @click="tabsel = index">
        {{tabpane.nav}}
      </a>
    </div>
    <!-- Tab panes -->
    <div class="ui bottom attached tab segment"
         v-for="(tabpane, index) in tabpanes"
         v-bind:class="{ 'active': tabsel == index }">
       {{tabpane.pane}}
    </div>
  </div>
  `,
      props: ["tabpanes"],
      data: function() {
        return {
          tabsel: 0
        }
      }
  });

  var app = new Vue({
    el: '#vueapp',
    data: {
      mytabpanes: [
        {nav: "tab1", pane:"tab1 content"},
        {nav: "tab2", pane:"tab2 content"},
        {nav: "tab3", pane:"tab3 content"},
        {nav: "tab4", pane:"tab4 content"},
      ]
    }
  });


The data passed to the component is an array consisting of nav and content of
the tabs. we use *v-for* to create the actual HTML of the tabs in the component.

We use the variable *tabsel* to indicate current selected tab.
When users click on the tab, update *tabsel* according to the tab index and
hence show the selected tab pane according to the value of *tabsel*.

The *active* class of tabs will also be updated according to *tabsel*.
Set *tabsel* to ``0`` in the initialization phase to make first tab as default
tab.

If you do not want to use Vue.js or any JavaScript code, see [2]_ for pure CSS
Semantic UI Basic tabs.

----

Tested on:

- ``Semantic UI 2.4.1``
- ``Vue.js 2.5.17``
- ``Chromium 69.0.3497.81 on Ubuntu 18.04 (64-bit)``

----

References:

.. [1] `Vue.js Component for Bootstrap Tab Panel <{filename}/articles/2018/10/12/vuejs-component-for-bootstrap-tab-panel%en.rst>`_
.. [2] `Pure CSS Semantic UI Basic Tabs <{filename}/articles/2018/09/30/css-only-semantic-ui-basic-tab%en.rst>`_
.. adsu:: 3
.. [3] `Components Basics — Vue.js <https://vuejs.org/v2/guide/components.html>`_
.. [4] `List Rendering — Vue.js <https://vuejs.org/v2/guide/list.html>`_

.. _Vue.js: https://vuejs.org/
.. _component: https://vuejs.org/v2/guide/components.html
.. _Semantic UI: https://semantic-ui.com/
.. _tab: https://semantic-ui.com/modules/tab.html

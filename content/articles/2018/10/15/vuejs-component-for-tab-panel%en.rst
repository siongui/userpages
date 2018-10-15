Vue.js Component for Tab Panel
##############################

:date: 2018-10-15T22:51+08:00
:tags: Vue.js, JavaScript, html, CSS, Tab Panel
:category: Vue.js
:summary: Implementation of reusable *Vue.js* component for tab panel.
:og_image: https://webdesignledger.com/wp-content/uploads/2015/04/08-semantic-ui-tabs-design.jpg
:adsu: yes


Tab panel implementation via reusable Vue.js_ component_ and CSS.
Click the following tabs to see the demo.

.. raw:: html

  <style>
  .nav-tabs {
    display: flex;
    flex-wrap: wrap;
    margin-bottom: 15px;
    border-bottom: 1px solid #ddd;
    list-style: none;
  }
  .nav-tabs > li {
    margin-bottom: 1px;
    margin-right: 1rem;
    line-height: 2rem;
  }
  .nav-tabs > li > a {
    cursor: pointer;
    text-decoration: none;
  }
  .tab-pane {
    padding: 1rem;
  }
  </style>

.. raw:: html

  <div id="vueapp" style="background-color: Azure; padding: 1rem;">
    <tab-panel :tabpanes="mytabpanes"></tab-panel>
  </div>

  <!-- you can put the following line at the end of html body -->
  <script src="https://cdn.jsdelivr.net/npm/vue@2.5.17/dist/vue.js"></script>

.. raw:: html

  <script>
  'use strict';

  Vue.component("tab-panel", {
      template: `
  <div>
    <!-- Nav tabs -->
    <ul class="nav-tabs">
      <li v-for="(tabpane, index) in tabpanes"
          @click="tabsel = index">
        <a>{{tabpane.nav}}</a>
      </li>
    </ul>
    <!-- Tab panes -->
    <div class="tab-content"
       <div class="tab-pane"
            v-for="(tabpane, index) in tabpanes"
            v-show="tabsel == index">
         {{tabpane.pane}}
       </div>
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

The following is the source code for above demo.


**HTML**:

.. code-block:: html

  <div id="vueapp">
    <tab-panel :tabpanes="mytabpanes"></tab-panel>
  </div>

  <!-- you can put the following line at the end of html body -->
  <script src="https://cdn.jsdelivr.net/npm/vue@2.5.17/dist/vue.js"></script>


We define a Vue.js component whose name is *tab-panel*, and we pass data of tabs
and panes to it.

.. adsu:: 2

**JavaScript**:

.. code-block:: javascript

  'use strict';

  Vue.component("tab-panel", {
      template: `
  <div>
    <!-- Nav tabs -->
    <ul class="nav-tabs">
      <li v-for="(tabpane, index) in tabpanes"
          @click="tabsel = index">
        <a>{{tabpane.nav}}</a>
      </li>
    </ul>
    <!-- Tab panes -->
    <div class="tab-content"
       <div class="tab-pane"
            v-for="(tabpane, index) in tabpanes"
            v-show="tabsel == index">
         {{tabpane.pane}}
       </div>
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

.. The *active* class of tabs will also be updated according to *tabsel*.

Set *tabsel* to ``0`` in the initialization phase to make first tab as default
tab.


**CSS**:

.. code-block:: css

  .nav-tabs {
    display: flex;
    flex-wrap: wrap;
    margin-bottom: 15px;
    border-bottom: 1px solid #ddd;
    list-style: none;
  }
  .nav-tabs > li {
    margin-bottom: 1px;
    margin-right: 1rem;
    line-height: 2rem;
  }
  .nav-tabs > li > a {
    cursor: pointer;
    text-decoration: none;
  }
  .tab-pane {
    padding: 1rem;
  }

Use flexbox to align the tabs in one row.

----

Tested on:

- ``Vue.js 2.5.17``
- ``Chromium 69.0.3497.81 on Ubuntu 18.04 (64-bit)``

----

References:

.. [1] `[Vue.js] Tab Panel <{filename}/articles/2017/05/16/vuejs-tab-panel%en.rst>`_
.. [2] `Vue.js Component for Semantic UI Basic Tabs <{filename}/articles/2018/10/14/vuejs-component-for-semantic-ui-basic-tab%en.rst>`_
.. adsu:: 3
.. [3] `Components Basics — Vue.js <https://vuejs.org/v2/guide/components.html>`_
.. [4] `List Rendering — Vue.js <https://vuejs.org/v2/guide/list.html>`_

.. _Vue.js: https://vuejs.org/
.. _component: https://vuejs.org/v2/guide/components.html
.. _Semantic UI: https://semantic-ui.com/
.. _tab: https://semantic-ui.com/modules/tab.html

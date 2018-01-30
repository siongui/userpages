Pure CSS Bulma Tabs
###################

:date: 2018-01-30T10:08+08:00
:tags: Pure CSS Toggle, toggle, toggleable, html, CSS, Tab Panel, Bulma
:category: CSS
:summary: CSS only *Bulma* tab panel implementation.
:og_image: https://s3-us-west-2.amazonaws.com/i.cdpn.io/139759.EaaZbP.25c822f2-0a7a-452e-bd7a-f8ef8ad4deba.png
:adsu: yes

CSS only Bulma_ tab_ panel implementation.
The tabs here is similar to `Bootstrap tab`_, but no JavaScript required.
Click the following tabs to see the demo.

.. raw:: html

  <style>
  a:hover { text-decoration: none; }
  input[type="radio"], .tab-pane { display: none; }

  /* the magic */
  #pic:checked ~ .tab-content > .content-pic,
  #music:checked ~ .tab-content > .content-music,
  #video:checked ~ .tab-content > .content-video,
  #doc:checked ~ .tab-content > .content-doc {
    display: block;
  }
  </style>

  <div style="margin: 2rem;">

  <!-- Nav tabs -->
  <input type="radio" id="pic" name="nav-tab">
  <input type="radio" id="music" name="nav-tab">
  <input type="radio" id="video" name="nav-tab">
  <input type="radio" id="doc" name="nav-tab">
  <div class="tabs">
    <ul>
      <li><label for="pic"><a>Pictures</a></label></li>
      <li><label for="music"><a>Music</a></label></li>
      <li><label for="video"><a>Videos</a></label></li>
      <li><label for="doc"><a>Documents</a></label></li>
    </ul>
  </div>

  <!-- Tab panes -->
  <div class="tab-content">
    <div class="tab-pane content-pic">My Pictures</div>
    <div class="tab-pane content-music">My Music</div>
    <div class="tab-pane content-video">My Videos</div>
    <div class="tab-pane content-doc">My Documents</div>
  </div>

  </div>
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.6.2/css/bulma.min.css">

The basic technique is the same as *pure CSS toggle HTML element* [3]_. We need:

1 - Visible HTML *label* elements, which is the navigation tabs.

2 - Invisible HTML *input* *radio* elements, referenced by the *label* elements.
The value of *name* attribute of these elements must be the same.

3 - Tab contents, invisible in the beginning.

When users click on the visible label elements, the corresponding input radio
element is checked. We use CSS3 **:checked** selector to select corresponding
tab content, and make the selected tab content visible.

The complete source code is as follows:

**HTML**:

.. code-block:: html

  <!-- Nav tabs -->
  <input type="radio" id="pic" name="nav-tab">
  <input type="radio" id="music" name="nav-tab">
  <input type="radio" id="video" name="nav-tab">
  <input type="radio" id="doc" name="nav-tab">
  <div class="tabs">
    <ul>
      <li><label for="pic"><a>Pictures</a></label></li>
      <li><label for="music"><a>Music</a></label></li>
      <li><label for="video"><a>Videos</a></label></li>
      <li><label for="doc"><a>Documents</a></label></li>
    </ul>
  </div>

  <!-- Tab panes -->
  <div class="tab-content">
    <div class="tab-pane content-pic">My Pictures</div>
    <div class="tab-pane content-music">My Music</div>
    <div class="tab-pane content-video">My Videos</div>
    <div class="tab-pane content-doc">My Documents</div>
  </div>

  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.6.2/css/bulma.min.css">

.. adsu:: 2

**CSS**:

.. code-block:: css

  input[type="radio"], .tab-pane { display: none; }

  /* the magic */
  #pic:checked ~ .tab-content > .content-pic,
  #music:checked ~ .tab-content > .content-music,
  #video:checked ~ .tab-content > .content-video,
  #doc:checked ~ .tab-content > .content-doc {
    display: block;
  }

- The first rule hide the tab content and input radio box.
- The magic is in second rule. We use **:checked** and general sibling selector
  (~) to make the user-selected tab content visible.

----

Tested on:

- ``Chromium 63.0.3239.132 on Ubuntu 17.10 (64-bit)``
- ``Bulma 0.6.2``

----

References:

.. adsu:: 3
.. [1] `Pure CSS Tab Panel <{filename}/articles/2017/05/21/css-only-tab-panel%en.rst>`_
.. [2] `[Vue.js] Bulma Tabs <{filename}/articles/2018/01/26/vuejs-bulma-tabs%en.rst>`_
.. [3] `Pure CSS Toggle (Show/Hide) HTML Element <{filename}/articles/2017/02/27/css-only-toggle-dom-element%en.rst>`_

.. _Bulma: https://bulma.io/
.. _tab: https://bulma.io/documentation/components/tabs/
.. _Bootstrap tab: https://getbootstrap.com/docs/3.3/javascript/#tabs

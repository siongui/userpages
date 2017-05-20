Pure CSS Tab Panel
##################

:date: 2017-05-21T07:40+08:00
:tags: Pure CSS Toggle, toggle, toggleable, html, CSS, Tab Panel
:category: CSS
:summary: CSS only tab panel implementation.
:og_image: https://www.xul.fr/images/tabpanel.jpg
:adsu: yes

CSS only tab panel implementation.
The tab panel here is similar to `Bootstrap tab`_, but no JavaScript required.
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
    cursor: pointer;
  }
  .tab-pane {
    padding: 1rem;
    display: none;
  }

  input[type="radio"] {
    display: none;
  }

  /* the magic */
  #tab-home:checked ~ .tab-content > .content-home,
  #tab-profile:checked ~ .tab-content > .content-profile {
    display: block;
  }
  </style>

  <div style="background-color: Azure; padding: 1rem;">

    <!-- Nav tabs -->
    <input type="radio" id="tab-home" name="nav-tab">
    <input type="radio" id="tab-profile" name="nav-tab">
    <ul class="nav-tabs">
      <li><label for="tab-home">Home</label></li>
      <li><label for="tab-profile">Profile</label></li>
    </ul>

    <!-- Tab panes -->
    <div class="tab-content">
      <div class="tab-pane content-home">My home</div>
      <div class="tab-pane content-profile">My profile</div>
    </div>

  </div>

The basic technique is the same as *pure CSS toggle HTML element* [7]_. We need:

- Visible HTML *label* elements, which is the navigation tabs.
- Invisible HTML *input* *radio* elements, referenced by the *label* elements.
  The value of *name* attribute of these elements must be the same.
- Tab contents, invisible in the beginning.

When users click on the visible label elements, the corresponding input radio
element is checked. We use CSS3 **:checked** selector to select corresponding
tab content, and make the selected tab content visible.

The complete source code is as follows:

**HTML**:

.. code-block:: html

  <!-- Nav tabs -->
  <input type="radio" id="tab-home" name="nav-tab">
  <input type="radio" id="tab-profile" name="nav-tab">
  <ul class="nav-tabs">
    <li><label for="tab-home">Home</label></li>
    <li><label for="tab-profile">Profile</label></li>
  </ul>

  <!-- Tab panes -->
  <div class="tab-content">
    <div class="tab-pane content-home">My home</div>
    <div class="tab-pane content-profile">My profile</div>
  </div>

.. adsu:: 2

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
    cursor: pointer;
  }
  .tab-pane {
    padding: 1rem;
    display: none;
  }

  input[type="radio"] {
    display: none;
  }

  /* the magic */
  #tab-home:checked ~ .tab-content > .content-home,
  #tab-profile:checked ~ .tab-content > .content-profile {
    display: block;
  }

- The first two rules use flexbox to align the tabs in one row.
- The third and fourth rules hide the tab content and input radio box.
- The magic is in last rule. We use **:checked** and general sibling selector
  (~) to make the user-selected tab content visible.

.. adsu:: 3

----

Tested on:

- ``Chromium Version 58.0.3029.110 Built on Ubuntu , running on Ubuntu 17.04 (64-bit)``

----

References:

.. [1] `有趣的CSS题目（8）：纯CSS的导航栏Tab切换方案 - WEB前端 - 伯乐在线 <http://web.jobbole.com/88478/>`_
.. [2] `有趣的CSS题目（10）：结构性伪类选择器 - WEB前端 - 伯乐在线 <http://web.jobbole.com/88981/>`_
.. [3] `巧用 CSS 的 :target 选择器，打造没有 JS 的 UI 效果 - WEB前端 - 伯乐在线 <http://web.jobbole.com/86796/>`_
.. [4] | `pure css tabs - Google search <https://www.google.com/search?q=pure+css+tabs>`_
       | `pure css tabs - DuckDuckGo search <https://duckduckgo.com/?q=pure+css+tabs>`_
       | `pure css tabs - Ecosia search <https://www.ecosia.org/search?q=pure+css+tabs>`_
       | `pure css tabs - Qwant search <https://www.qwant.com/?q=pure+css+tabs>`_
       | `pure css tabs - Bing search <https://www.bing.com/search?q=pure+css+tabs>`_
       | `pure css tabs - Yahoo search <https://search.yahoo.com/search?p=pure+css+tabs>`_
       | `pure css tabs - Baidu search <https://www.baidu.com/s?wd=pure+css+tabs>`_
       | `pure css tabs - Yandex search <https://www.yandex.com/search/?text=pure+css+tabs>`_
.. [5] `CSS3 :target Selector <https://www.w3schools.com/cssref/sel_target.asp>`_
.. [6] | `css target no jump - Google search <https://www.google.com/search?q=css+target+no+jump>`_
       | `css target no jump - DuckDuckGo search <https://duckduckgo.com/?q=css+target+no+jump>`_
       | `css target no jump - Ecosia search <https://www.ecosia.org/search?q=css+target+no+jump>`_
       | `css target no jump - Qwant search <https://www.qwant.com/?q=css+target+no+jump>`_
       | `css target no jump - Bing search <https://www.bing.com/search?q=css+target+no+jump>`_
       | `css target no jump - Yahoo search <https://search.yahoo.com/search?p=css+target+no+jump>`_
       | `css target no jump - Baidu search <https://www.baidu.com/s?wd=css+target+no+jump>`_
       | `css target no jump - Yandex search <https://www.yandex.com/search/?text=css+target+no+jump>`_
.. [7] `Pure CSS Toggle (Show/Hide) HTML Element <{filename}../../02/27/css-only-toggle-dom-element%en.rst>`_

.. _Bootstrap tab: http://getbootstrap.com/javascript/#tabs

Responsive Navigation Bar (Navbar)
##################################

:date: 2015-03-19 23:51
:tags: CSS, html, SASS, SCSS, toggle, toggleable, Responsive Web Design
:category: Web Development
:summary: Responsive top navigation bar (navbar) used in this website.
:adsu: yes


This post shows how to make a responsive top navbar like this website. Before
you start, please read [1]_ and [2]_ to know how to toogle dropdown menu with
CSS and how to create responsive grid layout. Both are used in our responsive
top navbar.

.. rubric:: `Demo <{filename}/code/css/responsive-navbar/layout.html>`_
      :class: align-center

Source Code for Demo (*HTML*):

.. show_github_file:: siongui userpages content/code/css/responsive-navbar/layout.html
.. adsu:: 2

Source Code for Demo (*SCSS*):

It is not a good idea to write CSS by hand in this case, so SCSS_ syntax is used
to make life easier.

.. show_github_file:: siongui userpages content/code/css/responsive-navbar/style.scss
.. adsu:: 3

Tested on:
  - ``Chromium Version 41.0.2272.76 Ubuntu 14.10 (64-bit)``
  - ``pyScss 1.3.4``

----

References
++++++++++

.. [1] `Toggle Element (Dropdown/Menu) Visibility with CSS <{filename}../../02/07/toogle-element-visibility-with-css%en.rst>`_

.. [2] `Simple Custom Responsive Grid Layout <{filename}../13/simple-custom-responsive-grid-layout%en.rst>`_

.. [3] `What’s the Deal With Display: Inline-Block? <http://designshack.net/articles/css/whats-the-deal-with-display-inline-block/>`_

.. [4] `Fighting the Space Between Inline Block Elements <https://css-tricks.com/fighting-the-space-between-inline-block-elements/>`_
.. adsu:: 4
.. [5] `vertical-align with bootstrap 3 <http://stackoverflow.com/questions/20547819/vertical-align-with-bootstrap-3>`_

.. [6] `How do you keep parents of floated elements from collapsing? <http://stackoverflow.com/questions/218760/how-do-you-keep-parents-of-floated-elements-from-collapsing>`_

.. [7] | `CSS 网格布局 - Google search <https://www.google.com/search?q=CSS+%E7%BD%91%E6%A0%BC%E5%B8%83%E5%B1%80>`_
       | `CSS 网格布局 - DuckDuckGo search <https://duckduckgo.com/?q=CSS+%E7%BD%91%E6%A0%BC%E5%B8%83%E5%B1%80>`_
       | `CSS 网格布局 - Ecosia search <https://www.ecosia.org/search?q=CSS+%E7%BD%91%E6%A0%BC%E5%B8%83%E5%B1%80>`_
       | `CSS 网格布局 - Bing search <https://www.bing.com/search?q=CSS+%E7%BD%91%E6%A0%BC%E5%B8%83%E5%B1%80>`_
       | `CSS 网格布局 - Yahoo search <https://search.yahoo.com/search?p=CSS+%E7%BD%91%E6%A0%BC%E5%B8%83%E5%B1%80>`_
       | `CSS 网格布局 - Baidu search <https://www.baidu.com/s?wd=CSS+%E7%BD%91%E6%A0%BC%E5%B8%83%E5%B1%80>`_
       | `CSS 网格布局 - Yandex search <https://www.yandex.com/search/?text=CSS+%E7%BD%91%E6%A0%BC%E5%B8%83%E5%B1%80>`_

.. _SCSS: http://sass-lang.com/

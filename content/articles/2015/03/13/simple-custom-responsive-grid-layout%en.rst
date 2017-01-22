Simple Custom Responsive Grid Layout
####################################

:date: 2015-03-13 23:23
:tags: CSS, html, SASS, SCSS
:category: Web Development
:summary: A simple custom responsive grid layout exmaple without CSS frameworks.
:adsu: yes


If simple design is adopted for your website, you do not need to use Bootstrap_
or `CSS frameworks`_ to make things complicated. This post shows how to write a
simple responsive layout by `CSS media query`_.

In the following demo, the layout consists of two columns (left column and right
column). If the screen width is larger than *767px*, left column will be *70%*
wide and right column will be *30%* wide. If the screen width is less than
*767px*, both left and right columns are 100% wide and left column will be on
top of right column.

.. rubric:: `Demo <{filename}/code/simple-custom-responsive-grid/layout.html>`_
      :class: align-center

Source Code for Demo (*HTML*):

.. show_github_file:: siongui userpages content/code/simple-custom-responsive-grid/layout.html

.. note::

  Add the following line in the *head* section of your HTML. Otherwise
  `CSS media query`_ will not work.

  .. code-block:: html

    <meta name="viewport" content="width=device-width, initial-scale=1">

.. note::

  Because ``display: inline-block`` is used for ``.left-column`` and
  ``.right-column``, there is whitespace issue [1]_ in such design.

  ``<!--`` and ``-->`` is used to prevent whitespace issue. Please refer to [2]_
  for details.

Source Code for Demo (*CSS*): Use `CSS media query`_

.. show_github_file:: siongui userpages content/code/simple-custom-responsive-grid/style.css

If you prefer to use SASS_ for your CSS writing, here is the SCSS equivalent of above CSS:

.. show_github_file:: siongui userpages content/code/simple-custom-responsive-grid/style.scss


Tested on: ``Chromium Version 41.0.2272.76 Ubuntu 14.10 (64-bit)``, ``pyScss 1.3.4``

----

References
++++++++++

.. [1] `What’s the Deal With Display: Inline-Block? <http://designshack.net/articles/css/whats-the-deal-with-display-inline-block/>`_

.. [2] `Fighting the Space Between Inline Block Elements <https://css-tricks.com/fighting-the-space-between-inline-block-elements/>`_

.. [3] `Should You Use Inline-Blocks As A Substitute For Floats? <http://www.vanseodesign.com/css/inline-blocks/>`_

.. [4] `Fluid + Responsive Grid Layouts in CSS <http://www.stephanboyer.com/post/41/fluid-responsive-grid-layouts-in-css>`_

.. [5] `Yet another CSS grid system <http://sans0r.github.io/yacgs/>`_

.. [6] `Simple pyScss Responsive Grid <https://github.com/siongui/scss-grid>`_

`CSS layout`_
~~~~~~~~~~~~~

.. [7] `A Complete Guide to Flexbox <http://css-tricks.com/snippets/css/a-guide-to-flexbox/>`_

.. [8] `vertical-align with bootstrap 3 <http://stackoverflow.com/questions/20547819/vertical-align-with-bootstrap-3>`_

.. [9] `How do you keep parents of floated elements from collapsing? <http://stackoverflow.com/questions/218760/how-do-you-keep-parents-of-floated-elements-from-collapsing>`_

.. [18] `Website Layout Tools Compared: Flexbox Vs. Susy – Smashing Magazine <https://www.smashingmagazine.com/2015/12/website-layout-tools-compared-flexbox-vs-susy/>`_

.. [19] `等高分栏布局小结 - WEB前端 - 伯乐在线 <http://web.jobbole.com/85031/>`_

.. [20] `圣杯布局小结 - WEB前端 - 伯乐在线 <http://web.jobbole.com/84993/>`

`scss media query`_
~~~~~~~~~~~~~~~~~~~

.. [10] `Responsive Web Design in Sass: Using media queries in Sass 3.2 <http://thesassway.com/intermediate/responsive-web-design-in-sass-using-media-queries-in-sass-32>`_

.. [11] `Write Better Media Queries with Sass <http://davidwalsh.name/write-media-queries-sass>`_

.. [12] `Approaches to Media Queries in Sass <http://css-tricks.com/approaches-media-queries-sass/>`_

`mobile responsive design`_
~~~~~~~~~~~~~~~~~~~~~~~~~~~

.. [13] `Bootstrap 3 Slide in Menu / Navbar on Mobile <http://stackoverflow.com/questions/20863288/bootstrap-3-slide-in-menu-navbar-on-mobile>`_

.. [14] `Bootstrap Tutorial – Creating a Responsive Navbar (Video) <http://bootstrapbay.com/blog/bootstrap-tutorial-navbar/>`_

.. [15] `How to Create Off Canvas Layouts with Susy <http://www.zell-weekeat.com/off-canvas-layouts-susy/>`_

.. [16] `Off The Beaten Canvas: Exploring The Potential Of The Off-Canvas Pattern <http://www.smashingmagazine.com/2014/02/24/off-the-beaten-canvas-exploring-the-potential-of-the-off-canvas-pattern/>`_

.. [17] `Implementing Off-Canvas Navigation For A Responsive Website <http://www.smashingmagazine.com/2013/01/15/off-canvas-navigation-for-responsive-website/>`_



.. _Bootstrap: http://getbootstrap.com/

.. _SASS: http://sass-lang.com/

.. _CSS media query: https://duckduckgo.com/?q=CSS+media+query

.. _CSS frameworks: https://duckduckgo.com/?q=CSS+frameworks

.. _CSS layout: https://duckduckgo.com/?q=CSS+layout

.. _scss media query: https://duckduckgo.com/?q=scss+media+query

.. _mobile responsive design: https://duckduckgo.com/?q=mobile+responsive+design

Pure CSS Responsive Nav Bar via Flexbox
#######################################

:date: 2017-04-20T22:25+08:00
:tags: CSS, Responsive Web Design, toggle, toggleable, html, Pure CSS Toggle
:category: CSS
:summary: CSS only responsive *nav bar* (*navigation bar*) via *flexbox*.
          Toggle mobile menu with CSS only. No JavaScript required.
          No CSS frameworks (Bootstrap, Bulma, etc.) required.
:og_image: https://www.w3schools.com/css/navbar_responsive_hor.jpg
:adsu: yes


CSS_ only responsive `nav bar`_ (*navigation bar*) via flexbox_.
Toggle mobile menu with CSS only.
No JavaScript_ required.
No `CSS frameworks`_ (Bootstrap_, Bulma_, etc.) required.

.. rubric:: Demo: Navbar of this webpage!
   :class: align-center

About two months ago, I made the pure CSS responsive navbar with Bulma
framework [1]_. Now I remove the dependency of Bulma framework and make a responsive
navbar with standard CSS flexbox [3]_.

I keep the name of the CSS classes the same as the name in Bulma to make Bulma
users feel familiar. To make the navbar responsive, you need to know the
technique of *CSS only toggle* [2]_. If not, please read references first.

For the CSS divider in navbar, please see [6]_.

**HTML**:

.. code-block:: html

  <nav class="nav">
    <div class="nav-left">
      <a class="nav-item" href="/">YOUR SITE NAME</a>
    </div>

    <label for="menu-toggle" class="nav-toggle">&equiv;</label>
    <input type="checkbox" id="menu-toggle" class="is-hidden"/>

    <div class="nav-right nav-menu">
      <a class="nav-item" href="/pages/about.html">About</a>
      <a class="nav-item" href="/archives.html">Archives</a>
      <a class="nav-item" href="/categories.html">Categories</a>
      <a class="nav-item" href="/tags.html">Tags</a>
      <a class="nav-item" href="/authors.html">Authors</a>

      <a class="nav-item" href="/zh/">中文</a>
      <a class="nav-item" href="/th/">ไทย</a>
    </div>
  </nav>

.. adsu:: 2

**SCSS**: You can `compile the SCSS code online`_.

.. code-block:: scss

  $mobile: 768px;
  $nav-height: 3.25rem;

  .is-hidden {
    display: none;
  }


  @media screen and (max-width: $mobile) {
    // if user touch nav-toggle, show nav-menu
    #menu-toggle:checked + .nav-menu {
      display: block;
    }
  }

  // ≡ is nav-toggle
  .nav-toggle {
    // hidden if not mobile
    @media (min-width: $mobile) {
      display: none;
    }

    &:hover {
      cursor: pointer;
    }

    font-size: $nav-height;

    align-items: center;
    display: flex;
    padding-right: 0.75rem;
  }

  .nav-item {
    align-items: center;
    display: flex;
    padding: 0.25rem 0.5rem;
  }

  .nav-left, .nav-right {
    align-items: stretch;
    display: flex;
    flex-wrap: wrap;
    flex-grow: 1;
  }

  .nav-left {
    justify-content: flex-start;
  }

  .nav-right {
    justify-content: flex-end;
  }

  // responsiveness
  .nav-menu {
    @media (max-width: $mobile) {
      &.nav-right {
        display: none;
        left: 0;
        right: 0;
        top: 100%;
        position: absolute;
        background-color: white;
        box-shadow: 0 4px 7px rgba(10, 10, 10, 0.1);
        .nav-item {
          border-top: 1px solid rgba(219, 219, 219, 0.5);
          padding: 0.75rem;
        }
      }
    }
  }

  .nav {
    display: flex;
    align-items: stretch;
    background-image: linear-gradient(to bottom left, white, #dddddd);
    min-height: $nav-height;

    // responsiveness
    position: relative;
  }


.. adsu:: 3

Tested on:
``Chromium Version 57.0.2987.98 Built on Ubuntu , running on Ubuntu 17.04 (64-bit)``

----

References
++++++++++

.. [1] `Pure CSS Bulma Responsive Nav Bar (Navigation Bar) <{filename}../../02/22/css-only-bulma-responsive-navbar%en.rst>`_
.. [2] `Pure CSS Toggle (Show/Hide) HTML Element <{filename}../../02/27/css-only-toggle-dom-element%en.rst>`_
.. [3] | `css flexbox - Google search <https://www.google.com/search?q=css+flexbox>`_
       | `A Complete Guide to Flexbox | CSS-Tricks <https://css-tricks.com/snippets/css/a-guide-to-flexbox/>`_
       | `Flexbox Cheatsheet | Hacker News <https://news.ycombinator.com/item?id=13877926>`__
       | `理解 Flexbox：你需要知道的一切 - WEB前端 - 伯乐在线 <http://web.jobbole.com/91035/>`_
       | `RAGrid: Powerful intrinsic pure CSS flexbox grid | Hacker News <https://news.ycombinator.com/item?id=14102320>`_
.. [4] `bulma/nav.sass at master · jgthms/bulma · GitHub <https://github.com/jgthms/bulma/blob/master/sass/components/nav.sass>`_
.. [5] | `TACHYONS - Css Toolkit <http://tachyons.io/>`_
       | `Mini.css: Minimal, responsive, style-agnostic CSS framework | Hacker News <https://news.ycombinator.com/item?id=14120796>`_
.. [6] `Responsive CSS Divider in Nav Bar <{filename}../21/responsive-css-separator-line-in-navbar%en.rst>`_
.. [7] | `Understanding Flexbox: Everything you need to know – freeCodeCamp <https://medium.freecodecamp.com/understanding-flexbox-everything-you-need-to-know-b4013d4dc9af>`_
       | `理解 Flexbox：你需要知道的一切 – 码农网 <http://www.codeceo.com/article/understanding-flexbox-everything-you-need-to-know.html>`_
       | `理解 Flexbox：你需要知道的一切 - WEB前端 - 伯乐在线 <http://web.jobbole.com/91035/>`_
       | `理解Flexbox：你需要知道的一切_Flexbox 教程_w3cplus <http://www.w3cplus.com/css3/understanding-flexbox-everything-you-need-to-know.html>`_
.. [8] `FlexBox Cheatsheet | Hacker News <https://news.ycombinator.com/item?id=14483429>`__

.. _CSS: https://www.google.com/search?q=CSS
.. _CSS frameworks: https://www.google.com/search?q=CSS+frameworks
.. _nav bar: https://www.google.com/search?q=navigation+bar
.. _flexbox: https://www.google.com/search?q=flexbox
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _Bulma: http://bulma.io/
.. _Bootstrap: http://getbootstrap.com/
.. _compile the SCSS code online: https://www.google.com/search?q=scss+to+css+online

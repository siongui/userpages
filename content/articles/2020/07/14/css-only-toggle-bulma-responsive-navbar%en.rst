Pure CSS Toggle Bulma Responsive Navbar
#######################################

:date: 2020-07-14T07:23+08:00
:tags: CSS, Responsive Web Design, toggle, toggleable, html, Pure CSS Toggle,
       Bulma, navbar
:category: CSS
:summary: CSS only responsive Bulma 0.9.0 navbar without JavaScript.
          Toggle mobile menu with CSS only. No JavaScript required.
:og_image: https://camo.githubusercontent.com/5f4c3abfdf781275ef90c4c7c3f5e7a73231a991/68747470733a2f2f692e696d6775722e636f6d2f6c6c61487163312e706e67
:adsu: yes


CSS_ only toggle responsive `Bulma navbar`_ without JavaScript_.

.. rubric:: `Demo 1 <{static}/code/css/bulma-responsive-navbar/index-bulma-0.9.0.html>`_
   :class: align-center

.. rubric:: `Demo 2 <{static}/code/css/bulma-responsive-navbar/index2-bulma-0.9.0.html>`_
   :class: align-center

Demo 1 is similar to the navbar style of Bulma 0.4.0., and Demo 2 is similar to
that of Bulma 0.9.0. The change of navbar style is achieved by the change of
placement of input_ checkbox element and slight change in CSS rules.

The technique is slightly modified from my previous post [1]_ and applied to the
latest Bulma 0.9.0 version. The difference from my previous post is that The
placement of invisible input_ checkbox element is changed so that the css
sibling and child selector can be used to toggle the menu and style of the
label_ element.

The key points of the technique [2]_:
  - A visible HTML label_ element (only visible on small screen < 1024px).
  - A invisible HTML input_ checkbox element, referenced by the label_ element.
  - The menu to be toggled
  - CSS rules to hide/show the menu according to whether the checkbox is
    checked.

When users click or touch on the visible label_ element, it will make the
invisible input_ checkbox checked/unchecked. And CSS rules will show the menu if
the checkbox is checked, or hide the menu if the checkbox is unchecked.

Full source code is as follows:

**Demo 1**:

.. show_github_file:: siongui userpages content/code/css/bulma-responsive-navbar/index-bulma-0.9.0.html

**Demo 2**:

.. show_github_file:: siongui userpages content/code/css/bulma-responsive-navbar/index2-bulma-0.9.0.html

----

Tested on: ``Chromium Version 83.0.4103.116 (Official Build) snap (64-bit)``, ``Ubuntu 20.04``, ``Bulma 0.9.0``

----

References
++++++++++

.. [1] `Pure CSS Bulma Responsive Nav Bar (Navigation Bar) <{filename}/articles/2017/02/22/css-only-bulma-responsive-navbar%en.rst>`_
.. [2] `CSS only menu toggle - no JavaScript required <http://www.outofscope.com/css-only-menu-toggle-no-javascript-required/>`_
.. [3] `css only toggle bulma navbar · siongui/paligo@eb3786d · GitHub <https://github.com/siongui/paligo/commit/eb3786d124df87025d09fe0b9dcd98eac87706b3>`_

.. _Bulma: http://bulma.io/
.. _CSS: https://www.google.com/search?q=CSS
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _nav bar: http://bulma.io/documentation/components/nav/
.. _Bulma navbar: https://versions.bulma.io/0.9.0/documentation/components/navbar/
.. _navigation bar: https://www.google.com/search?q=navigation+bar
.. _CSS only menu toggle: http://www.outofscope.com/css-only-menu-toggle-no-javascript-required/
.. _label: https://www.w3schools.com/TAGs/tag_label.asp
.. _input: https://www.w3schools.com/TAGs/tag_input.asp

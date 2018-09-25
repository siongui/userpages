Pure CSS Bootstrap Responsive Navbar (Navigation Bar)
#####################################################

:date: 2017-03-23T03:54+08:00
:tags: CSS, Responsive Web Design, toggle, toggleable, html, Pure CSS Toggle,
       navbar, Bootstrap
:category: CSS
:summary: CSS_ only responsive Bootstrap_ `navbar`_ (`navigation bar`_) without
          JavaScript_. Toggle mobile menu with CSS only. No JavaScript required.
:og_image: https://s3.envato.com/files/197871390/preview02.png
:adsu: yes


CSS_ only responsive Bootstrap_ `navbar`_ (`navigation bar`_) without
JavaScript_. Toggle mobile menu with CSS only. No JavaScript required.

.. rubric:: `Demo <{filename}/code/css/bootstrap-responsive-navbar/index.html>`_
   :class: align-center

I apply the technique of `CSS only menu toggle`_ to the responsive `navbar`_
code in Bootstrap_, and create this responsive navigation bar.

The key points of the technique [1]_ [2]_:
  - A visible HTML label_ element (only visible on small screen < 768px).
  - A invisible HTML input_ checkbox element, referenced by the label_ element.
  - The menu to be toggled (``ul.navbar-nav`` in demo)
  - A CSS rule to hide/show the menu according to whether the checkbox is
    checked.

When users click or touch on the visible label_ element, it will make the
invisible input_ checkbox checked/unchecked. And the CSS rule will show the menu
if the checkbox is checked, or hide the menu if the checkbox is unchecked.

.. adsu:: 2

Full source code is as follows:

.. show_github_file:: siongui userpages content/code/css/bootstrap-responsive-navbar/index.html
.. adsu:: 3

----

Tested on:
  - ``Chromium Version 56.0.2924.76 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``
  - ``Bootstrap 3.3.7``

----

References
++++++++++

.. [1] `CSS Only Bulma Responsive Nav Bar (Navigation Bar) <{filename}../../02/22/css-only-bulma-responsive-navbar%en.rst>`_
.. [2] `Pure CSS Toggle (Show/Hide) HTML Element <{filename}../../02/27/css-only-toggle-dom-element%en.rst>`_

.. _Bootstrap: http://getbootstrap.com/
.. _CSS: https://www.google.com/search?q=CSS
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _navbar: http://getbootstrap.com/components/#navbar
.. _navigation bar: https://www.google.com/search?q=navigation+bar
.. _CSS only menu toggle: http://www.outofscope.com/css-only-menu-toggle-no-javascript-required/
.. _label: https://www.w3schools.com/TAGs/tag_label.asp
.. _input: https://www.w3schools.com/TAGs/tag_input.asp

Pure CSS Bulma Responsive Nav Bar (Navigation Bar)
##################################################

:date: 2017-02-22T03:46+08:00
:tags: CSS, Responsive Web Design, toggle, toggleable, html, Pure CSS Toggle,
       Bulma
:category: CSS
:summary: CSS_ only responsive Bulma_ `nav bar`_ (`navigation bar`_) without
          JavaScript_. Toggle mobile menu with CSS only. No JavaScript required.
:adsu: yes


CSS_ only responsive Bulma_ `nav bar`_ (`navigation bar`_) without JavaScript_.
Toggle mobile menu with CSS only. No JavaScript required.

.. rubric:: `Demo <{filename}/code/css/bulma-responsive-navbar/index.html>`_
   :class: align-center

I apply the technique of `CSS only menu toggle`_ to the responsive `nav bar`_
code in Bulma_, and create this responsive navigation bar.

The key points of the technique [2]_:
  - A visible HTML label_ element (only visible on small screen < 768px).
  - A invisible HTML input_ checkbox element, referenced by the label_ element.
  - The menu to be toggled
  - A CSS rule to hide/show the menu according to whether the checkbox is
    checked.

When users click or touch on the visible label_ element, it will make the
invisible input_ checkbox checked/unchecked. And the CSS rule will show the menu
if the checkbox is checked, or hide the menu if the checkbox is unchecked.

.. adsu:: 2

Full source code is as follows:

.. show_github_file:: siongui userpages content/code/css/bulma-responsive-navbar/index.html
.. adsu:: 3

----

Tested on:
  - ``Chromium Version 55.0.2883.87 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``
  - ``Bulma 0.3.1``

----

References
++++++++++

.. [1] | `css only nav toggle - Google search <https://www.google.com/search?q=css+only+nav+toggle>`_
       | `css only nav toggle - DuckDuckGo search <https://duckduckgo.com/?q=css+only+nav+toggle>`_
       | `css only nav toggle - Ecosia search <https://www.ecosia.org/search?q=css+only+nav+toggle>`_
       | `css only nav toggle - Bing search <https://www.bing.com/search?q=css+only+nav+toggle>`_
       | `css only nav toggle - Yahoo search <https://search.yahoo.com/search?p=css+only+nav+toggle>`_
       | `css only nav toggle - Baidu search <https://www.baidu.com/s?wd=css+only+nav+toggle>`_
       | `css only nav toggle - Yandex search <https://www.yandex.com/search/?text=css+only+nav+toggle>`_
.. adsu:: 4
.. [2] `CSS only menu toggle - no JavaScript required <http://www.outofscope.com/css-only-menu-toggle-no-javascript-required/>`_
.. [3] `basic theme by Bulma · siongui/wat-pah-photiyan@7d3176c · GitHub <https://github.com/siongui/wat-pah-photiyan/commit/7d3176ce8e9754b4e4b556c19e9c96e30775f319>`_
.. [4] | `CSS 网格布局 - Google search <https://www.google.com/search?q=CSS+%E7%BD%91%E6%A0%BC%E5%B8%83%E5%B1%80>`_
       | `CSS 网格布局 - DuckDuckGo search <https://duckduckgo.com/?q=CSS+%E7%BD%91%E6%A0%BC%E5%B8%83%E5%B1%80>`_
       | `CSS 网格布局 - Ecosia search <https://www.ecosia.org/search?q=CSS+%E7%BD%91%E6%A0%BC%E5%B8%83%E5%B1%80>`_
       | `CSS 网格布局 - Bing search <https://www.bing.com/search?q=CSS+%E7%BD%91%E6%A0%BC%E5%B8%83%E5%B1%80>`_
       | `CSS 网格布局 - Yahoo search <https://search.yahoo.com/search?p=CSS+%E7%BD%91%E6%A0%BC%E5%B8%83%E5%B1%80>`_
       | `CSS 网格布局 - Baidu search <https://www.baidu.com/s?wd=CSS+%E7%BD%91%E6%A0%BC%E5%B8%83%E5%B1%80>`_
       | `CSS 网格布局 - Yandex search <https://www.yandex.com/search/?text=CSS+%E7%BD%91%E6%A0%BC%E5%B8%83%E5%B1%80>`_


.. _Bulma: http://bulma.io/
.. _CSS: https://www.google.com/search?q=CSS
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _nav bar: http://bulma.io/documentation/components/nav/
.. _navigation bar: https://www.google.com/search?q=navigation+bar
.. _CSS only menu toggle: http://www.outofscope.com/css-only-menu-toggle-no-javascript-required/
.. _label: https://www.w3schools.com/TAGs/tag_label.asp
.. _input: https://www.w3schools.com/TAGs/tag_input.asp

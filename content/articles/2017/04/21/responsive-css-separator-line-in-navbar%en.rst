Responsive CSS Divider in Nav Bar
#################################

:date: 2017-04-21T00:37+08:00
:tags: CSS, Responsive Web Design, html, navbar
:category: CSS
:summary: Responsive CSS separator line in *nav bar* (*navigation bar*).
          The CSS divider is vertical on tablet/desktop/wide screen and becomes
          horizontal on mobile/small screen.
:og_image: http://docs.themefuse.com/media/32/The-Core-Divider-5.jpg
:adsu: yes


Responsive CSS separator line in `nav bar`_ (*navigation bar*).
The CSS divider is vertical on tablet/desktop/wide screen and becomes horizontal
on mobile/small screen.

.. rubric:: Demo: See navbar of this webpage!
   :class: align-center

The CSS divider is used to separate items in responsive navbar [2]_.

**HTML**:

.. code-block:: html

  <span class="vertical-divider"></span>


**SCSS**: You can `compile the SCSS code online`_.

.. code-block:: scss

  $mobile: 768px;

  .vertical-divider {
    align-items: center;
    display: flex;
    margin: 0.5rem 0.25rem;
    border-right: 2px solid #fafafa;
    border-left: 2px solid #b4b4b4;
    @media (max-width: $mobile) {
      border-top: 2px solid #fafafa;
      border-bottom: 2px solid #b4b4b4;
      margin: 0;
    }
  }

.. adsu:: 2

Tested on:
``Chromium Version 57.0.2987.98 Built on Ubuntu , running on Ubuntu 17.04 (64-bit)``

----

References
++++++++++

.. [1] | `css vertical separator line - Google search <https://www.google.com/search?q=css+vertical+separator+line>`_
       | `css vertical separator line - DuckDuckGo search <https://duckduckgo.com/?q=css+vertical+separator+line>`_
       | `css vertical separator line - Ecosia search <https://www.ecosia.org/search?q=css+vertical+separator+line>`_
       | `css vertical separator line - Qwant search <https://www.qwant.com/?q=css+vertical+separator+line>`_
       | `css vertical separator line - Bing search <https://www.bing.com/search?q=css+vertical+separator+line>`_
       | `css vertical separator line - Yahoo search <https://search.yahoo.com/search?p=css+vertical+separator+line>`_
       | `css vertical separator line - Baidu search <https://www.baidu.com/s?wd=css+vertical+separator+line>`_
       | `css vertical separator line - Yandex search <https://www.yandex.com/search/?text=css+vertical+separator+line>`_

.. [2] `Pure CSS Responsive Nav Bar via Flexbox <{filename}../20/css-only-responsive-navbar-via-flexbox%en.rst>`_

.. _CSS: https://www.google.com/search?q=CSS
.. _nav bar: https://www.google.com/search?q=navigation+bar
.. _compile the SCSS code online: https://www.google.com/search?q=scss+to+css+online

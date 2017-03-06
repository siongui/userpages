[React] Tooltip
###############

:date: 2017-03-06T22:56+08:00
:tags: React, Tooltip, JavaScript, mouseenter, mouseleave, element offset, DOM,
       element position, JSX
:category: React
:summary: Simple tooltip_ implementation via React_ JavaScript_.
:og_image: https://facebook.github.io/react/img/logo.svg
:adsu: yes


See demo first. Move the cursor (mouse pointer) to hover over the blue text with
underline:

.. rubric:: `Demo <http://codepen.io/anon/pen/WpGjdQ?editors=0010>`_
   :class: align-center

The idea comes from the JavaScript_ implementation of my post [1]_. The
following is complete React_ source code.

**HTML**:

.. code-block:: html

  <div id="root">
    <!-- This div's content will be managed by React. -->
  </div>

**CSS**:

.. show_github_file:: siongui userpages content/code/react/tooltip/style.css
.. adsu:: 2

**JSX**:

.. show_github_file:: siongui userpages content/code/react/tooltip/app.jsx
.. adsu:: 3

----

Tested on: `CodePen <http://codepen.io/anon/pen/WpGjdQ?editors=0010>`__

----

References:

.. [1] `Pure CSS Tooltip and JavaScript Implementation <{filename}../04/css-only-tooltip-and-javascript-implementation%en.rst>`_

.. [2] | `react space - Google search <https://www.google.com/search?q=react+space>`_
       | `react space - DuckDuckGo search <https://duckduckgo.com/?q=react+space>`_
       | `react space - Ecosia search <https://www.ecosia.org/search?q=react+space>`_
       | `react space - Qwant search <https://www.qwant.com/?q=react+space>`_
       | `react space - Bing search <https://www.bing.com/search?q=react+space>`_
       | `react space - Yahoo search <https://search.yahoo.com/search?p=react+space>`_
       | `react space - Baidu search <https://www.baidu.com/s?wd=react+space>`_
       | `react space - Yandex search <https://www.yandex.com/search/?text=react+space>`_

.. [3] `reactjs - How to avoid JSX component from condensing when React.js rendering it? - Stack Overflow <http://stackoverflow.com/a/32690647>`_

.. _tooltip: https://www.google.com/search?q=tooltip
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _React: https://facebook.github.io/react/
.. _HTML: https://www.google.com/search?q=HTML
.. _CSS: https://www.google.com/search?q=CSS
.. _JSX: https://www.google.com/search?q=JSX

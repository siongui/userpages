[React] Toggle (Show/Hide) HTML Element
#######################################

:date: 2017-02-08T16:46+08:00
:tags: React, JavaScript, JSX, DOM, toggle, toggleable
:category: React
:summary: Toggle (Show/Hide) HTML_ `DOM element`_ in React_ JavaScript_.
:og_image: https://facebook.github.io/react/img/logo.svg
:adsu: yes


React_ is gaining more and more popularity now [1]_ [2]_, so I decide to try it
out with my favorite exercise: Toggle (Show/Hide) HTML_ `DOM element`_.

.. rubric:: `Demo <http://codepen.io/anon/pen/mRGqVd?editors=0010>`_
   :class: align-center

JSX_ code:

.. show_github_file:: siongui userpages content/code/react/toggle-element/app.jsx
.. adsu:: 2

The following is key points in the code:

- Two HTML_ *div* elements (line 24 and 25): one is control element and the
  other is content element. When users click on the control element, the content
  element will be shown if it's hidden, will be hidden if it's shown.
- A boolean state_ called ``isShow`` is defined (line 11). The onClick event of
  control element will toggle this state (line 15~19).
- If ``isShow`` is true, the *invisible* class will be added to content element
  (line 1~6), which makes the content element disappear.

CSS_ code:

.. show_github_file:: siongui userpages content/code/react/toggle-element/style.css

HTML_ code:

.. code-block:: html

  <div id="root">
    <!-- This div's content will be managed by React. -->
  </div>

----

Tested on: `CodePen <http://codepen.io/anon/pen/mRGqVd?editors=0010>`__

----

.. adsu:: 3

References:

.. [1] `Why I moved from Angular to React | Hacker News <https://news.ycombinator.com/item?id=13583059>`_
.. [2] `谁是 2016 年的 JavaScript 之最？ - 开源中国社区 <https://www.oschina.net/news/81155/who-is-javascript-new-star>`_

.. _React: https://facebook.github.io/react/
.. _HTML: https://www.google.com/search?q=HTML
.. _CSS: https://www.google.com/search?q=CSS
.. _JSX: https://www.google.com/search?q=JSX
.. _state: https://facebook.github.io/react/docs/state-and-lifecycle.html
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _GopherJS: http://www.gopherjs.org/
.. _DOM element: https://www.google.com/search?q=DOM+element
.. _Vue.js: https://vuejs.org/

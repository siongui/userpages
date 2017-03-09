[Vue.js] Animate.css Test Demo
##############################

:date: 2017-01-25T03:04+08:00
:tags: Vue.js, JavaScript, Animate.css, CSS, DOM, CSS Animation
:category: Vue.js
:summary: Use Animate.css_ and Vue.js_ to animate your webpage.
:og_image: https://camo.githubusercontent.com/224f79940611c6c12fb649128eca1cae31086d23/68747470733a2f2f7261776769742e636f6d2f7675656a732f617765736f6d652d7675652f6d61737465722f6c6f676f2e706e67
:adsu: yes


I simplify the `animations on GitHub Pages`_ of **Animate.css** and convert it
to Vue.js_ version.

.. rubric:: `Demo <{filename}/code/vuejs/animate.css-test/index.html>`_
   :class: align-center

We need only two directive_ to animate the text:

- v-model_: bind *select* element to the variable ``action``, which record the
            value of selected option.
- v-bind_: combine the class **animated** and the value in ``action`` to animate
           the text in *h1* element.

**Source code**:

.. adsu:: 2

.. show_github_file:: siongui userpages content/code/vuejs/animate.css-test/index.html

.. show_github_file:: siongui userpages content/code/vuejs/animate.css-test/app.js
.. adsu:: 3

For GopherJS_ version, see [1]_.

For AngularJS_ version, see [2]_.

----

Tested on:

- ``Chromium Version 55.0.2883.87 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``
- ``Vue.js 2.1.10``.

----

References:

.. [1] `[GopherJS] Animate.css Test Demo <{filename}../24/gopherjs-animate.css-test-demo%en.rst>`_

.. [2] `[AngularJS] Animate.css Test Demo <{filename}../26/angularjs-ng-animate.css-test-demo%en.rst>`_


.. _Vue.js: https://vuejs.org/
.. _Animate.css: https://daneden.github.io/animate.css/
.. _animations on GitHub Pages: https://daneden.github.io/animate.css/
.. _GopherJS: http://www.gopherjs.org/
.. _AngularJS: https://angularjs.org/
.. _directive: http://012.vuejs.org/guide/directives.html
.. _v-model: https://vuejs.org/v2/api/#v-model
.. _v-bind: https://vuejs.org/v2/api/#v-bind

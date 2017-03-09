[AngularJS] Animate.css Test Demo
#################################

:date: 2017-01-26T03:41+08:00
:tags: AngularJS, JavaScript, Animate.css, CSS, DOM, Directives, CSS Animation
:category: AngularJS
:summary: Use Animate.css_ and AngularJS_ to animate your webpage.
:og_image: https://angular.io/resources/images/logos/angular2/angular.svg
:adsu: yes


I simplify the `animations on GitHub Pages`_ of **Animate.css** and convert it
to AngularJS_ version.

.. rubric:: `Demo <{filename}/code/angularjs/animate.css-test/index.html>`_
   :class: align-center

We need only two directive_ to animate the text:

- ngModel_: bind *select* element to property named ``action``, which record the
            value of selected option.
- ngClass_: combine the class **animated** and the value in ``action`` to
            animate the text in *h1* element.

**Source code**:

.. adsu:: 2

.. show_github_file:: siongui userpages content/code/angularjs/animate.css-test/index.html

For GopherJS_ version, see [1]_.

For Vue.js_ version, see [2]_.

.. adsu:: 3

----

Tested on:

- ``Chromium Version 55.0.2883.87 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``
- ``AngularJS 1.6.1``.

----

References:

.. [1] `[GopherJS] Animate.css Test Demo <{filename}../24/gopherjs-animate.css-test-demo%en.rst>`_

.. [2] `[Vue.js] Animate.css Test Demo <{filename}../25/vuejs-animate.css-test-demo%en.rst>`_


.. _Vue.js: https://vuejs.org/
.. _AngularJS: https://angularjs.org/
.. _Animate.css: https://daneden.github.io/animate.css/
.. _animations on GitHub Pages: https://daneden.github.io/animate.css/
.. _GopherJS: http://www.gopherjs.org/
.. _directive: https://docs.angularjs.org/guide/directive
.. _ngClass: https://docs.angularjs.org/api/ng/directive/ngClass
.. _ngModel: https://docs.angularjs.org/api/ng/directive/ngModel

[AngularJS] Input Suggest Dropdown Menu
#######################################

:date: 2017-01-27T07:44+08:00
:tags: AngularJS, JavaScript, DOM, Directives, Web application, element offset,
       element position, dropdown menu
:category: AngularJS
:summary: Provide input suggest feature via AngularJS_. Usually used in
          dictionary application.
:og_image: https://angular.io/resources/images/logos/angular2/angular.svg
:adsu: yes


This is a simplified AngularJS_ version of my input suggest library written in
GopherJS_ [3]_. Usually used in dictionary application.

.. rubric:: `Demo <{filename}/code/angularjs/input-suggest/index.html>`_
   :class: align-center

The following are key aspects in the code:

- Use directive_ ngModel_ to bind *input* text element to property named
  ``word``, which record the value of user input.
- Use directive_ ngChange_ to monitor the changes of user input. If user input
  changes, ``inputChange()`` function calls ``suggestService`` to provide
  suggested words.
- ``suggestService`` generates random strings to provide suggested words. You
  should implement your own here.

**Source code**:

.. show_github_file:: siongui userpages content/code/angularjs/input-suggest/index.html
.. adsu:: 2
.. show_github_file:: siongui userpages content/code/angularjs/input-suggest/app.js
.. adsu:: 3
.. show_github_file:: siongui userpages content/code/angularjs/input-suggest/style.css

----

Tested on:

- ``Chromium Version 55.0.2883.87 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``
- ``AngularJS 1.6.1``.

----

References:

.. [1] `pali/input.html at master · siongui/pali · GitHub <https://github.com/siongui/pali/blob/master/dictionary/app/partials/input.html>`_

.. [2] `pali/inputSuggest.js at master · siongui/pali · GitHub <https://github.com/siongui/pali/blob/master/dictionary/app/scripts/directives/inputSuggest.js>`_

.. [3] `GitHub - siongui/gopherjs-input-suggest: input suggest menu via GopherJS <https://github.com/siongui/gopherjs-input-suggest>`_
       (`Demo <https://siongui.github.io/gopherjs-input-suggest/>`__)

.. [4] `angularjs - How to get element by classname or id - Stack Overflow <http://stackoverflow.com/questions/23609171/how-to-get-element-by-classname-or-id>`_

.. [5] `AngularJS: Error Reference: dupes <https://docs.angularjs.org/error/ngRepeat/dupes>`_

.. [6] `javascript - Angularjs get element position - Stack Overflow <http://stackoverflow.com/questions/27581260/angularjs-get-element-position>`_

.. _Vue.js: https://vuejs.org/
.. _AngularJS: https://angularjs.org/
.. _GopherJS: http://www.gopherjs.org/
.. _directive: https://docs.angularjs.org/guide/directive
.. _ngChange: https://docs.angularjs.org/api/ng/directive/ngChange
.. _ngModel: https://docs.angularjs.org/api/ng/directive/ngModel

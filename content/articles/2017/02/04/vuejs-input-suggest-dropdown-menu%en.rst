[Vue.js] Input Suggest Dropdown Menu
####################################

:date: 2017-02-04T05:27+08:00
:tags: Vue.js, JavaScript, DOM, Web application, element offset, dropdown menu,
       element position
:category: Vue.js
:summary: Provide input suggest feature via Vue.js_. Usually used in
          dictionary application.
:og_image: https://vuejs.org/images/logo.png
:adsu: yes


This is a simplified Vue.js_ version of my input suggest library written in
GopherJS_ [1]_. Usually used in dictionary application.

.. rubric:: `Demo <{filename}/code/vuejs/input-suggest/index.html>`_
   :class: align-center

The following are key aspects in the code:

- Use directive_ v-model_ to bind *input* text element to property named
  ``userinput``, which record the value of user input. The `.trim`_ modifier is
  also used to trim leading and trailing whitespaces of ``userinput``.
- watcher_ is used to monitor the changes of ``userinput`` [4]_. If
  ``userinput`` changes, we generate 10 random strings to provide suggested
  words. You should implement your own suggested words here.
- `v-bind:style`_, `v-show`_, and `computed properties`_ are used to provide
  style and visibility of ``div.suggest`` element.

**Source code**:

.. show_github_file:: siongui userpages content/code/vuejs/input-suggest/index.html
.. adsu:: 2
.. show_github_file:: siongui userpages content/code/vuejs/input-suggest/app.js
.. adsu:: 3
.. show_github_file:: siongui userpages content/code/vuejs/input-suggest/style.css

| For AngularJS_ version, see [5]_.

----

Tested on:

- ``Chromium Version 55.0.2883.87 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``
- ``Vue.js 2.1.10``.

----

References:

.. [1] `GitHub - siongui/gopherjs-input-suggest: input suggest menu via GopherJS <https://github.com/siongui/gopherjs-input-suggest>`_
       (`Demo <https://siongui.github.io/gopherjs-input-suggest/>`__)

.. [2] `vuejs autofocus - Google search <https://www.google.com/search?q=vuejs+autofocus>`_

       `vuejs autofocus - DuckDuckGo search <https://duckduckgo.com/?q=vuejs+autofocus>`_

       `vuejs autofocus - Bing search <https://www.bing.com/search?q=vuejs+autofocus>`_

       `vuejs autofocus - Yahoo search <https://search.yahoo.com/search?p=vuejs+autofocus>`_

       `vuejs autofocus - Baidu search <https://www.baidu.com/s?wd=vuejs+autofocus>`_

       `vuejs autofocus - Yandex search <https://www.yandex.com/search/?text=vuejs+autofocus>`_

.. [3] `Vuejs - How to Focus Inputs - Tutelage Systems <http://tutelagesystems.com/blog/2016-01-08-vuejs-how-to-focus-inputs>`_

.. [4] `[Vue.js] Input Text Element Change Event <{filename}../03/vuejs-input-change-event%en.rst>`_

.. [5] `[AngularJS] Input Suggest Dropdown Menu <{filename}../../01/27/angularjs-ng-input-suggest-dropdown-menu%en.rst>`_

.. _Vue.js: https://vuejs.org/
.. _AngularJS: https://angularjs.org/
.. _GopherJS: http://www.gopherjs.org/
.. _directive: https://www.google.com/search?q=vuejs+directive
.. _v-model: https://vuejs.org/v2/api/#v-model
.. _.trim: https://vuejs.org/v2/guide/forms.html#trim
.. _watcher: https://vuejs.org/v2/guide/computed.html#Watchers
.. _v-bind\:style: https://vuejs.org/v2/guide/class-and-style.html
.. _v-show: https://vuejs.org/v2/api/#v-show
.. _computed properties: https://vuejs.org/v2/guide/computed.html

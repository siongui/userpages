[AngularJS] Dropdown Menu Using Directive
#########################################

:tags: AngularJS, Directives, dropdown menu, without jQuery, JavaScript
:category: AngularJS
:summary: Use AngularJS_ powerful custom directive_ to implement
          `dropdown menu`_ without jQuery_ or any other library.
:adsu: yes


This post shows how to use AngularJS_ powerful directives_ to create
`dropdown menu`_ without using any library (of course, except AngularJS).

Please first see:

.. rubric:: `Demo <{filename}/code/angularjs/dropdown-menu/ngdropdown.html>`_
   :class: align-center

**JavaScript**: See `AngularJS: Developer Guide: Directives`_ and
`AngularJS: API: $compile`_ for how to write custom directives and API.

.. show_github_file:: siongui userpages content/code/angularjs/dropdown-menu/ngdropdown.js
.. adsu:: 2

**HTML**: From line 11 to 15, the menu is placed inside a *div* element, the
name ``dropdown`` of the custom directive is the attribute of the *div* element,
and we use three other attributes (*classlink*, *linktext*, *classmenu*) to pass
CSS class names and link text to the dropdown menu template in custom directive.

.. show_github_file:: siongui userpages content/code/angularjs/dropdown-menu/ngdropdown.html
.. adsu:: 3

**CSS**: Nothing special here. The same as you would do in a normal
non-angularjs dropdown menu.

.. show_github_file:: siongui userpages content/code/angularjs/dropdown-menu/style.css

Tested AngularJS version: ``1.3.11``, ``1.1.5``.

----

References:

.. [1] `angularjs dropdown directive hide when clicked outisde - Stack Overflow <http://stackoverflow.com/questions/14574365/angularjs-dropdown-directive-hide-when-clicking-outside>`_

.. _AngularJS: https://angularjs.org/
.. _directive: https://docs.angularjs.org/guide/directive
.. _directives: https://docs.angularjs.org/guide/directive
.. _dropdown menu: https://www.google.com/search?q=dropdown+menu
.. _jQuery: https://jquery.com/
.. _AngularJS\: Developer Guide\: Directives: https://docs.angularjs.org/guide/directive
.. _AngularJS\: API\: $compile: https://docs.angularjs.org/api/ng/service/$compile

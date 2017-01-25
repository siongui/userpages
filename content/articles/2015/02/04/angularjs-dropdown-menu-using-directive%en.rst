[AngularJS] Dropdown Menu Using Directive
#########################################

:tags: AngularJS, Directives, dropdown menu, without jQuery, JavaScript
:category: AngularJS
:summary: Use AngularJS powerful custom directive to implement dropdown menu without jQuery or any other library.
:adsu: yes


This post shows how to use `AngularJS <https://angularjs.org/>`_ powerful `directives <https://docs.angularjs.org/guide/directive>`_ to create dropdown menu without using any library (of course, except AngularJS).

Please first see:

.. rubric:: `Demo <{filename}ngdropdown.html>`_
   :class: align-center

js: See `Developer Guide: Directives <https://docs.angularjs.org/guide/directive>`_ and `API: $compile <https://docs.angularjs.org/api/ng/service/$compile>`_ for how to write custom directives and API.

.. show_github_file:: siongui userpages content/articles/2015/02/04/ngdropdown.js

.. adsu:: 2

html: From line 11 to 15, the menu is placed inside a *div* element, the name ``dropdown`` of the custom directive is the attribute of the *div* element, and we use three other attributes (*classlink*, *linktext*, *classmenu*) to pass css class names and link text to the dropdown menu template in custom directive.

.. show_github_file:: siongui userpages content/articles/2015/02/04/ngdropdown.html

css: Nothing special here. The same as you would do in a normal non-angularjs dropdown menu.

.. show_github_file:: siongui userpages content/articles/2015/02/04/style.css

Tested AngularJS version: ``1.3.11``, ``1.1.5``.

----

References:

[1] `angularjs dropdown directive hide when clicked outisde - Stack Overflow <http://stackoverflow.com/questions/14574365/angularjs-dropdown-directive-hide-when-clicking-outside>`_

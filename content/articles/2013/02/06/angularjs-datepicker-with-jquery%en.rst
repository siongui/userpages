[AngularJS] datepicker with jQuery
##################################

:date: 2013-02-06 02:38
:modified: 2015-02-23 19:47
:tags: JavaScript, jQuery, AngularJS
:category: AngularJS
:summary: jQuery datepicker wrapped in AngularJS directive.
:adsu: yes


The best way to implement a datepicker on website, in my opinion, is through
HTML5 input tag with *type=date*. This input type, however, is not
`well-supported`_. AngularJS_ is a very promising extension for HTML, so I did
some study on how to implement datepicker with AngularJS_. The easiest way is to
combine AngularJS directive and jQuery datepicker. The following is demo and
complete source code:

.. rubric:: `Demo <{filename}/code/angularjs-jquery-datepicker/directive.html>`_
      :class: align-center

.. show_github_file:: siongui userpages content/code/angularjs-jquery-datepicker/directive.html
.. adsu:: 2
.. show_github_file:: siongui userpages content/code/angularjs-jquery-datepicker/directive.js


Sample code tested on AngularJS_ ``1.0.4`` and ``1.3.11``

----

References:

.. [1] `AngularJS: Developer Guide: Directives <https://docs.angularjs.org/guide/directive>`_

.. [2] `Datepicker Widget | jQuery UI API Documentation <http://api.jqueryui.com/datepicker/>`_

.. [3] `jsFiddle demo <http://jsfiddle.net/nnsese/xB6c2/26/>`_

.. _AngularJS: https://angularjs.org/

.. _well-supported: http://www.w3schools.com/html/html_form_input_types.asp

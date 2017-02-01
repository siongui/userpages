[AngularJS] Get Element Offset (Position)
#########################################

:tags: AngularJS, element offset, element position, without jQuery, JavaScript
:category: AngularJS
:summary: Get the offset (position) of an AngularJS DOM element without jQuery
:adsu: yes


AngularJS_ doesn't provide `offset()`_ in jqLite_. If you don't want to include
jQuery_ for only `offset()`_, the following is a possible implementation of
`offset()`_:

.. code-block:: javascript

    function offset(elm) {
      try {return elm.offset();} catch(e) {}
      var rawDom = elm[0];
      var _x = 0;
      var _y = 0;
      var body = document.documentElement || document.body;
      var scrollX = window.pageXOffset || body.scrollLeft;
      var scrollY = window.pageYOffset || body.scrollTop;
      _x = rawDom.getBoundingClientRect().left + scrollX;
      _y = rawDom.getBoundingClientRect().top + scrollY;
      return { left: _x, top: _y };
    }

.. adsu:: 2

This function takes one argument, which is the element wrapped with jqLite_.
A try-catch block is added in the beginning to use the `offset()`_ of jQuery if
jQuery is included.

----

Reference:

.. adsu:: 3
.. [1] `getBoundingClientRect method <http://help.dottoro.com/ljvmcrrn.php>`_

.. _AngularJS: https://angularjs.org/
.. _offset(): https://api.jquery.com/offset/
.. _jqLite: https://docs.angularjs.org/api/ng/function/angular.element
.. _jQuery: https://jquery.com/

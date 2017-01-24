[AngularJS] Get Element Offset (Position)
#########################################

:tags: AngularJS, element offset, element position, without jQuery, JavaScript
:category: AngularJS
:summary: Get the offset (position) of an AngularJS DOM element without jQuery
:adsu: yes


`AngularJS <https://angularjs.org/>`_ doesn't provide `offset() <http://api.jquery.com/offset/>`__ in `jqLite <https://docs.angularjs.org/api/ng/function/angular.element>`__. If you don't want to include `jQuery <http://jquery.com/>`_ for only `offset() <http://api.jquery.com/offset/>`__, the following is a possible implementation of `offset() <http://api.jquery.com/offset/>`__:

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

This function takes one argument, which is the element wrapped with `jqLite <https://docs.angularjs.org/api/ng/function/angular.element>`__. A try-catch block is added in the beginning to use the `offset() <http://api.jquery.com/offset/>`__ of jQuery if jQuery is included.

.. adsu:: 2

----

Reference:

[1] `getBoundingClientRect method <http://help.dottoro.com/ljvmcrrn.php>`_

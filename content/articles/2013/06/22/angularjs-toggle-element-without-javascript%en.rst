[AngularJS] Toggle Element without JavaScript
#############################################

:tags: AngularJS, JavaScript, toggle, toggleable
:category: AngularJS
:summary: Example of toggling DOM_ element using AngularJS_,
          without any JavaScript_ code.
:adsu: yes


The power of AngularJS_ is best described by toggling element without writing
any JavaScript_ code. This example will show you how to do toggle-able effect in
AngularJS_ way.

Please first see:

.. rubric:: `Demo <{filename}ngtoggle.html>`_
   :class: align-center

The source code of above demo:

.. adsu:: 2
.. show_github_file:: siongui userpages content/articles/2013/06/22/ngtoggle.html
.. adsu:: 3

As you can see from the source code, we use a variable **isShow** and built-in
directives ng-init_, ng-click_, ng-show_, and ng-hide_ to achieve toggle-able
effect. The beauty of this solution is that you do not need to write any
JavaScript code, and the code is easily readable and understandable at the same
time!

.. _AngularJS: https://angularjs.org/
.. _DOM: https://www.google.com/search?q=DOM
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _ng-init: https://docs.angularjs.org/api/ng/directive/ngInit
.. _ng-click: https://docs.angularjs.org/api/ng/directive/ngClick
.. _ng-show: https://docs.angularjs.org/api/ng/directive/ngShow
.. _ng-hide: https://docs.angularjs.org/api/ng/directive/ngHide

[AngularJS] Toggle Element without JavaScript
#############################################

:tags: AngularJS, JavaScript, toggle, toggleable
:category: AngularJS
:summary: Example of toggling DOM element using AngularJS, without any JavaScript code.


The power of `AngularJS <https://angularjs.org/>`_ is best described by toggling element without writing any JavaScript code. This example will show you how to do toggle-able effect in AngularJS way.

Please first see:

.. rubric:: `Demo <{filename}ngtoggle.html>`_

The source code of above demo:

.. show_github_file:: siongui userpages content/articles/2013/06/22/ngtoggle.html

As you can see from the source code, we use a variable **isShow** and built-in directives `ng-init <https://docs.angularjs.org/api/ng/directive/ngInit>`_, `ng-click <https://docs.angularjs.org/api/ng/directive/ngClick>`_, `ng-show <https://docs.angularjs.org/api/ng/directive/ngShow>`_, and `ng-hide <https://docs.angularjs.org/api/ng/directive/ngHide>`_ to achieve toggle-able effect. The beauty of this solution is that you do not need to write any JavaScript code, and the code is easily readable and understandable at the same time!

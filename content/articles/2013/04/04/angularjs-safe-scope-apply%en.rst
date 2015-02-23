[AngularJS] Safe Scope.$apply Implementation (Error: $apply already in progress)
################################################################################

:date: 2013-04-04 02:57
:modified: 2015-02-23 11:20
:tags: JavaScript, AngularJS
:category: AngularJS
:summary: Safely execute an expression in angular from outside of the angular framework.


One of the headaches of writing one's own directives or services of AngularJS_
application is that get the following error message while the `$apply`_ method
of scope_ is called:

.. code-block:: txt

  Error: $apply already in progress

The following is my solution (call the following *safeApply* method instead of
$apply):

.. code-block:: javascript

  $scope.safeApply = function(fn) {
    var phase = this.$root.$$phase;
    if(phase == '$apply' || phase == '$digest')
      this.$eval(fn);
    else
      this.$apply(fn);
  };

  // OR

  function safeApply(scope, fn) {
    var phase = scope.$root.$$phase;
    if(phase == '$apply' || phase == '$digest')
      scope.$eval(fn);
    else
      scope.$apply(fn);
  }

The *fn* in above sample code could be AngularJS expression_ or JavaScript
function, depending on your need.

----

References:

.. [1] `my gist <https://gist.github.com/siongui/4969449>`_

.. [2] `angularjs - Prevent error $digest already in progress when calling $scope.$apply() - Stack Overflow <http://stackoverflow.com/questions/12729122/prevent-error-digest-already-in-progress-when-calling-scope-apply>`_

.. [3] `'Safe' $apply in Angular.JS <https://coderwall.com/p/ngisma/safe-apply-in-angular-js>`_


.. _AngularJS: https://angularjs.org/

.. _$apply: https://docs.angularjs.org/api/ng/type/$rootScope.Scope#$apply

.. _scope: https://docs.angularjs.org/guide/scope

.. _expression: https://docs.angularjs.org/guide/expression

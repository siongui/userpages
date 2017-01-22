[AngularJS] Update Property in Parent Scope
###########################################

:date: 2013-02-08 04:47
:modified: 2015-02-23 13:27
:tags: JavaScript, AngularJS
:category: AngularJS
:summary: Update parent scope from child scope by event dispatching and
          listening mechanism.
:adsu: yes

When developing AngularJS_ application, sometimes we need to update properties
in parent scope_ from child scope. Although properties inherited from parent
scope are accessible from child scope, they cannot be updated directly from
child scope. So how can we update properties in parent scope? After some
Googling, I found that the event dispatching and listening mechanism in
AngularJS can be used for our purpose. The following is an example
demonstrating how to do:

.. rubric:: `Demo <{filename}/code/angularjs-update-parent-scope/event.html>`_
      :class: align-center

.. show_github_file:: siongui userpages content/code/angularjs-update-parent-scope/event.html

.. show_github_file:: siongui userpages content/code/angularjs-update-parent-scope/event.js

In this example, we have a property called *parentProperty* in parent scope.
Initially the *parentProperty* is set to 0. When a user clicks on the button, an
*updateParentScopeEvent* event will be dispatched from child scope to parent
scope through `$emit`_ function. The parent scope registers a listener through
`$on`_ function. In the listener, the property in parent scope is updated, which
is the result we want.

This simple example demonstrates how to update properties in parent scope by
event dispatching and listening. Hope this would be helpful!

Sample code tested on AngularJS_ ``1.0.4`` and ``1.3.11``

----

References:

.. [1] `AngularJS: API: $rootScope.Scope <https://docs.angularjs.org/api/ng/type/$rootScope.Scope>`_

.. [2] `AngularJS: Developer Guide: Scopes <https://docs.angularjs.org/guide/scope>`_

.. _AngularJS: https://angularjs.org/

.. _scope: https://docs.angularjs.org/guide/scope

.. _$emit: https://docs.angularjs.org/api/ng/type/$rootScope.Scope#$emit

.. _$on: https://docs.angularjs.org/api/ng/type/$rootScope.Scope#$on

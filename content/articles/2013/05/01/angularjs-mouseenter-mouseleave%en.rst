[AngularJS] Incorrect ng-mouseenter ng-mouseleave and Solution
##############################################################

:date: 2013-05-01 06:53
:modified: 2015-02-24 14:22
:tags: AngularJS, JavaScript, Directives, mouseenter, mouseleave
:category: AngularJS
:summary: Solution to mouseenter and mouseleave event in old version AngularJS.
:adsu: yes


Old version AngularJS_ (such as 1.1.3) has incorrect implementation of
mouseenter_ and mouseleave_ directive_:

.. rubric:: `Demo <{filename}/code/angularjs-mouseenter-mouseleave/ngmouseenterleave.html>`_ of incorrect implementation
      :class: align-center

See also this
`pull request (fix(jqLite): mouseenter/-leave should not trigger on child elements) <https://github.com/angular/angular.js/pull/2134>`_
on AngularJS Github repo for more details.

The following is correct implementation of mouseenter and mouseleave directive:

.. show_github_file:: siongui userpages content/code/angularjs-mouseenter-mouseleave/mouseEnterLeave.js

The usage is the same as ngMouseenter_ and ngMouseleave_, except the name
changed to *mouseenter* and *mouseleave*. Also do not forget to include this
module in your AngularJS application.


.. _AngularJS: https://angularjs.org/

.. _mouseenter: http://api.jquery.com/mouseenter/

.. _mouseleave: http://api.jquery.com/mouseleave/

.. _directive: https://docs.angularjs.org/guide/directive

.. _ngMouseenter: https://docs.angularjs.org/api/ng/directive/ngMouseenter

.. _ngMouseleave: https://docs.angularjs.org/api/ng/directive/ngMouseleave

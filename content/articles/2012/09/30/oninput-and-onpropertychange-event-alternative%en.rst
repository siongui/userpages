[JavaScript] oninput and onpropertychange Event Alternative
###########################################################

:date: 2012-09-30 22:06
:modified: 2015-02-22 22:57
:tags: html, JavaScript, DOM
:category: JavaScript
:summary: Detect the content change of HTML input elements in a cross-browser
          and consistent manner.
:adsu: yes


If we want to detect the content change of HTML `input elements`_, like
textarea_, `input:text`_, `input:password`_, etc., we can use oninput_ event of
HTML `input elements`_. The oninput_ event, however, (and onpropertychange_
event in IE) is a disaster for web developers. Becasue

  - The oninput_ event is not supported in IE6, IE7, and IE8.

  - Even oninput_ event is supported in IE9, it is still buggy.

Although we can use onpropertychange_ event in IE browser, it is still very
difficult to maintain a consistent behavior. So here comes the question:

How do we detect the content change of HTML `input elements`_ in a cross-browser and consistent way?
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

The solution is inspired by the library in reference [3]_. It is simple, use
JavaScript built-in setTimeout_ or setInterval_ function to periodically check
the content of `input elements`_, and if the content changes, then an event
handler function is fired to handle this event. I wrote the following example to
demonstrate the solution:

.. adsu:: 2

.. rubric:: `Demo <{filename}/code/javascript-oninput-event-alternative/polling.html>`_
      :class: align-center

.. show_github_file:: siongui userpages content/code/javascript-oninput-event-alternative/polling.html
.. adsu:: 3
.. show_github_file:: siongui userpages content/code/javascript-oninput-event-alternative/polling.js

----

References:

.. [1] `oninput event | input event - Dottoro <http://help.dottoro.com/ljhxklln.php>`_

.. [2] `onpropertychange event | propertychange event - Dottoro <http://help.dottoro.com/ljufknus.php>`_

.. [3] `suggest.js - Input Suggest Library <http://www.enjoyxstudy.com/javascript/suggest/index.en.html>`_


.. _input elements: http://www.w3schools.com/html/html_forms.asp

.. _textarea: http://help.dottoro.com/ljtqbjui.php

.. _input\:text: http://help.dottoro.com/ljtdrupr.php

.. _input\:password: http://help.dottoro.com/ljevnnxp.php

.. _oninput: http://help.dottoro.com/ljhxklln.php

.. _onpropertychange: http://help.dottoro.com/ljufknus.php

.. _setTimeout: http://www.w3schools.com/js/js_timing.asp

.. _setInterval: http://www.w3schools.com/js/js_timing.asp

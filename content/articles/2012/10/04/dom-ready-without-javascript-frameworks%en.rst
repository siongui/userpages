DOM Ready without JavaScript Frameworks
#######################################

:date: 2012-10-04 23:22
:tags: JavaScript, DOM
:category: JavaScript
:summary: Check DOM readiness in vanilla JavaScript.


"`DOM ready`_" is an important event because HTML elements like *div* or *span*
can be accessed by JavaScript only when DOM is ready enough. Usually we can
initialize web application at "DOM ready" time. Various JavaScript frameworks
provide DOM ready event [1]_. Here, I want to discuss how to detect "DOM ready"
in a cross-browser way without using JavaScript Frameworks.

1. **Put and include JavaScript code right before the end of body tag**:

   This is simplest way to make sure DOM is ready when any JavaScript code is
   run. You do not have to write any code to detect whether DOM is ready, or
   reply on any "DOM ready" event of JavaScript frameworks. And your website
   will make users feels more responsive because users can see something on the
   browser screen before JavaScript code is run.

2. **Check the readiness of last DOM element**:

   If for some reason you don't want to put and include your JavaScript code
   before the end of *body* tag, here is an alternative - *Check the readiness
   of last element inside body tag*. The following JavaScript code snippet
   demonstrates how (assume the *id* of last element is *lastElement*):

   .. code-block:: javascript

     checkDOMReady();

     function checkDOMReady() {
       if (document.getElementById('lastElement')) {
         // DOM is ready now. Do something here.
       }
       else setTimeout(checkDOMReady, 50);
     }

   As you can see, a function is setup to access the last element. If the last
   element is not accessible, then the function sleep 50 ms and checks again
   until the last element is accessible. The checking interval (50 ms) could be
   made shorter if you like. In fact, you can check the *last element referenced
   in your code* instead of *last element inside body tag* if you want
   JavaScript code runs earlier.

Hope the above information is helpful for vanilla JavaScript lovers.

----

References:

.. [1] `DomReady Event Methods in JavaScript Frameworks <http://davidwalsh.name/javascript-domready>`_

.. [2] `Javascript DOM ready without an entire framework <http://stackoverflow.com/questions/2732171/javascript-dom-ready-without-an-entire-framework>`_

.. [3] `Relying on DOM readiness to invoke a function (instead of window.onload) <http://www.javascriptkit.com/dhtmltutors/domready.shtml>`_


.. _DOM ready: http://www.mootorial.com/wiki/mootorial/05-utilities/01-domready

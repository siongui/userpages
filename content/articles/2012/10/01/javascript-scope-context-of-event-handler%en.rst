[JavaScript] Scope (Context) of Event Handler Function
######################################################

:date: 2012-10-01 03:17
:modified: 2015-03-07 01:09
:tags: JavaScript, Function.prototype.bind()
:category: JavaScript
:summary: Discuss the scope (context) of event handler and how to change what
          "this" keyword refers to.


The scope (context) of event handler is headache for new JavaScript developers
(at least for me when I try to write JavaScript web application in the
beginning). Look at the following code snippet:

.. code-block:: javascript
  :linenos: table

  contact = function() {
    this.name = 'Bob';

    window.addEventListener("click", this.AddName, false);
  };

  contact.prototype.AddName = function() {
    alert(this.name);
  };

The problems in references [3]_ and [5]_ are basically the same as the above.
When JavaScript interpreter runs the line 8 of above code, error will be raised.
The JavaScript interpreter will tell you that *this.name* does not exist. You
may wonder that *this.name* is already assigned in line 2, why does the
JavaScript interpreter raise error?

The problem is that *this* keyword in line 8 refers to JavaScript built-in
object *window*, not refers to the *contact* function object. In line 2, you
assign the value *'Bob'* to the variable *name* in the scope (context) of
*contact* function object. As a result, the JavaScript interpreter raises the
error.

So how to make the *this* keyword in line 8 refer to *contact* function object?

For recent browsers which support Function.prototype.bind_ natively, you can
replace line 4:

.. code-block:: javascript

  window.addEventListener("click", this.AddName, false);

with the following line:

.. code-block:: javascript

  window.addEventListener("click", this.AddName.bind(this), false);

If you want to support old browsers like IE8, you can implement
`custom Function.prototype.bind`_, or you can use closures to change the scope
(context): line 4 can be replaced with the following code:

.. code-block:: javascript

  var self = this;
  window.addEventListener("click", function() {self.AddName();}, false);

Anonymous function and closures are used to make the *this* keyword in *AddName*
function refer to *contact* function object.

Other variant of the above problem is described in references [4]_. In summary,
there are two ways to solve the above problem:

  1. Function.prototype.bind_: Only provided in recent browsers, for old
     browsers, implement `custom Function.prototype.bind`_. See reference [1]_
     for more details.

  2. Closures: See references [2]_ and [6]_ for more details.

The questions on `Stack Overflow`_ in references [3]_, [4]_, [5]_ are very
helpful. Please read carefully.

----

References:

.. [1] `Function.prototype.bind() - JavaScript | MDN <https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Function/bind>`_

.. [2] `Objects, event handlers and "this" in javascript  –  Helephant.com <http://helephant.com/2008/04/26/objects-event-handlers-and-this-in-javascript/>`_

.. [3] `JavaScript Event Handler Scope - Stack Overflow <http://stackoverflow.com/questions/9488468/javascript-event-handler-scope>`_

.. [4] `oop - Global Javascript Event Handling Object Context - Stack Overflow <http://stackoverflow.com/questions/2241896/global-javascript-event-handling-object-context>`_

.. [5] `javascript - Event Handler Called With Wrong Context - Stack Overflow <http://stackoverflow.com/questions/6300817/event-handler-called-with-wrong-context>`_

.. [6] `Taming the Javascript event scope: Closures <http://www.tibobeijen.nl/blog/2009/07/27/taming-the-javascript-event-scope-closures/>`_


.. _Function.prototype.bind: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Function/bind

.. _custom Function.prototype.bind: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Function/bind#Polyfill

.. _Stack Overflow: http://stackoverflow.com/

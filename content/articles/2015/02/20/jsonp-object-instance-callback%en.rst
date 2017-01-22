Use Object Instance Function as JSONP Callback Function
#######################################################

:date: 2015-02-20 22:27
:tags: JavaScript, CORS, JSONP, Function.prototype.bind(), Google App Engine,
       Python, JSON
:category: JavaScript
:summary: Cross-domain requests by JSONP with object instance function as
          callback on Google App Engine for Python.
:adsu: yes


In my previous posts, a basic JSONP example [1]_, and another example of JSONP
with anonymous function as callback [2]_ are shown. In this post, an interesting
example of using object instance function (or method) as callback function in
JSONP will be shown. The is helpful if we want to make the code more
object-oriented.

The following is complete sample code.

*index.html* (run on client side, i.e., browser):

.. show_github_file:: siongui userpages content/code/jsonp-object-instance-callback/index.html

.. note::

  The *action* property of *form* element is not set in HTML. And we initialize
  an object at the *onload* event the *body* element. The form *onsubmit* event
  will be set in the constructor of the object.

*jsonp.js* (run on client side, i.e., browser):

.. show_github_file:: siongui userpages content/code/jsonp-object-instance-callback/jsonp.js

.. note::

  .. code-block:: javascript

    form.onsubmit = this.jsonp.bind(this);

  The bind_ function is used to keep **this** keyword refering to the object
  self when *jsonp* function is called.

  .. code-block:: javascript

    window['demoObjInGlobal'] = this;

  an alias name of the object instance in *global scope*. This is used to refer
  to this object instance in the communication of HTTP GET request. A better way
  here is to generate an un-used random name as alias name. If more than one
  instance of the object is needed, you have to generate different names for
  each object instances to prevent name conflict. For simplicity of the example,
  I use fixed name for the alias.

  .. code-block:: javascript

    var url = '/jsonp?callback=' +
        encodeURIComponent('demoObjInGlobal["callback"]') +
        '&input1=' + encodeURIComponent(input1value) +
        '&input2=' + encodeURIComponent(input2value);

  The alias name of the object instance is supplied to make the callback
  function of the object instance being executed after the *script* tag is
  appended to the *head* element. This is the main trick to make our object
  instance function to execute.

*jsonp.py* (run on server side, i.e., GAE Python):

.. show_github_file:: siongui userpages content/code/jsonp-object-instance-callback/jsonp.py

.. note::

  The above code on server side is basically the same as that in [2]_.

*app.yaml* (on server side, GAE Python config file):

.. show_github_file:: siongui userpages content/code/jsonp-object-instance-callback/app.yaml

In summary, the benefit of wrapping JSONP functionality inside an object is to
make the code modular, reusable, and more object-oriented.


Tested on: ``Ubuntu Linux 14.10``, ``Google App Engine Python SDK 1.9.18``

----

References:

.. [1] `JSONP on Google App Engine Python <{filename}jsonp-on-google-app-engine-python%en.rst>`_

.. [2] `JSONP with Anonymous Callback Function <{filename}jsonp-anonymous-callback-function%en.rst>`_

.. _bind: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Function/bind

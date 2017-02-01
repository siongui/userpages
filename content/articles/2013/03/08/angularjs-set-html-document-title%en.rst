[AngularJS] Set HTML Document Title
###################################

:tags: AngularJS, document title, html, JavaScript
:category: AngularJS
:summary: Set HTML document title in AngularJS
:adsu: yes


In pure JavaScript, the following way can be used to set HTML document title:

.. code-block:: javascript

    document.title = "myTitle";


However, this is not Angular way of doing programming, `AngularJS <http://angularjs.org/>`_ provide `$document <https://docs.angularjs.org/api/ng/service/$document>`_ for programmers to deal with HTML document. So I use the following syntax to set HTML document title:

.. code-block:: javascript

    $document.title = "myTitle"; // this doesn't work

.. adsu:: 2

But it does not work. Finally I found the correct way to set HTML document:

.. code-block:: javascript

    $document.prop('title', 'myTitle');


For those who has the same problem as me.

.. adsu:: 3

[Dart] Access HTML Data Attribute
#################################

:date: 2015-03-01 08:32
:tags: Dart, DOM, html
:category: Dart
:summary: Access data-* attribute of HTML elements in Dart
:adsu: yes


Explain by example, if you have following element in HTML:

.. code-block:: html

  <div id="myDiv" data-my-demo-value="value"></div>

Access the *value* by **dataset** property:

.. code-block:: dart

  var value = querySelector("#myDiv").dataset["myDemoValue"];

----

References:

.. [1] `HTML Global data-* Attributes - W3Schools <http://www.w3schools.com/tags/att_global_data.asp>`_

.. [2] `Getting HTML5 data-* attribute with Dart - Stack Overflow <http://stackoverflow.com/questions/20916927/getting-html5-data-attribute-with-dart>`_

.. [3] `dataset - dart:html.Element API Docs <https://api.dartlang.org/apidocs/channels/stable/dartdoc-viewer/dart:html.Element#id_dataset>`_

.. [4] Google search : `HTML Data Attribute <https://www.google.com/search?q=HTML+Data+Attribute>`_

.. [5] `[Golang] GopherJS DOM Example - Access HTML Data Attribute <{filename}../../../2016/01/12/gopherjs-dom-example-access-html-data-attribute%en.rst>`_

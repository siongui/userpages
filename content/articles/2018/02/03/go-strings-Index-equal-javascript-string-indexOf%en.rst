Golang strings.Index = JavaScript String indexOf()
##################################################

:date: 2018-02-03T04:51+08:00
:tags: Go, Golang, JavaScript, String Manipulation, Go strings Package
:category: Go
:summary: **Go** *strings.Index* is equivalent to
          **JavaScript** *String.prototype.indexOf()*
:og_image: https://www.safaribooksonline.com/library/view/javascript-the-missing/9780596515898/httpatomoreillycomsourceoreillyimages1489222.png
:adsu: yes


Go strings.Index_ = JavaScript `String.prototype.indexOf()`_

The following Go code (`Run Online <https://play.golang.org/p/nNFAwDXifW4>`__):

.. code-block:: go

  strings.Index("chicken", "ken")

is the same as the following JavaScript code (`Run Online <http://www.webtoolkitonline.com/javascript-tester.html>`__):

.. code-block:: javascript

  "chicken".indexOf("ken")

.. adsu:: 2

----

Moreover, Go strings.Contains_ = JavaScript `String.prototype.indexOf() != -1`_

The following Go code (`Run Online <https://play.golang.org/p/_sAqWYrDJnp>`__):

.. code-block:: go

  strings.Contains("seafood", "foo")

is the same as the following JavaScript code (`Run Online <http://www.webtoolkitonline.com/javascript-tester.html>`__):

.. code-block:: javascript

  "seafood".indexOf("foo") != -1

.. adsu:: 3

----

References:

.. [1] | `javascript string indexof - Google search <https://www.google.com/search?q=javascript+string+indexof>`_
       | `javascript string indexof - DuckDuckGo search <https://duckduckgo.com/?q=javascript+string+indexof>`_
       | `javascript string indexof - Ecosia search <https://www.ecosia.org/search?q=javascript+string+indexof>`_
       | `javascript string indexof - Qwant search <https://www.qwant.com/?q=javascript+string+indexof>`_
       | `javascript string indexof - Bing search <https://www.bing.com/search?q=javascript+string+indexof>`_
       | `javascript string indexof - Yahoo search <https://search.yahoo.com/search?p=javascript+string+indexof>`_
       | `javascript string indexof - Baidu search <https://www.baidu.com/s?wd=javascript+string+indexof>`_
       | `javascript string indexof - Yandex search <https://www.yandex.com/search/?text=javascript+string+indexof>`_

.. [2] `JavaScript String indexOf() Method <https://www.w3schools.com/jsref/jsref_indexof.asp>`_
.. [3] `func Index - Package strings - The Go Programming Language <https://golang.org/pkg/strings/#Index>`_

.. _String.prototype.indexOf(): https://www.w3schools.com/jsref/jsref_indexof.asp
.. _String.prototype.indexOf() != -1: https://www.w3schools.com/jsref/jsref_indexof.asp
.. _strings.Index: https://golang.org/pkg/strings/#Index
.. _strings.Contains: https://golang.org/pkg/strings/#Contains

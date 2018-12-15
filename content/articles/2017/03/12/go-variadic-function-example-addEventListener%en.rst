[Golang] Variadic Function Example - addEventListener
#####################################################

:date: 2017-03-12T02:42+08:00
:tags: Go, Golang, Variadic Function, GopherJS, Go to JavaScript
:category: Go
:summary: Variadic Function Example in Go_ - Implement addEventListener_ method
          via GopherJS_.
:adsu: yes


GopherJS_ [2]_ is a compiler from Go_ to JavaScript_, which helps you write
front-end Go_ code running in browsers.

To register event listeners, we can use the following code to call
addEventListener_ method:

.. code-block:: go

  foo := js.Global.Get("document").Call("getElementById", "foo")
  // register onclick event
  foo.Call("addEventListener", "click", func(event *js.Object) {
    // do something
  }, false)
  // or
  foo.Call("addEventListener", "click", func(event *js.Object) {
    // do something
  })

You can pass two or three arguments to addEventListener method. But the above
code is really ugly. We want to wrtie the code in idiomatic Go way, so I
implement the addEventListener as a Go `variadic function`_, which accepts two
or more arguments.

The implementation of addEventListener in Go_ via GopherJS is as follows:

.. code-block:: go

  func (o *Object) AddEventListener(t string, listener func(Event), args ...interface{}) {
  	if len(args) == 1 {
  		o.Call("addEventListener", t, listener, args[0])
  	} else {
  		o.Call("addEventListener", t, listener)
  	}
  }

.. adsu:: 2

where the type *Object* and *Event* is as follows:

.. code-block:: go

  type Object struct {
  	*js.Object
  }

  type Event struct {
  	*js.Object
  }

JavaScript object is ``*js.Object`` in GopherJS, so the type *Object* and
*Event* are in fact the wrapper for ``*js.Object``.

Now we can register the event listeners in the following idiomatic Go way:

.. code-block:: go

  foo.AddEventListener("click", func(e Event) {
    // do something
  })
  // or
  foo.AddEventListener("click", func(e Event) {
    // do something
  }, false)

You can pass more than three arguments, but the additional arguments after third
argument will not be used in my implementation.

.. adsu:: 3

----

Tested on: ``Ubuntu Linux 16.10``, ``Go 1.8``

----

References:

.. [1] | `variadic golang - Google search <https://www.google.com/search?q=variadic+golang>`_
       | `variadic golang - DuckDuckGo search <https://duckduckgo.com/?q=variadic+golang>`_
       | `variadic golang - Ecosia search <https://www.ecosia.org/search?q=variadic+golang>`_
       | `variadic golang - Qwant search <https://www.qwant.com/?q=variadic+golang>`_
       | `variadic golang - Bing search <https://www.bing.com/search?q=variadic+golang>`_
       | `variadic golang - Yahoo search <https://search.yahoo.com/search?p=variadic+golang>`_
       | `variadic golang - Baidu search <https://www.baidu.com/s?wd=variadic+golang>`_
       | `variadic golang - Yandex search <https://www.yandex.com/search/?text=variadic+golang>`_

.. [2] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [3] `[Golang] GopherJS Synonyms with JavaScript <{filename}../../../2016/01/29/go-gopherjs-synonyms-with-javascript%en.rst>`_

.. [4] `GitHub - siongui/godom: Make DOM manipultation in Go as similar to JavaScript as possible. (via GopherJS) <https://github.com/siongui/godom>`_

.. [5] `variadic function looping through channels : golang <https://www.reddit.com/r/golang/comments/6fv0r9/variadic_function_looping_through_channels/>`_
.. [6] `The Go variadic function, comparing errors in Golang, solving triangles & more : golang <https://old.reddit.com/r/golang/comments/a6dvwt/the_go_variadic_function_comparing_errors_in/>`_

.. _Go: https://golang.org/
.. _addEventListener: https://www.google.com/search?q=addEventListener
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _GopherJS: http://www.gopherjs.org/
.. _variadic function: https://www.google.com/search?q=variadic+function
.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js

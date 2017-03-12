[Golang] Pass Slice or Array to Variadic Function
#################################################

:date: 2017-03-13T05:16+08:00
:tags: Go, Golang, Variadic Function
:category: Go
:summary: Pass slice or array as variadic (...) parameter in Go_ programming
          language.
:adsu: yes

.. contents::

What is ... in Go Function Parameter? [2]_
++++++++++++++++++++++++++++++++++++++++++

dot dot dot (...) in final parameter of Go function means you can pass zero,
one, or more arguments for that parameter according to
`Function types in Go spec`_, and the function of such kind is usually called
`variadic function`_.

For example, assume we have the following function:

.. code-block:: go

  func giveMeString(s ...string)

We can invoke the function by:

.. code-block:: go

  giveMeString()
  giveMeString("hello")
  giveMeString("hello", "world")

The function can be invoked with more *string* arguments as you like. You can
`run above example in Go Playground`_ to fiddle the code by yourself.

One more well-known example is fmt.Println_ in Go Standard library. You can also
try it in `Go Playground`_.

.. adsu:: 2


Pass Slice or Array as Variadic Parameter [1]_
++++++++++++++++++++++++++++++++++++++++++++++

Continue the above example. Assune we have a slice of *[]string* as follows:

.. code-block:: go

  myStrSlice := []string{"hello", "world"}

How to pass the slice to the following function?

.. code-block:: go

  func giveMeString(s ...string)

Answer: append ... after the name of the slice.

.. code-block:: go

  giveMeString(myStrSlice...)

The same if we have the following array:

.. code-block:: go

  myStrArray := [2]string{"hello", "world"}

Pass it to the above variadic function via:

.. code-block:: go

  giveMeString(myStrArray[:]...)

You can try to `pass slice or array of above example on Go Playground`_.

.. adsu:: 3

----

Tested on: The `Go Playground`_

----

References:

.. [1] | `Passing arguments to ... parameters - The Go Programming Language Specification - The Go Programming Language <https://golang.org/ref/spec#Passing_arguments_to_..._parameters>`_
       | `go - How can I pass a slice as a variadic input? - Stack Overflow <http://stackoverflow.com/a/23724092>`_
.. [2] | `Function types - The Go Programming Language Specification - The Go Programming Language <https://golang.org/ref/spec#Function_types>`_
       | `go - dot dot dot in Golang. interface with empty braces - Stack Overflow <http://stackoverflow.com/a/23669857>`_
.. [3] `Go Slices: usage and internals - The Go Blog <https://blog.golang.org/go-slices-usage-and-internals>`_
.. [4] `[Golang] Variadic Function Example - addEventListener <{filename}../12/go-variadic-function-example-addEventListener%en.rst>`_
.. [5] `Go語言 函數參數點點點(...)意義 <{filename}../../02/07/go-function-argument-dot-dot-dot%zh.rst>`_

.. _Go: https://golang.org/
.. _Function types in Go spec: https://golang.org/ref/spec#Function_types
.. _variadic function: https://www.google.com/search?q=variadic+function
.. _run above example in Go Playground: https://play.golang.org/p/Qk2A6LbZfQ
.. _fmt.Println: https://golang.org/pkg/fmt/#Println
.. _Go Playground: https://play.golang.org/
.. _pass slice or array of above example on Go Playground: https://play.golang.org/p/UkXlDi0Lou

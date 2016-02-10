[Golang] Call a Struct and its Method by Name
#############################################

:date: 2016-02-11T02:58+08:00
:tags: Go, Golang, run-time reflection
:category: Go
:summary: Call a function (method_) of a struct_ by name during run-time in Go_.
          (`run-time reflection`_)

Call a function (method_) of a struct_ by name during run-time in Golang_.
(`run-time reflection`_)

`Run Code on Go Playground <https://play.golang.org/p/xeKdS7sh8E>`_

.. code-block:: go

  package main

  import (
          "fmt"
          "math"
          "reflect"
  )

  type Circle struct {
          r float64
  }

  func (c *Circle) Area() float64 {
          return math.Pi * c.r * c.r
  }

  func main() {
          c := Circle{1.2}

          // call Area() method
          fmt.Println(c.Area())

          // call Area() method by Name
          v := reflect.ValueOf(&c).MethodByName("Area").Call([]reflect.Value{})
          fmt.Println(v[0].Float())
  }

----

References:

.. [1] `golang reflect method example <https://www.google.com/search?q=golang+reflect+method+example>`_

.. [2] `Call a Struct and its Method by name in Go? - Stack Overflow <http://stackoverflow.com/questions/8103617/call-a-struct-and-its-method-by-name-in-go>`_

.. [3] `golang reflect example <https://www.google.com/search?q=golang+reflect+example>`_

.. [4] `python reflection class name <https://www.google.com/search?q=python+reflection+class+name>`_

.. [5] `introspection - Getting the class name of an instance in Python - Stack Overflow <http://stackoverflow.com/questions/510972/getting-the-class-name-of-an-instance-in-python>`_

.. [6] `reflect - The Go Programming Language <https://golang.org/pkg/reflect/>`_

.. [7] `The Laws of Reflection - The Go Blog <http://blog.golang.org/laws-of-reflection>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _struct: https://tour.golang.org/moretypes/2
.. _method: https://tour.golang.org/methods/1
.. _run-time reflection: http://blog.golang.org/laws-of-reflection

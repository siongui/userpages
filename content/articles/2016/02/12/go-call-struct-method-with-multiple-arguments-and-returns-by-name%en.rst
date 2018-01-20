[Golang] Call Struct Method With Multiple Arguments And Returns by Name
#######################################################################

:date: 2016-02-12T03:41+08:00
:tags: Go, Golang, run-time reflection
:category: Go
:summary: Call a function (method_), with multiple arguments and returns, of a
          struct_ by name during run-time in Go_. (`run-time reflection`_)
:adsu: yes

Call a function (method_), with multiple arguments and returns, of a struct_ by
name during run-time in Golang_. (`run-time reflection`_)

This is an advanced example of [1]_.

`Run Code on Go Playground <https://play.golang.org/p/y8rpCTmUdX>`_

.. code-block:: go

  package main

  import (
          "fmt"
          "reflect"
  )

  type DemoStruct struct {
          s string
          i int
          f float64
  }

  func (d *DemoStruct) DemoFunc(s string, i int, f float64) (string, int, float64, DemoStruct) {
          return s, i, f, *d
  }

  func main() {
          ds := DemoStruct{"sacca", 1, 3.14}

          // call DemoFunc() method
          fmt.Println(ds.DemoFunc("dukkha", 0, 1.414))

          // call DemoFunc() method by Name
          retv := reflect.ValueOf(&ds).MethodByName("DemoFunc").Call([]reflect.Value{
                  reflect.ValueOf("dukkha"),
                  reflect.ValueOf(13),
                  reflect.ValueOf(1.414),
          })
          fmt.Println(retv[0].String())
          fmt.Println(retv[0].Interface().(string))
          fmt.Println(retv[1].Int())
          fmt.Println(retv[1].Interface().(int))
          fmt.Println(retv[2].Float())
          fmt.Println(retv[2].Interface().(float64))
          fmt.Println(retv[3].Interface().(DemoStruct))
  }

.. adsu:: 2

----

References:

.. [1] `[Golang] Call a Struct and its Method by Name <{filename}../11/go-call-a-struct-and-its-method-by-name%en.rst>`_

.. [2] `golang reflect method example <https://www.google.com/search?q=golang+reflect+method+example>`_

.. [3] `Call a Struct and its Method by name in Go? - Stack Overflow <http://stackoverflow.com/questions/8103617/call-a-struct-and-its-method-by-name-in-go>`_

.. [4] `golang reflect example <https://www.google.com/search?q=golang+reflect+example>`_

.. [5] `python reflection class name <https://www.google.com/search?q=python+reflection+class+name>`_

.. [6] `introspection - Getting the class name of an instance in Python - Stack Overflow <http://stackoverflow.com/questions/510972/getting-the-class-name-of-an-instance-in-python>`_

.. [7] `reflect - The Go Programming Language <https://golang.org/pkg/reflect/>`_

.. [8] `The Laws of Reflection - The Go Blog <http://blog.golang.org/laws-of-reflection>`_

.. [9] `谈一谈Go的interface和reflect | legendtkl <http://legendtkl.com/2015/11/28/go-interface-reflect/>`_

.. [10] `为什么golang不能通过字符串来创建对象实例？ - Go 语言 - 知乎 <https://www.zhihu.com/question/25580049>`_

.. [11] `Go 反射实践及剖析 - 文章 - 伯乐在线 <http://blog.jobbole.com/108601/>`_
.. [12] `Could anyone help me with some reflection voodoo please? : golang <https://www.reddit.com/r/golang/comments/66qwet/could_anyone_help_me_with_some_reflection_voodoo/>`_
.. [13] | `Go Reflection: Creating Objects from Types — Part I (Primitive Types) : golang <https://www.reddit.com/r/golang/comments/7l4boc/go_reflection_creating_objects_from_types_part_i/>`_
        | `Go Reflection: Creating Objects from Types — Part I (Primitive Types) <https://medium.com/kokster/go-reflection-creating-objects-from-types-part-i-primitive-types-6119e3737f5d>`_
.. [14] | `Learning to Use Go Reflection – Capital One DevExchange – Medium : golang <https://www.reddit.com/r/golang/comments/7rjt1h/learning_to_use_go_reflection_capital_one/>`_
        | `Learning to Use Go Reflection – Capital One DevExchange – Medium <https://medium.com/capital-one-developers/learning-to-use-go-reflection-822a0aed74b7>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _struct: https://tour.golang.org/moretypes/2
.. _method: https://tour.golang.org/methods/1
.. _run-time reflection: http://blog.golang.org/laws-of-reflection

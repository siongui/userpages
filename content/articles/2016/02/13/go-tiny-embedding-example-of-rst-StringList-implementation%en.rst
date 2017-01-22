[Golang] Tiny Embedding Example of reStructuredText StringList Implementation
#############################################################################

:date: 2016-02-13T02:34+08:00
:tags: Go, Golang, Data Structure
:category: Go
:summary: Tiny Go_ embedding_ code example of analogy to StringList_
          implementation of reStructuredText_.
:adsu: yes

Tiny Golang_ embedding_ code example of analogy to StringList_ implementation
of reStructuredText_.

`Run Code on Go Playground <https://play.golang.org/p/5JaL_prBtA>`__

.. code-block:: go

  package main

  import "fmt"

  type ParentList struct {
          data   []int
          parent *ParentList
  }

  func NewParentList(data []int, parent *ParentList) ParentList {
          return ParentList{
                  data:   data,
                  parent: parent,
          }
  }

  func (p *ParentList) GetSlice(start, stop int) ParentList {
          return NewParentList(p.data[start:stop], p)
  }

  type ChildList struct {
          ParentList
          parent *ChildList
  }

  func NewChildList(data []int, parent *ChildList) ChildList {
          return ChildList{
                  ParentList: NewParentList(data, nil),
                  parent:     parent,
          }
  }

  func (c *ChildList) GetSlice(start, stop int) ChildList {
          return NewChildList(c.data[start:stop], c)
  }

  func main() {
          // Parent List
          b := NewParentList([]int{0, 1, 2, 3, 4, 5, 6, 7}, nil)
          bs := b.GetSlice(2, 5)
          fmt.Println(bs)
          fmt.Println(bs.parent)

          // Child List
          c := NewParentList([]int{7, 6, 5, 4, 3, 2, 1, 0}, nil)
          cs := c.GetSlice(2, 5)
          fmt.Println(cs)
          fmt.Println(cs.parent)
  }

----

References:

.. [1] `golang class inheritance <https://www.google.com/search?q=golang+class+inheritance>`_

.. [2] `Golang concepts from an OOP point of view - GitHub <https://github.com/luciotato/golang-notes/blob/master/OOP.md>`_

.. [3] `golang oop <https://www.google.com/search?q=golang+oop>`_

.. [4] `golang subclass <https://www.google.com/search?q=golang+subclass>`_

.. [5] `golang inheritance override <https://www.google.com/search?q=golang+inheritance+override>`_

.. [6] `golang embedding <https://www.google.com/search?q=golang+embedding>`_

.. [7] `golang interface inheritance <https://www.google.com/search?q=golang+interface+inheritance>`_

.. [8] `golang method overloading <https://www.google.com/search?q=golang+method+overloading>`_

.. [9] `Python docutils.statemachine.StringList Examples <http://www.programcreek.com/python/example/59097/docutils.statemachine.StringList>`_

.. [10] `docutils.statemachine.StringList - Nullege Python Samples <http://nullege.com/codes/search/docutils.statemachine.StringList>`_

.. [11] `another tiny (bad) example on Go Playground <https://play.golang.org/p/Z4rbNGff_d>`_

.. [12] `Go Object Oriented Design <https://nathany.com/good/>`_
        (`reddit <https://www.reddit.com/r/golang/comments/4bn34e/go_object_oriented_design/>`__)


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _embedding: https://www.google.com/search?q=golang+embedding
.. _reStructuredText: http://docutils.sourceforge.net/rst.html
.. _StringList: https://www.google.com/search?q=StringList+rst

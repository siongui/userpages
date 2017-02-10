[Golang] Synonyms with Python Class
###################################

:date: 2016-02-10T21:34+08:00
:tags: Go, Golang, Python, Data Structure
:category: Go
:summary: Synonyms - Python_ class_ translated to Go_
:adsu: yes

Simple example of writing Python_ class_ synonyms in Golang_.

.. list-table:: Python_ vs Golang_
   :header-rows: 1
   :class: table-syntax-diff

   * - Python_
     - Go_

   * - `Run Code on repl.it <https://repl.it/Bl4e/0>`_

       .. code-block:: python

         class Square:
             def __init__(self, x, y):
                 self.x = x
                 self.y = y

             def Area(self):
                 return self.x * self.y

         if __name__ == "__main__":
             sq = Square(2, 3)
             print(sq.Area())


     - `Run Code on Go Playground <https://play.golang.org/p/CToo3Co4Ta>`_

       .. code-block:: go

         package main

         import "fmt"

         type Square struct {
                 x int
                 y int
         }

         func (s *Square) Area() int {
                 return s.x * s.y
         }

         func main() {
                 // no constructor in Golang
                 sq := Square{2, 3}

                 // calculate area
                 fmt.Println(sq.Area())
         }

.. adsu:: 2

----

References:

.. [1] `golang class inheritance <https://www.google.com/search?q=golang+class+inheritance>`_

.. [2] `Golang concepts from an OOP point of view - GitHub <https://github.com/luciotato/golang-notes/blob/master/OOP.md>`_

.. [3] `Go Object Oriented Design <https://nathany.com/good/>`_
       (`reddit <https://www.reddit.com/r/golang/comments/4bn34e/go_object_oriented_design/>`__)


.. _Python: https://www.python.org/
.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _class: https://docs.python.org/2/tutorial/classes.html

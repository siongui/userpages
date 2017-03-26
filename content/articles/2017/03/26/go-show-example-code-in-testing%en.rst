[Golang] Write Example Code in Testing
######################################

:date: 2017-03-26T23:09+08:00
:tags: Go, Golang
:category: Go
:summary: Demonstrate the usage of your methods by writing example code in the
          testing of your package in Go programming language.
:adsu: yes


Demonstrate the usage of your methods by writing example code in the testing of
your package in Go programming language.

Wrong Way
+++++++++

When I want to show how to use Go methods in my blog posts, I usually put the
example code in the testing as follows:

**Method**: *hello.go*

.. code-block:: go

  package hello

  import (
  	"fmt"
  )

  func Hello() {
  	fmt.Println("hello world")
  }

**Usage**: *hello_test.go*

.. code-block:: go

  package hello

  import "testing"

  func TestHello(t *testing.T) {
  	// Example code here
  }

One day `@dchapes`_ told me `this is not a test`_, and `@shurcooL`_ showed me
how to write examples in idiomatic Go way.

.. adsu:: 2

Idiomatic Go Way
++++++++++++++++

`@shurcooL`_ told to `use Examples in testing package`_ [1]_. The code in above
`Wrong Way`_ section becomes:

**Method**: *hello.go*

.. code-block:: go

  package hello

  import (
  	"fmt"
  )

  func Hello() {
  	fmt.Println("hello world")
  }

**Usage**: *hello_test.go*

.. code-block:: go

  package hello

  import "testing"

  func ExampleHello(t *testing.T) {
  	// Example code here
  }

The standard Go testing_ package can runs and verifies example code if we follow
the conventions in the testing code. This is a better and idiomatic way to show
others how to use the methods in your package or code.

.. adsu:: 3

----

References:

.. [1] `Examples - testing - The Go Programming Language <https://golang.org/pkg/testing/#hdr-Examples>`_

.. _@dchapes: https://github.com/dchapes
.. _this is not a test: https://github.com/siongui/userpages/commit/77cd55346752ccaa2efa44b9084e97af81b664dd#commitcomment-21400283
.. _@shurcooL: https://github.com/shurcooL
.. _use Examples in testing package: https://github.com/siongui/userpages/commit/77cd55346752ccaa2efa44b9084e97af81b664dd#commitcomment-21401415
.. _testing: https://golang.org/pkg/testing/

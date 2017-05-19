[Golang] Integer Exponentiation
###############################

:date: 2017-05-20T03:17+08:00
:tags: Go, Golang, Algorithm, Math
:category: Go
:summary: Integer exponentiation in Go programming language.
:og_image: http://image.mathcaptain.com/cms/images/39/laws-of-exponents.JPG
:adsu: yes

Integer exponentiation in Go programming language.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/ljM_W1e6FZ>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  )

  // return a^n
  func Power(a, n uint) uint {
  	var i, result uint
  	result = 1
  	for i = 0; i < n; i++ {
  		result *= a
  	}
  	return result
  }

  func main() {
  	fmt.Println(Power(2, 0))
  	fmt.Println(Power(2, 1))
  	fmt.Println(Power(2, 2))
  	fmt.Println(Power(2, 3))
  }

.. adsu:: 2

Tested on: `Go Playground`_

----

References:

.. [1] | `golang integer power - Google search <https://www.google.com/search?q=golang+integer+power>`_
       | `golang integer power - DuckDuckGo search <https://duckduckgo.com/?q=golang+integer+power>`_
       | `golang integer power - Ecosia search <https://www.ecosia.org/search?q=golang+integer+power>`_
       | `golang integer power - Qwant search <https://www.qwant.com/?q=golang+integer+power>`_
       | `golang integer power - Bing search <https://www.bing.com/search?q=golang+integer+power>`_
       | `golang integer power - Yahoo search <https://search.yahoo.com/search?p=golang+integer+power>`_
       | `golang integer power - Baidu search <https://www.baidu.com/s?wd=golang+integer+power>`_
       | `golang integer power - Yandex search <https://www.yandex.com/search/?text=golang+integer+power>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/

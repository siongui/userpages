[Golang] Type Conversion between String and Integer
###################################################

:date: 2017-04-14T21:01+08:00
:tags: Go, Golang, Type Casting, Type Conversion
:category: Go
:summary: Go type conversion (type casting) between string and integer number.
:og_image: http://2.bp.blogspot.com/-O2uOFWMD_zs/Ve3Iq9rfX_I/AAAAAAAAPxE/ZKRsGxscEsY/s1600/go_lang_mascot_by_kirael_art-d7kunhu.gif
:adsu: yes

Example of Go type conversion (type casting) between *string* and integer
(*int64*).

Problem
+++++++

Assume we have the following variable of type *string*:

.. code-block:: go

  intNum := "100"

We will multiply the intNum by 2, convert the result back to *string*.

Solution
++++++++

- Convert *string* to *int64*:
    * Use strconv.ParseInt_
    * `Run Example on Go Playground <https://play.golang.org/p/cPdqJ-2yc->`__

- Convert *int64* to *string*:
    * Use strconv.FormatInt_
    * `Run Example on Go Playground <https://play.golang.org/p/xbwngzA7Mb>`__

.. adsu:: 2

The following code convert *string* to *int64*, multiply the number by 2, and
then convert the result from *int64* back to *string*.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/bvaoIu-1uS>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  	"strconv"
  )

  func main() {
  	intNum := "100"

  	// convert string to int
  	n, err := strconv.ParseInt(intNum, 10, 0)
  	if err != nil {
  		fmt.Println(err)
  		return
  	}

  	// convert int to string
  	result := strconv.FormatInt(n*2, 10)

  	fmt.Println("result: ", result)
  }

.. adsu:: 3

----

References:

.. [1] | `golang string to int - Google search <https://www.google.com/search?q=golang+string+to+int>`_
       | `golang string to int - DuckDuckGo search <https://duckduckgo.com/?q=golang+string+to+int>`_
       | `golang string to int - Ecosia search <https://www.ecosia.org/search?q=golang+string+to+int>`_
       | `golang string to int - Qwant search <https://www.qwant.com/?q=golang+string+to+int>`_
       | `golang string to int - Bing search <https://www.bing.com/search?q=golang+string+to+int>`_
       | `golang string to int - Yahoo search <https://search.yahoo.com/search?p=golang+string+to+int>`_
       | `golang string to int - Baidu search <https://www.baidu.com/s?wd=golang+string+to+int>`_
       | `golang string to int - Yandex search <https://www.yandex.com/search/?text=golang+string+to+int>`_
.. [2] `func ParseInt - strconv - The Go Programming Language <https://golang.org/pkg/strconv/#ParseInt>`_

.. [3] | `golang int to string - Google search <https://www.google.com/search?q=golang+int+to+string>`_
       | `golang int to string - DuckDuckGo search <https://duckduckgo.com/?q=golang+int+to+string>`_
       | `golang int to string - Ecosia search <https://www.ecosia.org/search?q=golang+int+to+string>`_
       | `golang int to string - Qwant search <https://www.qwant.com/?q=golang+int+to+string>`_
       | `golang int to string - Bing search <https://www.bing.com/search?q=golang+int+to+string>`_
       | `golang int to string - Yahoo search <https://search.yahoo.com/search?p=golang+int+to+string>`_
       | `golang int to string - Baidu search <https://www.baidu.com/s?wd=golang+int+to+string>`_
       | `golang int to string - Yandex search <https://www.yandex.com/search/?text=golang+int+to+string>`_
.. [4] `func FormatInt - strconv - The Go Programming Language <https://golang.org/pkg/strconv/#FormatInt>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _strconv.ParseInt: https://golang.org/pkg/strconv/#ParseInt
.. _strconv.FormatInt: https://golang.org/pkg/strconv/#FormatInt

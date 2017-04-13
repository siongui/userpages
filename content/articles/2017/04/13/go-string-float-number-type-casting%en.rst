[Golang] Type Conversion between String and Floating Number
###########################################################

:date: 2017-04-13T14:10+08:00
:tags: Go, Golang, Type Casting, Type Conversion
:category: Go
:summary: Go type conversion (type casting) between string and floating number.
:og_image: http://2.bp.blogspot.com/-O2uOFWMD_zs/Ve3Iq9rfX_I/AAAAAAAAPxE/ZKRsGxscEsY/s1600/go_lang_mascot_by_kirael_art-d7kunhu.gif
:adsu: yes

Example of Go type conversion (type casting) between *string* and floating
number (*float64*).

Problem
+++++++

Assume we have the following two variable of type *string*:

.. code-block:: go

  width := "560"
  height := "315"

We want to calculate the `aspect ratio`_, i.e., width/height, and then represent
the ratio as percents using the symbol %.

The final result will be (of type *string*):

.. code-block:: go

  56.25%

Solution
++++++++

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/q6AKBn5n-T>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  	"strconv"
  )

  func main() {
  	width := "560"
  	height := "315"

  	// convert string to float64
  	w, err := strconv.ParseFloat(width, 64)
  	if err != nil {
  		fmt.Println(err)
  	}
  	h, err := strconv.ParseFloat(height, 64)
  	if err != nil {
  		fmt.Println(err)
  	}

  	// make calculation and then convert float64 to string
  	aspectRatio := strconv.FormatFloat(h*100/w, 'f', -1, 64) + "%"

  	fmt.Println("aspect ratio: ", aspectRatio)
  }

.. adsu:: 2

----

References:

.. [1] | `golang string to float - Google search <https://www.google.com/search?q=golang+string+to+float>`_
       | `golang string to float - DuckDuckGo search <https://duckduckgo.com/?q=golang+string+to+float>`_
       | `golang string to float - Ecosia search <https://www.ecosia.org/search?q=golang+string+to+float>`_
       | `golang string to float - Qwant search <https://www.qwant.com/?q=golang+string+to+float>`_
       | `golang string to float - Bing search <https://www.bing.com/search?q=golang+string+to+float>`_
       | `golang string to float - Yahoo search <https://search.yahoo.com/search?p=golang+string+to+float>`_
       | `golang string to float - Baidu search <https://www.baidu.com/s?wd=golang+string+to+float>`_
       | `golang string to float - Yandex search <https://www.yandex.com/search/?text=golang+string+to+float>`_
.. [2] `func ParseFloat - strconv - The Go Programming Language <https://golang.org/pkg/strconv/#ParseFloat>`_

.. [3] | `golang float to string - Google search <https://www.google.com/search?q=golang+float+to+string>`_
       | `golang float to string - DuckDuckGo search <https://duckduckgo.com/?q=golang+float+to+string>`_
       | `golang float to string - Ecosia search <https://www.ecosia.org/search?q=golang+float+to+string>`_
       | `golang float to string - Qwant search <https://www.qwant.com/?q=golang+float+to+string>`_
       | `golang float to string - Bing search <https://www.bing.com/search?q=golang+float+to+string>`_
       | `golang float to string - Yahoo search <https://search.yahoo.com/search?p=golang+float+to+string>`_
       | `golang float to string - Baidu search <https://www.baidu.com/s?wd=golang+float+to+string>`_
       | `golang float to string - Yandex search <https://www.yandex.com/search/?text=golang+float+to+string>`_
.. [4] `func FormatFloat - strconv - The Go Programming Language <https://golang.org/pkg/strconv/#FormatFloat>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _aspect ratio: https://www.google.com/search?q=aspect+ratio

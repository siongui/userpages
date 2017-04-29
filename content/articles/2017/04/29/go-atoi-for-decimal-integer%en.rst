[Golang] Atoi for Decimal Integer
#################################

:date: 2017-04-29T22:35+08:00
:tags: Go, Golang, Go strconv Package, String Manipulation
:category: Go
:summary: My implementation of *atoi* function, which converts a string
          representing **decimal number** to its corresponding *int* type,
          in Go_ programmming language.
:og_image: http://www.unixstickers.com/image/cache/data/stickers/golang/Go-brown-side.sh-600x600.png
:adsu: yes


Introduction
++++++++++++

I read the post of *Write your own atoi()* in GeeksforGeeks [1]_, and feel it is
a good exercise to implement my own atoi_ in Go_, that takes a string (which
represents an decimal integer) as an argument and returns its value in *int*
type.

Go standard library already provides this function, see strconv.Atoi_ for more
information.


atoi Implementation
+++++++++++++++++++

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/HGRXG2Vv61>`_
      :class: align-center

.. code-block:: go

  import (
  	"errors"
  )

  var m = map[rune]int{
  	'0': 0,
  	'1': 1,
  	'2': 2,
  	'3': 3,
  	'4': 4,
  	'5': 5,
  	'6': 6,
  	'7': 7,
  	'8': 8,
  	'9': 9,
  }

  // input: string of decimal integer number
  func Atoi(ds string) (di int, err error) {
  	if len(ds) == 0 {
  		err = errors.New("input length is zero")
  		return
  	}

  	sign := 1
  	if ds[0] == '-' {
  		sign = -1
  		ds = ds[1:]
  		if len(ds) == 0 {
  			err = errors.New("invalid input: -")
  			return
  		}
  	}

  	for _, digit := range ds {
  		if d, ok := m[digit]; ok {
  			di = di*10 + d
  		} else {
  			err = errors.New("invalid digit: " + string(digit))
  			return
  		}
  	}

  	di = di * sign
  	return
  }

.. adsu:: 2

----

Tested on: `Go Playground`_

----

References
++++++++++

.. [1] `Write your own atoi() - GeeksforGeeks <http://www.geeksforgeeks.org/write-your-own-atoi/>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
.. _strconv.Atoi: https://golang.org/pkg/strconv/#Atoi
.. _atoi: https://www.google.com/search?q=atoi

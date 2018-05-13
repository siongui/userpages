[Golang] Pretty Print Variable (struct, map, array, slice)
##########################################################

:date: 2016-01-30T02:21+08:00
:modified: 2018-05-13T22:09+08:00
:tags: pretty print, Go, Golang, JSON
:category: Go
:summary: Pretty print variable (struct, map, array, slice) in Go.
:og_image: https://cdn-images-1.medium.com/max/600/1*i2skbfmDsHayHhqPfwt6pA.png
:adsu: yes

Pretty print variable (struct, map, array, slice) in Golang.
The easiest way is through MarshalIndent_ function in Go standard
`encoding/json`_ package.

.. code-block:: go

  import (
  	"encoding/json"
  	"fmt"
  )

  func PrettyPrint(v interface{}) (err error) {
  	b, err := json.MarshalIndent(v, "", "  ")
  	if err != nil {
  		return
  	}
  	fmt.Println(string(b))
  	return
  }

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/2IV4QM3hx87>`__
   :class: align-center

Tested on: `Go Playground`_

.. adsu:: 2

.. _MarshalIndent: https://golang.org/pkg/encoding/json/#MarshalIndent
.. _encoding/json: https://golang.org/pkg/encoding/json/
.. _Go Playground: https://play.golang.org/

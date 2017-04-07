[Golang] Pretty Print Variable (struct, map, array, slice)
##########################################################

:date: 2016-01-30T02:21+08:00
:modified: 2017-04-08T01:47+08:00
:tags: pretty print, Go, Golang, JSON
:category: Go
:summary: Pretty print variable (struct, map, array, slice) in Go_.
:adsu: yes

Pretty print variable (struct, map, array, slice) in Golang_.
The easiest way is through MarshalIndent_ function in json_ package.

.. code-block:: go

  import "encoding/json"

  func PrettyPrint(v interface{}) {
  	b, _ := json.MarshalIndent(v, "", "  ")
  	println(string(b))
  }

`Run Example on Go Playground <https://play.golang.org/p/ZpjQoxMinU>`_

Tested on:

- ``Ubuntu Linux 15.10``, ``Go 1.5.3``.
- Go Playground

.. adsu:: 2

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _MarshalIndent: https://golang.org/pkg/encoding/json/#MarshalIndent
.. _json: https://golang.org/pkg/encoding/json/

[Golang] Pretty Print Variable (struct, map, array, slice)
##########################################################

:date: 2016-01-30T02:21+08:00
:tags: pretty print, Go, Golang
:category: Go
:summary: Pretty print variable (struct, map, array, slice) in Go_.
:adsu: yes

Pretty print variable (struct, map, array, slice) in Golang_.
The easiest way is through MarshalIndent_ function in json_ package.

.. code-block:: go

  import "encoding/json"

  // assume we want to pretty print *mystruct*
  b, _ := json.MarshalIndent(mystruct, "", "  ")
  println(string(b))


Tested on: ``Ubuntu Linux 15.10``, ``Go 1.5.3``.

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _MarshalIndent: https://golang.org/pkg/encoding/json/#MarshalIndent
.. _json: https://golang.org/pkg/encoding/json/

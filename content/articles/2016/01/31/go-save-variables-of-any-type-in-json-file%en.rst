[Golang] Save Variables of Any Type in JSON file
################################################

:date: 2016-01-31T02:04+08:00
:tags: Go, Golang, JSON
:category: Go
:summary: Save variables of any type (struct, map, array, slice, etc.) in JSON_
          format file via Go_ programming language.
:adsu: yes


Save variables of any type (struct, map, array, slice, etc.) in JSON_ format
file via Golang_.

.. code-block:: go

  import "os"
  import "encoding/json"

  func SaveJsonFile(v interface{}, path string) {
          fo, err := os.Create(path)
          if err != nil {
                  panic(err)
          }
          defer fo.Close()
          e := json.NewEncoder(fo)
          if err := e.Encode(v); err != nil {
                  panic(err)
          }
  }

  // assume we want to save *mystruct* in ``/home/myaccount/mystruct.json``
  SaveJsonFile(mystruct, "/home/myaccount/mystruct.json")


Tested on: ``Ubuntu Linux 15.10``, ``Go 1.5.3``.

.. adsu:: 2

----

References:

.. [1] `json - The Go Programming Language <https://golang.org/pkg/encoding/json/>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _JSON: http://json.org/

[Golang] Save JSON Data in Directory
####################################

:date: 2015-03-28 19:59
:tags: Go, Golang, JSON
:category: Go
:summary: Save JSON-format data in the directory using Go programming language.
:adsu: yes


This post shows how to convert data to JSON_ format and save the data in the
directory, and if the directory does not exist, it will be created first.

Souce Code
++++++++++

.. show_github_file:: siongui userpages content/code/go-save-json-in-folder/savelink.go

.. adsu:: 2

Test
++++

.. show_github_file:: siongui userpages content/code/go-save-json-in-folder/savelink_test.go


Tested on: ``Ubuntu Linux 14.10``, ``Go 1.4``.

----

References:

.. [1] `go - How to check whether a file or directory denoted by a path exists in golang? - Stack Overflow <http://stackoverflow.com/questions/10510691/how-to-check-whether-a-file-or-directory-denoted-by-a-path-exists-in-golang>`_

.. [2] `go - os.MkDir and os.MkDirAll permission value? - Stack Overflow <http://stackoverflow.com/questions/14249467/os-mkdir-and-os-mkdirall-permission-value>`_

.. [3] `How to efficiently concatenate strings in Go? - Stack Overflow <http://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go>`_


.. _JSON: http://json.org/

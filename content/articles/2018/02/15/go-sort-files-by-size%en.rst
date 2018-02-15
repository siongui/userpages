[Golang] Sort Files by Size
###########################

:date: 2018-02-15T09:46+08:00
:tags: Go, Golang, Go sort
:category: Go
:summary: Sort files by size in Go. Use *sort.Slice* on the slice of
          *os.FileInfo* returned by *ioutil.ReadDir*.
:og_image: https://pbs.twimg.com/media/CxQV-tmUUAEDvzF.jpg
:adsu: yes

Sort files in a directory by size in Go. The steps are:

1. Use `ioutil.ReadDir`_ to get files in a directory (not including
   sub-directories).
2. Use `sort.Slice`_ to sort the files by size.

The following are complete source code.

.. show_github_file:: siongui userpages content/code/go/file-sort-by-size/sortfile.go
.. adsu:: 2
.. show_github_file:: siongui userpages content/code/go/file-sort-by-size/sortfile_test.go

If you want to know how to sort prior to Go 1.8, see [1]_.

If you want to find all files in a directory and all its sub-directories, see
[2]_.

Tested on:

- ``Ubuntu Linux 17.10``, ``Go 1.9.3``

.. - `Go Playground`_

----

References:

.. adsu:: 2
.. [1] `[Golang] Sort String by Character <{filename}/articles/2017/05/07/go-sort-string-slice-of-rune%en.rst>`_
.. [2] `[Golang] Walk All Files in Directory <{filename}/articles/2016/02/04/go-walk-all-files-in-directory%en.rst>`_

.. _Go: https://golang.org/
.. _Go Playground: https://play.golang.org/
.. _sort.Slice: https://golang.org/pkg/sort/#Slice
.. _ioutil.ReadDir: https://golang.org/pkg/io/ioutil/#ReadDir

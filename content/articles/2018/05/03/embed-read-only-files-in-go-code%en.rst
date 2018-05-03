Embed Read-Only Files in Go Code
################################

:date: 2018-05-03T23:49+08:00
:tags: Go, Golang, GopherJS, Full-Stack Go, Full-Stack Golang,
       Embed File in Go Code
:category: Go
:summary: Embed read-only files in Go code.
:og_image: https://newrelic.com/assets/pages/golang/go-mascot.svg
:adsu: yes


The motivation is to embed data files of succinct trie in Go code and compiled
to JavaScript via GopherJS_. But the method here can not only used in frontend,
but also in local Go program. There are many third-party Go packages that can
help you to do the same thing, some of them can even let you write content to
the files. But they are too complicated for me. I just need a simple package
which embed files in Go code, and read the files only (no need to write files).
So I write this simple tool *goef* [1]_ to help me.

The features are:

- Files are read-only.
- Used in front-end code via GopherJS_, or local Go program.
- Can be included in your Go package, or put in a separate package.

The idea of *goef* is simple. Encode the content of the files in Base64 format
and put the (name, content) of the files in the (key, value) of Go built-in map_
structure. Then implement a *ReadFile* method which has the same usage as
`ioutil.ReadFile`_ in Go standard library.

For usage details, please visit GitHub repo.

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.10.2``.

**References**:

.. [1] `GitHub - siongui/goef: Embed file in your Go code <https://github.com/siongui/goef>`_

.. _GopherJS: https://github.com/gopherjs/gopherjs
.. _map: https://blog.golang.org/go-maps-in-action
.. _ioutil.ReadFile: https://golang.org/pkg/io/ioutil/#ReadFile

Pass Command Line Arguments (Flags) in Go Test
##############################################

:date: 2017-04-28T03:08+08:00
:tags: Go, Golang, Commandline, Go flag Package, Go testing
:category: Go
:summary: Pass command line arguments (flags) in Go test.
:og_image: https://newrelic.com/assets/pages/golang/go-mascot.svg
:adsu: yes


This post show how to pass command line arguments (flags) in Golang test.

The answers found in Google search [1]_ and Stack Overflow [2]_ are not working
for ``Go 1.8.1``. Finally I found the issues [3]_ and figure out how to pass
arguments correctly. The following is howto.

In Go test code:

.. code-block:: go

  package goef

  import (
  	"flag"
  	"testing"
  )

  var pkgdir = flag.String("pkgdir", "", "dir of package containing embedded files")

  func TestGenerateGoPackage(t *testing.T) {
  	t.Log(*pkgdir)
  }

Note that you do not need to call `flag.Parse()`_ in the test code.

You can pass the command line arguments as follows:

.. code-block:: bash

  $ export PKGDIR=${GOPATH}/src/github.com/siongui/myvfs
  $ go test -v embed.go buildpkg.go buildpkg_test.go -args -pkgdir=${PKGDIR}

Note that if ``-args`` is not in ``go test`` command, the *pkgdir* string
variable in the test code will be empty.

.. adsu:: 2

----

**Appendix**

You can also use `environment variable`_ to pass arguments in Go test.

In Go test code:

.. code-block:: go

  package goef

  import (
  	"os"
  	"testing"
  )

  func TestGenerateGoPackage(t *testing.T) {
  	t.Log(os.Getenv("PKGDIR"))
  }

Pass arguments as follows:

.. code-block:: bash

  $ export PKGDIR=${GOPATH}/src/github.com/siongui/myvfs
  $ go test -v embed.go buildpkg.go buildpkg_test.go

.. adsu:: 3

----

Tested on: ``Ubuntu Linux 17.04``, ``Go 1.8.1``.

----

References:

.. [1] | `go test pass flags - Google search <https://www.google.com/search?q=go+test+pass+flags>`_
       | `golang test flag - Google search <https://www.google.com/search?q=golang+test+flag>`_
       | `go test custom flags - Google search <https://www.google.com/search?q=go+test+custom+flags>`_

.. [2] | `Custom command line flags in Go's unit tests - Stack Overflow <http://stackoverflow.com/questions/27342973/custom-command-line-flags-in-gos-unit-tests>`_
       | `Provide additional/custom flag to go test tool - Google Groups <https://groups.google.com/d/topic/golang-nuts/X9x4tNVqK-8>`_
       | `testing - Process command line arguments in go test - Stack Overflow <http://stackoverflow.com/questions/21350962/process-command-line-arguments-in-go-test>`_
       | `flag.Parse in tests - Google Groups <https://groups.google.com/d/topic/golang-nuts/P6EdEdgvDuc>`_

.. [3] | `cmd/go: Test tool eats known flags, docs say it doesn't · Issue #12177 · golang/go · GitHub <https://github.com/golang/go/issues/12177>`_
       | `cmd/go: add -args to 'go test' to resolve -v ambiguity (Ief9e830a) · Gerrit Code Review <https://go-review.googlesource.com/c/17775/>`_

.. [4] | `[Golang] Read Command-line Arguments Example <{filename}../../../2015/02/18/go-parse-command-line-arguments%en.rst>`_
       | `[Golang] Parse Command Line Arguments - String Variable <{filename}../../../2016/12/21/go-parse-commandline-arguments-string-variable%en.rst>`_

.. [5] `Good resources for testing in Go : golang <https://old.reddit.com/r/golang/comments/9zri71/good_resources_for_testing_in_go/>`_
.. [6] `TIL: 'testing' package has a '--short' flag : golang <https://old.reddit.com/r/golang/comments/a1iuhg/til_testing_package_has_a_short_flag/>`_
.. [7] `Dependency management in Go : golang <https://old.reddit.com/r/golang/comments/a1ycyk/dependency_management_in_go/>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _flag.Parse(): https://golang.org/pkg/flag/#Parse
.. _environment variable: https://www.google.com/search?q=environment+variable

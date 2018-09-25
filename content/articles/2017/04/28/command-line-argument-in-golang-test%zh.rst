Go語言測試中傳入命令行參數
##########################

:date: 2018-09-26T02:45+08:00
:tags: Go語言, 命令行, 測試
:category: Go語言
:summary: 在Go語言測試裡傳入命令行(command-line)參數。
:og_image: https://newrelic.com/assets/pages/golang/go-mascot.svg
:adsu: yes


本文展示該如何在Golang測試裡，傳入命令行(command-line)參數。

在Google搜尋 [1]_ 以及 Stack Overflow [2]_ 裡找到的答案無法在 ``Go 1.8.1`` 使用。
最後我發現這個問題 [3]_ 並搞清楚該如何正確地傳入參數。下面就是解答。

在Golang測試程式碼：

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

注意！你並不需要在測試程式碼裡呼叫 `flag.Parse()`_ 。

你可以以如下的方法傳入命令行參數：

.. code-block:: bash

  $ export PKGDIR=${GOPATH}/src/github.com/siongui/myvfs
  $ go test -v embed.go buildpkg.go buildpkg_test.go -args -pkgdir=${PKGDIR}

注意！如果 ``-args`` 不在 ``go test`` 命令裡，測試程式碼裡的 *pkgdir* 字串變數會是空的。

.. adsu:: 2

----

**附錄**

你也可以透過 `environment variable`_ 來在Golang測試裡傳參數。

在Golang測試碼：

.. code-block:: go

  package goef

  import (
  	"os"
  	"testing"
  )

  func TestGenerateGoPackage(t *testing.T) {
  	t.Log(os.Getenv("PKGDIR"))
  }

可以以如下的方式傳參數：

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

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _flag.Parse(): https://golang.org/pkg/flag/#Parse
.. _environment variable: https://www.google.com/search?q=environment+variable

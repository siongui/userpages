[Golang] Embed File in String Literal
#####################################

:date: 2017-04-26T22:43+08:00
:tags: Go, Golang, Embed File in Go Code
:category: Go
:summary: Embed any type of file in string literal in Go code.
:og_image: https://newrelic.com/assets/pages/golang/go-mascot.svg
:adsu: yes

Embed any type of file, no matter it is data/assets/resources/binary/etc., in
`string literal`_ in Go code.

When I tried googling to find the way to embed files directly in Go code,
I found that if we combine the `string literal`_ with `encoding/hex`_ package in
Go standard library, we can embed any type of file in Go code.

The first step is to convert the byte sequences of a file to hex value:

.. code-block:: go

  import (
  	"encoding/hex"
  	"io/ioutil"
  )

  func FileToStringLiteral(filepath string) (sl string, err error) {
  	bs, err := ioutil.ReadFile(filepath)
  	if err != nil {
  		return
  	}
  	sl = hex.EncodeToString(bs)
  	return
  }

I create a file, the content of which is ``hello world\n``, the hex value
returned by above method is ``68656c6c6f20776f726c640a``. Now we embed the file
content in string literal as follows:

.. code-block:: go

  import (
  	"encoding/hex"
  )

  const myFile = "68656c6c6f20776f726c640a"

  func ReadMyFile() ([]byte, error) {
  	return hex.DecodeString(myFile)
  }

We can access the content of the file as follows:

.. code-block:: go

  b, err := ReadMyFile()
  if err != nil {
  	panic(err)
  }

  print(string(b))

The printed output will be ``hello world\n``.

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 17.04``, ``Go 1.8.1``.

----

References:

.. [1] | `golang literal byte array - Google search <https://www.google.com/search?q=golang+literal+byte+array>`_
       | `golang literal byte array - DuckDuckGo search <https://duckduckgo.com/?q=golang+literal+byte+array>`_
       | `golang literal byte array - Ecosia search <https://www.ecosia.org/search?q=golang+literal+byte+array>`_
       | `golang literal byte array - Qwant search <https://www.qwant.com/?q=golang+literal+byte+array>`_
       | `golang literal byte array - Bing search <https://www.bing.com/search?q=golang+literal+byte+array>`_
       | `golang literal byte array - Yahoo search <https://search.yahoo.com/search?p=golang+literal+byte+array>`_
       | `golang literal byte array - Baidu search <https://www.baidu.com/s?wd=golang+literal+byte+array>`_
       | `golang literal byte array - Yandex search <https://www.yandex.com/search/?text=golang+literal+byte+array>`_
.. [2] | `golang byte literal - Google search <https://www.google.com/search?q=golang+byte+literal>`_
       | `golang byte literal - DuckDuckGo search <https://duckduckgo.com/?q=golang+byte+literal>`_
       | `golang byte literal - Ecosia search <https://www.ecosia.org/search?q=golang+byte+literal>`_
       | `golang byte literal - Qwant search <https://www.qwant.com/?q=golang+byte+literal>`_
       | `golang byte literal - Bing search <https://www.bing.com/search?q=golang+byte+literal>`_
       | `golang byte literal - Yahoo search <https://search.yahoo.com/search?p=golang+byte+literal>`_
       | `golang byte literal - Baidu search <https://www.baidu.com/s?wd=golang+byte+literal>`_
       | `golang byte literal - Yandex search <https://www.yandex.com/search/?text=golang+byte+literal>`_
.. [3] `String literals - The Go Programming Language Specification - The Go Programming Language <https://golang.org/ref/spec#String_literals>`_
.. [4] `What is a string? - Strings, bytes, runes and characters in Go - The Go Blog <https://blog.golang.org/strings#TOC_2.>`_
.. [5] | `golang print hex string - Google search <https://www.google.com/search?q=golang+print+hex+string>`_
       | `golang print hex string - DuckDuckGo search <https://duckduckgo.com/?q=golang+print+hex+string>`_
       | `golang print hex string - Ecosia search <https://www.ecosia.org/search?q=golang+print+hex+string>`_
       | `golang print hex string - Qwant search <https://www.qwant.com/?q=golang+print+hex+string>`_
       | `golang print hex string - Bing search <https://www.bing.com/search?q=golang+print+hex+string>`_
       | `golang print hex string - Yahoo search <https://search.yahoo.com/search?p=golang+print+hex+string>`_
       | `golang print hex string - Baidu search <https://www.baidu.com/s?wd=golang+print+hex+string>`_
       | `golang print hex string - Yandex search <https://www.yandex.com/search/?text=golang+print+hex+string>`_
.. [6] `hex - The Go Programming Language <https://golang.org/pkg/encoding/hex/>`_
.. [7] `GitHub - siongui/goef: Embed file in your Go code <https://github.com/siongui/goef>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _string literal: https://golang.org/ref/spec#String_literals
.. _encoding/hex: https://golang.org/pkg/encoding/hex/

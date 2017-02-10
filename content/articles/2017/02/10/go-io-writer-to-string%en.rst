[Golang] io.Writer to string
############################

:date: 2017-02-10T09:33+08:00
:tags: Go, Golang, String Manipulation
:category: Go
:summary: Get string_ from io.Writer_ (io_ writer) in Go_ programming language.
:og_image: http://www.unixstickers.com/image/cache/data/stickers/golang/Go-brown-side.sh-600x600.png
:adsu: yes

Use Buffer_ in Go bytes_ package to get string_ from io.Writer_ (io_ Writer).
Common pattern as follows:

.. code-block:: go

  import "bytes"

  var b bytes.Buffer
  fn(b)
  b.String()


`text/template`_ Example
++++++++++++++++++++++++

`Run code on Go Playground <https://play.golang.org/p/68YQg5KlBT>`_

.. adsu:: 2

.. code-block:: go

  package main

  import (
  	"bytes"
  	"fmt"
  	"text/template"
  )

  type TmplData struct {
  	Name string
  }

  var testTmpl = `Hello {{.Name}}`

  func main() {
  	data := TmplData{Name: "World"}
  	tmpl, err := template.New("test").Parse(testTmpl)
  	if err != nil {
  		fmt.Println(err)
  	}

  	var b bytes.Buffer
  	err = tmpl.Execute(&b, &data)

  	fmt.Println(b.String())
  }

.. adsu:: 3

Final output:

.. code-block:: txt

  Hello World

----

Tested on:

- ``Ubuntu Linux 16.10``
- ``Go 1.7.5``
- `Go Playground`_

----

References:

.. adsu:: 4

.. [1] | `io writer to string golang - Google search <https://www.google.com/search?q=io+writer+to+string+golang>`_
       | `io writer to string golang - DuckDuckGo search <https://duckduckgo.com/?q=io+writer+to+string+golang>`_
       | `io writer to string golang - Ecosia search <https://www.ecosia.org/search?q=io+writer+to+string+golang>`_
       | `io writer to string golang - Bing search <https://www.bing.com/search?q=io+writer+to+string+golang>`_
       | `io writer to string golang - Yahoo search <https://search.yahoo.com/search?p=io+writer+to+string+golang>`_
       | `io writer to string golang - Baidu search <https://www.baidu.com/s?wd=io+writer+to+string+golang>`_
       | `io writer to string golang - Yandex search <https://www.yandex.com/search/?text=io+writer+to+string+golang>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _string: https://www.google.com/search?q=golang+string
.. _io.Writer: https://golang.org/pkg/io/#Writer
.. _io: https://golang.org/pkg/io/
.. _Go playground: https://play.golang.org/
.. _Buffer: https://golang.org/pkg/bytes/#Buffer
.. _bytes: https://golang.org/pkg/bytes/
.. _text/template: https://golang.org/pkg/text/template/

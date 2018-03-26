[Golang] Move File to Another Directory
#######################################

:date: 2018-03-26T23:11+08:00
:tags: Go, Golang
:category: Go
:summary: Move file to another folder via Go *os.Rename* method.
:og_image: https://cdn2.hubspot.net/hubfs/3917309/old-assets/old-theme/Images/golang-gopher-laptop.png
:adsu: yes


The shell command mv_ is used to move files in Linux system. I want to move file
directly in Go, so I searched *golang os.Move* [1]_ and found that os.Rename_ is
the method I need. Let's try to see how it works.

I have one file and one directory as follows:

.. code-block:: bash

  .
  ├── hello.txt
  └── testdir/

I want to move *hello.txt* to *testdir/*, so I write the following code:

.. code-block:: go

  package main

  import (
  	"os"
  )

  func main() {
  	err := os.Rename("hello.txt", "testdir/")
  	if err != nil {
  		panic(err)
  	}
  }

Seems reasonable, because this is how we use mv_ command under Linux shell.
When I run the Go code, panic happened:

.. code-block:: bash

  panic: rename hello.txt testdir/: file exists

So we cannot just put the name of *testdir/* in the new path of os.Rename_. We
have to specify the whole new path, including the file name:

.. code-block:: go

  package main

  import (
  	"os"
  )

  func main() {
  	err := os.Rename("hello.txt", "testdir/hello.txt")
  	if err != nil {
  		panic(err)
  	}
  }

And we get the result as expected:

.. code-block:: bash

  .
  └── testdir/
      └── hello.txt

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.10``

----

References:

.. [1] | `golang os.move - Google search <https://www.google.com/search?q=golang+os.move>`_
       | `golang os.move - DuckDuckGo search <https://duckduckgo.com/?q=golang+os.move>`_
       | `golang os.move - Ecosia search <https://www.ecosia.org/search?q=golang+os.move>`_
       | `golang os.move - Qwant search <https://www.qwant.com/?q=golang+os.move>`_
       | `golang os.move - Bing search <https://www.bing.com/search?q=golang+os.move>`_
       | `golang os.move - Yahoo search <https://search.yahoo.com/search?p=golang+os.move>`_
       | `golang os.move - Baidu search <https://www.baidu.com/s?wd=golang+os.move>`_
       | `golang os.move - Yandex search <https://www.yandex.com/search/?text=golang+os.move>`_
.. [2] `func Rename - os - The Go Programming Language <https://golang.org/pkg/os/#Rename>`_
.. [3] | `ubuntu tree command - Google search <https://www.google.com/search?q=ubuntu+tree+command>`_
       | `ubuntu tree command - DuckDuckGo search <https://duckduckgo.com/?q=ubuntu+tree+command>`_
       | `ubuntu tree command - Ecosia search <https://www.ecosia.org/search?q=ubuntu+tree+command>`_
       | `ubuntu tree command - Qwant search <https://www.qwant.com/?q=ubuntu+tree+command>`_
       | `ubuntu tree command - Bing search <https://www.bing.com/search?q=ubuntu+tree+command>`_
       | `ubuntu tree command - Yahoo search <https://search.yahoo.com/search?p=ubuntu+tree+command>`_
       | `ubuntu tree command - Baidu search <https://www.baidu.com/s?wd=ubuntu+tree+command>`_
       | `ubuntu tree command - Yandex search <https://www.yandex.com/search/?text=ubuntu+tree+command>`_

.. _mv: https://linux.die.net/man/1/mv
.. _os.Rename: https://golang.org/pkg/os/#Rename

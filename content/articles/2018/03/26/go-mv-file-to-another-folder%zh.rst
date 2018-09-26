[Go語言] 移動檔案到別的目錄
###########################

:date: 2018-09-27T01:49+08:00
:tags: Go語言
:category: Go語言
:summary: 透過 Go *os.Rename* 方法將檔案移動到另一個目錄。
:og_image: https://cdn2.hubspot.net/hubfs/3917309/old-assets/old-theme/Images/golang-gopher-laptop.png
:adsu: yes


mv_ 這個shell指令在Linux系統裡用來移動檔案。我想要在Go語言程式碼裡直接移動檔案，
所以我搜尋了 *golang os.Move* [1]_ 然後發現 os.Rename_ 是我所需要的方法。
讓我們試看看它是如何運作的。

如下，有一個檔案跟一個目錄：

.. code-block:: bash

  .
  ├── hello.txt
  └── testdir/

我想要移動 *hello.txt* 到 *testdir/* 目錄下，我寫了下面的程式碼：

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

看起來很合理，因為這是我們在Linux shell利用 mv_ 指令的方式。當我運行Go程式碼時，
發生了panic：

.. code-block:: bash

  panic: rename hello.txt testdir/: file exists

所以我們不能只是把 *testdir/* 放在 os.Rename_ 的新路徑參數。
我們必須指定完整的新路徑，包含檔名：

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

然後我們就能得到我們想要的結果：

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

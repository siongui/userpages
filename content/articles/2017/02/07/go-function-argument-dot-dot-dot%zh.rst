Go語言 函數參數點點點(...)意義
##############################

:date: 2017-02-07T10:25+08:00
:tags: Web開發, Go語言
:category: Go語言
:summary: `Go語言`_ 函數參數(`function argument`_) 點點點(...)意義
:adsu: yes


`Go語言`_ 函數參數(`function argument`_) 點點點(...)意義是，此函數可以傳入零個，
一個，或多個參數，此函數稱為 `可變參數函數`_ (`variadic function`_)。

舉例來說，假設有一個函數：

.. code-block:: go

  func giveMeString(s ...string)

則我們在呼叫此函數時，可以這樣呼叫：

.. code-block:: go

  giveMeString()
  giveMeString("hello")
  giveMeString("hello", "world")

以上呼叫函數的方式皆可，再多傳幾個 *string* 型態的參數都可以。
在Go語言標準函式庫裡，最有名的例子是 `fmt.Printf`_ ，相信大家都有用過。

.. adsu:: 2

同場加映
++++++++

延續上面的例子，假設有一個 *[]string* 型態的變數想要傳入上面的函數，
該怎樣做呢？

比方說：

.. code-block:: go

  s := []string{"hello", "world"}

上面這個變數該如何傳入呢？ 答案：

.. code-block:: go

  s := []string{"hello", "world"}
  giveMeString(s...)

在變數名稱後一樣加三個點就可以傳入。

.. adsu:: 3

----

References:

.. [1] | `golang arguments dot - Google search <https://www.google.com/search?q=golang+arguments+dot>`_
       | `golang arguments dot - DuckDuckGo search <https://duckduckgo.com/?q=golang+arguments+dot>`_
       | `golang arguments dot - Ecosia search <https://www.ecosia.org/search?q=golang+arguments+dot>`_
       | `golang arguments dot - Bing search <https://www.bing.com/search?q=golang+arguments+dot>`_
       | `golang arguments dot - Yahoo search <https://search.yahoo.com/search?p=golang+arguments+dot>`_
       | `golang arguments dot - Baidu search <https://www.baidu.com/s?wd=golang+arguments+dot>`_
       | `golang arguments dot - Yandex search <https://www.yandex.com/search/?text=golang+arguments+dot>`_

.. [2] `go - dot dot dot in Golang. interface with empty braces - Stack Overflow <http://stackoverflow.com/a/23669857>`_

.. _Go: https://golang.org/
.. _Go語言: https://golang.org/
.. _function argument: https://www.google.com/search?q=function+argument
.. _可變參數函數: https://zh.wikipedia.org/wiki/%E5%8F%AF%E8%AE%8A%E5%8F%83%E6%95%B8%E5%87%BD%E6%95%B8
.. _variadic function: https://en.wikipedia.org/wiki/Variadic_function
.. _fmt.Printf: https://golang.org/pkg/fmt/#Printf

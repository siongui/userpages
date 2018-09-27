[Go語言] 刪除字串前綴或後綴
###########################

:date: 2018-09-28T00:38+08:00
:tags: Go語言, 字串處理
:category: Go語言
:summary: 利用Go語言來刪除字串前綴或後綴。
:og_image: https://d2d42mpnbqmzj3.cloudfront.net/images/stories/doc-excel/remove-prefix-suffix/doc-remove-prefix-1.png
:adsu: yes


刪除字串前綴
++++++++++++

Q: 如何從 ``sammadiṭṭhi`` 中刪除開頭的 ``samma`` ？

A: 利用 strings_ package 裡的 TrimPrefix_ function

   .. code-block:: go

     result := strings.TrimPrefix("sammadiṭṭhi", "samma")

   看 `這裡 <https://play.golang.org/p/sOIZognDV6>`__ 的官方 TrimPrefix_ 範例。

.. adsu:: 2

刪除字串後綴
++++++++++++

Q: 如何從 ``sammadiṭṭhi`` 中刪除結尾的 ``diṭṭhi`` ？

A: 利用 strings_ package 裡的 TrimSuffix_ function

   .. code-block:: go

     result := strings.TrimSuffix("sammadiṭṭhi", "diṭṭhi")

   看 `這裡 <https://play.golang.org/p/9DR1iBH8O4>`__ 的官方 TrimSuffix_ 範例。

----

Tested on: ``Ubuntu Linux 16.10``, ``Go 1.7.4``.

----

.. adsu:: 3

References:

.. [1] | `golang remove prefix - Google search <https://www.google.com/search?q=golang+remove+prefix>`_
       | `golang remove prefix - DuckDuckGo search <https://duckduckgo.com/?q=golang+remove+prefix>`_
       | `golang remove prefix - Ecosia search <https://www.ecosia.org/search?q=golang+remove+prefix>`_
       | `golang remove prefix - Bing search <https://www.bing.com/search?q=golang+remove+prefix>`_
       | `golang remove prefix - Yahoo search <https://search.yahoo.com/search?p=golang+remove+prefix>`_
       | `golang remove prefix - Baidu search <https://www.baidu.com/s?wd=golang+remove+prefix>`_
       | `golang remove prefix - Yandex search <https://www.yandex.com/search/?text=golang+remove+prefix>`_
       | `func TrimPrefix - Package strings - The Go Programming Language <https://golang.org/pkg/strings/#TrimPrefix>`_
       | `func TrimSuffix - Package strings - The Go Programming Language <https://golang.org/pkg/strings/#TrimSuffix>`_

.. [2] | `golang remove prefix from string - Google search <https://www.google.com/search?q=golang+remove+prefix+from+string>`_
       | `golang remove prefix from string - DuckDuckGo search <https://duckduckgo.com/?q=golang+remove+prefix+from+string>`_
       | `golang remove prefix from string - Ecosia search <https://www.ecosia.org/search?q=golang+remove+prefix+from+string>`_
       | `golang remove prefix from string - Bing search <https://www.bing.com/search?q=golang+remove+prefix+from+string>`_
       | `golang remove prefix from string - Yahoo search <https://search.yahoo.com/search?p=golang+remove+prefix+from+string>`_
       | `golang remove prefix from string - Baidu search <https://www.baidu.com/s?wd=golang+remove+prefix+from+string>`_
       | `golang remove prefix from string - Yandex search <https://www.yandex.com/search/?text=golang+remove+prefix+from+string>`_

.. [3] | `golang trim prefix from string - Google search <https://www.google.com/search?q=golang+trim+prefix+from+string>`_
       | `golang trim prefix from string - DuckDuckGo search <https://duckduckgo.com/?q=golang+trim+prefix+from+string>`_
       | `golang trim prefix from string - Ecosia search <https://www.ecosia.org/search?q=golang+trim+prefix+from+string>`_
       | `golang trim prefix from string - Bing search <https://www.bing.com/search?q=golang+trim+prefix+from+string>`_
       | `golang trim prefix from string - Yahoo search <https://search.yahoo.com/search?p=golang+trim+prefix+from+string>`_
       | `golang trim prefix from string - Baidu search <https://www.baidu.com/s?wd=golang+trim+prefix+from+string>`_
       | `golang trim prefix from string - Yandex search <https://www.yandex.com/search/?text=golang+trim+prefix+from+string>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _TrimPrefix: https://golang.org/pkg/strings/#TrimPrefix
.. _TrimSuffix: https://golang.org/pkg/strings/#TrimSuffix
.. _strings: https://golang.org/pkg/strings/

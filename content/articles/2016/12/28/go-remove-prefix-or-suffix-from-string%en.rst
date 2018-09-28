[Golang] Trim Prefix or Suffix from String
##########################################

:date: 2016-12-28T19:54+08:00
:tags: Go, Golang, String Manipulation, Go strings Package
:category: Go
:summary: Remove prefix or suffix from string via Go_ programming language.
:og_image: https://d2d42mpnbqmzj3.cloudfront.net/images/stories/doc-excel/remove-prefix-suffix/doc-remove-prefix-1.png
:adsu: yes


Trim/Remove Prefix
++++++++++++++++++

Q: How to trim leading ``samma`` from ``sammadiṭṭhi``?

A: Use TrimPrefix_ function in strings_ package

   .. code-block:: go

     result := strings.TrimPrefix("sammadiṭṭhi", "samma")

   See `here <https://play.golang.org/p/sOIZognDV6>`__ for official example of
   TrimPrefix_.


Trim/Remove Suffix
++++++++++++++++++

Q: How to trim trailing ``diṭṭhi`` from ``sammadiṭṭhi``?

A: Use TrimSuffix_ function in strings_ package

   .. code-block:: go

     result := strings.TrimSuffix("sammadiṭṭhi", "diṭṭhi")

   See `here <https://play.golang.org/p/9DR1iBH8O4>`__ for official example of
   TrimSuffix_.

----

Tested on: ``Ubuntu Linux 16.10``, ``Go 1.7.4``.

----

.. adsu:: 2

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

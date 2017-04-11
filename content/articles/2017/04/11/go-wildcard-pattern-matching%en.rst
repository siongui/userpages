[Golang] Wildcard Pattern Matching
##################################

:date: 2017-04-11T22:04+08:00
:tags: Go, Golang, String Manipulation, Algorithm
:category: Go
:summary: Given a text and a wildcard pattern, implement wildcard pattern
          matching algorithm that finds if wildcard pattern is matched with
          text in Go programming language.
:og_image: https://d1ohg4ss876yi2.cloudfront.net/preview/golang.png
:adsu: yes


Question
++++++++

Given a text and a wildcard pattern, implement wildcard pattern matching
algorithm that finds if wildcard pattern is matched with text. The matching
should cover the entire text (not partial text).

The wildcard pattern can include the characters ‘?’ and ‘*’

  ‘?’ – matches any single character

  ‘*’ – Matches any sequence of characters (including the empty sequence)

(Question copyed from [2]_)


Answer
++++++

For detailed explanation, please read the `original post`_. Here only the Go
implementation code is provided.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/ma1OSOR-L2>`__
   :class: align-center

.. show_github_file:: siongui userpages content/code/go/wildcard-pattern-matching/wpm.go
.. adsu:: 2

Testing:

.. show_github_file:: siongui userpages content/code/go/wildcard-pattern-matching/wpm_test.go
.. adsu:: 3

----

Tested on:

- ``Ubuntu Linux 16.10``, ``Go 1.8.1``
- `Go Playground`_

----

References:

.. [1] | `Wildcard Pattern Matching - Google search <https://www.google.com/search?q=Wildcard+Pattern+Matching>`_
       | `Wildcard Pattern Matching - DuckDuckGo search <https://duckduckgo.com/?q=Wildcard+Pattern+Matching>`_
       | `Wildcard Pattern Matching - Ecosia search <https://www.ecosia.org/search?q=Wildcard+Pattern+Matching>`_
       | `Wildcard Pattern Matching - Qwant search <https://www.qwant.com/?q=Wildcard+Pattern+Matching>`_
       | `Wildcard Pattern Matching - Bing search <https://www.bing.com/search?q=Wildcard+Pattern+Matching>`_
       | `Wildcard Pattern Matching - Yahoo search <https://search.yahoo.com/search?p=Wildcard+Pattern+Matching>`_
       | `Wildcard Pattern Matching - Baidu search <https://www.baidu.com/s?wd=Wildcard+Pattern+Matching>`_
       | `Wildcard Pattern Matching - Yandex search <https://www.yandex.com/search/?text=Wildcard+Pattern+Matching>`_
.. [2] `Wildcard Pattern Matching - GeeksforGeeks <http://www.geeksforgeeks.org/wildcard-pattern-matching/>`_
.. [3] `Wildcard Pattern Matching - Techie Delight <http://www.techiedelight.com/wildcard-pattern-matching/>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
.. _original post: http://www.geeksforgeeks.org/wildcard-pattern-matching/

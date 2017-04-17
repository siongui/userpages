[Golang] Check if Integer is Palindromic Number
###############################################

:date: 2017-04-15T09:58+08:00
:tags: Go, Golang, Type Casting, Type Conversion, Algorithm
:category: Go
:summary: Check if an integer number is palindrome in Go programming language.
:og_image: https://www.calhoun.io/content/images/2016/10/Let-s-Learn-Algorithms---Bubble-Sort-impl.png
:adsu: yes

Check if an decimal integer number (*int64*) is palindrome_ in Golang_.

The algorithm:

1. Get absolute value of the integer (*int64*).
2. Convert the *int64* type to *string* type using strconv.FormatInt_.
3. Check if string[0] == string[n-1] where n is the length of the string, then
   check if string[1] == string[n-2] ... and so on, until the middle digit of
   string. If all are true, return *true*. Otherwise return *false*.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/Rji0XhRd5M>`__
   :class: align-center

.. show_github_file:: siongui userpages content/code/go/palindrome-number/decimal.go
.. adsu:: 2

Testing:

.. show_github_file:: siongui userpages content/code/go/palindrome-number/decimal_test.go
.. adsu:: 2

Tested on:

- ``Ubuntu Linux 16.10``, ``Go 1.8.1``
- Go Playground

----

References:

.. [1] | `golang absolute value - Google search <https://www.google.com/search?q=golang+absolute+value>`_
       | `golang absolute value int - Google search <https://www.google.com/search?q=golang+absolute+value+int>`_

.. [2] `[Golang] Type Conversion between String and Integer <{filename}../14/go-string-int-type-casting%en.rst>`_

.. [3] | `c2go v0.8.3 - now entirely written in Go :) : golang <https://www.reddit.com/r/golang/comments/65bc40/c2go_v083_now_entirely_written_in_go/>`_
       | `GitHub - elliotchance/c2go: ⚖️ A tool for converting C to Go. <https://github.com/elliotchance/c2go>`_
       | `C programming examples | Programming Simplified <http://www.programmingsimplified.com/c-program-examples>`_
       | `Palindrome Numbers | Programming Simplified <http://www.programmingsimplified.com/c/source-code/c-program-palindrome-number>`_

.. [4] | `algorithm - How do I check if a number is a palindrome? - Stack Overflow <http://stackoverflow.com/questions/199184/how-do-i-check-if-a-number-is-a-palindrome>`_
       | `Largest palindrome product - Problem 4 - Project Euler <https://projecteuler.net/index.php?section=problems&id=4>`_

.. [5] | `A Competitive Programmer's Handbook | Hacker News <https://news.ycombinator.com/item?id=14115826>`_
       | `Competitive Programmer's Handbook <https://cses.fi/book.html>`_

.. [6] | `Data structures and algorithms interview questions and their solutions | Hacker News <https://news.ycombinator.com/item?id=14128145>`_
       | `500 Data structures and algorithms interview questions and their solutions - Techie Delight - Quora <https://techiedelight.quora.com/500-Data-structures-and-algorithms-interview-questions-and-their-solutions>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _strconv.FormatInt: https://golang.org/pkg/strconv/#FormatInt
.. _palindrome: https://www.google.com/search?q=palindrome+number

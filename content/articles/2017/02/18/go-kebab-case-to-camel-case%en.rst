[Golang] kebab-case to camelCase
################################

:date: 2017-02-18T16:52+08:00
:tags: Go, Golang, String Manipulation
:category: Go
:summary: Convert the `kebab-case`_ to camelCase_ in Go_ programming language.
:adsu: yes

Convert the `kebab-case`_ (also called spinal-case, Train-Case, or Lisp-case)
[4]_ string to camelCase_ [3]_ in Golang_.

The motivation is to convert the property name of CSS (kebab-case) to element's
inline style attribute (camelCase) manipulated by JavaScript.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/Qs4LSWUkcR>`_
      :class: align-center

.. show_github_file:: siongui userpages content/code/go/kebab-case-to-camelCase/converter.go
.. adsu:: 2

Use range_ keyword [5]_ to iterate over the string. If **-** is met, *isToUpper*
is set to *true*. In next iteration convert the letter to upper case by
strings.ToUpper_, and set *isToUpper* as *false*.

.. show_github_file:: siongui userpages content/code/go/kebab-case-to-camelCase/converter_test.go
.. adsu:: 3

----

Tested on:

- ``Ubuntu Linux 16.10``, ``Go 1.8``
- The `Go Playground`_

----

References:

.. [1] | `variable naming conventions - Google search <https://www.google.com/search?q=variable+naming+conventions>`_
       | `variable naming conventions - DuckDuckGo search <https://duckduckgo.com/?q=variable+naming+conventions>`_
       | `variable naming conventions - Ecosia search <https://www.ecosia.org/search?q=variable+naming+conventions>`_
       | `variable naming conventions - Bing search <https://www.bing.com/search?q=variable+naming+conventions>`_
       | `variable naming conventions - Yahoo search <https://search.yahoo.com/search?p=variable+naming+conventions>`_
       | `variable naming conventions - Baidu search <https://www.baidu.com/s?wd=variable+naming+conventions>`_
       | `variable naming conventions - Yandex search <https://www.yandex.com/search/?text=variable+naming+conventions>`_

.. [2] `Naming convention (programming) - Wikipedia <https://en.wikipedia.org/wiki/Naming_convention_(programming)>`_

.. [3] `Camel case - Wikipedia <https://en.wikipedia.org/wiki/Camel_case>`_

.. [4] | kebab-case(spinal-case, Train-case)
       | `programming languages - What's the name for hyphen-separated case? - Stack Overflow <https://stackoverflow.com/questions/11273282/whats-the-name-for-hyphen-separated-case>`_
       | `Letter case - Wikipedia <https://en.wikipedia.org/wiki/Letter_case#Special_case_styles>`_

.. [5] `[Golang] Iterate Over UTF-8 Strings (non-ASCII strings) <{filename}../../../2016/02/03/go-iterate-over-utf8-non-ascii-string%en.rst>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
.. _kebab-case: https://stackoverflow.com/questions/11273282/whats-the-name-for-hyphen-separated-case
.. _camelCase: https://en.wikipedia.org/wiki/Camel_case
.. _for: https://tour.golang.org/flowcontrol/1
.. _range: https://github.com/golang/go/wiki/Range
.. _strings.ToUpper: https://golang.org/pkg/strings/#ToUpper

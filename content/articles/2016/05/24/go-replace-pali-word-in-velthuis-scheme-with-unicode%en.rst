[Golang] Replace Pāli Word in Velthuis Scheme With Unicode
##########################################################

:date: 2016-05-24T03:56+08:00
:tags: Go, Golang, Regular Expression, String Manipulation, File Input/Output
:category: Go
:summary: Find `Pāli`_ word in `Velthuis scheme`_, and replace them with
          unicode_ via Go_ programming language.
:adsu: yes


Find `Pāli`_ word in `Velthuis scheme`_, and replace them with unicode_
via Golang_ (Go_ programming language).

.. show_github_file:: siongui userpages content/code/go-velthuis-to-unicode/replace.go

.. show_github_file:: siongui userpages content/code/go-velthuis-to-unicode/replace_test.go

----

Tested on: ``Ubuntu Linux 16.04``, ``Go 1.6.2``.

----

References:

.. [1] `[Golang] Find Pāli Word in Velthuis Scheme <{filename}../../03/17/go-find-pali-word-in-velthuis-scheme%en.rst>`_

.. [2] `regex - Regular Expression to match only alphabetic characters - Stack Overflow <http://stackoverflow.com/questions/6067592/regular-expression-to-match-only-alphabetic-characters>`_

.. [3] `[Golang] Iterate Over UTF-8 Strings (non-ASCII strings) <{filename}../../02/03/go-iterate-over-utf8-non-ascii-string%en.rst>`_

.. [4] `GitHub - matko/emacs-pali-velthuis: emacs input method to write pali in the latin alphabet augmented with diacritical marks, using the Velthuis method <https://github.com/matko/emacs-pali-velthuis>`_

.. [5] `regexp - The Go Programming Language <https://golang.org/pkg/regexp/>`_

       `strings - The Go Programming Language <https://golang.org/pkg/strings/>`_

       `ioutil - The Go Programming Language <https://golang.org/pkg/io/ioutil/>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Pāli: https://en.wikipedia.org/wiki/Pali
.. _Velthuis scheme: http://www.accesstoinsight.org/abbrev.html#velthuis
.. _unicode: https://en.wikipedia.org/wiki/Unicode

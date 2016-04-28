[Golang] Wrap Pāli Words in Span Element
########################################

:date: 2016-04-28T20:58+08:00
:tags: Go, Golang, html, Regular Expression, String Manipulation
:category: Go
:summary: Wrap `Pāli`_ words in span_ element via `regular expression`_ and Go_.


Wrap `Pāli`_ words in span_ element via `regular expression`_ and Golang_ (Go_
programming language).

Source code:

.. show_github_file:: siongui userpages content/code/go-wrap-word-in-span/pali.go

.. show_github_file:: siongui userpages content/code/go-wrap-word-in-span/pali_test.go

Output:

.. code-block:: txt

  === RUN   TestWrapPaliWordsInSpan
  --- PASS: TestWrapPaliWordsInSpan (0.00s)
  	pali_test.go:11: 1.
  		<span>Manopubbaṅgamā</span> <span>dhammā</span>, <span>manoseṭṭhā</span> <span>manomayā</span>;
  		<span>Manasā</span> <span>ce</span> <span>paduṭṭhena</span>, <span>bhāsati</span> <span>vā</span> <span>karoti</span> <span>vā</span>;
  		<span>Tato</span> <span>naṃ</span> <span>dukkhamanveti</span>, <span>cakkaṃva</span> <span>vahato</span> <span>padaṃ</span>.
  PASS

----

Tested on: ``Ubuntu Linux 16.04``, ``Go 1.6.2``.

----

References:

.. [1] `[JavaScript] Wrap Pāli Words in Span Element <{filename}../02/javascript-wrap-pali-words-in-span-element%en.rst>`_

.. [2] `regexp - The Go Programming Language <https://golang.org/pkg/regexp/>`_

.. [3] `[Golang] Find Pāli Word in Velthuis Scheme <{filename}../../03/17/go-find-pali-word-in-velthuis-scheme%en.rst>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Pāli: https://en.wikipedia.org/wiki/Pali
.. _span: http://www.w3schools.com/tags/tag_span.asp
.. _regular expression: https://www.google.com/search?q=regular+expression

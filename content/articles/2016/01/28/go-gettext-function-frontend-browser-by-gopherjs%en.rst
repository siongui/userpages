[Golang] gettext Function on Frontend (Browser) by GopherJS
###########################################################

:date: 2016-01-28T08:46+08:00
:tags: Go, Golang, i18n, Locale, JSON, gettext, GopherJS, Go to JavaScript,
       DOM, html, Web application, web
:category: GopherJS
:summary: Implement `gettext function`_ on front-end (browsers) by GopherJS_.
          The *gettext* function translates a text string into the user's
          native language.


Introduction
++++++++++++

Implement `gettext function`_ on front-end (browsers) by Go_ and  GopherJS_,
which compiles Golang_ program to JavaScript_. The *gettext* function translates
a text string into the user's native language. Before you start, you need to
prepare PO_ files (see [4]_, [5]_, [6]_) and convert the PO files to JSON_
format (see [7]_). The translations in JSON_ format will be used by our
*gettext* function.

Install GopherJS_
+++++++++++++++++

Run the following command to install GopherJS_:

.. code-block:: bash

  $ go get -u github.com/gopherjs/gopherjs

Source Code
+++++++++++

HTML file for demo:

.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/jsgettext/index.html

In our HTML file, we wrap a text string to be translated in *div* element, and
set the *data-default-string* attribute of *div* element to the text string. For
example, **Home** is a text string to be translated, it will be wrapped as:

.. code-block:: html

  <div data-default-string="Home">Home</div>

You can also wrap a text string in *span* element or other elements instead of
*div*.

.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/jsgettext/jsgettext-raw.go

We put translations converted from PO files in the source code. The translations
are used by *gettext* function to translate a text string.

When a user clicks the language button, we will select the *div* elements with
*data-default-string* attribute. The string in the attribute will be translated
by *gettext* function and replace the original string in the *div*.

Compile the Go_ code to JavaScript_ by:

.. code-block:: bash

  $ gopherjs build jsgettext-raw.go -o demo.js

Put *demo.js* together with the *index.html* in the same directory. Open the
*index.html* with your browser. Click language buttons to translate strings.

----

Appendix
++++++++

If you want to use `GopherJS bindings for the JavaScript DOM APIs`_ to write
idiomatic Go_ code, install DOM_ binding by:

.. code-block:: bash

  $ go get -u honnef.co/go/js/dom

And compile the following code to JavaScript instead of above:

.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/jsgettext/jsgettext.go

Note that if you wrap strings in *span* or other elements instead of *div*,
remember to modify the above code accordingly.

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.5.3``,
``Chromium Version 48.0.2564.82 Ubuntu 15.10 (64-bit)``.

----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] `Bindings · gopherjs/gopherjs Wiki · GitHub <https://github.com/gopherjs/gopherjs/wiki/bindings>`_

.. [3] `dom - GopherJS bindings for the JavaScript DOM APIs <https://godoc.org/honnef.co/go/js/dom>`_
       (`GitHub <https://github.com/dominikh/go-js-dom>`__)

.. [4] `Internationalization (i18n) of Web Application by GNU gettext Tools <{filename}../07/i18n-web-application-by-gnu-gettext-tools%en.rst>`_

.. [5] `i18n Golang Web Application by gettext and html/template <{filename}../19/i18n-go-web-application-by-gettext-html-template%en.rst>`_

.. [6] `xgettext Extract Translatable Strings From Golang html/template <{filename}../19/xgettext-extract-translatable-string-from-go-html-template%en.rst>`_

.. [7] `[Golang] Convert PO file to JSON Format <{filename}../27/go-convert-po-file-to-json%en.rst>`_

.. [8] `queryselector <https://www.google.com/search?q=queryselector>`__

.. [9] `Document.querySelector() - Web APIs | MDN <https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelector>`_

.. [10] `HTML DOM querySelector() Method <http://www.w3schools.com/jsref/met_document_queryselector.asp>`_

.. [11] `CSS Selectors Reference <http://www.w3schools.com/cssref/css_selectors.asp>`_

.. [12] `queryselector attribute selector <https://www.google.com/search?q=queryselector+attribute+selector>`_

.. [13] `javascript - How to use querySelectorAll only for elements that have a specific attribute set? - Stack Overflow <http://stackoverflow.com/questions/10777684/how-to-use-queryselectorall-only-for-elements-that-have-a-specific-attribute-set>`_

.. [14] `javascript - document.querySelector multiple data-attributes in one element - Stack Overflow <http://stackoverflow.com/questions/29937768/document-queryselector-multiple-data-attributes-in-one-element>`_

.. [15] `json - The Go Programming Language <https://golang.org/pkg/encoding/json/#example_Decoder_Decode_stream>`_

.. [16] `golang map key exists <https://www.google.com/search?q=golang+map+key+exists>`_

.. [17] `dictionary - How to check if a map contains a key in go? - Stack Overflow <http://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go>`_

.. [18] `javascript gettext <https://www.google.com/search?q=javascript+gettext>`_

.. [19] `Jed | Gettext Style i18n for Modern JavaScript Apps <https://slexaxton.github.io/Jed/>`_

.. [20] `javascript - .po files and gettext VS JSON and custom i18n library? - Stack Overflow <http://stackoverflow.com/questions/8121575/po-files-and-gettext-vs-json-and-custom-i18n-library>`_

.. [21] `javascript gettext frontend <https://www.google.com/search?q=javascript+gettext+frontend>`_

.. [22] `gettext - How to split frontend and backend translations? - Stack Overflow <http://stackoverflow.com/questions/25365064/how-to-split-frontend-and-backend-translations>`_


.. _GopherJS bindings for the JavaScript DOM APIs: https://godoc.org/honnef.co/go/js/dom
.. _GopherJS: http://www.gopherjs.org/
.. _DOM: https://developer.mozilla.org/en-US/docs/Web/API/Document_Object_Model
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _gettext: https://www.gnu.org/software/gettext/
.. _locale: https://en.wikipedia.org/wiki/Locale
.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _PO: https://www.gnu.org/software/gettext/manual/html_node/PO-Files.html
.. _JSON: https://www.google.com/search?q=JSON
.. _gettext function: http://linux.die.net/man/3/gettext

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js

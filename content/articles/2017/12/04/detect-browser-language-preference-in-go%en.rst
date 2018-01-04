Detect Browser Language Preference in Go
########################################

:date: 2017-12-26T23:09+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go, i18n,
       Locale
:category: Frontend Programming in Go
:summary: Know which language(s) users prefer on browsers in Go/GopherJS.
:og_image: https://zzz.buzz/content/2016/01/navigator-languages-chrome.png
:adsu: yes


To build a multilingual website, it is important to know language preferences
of users. There are two ways for frontend programmers to detect browser language
preference:

- Via NavigatorLanguage_ API: This API contains properties (*language* and
  *languages*) to indicate the preferred language of the browser user. This post
  will mainly focus on how to use this API.
- Via `Accept-Language`_ header in HTTP request: There is no way to directly
  access this header in your code. You can, however, run a server which returns
  the headers to client browser via JSONP. See [3]_ for more details.

.. contents:: **Table of Content**

Access NavigatorLanguage API
============================

`navigator.language`_ returns the preferred language of the user, while
`navigator.languages`_ returns languages known to the user.


JavaScript
++++++++++

.. code-block:: javascript

  console.log(navigator.language);
  console.log(navigator.languages);

The return value of *navigator.language* is string, and the return value of
*navigator.languages* is an array.


GopherJS
++++++++

The above code in Go/GopherJS is as follows:

.. code-block:: go

  import (
  	"github.com/gopherjs/gopherjs/js"
  )

  println(js.Global.Get("navigator").Get("language").String())
  println(js.Global.Get("navigator").Get("languages").String())

.. rubric:: `Run Code on GopherJS Playground <https://gopherjs.github.io/playground/#/jPwHSssUd6>`__
   :class: align-center

Note that the return value of both are strings because of the constraint of
GopherJS compiler.

.. adsu:: 2


GopherJS + godom
++++++++++++++++

To make your code more readable, we can prettify the above code with godom_:

.. code-block:: go

  import (
  	. "github.com/siongui/godom"
  )

  println(Window.Navigator().Language())
  println(Window.Navigator().Languages())

The full code example of this post is `on my GitHub`_.

----

References:

.. [1] `[Golang] Detect Browser Language Preference by window.navigator.language <{filename}../../../2016/01/24/go-detect-browser-language-preference%en.rst>`_
.. [2] `[Golang] GopherJS Synonyms with JavaScript <{filename}../../../2016/01/29/go-gopherjs-synonyms-with-javascript%en.rst>`_
.. [3] `[Golang] Access HTTP Request Header (Accept-Language) by JSONP <{filename}../../../2016/01/24/go-http-request-header-by-jsonp-gopherjs%en.rst>`_

.. _GopherJS: http://www.gopherjs.org/
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _Go: https://golang.org/
.. _godom: https://github.com/siongui/godom
.. _on my GitHub: https://github.com/siongui/frontend-programming-in-go/tree/master/013-navigator-language
.. _NavigatorLanguage: https://developer.mozilla.org/en-US/docs/Web/API/NavigatorLanguage
.. _Accept-Language: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Language
.. _navigator.language: https://developer.mozilla.org/en-US/docs/Web/API/NavigatorLanguage/language
.. _navigator.languages: https://developer.mozilla.org/en-US/docs/Web/API/NavigatorLanguage/languages

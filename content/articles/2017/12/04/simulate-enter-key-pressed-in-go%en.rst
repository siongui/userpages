Simulate Enter Key Pressed in Go
################################

:date: 2020-07-15T08:31+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go,
       Keyboard Event
:category: Frontend Programming in Go
:summary: Programmatically fire enter key event in Go/GopherJS.
:og_image: https://i.stack.imgur.com/FgAVt.png
:adsu: yes


Simulate enter key pressed via Go_/ GopherJS_, i.e., programmatically fire enter
key event in browsers without user intervention.


JavaScript
++++++++++

From [2]_ we know JavaScript code for firing enter keyevent would be:

.. code-block:: javascript

  const ke = new KeyboardEvent("keyup", {keyCode: 13});
  document.body.dispatchEvent(ke);


GopherJS
++++++++

The above code in Go/GopherJS is as follows:

.. code-block:: go

  import (
  	"github.com/gopherjs/gopherjs/js"
  )

  option := js.Global.Get("Object").New()
  option.Set("keyCode", 13)
  ke := js.Global.Get("KeyboardEvent").New("keyup", option)
  js.Global.Get("document").Get("body").Call("dispatchEvent", ke)


GopherJS + godom
++++++++++++++++

To make your code more readable, we can prettify the above code with godom_:

.. code-block:: go

  import (
  	. "github.com/siongui/godom"
  )

  option := Window.Get("Object").New()
  option.Set("keyCode", 13)
  ke := Window.Get("KeyboardEvent").New("keyup", option)
  Document.Get("body").Call("dispatchEvent", ke)

----

References:

.. [1] | `javascript fire enter key event - Google search <https://www.google.com/search?q=javascript+fire+enter+key+event>`_
       | `javascript fire enter key event - DuckDuckGo search <https://duckduckgo.com/?q=javascript+fire+enter+key+event>`_
       | `javascript fire enter key event - Ecosia search <https://www.ecosia.org/search?q=javascript+fire+enter+key+event>`_
       | `javascript fire enter key event - Qwant search <https://www.qwant.com/?q=javascript+fire+enter+key+event>`_
       | `javascript fire enter key event - Bing search <https://www.bing.com/search?q=javascript+fire+enter+key+event>`_
       | `javascript fire enter key event - Yahoo search <https://search.yahoo.com/search?p=javascript+fire+enter+key+event>`_
       | `javascript fire enter key event - Baidu search <https://www.baidu.com/s?wd=javascript+fire+enter+key+event>`_
       | `javascript fire enter key event - Yandex search <https://www.yandex.com/search/?text=javascript+fire+enter+key+event>`_
.. [2] `javascript - jquery (or pure js) simulate enter key pressed for testing - Stack Overflow <https://stackoverflow.com/a/18937620>`_
.. [3] `Frontend Programming in Go <{filename}/articles/2017/12/04/frontend-programming-in-go%en.rst>`_
.. [4] `Synonyms - Go and JavaScript <{filename}/articles/2017/12/04/synonyms-go-and-javascript%en.rst>`_
.. [5] `Watch Data Change With Options via Go and Vue.js <{filename}/articles/2017/12/04/watch-data-change-with-option-via-gopherjs-vue%en.rst>`_
.. [6] `KeyboardEvent() - Web APIs | MDN <https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/KeyboardEvent>`_
.. [7] `JavaScript Tips and Gotchas · gopherjs/gopherjs Wiki · GitHub <https://github.com/gopherjs/gopherjs/wiki/JavaScript-Tips-and-Gotchas>`_

.. _Go: https://golang.org/
.. _GopherJS: http://www.gopherjs.org/
.. _godom: https://github.com/siongui/godom

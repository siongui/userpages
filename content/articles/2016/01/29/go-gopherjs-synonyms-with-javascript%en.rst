[Golang] GopherJS Synonyms with JavaScript
##########################################

:date: 2016-01-29T04:26+08:00
:modified: 2016-01-30T05:57+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, DOM, JavaScript
:category: GopherJS
:summary: Synonyms - Go_/GopherJS_ idioms and snippets translated to JavaScript_


Inspired by [4]_, Golang_/GopherJS_ (with / without
`GopherJS bindings for the JavaScript DOM APIs`_) synonyms with JavaScript_.


.. list-table:: JavaScript_ vs GopherJS_ vs GopherJS with `DOM binding`_
   :header-rows: 1
   :class: table-syntax-diff

   * - JavaScript_
     - GopherJS_
     - GopherJS with `DOM binding`_

   * - .. code-block:: javascript

         // no need to import anything

     - .. code-block:: go

         import "github.com/gopherjs/gopherjs/js"

     - .. code-block:: go

         import "honnef.co/go/js/dom"

   * - The window_ object

       .. code-block:: javascript

         window

     - type Object_

       .. code-block:: go

         js.Global

     - `GetWindow()`_ function

       .. code-block:: go

         dom.GetWindow()

   * - `alert()`_

       .. code-block:: javascript

         alert("Hello World")

     - |

       .. code-block:: go

         js.Global.Call("alert", "Hello World")

     - |

       .. code-block:: go

         dom.GetWindow().Alert("Hello World")

   * - The HTML DOM_ document_ Object

       .. code-block:: javascript

         document

     - |

       .. code-block:: go

         js.Global.Get("document")

     - |

       .. code-block:: go

         dom.GetWindow().Document()

   * - navigator_ Object

       .. code-block:: javascript

         window.navigator

     - |

       .. code-block:: go

         js.Global.Get("navigator")

     - |

       .. code-block:: go

         dom.GetWindow().Navigator()

   * - NavigatorLanguage_ API

       .. code-block:: javascript

         window.navigator.language
         window.navigator.languages

     - |

       .. code-block:: go

         js.Global.Get("navigator").Get("language").String()
         js.Global.Get("navigator").Get("languages").String()

     - |

       .. code-block:: go

         // not implemented

   * - `getElementById()`_

       .. code-block:: javascript

         element = document.getElementById("foo");

     - |

       .. code-block:: go

         element := js.Global.Get("document").Call("getElementById", "foo")

     - |

       .. code-block:: go

         element := dom.GetWindow().Document().GetElementByID("foo")

   * - HTML DOM innerHTML_ property

       .. code-block:: javascript

         // read
         var value = document.getElementById("foo").innerHTML;
         // set
         document.getElementById("foo").innerHTML = "new value";

     - |

       .. code-block:: go

         element := js.Global.Get("document").Call("getElementById", "foo")
         // read
         value := element.Get("innerHTML").String()
         // set
         element.Set("innerHTML", "new value")

     - |

       .. code-block:: go

         element := dom.GetWindow().Document().GetElementByID("foo")
         // read
         value := element.InnerHTML()
         // set
         element.SetInnerHTML("new value")

   * - HTML DOM textContent_ property

       .. code-block:: javascript

         // read
         var value = document.getElementById("foo").textContent;
         // set
         document.getElementById("foo").textContent = "new value";

     - |

       .. code-block:: go

         element := js.Global.Get("document").Call("getElementById", "foo")
         // read
         value := element.Get("textContent").String()
         // set
         element.Set("textContent", "new value")

     - |

       .. code-block:: go

         element := dom.GetWindow().Document().GetElementByID("foo")
         // read
         value := element.TextContent()
         // set
         element.SetTextContent("new value")

   * - Event binding - `addEventListener()`_

       .. code-block:: javascript

         var foo = document.getElementById("foo");
         // register onclick event
         foo.addEventListener("click", function(event) {
           // do something
           event.preventDefault()
         }, false);

     - |

       .. code-block:: go

         foo := js.Global.Get("document").Call("getElementById", "foo")
         // register onclick event
         foo.Call("addEventListener", "click", func(event *js.Object) {
           // do something
           event.Call("preventDefault")
         }, false)

     - |

       .. code-block:: go

         foo := dom.GetWindow().Document().GetElementByID("foo")
         // register onclick event
         foo.AddEventListener("click", false, func(event dom.Event) {
           // do something
           event.PreventDefault()
         })


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

.. [4] `Synonyms - Dart, JavaScript, C#, Python | Dart <https://www.dartlang.org/docs/synonyms/>`_

.. _GopherJS: http://www.gopherjs.org/
.. _DOM binding: https://godoc.org/honnef.co/go/js/dom
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _window: http://www.w3schools.com/jsref/obj_window.asp
.. _Object: https://godoc.org/github.com/gopherjs/gopherjs/js#Object
.. _GetWindow(): https://godoc.org/honnef.co/go/js/dom#GetWindow
.. _document: http://www.w3schools.com/jsref/dom_obj_document.asp
.. _GopherJS bindings for the JavaScript DOM APIs: https://godoc.org/honnef.co/go/js/dom
.. _DOM: https://developer.mozilla.org/en-US/docs/Web/API/Document_Object_Model
.. _alert(): http://www.w3schools.com/jsref/met_win_alert.asp
.. _navigator: https://developer.mozilla.org/en-US/docs/Web/API/Navigator
.. _NavigatorLanguage: https://developer.mozilla.org/en-US/docs/Web/API/NavigatorLanguage
.. _getElementById(): https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementById
.. _innerHTML: http://www.w3schools.com/jsref/prop_html_innerhtml.asp
.. _textContent: http://www.w3schools.com/jsref/prop_node_textcontent.asp
.. _addEventListener(): https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js

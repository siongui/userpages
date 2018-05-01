Toggle (Show/Hide) HTML Element via Go and Vue.js
#################################################

:date: 2018-01-21T08:23+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go,
       Vue.js, gopherjs-vue
:category: Frontend Programming in Go
:summary: Toggle (Show/Hide) HTML DOM element via Go/GopherJS/gopherjs-vue.
:og_image: https://cdn.scotch.io/1/PNl29thbQ72wcrx4Icoq_build-a-single-page-app-with-go-echo-vue.png.jpg
:adsu: yes

Write only Go code to toggle (Show/Hide) HTML DOM element via Vue.js_. We need
gopherjs-vue_ package, which is Vue.js bindings for GopherJS, to help us. The
example here is the same as my previous post named
*[Vue.js] Toggle (Show/Hide) HTML Element* [1]_ except that the JavaScript code
is replaced with Go code. Please see the demo, JavaScript code, and explanation
in my previous post first.

**HTML**

.. code-block:: html

  <div id="app">
    <div v-on:click="isShow = !isShow">
      click me to toggle (show/hide) HTML element
    </div>
    <div v-show="isShow">I am element being toggled</div>
  </div>

**Go**

.. code-block:: go

  package main

  import (
  	"github.com/gopherjs/gopherjs/js"
  	"github.com/oskca/gopherjs-vue"
  )

  type Model struct {
  	*js.Object      // this is needed for bidirectional data bindings
  	IsShow     bool `js:"isShow"`
  }

  func main() {
  	m := &Model{
  		Object: js.Global.Get("Object").New(),
  	}
  	// field assignment is required in this way to make data passing works
  	m.IsShow = false
  	// create the VueJS viewModel using a struct pointer
  	vue.New("#app", m)
  }

The full code example can be found `on my GitHub`_.

.. adsu:: 2

----

References:

.. [1] `[Vue.js] Toggle (Show/Hide) HTML Element <{filename}/articles/2017/02/07/vuejs-toggle-dom-element%en.rst>`_

.. _gopherjs-vue: https://github.com/oskca/gopherjs-vue
.. _Vue.js: https://vuejs.org/
.. _on my GitHub: https://github.com/siongui/frontend-programming-in-go/tree/master/022-toggle-element-gopherjs-vue

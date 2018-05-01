Show keyCode of Pressed Key via Go and Vue.js
#############################################

:date: 2018-01-23T09:39+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go,
       Vue.js, gopherjs-vue, Keyboard Event
:category: Frontend Programming in Go
:summary: Show keyCode of pressed key of focused HTML *input* element
          via Go/GopherJS/gopherjs-vue.
:og_image: https://cdn.scotch.io/1/PNl29thbQ72wcrx4Icoq_build-a-single-page-app-with-go-echo-vue.png.jpg
:adsu: yes

Example for getting *keyCode* of pressed key of focused HTML *input* element via
gopherjs-vue_.

See the following the HTML code:

**HTML**

.. code-block:: html

  <div id="vueapp">
    <input type="text" @keyup="ShowKeyCode" placeholder="Press Any Key Here" />
    <div>The keyCode of the keypress: {{ keypressed }}</div>
  </div>

When the *input* element is focused and user press some key, the *keyCode* of
the pressed key will be shown below the *input* element.

The following is the Go code for above HTML code:

**Go**

.. code-block:: go

  package main
  
  import (
  	"github.com/gopherjs/gopherjs/js"
  	"github.com/oskca/gopherjs-vue"
  )
  
  type Model struct {
  	*js.Object        // this is needed for bidirectional data bindings
  	Keypressed string `js:"keypressed"`
  }
  
  // this would be recognized as Show in html
  func (m *Model) ShowKeyCode(event *js.Object) {
  	m.Keypressed = event.Get("keyCode").String()
  }
  
  func main() {
  	m := &Model{
  		Object: js.Global.Get("Object").New(),
  	}
  	// field assignment is required in this way to make data passing works
  	m.Keypressed = ""
  	// create the VueJS viewModel using a struct pointer
  	vue.New("#vueapp", m)
  }

The above Go code is equivalent to the following JavaScript code:

**JavaScript**

.. code-block:: javascript

  'use strict';
  
  new Vue({
    el: '#vueapp',
    data: {
      keypressed: ""
    },
    methods: {
      ShowKeyCode: function (event) {
        this.keypressed = event.keyCode;
      }
    }
  });

The full code example can be found `on my GitHub`_.

.. adsu:: 2

----

References:

.. [1] `[Vue.js] Keyboard Event (Arrow Key Example) <{filename}/articles/2017/12/06/vuejs-keyboard-event-arrow-key-example%en.rst>`_

.. _gopherjs-vue: https://github.com/oskca/gopherjs-vue
.. _Vue.js: https://vuejs.org/
.. _on my GitHub: https://github.com/siongui/frontend-programming-in-go/tree/master/023-keyboard-event-keyCode-gopherjs-vue

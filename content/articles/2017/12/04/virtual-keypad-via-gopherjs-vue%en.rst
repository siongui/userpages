Virtual Keyboard via Go and Vue.js
##################################

:date: 2018-03-08T22:50+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go,
       Vue.js, IME, Pāli Input Method, Virtual Keyboard/Keypad
:category: Frontend Programming in Go
:summary: Virtual kayboard/keypad via Go/GopherJS/gopherjs-vue.
:og_image: http://2.bp.blogspot.com/-R0eGgtfg83c/VDtjrYbT0oI/AAAAAAAADXY/TvNcOzfAl7o/s1600/2014-10-13-012922_1366x768_scrot.png
:adsu: yes

Virtual kayboard/keypad via Go/GopherJS/gopherjs-vue_ (Go binding for Vue.js_).

.. rubric:: `Demo <{filename}/code/vuejs/virtual-keyboard/index.html>`_
   :class: align-center

I will list the main logic here. The link of full source code is available
`on my GitHub repo`_.

**HTML**

.. code-block:: html

  <div id="container">
    <input type="text" v-model="paliWord"><br>
    <div class="keypad">
      <div><input @click="paliWord += letter" :value="letter" type="button" v-for="letter in row1letters"></div>
      <div><input @click="paliWord += letter" :value="letter" type="button" v-for="letter in row2letters"></div>
      <div><input @click="paliWord += letter" :value="letter" type="button" v-for="letter in row3letters"></div>
      <div><input @click="paliWord += letter" :value="letter" type="button" v-for="letter in row4letters"></div>
    </div>
  </div>

**Go**

.. code-block:: go

  package main

  import (
  	"github.com/gopherjs/gopherjs/js"
  	"github.com/oskca/gopherjs-vue"
  )

  type Model struct {
  	*js.Object           // this is needed for bidirectional data bindings
  	PaliWord    string   `js:"paliWord"`
  	Row1letters []string `js:"row1letters"`
  	Row2letters []string `js:"row2letters"`
  	Row3letters []string `js:"row3letters"`
  	Row4letters []string `js:"row4letters"`
  }

  func main() {
  	m := &Model{
  		Object: js.Global.Get("Object").New(),
  	}
  	// field assignment is required in this way to make data passing works
  	m.PaliWord = ""
  	m.Row1letters = []string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p"}
  	m.Row2letters = []string{"a", "s", "d", "f", "g", "h", "j", "k", "l"}
  	m.Row3letters = []string{"z", "x", "c", "v", "b", "n", "m"}
  	m.Row4letters = []string{"ā", "ḍ", "ī", "ḷ", "ṁ", "ṃ", "ñ", "ṇ", "ṭ", "ū", "ŋ", "ṅ"}

  	// create the VueJS viewModel using a struct pointer
  	vue.New("#container", m)
  }

The JavaScript equivalent is available on my previous post [1]_.

.. adsu:: 2

----

References:

.. [1] `[Vue.js] Virtual Keyboard <{filename}/articles/2017/01/21/vuejs-virtual-keypad%en.rst>`_

.. _gopherjs-vue: https://github.com/oskca/gopherjs-vue
.. _Vue.js: https://vuejs.org/
.. _on my GitHub repo: https://github.com/siongui/frontend-programming-in-go/tree/master/027-virtula-keypad-gopherjs-vue

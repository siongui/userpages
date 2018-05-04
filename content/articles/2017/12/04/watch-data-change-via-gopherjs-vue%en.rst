Watch Data Change via Go and Vue.js
###################################

:date: 2018-05-04T23:34+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go,
       Vue.js, gopherjs-vue
:category: Frontend Programming in Go
:summary: Watch and react to data changes via Go/GopherJS/gopherjs-vue.
:og_image: https://cdn-images-1.medium.com/max/1024/1*kzBgrHhQRwEycJV8YQYR0A.png
:adsu: yes

To watch and react to data changes is a must-have feature in modern front-end
programming. Here we will show how to write Go code to watch data changes via
Go/GopherJS/gopherjs-vue_ (Go binding for Vue.js_).

See demo first:

.. raw:: html

  <div id="vueapp" style="background: aqua; padding: .5rem;">
    <input v-model="userinput" placeholder="Input something">
    <p>New Value: {{ newvalue }}</p>
    <p>Old Value: {{ oldvalue }}</p>
  </div>
  <script src="https://cdn.jsdelivr.net/npm/vue@2.5.16/dist/vue.js"></script>
  <script>
  new Vue({
    el: '#vueapp',
    data: {
      userinput: "",
      newvalue: "",
      oldvalue: ""
    },
    watch: {
      userinput: function(newVal, oldVal) {
        this.newvalue = newVal;
        this.oldvalue = oldVal;
      }
    }
  });
  </script>

The demo is simple. Just watch user input and print the new value and old value.
The source code of demo in HTML and JavaScript/Vue.js is:

**HTML**

.. code-block:: html

  <div id="vueapp">
    <input v-model="userinput" placeholder="Input something">
    <p>New Value: {{ newvalue }}</p>
    <p>Old Value: {{ oldvalue }}</p>
  </div>
  <script src="https://cdn.jsdelivr.net/npm/vue@2.5.16/dist/vue.js"></script>

**JavaScript/Vue.js**

.. code-block:: javascript

  new Vue({
    el: '#vueapp',
    data: {
      userinput: "",
      newvalue: "",
      oldvalue: ""
    },
    watch: {
      userinput: function(newVal, oldVal) {
        this.newvalue = newVal;
        this.oldvalue = oldVal;
      }
    }
  });

Now we will translate the JavaScript/Vue.js part into Go. The HTML code is the
same except you do not have to insert vue.js cdn in your HTML (vue.js code is
already included in gopherjs-vue and will be automatically compiled and included
in final code.) The following is Go equivalent code.

**Go**: *go get* *github.com/gopherjs/gopherjs* and
*github.com/oskca/gopherjs-vue* before compiling the Go code.

.. code-block:: go

  package main

  import (
  	"github.com/gopherjs/gopherjs/js"
  	"github.com/oskca/gopherjs-vue"
  )

  type Model struct {
  	*js.Object        // this is needed for bidirectional data bindings
  	UserInput  string `js:"userinput"`
  	OldValue   string `js:"oldvalue"`
  	NewValue   string `js:"newvalue"`
  }

  func main() {
  	m := &Model{
  		Object: js.Global.Get("Object").New(),
  	}
  	// field assignment is required in this way to make data passing works
  	m.UserInput = ""
  	m.OldValue = ""
  	m.NewValue = ""

  	// create the VueJS viewModel using a struct pointer
  	app := vue.New("#vueapp", m)

  	// add watcher
  	app.Call("$watch", "userinput", func(newVal, oldVal string) {
  		m.OldValue = oldVal
  		m.NewValue = newVal
  	})
  }

The link of full source code is available `on my GitHub repo`_.

.. adsu:: 2

Tested on:

- ``Chromium 65.0.3325.181 on Ubuntu 17.10 (64-bit)``
- ``Go 1.10.2``
- ``GopherJS 1.10-3``

----

References:

.. [1] `Online Sieve of Eratosthenes Demo via Go and Vue.js <{filename}sieve-of-eratosthenes-via-gopherjs-vue%en.rst>`_
.. [2] `Computed vs Watched Property - Computed Properties and Watchers — Vue.js <https://vuejs.org/v2/guide/computed.html#Computed-vs-Watched-Property>`_

.. _gopherjs-vue: https://github.com/oskca/gopherjs-vue
.. _Vue.js: https://vuejs.org/
.. _on my GitHub repo: https://github.com/siongui/frontend-programming-in-go/tree/master/031-watch-data-change-gopherjs-vue

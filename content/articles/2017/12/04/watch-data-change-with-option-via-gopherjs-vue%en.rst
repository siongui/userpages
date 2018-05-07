Watch Data Change With Options via Go and Vue.js
################################################

:date: 2018-05-07T23:33+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go,
       Vue.js, gopherjs-vue
:category: Frontend Programming in Go
:summary: Run data change watchers with option via Go/GopherJS/gopherjs-vue.
:og_image: https://cdn-images-1.medium.com/max/1024/1*kzBgrHhQRwEycJV8YQYR0A.png
:adsu: yes


In my `previous post`_ [1]_, we showed how to watch data changes using
Go/GopherJS/gopherjs-vue_. But sometimes we need to run watchers with option.
For example, we may want to run watchers immediately after initialization, and
default behavior does not run waters on initialization. In this case, we need to
set options as well when waters are set.

The code to set options is basically the same as the code of my previous post,
except an object representing options is declared and passed as argument while
setting the watchers.

The following patch shows the code difference that adding *{immediate: true}*
option to the watcher, which means running the watchers immediately after
initialization.

.. code-block:: go

  --- 031-watch-data-change-gopherjs-vue/app.go	2018-05-04 09:09:28.674070365 +0800
  +++ 032-watch-data-change-with-option-gopherjs-vue/app.go	2018-05-07 22:37:18.242253159 +0800
  @@ -24,8 +24,10 @@
   	// create the VueJS viewModel using a struct pointer
   	app := vue.New("#vueapp", m)

  +	option := js.Global.Get("Object").New()
  +	option.Set("immediate", true)
   	app.Call("$watch", "userinput", func(newVal, oldVal string) {
   		m.OldValue = oldVal
   		m.NewValue = newVal
  -	})
  +	}, option)
   }


The full code example is available `on my GitHub repo`_.

For other example and demo of using watchers with option using gopherjs-vue, see
`online Sieve of Eratosthenes`_ [2]_.

.. adsu:: 2

Tested on:

- ``Chromium 65.0.3325.181 on Ubuntu 17.10 (64-bit)``
- ``Go 1.10.2``
- ``GopherJS 1.10-3``

----

References:

.. [1] `Watch Data Change via Go and Vue.js <{filename}watch-data-change-via-gopherjs-vue%en.rst>`_
.. [2] `Online Sieve of Eratosthenes Demo via Go and Vue.js <{filename}sieve-of-eratosthenes-via-gopherjs-vue%en.rst>`_

.. _previous post: {filename}watch-data-change-via-gopherjs-vue%en.rst
.. _gopherjs-vue: https://github.com/oskca/gopherjs-vue
.. _Vue.js: https://vuejs.org/
.. _on my GitHub repo: https://github.com/siongui/frontend-programming-in-go/tree/master/032-watch-data-change-with-option-gopherjs-vue
.. _online Sieve of Eratosthenes: {filename}sieve-of-eratosthenes-via-gopherjs-vue%en.rst

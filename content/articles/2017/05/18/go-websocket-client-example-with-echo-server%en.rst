[GopherJS] WebSocket Client Example With Echo Server
####################################################

:date: 2017-05-18T02:28+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, WebSocket
:category: GopherJS
:summary: Simple example of a WebSocket client via GopherJS. Connect and send
          a message to the WebSocket server that echos everything from clients.
:og_image: https://1.bp.blogspot.com/-6DyCH5hPimo/WFYqCCVdALI/AAAAAAAAVss/qm71lQDiTPQUCtuFESwtRiZ07FLa4EXBwCLcB/s1600/overview.png
:adsu: yes


Simple example of a WebSocket client via GopherJS and Golang.
Connect and send a message to the WebSocket server that echos everything from
clients.

The WebSocket echo server is provided by *websocket.org* and the address is
*wss://echo.websocket.org*.

.. rubric:: `Run Code on GopherJS Playground <https://gopherjs.github.io/playground/#/huwdaAE0aJ>`_
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"

  	"github.com/gopherjs/gopherjs/js"
  )

  func main() {
  	ws := js.Global.Get("WebSocket").New("wss://echo.websocket.org")

  	ws.Call("addEventListener", "open", func(e *js.Object) {
  		fmt.Println("Connection open ...")
  		ws.Call("send", "Hello WebSockets!")
  	})

  	ws.Call("addEventListener", "message", func(e *js.Object) {
  		fmt.Println("Received Message: " + e.Get("data").String())
  		ws.Call("close")
  	})

  	ws.Call("addEventListener", "close", func(e *js.Object) {
  		fmt.Println("Connection closed.")
  	})
  }

Console output:

.. code-block:: txt

  Connection open ...
  Received Message: Hello WebSockets!
  Connection closed.

.. adsu:: 2

----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] `JavaScript Tips and Gotchas · gopherjs/gopherjs Wiki · GitHub <https://github.com/gopherjs/gopherjs/wiki/JavaScript-Tips-and-Gotchas>`_

.. [3] | `WebSocket 教程 - 阮一峰的网络日志 <http://www.ruanyifeng.com/blog/2017/05/websocket.html>`_
       | `WebSocket 教程 - WEB前端 - 伯乐在线 <http://web.jobbole.com/91321/>`_
.. [4] `A VPN-server, using websockets : golang <https://old.reddit.com/r/golang/comments/9zqn0d/a_vpnserver_using_websockets/>`_
.. [5] `peer to peer connection : golang <https://old.reddit.com/r/golang/comments/9zozqy/peer_to_peer_connection/>`_

.. _GopherJS: http://www.gopherjs.org/
.. _WebSocket: https://www.google.com/search?q=WebSocket

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js

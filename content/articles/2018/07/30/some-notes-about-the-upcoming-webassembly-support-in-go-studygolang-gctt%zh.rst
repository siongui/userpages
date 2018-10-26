關於 Go 即將支持的 WebAssembly 的一些注意事項
#############################################

:date: 2018-07-14T12:06:09+08:00
:tags: Go語言, Go語言中文網, WebAssembly, 轉錄
:category: Go語言
:author: themoonbear(譯者)
:summary: 這是一篇關於 webassembly 的即時記錄，它的目的是給我做個備忘而不僅僅是如何使用它的教程。
:og_image: https://static.oschina.net/uploads/space/2018/0415/083556_gBeH_2720166.png


這是一篇關於 webassembly 的即時記錄，它的目的是給我做個備忘而不僅僅是如何使用它的教程。

即將發佈的 Go 1.11 版本將支持 Wasm。@neelance 做了大部分的實施工作。對 wasm 的支持已經可以通過他在 github 上的工作分支進行測試。

看 `這篇文章`_ 瞭解更多信息


工具鏈設置
++++++++++

要從 go 源碼生產一個 wasm 文件，您需要從源碼獲取並為 go 工具集打補丁：

.. code-block:: bash

  ~ mkdir ~/gowasm
  ~ git clone https://go.googlesource.com/go ~/gowasm
  ~ cd ~/gowasm
  ~ git remote add neelance https://github.com/neelance/go
  ~ git fetch --all
  ~ git checkout wasm-wip
  ~ cd src
  ~ ./make.bash

然後使用這個版本的 Go，把 ``GOROOT`` 指到 ``~/gowasm`` 並使用 ``~/gowasm/bin/go`` 這個二進制文件。


第一個例子
++++++++++

按照慣例，讓我們先來寫個 “hello world”：

.. code-block:: go

  package main

  import "fmt"

  func main() {
  	fmt.Println("Hello World!")
  }

然後編譯出文件並命名 ``example.wasm``:

.. code-block:: bash

  GOROOT=~/gowasm GOARCH=wasm GOOS=js ~/gowasm/bin/go build -o example.wasm main.go


運行這個例子
++++++++++++

這是 `官方文檔`_ 的節選：

  雖然有計畫允許 WebAssembly 模塊像 ES6 模塊(...)那樣加載，但目前 WebAssembly 必須由 JavaScript 加載並編譯，對於基本的加載，有如下三個步驟：

  - 將 .wasm 字節轉化為數組或 ArrayBuffer
  - 將字節編譯為 WebAssemly.Module
  - 使用導入實例化的 WebAssembly.Module 以獲取可調用的導出

幸運的是，Go 的作者已經提供了一個 Javascript 加載器 ``~/gowasm/misc/wasm/wasm_exec.js`` 來簡化這個過程。它附帶一個 HTML 文件，負責粘貼瀏覽器中所有的內容。

要實際運行我們的文件，需要將以下文件複製到一個目錄中並作為 Web 服務啟動：

.. code-block:: bash

  ~ mkdir ~/wasmtest
  ~ cp ~/gowasm/misc/wasm/wasm_exec.js ~/wasmtest
  ~ cp ~/gowasm/misc/wasm/wasm_exec.html ~/wasmtest/index.html
  ~ cp example.wasm ~/wasmtest

然後編輯這個 ``index.html`` 文件去運行正確的例子：

.. code-block:: javascript

  // ...
  WebAssembly.instantiateStreaming(fetch("example.wasm"), go.importObject).then((result) => {
  	mod = result.module;
  	inst = result.instance;
  	document.getElementById("runButton").disabled = false;
  });
  // ...

理論上，任何 web 服務都可以運行它，但是當我們試著用 ``caddy`` 運行它時遇到一個問題。這個 javascript 加載器需要服務發送這個 wasm 文件的正確 mime 類型給它。

這有一個快速的破解方法來運行我們的測試：為我們的 wasm 文件寫個帶有特殊處理的 Go 服務。

.. code-block:: go

  package main

  import (
  	"log"
  	"net/http"
  )

  func wasmHandler(w http.ResponseWriter, r *http.Request) {
  	w.Header().Set("Content-Type", "application/wasm")
  	http.ServeFile(w, r, "example.wasm")
  }
  func main() {
  	mux := http.NewServeMux()
  	mux.Handle("/", http.FileServer(http.Dir(".")))
  	mux.HandleFunc("/example.wasm", wasmHandler)
  	log.Fatal(http.ListenAndServe(":3000", mux))
  }

*注意* 設置一個特殊的路由器來處理所有的 wasm 文件沒什麼大不了，如我所說，這是一個 POC，這篇文章只是關於它的附註。

然後使用 ``go run server.go`` 來啟動服務，並打開瀏覽器訪問 ``http://localhost:3000`` 。

打開控制台看看！


和瀏覽器交互
++++++++++++

讓我們和世界互動。


解決 DOM 問題
=============

``syscall/js`` 包中包含允許通過 javascript API 與 DOM 交互的函數。要獲取此包的文檔，只需運行：

.. code-block:: bash

  GOROOT=~/gowasm godoc -http=:6060

然後用瀏覽器訪問 ``http://localhost:6060/pkg/syscall/js/`` 。

讓我們寫個簡單的 HTML 文件來顯示一個輸入框。然後從 webassembly，我們給這個元素綁定一個事件，並在監聽到事件時觸發一個動作。

編輯 ``index.html`` 並把代碼放在 ``run`` 按鈕下面：

.. code-block:: html

  	<button onClick="run();" id="runButton" disabled>Run</button>
  	<input type="number" id="myText" value="" />
  </body>

然後修改 Go 文件：

.. code-block:: go

  package main

  import "fmt"

  func main() {
  	c := make(chan struct{}, 0)
  	cb = js.NewCallback(func(args []js.Value) {
  		move := js.Global.Get("document").Call("getElementById", "myText").Get("value").Int()
  		fmt.Println(move)
  	})
  	js.Global.Get("document").Call("getElementById", "myText").Call("addEventListener", "input", cb)
  	// The goal of the channel is to wait indefinitly
  	// Otherwise, the main function ends and the wasm modules stops
  	<-c
  }

像以前一樣編譯文件並刷新瀏覽器……打開控制台然後輸入一個數字……瞧瞧


暴露函數
++++++++

這有點辣手……我沒有找到任何簡單的方法將一個 Go 函數暴露給 Javascript 生態系統。我們需要做的是在 Go 文件中創建一個 ``Callback`` 對象並指定到一個 Javascript 對象。

為得到返回結果，我們不能返回一個值給 callback 而是使用 Javascript 對象代替。

這是新的 Go 代碼：

.. code-block:: go

  package main

  import (
  	"syscall/js"
  )

  func main() {
  	c := make(chan struct{}, 0)
  	add := func(i []js.Value) {
  		js.Global.Set("output", js.ValueOf(i[0].Int()+i[1].Int()))
  	}
  	js.Global.Set("add", js.NewCallback(add))
  	<-c
  }

現在編譯並運行代碼。打開瀏覽器和控制台。

如果你輸入 ``output`` 將返回 ``Object not found`` 。現在您輸入 ``add(2,3)`` 和 ``output`` ...應該得到 ``5`` 。

這不是很優雅的交互方式，但它按預期運行。


結論
++++

Go 對 wasm 的支持剛剛開始，但正大力發展。許多功能現在都可運行。我甚至可以在瀏覽器運行一個完整的遞歸神經網絡，這歸功於 Gorgonia。我將稍後講解這些。


----

via: `Some notes about the upcoming WebAssembly support in Go <https://blog.owulveryck.info/2018/06/08/some-notes-about-the-upcoming-webassembly-support-in-go.html>`_

作者： `Parikshit Agnihotry`_  譯者： themoonbear_  校對： polaris1119_

本文由 GCTT_ 原創編譯， `Go語言中文網`_ 榮譽推出

.. note::

  | 本文由 GCTT 原創翻譯， `Go語言中文網`_ 首發。也想加入譯者行列，為開源做一些自己的貢獻麼？歡迎加入 GCTT_ ！
  | 翻譯工作和譯文發表僅用於學習和交流目的，翻譯工作遵照 `CC-BY-NC-SA 協議規定`_ ，如果我們的工作有侵犯到您的權益，請及時聯繫我們。
  | 歡迎遵照 `CC-BY-NC-SA 協議規定`_ 轉載，敬請在正文中標注並保留原文/譯文鏈接和作者/譯者等信息。

----

- `关于 Go 即将支持的 WebAssembly 的一些注意事项  - Go语言中文网 - Golang中文社区 <https://studygolang.com/articles/13611>`_
- `關於 Go 即將支援的 WebAssembly 的一些注意事項 - 掃文資訊 <https://tw.saowen.com/a/f57e05ac77d1c7ab38d937ac951af71a229fdecddfe3054a87932ed47b562518>`_
- `如何將 Go 程式編譯成 WebAssembly | Tsung's Blog <https://blog.longwin.com.tw/2018/09/how-to-compile-golang-webassembly-wasm-2018/>`_

.. _這篇文章: https://blog.gopheracademy.com/advent-2017/go-wasm/
.. _官方文檔: https://webassembly.org/getting-started/js-api/
.. _Parikshit Agnihotry: https://medium.com/@parikshit
.. _themoonbear: https://github.com/themoonbear
.. _polaris1119: https://github.com/polaris1119
.. _GCTT: https://github.com/studygolang/GCTT
.. _Go語言中文網: https://studygolang.com/articles/13611
.. _CC-BY-NC-SA 協議規定: http://creativecommons.org/licenses/by-nc-sa/3.0/deed.zh

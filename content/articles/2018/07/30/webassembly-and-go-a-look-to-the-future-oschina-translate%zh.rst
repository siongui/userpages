WebAssembly 和 Go：對未來的觀望
###############################

:date: 2018-07-25T17:55+08:00
:tags: Go語言, 轉錄, 開源中國, WebAssembly
:category: Go語言
:author: 開源中國翻譯
:summary: 我反對學習 JavaScript 還有前端開發已經不是秘密了。事實上，在 CSS 出現前我就學會了 HTML，不過 JavaScript 是我做 Web 開發好久後的事情了。當看到現代 Web 的發展時，我感到不寒而栗。這個生態對於脫離已久的我來說是如此迷茫。Node, webpack, yarn, npm, frameworks, UMD, AMD，我的天啊！
:og_image: https://oscimg.oschina.net/oscnet/66cd656f4ef5646d1268db73cd5c7343d27.jpg

.. note::

  | 英文原文： `Web Assembly and Go: A look to the future`_
  | 轉錄來源： `WebAssembly 和 Go：对未来的观望 - 开源中国`_
  | 參與翻譯 (5人) : `琪花亿草`_, xiaoaiwhc1_, lnovonl_, kevinlinkai_, `fly_白`_

我反對學習 JavaScript 還有前端開發已經不是秘密了。事實上，在 CSS 出現前我就學會了 HTML，不過 JavaScript 是我做 Web 開發好久後的事情了。當看到現代 Web 的發展時，我感到不寒而栗。這個生態對於脫離已久的我來說是如此迷茫。Node, webpack, yarn, npm, frameworks, UMD, AMD，我的天啊！

目前我關注 WebAssembly 也已經有段時間了，期望它能讓我在沒有典型 JavaScript 構建的情況下編寫 Web 應用程序。

當聽到 WebAssembly(wasm) 最近支持 Go 語言時，我知道實驗的時機已經成熟，並且迫切期待嘗試。在嘗試之前我讀了些好文章，而這篇文章將記錄我的一些體驗。

為了用 Go 來寫 wasm，你需要先下載 Go 源碼並編譯好。從 Go 1.11 開始，WebAssembly 將被原生支持，但現在還沒有 release。

你可以按照 `這裏`_ 的步驟來編譯 Go。因為 Go 本身也是用 Go 語言實現的，所以在編譯之前你需要先有一個可以正常工作的 Go 二進制版本來自舉自己。最終，你系統裏會有兩個不同的 Go 版本。注意：如果你後面忘了你系統裏安裝了兩個版本的 Go，那可能會給你造成一些困擾。可以使用 direnv_ 來管理 Go 版本，這樣你就可以為不同的項目來配置不同的 Go 了。

安裝最新的 Go 後，就可以體驗 WebAssembly 了。你需要一個 HTML 文件和一個 JavaScript 腳本來加載生成的 wasm 文件。這些都包含在 Go 安裝路徑下的 misc/wasm 目錄裏。你可以複製它們到項目目錄，修改它們以加載你的 wasm 文件。

我的第一個項目有點雄心勃勃，我打算用 Go 語言構建一個看起來像 `Web 組件`_ 的東西，編譯成 WebAssembly。我並沒有把整件事做完，因為我被每件事要如何都做得好弄得心煩意亂。

首先，我將 GOROOT/misc/wasm 中的 HTML 和 JavaScript 文件複製到一個新目錄中，並添加了一個 main.go 文件。根據我預先想好的計劃，我把 HTML 放進 DOM 的一個現有節點，這個 DOM 要在 HTML 中聲明。所以我創建了一個帶有 thing 作為 ID 的 HTML section 標簽。

.. code-block:: html

  <section class="main" id="thing" >Please wait...</section>

我在 HTML 文件底部的腳本標簽上面插入了這個。接下來，我知道我想程序化地替換這個節點，所以我查找了 Go 的 wasm 庫中與 DOM 交互的語法。為 Go 添加了一個 syscall/js 包，允許與 DOM 進行交互。我使用了這段 Go 代碼得到了一個 HTML 帶有 thing 作為 ID 的節點的引用：

.. code-block:: go

  el := js.Global.Get("document").Call("getElementById", "thing")

現在我有一個空DOM節點的引用，我可以使用渲染的HTML來填充。因此下一步其實就是創建一些HTML並將其填充進去。

我將著名的TodoMVC應用作為靈感。首先我創建兩個文件：todo.go和todolist.go。這些文件包含一些Go結構來表示Todo事項，和Todo事項列表。

.. code-block:: go

  type Todo struct {
  	Title     string
  	Completed bool
  	//Root      js.Value
  	tree *vdom.Tree
  }

  type TodoList struct {
  	Todos []Todo
  	Component
  }

  type Component struct {
  	Name     string
  	Root     js.Value
  	Tree     *vdom.Tree
  	Template string
  }

我也有點自大，開始將東西提取到Component類型中，並認為我可以將它嵌入到我的自定義類型中，以便向它們提供Web 組件功能。 *我沒有完成這個想法。。。在後文你會看到原因。*

這些自定義Go類型每一個都有一個Render()方法和模板：

.. code-block:: go

  var todolisttemplate = `<ul>
  {{range $i, $x := $.Todos}}
  	{{$x.Render}}
  {{end}}
  </ul>`

|

.. code-block:: go

  func (todoList *TodoList) Render() error {

  	tmpl, err := template.New("todolist").Parse(todoList.Template)
  	if err != nil {
  		return err
  	}
  	// Execute the template with the given todo and write to a buffer
  	buf := bytes.NewBuffer([]byte{})
  	if err := tmpl.Execute(buf, todoList); err != nil {
  		return err
  	}
  	// Parse the resulting html into a virtual tree
  	newTree, err := vdom.Parse(buf.Bytes())
  	if err != nil {
  		return err
  	}

  	if todoList.Tree != nil {
  		// Calculate the diff between this render and the last render
  		//	patches, err := vdom.Diff(todo.tree, newTree)
  }		//	if err != nil {
  		//		return err
  		//	}

  		// Effeciently apply changes to the actual DOM
  		//		if err := patches.Patch(todo.Root); err != nil {
  		//			return err
  		//		}
  	} else {

  		todoList.Tree = newTree
  	}
  	// Remember the virtual DOM state for the next render to diff against
  	todoList.Tree = newTree

  	todoList.Root.Set("innerHTML", string(newTree.HTML()))
  	return nil
  }

我的想法是用我找到的 vdom_ 包來做這些渲染，這樣的話渲染的效率會更高一些。這就是我遇到的第一個問題。

GopherJS和Go/wasm之間的區別
+++++++++++++++++++++++++++

vdom包專為 GopherJS_ 而寫，而GopherJS是一個從Go到Javascript的轉譯器。基於便捷，GopherJS使用js.Object類型。Go的新wasm庫syscall/js使用js.Value類型。它們精神上是相似的，但在實現上大為不同。這意味著我使用vdom渲染的想法是行不通的，除非我將vdom使用的js.Object移植到使用js.Value。盡管vdom的tree.HTML()函數在不用修改的情況下就可以運行，因此我可以將HTML節點的內部HTML設置為vdom解析出的內容。Render()函數解析Go結構模板，將Go結構的實例作為上下文來傳值。然後它用vdom庫創建一個解析dom樹，而且在函數的最後一行渲染樹：

.. code-block:: go

  todoList.Root.Set("innerHTML", string(newTree.HTML()))

此時，我已經有了一個可以運行的Go/wasm原型，沒有連接任何事件。但是它確實可以渲染成dom並顯示在瀏覽器。這是巨大的一步；我當時很興奮。

我創建了一個Makefile，這樣我就不用一次又一次的輸入冗長的編譯命令：

.. code-block:: makefile

  wasm2:
  	GOROOT=~/gowasm GOARCH=wasm GOOS=js ~/gowasm/bin/go build -o example.wasm markdown.go

  wasm:
  	GOROOT=~/gowasm GOARCH=wasm GOOS=js ~/gowasm/bin/go build -o example.wasm .

  build-server:
  	go build -o server-app server/server.go

  run: build-server wasm
  	./server-app

基於現在的Web Assembly狀態，這個makefile也指出了一個至關重要的問題。新型瀏覽器會忽略WASM文件，除非給他們提供合適的MIME類型。 `這篇文章`_ 有一個簡單的HTTP文件服務器，它為web assembly文件設置了正確的MIME類型。我將其複製到我的項目，並將其用於應用中。如果你的web服務器確實為.wasm文件配置好了，那麼你就不需要自定義服務器。


提出挑戰
++++++++

在這一點上，我意識到Web Assembly可以正常運行，而也許更重要的是：GopherJS的很多代碼很少甚至不用修改就可以在Web Assembly可以正常運行。我給自己提出挑戰（ `nerd sniped`_ ）。我嘗試的下一件事情是找一個 vecty_ 應用並編譯它。由於vecty是專為GopherJS所寫，而且使用了js.Object類型而不是js.Vaule，因此要想失敗很困難。為了讓vecty在wasm中編譯，我 `fork了vecty`_ ，然後做了一些修改，一些處理，並註釋了很多代碼。

最終的結果就是放在在vecty/example目錄中的markdown編輯器可以在Web Assembly中完美運行。本文有點冗長，因此我會讓你 `在這`_ 看源碼。總結：它與GopherJS版本幾乎完全相同，但是在main()退出的時候web assembly也會退出，因此為了阻止退出並保持應用運行，我在main()結尾添加了一個空的通道接收。


事件
++++

Go 的 syscall/js 使用了一個非常不同的方法來進行事件註冊，我不得不修改 vecty 的事件 `註冊代碼`_ 才能使用 wasm 新的回調註冊，在這裏我花了非常多的時間。不過直到現在，這個方法工作的還不錯。


結論
++++

.. image:: https://oscimg.oschina.net/oscnet/66cd656f4ef5646d1268db73cd5c7343d27.jpg
   :alt: WebAssembly 和 Go：對未來的觀望
   :align: center

通過對這些事件課程的學習，我認定 WebAssembly 就是 Web 開發的未來。它可以使用任何語言作為“前端語言”來進行 Web 開發，然後編譯為 wasm 就可以了。這給像我一樣並不想再學習 Javascript，而可以使用自己喜歡的語言來進行 Web 開發的人帶來了很多好處。

*你可以從* `這裏 <https://github.com/bketelsen/wasmplay>`__ *下載源代碼，不過記住：風險自擔。*

.. note::

  | 本文中的所有譯文僅用於學習和交流目的，轉載請務必註明文章譯者、出處、和本文鏈接。
  | 我們的翻譯工作遵照 `CC 協議`_ ，如果我們的工作有侵犯到您的權益，請及時聯系我們。

----

- `WebAssembly 和 Go：對未來的觀望 - 掃文資訊 <https://tw.saowen.com/a/08a24b7f3544d5f7e7668eccbf4c3a075880b814f1a5b0310f67f3b49dc0c42f>`_
- `如何將 Go 程式編譯成 WebAssembly | Tsung's Blog <https://blog.longwin.com.tw/2018/09/how-to-compile-golang-webassembly-wasm-2018/>`_

.. _Web Assembly and Go\: A look to the future: https://brianketelsen.com/web-assembly-and-go-a-look-to-the-future/
.. _WebAssembly 和 Go：对未来的观望 - 开源中国: https://www.oschina.net/translate/webassembly-and-go-a-look-to-the-future
.. _琪花亿草: https://my.oschina.net/u/146173
.. _xiaoaiwhc1: https://my.oschina.net/xiaoaiwhc
.. _lnovonl: https://my.oschina.net/u/3680260
.. _kevinlinkai: https://my.oschina.net/u/2267262
.. _fly_白: https://my.oschina.net/flyWhite
.. _這裏: https://golang.org/doc/install/source
.. _direnv: http://direnv.net/
.. _Web 組件: https://www.webcomponents.org/
.. _vdom: https://github.com/albrow/vdom
.. _GopherJS: https://gopherjs.org/
.. _這篇文章: https://blog.owulveryck.info/2018/06/08/some-notes-about-the-upcoming-webassembly-support-in-go.html
.. _nerd sniped: https://xkcd.com/356/
.. _vecty: https://github.com/gopherjs/vecty
.. _fork了vecty: https://github.com/gowasm/vecty
.. _在這: https://github.com/bketelsen/wasmplay/tree/master/markdownvecty
.. _註冊代碼: https://github.com/gowasm/vecty/blob/wasm-wip/dom.go#L231
.. _CC 協議: http://zh.wikipedia.org/wiki/Wikipedia:CC

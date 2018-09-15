Go 1.11 正式發佈：添加對模塊 "modules" 的初步支持
#################################################

:date: 2018-08-25T00:00+08:00
:tags: Go語言, 轉錄, 開源中國
:category: Go語言
:author: 局長(開源中國)
:summary: 美國當地時間8月24日，Go 開發團隊 `宣佈推出 Go 1.11 正式版`_ 。
:og_image: https://www.teamwork.com/assets/images/pages/internship/golang-icon.png


美國當地時間8月24日，Go 開發團隊 `宣佈推出 Go 1.11 正式版`_ 。

下載地址 >>> https://golang.org/dl/ or https://golang.google.cn/dl/（免梯子）

看看有哪些值得關注的更新 ——

新版本在工具鏈、運行時和庫都有許多變化和改進。當然，兼容性方面依然會兼容舊版本。此外有兩個最值得關注且激動人心的新特性 —— 對模塊(modules)和 WebAssembly 的支持。

Go Modules
==========

此版本增加了對被稱作“模塊(Go Modules)”的初步支持，這是 GOPATH 的替代方案，集成了對版本控制和軟件包分發的支持。該功能目前仍處於實驗性階段，並且仍有一些可能會影響使用的問題，因此請隨意使用 `問題跟蹤器`_ 進行反饋或查閱解決方案。

WebAssembly
===========

Go 1.11 還為 WebAssembly (js/wasm) 添加了一個處於實驗性階段的端口。它使得開發者可將 Go 程序編譯為與四個主流 Web 瀏覽器兼容的二進制格式。可以在 `webassembly.org`_ 上閱讀有關 WebAssembly（縮寫為"WASM"）的更多信息。

Go 程序現在可被編譯為一個 WebAssembly 模塊，該模塊包括用於 goroutine 調度、垃圾收集、映射等的 Go 運行時。因此，生成的模塊大小約為 2MB，或壓縮後為 500KB 左右。Go 程序也可以使用新的實驗性 `syscall/js`_ 包調用 JavaScript。有關二進制包大小和與其他語言互操作的問題尚未成為優先處理事項，不過會在將來的版本中得到解決。

有關 Go 1.11 中變更的更多詳細信息，請參閱發行說明 >>> https://golang.org/doc/go1.11

|
| 本站文章除註明轉載外，均為本站原創或編譯。歡迎任何形式的轉載，但請務必註明出處，尊重他人勞動共創開源社區。
| 轉載請註明：文章轉載自 開源中國社區 [http://www.oschina.net]
| 本文標題：Go 1.11 正式發佈：添加對模塊 "modules" 的初步支持
| 本文地址：https://www.oschina.net/news/99309/go-1-11-released
|

.. _宣佈推出 Go 1.11 正式版: https://blog.golang.org/go1.11
.. _問題跟蹤器: https://golang.org/issue/new
.. _webassembly.org: https://webassembly.org/
.. _syscall/js: https://golang.org/pkg/syscall/js/

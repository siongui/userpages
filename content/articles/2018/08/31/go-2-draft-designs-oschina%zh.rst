Go 公佈 2.0 設計草案：主打規模化和擴展性，支持泛型
##################################################

:date: 2018-08-31T00:00+08:00
:tags: Go語言, 轉錄, 開源中國
:category: Go語言
:author: 王練(開源中國)
:summary: 去年7月，Go 語言官博就曾透露 `Go 2 開發計畫`_ ，並表示 Go 2 的目標就是解決 Go 1.x 在規模化方面做的還不夠好的地方。隨著時間的推進，開發團隊已著手準備 2.0 版本的開發工作，並公佈了 `設計草案`_ ，供社區討論和反饋，以促進最終的語言設計。
:og_image: https://oscimg.oschina.net/oscnet/787971faf8cdd0a924b163eb4c48aa3a95b.jpg


去年7月，Go 語言官博就曾透露 `Go 2 開發計畫`_ ，並表示 Go 2 的目標就是解決 Go 1.x 在規模化方面做的還不夠好的地方。隨著時間的推進，開發團隊已著手準備 2.0 版本的開發工作，並公佈了 `設計草案`_ ，供社區討論和反饋，以促進最終的語言設計。

.. image:: https://oscimg.oschina.net/oscnet/787971faf8cdd0a924b163eb4c48aa3a95b.jpg
   :alt: Go 公佈 2.0 設計草案：主打規模化和擴展性，支持泛型
   :align: center

設計草案包含三個方面，錯誤處理、錯誤值和泛型，並針對各個方面進行了詳細的概述和改進草案。大致總結如下：


一、錯誤處理（Error handling）
==============================

為擴展至大型代碼庫，Go 程序必須是輕量級的，不會過度重複，且具備穩健性，能夠優雅地處理出現的錯誤。

目前 Go 檢查錯誤的代碼太多，但處理這些錯誤的代碼卻嚴重不足。對於 Go 2，開發團隊希望錯誤檢查更加輕量級，減少用於錯誤檢查的 Go 程序的文本量。此外，還能更加方便地編寫錯誤處理程序，提高開發者處理錯誤的可能性。

為避免處理重複異常，錯誤檢查和錯誤處理還必須是顯性的，在程序文本中可見。

參考示例：

.. code-block:: go

  func main() {
  	hex, err := ioutil.ReadAll(os.Stdin)
  	if err != nil {
  		log.Fatal(err)
  	}

  	data, err := parseHexdump(string(hex))
  	if err != nil {
  		log.Fatal(err)
  	}

  	os.Stdout.Write(data)
  }

簡化後：

.. code-block:: go

  func main() {
  	handle err {
  		log.Fatal(err)
  	}

  	hex := check ioutil.ReadAll(os.Stdin)
  	data := check parseHexdump(string(hex))
  	os.Stdout.Write(data)
  }


二、錯誤值（Error values）
==========================

大型程序必須能夠以編程方式測試和響應錯誤，並且還能很好地報告它們。

目前的各種流行的助手工具包添加了超出標準錯誤接口的功能，但它們以不兼容的方式執行。對於 Go 2，開發團隊考慮將“可選接口”標準化，以允許這些工具包進行互操作，並慢慢減少對它們的需求。

改進主要包含兩個目標：一是讓程序的錯誤檢查更容易，更不容易出錯，以提高程序的錯誤處理和穩健性；二是希望能夠以標準格式打印包含額外細節的錯誤。

.. code-block:: txt

  // Is reports whether err or any of the errors in its chain is equal to target.
  func Is(err, target error) bool// As checks whether err or any of the errors in its chain is a value of type E.// If so, it returns the discovered value of type E, with ok set to true.// If not, it returns the zero value of type E, with ok set to false.
  func As(type E)(err error) (e E, ok bool)


三、泛型（Generics）
====================

想要擴展到大型代碼庫，代碼的可重用性非常重要。

Go 團隊在早期其實一直有在調查和討論“泛型”的可能性設計，但由於種種原因，Go 1 更多的是確保能快速構建包含很多獨立軟件包的程序。

Go 2 的目標是通過允許帶有類型參數的參數多態來解決編寫 Go 庫的問題，這些問題抽象出了不必要的類型細節。

此外，除了預期的容器類型之外，開發團隊還希望能編寫有用的庫來操作任意的 map 和 channel 值。理想方案是編寫能夠同時操作 [ ]byte 和 string 值的多態函數。

.. code-block:: go

  type List(type T) []T

  func Keys(type K, V)(m map[K]V) []K


`更多細節請查閱設計草案頁面。`_

|
| 本站文章除註明轉載外，均為本站原創或編譯。歡迎任何形式的轉載，但請務必註明出處，尊重他人勞動共創開源社區。
| 轉載請註明：文章轉載自 開源中國社區 [http://www.oschina.net]
| 本文標題：Go 公佈 2.0 設計草案：主打規模化和擴展性，支持泛型
| 本文地址：https://www.oschina.net/news/99490/go-2-draft-designs
|

----

- `Sydney Golang Meetup - Rob Pike - Go 2 Draft Specifications : golang <https://redd.it/9www5f>`_
- `Chris Siebenmann: Go 2 contracts are too clever (via Golang Ninjas newsletter) : golang <https://old.reddit.com/r/golang/comments/a06v2g/chris_siebenmann_go_2_contracts_are_too_clever/>`_
- `Golang made teeny-tiny: Go compiler for microcontrollers : golang <https://old.reddit.com/r/golang/comments/a0v7tu/golang_made_teenytiny_go_compiler_for/>`_
- `Go 2, here we come! : golang <https://old.reddit.com/r/golang/comments/a1j3h6/go_2_here_we_come/>`_
- `Go 2, here we come | Hacker News <https://news.ycombinator.com/item?id=18561587>`_

.. _Go 2 開發計畫: https://www.oschina.net/news/86774/toward-go2
.. _設計草案: https://go.googlesource.com/proposal/+/master/design/go2draft.md
.. _更多細節請查閱設計草案頁面。: https://go.googlesource.com/proposal/+/master/design/go2draft.md

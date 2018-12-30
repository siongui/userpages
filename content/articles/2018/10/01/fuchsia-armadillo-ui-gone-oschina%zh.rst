Google 移除了 Fuchsia 中代號 Armadillo 的系統 UI
################################################

:date: 2018-12-30
:tags: 轉錄, 開源中國, Fuchsia
:category: 轉錄
:author: 局长(開源中國)
:summary: 過去兩年，谷歌一直在開發一個名為 Fuchsia 的全新開源操作系統。對於這款操作系統，谷歌方面主動透露的消息並不多，但它是作為一個開源項目開發的，因此外界總能跟蹤它的開發進度。
:og_image: https://oscimg.oschina.net/oscnet/7456103879d85a746f86e22fea4e9bb871b.jpg

過去兩年，谷歌一直在開發一個名為 Fuchsia 的全新開源操作系統。對於這款操作系統，谷歌方面主動透露的消息並不多，但它是作為一個開源項目開發的，因此外界總能跟蹤它的開發進度。

與 Android 和 Chrome OS 不同，Fuchsia 並不基於 Linux 內核 —— 它使用了一個全新的、谷歌開發的微內核，稱作 " Zircon_ "。Fuchsia 不僅“拋棄”Linux 內核，還可以不使用 GPL 開源許可證：該系統使用 `BSD 3 clause`_ , MIT_ , 和 `Apache 2.0`_ 三者組合的開源許可證。

.. _Zircon: https://www.oschina.net/news/88782/google-renames-magenta-kernel-to-zircon
.. _BSD 3 clause: https://en.wikipedia.org/wiki/BSD_licenses#3-clause
.. _MIT: https://en.wikipedia.org/wiki/MIT_License
.. _Apache 2.0: https://en.wikipedia.org/wiki/Apache_License#Version_2.0

去年它 `被發現`_ 引入了一個新系統 UI。Fuchsia 的 UI 層使用的是 Dart 語言開發的 Flutter SDK，Flutter 可以提供跨平台的在 Android 和 iOS 上運行的代碼。系統 UI 代號為 Armadillo。有人甚至設法在 Pixelbook 上演示了 Armadillo。

.. _被發現: https://www.solidot.org/story?sid=52333

.. image:: https://oscimg.oschina.net/oscnet/7456103879d85a746f86e22fea4e9bb871b.jpg
   :alt: Google 移除了 Fuchsia 中代號 Armadillo 的系統 UI
   :align: center

.. container:: align-center

  名為 Armadillo 的 Fuchsia 系統 UI 運行截圖

然而對於這樣一個有著不錯視覺效果的新系統 UI，谷歌似乎不太滿意。有人發現谷歌在最近的代碼變更中 `完全移除了`_ Armadillo，開發者甚至起了一個標題叫“ `Armadillo fainted!`_ ”，Armadillo 現在被面向開發者的 Shell Ermine 替代了。

.. _完全移除了: https://tech.slashdot.org/story/18/12/28/2155257/everything-we-knew-about-fuchsias-ui-armadillo-is-gone
.. _Armadillo fainted!: https://fuchsia-review.googlesource.com/c/topaz/+/235313

對於未來，Fuchsia 大部分 UI 開發的工作都可在名為 “ `vendor/google`_ ” 的封閉源代碼庫中找到。而根據公共代碼的評論，我們能獲知至少有三個新的 “shells” (即系統的 UI)在開發中，分別是 Redditor mishudark,  `Dragonglass 和 Flamingo`_ 。

.. _vendor/google: https://fuchsia-review.googlesource.com/c/topaz/+/134122/7#message-63f260f8522ceebc32af082f5d11a16da59f2a07
.. _Dragonglass 和 Flamingo: https://fuchsia-review.googlesource.com/c/peridot/+/229297

..
  .. image:: 
   :alt: 
   :align: center

.. highlights::

  | 本站文章除註明轉載外，均為本站原創或編譯。歡迎任何形式的轉載，但請務必註明出處，尊重他人勞動共創開源社區。
  | 轉載請註明：文章轉載自 開源中國社區 [https://www.oschina.net]
  | 本文標題：Google 移除了 Fuchsia 中代號 Armadillo 的系統 UI
  | 本文地址：https://www.oschina.net/news/103168/fuchsia-armadillo-ui-gone

----

`Solidot | Google 移除了 Fuchsia 中代号 Armadillo 的系统 UI <https://www.solidot.org/story?sid=59128>`_


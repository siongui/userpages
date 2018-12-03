Fuchsia 操作系統的四層結構設計
##############################

:date: 2018-11-29T09:59+08:00
:tags: 轉錄, 開源中國, Fuchsia
:category: 轉錄
:author: h4cd(開源中國)
:summary: 最近很多朋友很關注 Fuchsia 操作系統，也有很多朋友關心它的開發，今天我們就來講講它的結構，在描述 Fuchsia OS的結構的時候，谷歌使用了“層蛋糕”的比喻，原文是 Layer Cake。Fuchsia OS是整個項目的名稱和成品的名稱，但在其下它被分成4個不同的層次
:og_image: https://oscimg.oschina.net/oscnet/3ee747860fd479212cd8dc96058bdf241b4.jpg

最近很多朋友很關注 Fuchsia 操作系統，也有很多朋友關心它的開發，今天我們就來講講它的結構，在描述 Fuchsia OS的結構的時候，谷歌使用了“層蛋糕”的比喻，原文是 Layer Cake。Fuchsia OS是整個項目的名稱和成品的名稱，但在其下它被分成4個不同的層次，如下圖所示：

.. image:: https://oscimg.oschina.net/oscnet/3ee747860fd479212cd8dc96058bdf241b4.jpg
   :alt: Fuchsia 操作系統的四層結構設計
   :align: center

第一層：也是最底下一層，是構建 Fuchsia OS 的基石，Zircon 內核，去年的新聞是叫 Magenta，但是後來改為了 Zircon 這個名字，這是一個由Google全新設計的新內核，主要處理硬件訪問和軟件之間的通信。

對於不太瞭解內核作用的同學簡而言之，Zircon之於Fuchsia，恰如Linux之餘於Android。Linux內核驅動了多個操作系統，很多操作系統構建在它之上，比如 Ubuntu、Android、Manjaro、ArchLinux、Debian、Red Hat、SUSE 甚至 Chrome OS ，所以我們也可以大膽預測，如果未來Fuchsia OS 發展良好， Zircon 內核也被證明好用，那麼很有可能有更多的操作系統採用這一新內核。

第二層：也是直接構建在 Zircon 上的一層名叫 Garnet。 Garnet 包含各種操作系統所需的各種底層功能，包括硬件的驅動程序（網絡，圖形等）和軟件安裝。這一層最激動人心的事情是 Escher（圖形渲染器），Amber（Fuchsia 的更新程序）和Xi Core，它是Xi文本和代碼編輯器的底層引擎（今年早些時候已經發佈了）。

第三層：Peridot 是接下來的這一層，主要處理Fuchsia的模塊化應用程序設計， Peridot的另外兩個主要組件直接用於模塊。 Ledger 可以跨設備保存您在應用/模塊中的位置，並同步到您的Google帳戶。Maxwell 是一個更複雜的主題，需要更多進一步的深入研究，但是 Maxwell 極有可能是讓 Fuchsia 充分施展魔力的點睛之筆，可以提前透露的是，Maxwell 的厲害之處包括 Kronk，也是大家熟知的 Google Assistant。

第四層：Topaz，是這個 Layer Cake 蛋糕的頂層，也是對開發者和用戶直接影響最大的一層。Topaz 提供 Flutter 支持，而有了Flutter 的支持，各種華麗的應用程序，可以幫助充實地提供日常使用的功能齊全的應用程序。比如，現在最令人印象深刻的當然是 Armadillo UI，它是 Fuchsia 主要用戶界面和主屏幕。

可以做一個類比，Topaz 這一層在 Android 中可以找到一個對照，這將是你的必備應用程序，如聯繫人，音樂，文件管理器和文本編輯器 Xi（Topaz中的可視前端連接到Garnet的後端）。即使沒有你需要的東西，你也可以簡單方便地安裝。

從表面上看，Fuchsia OS 的層次設計似乎更適合團隊組織。它還有助於將代碼劃分成不同的更易於理解和開發的部分。當然，還有比這更重要的原因，硬件供應商也可以擁有自己可以掌控的層次結構。這意味著公司將能夠用他們自己的修改版本來替換四個層中的一個（或多個）。

大多數手機製造商會定製Android用戶體驗，以便從競爭中脫穎而出，而不是使用 Google 的默認界面設計。自定義設計的能力進一步表明 Google 正在借鑑從 Android 中得到的經驗。Fuchsia 和 Android 使供應商更容易使用他們的自定義設置和UI設計，而不影響系統的其餘部分。例如，三星可以用 TouchWiz 主題版本取代 Topaz 層，HTC、華碩和其他手機製造也可以。

此外，Android 團隊無法預料像 Amazon 這樣的公司會大量修改 Android 來用作 Kindle Fire 設備的操作系統。使用 Fuchsia OS 的話，同樣的事情變得更簡單，比如用 Amazon Web Services 和 Alexa 替換 Peridot 的Google Cloud 和 Google 智能助理，當然，用 亞馬遜的設計語言取代 Topaz ，但是不影響獲得 Fuchsia 的 Zircon 和 Garnet 的更新。

總而言之，從 Fuchsia 的設計可以看出，Google 正在銘記 其 Android 團隊學到的很多寶貴教訓與經驗。與 Android 的撕裂的生態系統相比，將 Fuchsia OS 分成僅向上構建的層次結構應有助於確保可更新性和統一性。當然，Google 仍然可以改變這種設計，只有時間才能證明 Fuchsia OS 的計畫是否會取得成功。

來源： `Fuchsia OS 中文社區`_

.. highlights::

  | 本站文章除註明轉載外，均為本站原創或編譯。歡迎任何形式的轉載，但請務必註明出處，尊重他人勞動共創開源社區。
  | 轉載請註明：文章轉載自 開源中國社區 [https://www.oschina.net]
  | 本文標題：Fuchsia 操作系統的四層結構設計
  | 本文地址：https://www.oschina.net/news/102223/the-4-layers-of-fuchsia

.. _Fuchsia OS 中文社區: https://fuchsia-china.com/the-4-layers-of-fuchsia/zh-cn/

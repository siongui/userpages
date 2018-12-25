Ubuntu 2018 回顧：從內存洩露到“感人”的 LTS 版本
###############################################

:date: 2018-12-24T08:33+08:00
:tags: 轉錄, 開源中國, Ubuntu Linux
:category: Linux
:author: 局长(開源中國)
:summary: Ubuntu 從 2018 年開始就一直十分活躍 —— 因為在 2018 年到來前的兩個月發佈的 Ubuntu 17.10 ‘Artful Aardvark’ 帶來了一波接著一波的討論。但隨著一個新的長期支持版本迫在眉睫、對重新設想的安裝器的爭論不休，以及 GNOME Shell 的大量內存洩露問題……2018 年對於 Ubuntu 工程師來說並不容易。
:og_image: https://oscimg.oschina.net/oscnet/30ad5cd5bb96627a3b8d3510ac74389bf4d.jpg


Ubuntu 從 2018 年開始就一直十分活躍 —— 因為在 2018 年到來前的兩個月發佈的 Ubuntu 17.10 ‘Artful Aardvark’ 帶來了一波接著一波的討論。但隨著一個新的長期支持版本迫在眉睫、對重新設想的安裝器的爭論不休，以及 GNOME Shell 的大量內存洩露問題……2018 年對於 Ubuntu 工程師來說並不容易。

本文將和你回顧一下 Ubuntu 在 2018 年的關鍵事件……

Ubuntu in 2018
++++++++++++++

一月：To-Do App
===============

.. image:: https://oscimg.oschina.net/oscnet/30ad5cd5bb96627a3b8d3510ac74389bf4d.jpg
   :alt: Ubuntu 2018 回顧：從內存洩露到“感人”的 LTS 版本
   :align: center

1月份，開發者確認計畫 `將 GNOME To-Do 應用程序添加`_ 到 Ubuntu 18.04 的默認應用中。GNOME To Do 提供基本的待辦事項和筆記功能，還整合了多種在線服務，使其成為以生產力為重點的 LTS 版本的完美補充。

另外，Ubuntu 開發者也證實，他們會在 Ubuntu 18.04 中 `使用 Xorg 作為默認顯示服務器`_ 。當然，Wayland 仍然是預裝的，因為 Ubuntu 17.10 與 Wayland 的搭配嘗試取得了不錯的效果。但在全新的安裝系統中，Xorg 將成為默認的選擇。

說到 Ubuntu 17.10，1月11日 Canonical `重新發佈了`_ Ubuntu 17.10，主要是因為在 `某些筆記本電腦上出現 BIOS 固件問題`_ ，導致其電腦“變磚”。

在社區的其他地方，GNOME 確認其計畫從 `Nautilus 中刪除對桌面圖標的支持`_ 。這一有爭議的變化將使 Ubuntu 選擇在 2018 年的所有版本中發佈舊版本的著名文件管理器。

.. _將 GNOME To-Do 應用程序添加: https://www.oschina.net/news/92865/ubuntu-18-04-new-app-to-do
.. _使用 Xorg 作為默認顯示服務器: https://www.oschina.net/news/92814/xorg-will-default-display-server-ubuntu-18-04-lts
.. _重新發佈了: https://www.oschina.net/news/92400/ubuntu-17-10-fix-bios-released
.. _某些筆記本電腦上出現 BIOS 固件問題: https://www.oschina.net/news/91749/ubuntu-17-10-stop-download
.. _Nautilus 中刪除對桌面圖標的支持: https://www.omgubuntu.co.uk/2018/01/gnome-desktop-icons-removed-3-28


二月：最小化安裝(Minimal Install)
=================================

Ubuntu `開始收集用戶數據`_ 的消息讓很多人感到擔憂，但是一種前瞻性和誠實的方法緩解了任何可能引起爭議的問題 —— Canonical 在即將發佈的 Ubuntu 18.04 LTS 版本中 `提供了一個“最小化安裝”選項`_ 。

二月份對於 Snapcraft 來說也是一個淘汰賽，這是 Canonical 新的應用程序格式。今年大量軟件廠商都採用了這種格式，包括微軟、Mozilla 和 JetBrains。

.. _開始收集用戶數據: https://www.oschina.net/news/93409/ubuntu-data-collection-opt-out
.. _提供了一個“最小化安裝”選項: https://www.oschina.net/news/93473/ubuntu-18-04-minimal-install


三月：大規模的內存洩露問題
==========================

在發現一個主要的 `GNOME Shell 內存洩露`_ 問題之後，三月份，GNOME Shell 桌面成為大家關注的焦點。Linux 社區陷入困境，Ubuntu 開發者也加入到 GNOME 開發團隊以幫助縮小範圍 `並解決問題`_ 。在同一個月，Ubuntu 設計師 `推出了`_ 一款基於海狸的默認壁紙，而 Ubuntu 工程師則將最小的安裝映像 `優化得更小`_ ，ISO 僅為 28MB。

.. _GNOME Shell 內存洩露: https://www.oschina.net/news/94451/gnome-shell-has-a-memory-leak
.. _並解決問題: https://www.oschina.net/news/95458/gnome-shell-memory-leak-fix
.. _推出了: https://www.oschina.net/news/94484/ubuntu-18-04-default-wallpaper
.. _優化得更小: https://www.oschina.net/news/94287/ubuntu-has-made-its-minimal-images-even-more-minimal-just-28mb


四月：Ubuntu 18.04 LTS 發佈
===========================

4月份，萬眾期待的 Ubuntu 18.04 LTS 長期支持版 `終於發佈了`_ ，並給那些選擇從 16.04 升級到 Ubuntu 18.04 的用戶帶來了驚喜：新的桌面、新的工作流程，以及全新的一切。

憑藉由 Canonical 提供 5 年的技術支持（上個月 Canonical 創始人兼 CEO Mark Shuttleworth `宣佈`_ Ubuntu 18.04 LTS Linux 支持時間延長到十年），Ubuntu 18.04 LTS 仍然是值得推薦的首選版本。而 Bionic Beaver 的到來也為大家提供了很多幫助，包括 Nautilus 的全新面貌，以及更新的 Ambiance GTK 主題。

此外，在4月份發佈的 Ubuntu 18.04 中還看到了對自動“最小化安裝”選項以及 Canonical 的 Livepatch 服務的集成支持，該服務允許用戶安裝新的 Linux 內核補丁而無需重新啟動系統。

.. _終於發佈了: https://www.oschina.net/news/95551/ubuntu-1804-lts-released
.. _宣佈: https://www.serverwatch.com/server-news/canonical-extends-ubuntu-18.04-lts-linux-support-to-10-years.html


五月：到達和離開(Arrivals & Departures)
=======================================

Mark Shuttleworth 在五月份表示，將會在“下一代” Ubuntu 安裝程序中使用 Electron, Curtin 和 Snappy 等現代技術。雖然新的安裝程序在撰寫本文時尚未實現，但開發團隊保證會在不久的將來推出。

最後，Ubuntu MATE 和 Ubuntu Budgie 表示將不再提供 `32 位 ISO`_ 。他們不是第一個（Ubuntu 在 2017 年放棄了 32 位版本），也不是最後一個：Xubuntu 在後來也宣佈不再支持 32 位，並擁抱唯一的未來， 64位系統。

.. _32 位 ISO: https://www.omgubuntu.co.uk/2018/05/ubuntu-mate-budgie-32-bit-iso-dropped


六月：Minty Fresh
=================

6月份，Linux Mint 19 ‘Tara’ `正式推出`_ ，這是基於 Ubuntu 18.04 LTS 的流行 Linux 發行版的第一個版本。該版本原生提供對 Flatpak 的支持，帶來了新的歡迎應用程序和新的 GTK 主題。

在其他地方，粉絲在軟件包中發現了一些小彩蛋 —— 英特爾的第7代NUC被 `標記為"Ubuntu Certified"`_ 。

.. _正式推出: https://www.oschina.net/news/97556/linux-mint-19-released
.. _標記為"Ubuntu Certified": https://www.omgubuntu.co.uk/2018/06/intels-7th-gen-nucs-are-now-ubuntu-certified


七月：安全升級版本發佈
======================

7月為 Ubuntu 18.04 LTS 發佈了第一個維護版本。新的 ISO 映像包含許多關鍵錯誤修復和安全補丁。這也是 Canonical 向 Ubuntu 16.04 LTS 發送“升級信號”的月份。

同一個月份， `SUSE Linux 以 25 億美元被 EQT 收購`_ 。

.. _SUSE Linux 以 25 億美元被 EQT 收購: https://www.oschina.net/news/97651/micro-focus-sell-suse-to-eqt


八月：使用 Ubuntu 的筆記本電腦大量上市
======================================

戴爾提供便宜的 XPS 13（系統由 Ubuntu 提供支持）大規模開售，這是一款極為纖薄的機型，具有頂級規格製造的屏幕和消費者更易接受的價格。

而 Ubuntu 16.04 的第五個維護版也在這個月 `發佈`_ ，以“安撫”那些不想很快失去 Unity 桌面的用戶。這也是“新面貌”的月份，Ubuntu 18.10 帶來了新的‘Yaru’主題，而郵件客戶端 Thunderbird_ 也有了新的界面、logo 和新功能。

.. _發佈: https://www.oschina.net/news/98616/ubuntu-16-04-5-lts-released
.. _Thunderbird: https://www.omgubuntu.co.uk/2018/08/thunderbird-60-release-features


九月：軟件中心改造
==================

流行的發行版 KDE Neon 升級到了 Ubuntu 18.04 LTS。除了安全性和性能改進之外，升級還使得 Neon 用戶可訪問許多核心應用程序的更新版本。九月還看到了 Ubuntu 18.10 `新透露的壁紙`_ 。

在博客上，我們還看到了 Ubuntu 的主要 `軟件中心改造`_ 計畫，該計畫將應用程序列表與策劃的編輯內容混合在一起，像 Apple App Store 一樣。

.. _新透露的壁紙: https://www.oschina.net/news/100395/ubuntu-18-10-released
.. _軟件中心改造: https://www.omgubuntu.co.uk/2018/09/canonical-wants-to-make-its-software-center-more-like-the-apple-app-store


十月：Ubuntu 18.10 發佈
=======================

Ubuntu 18.10 `如約發佈`_ ，主要是對 GNOME Shell 桌面的迭代改進，以及新的 GTK 主題，對此感到興奮的用戶也許並不是很多。

而今年早些時候才開始收集用戶數據的 Ubuntu，直到10月，Canonical 才終於準備通過推出專用的 `Ubuntu 用戶統計網頁`_ 分享細節。 

對了，這個月還有一件可以載入“史冊”的開源大事 —— `IBM 收購了 Red Hat`_ 。

.. _如約發佈: https://www.oschina.net/news/100966/ubuntu-18-10-released
.. _Ubuntu 用戶統計網頁: https://www.omgubuntu.co.uk/2018/10/ubuntu-user-statistics-revealed
.. _IBM 收購了 Red Hat: https://www.oschina.net/news/101256/ibm-to-acquire-red-hat


十一月：敬業的工程師
====================

Ubuntu `推出了 Ubuntu 19.04 的每日構建版`_ 。而三星則公佈了一個新項目 —— `Linux on DeX`_ ，Linux on DeX 可讓你隨時隨地享受 Linux 環境。

.. _推出了 Ubuntu 19.04 的每日構建版: https://www.oschina.net/news/101781/ubuntu-19-04-daily-can-download-now
.. _Linux on DeX: https://www.linuxondex.com/

.. image:: https://oscimg.oschina.net/oscnet/7a7d8b7161af0a1be436a109a563beeeb81.jpg
   :alt: Ubuntu 2018 回顧：從內存洩露到“感人”的 LTS 版本
   :align: center

通過該應用三星手機可以啟動 Linux 容器，然後再連接顯示器，就會變成 Ubuntu 桌面環境，從而在手機上達到 PC 開發的體驗。不過 Linux on DeX 僅支持一個定製的 Ubuntu 鏡像（Canonical 提供的 Ubuntu 16.04 LTS 版本）。


十二月：FOSS 慶祝活動
=====================

各種基於 Ubuntu 的發行版都發佈了更新。對於 Ubuntu 來說，各種工作仍在進行，計畫改變 Alt + Tab 的行為，將 GNOME 時鐘和跟蹤器添加到默認安裝，並創建一個特殊的設置區域來管理 Canonical Livepatch 等。


在其他風味版本的相關新聞中，Xubuntu 宣佈它從 19.04 起或更高版本將不再提供 32 位版本。

如有遺漏的重大事件，歡迎大家在評論區補充~！

編譯自： omgubuntu_

..
  .. image:: 
   :alt: 
   :align: center


.. highlights::

  | 本站文章除註明轉載外，均為本站原創或編譯。歡迎任何形式的轉載，但請務必註明出處，尊重他人勞動共創開源社區。
  | 轉載請註明：文章轉載自 開源中國社區 [https://www.oschina.net]
  | 本文標題：Ubuntu 2018 回顧：從內存洩露到“感人”的 LTS 版本
  | 本文地址：https://www.oschina.net/news/102975/ubuntu-in-2018-recap

.. _omgubuntu: https://www.omgubuntu.co.uk/2018/12/ubuntu-in-2018-recap

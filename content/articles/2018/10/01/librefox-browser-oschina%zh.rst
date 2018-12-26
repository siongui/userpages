Librefox：具有隱私增強功能的Firefox
###################################

:date: 2018-12-25T16:51+08:00
:tags: 轉錄, 開源中國
:category: 轉錄
:author: 程六金(開源中國)
:summary: Librefox 是一個開源項目，旨在為任何人提供隱私保護和增強安全的 Firefox 瀏覽器代替品。
:og_image: https://www.ghacks.net/wp-content/uploads/2018/12/librefox-firefox-privacy.png

Librefox 是一個開源項目，旨在為任何人提供隱私保護和增強安全的 Firefox 瀏覽器代替品。

Librefox 不是 Firefox 的分支，Librefox 使用核心的 Firefox 功能，並對瀏覽器在隱私保護，安全性方面進行優化增強。該項目使用 Ghacks user.js 和其他功能創建一個瀏覽器。提供更好的隱私保護和開箱即用的安全性。

Librefox 適用於 Windows，Mac 和 Linux 設備。在 Window 上，無需安裝即可運行。Librefox 項目團隊創建了基於 Firefox 穩定版的擴展插件和 Librefox 的 Beta 版。

在使用時，整體界面與 Firefox 並無差異，因為整個瀏覽器還是運行在 Firefox 之上。

.. image:: https://www.ghacks.net/wp-content/uploads/2018/12/librefox-firefox-privacy.png
   :alt: Librefox：具有隱私增強功能的Firefox
   :align: center

Librefox 開發團隊從 Firefox 中刪除了一些組件：從瀏覽器中刪除了『不尊重隱私』的更新程序、崩潰報告組件和其他附加組件。Firefox 默認使用的通信組件也被刪除：

.. code-block:: bash

  The objective is zero unauthorized connection (ping/telemetry/Mozilla/Google...).

Lirefox 沒有附加組件，目前開發團隊已經為 LibreFox 創建了幾個擴展組件，團隊建議安裝。建議安裝的組件已經通過了相關的代碼審查。

Librefox 擴展為 Firefox 提供了一個暗黑主題，HTTP Watcher 和 Relaod 按鈕。推薦的第三方擴展包括：uBlock Origin、Cookies Master、First Party Isolation、User Agent Platform Spoofer 和 Brower Plugs Privacy Firefox。擴展相關鏈接在 `項目網站`_ 上。

初次之外，Librefox 有一個 Extensions Firefox 的實驗性功能，但默認情況下它被禁用。他旨在在全局管理加載項，允許或者禁止相關擴展連接。

想要瞭解更多 Firefox 和 Librefox 之間的差異可以從檢查 mozilla.cfg 和 policies.json開始。對於隱私安全有需求的Firefox用戶來說，Librefox可能值得一試。

項目地址： `Librefox: Librefox, patching Firefox for an enforced privacy and security`_

..
  .. image:: 
   :alt: 
   :align: center

.. highlights::

  | 本站文章除註明轉載外，均為本站原創或編譯。歡迎任何形式的轉載，但請務必註明出處，尊重他人勞動共創開源社區。
  | 轉載請註明：文章轉載自 開源中國社區 [https://www.oschina.net]
  | 本文標題：Librefox：具有隱私增強功能的Firefox
  | 本文地址：https://www.oschina.net/news/103014/librefox-browser

.. _項目網站: https://github.com/intika/Librefox/#librefox-addons-
.. _Librefox\: Librefox, patching Firefox for an enforced privacy and security: https://github.com/intika/Librefox

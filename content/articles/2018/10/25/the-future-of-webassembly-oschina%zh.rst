WebAssembly 的未來：將逐漸解鎖整個“技能樹”
##########################################

:date: 2018-10-25T00:00+08:00
:tags: Go語言, 轉錄, 開源中國, WebAssembly
:category: Go語言
:author: 王練(開源中國)
:summary: WebAssembly 在2017年受到主流瀏覽器的支持，並發布了 MVP 版本（Minimum Viable Product，最簡可行產品）。雖然 WebAssembly 社區組致力於使 WebAssembly 向後相容，讓現在開發的 WebAssembly 應用程式，仍可以在未來的瀏覽器中運作，但這並不代表 WebAssembly 功能已經完整，MVP 絕非最終版本。相反，WebAssembly 還將增加許多新功能，而這些新功能將從根本上顛覆 WebAssembly 所能實現的工作。
:og_image: https://oscimg.oschina.net/oscnet/42ccc510cfd9d1d2670037cba32f3007d47.jpg


WebAssembly 在2017年受到主流瀏覽器的支持，並發布了 MVP 版本（Minimum Viable Product，最簡可行產品）。雖然 WebAssembly 社區組致力於使 WebAssembly 向後相容，讓現在開發的 WebAssembly 應用程式，仍可以在未來的瀏覽器中運作，但這並不代表 WebAssembly 功能已經完整，MVP 絕非最終版本。相反，WebAssembly 還將增加許多新功能，而這些新功能將從根本上顛覆 WebAssembly 所能實現的工作。

為消除人們對 WebAssembly 的誤解，WebAssembly 社區組以 RPG 遊戲中人物養成的“技能樹”形式，對 `WebAssembly 的未來發展路徑`_ 做了非常詳細的解釋。他們表示目前已經完全掌握這些技能中的前幾項，後續需要慢慢解鎖整個技能樹。

.. image:: https://oscimg.oschina.net/oscnet/42ccc510cfd9d1d2670037cba32f3007d47.jpg
   :alt: WebAssembly 的未來：將逐漸解鎖整個“技能樹”
   :align: center

WebAssembly 在 MVP 版本階段滿足 4 個基本技能要求：編譯、快速執行、壓縮和線性內存分配。使用 WebAssembly 的人知道他們不想只支持 C 和 C ++，而是希望能夠將許多不同的語言編譯為 WebAssembly 。經 WebAssembly 編譯器編譯的應用需要能夠快速執行，滿足需求。而為了加速載入速度，WebAssembly 還需具備壓縮能力，減少使用者的等待時間。另外，WebAssembly 需要有別於 JavaScript 使用內存方式，能夠直接管理使用的內存，在加上安全因素的考量，WebAssembly 採用線性內存模式。

.. image:: https://oscimg.oschina.net/oscnet/48a0162ed0e0f848088cecfa4c9c0d6b5b4.jpg
   :alt: WebAssembly 的未來：將逐漸解鎖整個“技能樹”
   :align: center

社區組表示，WebAssembly 的下一個目標是平滑運行那些更重的應用程序，比如 Photoshop、Gmail 等。為確保此類應用能在瀏覽器中運行良好，他們需要解鎖新一批的“技能”，包括支持多線程、SIMD（單指令流多數據流）、64位尋址、流式編譯（在下載的同時編譯 WebAssembly 文件）、分層編譯器、隱式 HTTP 緩存以及一些其他改進。

.. image:: https://oscimg.oschina.net/oscnet/586022e01ab45f31b0d35d223ea932a8827.jpg
   :alt: WebAssembly 的未來：將逐漸解鎖整個“技能樹”
   :align: center

一旦以上功能全部就位，WebAssembly 又將進入下一個階段 —— 與 JavaScript 互操作，包括 JS 和 WebAssembly 之間的快速調用、簡便的數據交換、ES模塊集成、工具鏈集成和向後兼容性。

.. image:: https://oscimg.oschina.net/oscnet/5d163b3cfd5ecd82788682365c270636ebe.jpg
   :alt: WebAssembly 的未來：將逐漸解鎖整個“技能樹”
   :align: center

此外，他們還想在 WebAssembly 中重寫 JavaScript 框架的大部分內容，並使靜態類型的 compile-to-js 語言可編譯為 WebAssembly 。想要實現這兩個目標，WebAssembly 還需要支持高階的語言功能，包括垃圾回收、異常處理、調試以及尾調用（Tail calls）。

.. image:: https://oscimg.oschina.net/oscnet/c84da0e9a6a4992b6a2219bc3b6f65dc950.jpg
   :alt: WebAssembly 的未來：將逐漸解鎖整個“技能樹”
   :align: center

更多功能和具體進度，可查閱 `博客原文`_ 。

|
| 本站文章除註明轉載外，均為本站原創或編譯。歡迎任何形式的轉載，但請務必註明出處，尊重他人勞動共創開源社區。
| 轉載請註明：文章轉載自 開源中國社區 [http://www.oschina.net]
| 本文標題：WebAssembly 的未來：將逐漸解鎖整個“技能樹”
| 本文地址：https://www.oschina.net/news/101159/the-future-of-webassembly

.. _WebAssembly 的未來發展路徑: https://hacks.mozilla.org/2018/10/webassemblys-post-mvp-future/
.. _博客原文: https://hacks.mozilla.org/2018/10/webassemblys-post-mvp-future/

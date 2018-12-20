Google 被指採用不正當手段搞垮微軟 Edge
######################################

:date: 2018-12-19T08:36+08:00
:tags: 轉錄, 開源中國
:category: 轉錄
:author: h4cd(開源中國)
:summary: 微軟 Edge 團隊工程師指責 Google 採用骯髒手法使得 Edge 最終無奈潦草收場。
:og_image: https://oscimg.oschina.net/oscnet/9d3d169a063f660a777c9c40e8f1fbdb305.jpg

微軟 Edge 團隊工程師指責 Google 採用骯髒手法使得 Edge 最終無奈潦草收場。

.. image:: https://oscimg.oschina.net/oscnet/9d3d169a063f660a777c9c40e8f1fbdb305.jpg
   :alt: Google 被指採用不正當手段搞垮微軟 Edge
   :align: center

前幾天微軟通告未來將採用 Chromium 內核開發桌面版 Edge 瀏覽器。雖然官方說得好聽，解釋採用 Chromium 是為了讓用戶有更好的 Web 兼容性，也為 Web 開發者減少碎片化帶來的複雜性，但是外界其實將這視為微軟在幾年的苦心經營後，認識到 Edge 沒法與 Chrome 抗衡，於是改變策略，使用其近幾年在開源領域的成功經驗——用人家的開源技術做一個同類產品再去幹它，這樣的成功已經在其 TypeScript 與 VS Code 等項目上屢試不爽。

然而，近日用戶 JoshuJB `在 Haker News 上發帖道明了“真相”`_ ，JoshuJB 自稱是 Edge 團隊工程師，他表示之所以 Edge 落得如此下場，其中一大原因是谷歌採用了不正常的競爭手段。

比如，他們在 YouTube 視頻中添加了一個隱藏的空 div，這使得微軟無法啟用 fast-path 硬件加速。“在此之前，我們相當先進的視頻加速使我們在電池播放時間方面遠遠領先於 Chrome，但就在他們開始在 YouTube 上做手腳時，他們就開始宣傳 Chrome 觀看視頻時電池續航能力比 Edge 強”，JoshuJB 補充到：“當我們反饋了該問題之後，YouTube 拒絕了刪除隱藏的空 div 的請求，而且沒有對此進一步做出解釋”。

帖子下有人將該情況與 AMD 和 Intel 反壟斷案聯繫起來，當年 AMD 聲稱 Intel 的 C 編譯器故意削弱 AMD 處理器的性能，以此來幫助 Intel 在處理器競爭中取勝。

這樣的手段是真的髒，不過真相如何還沒有定論。目前谷歌方面還沒有做出回應，估計該事件將繼續發酵一陣子。

.. highlights::

  | 本站文章除註明轉載外，均為本站原創或編譯。歡迎任何形式的轉載，但請務必註明出處，尊重他人勞動共創開源社區。
  | 轉載請註明：文章轉載自 開源中國社區 [https://www.oschina.net]
  | 本文標題：Google 被指採用不正當手段搞垮微軟 Edge
  | 本文地址：https://www.oschina.net/news/102831/engineer-accuses-google-of-sabotaging-ms-edge

----

- `Solidot | 前微软实习生指责 Google 破坏微软的浏览器 <https://www.solidot.org/story?sid=58993>`_
- `Solidot | Google 否认修改 YouTube 代码去破坏 Microsoft Edge  <https://www.solidot.org/story?sid=59016>`_

.. _在 Haker News 上發帖道明了“真相”: https://news.ycombinator.com/item?id=18697824

CSS Grid佈局這樣玩
##################

:date: 2017-04-23
:author: 大漠(譯者)
:tags: CSS, Web開發, 響應式網頁設計(Responsive Web Design), 轉錄
:category: CSS
:summary: 自從去年下半年開始，CSS Grid佈局的相關教程在互聯網上就鋪天蓋地，可謂是聲勢浩大。就針對於Web佈局而言，個人認為Grid佈局將是Web佈局的神器，它改變了以往任何一種佈局方式或者方法。不管以前的採用什麼佈局方法都可以說是一維的佈局方式，而Grid最大的特色，採用了二維佈局。@Rachel Andrew也一直致力於完善Grid的規範。
:og_image: https://cdn-images-1.medium.com/max/1200/1*8xM-wxiBuNnhYIG6vFem9Q.png

自從去年下半年開始，CSS Grid佈局的相關教程在互聯網上就鋪天蓋地，可謂是聲勢浩大。就針對於Web佈局而言，個人認為Grid佈局將是Web佈局的神器，它改變了以往任何一種佈局方式或者方法。不管以前的採用什麼佈局方法都可以說是一維的佈局方式，而Grid最大的特色，採用了二維佈局。@Rachel Andrew也一直致力於完善Grid的規範。

就我個人而言，我也一直在不斷的關注這個佈局利器的相關更新，自從最初規範的出來，到目前規範的完善。在站上也不斷的在更新 `CSS Grid佈局`_ 的使用。雖然這方向的教程已經很多了，但各有千秋，我追求以最簡單，最直接的方式來闡述它的使用方式方法。讓初學者能盡快的掌握其使用規則。

前段時間 `@Mirza Joldic`_ 在Medium上發佈了 `一篇文章`_ ，通過幾個Gif動態非常形象的闡述了CSS Grid的幾個核心概念以及使用方法，今天我就借花獻佛，用這幾張圖讓初學者快速掌握CSS Grid的核心概念和使用技巧。


Web佈局的歷史演變
+++++++++++++++++

自從Web出來至今，Web的佈局也經過了幾個演變，下圖可以一目了然：

.. image:: https://cdn-images-1.medium.com/max/800/1*phV0oLsKV_qVjFVv5lY1vw.png
   :alt: CSS Grid佈局這樣玩
   :align: center

有關於Web佈局的演變史，去年也整理過一篇相關的文章簡單的闡述了這方面的故事，如果你感興趣的話， `可以點擊這裡`_ 進行了解。在Web的學習過程中， `學習Web佈局`_ 是一個不可避免的過程，而隨著前端技術的日新月異的變化，佈局方式也在不斷的更新，早在2013年@Peter Gasston就對 `CSS佈局的未來趨勢`_ 就做過預判斷，文章中就提供了CSS Grid的佈局。如果今天來看，這種趨勢的預判是正確的，特別是今年3月份之後，各大主流瀏覽器都發佈了對CSS Grid的支持。既然如此，學習CSS Grid相關的知識就很有必要。

既然掌握CSS Grid很有必要，那用什麼樣的方式能最快的掌握CSS Grid相關的知識呢？這很重要。 特別是@Mirza Joldic在Medium上發佈的文章，裡面的動圖讓我耳目一新，通過簡單的幾張圖，就把CSS Grid的幾個核心介紹的非常清楚，我覺得很有必要拿出來與大家分享。

在繼續下面的內容之前，再次感謝@Mirza Joldic的付出。那咱們就不說廢話了，開始今天的學習之旅。

----

如需轉載，煩請註明出處：http://www.w3cplus.com/css3/playing-with-css-grid-layout.html

轉載：

- `Playing with CSS Grid Layout – Mirza Joldic – Medium <https://medium.com/@purplecones/playing-with-css-grid-layout-a75836098370>`_
- `CSS Grid布局这样玩_CSS3 Grid Layout, Grid, Layout, CSS3 教程_w3cplus <https://www.w3cplus.com/css3/playing-with-css-grid-layout.html>`_
- `CSS Grid布局这样玩 - WEB前端 - 伯乐在线 <http://web.jobbole.com/91146/>`_

.. _CSS Grid佈局: https://www.w3cplus.com/blog/tags/356.html
.. _@Mirza Joldic: https://medium.com/@purplecones
.. _一篇文章: https://medium.com/@purplecones/playing-with-css-grid-layout-a75836098370
.. _可以點擊這裡: https://www.w3cplus.com/css/css-layout-model.html
.. _學習Web佈局: https://www.w3cplus.com/css/learn-css-layout.html
.. _CSS佈局的未來趨勢: https://www.w3cplus.com/css3/future-css-layouts.html

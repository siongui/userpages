Sticky Footer，完美的絕對底部
#############################

:date: 2017-04-14T08:49:02.866Z
:author: NoNo Lee
:tags: CSS, Web開發, 響應式網頁設計(Responsive Web Design), 轉錄
:category: CSS
:summary: Sticky Footer，完美的絕對底部
:og_image: https://misc.aotu.io/liqinuo/sticky_cover.png

.. image:: https://misc.aotu.io/liqinuo/sticky_cover.png
   :alt: sticky footer

| 轉載自 `凹凸实验室`_: `Sticky Footer，完美的绝对底部`_
| by NONO_ on 2017-04-13


寫在前面
++++++++

做過網頁開發的同學想必都遇到過這樣尷尬的排版問題：

在主體內容不足夠多或者未完全加載出來之前，就會導致出現（圖一）的這種情況，原因是因為沒有足夠的垂直空間使得頁腳推到瀏覽器窗口最底部。但是，我們期望的效果是頁腳應該一直處於頁面最底部（如圖二）：

.. image:: https://misc.aotu.io/liqinuo/sticky_01.png
   :alt: sticky footer

筆者最近在項目中也遇到過這樣的場景，在尋找最佳解決方案的過程中，瞭解到了 「Sticky Footer」 這個名詞。

本文將帶大家重新認識這個常見的網頁效果，以及一些可行的實現方案。


什麼是 「Sticky Footer」
++++++++++++++++++++++++

所謂 「Sticky Footer」，並不是什麼新的前端概念和技術，它指的就是一種網頁效果：

如果頁面內容不足夠長時，頁腳固定在瀏覽器窗口的底部；如果內容足夠長時，頁腳固定在頁面的最底部。

總而言之，就是頁腳一直處於最底，效果大致如圖所示：

.. image:: https://misc.aotu.io/liqinuo/sticky_02.png
   :alt: sticky footer

當然，實現這種效果的方法有很多種，其中有通過腳本計算的，有通過 CSS 處理的，腳本計算的方案我們不在本文探討。

下面我們看看有哪些通過 CSS 可以實現且適用於移動端開發的方案，並分析其中的利弊。


如何實現
++++++++

假設我們頁面的 HTML 結構是這樣：

.. code-block:: html

  <div class="wrapper">
      <div class="content"><!-- 頁面主體內容區域 --></div>
      <div class="footer"><!-- 需要做到 Sticky Footer 效果的頁腳 --></div>
  </div>


實現方案一：absolute
++++++++++++++++++++

通過絕對定位處理應該是常見的方案，只要使得頁腳一直定位在主容器預留佔位位置。

.. code-block:: css

  html, body {
      height: 100%;
  }
  .wrapper {
      position: relative;
      min-height: 100%;
      padding-bottom: 50px;
      box-sizing: border-box;
  }
  .footer {
      position: absolute;
      bottom: 0;
      height: 50px;
  }

這個方案需指定 html、body 100% 的高度，且 content 的 padding-bottom 需要與 footer 的 height 一致。


實現方案二：calc
++++++++++++++++

通過計算函數 calc 計算（視窗高度 - 頁腳高度）賦予內容區最小高度，不需要任何額外樣式處理，代碼量最少、最簡單。

.. code-block:: css

  .content {
      min-height: calc(100vh - 50px);
  }
  .footer {
      height: 50px;
  }

如果不需考慮 calc() 以及 vh 單位的兼容情況，這是個很理想的實現方案。

同樣的問題是 footer 的高度值需要與 content 其中的計算值一致。


實現方案三：table
+++++++++++++++++

通過 table 屬性使得頁面以表格的形態呈現。

.. code-block:: css

  html, body {
      height: 100%;
  }
  .wrapper {
      display: table;
      width: 100%;
      min-height: 100%;
  }
  .content {
      display: table-row;
      height: 100%;
  }

需要注意的是，使用 table 方案存在一個比較常見的樣式限制，通常 margin、padding、border 等屬性會不符合預期。

筆者不建議使用這個方案。當然，問題也是可以解決的：別把其他樣式寫在 table 上。


實現方案四：Flexbox
+++++++++++++++++++

Flexbox 是非常適合實現這種效果的，使用 Flexbox 實現不僅不需要任何額外的元素，而且允許頁腳的高度是可變的。

雖然大多數 Flexbox 佈局常用於水平方向佈局，但別忘了實際上它也可用於垂直佈局，所以你需要做的是將垂直部分包裝在一個 Flex 容器中，並選擇要擴展的部分，他們將自動佔用其容器中的所有可用空間。

.. code-block:: css

  html {
      height: 100%;
  }
  body {
      min-height: 100%;
      display: flex;
      flex-direction: column;
  }
  .content {
      flex: 1;
  }

需要注意的是想要兼容各種系統設備，需要兼顧 flex 的兼容寫法。


寫在最後
++++++++

以上幾種實現方案，筆者都在項目中嘗試過，每個實現的方法其實大同小異，同時也都有自己的利弊。

其中有的方案存在限制性問題，需要固定頁腳高度；其中有的方案需要添加額外的元素或者需要 Hack 手段。同學們可以根據頁面具體需求，選擇最適合的方案。

當然，技術是不斷更新的，也許還有很多不同的、更好的方案。但相信大家最終目都是一樣的，為了更好的用戶體驗！

| 參考資料：
| https://css-tricks.com/couple-takes-sticky-footer/
| http://www.w3cplus.com/css3/css-secrets/sticky-footers.html

感謝您的閱讀，本文由 凹凸實驗室 版權所有。如若轉載，請註明出處：凹凸實驗室（https://aotu.io/notes/2017/04/13/Sticky-footer/）

.. _NONO: https://github.com/liqinuo
.. _凹凸实验室: https://aotu.io/
.. _Sticky Footer，完美的绝对底部: https://aotu.io/notes/2017/04/13/Sticky-footer/

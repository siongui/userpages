CSS Grid VS Flexbox：實例對比
#############################

:date: 2017-04-17T19:21+08:00
:author: IT程序狮(譯者)
:tags: CSS, Web開發, 響應式網頁設計(Responsive Web Design), 轉錄
:category: CSS
:summary: CSS Grid VS Flexbox：實例對比
:og_image: http://upload-images.jianshu.io/upload_images/675733-cad4d95cbb4c9fc0.jpg?imageMogr2/auto-orient/strip%7CimageView2/2

.. image:: http://upload-images.jianshu.io/upload_images/675733-cad4d95cbb4c9fc0.jpg?imageMogr2/auto-orient/strip%7CimageView2/2
   :alt: CSS Grid VS Flexbox：實例對比
   :align: center

- DEMO 地址：【 `傳送門 <http://demo.tutorialzine.com/2017/03/css-grid-vs-flexbox/>`__ 】
- 示例下載地址： 【 `傳送門 <http://demo.tutorialzine.com/2017/03/css-grid-vs-flexbox/css-grid-vs-flexbox.zip>`__ 】

----

不久以前，所有 HTML 頁面的佈局還都是通過 tables、floats 以及其他的 CSS 屬性來完成的。面對複雜頁面的佈局，卻沒有很好的辦法。

然而 Flexbox_ 的出現，便輕鬆的解決了複雜的 Web 佈局。它是一種專注於創建穩定的響應式頁面的佈局模式，並可以輕鬆地正確對齊元素及其內容。如今已是大多數 Web 開發人員首選的 CSS 佈局方式。

現在，又出現了一個構建 HTML 最佳佈局體系的新競爭者。（冠軍頭銜正在爭奪中…）它就是強大的 `CSS Grid`_ 佈局。直到本月月底，它也將在 `Firefox 52`_ 和 `Chrome 57`_ 上得到原生支持，相信不久也會得到其他瀏覽器兼容性的支持。


基本佈局測試
++++++++++++

要瞭解這兩個體系構建佈局的方式，我們將通過相同的 HTML 頁面，利用不同的佈局方式 （即 Flexbox 與 CSS Grid）為大家區分。

同時，你也可以通過文章頂部附近的下載按鈕，下載演示項目進行對比，或者通過在線演示來察看它們：

.. image:: http://upload-images.jianshu.io/upload_images/675733-1a128310e884a63e.png?imageMogr2/auto-orient/strip%7CimageView2/2
   :alt: CSS Grid VS Flexbox：實例對比
   :align: center

該頁面的設計相對比較簡單 – 它是由一個居中的容器組成，在其內部則包含了標頭、主要內容部分、側邊欄和頁腳。接下來，我們要完成同時保持 CSS 和 HTML 儘可能整潔的挑戰事項：

1. 在佈局中將四個主要的部分進行定位。
2. 將頁面變為響應式頁面；
3. 對齊標頭：導航朝左對齊，按鈕向右對齊。

如你所見，為了便於比較，我們將所有事項從簡處理。那麼，讓我們從第一個挑戰事項開始吧！


挑戰 1：定位頁面部分
++++++++++++++++++++

Flexbox 解決方案
================

我們將從 Flexbox 解決方案開始。我們將為容器添加 ``display: flex`` 來指定為 Flex 佈局，並指定子元素的垂直方向。

.. code-block:: css

  .container {
      display: flex;
      flex-direction: column;
  }

現在我們需要使主要內容部分和側邊欄彼此相鄰。由於 Flex 容器通常是單向的，所以我們需要添加一個包裝器元素。

.. code-block:: html

  <header></header>
  <div class="main-and-sidebar-wrapper">
      <section class="main"></section>
      <aside class="sidebar"></aside>
  </div>
  <footer></footer>

然後，我們給包裝器在反向添加 ``display: flex`` 和 ``flex-direction`` 屬性。

.. code-block:: css

  .main-and-sidebar-wrapper {
      display: flex;
      flex-direction: row;
  }

最後一步，我們將設置主要內容部分與側邊欄的大小。通過 Flex 實現後，主要內容部分會比側邊欄大三倍。

.. code-block:: css

  .main {
      flex: 3;
      margin-right: 60px;
  }
  .sidebar {
     flex: 1;
  }

如你所見，Flex 將其很好的實現了出來，但是仍需要相當多的 CSS 屬性，並借助了額外的 HTML 元素。那麼，讓我們看看 CSS Grid 如何實現的。


CSS Grid 解決方案
=================

針對本項目，有幾種不同的 CSS Grid 解決方法，但是我們將使用 `網格模板區域`_ 語法來實現，因為它似乎最適合我們要完成的工作。

首先，我們將定義四個網格區域，所有的頁面各一個：

.. code-block:: html

  <header></header>
  <!-- Notice there isn't a wrapper this time -->
  <section class="main"></section>
  <aside class="sidebar"></aside>
  <footer></footer>

.. code-block:: css

  header {
      grid-area: header;
  }
  .main {
      grid-area: main;
  }
  .sidebar {
      grid-area: sidebar;
  }
  footer {
      grid-area: footer;
  }

然後，我們會設置網格並分配每個區域的位置。初次接觸 Grid 佈局的朋友，可能感覺以下的代碼會有些複雜，但當你瞭解了網格體系，就很容易掌握了。

.. code-block:: css

  .container {
      display: grid;

      /*     Define the size and number of columns in our grid.
      The fr unit works similar to flex:
      fr columns will share the free space in the row in proportion to their value.
      We will have 2 columns - the first will be 3x the size of the second.  */
      grid-template-columns: 3fr 1fr;

      /*     Assign the grid areas we did earlier to specific places on the grid.
          First row is all header.
          Second row is shared between main and sidebar.
          Last row is all footer.  */
      grid-template-areas:
          "header header"
          "main sidebar"
          "footer footer";

      /*  The gutters between each grid cell will be 60 pixels. */
      grid-gap: 60px;
  }

就是這樣！ 我們現在將遵循上述結構進行佈局，甚至不需要我們處理任何的 margins 或 paddings 。


挑戰 2：將頁面變為響應式頁面
++++++++++++++++++++++++++++

Flexbox 解決方案
================

這一步的執行與上一步密切相關。對於 Flexbox 解決方案，我們將更改包裝器的 ``flex-direction`` 屬性，並調整一些 margins。

.. code-block:: css

  @media (max-width: 600px) {
      .main-and-sidebar-wrapper {
          flex-direction: column;
      }

      .main {
          margin-right: 0;
          margin-bottom: 60px;
      }
  }

由於網頁比較簡單，所以我們在媒體查詢上不需要太多的重寫。但是，如果遇見更為複雜的佈局，那麼將會重新的定義相當多的內容。


CSS Grid 解決方案
=================

由於我們已經定義了網格區域，所以我們只需要在媒體查詢中重新排序它們。 我們可以使用相同的列設置。

.. code-block:: css

  @media (max-width: 600px) {
      .container {
      /*  Realign the grid areas for a mobile layout. */
          grid-template-areas:
              "header header"
              "main main"
              "sidebar sidebar"
              "footer footer";
      }
  }

或者，我們可以從頭開始重新定義整個佈局。

.. code-block:: css

  @media (max-width: 600px) {
      .container {
          /*  Redefine the grid into a single column layout. */
          grid-template-columns: 1fr;
          grid-template-areas:
              "header"
              "main"
              "sidebar"
              "footer";
      }
  }


挑戰 3：對齊標頭組件
++++++++++++++++++++

Flexbox 解決方案
================

我們的標頭包含了導航和一個按鈕的相關鏈接。我們希望導航朝左對齊，按鈕向右對齊。而導航中的鏈接務必正確對齊，且彼此相鄰。

.. code-block:: html

  <header>
      <nav>
          <li><a href="#"><h1>Logo</h1></a></li>
          <li><a href="#">Link</a></li>
          <li><a href="#">Link</a></li>
      </nav>
      <button>Button</button>
  </header>

我們曾在一篇較早的文章中使用 Flexbox 做了類似的佈局： `響應式標頭最簡單的製作方法`_ 。這個技術很簡單：

.. code-block:: css

  header {
      display: flex;
      justify-content: space-between;
  }

現在導航列表和按鈕已正確對齊。下來我們將使 ``<nav>`` 內的 items 進行水平移動。這裡最簡單的方法就是使用 ``display：inline-block`` 屬性，但目前我們需要使用一個 Flexbox 解決方案：

.. code-block:: css

  header nav {
      display: flex;
      align-items: baseline;
  }

僅兩行代碼就搞定了！ 還不錯吧。接下來讓我們看看如何使用 CSS Grid 解決它。


CSS Grid 解決方案
=================

為了拆分導航和按鈕，我們要為標頭定義 ``display: grid`` 屬性，並設置一個 2 列的網格。同時，我們還需要兩行額外的 CSS 代碼，將它們定位在相應的邊界上。

.. code-block:: css

  header{
      display: grid;
      grid-template-columns: 1fr 1fr;
  }
  header nav {
      justify-self: start;
  }
  header button {
      justify-self: end;
  }

至於導航中的內鏈 - 這是我們使用 CSS grid 最好的佈局展示：

.. image:: http://upload-images.jianshu.io/upload_images/675733-bd93d208a446870e.png?imageMogr2/auto-orient/strip%7CimageView2/2
   :alt: CSS Grid VS Flexbox：實例對比
   :align: center

雖然鏈接為內鏈形式，但它們不能正確的對齊。由於 CSS grid 不具備基線選項（不像 Flexbox 具備的 ``align-items`` 屬性），所以我們只能再定義一個子網格。

.. code-block:: css

  header nav {
      display: grid;
      grid-template-columns: auto 1fr 1fr;
      align-items: end;
  }

CSS grid 在此步驟中，存在一些明顯的佈局上的缺陷。但你也不必過於驚訝。因為它的目標是對齊容器，而不是內部的內容。所以，用它來處理收尾工作，或許不是很好的選擇哦。


結論
++++

如果你已經瀏覽完整篇文章，那麼結論不會讓你感到意外。事實上，並不存在最好的佈局方式。Flexbox 和 CSS grid 是兩種不同的佈局形式，我們應該根據具體的場景將它們搭配使用，而不是相互替代。

對於那些跳過文章只想看結論的朋友（不用擔心，我們也這樣做），這裡是通過實例比較後的總結：

1. CSS grids 適用於佈局大畫面。它們使頁面的佈局變得非常容易，甚至可以處理一些不規則和非對稱的設計。
2. Flexbox 非常適合對齊元素內的內容。你可以使用 Flex 來定位設計上一些較小的細節。
3. 2D 佈局適合使用 CSS grids（行與列）。
4. Flexbox 適用於單一維度的佈局（行或列）。
5. 共同學習並使用它們。

----

感謝你的閱讀。若你有所收穫，歡迎點讚與分享。

**註：**

1. 本文版權歸原作者所有，僅用於學習與交流。
2. 如需轉載譯文，煩請按下方註明出處信息，謝謝！

     | 英文原文： `CSS Grid VS Flexbox: A Practical Comparison`_
     | 作者：Danny Markov
     | 譯者：IT程序獅
     | 譯文地址：http://www.jianshu.com/p/6262c3e48443

----

轉載： `CSS Grid VS Flexbox：实例对比 - WEB前端 - 伯乐在线 <http://web.jobbole.com/91120/>`_

參考：

`CSS3 Flexbox属性可视化指南 <http://www.css88.com/archives/5744>`_

.. _Flexbox: https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Flexible_Box_Layout/Using_CSS_flexible_boxes
.. _CSS Grid: https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Grid_Layout
.. _Firefox 52: https://hacks.mozilla.org/2017/03/firefox-52-introducing-web-assembly-css-grid-and-the-grid-inspector/
.. _Chrome 57: https://developers.google.com/web/updates/2017/03/nic57
.. _網格模板區域: https://developer.mozilla.org/en-US/docs/Web/CSS/grid-template-areas
.. _響應式標頭最簡單的製作方法: http://tutorialzine.com/2016/02/quick-tip-easiest-way-to-make-responsive-headers/
.. _CSS Grid VS Flexbox\: A Practical Comparison: http://tutorialzine.com/2017/03/css-grid-vs-flexbox/

利用視口單位實現適配佈局
########################

:date: 2017-05-02T13:06:51.654Z
:author: Tingglelaoo
:tags: CSS, Web開發, 響應式網頁設計(Responsive Web Design), 轉錄
:category: CSS
:summary: 近年來，隨著移動端對視口單位的支持越來越成熟、廣泛，使得我們可以嘗試一種新的辦法去真正地適配所有設備尺寸。
:og_image: https://cdn.rawgit.com/o2team/misc/gh-pages/Tingglelaoo/CSS_viewport_units_900x500.jpg

.. image:: https://cdn.rawgit.com/o2team/misc/gh-pages/Tingglelaoo/CSS_viewport_units_900x500.jpg
   :alt: 利用視口單位實現適配佈局
   :align: center

| 轉載自 `凹凸实验室`_: `利用视口单位实现适配布局`_
| by Tingglelaoo_ on 2017-04-28

響應式佈局的實現依靠媒體查詢（ Media Queries ）來實現，選取主流設備寬度尺寸作為斷點針對性寫額外的樣式進行適配，但這樣做會比較麻煩，只能在選取的幾個主流設備尺寸下呈現完美適配。
即使是通過 rem 單位來實現適配，也是需要內嵌一段腳本去動態計算根元素大小。

近年來，隨著移動端對視口單位的支持越來越成熟、廣泛，使得我們可以嘗試一種新的辦法去真正地適配所有設備尺寸。


認識視口單位（ Viewport units )
+++++++++++++++++++++++++++++++

首先，我們要瞭解什麼是視口。

在業界，極為推崇的一種理論是 Peter-Paul Koch (江湖人稱「PPK大神」)提出的關於視口的 `解釋`_ ——在桌面端，視口指的是在桌面端，指的是瀏覽器的可視區域；而在移動端較為複雜，它涉及到三個視口：分別是 Layout Viewport（佈局視口）、 Visual Viewport（視覺視口）、Ideal Viewport。

而視口單位中的「視口」，在桌面端，毫無疑問指的就是瀏覽器的可視區域；但是在移動端，它指的則是三個 Viewport 中的 Layout Viewport 。

.. _解釋: http://weizhifeng.net/viewports.html

.. container:: align-center

   .. image:: https://misc.aotu.io/Tingglelaoo/viewport.jpg
      :alt: 視口單位中的「視口」
      :align: center

   視口單位中的「視口」


根據 `CSS3規範`_ ，視口單位主要包括以下4個：

.. _CSS3規範: https://drafts.csswg.org/css-values-3/#viewport-relative-lengths

- vw : 1vw 等於視口寬度的1%
- vh : 1vh 等於視口高度的1%
- vmin : 選取 vw 和 vh 中最小的那個
- vmax : 選取 vw 和 vh 中最大的那個

視口單位區別於 ``%`` 單位，視口單位是依賴於視口的尺寸，根據視口尺寸的百分比來定義的；而 ``%`` 單位則是依賴於元素的祖先元素。

.. container:: align-center

   .. image:: https://misc.aotu.io/Tingglelaoo/vw_vh.jpg
      :alt: 用視口單位度量，視口寬度為100vw，高度為100vh（左側為豎屏情況，右側為橫屏情況）
      :align: center

   用視口單位度量，視口寬度為100vw，高度為100vh（左側為豎屏情況，右側為橫屏情況）

例如，在桌面端瀏覽器視口尺寸為650px，那麼 1vw = 650 * 1% = 6.5px（這是理論推算的出，如果瀏覽器不支持0.5px，那麼實際渲染結果可能是7px）。


兼容性
++++++

其兼容性如下圖所示，可以知道：在移動端 iOS 8 以上以及 Android 4.4 以上獲得支持，並且在微信 x5 內核中也得到完美的全面支持。

.. container:: align-center

   .. image:: https://misc.aotu.io/Tingglelaoo/caniuse_viewport.png
      :alt: 截圖來自Can I Use
      :align: center

   截圖來自 `Can I Use <http://caniuse.com/#search=vm>`_

.. container:: align-center

   .. image:: https://misc.aotu.io/Tingglelaoo/wechat.jpg
      :alt: 截圖來自X5內核－Can I Use
      :align: center

   截圖來自 `X5內核－Can I Use <http://res.imtt.qq.com/tbs/incoming20160419/home.html>`_


利用視口單位適配頁面
++++++++++++++++++++

對於移動端開發來說，最為重要的一點是如何適配頁面，實現多終端的兼容，不同的適配方式各有千秋，也各有缺點。

就主流的響應式佈局、彈性佈局來說，通過 Media Queries 實現的佈局需要配置多個響應斷點，而且帶來的體驗也對用戶十分的不友好：佈局在響應斷點範圍內的分辨率下維持不變，而在響應斷點切換的瞬間，佈局帶來斷層式的切換變化，如同卡帶的唱機般「咔咔咔」地一下又一下。

而通過採用rem單位的動態計算的彈性佈局，則是需要在頭部內嵌一段腳本來進行監聽分辨率的變化來動態改變根元素字體大小，使得 CSS 與 JS 耦合了在一起。

有沒有辦法能夠解決這樣的問題呢？

答案是肯定的，通過利用視口單位實現適配的頁面，是既能解決響應式斷層問題，又能解決腳本依賴的問題的。


做法一：僅使用vw作為CSS單位
+++++++++++++++++++++++++++

在僅使用 vw 單位作為唯一應用的一種 CSS 單位的這種做法下，我們遵守：

1. 對於設計稿的尺寸轉換為vw單位，我們使用Sass函數編譯

.. code-block:: scss

  //iPhone 6尺寸作為設計稿基準
  $vm_base: 375;
  @function vw($px) {
      @return ($px / 375) * 100vw;
  }

2. 無論是文本還是佈局高寬、間距等都使用 vw 作為 CSS 單位

.. code-block:: scss

  .mod_nav {
      background-color: #fff;
      &_list {
          display: flex;
          padding: vm(15) vm(10) vm(10); // 內間距
          &_item {
              flex: 1;
              text-align: center;
              font-size: vm(10); // 字體大小
              &_logo {
                  display: block;
                  margin: 0 auto;
                  width: vm(40); // 寬度
                  height: vm(40); // 高度
                  img {
                      display: block;
                      margin: 0 auto;
                      max-width: 100%;
                  }
              }
              &_name {
                  margin-top: vm(2);
              }
          }
      }
  }

3. 1物理像素線（也就是普通屏幕下 1px ，高清屏幕下 0.5px 的情況）採用 transform 屬性 scale 實現。

.. code-block:: scss

  .mod_grid {
      position: relative;
      &::after {
          // 實現1物理像素的下邊框線
          content: '';
          position: absolute;
          z-index: 1;
          pointer-events: none;
          background-color: #ddd;
          height: 1px;
          left: 0;
          right: 0;
          top: 0;
          @media only screen and (-webkit-min-device-pixel-ratio: 2) {
              -webkit-transform: scaleY(0.5);
              -webkit-transform-origin: 50% 0%;
          }
      }
      ...
  }

4. 對於需要保持高寬比的圖，應改用 padding-top 實現

.. code-block:: scss

  .mod_banner {
      position: relative;
      padding-top: percentage(100/700); // 使用padding-top
      height: 0;
      overflow: hidden;
      img {
          width: 100%;
          height: auto;
          position: absolute;
          left: 0;
          top: 0;
      }
  }

由此，我們能夠實現一個常見佈局的頁面效果如下：

.. container:: align-center

   .. image:: https://misc.aotu.io/Tingglelaoo/layout.jpg
      :alt: 體驗地址點擊此處
      :align: center

   體驗地址 `點擊此處 <https://jdc.jd.com/demo/ting/vw_layout.html>`_


做法二：搭配vw和rem，佈局更優化
+++++++++++++++++++++++++++++++

這樣的頁面雖然看起來適配得很好，但是你會發現由於它是利用視口單位實現的佈局，依賴於視口大小而自動縮放，無論視口過大還是過小，它也隨著視口過大或者過小，失去了最大最小寬度的限制。

當然，你可以不在乎這樣微小的不友好用戶體驗，但我們還是嘗試下追求修復這樣的小瑕疵吧。

於是，聯想到不如結合rem單位來實現佈局？rem 彈性佈局的核心在於動態改變根元素大小，那麼我們可以通過：

1. 給根元素大小設置隨著視口變化而變化的 vw 單位，這樣就可以實現動態改變其大小。
2. 限制根元素字體大小的最大最小值，配合 body 加上最大寬度和最小寬度

這樣我們就能夠實現對佈局寬度的最大最小限制。因此，根據以上條件，我們可以得出代碼實現如下：

.. code-block:: scss

  // rem 單位換算：定為 75px 只是方便運算，750px-75px、640-64px、1080px-108px，如此類推
  $vm_fontsize: 75; // iPhone 6尺寸的根元素大小基準值
  @function rem($px) {
       @return ($px / $vm_fontsize ) * 1rem;
  }

  // 根元素大小使用 vw 單位
  $vm_design: 750;
  html {
      font-size: ($vm_fontsize / ($vm_design / 2)) * 100vw;
      // 同時，通過Media Queries 限制根元素最大最小值
      @media screen and (max-width: 320px) {
          font-size: 64px;
      }
      @media screen and (min-width: 540px) {
          font-size: 108px;
      }
  }

  // body 也增加最大最小寬度限制，避免默認100%寬度的 block 元素跟隨 body 而過大過小
  body {
      max-width: 540px;
      min-width: 320px;
  }

這裡就不再給出截圖，但你可以 `點擊此處在線地址 <https://jdc.jd.com/demo/ting/vw_rem_layout.html>`_ 進行體驗。


小結
++++

相對於做法一，個人比較推崇做法二，有以下兩點原因：

第一，做法二相對來說用戶視覺體驗更好，增加了最大最小寬度的限制；

第二，更重要是，如果選擇主流的rem彈性佈局方式作為項目開發的適配頁面方法，那麼做法二更適合於後期項目從 rem 單位過渡到 vw 單位。只需要通過改變根元素大小的計算方式，你就可以不需要其他任何的處理，就無縫過渡到另一種CSS單位，更何況vw單位的使用必然會成為一種更好適配方式，目前它只是礙於兼容性的支持而得不到廣泛的應用。


後語
++++

這是筆者在偶然中閱讀到 `[翻譯]使用VH和VW實現真正的流體排版`_ 這一篇文章得到的感悟與成果，也滿心歡喜地期待這篇文章同樣能夠帶給讀者一些啟發，並提出一些的vw單位使用秘笈來交流交流～:）


參考文檔
++++++++

- `基於視口單位的網頁排版 <https://github.com/dwqs/blog/issues/5>`_
- `(轉）基於視口單位的網頁排版 <http://www.open-open.com/lib/view/open1464136989764.html>`_
- `[翻譯]使用VH和VW實現真正的流體排版`_

感謝您的閱讀，本文由 `凹凸實驗室`_ 版權所有。如若轉載，請註明出處：凹凸實驗室（https://aotu.io/notes/2017/04/28/2017-4-28-CSS-viewport-units/）


`利用视口单位实现适配布局 - WEB前端 - 伯乐在线 <http://web.jobbole.com/91190/>`_

.. _Tingglelaoo: https://github.com/Tingglelaoo
.. _凹凸实验室: https://aotu.io/
.. _凹凸實驗室: https://aotu.io/
.. _利用视口单位实现适配布局: https://aotu.io/notes/2017/04/28/2017-4-28-CSS-viewport-units/
.. _[翻譯]使用VH和VW實現真正的流體排版: http://www.cnblogs.com/wengxuesong/archive/2016/05/16/5497653.html

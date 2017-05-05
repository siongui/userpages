一個例子上手SVG動畫
###################

:date: 2017-05-04T11:40:28.853Z
:author: Simba
:tags: CSS, Web開發, 響應式網頁設計(Responsive Web Design), 轉錄
:category: CSS
:summary: CSS3動畫已足夠強大，不過還是有一些它做不到的地方。配合SVG，讓Web動效有更多的可能性。這次要做的效果是一個loading動畫（如圖）：其中旋轉通過CSS來完成，但是旋轉之後圓弧縮短變成笑臉的嘴巴需要借助SVG來實現。
:og_image: https://o2team.github.io/misc/Simbachen/svg/banner.png

.. image:: https://o2team.github.io/misc/Simbachen/svg/banner.png
   :alt: 一個例子上手SVG動畫
   :align: center

| 轉載自 `凹凸实验室`_: `一个例子上手SVG动画`_
| by Simba_ on 2017-05-04

CSS3動畫已足夠強大，不過還是有一些它做不到的地方。配合SVG，讓Web動效有更多的可能性。這次要做的效果是一個loading動畫（如圖）：其中旋轉通過CSS來完成，但是旋轉之後圓弧縮短變成笑臉的嘴巴需要借助SVG來實現。

.. image:: https://o2team.github.io/misc/Simbachen/svg/loading.gif
   :alt: 一個例子上手SVG動畫


Step1、聲明SVG視口
++++++++++++++++++

.. code-block:: html

  <svg width="100" height=“100”></svg>

指定一個寬高都為100像素的區域，width=”100”和width=”100px”是等價的，當然也可以使用其他的合法單位，例如cm、mm、em等。

閱讀器會設置一個默認的坐標系統，見圖：左上角為原點，其中水平（x）坐標向右遞增，垂直（y）坐標向下遞增。

.. image:: https://o2team.github.io/misc/Simbachen/svg/view.png
   :alt: 一個例子上手SVG動畫

..

  在沒有指定的情況下，所有的的數值默認單位都是像素。


Step2、繪制購物袋
+++++++++++++++++

購物袋由兩個部分組成，先畫下面的主體：

.. code-block:: html

  <path d="M 20 40 L 80 40 L 80 90 A 10 10 90 0 1 70 100 L 30 100 A 10 10 90 0 1 20 90" style="fill: #e9e8ee;"/>

任何形狀都可以使用路徑元素畫出，描述輪廓的數據放在它的d屬性中。

a. 樣式中的fill用來設置填充色。

b. 路徑數據由命令和坐標構成:

.. list-table::
   :header-rows: 1

   * - 指令
     - 說明

   * - M 20 40
     - 表示移動畫筆到(20,40)

   * - L 80 40
     - 表示繪制一條線到(80, 40)

   * - A 10 10 90 0 1 70 100
     - 繪制一個橢圓弧

..

  圓弧命令以字母A開始，後面緊跟著7個參數，這7個參數分別用來表示：

  - 橢圓的x半徑和y半徑
  - 橢圓的x軸旋轉角度
  - 圓弧的角度小於180度，為0；大於或等於180度，則為1
  - 以負角度繪制為0，否則為1
  - 終點的x、y坐標

.. image:: https://o2team.github.io/misc/Simbachen/svg/cart.png
   :alt: 一個例子上手SVG動畫

接下來繪制購物袋上面的部分

.. code-block:: html

  <path d="M 35 40 A 15 15 180 1 1 65 40" style="fill: none; stroke: #e9e8ee; stroke-width: 5;” />

上面的部分是一個半圓弧，我同樣用路徑來畫出，也可以使用基礎形狀來完成。

樣式中的 ``stoke`` 和 ``stroke-width`` 分別用來設置描邊色和描邊的寬度。

.. image:: https://o2team.github.io/misc/Simbachen/svg/cart2.png
   :alt: 一個例子上手SVG動畫


Step3、繪制眼睛
+++++++++++++++

.. code-block:: html

  <circle cx=“40"cy="60" r="2.5" style="fill: #fff;"/>
  <circle cx="60" cy="60" r="2.5" style="fill: #fff;"/>

使用基礎形狀，畫兩個個小圓點。四個屬性分別是位置坐標、半徑和填充顏色。

.. image:: https://o2team.github.io/misc/Simbachen/svg/eye.png
   :alt: 一個例子上手SVG動畫


Step4、繪制嘴巴
+++++++++++++++

.. code-block:: html

  <circle cx="50" cy="70" r="15" style="fill: none; stroke: #fff; stroke-width: 5; stroke-linecap: round;transform: rotate(280deg); transform-origin: 50% 50%; stroke-dashoffset: -23; stroke-dasharray: 42, 95;”>

嘴巴是一段圓弧，我繪制了一個圓，然後描邊了其中的一段，並且做了一個旋轉，來讓它的角度處於正確的位置。

  - ``stroke-linecap`` ：用來定義開放路徑的終結,可選round|butt|square
  - ``stroke-dasharray`` ：用來創建虛線
  - ``stroke-dashoffset`` ：設置虛線位置的起始偏移值，在下一步驟裏，它會和stroke-dasharray一起用來實現動效。

.. image:: https://o2team.github.io/misc/Simbachen/svg/mouth.png
   :alt: 一個例子上手SVG動畫


Step5、給嘴巴部分添加動效
+++++++++++++++++++++++++

.. code-block:: css

  @keyframes mouth {
  0% {
  transform: rotate(-80deg);
  stroke-dasharray: 60, 95;
  stroke-dashoffset: 0;
  }
  40% {
  transform: rotate(280deg);
  stroke-dasharray: 60, 95;
  stroke-dashoffset: 0;
  }
  70%, 100% {
  transform: rotate(280deg);
  stroke-dashoffset: -23;
  stroke-dasharray: 42, 95;
  }
  }

動畫分為兩個部分：

1. 圓弧旋轉
2. 旋轉之後縮短變形

..

  在一個循環裏，最後留有30%的時間保持一個停留。

.. image:: https://o2team.github.io/misc/Simbachen/svg/mouth.gif
   :alt: 一個例子上手SVG動畫


Step6、給眼睛添加動畫
+++++++++++++++++++++

兩只眼睛都是沿著圓弧運動 ，例如左眼，首先用一個路徑來規定它的運動軌跡：

.. code-block:: html

  <path id="eyeright"  d="M 40 60 A 15 15 180 0 1 60 60" style="fill: none; stroke-width: 0;"/>

然後使用animateMotion來設置動畫：

.. code-block:: html

  <circle class="eye" cx="" cy="" r="2.5" style="fill: #fff;">
    <animateMotion
      dur="0.8s"
      repeatCount="indefinite"
      keyPoints="0;0;1;1"
      keyTimes="0;0.3;0.9;1"
      calcMode="linear">
      <mpath xlink:href="#eyeleft"/>
    </animateMotion>
  </circle>

..

  - ``dur`` ：動畫的時間
  - ``repeatCount`` ：重復次數
  - ``keyPoints`` ：運動路徑的關鍵點
  - ``timePoints`` ：時間的關鍵點
  - ``calcMode`` ：控制動畫的運動速率的變化，discrete | linear | paced | spline四個屬性可選
  - ``mpath`` ：指定一個外部定義的路徑

.. image:: https://o2team.github.io/misc/Simbachen/svg/eye.gif
   :alt: 一個例子上手SVG動畫


Step7、將不同部位的動畫組合到一起
+++++++++++++++++++++++++++++++++

- 眼睛的動畫是從嘴巴旋轉完成開始，到嘴巴變形完成結束，因此和嘴巴的動畫一樣，我設置了四個對應的關鍵時間點。
- 為了讓銜接更順暢，眼睛的動畫開始比嘴巴變形開始稍微提前了一點點。

.. image:: https://o2team.github.io/misc/Simbachen/svg/end.gif
   :alt: 一個例子上手SVG動畫

參考：

- `MDN-SVG文檔`_
- 《SVG精髓》- 人民郵電出版社

感謝您的閱讀，本文由 `凹凸實驗室`_ 版權所有。如若轉載，請註明出處：凹凸實驗室（https://aotu.io/notes/2017/05/04/example-for-svg-animation/）


`一个例子上手 SVG 动画 - WEB前端 - 伯乐在线 <http://web.jobbole.com/91221/>`_

.. _Simba: https://github.com/Simbachen
.. _凹凸实验室: https://aotu.io/
.. _凹凸實驗室: https://aotu.io/
.. _一个例子上手SVG动画: https://aotu.io/notes/2017/05/04/example-for-svg-animation/
.. _MDN-SVG文檔: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/animateMotion

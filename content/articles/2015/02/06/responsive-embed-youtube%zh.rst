利用CSS根據螢幕大小自動調整網站內嵌的Youtube影片大小
####################################################

:date: 2015-02-06
:modified: 2017-04-26T23:09+08:00
:tags: 響應式網頁設計(Responsive Web Design), Web開發, CSS
:category: 響應式網頁設計(Responsive Web Design)
:summary: 現今移動裝置盛行，有各種不同螢幕大小的裝置來瀏覽網站，
          若在網站上要內嵌入Youtube的影片，
          要如何讓該影片可以根據瀏覽裝置的螢幕大小自動調整影片嵌入大小呢？
:og_image: http://1.bp.blogspot.com/-mGHl5PJ7UIY/Vn9tWZo7ocI/AAAAAAAAA-c/EUO2i-SQbYU/s640/responsive-youtube-video-embed.png
:adsu: yes


響應式嵌入影片
~~~~~~~~~~~~~~

在自己或是幫別人製作的網站上，常常會需要嵌入Youtube影片，現今移動裝置盛行，
有各種不同螢幕大小的裝置來瀏覽網站，若在網站上要內嵌入Youtube的影片，
該如何讓該影片可以根據瀏覽裝置的螢幕大小自動調整影片嵌入大小呢？

在請教 `Google大神`_ 之後，答案出乎意料的簡單。以下是文章 [1]_
裡頭所提出的解決方案：

首先先把下列CSS程式碼貼到網站的CSS檔裡：

.. code-block:: css

  .video-container {
    position: relative;
    padding-bottom: 56.25%;
    padding-top: 30px;
    height: 0;
    overflow: hidden;
  }

  .video-container iframe,
  .video-container object,
  .video-container embed {
      position: absolute;
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
  }

接著把內嵌影片的iframe程式碼包在 **div** 裡面，並且把此 **div** 的
*class* 設成 ``video-container`` ，以下為範例：

.. code-block:: html

  <div class="video-container">
    <iframe width="560" height="315" src="https://www.youtube.com/embed/2KmHtgtEZ1s" frameborder="0" allowfullscreen></iframe>
  </div>

如此一來就內嵌的影片大小就會自動根據螢幕寬度調整了！

.. adsu:: 2

同場加映
++++++++

若想要更進一步，在螢幕寬度 ``559px`` 以下Youtube影片才自動調整大小，螢幕寬度
``560px`` 以上則不變動，則我們可利用 `CSS media query`_
把上面CSS程式碼改成如下：

.. code-block:: css

  @media (max-width: 559px) {
    .video-container {
      position: relative;
      padding-bottom: 56.25%;
      padding-top: 30px;
      height: 0;
      overflow: hidden;
    }

    .video-container iframe,
    .video-container object,
    .video-container embed {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
    }
  }

這樣一來就可以達到我們想要的效果！

.. adsu:: 3

以下為實際demo，在螢幕寬度 ``560px`` 以上，地圖寬度維持不變，
螢幕寬度若寬度不足 ``560px`` ，則會自動縮小寬度。

.. raw:: html

  <style>
  @media (max-width: 559px) {
    .video-container {
      position: relative;
      padding-bottom: 56.25%;
      padding-top: 30px;
      height: 0;
      overflow: hidden;
    }

    .video-container iframe,
    .video-container object,
    .video-container embed {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
    }
  }
  </style>

  <div class="video-container">
    <iframe width="560" height="315" src="https://www.youtube.com/embed/2KmHtgtEZ1s" frameborder="0" allowfullscreen></iframe>
  </div>

.. adsu:: 4

----

**你可能也會有興趣...**

- `利用CSS根據螢幕大小自動調整網站內嵌的Google地圖(Goole Maps)大小 <{filename}../../../2017/02/23/responsive-embed-google-maps%zh.rst>`_

----

參考：

.. [1] `Responsive Youtube Embed <http://avexdesigns.com/responsive-youtube-embed/>`_

.. _Google大神: https://www.google.com/search?q=responsive+youtube+embed
.. _CSS media query: https://developer.mozilla.org/en-US/docs/Web/Guide/CSS/Media_queries

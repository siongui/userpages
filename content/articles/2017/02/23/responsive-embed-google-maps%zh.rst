利用CSS根據螢幕大小自動調整網站內嵌的Google地圖(Goole Maps)大小
###############################################################

:date: 2017-02-23T03:55+08:00
:tags: 響應式網頁設計(Responsive Web Design), Web開發, CSS
:category: 響應式網頁設計(Responsive Web Design)
:summary: 現今移動裝置盛行，有各種不同螢幕大小的裝置來瀏覽網站，
          若在網站上要內嵌入Google地圖（ `Google Maps`_ ），
          如何讓地圖可以根據瀏覽裝置的螢幕大小自動調整地圖嵌入大小呢？
:adsu: yes


`響應式`_ 嵌入Google地圖（ `Google Maps`_ ）
++++++++++++++++++++++++++++++++++++++++++++

在自己或是幫別人製作的網站上，常常會需要嵌入Google地圖（ `Google Maps`_ ），
現今移動裝置盛行，有各種不同螢幕大小的裝置來瀏覽網站，若在網站上要內嵌入
Google地圖（ `Google Maps`_ ），
該如何讓地圖可以根據瀏覽裝置的螢幕大小自動調整地圖嵌入大小呢？

在請教 `Google大神`_ 之後，答案很簡單。以下是文章 [2]_
裡頭所提出的解決方案：

首先先把下列CSS程式碼貼到網站的CSS檔裡：

.. code-block:: css

  .google-maps {
    position: relative;
    padding-bottom: 75%; /* 此為地圖長寬比 */
    height: 0;
    overflow: hidden;
  }
  .google-maps iframe {
    position: absolute;
    top: 0;
    left: 0;
    width: 100% !important;
    height: 100% !important;
  }

接著把內嵌地圖的iframe程式碼包在 **div** 裡面，並且把此 **div** 的
*class* 設成 ``google-maps`` ，以下為範例：

.. code-block:: html

  <div class="google-maps">
    <iframe src="https://www.google.com/maps/embed?pb=!1m14!1m8!1m3!1d15403.525007372726!2d105.4270337!3d15.1648648!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x0%3A0xfe18c74d9fd44fd5!2sWat+Pa+Phothiyan!5e0!3m2!1sen!2stw!4v1487781323266" width="600" height="450" frameborder="0" style="border:0" allowfullscreen></iframe>
  </div>

如此一來就內嵌的地圖大小就會自動根據螢幕寬度調整了！

.. adsu:: 2


同場加映
++++++++

若想要更進一步，在螢幕寬度 ``599px`` 以下Google地圖才自動調整大小，螢幕寬度
``600px`` 以上則不變動，則我們可利用 `CSS media query`_
把上面CSS程式碼改成如下：

.. code-block:: css

  @media (max-width: 599px) {
    .google-maps {
      position: relative;
      padding-bottom: 75%; /* 此為地圖長寬比 */
      height: 0;
      overflow: hidden;
    }
    .google-maps iframe {
      position: absolute;
      top: 0;
      left: 0;
      width: 100% !important;
      height: 100% !important;
    }
  }

這樣一來就可以達到我們想要的效果！

.. adsu:: 3

以下為實際demo，在螢幕寬度 ``600px`` 以上，地圖寬度維持 ``600px``，
螢幕寬度若寬度不足 ``600px`` ，則會自動縮小寬度。

.. raw:: html

  <style>
  @media (max-width: 599px) {
    .google-maps {
      position: relative;
      padding-bottom: 75%;
      height: 0;
      overflow: hidden;
    }
    .google-maps iframe {
      position: absolute;
      top: 0;
      left: 0;
      width: 100% !important;
      height: 100% !important;
    }
  }
  </style>

  <div class="google-maps">
    <iframe src="https://www.google.com/maps/embed?pb=!1m14!1m8!1m3!1d15403.525007372726!2d105.4270337!3d15.1648648!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x0%3A0xfe18c74d9fd44fd5!2sWat+Pa+Phothiyan!5e0!3m2!1sen!2stw!4v1487781323266" width="600" height="450" frameborder="0" style="border:0" allowfullscreen></iframe>
  </div>

.. adsu:: 4

----

**你可能也會有興趣...**

- `利用CSS根據螢幕大小自動調整網站內嵌的Youtube影片大小 <{filename}../../../2015/02/06/responsive-embed-youtube%zh.rst>`_

----

參考：

.. [1] | `responsive google map embed - Google search <https://www.google.com/search?q=responsive+google+map+embed>`_
       | `responsive google map embed - DuckDuckGo search <https://duckduckgo.com/?q=responsive+google+map+embed>`_
       | `responsive google map embed - Ecosia search <https://www.ecosia.org/search?q=responsive+google+map+embed>`_
       | `responsive google map embed - Bing search <https://www.bing.com/search?q=responsive+google+map+embed>`_
       | `responsive google map embed - Yahoo search <https://search.yahoo.com/search?p=responsive+google+map+embed>`_
       | `responsive google map embed - Baidu search <https://www.baidu.com/s?wd=responsive+google+map+embed>`_
       | `responsive google map embed - Yandex search <https://www.yandex.com/search/?text=responsive+google+map+embed>`_
.. [2] `How to Embed Google Maps in Responsive Websites <https://www.labnol.org/internet/embed-responsive-google-maps/28333/>`_

.. _Google Maps: https://maps.google.com/
.. _響應式: https://zh.wikipedia.org/wiki/%E5%93%8D%E5%BA%94%E5%BC%8F%E7%BD%91%E9%A1%B5%E8%AE%BE%E8%AE%A1
.. _Google大神: https://www.google.com/search?q=responsive+google+map+embed
.. _CSS media query: https://developer.mozilla.org/en-US/docs/Web/Guide/CSS/Media_queries

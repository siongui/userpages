OpenJDK 中髒話太多，開發者看不下去了
####################################

:date: 2018-12-17T08:36+08:00
:tags: 轉錄, 開源中國
:category: 轉錄
:author: h4cd(開源中國)
:summary: 近日有開發者提交了一個 issue JDK-8215217，他指出在 OpenJDK 源碼中存在許多髒話，應該將它們刪除。
:og_image: https://oscimg.oschina.net/oscnet/b03c19eb2674cde04be7f0f946ee2a1fc7e.jpg

近日有開發者提交了一個 issue `JDK-8215217`_ ，他指出在 OpenJDK 源碼中存在許多髒話，應該將它們刪除。

.. image:: https://oscimg.oschina.net/oscnet/b03c19eb2674cde04be7f0f946ee2a1fc7e.jpg
   :alt: OpenJDK 中髒話太多，開發者看不下去了
   :align: center

提交者表示，由於是在專業環境中使用 OpenJDK，因此留著這些髒話是十分不妥的，同時他還上傳了一個刪除髒話的變更集。 

此外他還指出了另外一些存在髒話的地方，但他覺得這些位置的髒話似乎不太可能被刪除：

- 用於測試的大字符串。

- hb-private.hh 的“公共區域訪問保護”類型。

OpenJDK 開發者 Adam Farley 對此作出了回應，表示在與社區討論後他們達成了三個決定：

- "Damn" 和 "Crap" 並不是髒話。

- 四個 f-XXXX 中有三個位於 jszip.js，應該在上游糾正（將跟進）。

- BitArray.java 中的 f-XXXX 和 SoftChannel.java 中的粗俗用語確實是髒話，應該被刪除。

Adam 表示將創建並上載新的 webrev，該問題也已被標記為 RESOLVED。

..
  .. image:: 
   :alt: 
   :align: center

.. highlights::

  | 本站文章除註明轉載外，均為本站原創或編譯。歡迎任何形式的轉載，但請務必註明出處，尊重他人勞動共創開源社區。
  | 轉載請註明：文章轉載自 開源中國社區 [https://www.oschina.net]
  | 本文標題：OpenJDK 中髒話太多，開發者看不下去了
  | 本文地址：https://www.oschina.net/news/102757/jdk-8215217-swear-words

----

`Solidot | OpenJDK 代码被指含有太多的脏话 <https://www.solidot.org/story?sid=58969>`_

.. _JDK-8215217: https://bugs.openjdk.java.net/browse/JDK-8215217

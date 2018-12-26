Ant Design 聖誕“彩蛋”炸雷，開源項目為何失控了？
###############################################

:date: 2018-12-25T14:57+08:00
:tags: 轉錄, 開源中國
:category: 轉錄
:author: 局长(開源中國)
:summary: 喜慶洋洋的聖誕節，技術圈也十分熱鬧。因為今天開源屆被一個前端 UI 框架的“彩蛋”掀翻天了。
:og_image: https://static.oschina.net/uploads/space/2018/1225/133646_txSi_2720166.jpg

喜慶洋洋的聖誕節，技術圈也十分熱鬧。因為今天開源屆被一個前端 UI 框架的“彩蛋”掀翻天了。

事件起因是螞蟻金服的 Ant Design 框架，開發者別出心裁地在框架代碼中埋下了一個會在聖誕節當天觸發的“彩蛋”。

.. image:: https://static.oschina.net/uploads/space/2018/1225/133646_txSi_2720166.jpg
   :alt: Ant Design 聖誕“彩蛋”炸雷，開源項目為何失控了？
   :align: center

如圖中所看到的，本來這應該是一個正常的藍色按鈕，而上面的 **「積雪」效果** 就是 Ant Design 今天的聖誕節彩蛋之一。

.. image:: https://static.oschina.net/uploads/space/2018/1225/134123_j4K1_2720166.jpg
   :alt: Ant Design 聖誕“彩蛋”炸雷，開源項目為何失控了？
   :align: center

|

.. image:: https://static.oschina.net/uploads/space/2018/1225/134124_q2n1_2720166.jpg
   :alt: Ant Design 聖誕“彩蛋”炸雷，開源項目為何失控了？
   :align: center

眾所周知 Ant Deign 有著大量的用戶群體，而這個“彩蛋”呈現出來的最終效果卻是猶如一顆“毒蛋”，說是一場災難也不為過，最後也把眾多的框架使用者給坑了 —— 大量前端開發者紛紛前往 Ant Design 的 GitHub 倉庫留言、吐槽……

.. image:: https://static.oschina.net/uploads/space/2018/1225/134635_GYDt_2720166.png
   :alt: Ant Design 聖誕“彩蛋”炸雷，開源項目為何失控了？
   :align: center

|

.. image:: https://static.oschina.net/uploads/space/2018/1225/134721_UQJR_2720166.jpg
   :alt: Ant Design 聖誕“彩蛋”炸雷，開源項目為何失控了？
   :align: center

|

.. image:: https://static.oschina.net/uploads/space/2018/1225/135616_VOWn_2720166.jpg
   :alt: Ant Design 聖誕“彩蛋”炸雷，開源項目為何失控了？
   :align: center

|

.. image:: https://static.oschina.net/uploads/space/2018/1225/135617_LkAy_2720166.jpg
   :alt: Ant Design 聖誕“彩蛋”炸雷，開源項目為何失控了？
   :align: center

Ant Design 的核心維護人員面對用戶突如其來的質疑和吐槽， `立馬做出了回應`_ ， 並提供了修復的方法：

.. _立馬做出了回應: https://zhuanlan.zhihu.com/p/53214213

.. highlights::

  關於 Ant Design 聖誕彩蛋，起源自 2018 年 9 月 10 日我的一次提交： `add christmas easter egg · ant-design/ant-design@00aebeb`_ ，代碼實現會在 12 月 25 日當天給所有按鈕添加積雪效果，並增加 `Ho Ho Ho!` 的瀏覽器默認提示信息。這完全是我個人的一意孤行且愚蠢的決定，是我的錯誤給大家造成了不良影響，非常抱歉。

  **如何修復這個問題？**

  影響範圍：3.9.3、3.10.0~3.10.9、3.11.0~3.11.5

  我們已經發佈了修訂版本：3.9.4、3.10.10、3.11.6，各位請更新至相應的版本即可。使用了語義化版本的直接重新安裝 node_modules 並重新下載即可。

  **代碼裡還有其他彩蛋麼？**

  沒有。

  **未來還會有類似的問題麼？**

  不會。我們是開源軟件，請像這一次一樣持續監督我們。

.. _add christmas easter egg · ant-design/ant-design@00aebeb: https://github.com/ant-design/ant-design/commit/00aebeb9756afecc884ad48486084836b9a2707a

玉伯也對該事件 `進行了回應`_ ：

.. _進行了回應: https://www.zhihu.com/question/306858501/answer/559312342
 
.. image:: https://oscimg.oschina.net/oscnet/up-cl8tfxcqerjzr9gwfeuleqdpf6tnhgmu.jpg
   :alt: Ant Design 聖誕“彩蛋”炸雷，開源項目為何失控了？
   :align: center

對於 Ant Design 這次的聖誕節彩蛋，你怎麼看待？

都說開源軟件會因為受到監督而安全，但這次事件似乎反應了開源項目的一個典型困境 —— 其支持者聲稱由於開源軟件會因開源而公開透明，並因持續受到監督而安全。但在實際過程中，多數人並不會審查代碼。即使有人發現問題，也因為不受重視而無法解決，最終就有可能導致嚴重的後果。

對此你如何看待？歡迎留言共同討論。

開源中國亦將會持續跟蹤報導該事件。

..
  .. image:: 
   :alt: 
   :align: center

.. highlights::

  | 本站文章除註明轉載外，均為本站原創或編譯。歡迎任何形式的轉載，但請務必註明出處，尊重他人勞動共創開源社區。
  | 轉載請註明：文章轉載自 開源中國社區 [https://www.oschina.net]
  | 本文標題：Ant Design 聖誕“彩蛋”炸雷，開源項目為何失控了？
  | 本文地址：https://www.oschina.net/news/103025/ant-design-festival-code

----

`Solidot | 支付宝开源前端加入圣诞节彩蛋，被紧急撤回 <https://www.solidot.org/story?sid=59073>`_


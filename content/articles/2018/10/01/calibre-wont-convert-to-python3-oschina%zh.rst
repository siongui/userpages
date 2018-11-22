不兼容惹的禍，Calibre 作者拒絕遷移至 Python 3
#############################################

:date: 2018-11-22T10:01+08:00
:tags: 轉錄, 開源中國, Linux, Python
:category: Python
:author: 王練(開源中國)
:summary: 開源電子書管理器 Calibre 的作者 Kavid Goyal 近日公開了自己在 2017 年對用戶要求升級至 Python 3 的回應，以表明 Calibre 不會直接遷移至 Python 3 的態度。
:og_image: https://oscimg.oschina.net/oscnet/88bcb143b0fdfe8b3ede9985aac837a4b02.jpg


開源電子書管理器 Calibre 的作者 Kavid Goyal 近日公開了自己在 2017 年對用戶要求升級至 Python 3 的 `回應`_ ，以表明 Calibre 不會直接遷移至 Python 3 的態度。

2017年8月，有用戶提交反饋稱由於 Python 2 將於2020年停止支持，Calibre 需要升級為 Python 3 。Kavid Goyal 將該問題的狀態更改為“不會修復”，並回應道：“沒有必要，我完全有能力自己維護 Python 2 。這樣做的工作量比考慮遷移整個代碼庫要少得多。”

.. image:: https://oscimg.oschina.net/oscnet/88bcb143b0fdfe8b3ede9985aac837a4b02.jpg
   :alt: 不兼容惹的禍，Calibre 作者拒絕遷移至 Python 3
   :align: center

隨著前幾個月 Guido van Rossum 對 `Python 2.7 將於2020年1月1日終止支持`_ 的證實，Calibre 再次收到了不少要求升級的需求，因此 Kavid Goyal 選擇將之前的 `回應公開`_ ，以作統一回覆。

此外，Kavid Goyal 在 GitHub 上補充道，“我的目標是以與 Python 2 相同的方式在 Python 3 上運行 Calibre 的代碼。計畫在2019年底或最遲在2020年年初，完全兼容 Python 3 。”

Python 3.0 發佈至今已有十年，許多軟件和項目已完成遷移。但由於不向後兼容，也有部分開發者選擇不去跟進。

.. highlights::

  | 本站文章除註明轉載外，均為本站原創或編譯。歡迎任何形式的轉載，但請務必註明出處，尊重他人勞動共創開源社區。
  | 轉載請註明：文章轉載自 開源中國社區 [http://www.oschina.net]
  | 本文標題：不兼容惹的禍，Calibre 作者拒絕遷移至 Python 3
  | 本文地址：https://www.oschina.net/news/102013/calibre-wont-convert-to-python3

.. _回應: https://bugs.launchpad.net/calibre/+bug/1714107
.. _Python 2.7 將於2020年1月1日終止支持: https://www.oschina.net/news/94198/python-2-7-quit
.. _回應公開: https://www.developpez.com/actu/233362/Calibre-le-gestionnaire-open-source-de-livres-numeriques-ne-va-pas-migrer-a-Python-3-car-l-auteur-s-estime-capable-de-maintenir-Python-2

Python 編程語言實行儘可能成熟、穩定的新管理模式
###############################################

:date: 2018-12-22T14:09+08:00
:tags: 轉錄, 開源中國, Python
:category: 轉錄
:author: 程六金(開源中國)
:summary: 在創始人和首席執行官 Guido van Rossum 辭去 BDFL 後，Python 軟件基金會已經確定了 Python 的新管理模式。
:og_image: https://upload.wikimedia.org/wikipedia/commons/f/f8/Python_logo_and_wordmark.svg


在創始人和首席執行官 Guido van Rossum 辭去 BDFL 後，Python 軟件基金會已經確定了 Python 的新管理模式。

新管理模型將依賴於一個由五人組成的指導委員會來建立標準實踐，以便為 Python 引入新功能。該提案被設計為“無趣”、全面、靈活和輕量。

『我們不是管理專家，我們認為 Python 並不是一個好的實驗對象，去實踐新的或者沒有經過證實的管理模式』，Nathaniel Smith 和 Donald Stufft 在 Python `文檔中`_ 解釋道。『所以這個模式是儘可能地成熟的、眾所周知的，且經過測試。大多數成功的 F/OSS 項目中最常見的是，一個大多數不干涉的理事會，而且低級的細節直接來自 Django 的管理模型』。

指導委員會將作為 Python 的“最終上訴法院”，並將對決策過程擁有廣泛的權力，包括接受或拒絕 PEP 的能力（Python 功能增強建議），執行和更新項目的行為準則，創建子委員會和管理項目資產。但 Nathaniel Smith 和 Donald Stufft 說，理事會的預期目標是採取更多不干涉和偶爾干涉的方式來發揮其權力。

『理事會應該儘可能少地尋找使用這些權力的方法。』Nathaniel Smith 和 Donald Stufft。『最好定義 PEP 決策的標準流程（例如，接受其他 801x 系列 PEP 之一），而不是對單個 PEP 進行裁決。建立行為準則委員會比制定個案更好。等等。為了利用其權力，理事會投票。每個理事會成員都必須投票或明確棄權。在特定投票中有利益衝突的成員必須棄權。通過需要大多數非棄權理事會成員的支持。』

指導委員會的任務是提供一種可訪問的，可維護的，形式化的引入變更的方法，基於『一般的理念，即將大的變化分成一系列可以獨立審查的小變化：而不是試圖在一個 PEP 中做所有事情，我們專注於為進一步的治理決策提供最小但堅實的基礎。』

..
  .. image:: 
   :alt: 
   :align: center

.. highlights::

  | 本站文章除註明轉載外，均為本站原創或編譯。歡迎任何形式的轉載，但請務必註明出處，尊重他人勞動共創開源社區。
  | 轉載請註明：文章轉載自 開源中國社區 [https://www.oschina.net]
  | 本文標題：Python 編程語言實行儘可能成熟、穩定的新管理模式
  | 本文地址：https://www.oschina.net/news/102915/python-programming-language-gets-new-governance-model

.. _文檔中: https://www.python.org/dev/peps/pep-8016/

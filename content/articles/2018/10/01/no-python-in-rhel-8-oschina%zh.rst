RHEL 8 中不再默認系統 Python 版本
#################################

:date: 2018-12-20T12:08+08:00
:tags: 轉錄, 開源中國, Python
:category: 轉錄
:author: h4cd(開源中國)
:summary: RHEL 8 中不再默認 Python 版本。
:og_image: https://oscimg.oschina.net/oscnet/877c33069ef330968126852c93f18b34267.jpg

RHEL 8 中不再默認 Python 版本。

.. image:: https://oscimg.oschina.net/oscnet/877c33069ef330968126852c93f18b34267.jpg
   :alt: RHEL 8 中不再默認系統 Python 版本
   :align: center

包括 RHEL 在內的大多數 Linux 發行版，用戶除非離開系統的包管理器環境，否則一般是被默認限定在系統提供的 Python 版本中。對於像 Ruby、Node、Perl 與 PHP 在內的許多工具來說這都是習以為常的，但是 Python 的情況會比較複雜一些，因為許多 Linux 工具（如 yum）都直接依賴於 Python。

根據 PEP 394，目前 /usr/bin/python 默認是指 Python2，也就是 “Python”這一命令或者 Python 解釋器將默認指向 Python2 版本。

Red Hat 官方 `在其開發者博客中發文`_ 稱，針對這一點，為了改善用戶體驗，從 RHEL 8 Beta 開始不再強調“系統 Python”，不再默認一個 Python 版本。他們使用模塊化的 Application Streams 設計，結合 Python 可多版本同時安裝的特點，將為用戶提供多個版本 Python 的選項，並且可以從標準存儲庫輕鬆安裝到標準位置，用戶可以選擇他們想要在任何給定用戶空間中運行的 Python 版本。

Application Streams 是在 RHEL 8 中引入的一類存儲庫，它提供用戶可能希望在給定用戶空間中運行的所有應用程序，它是在物理存儲庫中創建的多個虛擬存儲庫。

這種變化之後，用戶想要使用 Python，需要直接指定 Python3 或者 Python2，而不是直接 Python。同時 yum install python 將返回 404，因為它同樣需要指定安裝版本。建議使用 yum install @python36 或 yum install @python27 安裝推薦軟件包，而如果只需要 Python 二進制文件，則可以使用 yum install python3 或 yum install python2。此外，pip 等工具也有變化，比如 Python3 將安裝在 pip3 路徑下，而不是沒有版本指定的 pip 路徑。

Red Hat 解釋，除了提升用戶體驗，這種方案還讓方便了系統維護人員，因為不會被鎖定在系統中老版本的 Python 上，那麼他們可以自由地利用新版本的語言功能與性能改進等優勢。

..
  .. image:: 
   :alt: 
   :align: center

.. highlights::

  | 本站文章除註明轉載外，均為本站原創或編譯。歡迎任何形式的轉載，但請務必註明出處，尊重他人勞動共創開源社區。
  | 轉載請註明：文章轉載自 開源中國社區 [https://www.oschina.net]
  | 本文標題：RHEL 8 中不再默認系統 Python 版本
  | 本文地址：https://www.oschina.net/news/102857/no-python-in-rhel-8

.. _在其開發者博客中發文: https://developers.redhat.com/blog/2018/11/27/what-no-python-in-rhel-8-beta/

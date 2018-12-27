2018年 Wayland/Weston 發展數據一覽
##################################

:date: 2018-12-27T10:06+08:00
:tags: 轉錄, 開源中國
:category: 轉錄
:author: 程六金(開源中國)
:summary: 今年對於 Wayland 來說非常有意思，GNOME Shell 和 KDE Plasma 支持相繼成熟，較小的 i3-inspired Sway 也接近了1.0版本，NVIDIA 致力於 EGLStreams ，已完成對 KWin 的支持。與此同時，三星放棄了 Wayland 的開發商來作為OSG 重組的一部分，這對上游的項目做出了巨大貢獻。以下是 Wayland/Weston 2018 的數據。
:og_image: https://phoronix.com/misc/wayland-eoy2018/commits_by_year.png


今年對於 Wayland 來說非常有意思，GNOME Shell 和 KDE Plasma 支持相繼成熟，較小的 i3-inspired Sway 也接近了1.0版本，NVIDIA 致力於 EGLStreams ，已完成對 KWin 的支持。與此同時，三星放棄了 Wayland 的開發商來作為OSG 重組的一部分，這對上游的項目做出了巨大貢獻。以下是 Wayland/Weston 2018 的數據。

Wayland 本身並沒有很多的 Commit 活動，因為其核心已經相當成熟，開發商今年在定時發佈節奏上，與新的 Weston 版本聯合起來。

.. image:: https://phoronix.com/misc/wayland-eoy2018/commits_by_year.png
   :alt: 2018年 Wayland/Weston 發展數據一覽
   :align: center

迄今為止，Wayland 的 Git 存儲庫僅提交了 70 次，去年是 78 次，而前幾年有超過 150 次 提交。今年還有 792 個新代碼行數和 357 個刪除的代碼行數。這是歷史上一個低點。

今年對Wayland的貢獻最多的貢獻者是 Derek Foreman（三星），其次 Daniel Stone、Emil Velikov、Pekka Paalanen、Simon Ser、Michal Srb。通過 GitStats_ 可以看到更多關於 Wayland 數據。

除了所有其他桌面的 Wayland，Weston 今年只有 283 次提交，而去年有 437 次，而前幾年有 600多次提交。回到 2010 年，當時有 248 次。增加了 9.6k 行代碼，刪除了 3.6k 行代碼，自 2009年以來的最低點。

今年對 Weston 貢獻最多的貢獻者是 Collabora 的 Pekka Paalanen，其次是 Daniel Stone、Emre Ucan、Derek Foreman、Marius Vlad 和 Alexandros Frantzis。Weston 本身最多有 153343行代碼。你可以通過 `這些數據`_ 瞭解更多 Weston 的 Git 活動細節。

原文鏈接：https://www.phoronix.com/scan.php?page=news_item&px=Wayland-Weston-EOY-2018-Stats

..
  .. image:: 
   :alt: 
   :align: center

.. highlights::

  | 本站文章除註明轉載外，均為本站原創或編譯。歡迎任何形式的轉載，但請務必註明出處，尊重他人勞動共創開源社區。
  | 轉載請註明：文章轉載自 開源中國社區 [https://www.oschina.net]
  | 本文標題：2018年 Wayland/Weston 發展數據一覽
  | 本文地址：https://www.oschina.net/news/103073/wayland-weston-eoy-2018-stats

.. _GitStats: https://phoronix.com/misc/wayland-eoy2018/index.html
.. _這些數據: https://phoronix.com/misc/weston-eoy2018/index.html

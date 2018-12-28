Adiantum 和 Streebog 開始向用戶推送
###################################

:date: 2018-12-28T17:06+08:00
:tags: 轉錄, 開源中國, Linux
:category: 轉錄
:author: 程六金(開源中國)
:summary: Linux 4.21 內核的加密子系統更新在今天早上推送。
:og_image: https://upload.wikimedia.org/wikipedia/commons/3/35/Tux.svg


Linux 4.21 內核的加密子系統更新在今天早上推送。

在此次推送中， Adiantum_ 作為 Google Speck 計畫的代替品。用於假幣在 CPU/Soc 上缺少原生加密擴展的低端 Android Go 設備的數據加密。 `Adiantum 的表現超過了 Speck`_ ，最重要的是 Speck 是美國國家安全局開發的，而 Adiantum 不是。

Adiantum這次提交代碼的一部分，是把 Adiantum 請求構建在 XChaCha12/XChaCha20 之上。通過這種 Adiantum ，還可以在 ChaCha20 上進行各種性能改進，包括 ARM64和 x86，HPON的NEON/SSE2/AVX2加速以及其他加密性能工作。

提交代碼的另外一部分是 `Adiantum對fscrypt的支持`_ ，因此 Adiantum 對 EXT4 和 F2FS 之類也可以提供加密支持。對於Linux 4.21，預計這些 fscrypt 變化也將被加入。

另外 `一部分`_ 是添加了 `Streebog散列函數`_ 。Streebog 由俄羅斯 FSB 安全服務和其他組織開發。

有關 Linux4.21所有加密改進信息，請 `查看這裡`_ 。

英文原文：https://www.phoronix.com/scan.php?page=news_item&px=Linux-4.21-Crypto

..
  .. image:: 
   :alt: 
   :align: center

.. highlights::

  | 本站文章除註明轉載外，均為本站原創或編譯。歡迎任何形式的轉載，但請務必註明出處，尊重他人勞動共創開源社區。
  | 轉載請註明：文章轉載自 開源中國社區 [https://www.oschina.net]
  | 本文標題：Adiantum 和 Streebog 開始向用戶推送
  | 本文地址：https://www.oschina.net/news/103105/linux-4.21-crypto

.. _Adiantum: https://www.phoronix.com/scan.php?page=search&q=Adiantum
.. _Adiantum 的表現超過了 Speck: https://www.phoronix.com/scan.php?page=news_item&px=Adiantum-Crypto-Linux-Coming
.. _Adiantum對fscrypt的支持: https://www.phoronix.com/scan.php?page=news_item&px=Adiantum-Fscrypt-Linux-4.21
.. _一部分: https://www.phoronix.com/scan.php?page=news_item&px=Linux-4.21-Streebog-Crypto
.. _Streebog散列函數: https://www.phoronix.com/scan.php?page=news_item&px=Linux-4.21-Streebog-Crypto
.. _查看這裡: https://lkml.org/lkml/2018/12/26/121

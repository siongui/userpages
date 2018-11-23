製作Ubuntu Linux開機及備份外接USB碟
###################################

:date: 2018-11-24T04:16+08:00
:tags: Linux, Ubuntu Linux
:category: Linux
:summary: 製作Ubuntu Linux開機外接USB碟，並把資料備份在此碟內。
:og_image: http://4.bp.blogspot.com/__fGE9uGYw5A/S0WDm4UdKgI/AAAAAAAABEI/xitst6i9aiQ/s400/Bootable+Ubuntu+USB+Flash.jpg
:adsu: yes


**本文會將USB碟資料全部銷毀，在往下執行前請先備份USB碟裡的資料**

現今有許多USB外接硬碟以及USB外接隨身碟，使用上很方便。
本文介紹利用這些外接USB碟來製作Ubuntu Linux開機USB碟，並把資料儲存在其內。

步驟：

1. 製作Ubuntu Linux開機USB碟(*本步驟會銷毀USB碟內資料，請注意*)：下載Ubuntu
   Linux ISO檔，並參閱 [3]_ 或 [4]_ 來製作開機USB碟。

2. 利用 fdisk_ 指令把剩下的空間建立新的partition，並用 mkfs.ext4_
   或類似指令格式化新建立的partition。然後再把資料 copy 進來即可。

----

.. adsu:: 2

參考：

.. [1] `全新安裝Ubuntu Linux注意事項 <{filename}/articles/2016/02/29/clean-install-ubuntu-linux-notes%zh.rst>`_
.. [2] `kazam 螢幕錄影程式 (Ubuntu Linux 17.10) <{filename}/articles/2018/01/28/kazam-screen-recorder-ubuntu-linux-17-10%zh.rst>`_
.. [3] `How to Create Bootable Ubuntu USB flash Drive from Terminal <https://linoxide.com/linux-how-to/create-bootable-ubuntu-usb-flash-drive-terminal/>`_
.. [4] `How to Create Bootable USB Disk / DVD on Ubuntu / Linux Mint <https://www.linuxtechi.com/create-bootable-usb-disk-dvd-ubuntu-linux-mint/>`_

.. _fdisk: https://duckduckgo.com/?q=fdisk+linux
.. _mkfs.ext4: https://duckduckgo.com/?q=mkfs.ext4

全新安裝Ubuntu Linux注意事項
############################

:date: 2016-02-29T01:39+08:00
:modified: 2020-08-31T07:01+08:00
:tags: Linux, apt command, Ubuntu Linux
:category: Linux
:summary: 全新安裝Ubuntu Linux注意事項
:og_image: http://design.ubuntu.com/wp-content/uploads/logo-ubuntu_st_no%C2%AE-black_orange-hex.jpg
:adsu: yes


- 下載 `Ubuntu 20.04.1 LTS (Focal Fossa) <https://releases.ubuntu.com/20.04/>`_

- 檢查ISO檔
  (`HowToMD5SUM - Community Help Wiki <https://help.ubuntu.com/community/HowToMD5SUM>`_):

  .. code-block:: bash

    $ md5sum -c MD5SUMS

- 開啟 Startup Disk Creator燒錄
  (`How to create a bootable USB stick on Ubuntu <https://www.ubuntu.com/download/desktop/create-a-usb-stick-on-ubuntu>`_).

  * 會毀掉整個USB裝置裡的資料，並佔用2GB多的空間(佔用sdb1, sdb2)

    .. code-block:: bash

      $ sudo fdisk /dev/sdb
      $ sudo mkfs.ext4 /dev/sdb3

- `ubuntu fail to install grub lenovo - Google search <https://www.google.com/search?q=ubuntu+fail+to+install+grub+lenovo>`_,
  `ubuntu fail to install grub - Google search <https://www.google.com/search?q=ubuntu+fail+to+install+grub>`_

  * BIOS裡選legacy,不選UEFI
  * 安裝時問是否使用UEFI,選擇Go back(不要用UEFI)

.. adsu:: 2

- 更新系統 [1]_

  .. code-block:: bash

    $ sudo apt-get update && sudo apt-get upgrade && sudo apt-get dist-upgrade

- `git autocomplete <https://www.google.com/search?q=git+autocomplete>`_

  .. code-block:: bash

    $ sudo apt-get install git bash-completion

- 傳統中文 language pack

  .. code-block:: bash

    $ sudo apt-get install language-pack-zh-hant

- 工具(Vim_, Chromium_)

  .. code-block:: bash

    $ sudo apt-get install vim
    $ sudo apt-get install chromium-browser

.. adsu:: 3

- Chromium extension and app

  * `uBlock Origin <https://chrome.google.com/webstore/detail/ublock-origin/cjpalhdlnbpafiamejdnhcphjbkeiagm?hl=en>`_
  * `Google Translate <https://chrome.google.com/webstore/detail/google-translate/aapbdbdomjkkjkaonfhkkikfgjllcleb?hl=en>`_
  * `新同文堂 <https://chrome.google.com/webstore/detail/new-tong-wen-tang/ldmgbgaoglmaiblpnphffibpbfchjaeg?hl=zh-TW>`_
  * `LINE <https://chrome.google.com/webstore/detail/line/menkifleemblimdogmoihpfopnplikde?hl=en>`_

- `ubuntu chinese input - Google search <https://www.google.com/search?q=ubuntu+chinese+input>`_,
  `[Ubuntu] 在 Ubuntu 18.04 中新增新酷音輸入法 | by MuLong PuYang | Medium <https://medium.com/@racktar7743/ubuntu-%E5%9C%A8-ubuntu-18-04-%E4%B8%AD%E6%96%B0%E5%A2%9E%E6%96%B0%E9%85%B7%E9%9F%B3%E8%BC%B8%E5%85%A5%E6%B3%95-4aa85782f656>`_

- Lenovo筆電使用rtl8821ae無線網路斷線:
  請看 `台灣最佳CP值無作業系統(OS)筆記型電腦 <{filename}../26/best-cp-no-os-notebook-in-taiwan%zh.rst>`_

- 複製檔案數量多時，請看
  `[Bash] Copy Large Number of Files on Linux <{filename}../../12/20/bash-copy-large-number-of-files-on-linux%en.rst>`_

- 複製檔案時，注意權限問題:

  .. code-block:: bash

    $ sudo cp -r src dst
    $ sudo chown -R usr:grp *


----

.. adsu:: 4

參考：

.. [1] `update ubuntu system command line - Google search <https://www.google.com/search?q=update+ubuntu+system+command+line>`_

       `update ubuntu system command line - DuckDuckGo search <https://duckduckgo.com/?q=update+ubuntu+system+command+line>`_

       `update ubuntu system command line - Bing search <https://www.bing.com/search?q=update+ubuntu+system+command+line>`_

       `update ubuntu system command line - Yahoo search <https://search.yahoo.com/search?p=update+ubuntu+system+command+line>`_

       `update ubuntu system command line - Baidu search <https://www.baidu.com/s?wd=update+ubuntu+system+command+line>`_

       `update ubuntu system command line - Yandex search <https://www.yandex.com/search/?text=update+ubuntu+system+command+line>`_

       `command line - Update the system from terminal - Ask Ubuntu <http://askubuntu.com/questions/462449/update-the-system-from-terminal>`_



.. _Vim: http://www.vim.org/
.. _Chromium: https://www.chromium.org/

kazam 螢幕錄影程式 (Ubuntu Linux 17.10)
#######################################

:date: 2018-01-28T04:36+08:00
:tags: Linux, Ubuntu Linux
:category: Linux
:summary: 在 Ubuntu Linux 17.10 下安裝螢幕錄影程式 kazam 時遇到的麻煩
          ，以及解決方法。
:og_image: https://1.bp.blogspot.com/-JCi7zJYiL_g/UTJ7cK845OI/AAAAAAAACbA/t226VUrpGwg/s1000/kazam-1.jpg
:adsu: yes

`Ubuntu Linux 17.10`_ 將 Xorg 換成 Wayland ，讓原本用的好好的螢幕錄影程式 kazam_
不能用了。主要有兩個問題：

1. kazam 無法啟動 [2]_ 解法： `將快捷鍵停用`_

2. kazam 錄影時只有聲音沒有畫面 [3]_ 解法： `改回用 Xorg`_

----

參考：

.. [1] | `ubuntu 17.10 screen recorder - Google search <https://www.google.com/search?q=ubuntu+17.10+screen+recorder>`_
       | `ubuntu 17.10 screen recorder - DuckDuckGo search <https://duckduckgo.com/?q=ubuntu+17.10+screen+recorder>`_
       | `ubuntu 17.10 screen recorder - Ecosia search <https://www.ecosia.org/search?q=ubuntu+17.10+screen+recorder>`_
       | `ubuntu 17.10 screen recorder - Qwant search <https://www.qwant.com/?q=ubuntu+17.10+screen+recorder>`_
       | `ubuntu 17.10 screen recorder - Bing search <https://www.bing.com/search?q=ubuntu+17.10+screen+recorder>`_
       | `ubuntu 17.10 screen recorder - Yahoo search <https://search.yahoo.com/search?p=ubuntu+17.10+screen+recorder>`_
       | `ubuntu 17.10 screen recorder - Baidu search <https://www.baidu.com/s?wd=ubuntu+17.10+screen+recorder>`_
       | `ubuntu 17.10 screen recorder - Yandex search <https://www.yandex.com/search/?text=ubuntu+17.10+screen+recorder>`_

.. [2] | `Kazam not working on Python 3.5.3 · Issue #4 · sconts/kazam · GitHub <https://github.com/sconts/kazam/issues/4>`_
       | `python3 - kazam fails with "PyGIWarning: Gtk was imported without specifying a version first..." - Ask Ubuntu <https://askubuntu.com/questions/982233/kazam-fails-with-pygiwarning-gtk-was-imported-without-specifying-a-version-fir>`_
.. adsu:: 2
.. [3] | `software recommendation - Record screen, microphone and web camera in Ubuntu 17.10 (Wayland Session) - Ask Ubuntu <https://askubuntu.com/questions/970524/record-screen-microphone-and-web-camera-in-ubuntu-17-10-wayland-session>`_
       | `How do you switch from Wayland back to Xorg in Ubuntu 17.10? - Ask Ubuntu <https://askubuntu.com/questions/961304/how-do-you-switch-from-wayland-back-to-xorg-in-ubuntu-17-10>`_

.. _Ubuntu Linux 17.10: https://www.ubuntu.com/desktop/1710
.. _kazam: https://github.com/sconts/kazam
.. _將快捷鍵停用: https://askubuntu.com/a/989341
.. _改回用 Xorg: https://askubuntu.com/a/968265

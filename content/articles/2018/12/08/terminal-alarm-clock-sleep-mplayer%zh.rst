使用終端機命令行來做鬧鐘
########################

:date: 2018-12-08T00:57+08:00
:tags: Linux, 命令行
:category: Linux
:summary: 利用 *sleep* 以及 *mplayer* 指令來來製作終端機鬧鐘。
:og_image: https://s10629.pcdn.co/wp-content/pictures/2011/10/Open_Terminal2.png
:adsu: yes

利用 sleep_ 以及 mplayer_ 指令來來製作不需要圖形界面的終端機(terminal)鬧鐘。
首先我不知道該怎樣做，搜尋了一下 [1]_ ，發現了 [2]_ 這篇文章，
講述了一個很簡單的方式實現終端機鬧鐘：

.. code-block:: bash

  $ sleep 5m && mplayer /path/to/your.mp3

上述指令在按下enter之後五分鐘，就會播放指定的mp3檔案！簡單又方便的小鬧鐘！

----

參考：

.. [1] `command line alarm clock at DuckDuckGo <https://duckduckgo.com/?q=command+line+alarm+clock>`_
.. [2] `Terminal Tricks: Use the Terminal as an alarm clock <https://helpdeskgeek.com/linux-tips/terminal-tricks-use-the-terminal-as-an-alarm-clock/>`_

.. _sleep: https://linux.die.net/man/1/sleep
.. _mplayer: https://linux.die.net/man/1/mplayer

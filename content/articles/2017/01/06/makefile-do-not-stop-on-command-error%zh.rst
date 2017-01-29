Makefile 執行指令遇到錯誤時，繼續執行
#####################################

:date: 2017-01-06T21:39+08:00
:tags: Linux, Makefile
:category: Linux
:summary: 執行指令遇到錯誤時，繼續執行Makefile裡剩下的指令，而不停止執行。
:adsu: yes


舉例說明，假設執行下面指令：

.. code-block:: bash

  sudo apt-get install opencc
  go get -u github.com/siongui/go-opencc

若第一行執行時發生錯誤（比方說安裝失敗），第二行將不會執行。

從 [2]_ 的回答，若在第一行前面加 ``-`` ，則第一行執行失敗時
，仍然會繼續往下執行，不會停止：

.. code-block:: bash

  -sudo apt-get install opencc
  go get -u github.com/siongui/go-opencc

.. adsu:: 2

----

References:

.. [1] `makefile do not stop on error - Google search <https://www.google.com/search?q=makefile+do+not+stop+on+error>`_

       `makefile do not stop on error - DuckDuckGo search <https://duckduckgo.com/?q=makefile+do+not+stop+on+error>`_

       `makefile do not stop on error - Bing search <https://www.bing.com/search?q=makefile+do+not+stop+on+error>`_

       `makefile do not stop on error - Yahoo search <https://search.yahoo.com/search?p=makefile+do+not+stop+on+error>`_

       `makefile do not stop on error - Baidu search <https://www.baidu.com/s?wd=makefile+do+not+stop+on+error>`_

       `makefile do not stop on error - Yandex search <https://www.yandex.com/search/?text=makefile+do+not+stop+on+error>`_

.. [2] `makefile - Make: how to continue after a command fails? - Stack Overflow <http://stackoverflow.com/a/2670143>`_

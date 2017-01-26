Vim編輯器移除每行開頭的tab空白
##############################

:date: 2017-01-27T02:58+08:00
:tags: Linux, Vim
:category: Linux
:summary: Vim_ 編輯器: 移除每行開頭的 `tab空白`_
:og_image: https://camo.githubusercontent.com/1cbe7de4dee3a64e32dabbb1dad6fc692e5a2b89/687474703a2f2f692e696d6775722e636f6d2f486e7047472e706e67
:adsu: yes

Vim_ 編輯器: 移除每行開頭的 `tab空白`_

.. code-block:: vim

  :%s/^\t//gc

其中， ``%s`` 意思是取代， ``^\t`` 意思是開頭的 `tab空白`_ ，
``^`` 意思是開頭， ``\t`` 是 `tab空白`_ 。
``//`` 斜線中間沒有任何東西，意思就是移除。
``g`` 意思是 global，意思是找出檔案裡所有的 ``^\t`` ，
若沒有 ``g`` ，則只會取代找到的第一個 ``^\t`` 。
``c`` 意思是 confirm，意思要取代前先問是否要取代。

.. adsu:: 2

若要把每行開頭的 `tab空白`_ 取代成兩個一般空白，
可用下面指令：

.. code-block:: vim

  :%s/^\t/  /gc

----

參考：

.. [1] `vim replace gc - Google search <https://www.google.com/search?q=vim+replace+gc>`_

       `vim replace gc - DuckDuckGo search <https://duckduckgo.com/?q=vim+replace+gc>`_

       `vim replace gc - Bing search <https://www.bing.com/search?q=vim+replace+gc>`_

       `vim replace gc - Yahoo search <https://search.yahoo.com/search?p=vim+replace+gc>`_

       `vim replace gc - Baidu search <https://www.baidu.com/s?wd=vim+replace+gc>`_

       `vim replace gc - Yandex search <https://www.yandex.com/search/?text=vim+replace+gc>`_

.. _Vim: http://www.vim.org/
.. _tab空白: https://zh.wikipedia.org/wiki/%E8%A3%BD%E8%A1%A8%E9%8D%B5

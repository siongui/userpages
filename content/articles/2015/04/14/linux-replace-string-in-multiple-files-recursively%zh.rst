Linux系統，取代目錄(含子目錄)下所有檔案裡字串
#############################################

:date: 2015-04-14 01:03
:tags: Linux, 字串處理
:category: Linux
:summary: 將某個目錄(含子目錄)下，所有檔案裡的某個字串(string)取代成另一個字串。

問題
++++

將 *content* 目錄下 ``http://forestmeditation.com/audio/files/``
這個字串取代成 ``/7rsk9vjkm4p8z5xrdtqc/audio/ForestMeditation/``
，該如何做呢？

解答
++++

經過Google搜尋後，找到 [1]_ 跟 [2]_ ，可利用 find_ 跟 sed_
來達成任務：

先進入 *content* 目錄，然後打：

.. code-block:: bash

  find . -type f -exec sed -i 's/http:\/\/forestmeditation.com\/audio\/files\//\/7rsk9vjkm4p8z5xrdtqc\/audio\/ForestMeditation\//g' {} +

(利用 ``\`` 來 escape_ ``/``)

----

References:

.. [1] `text processing - How can I replace a string in a file(s)? - Unix & Linux Stack Exchange <http://unix.stackexchange.com/questions/112023/how-can-i-replace-a-string-in-a-files>`_

.. [2] `How to replace a string in multiple files in linux command line - Stack Overflow <http://stackoverflow.com/questions/11392478/how-to-replace-a-string-in-multiple-files-in-linux-command-line>`_

.. [3] `bash - Escape a string for a sed replace pattern - Stack Overflow <http://stackoverflow.com/questions/407523/escape-a-string-for-a-sed-replace-pattern>`_

.. [4] `三分钟教你轻松掌握 grep 命令中的正则表达式 - 博客 - 伯乐在线 <http://blog.jobbole.com/98134/>`_


.. _find: http://content.hccfl.edu/pollock/Unix/FindCmd.htm

.. _sed: http://www.grymoire.com/Unix/Sed.html

.. _escape: http://stackoverflow.com/questions/407523/escape-a-string-for-a-sed-replace-pattern

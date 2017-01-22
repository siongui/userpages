[sed] 在匹配前插入空行
######################

:date: 2017-01-23T05:03+08:00
:tags: Linux, sed
:category: Linux
:summary: 利用 sed_ 搜尋字串，若匹配則在此字串前插入空行
:adsu: yes


假設在某個目錄下，有許多副檔名為 `.rst` 的檔案，
若檔案裡某一行開頭為 ``:adsu`` ，則在此行前插入空行。

舉例來說，假設檔案內容如下：

.. code-block:: txt

  hello
  world
  :adsu
  你好
  世界

經過 sed_ 處理後，變成如下：

.. code-block:: txt

  hello
  world

  :adsu
  你好
  世界

利用 Bash_ script以及 sed_ 來達成此任務：

.. code-block:: bash

  #!/bin/bash

  rstdir=content/articles/

  for path in $(find $rstdir -type f -name "*.rst")
  do
    sed -i '0,/:adsu/s/:adsu/\n&/' $path
  done


----

References:

.. [1] `linux - Command to insert lines before first match - Stack Overflow <http://stackoverflow.com/questions/30386483/command-to-insert-lines-before-first-match>`_

.. [2] `[Bash] Find Redundant Files Saved by Chrome <{filename}../../../2016/05/23/bash-find-redundant-files-saved-by-chrome%en.rst>`_

.. _sed: https://www.google.com/search?q=sed
.. _Bash: https://www.google.com/search?q=Bash

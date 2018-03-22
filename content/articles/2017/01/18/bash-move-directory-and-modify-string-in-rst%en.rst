[Bash] Move Directories and Modify Path in Files
################################################

:date: 2017-01-18T23:24+08:00
:tags: sed, find command, Bash, Commandline, String Manipulation,
       List Files in Directory, Regular Expression
:category: Bash
:summary: Replace specific pattern of strings in HTML files under directory via
          sed_.
:og_image: https://image.slidesharecdn.com/linuxtraining-130801203338-phpapp01/95/linux-training-15-638.jpg
:adsu: yes


Question
++++++++

In the ``code/`` directory, there are sub-directories like the following:

  | bash/
  | bash-big5-to-utf8/
  | bash-redundant-file/
  | bash-wget/

I want to move the ``bash-*/`` sub-directories to ``bash/`` sub-directory and
remove the ``bash-`` prefix in the name of ``bash-*/`` sub-directories. The
result will be

  | bash/big5-to-utf8/
  | bash/redundant-file/
  | bash/wget/

And also modify the string ``bash-*/`` to ``bash/*/`` via sed_ in the
reStructuredText_ files in ``articles`` directory.

.. adsu:: 2

Answer
++++++

.. show_github_file:: siongui userpages content/code/bash/move-dir-modify-rst/run.sh
.. adsu:: 3

----

References:

.. [1] `bash list directories - Google search <https://www.google.com/search?q=bash+list+directories>`_

       `bash list directories - DuckDuckGo search <https://duckduckgo.com/?q=bash+list+directories>`_

       `bash list directories - Bing search <https://www.bing.com/search?q=bash+list+directories>`_

       `bash list directories - Yahoo search <https://search.yahoo.com/search?p=bash+list+directories>`_

       `bash list directories - Baidu search <https://www.baidu.com/s?wd=bash+list+directories>`_

       `bash list directories - Yandex search <https://www.yandex.com/search/?text=bash+list+directories>`_

       `directory - Listing only directories using ls in bash: An examination - Stack Overflow <http://stackoverflow.com/a/17009555>`_

.. [2] `bash basename - Google search <https://www.google.com/search?q=bash+basename>`_

       `bash basename - DuckDuckGo search <https://duckduckgo.com/?q=bash+basename>`_

       `bash basename - Bing search <https://www.bing.com/search?q=bash+basename>`_

       `bash basename - Yahoo search <https://search.yahoo.com/search?p=bash+basename>`_

       `bash basename - Baidu search <https://www.baidu.com/s?wd=bash+basename>`_

       `bash basename - Yandex search <https://www.yandex.com/search/?text=bash+basename>`_

       `linux - Extract File Basename Without Path and Extension in Bash - Stack Overflow <http://stackoverflow.com/a/2664746>`_

.. [3] `bash string starts with - Google search <https://www.google.com/search?q=bash+string+starts+with>`_

       `bash string starts with - DuckDuckGo search <https://duckduckgo.com/?q=bash+string+starts+with>`_

       `bash string starts with - Bing search <https://www.bing.com/search?q=bash+string+starts+with>`_

       `bash string starts with - Yahoo search <https://search.yahoo.com/search?p=bash+string+starts+with>`_

       `bash string starts with - Baidu search <https://www.baidu.com/s?wd=bash+string+starts+with>`_

       `bash string starts with - Yandex search <https://www.yandex.com/search/?text=bash+string+starts+with>`_

       `In bash, how can I check if a string begins with some value? - Stack Overflow <http://stackoverflow.com/questions/2172352/in-bash-how-can-i-check-if-a-string-begins-with-some-value>`_

.. [4] `bash remove prefix - Google search <https://www.google.com/search?q=bash+remove+prefix>`_

       `bash remove prefix - DuckDuckGo search <https://duckduckgo.com/?q=bash+remove+prefix>`_

       `bash remove prefix - Bing search <https://www.bing.com/search?q=bash+remove+prefix>`_

       `bash remove prefix - Yahoo search <https://search.yahoo.com/search?p=bash+remove+prefix>`_

       `bash remove prefix - Baidu search <https://www.baidu.com/s?wd=bash+remove+prefix>`_

       `bash remove prefix - Yandex search <https://www.yandex.com/search/?text=bash+remove+prefix>`_

       `bash - remove a fixed prefix/suffix from a string - Stack Overflow <http://stackoverflow.com/a/16623897>`_

.. [5] `bash dirname - Google search <https://www.google.com/search?q=bash+dirname>`_

       `bash dirname - DuckDuckGo search <https://duckduckgo.com/?q=bash+dirname>`_

       `bash dirname - Bing search <https://www.bing.com/search?q=bash+dirname>`_

       `bash dirname - Yahoo search <https://search.yahoo.com/search?p=bash+dirname>`_

       `bash dirname - Baidu search <https://www.baidu.com/s?wd=bash+dirname>`_

       `bash dirname - Yandex search <https://www.yandex.com/search/?text=bash+dirname>`_

.. [6] `bash test if directory exists - Google search <https://www.google.com/search?q=bash+test+if+directory+exists>`_

       `bash test if directory exists - DuckDuckGo search <https://duckduckgo.com/?q=bash+test+if+directory+exists>`_

       `bash test if directory exists - Bing search <https://www.bing.com/search?q=bash+test+if+directory+exists>`_

       `bash test if directory exists - Yahoo search <https://search.yahoo.com/search?p=bash+test+if+directory+exists>`_

       `bash test if directory exists - Baidu search <https://www.baidu.com/s?wd=bash+test+if+directory+exists>`_

       `bash test if directory exists - Yandex search <https://www.yandex.com/search/?text=bash+test+if+directory+exists>`_

.. [7] `bash string replace character - Google search <https://www.google.com/search?q=bash+string+replace+character>`_

       `bash string replace character - DuckDuckGo search <https://duckduckgo.com/?q=bash+string+replace+character>`_

       `bash string replace character - Bing search <https://www.bing.com/search?q=bash+string+replace+character>`_

       `bash string replace character - Yahoo search <https://search.yahoo.com/search?p=bash+string+replace+character>`_

       `bash string replace character - Baidu search <https://www.baidu.com/s?wd=bash+string+replace+character>`_

       `bash string replace character - Yandex search <https://www.yandex.com/search/?text=bash+string+replace+character>`_

       `linux - Replacing some chars with another - Stack Overflow <http://stackoverflow.com/a/27369375>`_

.. [8] `sed process multiple files - Google search <https://www.google.com/search?q=sed+process+multiple+files>`_

       `sed process multiple files - DuckDuckGo search <https://duckduckgo.com/?q=sed+process+multiple+files>`_

       `sed process multiple files - Bing search <https://www.bing.com/search?q=sed+process+multiple+files>`_

       `sed process multiple files - Yahoo search <https://search.yahoo.com/search?p=sed+process+multiple+files>`_

       `sed process multiple files - Baidu search <https://www.baidu.com/s?wd=sed+process+multiple+files>`_

       `sed process multiple files - Yandex search <https://www.yandex.com/search/?text=sed+process+multiple+files>`_

       `sed - Change multiple files - Stack Overflow <http://stackoverflow.com/a/30717770>`_

.. [9] `sed single quote - Google search <https://www.google.com/search?q=sed+single+quote>`_

       `sed single quote - DuckDuckGo search <https://duckduckgo.com/?q=sed+single+quote>`_

       `sed single quote - Bing search <https://www.bing.com/search?q=sed+single+quote>`_

       `sed single quote - Yahoo search <https://search.yahoo.com/search?p=sed+single+quote>`_

       `sed single quote - Baidu search <https://www.baidu.com/s?wd=sed+single+quote>`_

       `sed single quote - Yandex search <https://www.yandex.com/search/?text=sed+single+quote>`_

       `escaping - How to escape single quote in sed? - Stack Overflow <http://stackoverflow.com/a/24509931>`_


.. _sed: https://www.google.com/search?q=sed
.. _reStructuredText: https://www.google.com/search?q=reStructuredText

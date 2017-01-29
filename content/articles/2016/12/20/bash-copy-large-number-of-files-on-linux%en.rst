[Bash] Copy Large Number of Files on Linux
##########################################

:date: 2016-12-20T20:36+08:00
:tags: Bash, Commandline, tar command
:category: Bash
:summary: Copy large number of files efficiently via Bash_ script on Linux.
:adsu: yes


If you want to copy large number of files, it's not a good idea to use
``cp -r`` command directly because it will take a lot of time. It's better to
use tar_ command to archive the files into a single file and then extract the
file to the destination directory you want. The following is the Bash_ script
to demonstrate how to copy large number of files efficiently:

.. show_github_file:: siongui userpages content/code/bash/copy-large-number-of-files/cpn.sh
.. adsu:: 2

----

Tested on: ``Ubuntu Linux 16.04``, ``bash 4.3.46(1)``.

----

References:

.. [1] `copy large number of files linux - Google search <https://www.google.com/search?q=copy+large+number+of+files+linux>`_

       `copy large number of files linux - DuckDuckGo search <https://duckduckgo.com/?q=copy+large+number+of+files+linux>`_

       `copy large number of files linux - Bing search <https://www.bing.com/search?q=copy+large+number+of+files+linux>`_

       `copy large number of files linux - Yahoo search <https://search.yahoo.com/search?p=copy+large+number+of+files+linux>`_

       `copy large number of files linux - Baidu search <https://www.baidu.com/s?wd=copy+large+number+of+files+linux>`_

       `copy large number of files linux - Yandex search <https://www.yandex.com/search/?text=copy+large+number+of+files+linux>`_

.. [2] `[Bash] Move Large Number of Files via tar Command <{filename}../../04/29/bash-move-large-number-of-files-via-tar%en.rst>`_


.. _Bash: https://www.google.com/search?q=Bash
.. _tar: http://www.tecmint.com/18-tar-command-examples-in-linux/

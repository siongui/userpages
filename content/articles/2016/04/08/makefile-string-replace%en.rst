[Makefile] String Replacement Example
#####################################

:date: 2016-04-08T03:41+08:00
:tags: Bash, Makefile, Commandline, String Manipulation
:category: Bash
:summary: Example for how to replace a string in a variable in Makefile_.

| First example: replace ``.`` with ``-``
| i.e.,
| ``mn.004.contrast-reading.html`` => ``mn-004-contrast-reading-html``
|
|
| Second example: replace ending ``.html`` with ``%zh.rst``
| i.e.,
| ``mn.004.contrast-reading.html`` => ``mn.004.contrast-reading%zh.rst``

----

Source code:

.. show_github_file:: siongui userpages content/code/make-replace-string/Makefile

Output:

.. code-block:: txt

  mn-004-contrast-reading-html
  mn.004.contrast-reading%zh.rst

----

Tested on ``Ubuntu Linux 15.10``, `GNU make 4.0-8.2`_.

----

References:

.. [1] `shell replace character in string - Google search <https://www.google.com/search?q=shell+replace+character+in+string>`_

       `shell replace character in string - DuckDuckGo search <https://duckduckgo.com/?q=shell+replace+character+in+string>`_

       `shell replace character in string - Bing search <https://www.bing.com/search?q=shell+replace+character+in+string>`_

       `shell replace character in string - Yahoo search <https://search.yahoo.com/search?p=shell+replace+character+in+string>`_

       `shell replace character in string - Baidu search <https://www.baidu.com/s?wd=shell+replace+character+in+string>`_

       `shell replace character in string - Yandex search <https://www.yandex.com/search/?text=shell+replace+character+in+string>`_

.. [2] `makefile replace string in variable - Google search <https://www.google.com/search?q=makefile+replace+string+in+variable>`_

       `makefile replace string in variable - DuckDuckGo search <https://duckduckgo.com/?q=makefile+replace+string+in+variable>`_

       `makefile replace string in variable - Bing search <https://www.bing.com/search?q=makefile+replace+string+in+variable>`_

       `makefile replace string in variable - Yahoo search <https://search.yahoo.com/search?p=makefile+replace+string+in+variable>`_

       `makefile replace string in variable - Baidu search <https://www.baidu.com/s?wd=makefile+replace+string+in+variable>`_

       `makefile replace string in variable - Yandex search <https://www.yandex.com/search/?text=makefile+replace+string+in+variable>`_

       `GNU make - Functions for Transforming Text <ftp://ftp.gnu.org/old-gnu/Manuals/make-3.79.1/html_chapter/make_8.html>`_

.. [3] `makefile file name manipulation · twnanda/twnanda@0323a19 · GitHub <https://github.com/twnanda/twnanda/commit/0323a193209a72041d7edb9e571125a5ce033844>`_

.. _Makefile: https://www.google.com/search?q=Makefile
.. _notdir: https://www.gnu.org/software/make/manual/html_node/File-Name-Functions.html
.. _GNU make 4.0-8.2: http://packages.ubuntu.com/wily/make

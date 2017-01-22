[Makefile] Get Prefix or Suffix from String
###########################################

:date: 2016-12-29T21:28+08:00
:tags: Bash, Makefile, Commandline, String Manipulation, Ubuntu Linux,
       tail command, head command
:category: Bash
:summary: Get prefix_ or suffix_ of a string in Makefile_.
:adsu: yes


**Source code**:

.. show_github_file:: siongui userpages content/code/make/string-prefix-suffix/Makefile


shell command head_/tail_ are used to get prefix_/suffix_ from string.

**Output**:

.. code-block:: txt

  string: 1.0.4-1ubuntu0.16.10.1
  prefix: 1.0.4
  suffix: 0.16.10.1

----

Tested on `Ubuntu Linux 16.10`_, `GNU make 4.1-9`_.

----

References:

.. [1] `[Makefile] Check Installed Package Version on Ubuntu Linux <{filename}../23/makefile-check-installed-package-version-on-ubuntu-linux%en.rst>`_

.. [2] `makefile string starts with - Google search <https://www.google.com/search?q=makefile+string+starts+with>`_

       `makefile string starts with - DuckDuckGo search <https://duckduckgo.com/?q=makefile+string+starts+with>`_

       `makefile string starts with - Bing search <https://www.bing.com/search?q=makefile+string+starts+with>`_

       `makefile string starts with - Yahoo search <https://search.yahoo.com/search?p=makefile+string+starts+with>`_

       `makefile string starts with - Baidu search <https://www.baidu.com/s?wd=makefile+string+starts+with>`_

       `makefile string starts with - Yandex search <https://www.yandex.com/search/?text=makefile+string+starts+with>`_

       `Get the first letter of a make variable - Stack Overflow <http://stackoverflow.com/questions/12798666/get-the-first-letter-of-a-make-variable>`_

       `makefile if else - Google search <https://www.google.com/search?q=makefile+if+else>`_

       `ugly fix for OpenCC build issue · siongui/pali@f2d3adf · GitHub <https://github.com/siongui/pali/commit/f2d3adf00117de0b146b3e30c4f0955d33205c02>`_

.. _Makefile: https://www.google.com/search?q=Makefile
.. _prefix: https://www.google.com/search?q=prefix
.. _suffix: https://www.google.com/search?q=suffix
.. _head: https://linux.die.net/man/1/head
.. _tail: https://linux.die.net/man/1/tail
.. _Ubuntu Linux 16.10: http://releases.ubuntu.com/16.10/
.. _GNU make 4.1-9: https://www.gnu.org/software/make/

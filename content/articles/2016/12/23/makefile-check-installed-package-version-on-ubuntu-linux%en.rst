[Makefile] Check Installed Package Version on Ubuntu Linux
##########################################################

:date: 2016-12-23T21:33+08:00
:tags: Bash, Makefile, Commandline, String Manipulation, Ubuntu Linux,
       grep command, cut command, apt command
:category: Bash
:summary: Given an installed package name on Ubuntu Linux, find its version in
          Makefile_.


Given an installed package name on Ubuntu Linux, find its version in Makefile_.

Source code:

.. show_github_file:: siongui userpages content/code/make/ubuntu-package-version/Makefile

Note that `strip text function`_ is used to remove spaces of the string.

Output:

.. code-block:: txt

  package name: python
  package version: 2.7.11-2

----

Tested on `Ubuntu Linux 16.10`_, `GNU make 4.1-9`_.

----

References:

.. [1] `ubuntu package version command line - Google search <https://www.google.com/search?q=ubuntu+package+version+command+line>`_

       `ubuntu check package version command line - Google search <https://www.google.com/search?q=ubuntu+check+package+version+command+line>`_

       `package management - How do I get the version of an application from the command line? - Ask Ubuntu <http://askubuntu.com/a/441005>`_

.. [2] `makefile run command - Google search <https://www.google.com/search?q=makefile+run+command>`_

       `osx - How to use shell commands in Makefile - Stack Overflow <http://stackoverflow.com/questions/10024279/how-to-use-shell-commands-in-makefile>`_

       `linux - How to use echo in shell command from Makefile? - Super User <http://superuser.com/questions/945148/how-to-use-echo-in-shell-command-from-makefile>`_

.. [3] `makefile strip spaces - Google search <https://www.google.com/search?q=makefile+strip+spaces>`_

       `GNU make: Text Functions <https://www.gnu.org/software/make/manual/html_node/Text-Functions.html>`_


.. _Makefile: https://www.google.com/search?q=Makefile
.. _strip text function: https://www.gnu.org/software/make/manual/html_node/Text-Functions.html
.. _Ubuntu Linux 16.10: http://releases.ubuntu.com/16.10/
.. _GNU make 4.1-9: https://www.gnu.org/software/make/

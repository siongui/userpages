[Makefile] Create Directory If Not Exist
########################################

:date: 2016-01-30T02:02+08:00
:tags: Bash, Makefile, Commandline
:category: Bash
:summary: Makefile - Create a directory if it does not exist. Otherwise do
          nothing.
:adsu: yes


Write a command line in Makefile: Create a directory if it does not exist.
Otherwise do nothing.

.. code-block:: bash

  MYDIR=/home/myaccount/dev
  [ -d $(MYDIR) ] || mkdir -p $(MYDIR)


Tested on ``Ubuntu Linux 15.10``.

.. adsu:: 2

[Bash] Remove Execute Permission of Files in Directory Recursively
##################################################################

:date: 2016-05-09T21:57+08:00
:tags: Bash, Commandline, find command, List Files in Directory, chmod command
:category: Bash
:summary: Remove execute(x) permission of files in directory recursively via
          chmod_ command and Bash_ script.
:adsu: yes


Remove execute(x) permission of files in directory recursively via chmod_
command and Bash_ script.

.. code-block:: sh

  #!/bin/bash

  # $1 is the directory in which files to be processed
  for path in $(find $1 -type f)
  do
    chmod -x ${path}
  done

----

.. adsu:: 2

References:

.. [1] `[Bash] Process HTML Files Format in Volumn <{filename}../../04/26/bash-process-html-files-format-in-volumn%en.rst>`_


.. _Bash: https://www.google.com/search?q=Bash
.. _chmod: https://www.google.com/search?q=chmod

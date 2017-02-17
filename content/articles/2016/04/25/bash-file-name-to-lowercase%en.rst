[Bash] Rename Files in Directory to Lowercase
#############################################

:date: 2016-04-25T14:38+08:00
:tags: Bash, Commandline, String Manipulation, find command, tr command,
       List Files in Directory
:category: Bash
:summary: Convert the name of files in directory to lowercase via Bash_ script.
:adsu: yes


Convert the name of files in directory to lowercase via Bash_ script.

.. code-block:: sh

  #!/bin/bash

  # $1 is the directory in which files to be renamed to lowercase
  for path in $(find $1 -type f)
  do
    dir=$(dirname ${path})
    file=$(basename ${path} | tr "[:upper:]" "[:lower:]")
    mv ${path} ${dir}/${file}
  done

.. adsu:: 2

----

References:

.. [1] `[Bash] List All Files in Directory Recursively and Rename <{filename}../../../2015/02/02/bash-list-files-recursively-and-rename%en.rst>`_

.. [2] | `bash to lowercase - Google search <https://www.google.com/search?q=bash+to+lowercase>`_
       | `Converting string to lower case in Bash shell scripting - Stack Overflow <http://stackoverflow.com/questions/2264428/converting-string-to-lower-case-in-bash-shell-scripting>`_

.. [3] | `bash basename dirname - Google search <https://www.google.com/search?q=bash+basename+dirname>`_
       | `linux - Bash - get last dirname/filename in a file path argument - Stack Overflow <http://stackoverflow.com/questions/3294072/bash-get-last-dirname-filename-in-a-file-path-argument>`_
.. adsu:: 3
.. [4] | `bash path join - Google search <https://www.google.com/search?q=bash+path+join>`_
       | `linux - How to concatenate two strings to build a complete path - Stack Overflow <http://stackoverflow.com/questions/11226322/how-to-concatenate-two-strings-to-build-a-complete-path>`_

.. _Bash: https://www.google.com/search?q=Bash

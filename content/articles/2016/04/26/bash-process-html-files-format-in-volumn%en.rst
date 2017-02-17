[Bash] Process HTML Files Format in Volumn
##########################################

:date: 2016-04-26T21:26+08:00
:tags: Bash, Commandline, String Manipulation, find command, tr command,
       List Files in Directory, sed, iconv command, remove trailing newline,
       remove carriage return
:category: Bash
:summary: Convert the format (Big5_ encoding to UTF-8_, remove `DOS newline`_ in
          file, replace string *big5* with *UTF-8*, and append `UNIX newline`_
          to end of file) of HTML files in directory via Bash_ script.
:adsu: yes


Convert the format (Big5_ encoding to UTF-8_, remove `DOS newline`_ in file,
replace string *big5* with *UTF-8*, and append `UNIX newline`_ to end of file)
of HTML files in directory via Bash_ script.

.. code-block:: sh

  #!/bin/bash

  # $1 is the directory in which files to be processed
  for path in $(find $1 -type f)
  do
    echo -e "\033[92mProcessing ${path} ...\033[0m"
    # big5 to utf8
    iconv -f big5 -t utf-8 ${path} > tmp.html
    if [ $? -ne 0 ]; then
      # fail to convert big5 to UTF-8
      continue
    fi
    mv tmp.html ${path}

    # remove dos newline
    tr -d '\015' <${path} > tmp.html
    mv tmp.html ${path}

    # html meta big5 to UTF-8
    sed 's/big5/UTF-8/' -i ${path}

    # append newline to end of file
    sed -i -e '$a\' ${path}
  done

.. adsu:: 2

----

References:

.. [1] `[Bash] Rename Files in Directory to Lowercase <{filename}../25/bash-file-name-to-lowercase%en.rst>`_

.. [2] `[Bash] Encoding Conversion, Newline Manipulation, String Replacement of File <{filename}../20/bash-file-encoding-conversion-newline-manipulation-string-replacement%en.rst>`_

.. [3] | `bash command success - Google search <https://www.google.com/search?q=bash+command+success>`_
       | `How to check if a command succeeded? - Ask Ubuntu <http://askubuntu.com/questions/29370/how-to-check-if-a-command-succeeded>`_

.. [4] | `bash for loop continue - Google search <https://www.google.com/search?q=bash+for+loop+continue>`_
       | `Break and continue <http://tldp.org/LDP/Bash-Beginners-Guide/html/sect_09_05.html>`_


.. _Bash: https://www.google.com/search?q=Bash
.. _Big5: https://en.wikipedia.org/wiki/Big5
.. _UTF-8: https://en.wikipedia.org/wiki/UTF-8
.. _DOS newline: https://en.wikipedia.org/wiki/Newline
.. _UNIX newline: https://en.wikipedia.org/wiki/Newline

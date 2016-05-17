[Bash] HTML to PDF via wkhtmltopdf
##################################

:date: 2016-05-17T23:12+08:00
:tags: Bash, Commandline, find command, List Files in Directory, wkhtmltopdf
:category: Bash
:summary: Convert HTML files in directory to PDF recursively via wkhtmltopdf_
          and Bash_ script.


Convert HTML files in directory to PDF recursively via wkhtmltopdf_ and Bash_
script. Please `install wkhtmltopdf`_ first.

.. code-block:: sh

  #!/bin/bash

  # $1 is the directory in which files to be processed
  # $2 is the css file
  for path in $(find $1 -name "*.html")
  do
    echo -e "\033[92mProcessing ${path}\033[0m"
    wkhtmltopdf ${path} --disable-javascript --user-style-sheet $2 "${path%.html}.pdf"
  done

If you get error message ``wkhtmltopdf: cannot connect to X server``, see
`this question on Stack Overflow`_ [3]_ for solution.

----

Tested on: ``Ubuntu Linux 16.04``, ``wkhtmltopdf 0.12.2.4``

----

References:

.. [1] `html to pdf command line linux - Google search <https://www.google.com/search?q=html+to+pdf+command+line+linux>`_

       `How to convert a HTML file to PDF (with colors) - Ask Ubuntu <http://askubuntu.com/questions/320195/how-to-convert-a-html-file-to-pdf-with-colors>`_

.. [2] `bash file extension change - Google search <https://www.google.com/search?q=bash+file+extension+change>`_

       `bash - How do I rename the extension for a batch of files? - Stack Overflow <http://stackoverflow.com/questions/1224766/how-do-i-rename-the-extension-for-a-batch-of-files>`_

.. [3] `xserver - wkhtmltopdf: cannot connect to X server - Stack Overflow <http://stackoverflow.com/questions/9604625/wkhtmltopdf-cannot-connect-to-x-server>`_


.. _Bash: https://www.google.com/search?q=Bash
.. _wkhtmltopdf: http://wkhtmltopdf.org/
.. _this question on Stack Overflow: http://stackoverflow.com/questions/9604625/wkhtmltopdf-cannot-connect-to-x-server
.. _install wkhtmltopdf: https://www.google.com/search?q=install+wkhtmltopdf

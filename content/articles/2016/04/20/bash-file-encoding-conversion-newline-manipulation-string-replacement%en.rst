[Bash] Encoding Conversion, Newline Manipulation, String Replacement of File
############################################################################

:date: 2016-04-20T04:45+08:00
:tags: Commandline, String Manipulation, Bash, iconv command, sed, tr command,
       remove trailing newline, remove carriage return
:category: Bash
:summary: Convert encoding of file from Big5_ to UTF-8_,  remove `DOS newline`_
          in file, replace string `big5` with `UTF-8`, and append
          `UNIX newline`_ to end of file.
:adsu: yes


Convert encoding of file from Big5_ to UTF-8_ [4]_,  remove `DOS newline`_ in
file [2]_, replace string `big5` with `UTF-8` via sed_ command, and append
`UNIX newline`_ to end of file. [3]_

.. code-block:: sh

  #!/bin/bash

  echo "$1: big5 to utf-8"
  iconv -f big5 -t utf-8 $1 > tmp.html
  mv tmp.html $1

  echo "$1: remove dos newline"
  tr -d '\015' <$1 > tmp.html
  mv tmp.html $1

  echo "$1: html meta big5 to UTF-8"
  sed 's/big5/UTF-8/' -i $1

  echo "$1: append newline to end of file"
  sed -i -e '$a\' $1


----

References:

.. [1] `bash read arguments - Google search <https://www.google.com/search?q=bash+read+arguments>`_

.. [2] `bash dos to unix - Google search <https://www.google.com/search?q=bash+dos+to+unix>`_

       `linux - How to convert DOS/Windows newline (CRLF) to Unix newline (\n) in bash script? - Stack Overflow <http://stackoverflow.com/questions/2613800/how-to-convert-dos-windows-newline-crlf-to-unix-newline-n-in-bash-script>`_

.. [3] `bash append newline to end of file - Google search <https://www.google.com/search?q=bash+append+newline+to+end+of+file>`_

       `bash - How to add a newline to the end of a file? - Unix & Linux Stack Exchange <http://unix.stackexchange.com/questions/31947/how-to-add-a-newline-to-the-end-of-a-file>`_

.. [4] `[Bash] Convert Files in Directory From Big5 to UTF-8 <{filename}../../03/18/bash-convert-files-in-directory-from-big5-to-utf8%en.rst>`_


.. _Big5: https://en.wikipedia.org/wiki/Big5
.. _UTF-8: https://en.wikipedia.org/wiki/UTF-8
.. _iconv: http://linux.die.net/man/1/iconv
.. _sed: http://www.grymoire.com/Unix/Sed.html
.. _DOS newline: https://en.wikipedia.org/wiki/Newline
.. _UNIX newline: https://en.wikipedia.org/wiki/Newline

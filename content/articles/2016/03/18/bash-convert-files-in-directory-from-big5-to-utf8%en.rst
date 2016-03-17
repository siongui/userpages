[Bash] Convert Files in Directory From Big5 to UTF-8
####################################################

:tags: find command, Commandline, String Manipulation, List Files in Directory,
       Bash, iconv command
:category: Bash
:summary: Convert all files in a directory recursively (i.e., including
          subdirectories of the directory) from Big5_ encoding to UTF-8_ via
          iconv_ command.


Question:

  A lot of HTML files under *dhpstory* directory. The encoding of the files is
  Big5_. Converted them to UTF-8_ encoding via iconv_ command.

Answer:

.. show_github_file:: siongui userpages content/code/bash-big5-to-utf8/convert.sh


----

References:

.. [1] `[Bash] List All Files in Directory Recursively and Rename <{filename}../../../2015/02/02/bash-list-files-recursively-and-rename%en.rst>`_

.. _Big5: https://en.wikipedia.org/wiki/Big5
.. _UTF-8: https://en.wikipedia.org/wiki/UTF-8
.. _iconv: http://linux.die.net/man/1/iconv

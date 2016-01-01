[Bash] List All Files in Directory Recursively and Rename
#########################################################

:tags: sed, find, Bash, Commandline, String Manipulation
:category: Bash
:summary: List all files in a directory recursively (i.e., including
          subdirectories of the directory), and use sed to rename the files.


Question:

  I have a lot of files under *content/articles* directory. The filenames are something like
  ``abc-def-ghi#en.rst``. I want to rename them all and replace the ``#`` with
  ``%``. For example, ``abc-def-ghi#en.rst`` will be renamed as
  ``abc-def-ghi%en.rst``, how do I do it?

Answer:

.. show_github_file:: siongui userpages content/articles/2015/02/02/list-files-recursively-rename.sh


----

References:

.. [1] `List all files in a directory recursively but exclude directories themselves <http://unix.stackexchange.com/questions/76855/list-all-files-in-a-directory-recursively-but-exclude-directories-themselves>`_

.. [2] `Using sed to mass rename files <http://stackoverflow.com/questions/2372719/using-sed-to-mass-rename-files>`_

.. [3] `How to increment a variable in bash? <http://askubuntu.com/questions/385528/how-to-increment-a-variable-in-bash>`_

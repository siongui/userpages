Find Image With Specific Dimension Using Command Line
#####################################################

:date: 2021-12-19T21:30+08:00
:tags: sed, find command, Bash, Commandline, String Manipulation, ImageMagick,
       grep command
:category: Bash
:summary: Find images with specific dimension using Linux command line.
:og_image: https://cdn.pixabay.com/photo/2013/07/12/14/48/bash-148836_1280.png
:adsu: yes


**Question**:

  Find JPEG images with size 1080x1920 using Linux command line tool.

**Answer**:

Install ImageMagick_:

.. code-block:: bash

  $ sudo apt-get install imagemagick

Use the following command to find all ``*.jpg`` files with size 1080x1920: [1]_

.. code-block:: bash

  $ find . -iname "*.jpg" -type f -exec identify -format '%i %wx%h\n' '{}' \; | grep '1080x1920'


To move the 1080x1920 ``*.jpg`` files to another directory, run: [2]_

.. code-block:: bash

  $ find . -iname "*.jpg" -type f -exec identify -format '%i %wx%h\n' '{}' \; | grep '1080x1920' | sed -e "s/ 1080x1920$//" | xargs -I {} mv {} /path/to/destination/directory/

----

References:

.. [1] `How to find all images with a certain pixel size using command line? - Ask Ubuntu <https://askubuntu.com/questions/238136/how-to-find-all-images-with-a-certain-pixel-size-using-command-line>`_
.. [2] `Remove a fixed prefix/suffix from a string in Bash - Stack Overflow <https://stackoverflow.com/questions/16623835/remove-a-fixed-prefix-suffix-from-a-string-in-bash>`_

.. _ImageMagick: https://imagemagick.org/
.. _sed: https://www.google.com/search?q=sed

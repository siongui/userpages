Update Modification Time According to EXIF Data
###############################################

:date: 2021-12-31T00:28+08:00
:modified: 2022-01-02T22:51+08:00
:tags: find command, Bash, Commandline, EXIF
:category: Bash
:summary: Update modification time of photo/video according to EXIF data using
          command line.
:og_image: https://upload.wikimedia.org/wikipedia/commons/5/59/Exif_Metadata_logo.svg
:adsu: yes


**Question**:

  Find photo/video and update modification time (mtime) of the photo/video
  according to EXIF data.

**Answer**:

Install *exiftool*:

.. code-block:: bash

  $ sudo apt-get install libimage-exiftool-perl

Find photo/video. The ``-i`` in ``-iregex`` means case insensitive.

.. code-block:: bash

  $ find . -type f -iregex '.*\.\(jpg\|gif\|png\|jpeg\|mov\|mp4\|heic\)$'

Update modification time according to EXIF data.

.. code-block:: bash

  $ find . -type f -iregex '.*\.\(jpg\|gif\|png\|jpeg\|mov\|mp4\|heic\)$' | xargs -I {} exiftool "-DateTimeOriginal>FileModifyDate" {}

----

Instead of *exiftool*, you can also use *jhead* to update mtime.

.. code-block:: bash

  $ sudo apt-get install jhead

Use *jhead* along with *find* command:

.. code-block:: bash

  $ find . -type f -iregex '.*\.\(jpg\|gif\|png\|jpeg\|mov\|mp4\|heic\)$' -exec jhead -ft {} +

----

References:

.. [1] `images - Change file created date from JPEG EXIF metadata - Unix & Linux Stack Exchange <https://unix.stackexchange.com/questions/89264/change-file-created-date-from-jpeg-exif-metadata>`_
.. [2] `regular expression - How to use find command to search for multiple extensions - Unix & Linux Stack Exchange <https://unix.stackexchange.com/questions/15308/how-to-use-find-command-to-search-for-multiple-extensions>`_
.. [3] `Is there a free program to (batch) change photo file's date to match EXIF? - Photography Stack Exchange <https://photo.stackexchange.com/questions/27245/is-there-a-free-program-to-batch-change-photo-files-date-to-match-exif>`_
.. [4] `timestamp - How can I change the date modified/created of a file? - Ask Ubuntu <https://askubuntu.com/questions/62492/how-can-i-change-the-date-modified-created-of-a-file>`_

[Bash] Convert wav to mp3 via ffmpeg
####################################

:date: 2016-02-22T07:07+08:00
:tags: Bash, Commandline, Ubuntu Linux, apt command
:category: Bash
:summary: Write a bash_ script to convert wav to mp3 via ffmpeg_ on Ubuntu Linux
          15.10.
:adsu: yes


Question
++++++++

  Convert ``9dt/9dtXXX.wav`` to ``output/9dtXXX.mp3``, where XXX are numbers
  from 000 to 094.

Answer
++++++

Install ffmpeg_:

.. code-block:: bash

  # install ffmpeg on Ubuntu Linux 15.10
  $ sudo apt-get install ffmpeg

The bash_ script to convert wav to mp3 via ffmpeg_:

.. code-block:: bash

  #!/bin/bash

  OUTPUTDIR="output"

  mkdir -p $OUTPUTDIR
  for i in {000..094}
  do
    ffmpeg -i 9dt/9dt$i.wav $OUTPUTDIR/9dt$i.mp3
  done

Save the above bash_ script as *wav2mp3.sh*

Run the script by:

.. code-block:: bash

  $ bash wav2mp3.sh

----

Tested on ``Ubuntu Linux 15.10``

----

References:

.. [1] Google search: `linux convert wav to mp3 <https://www.google.com/search?q=linux+convert+wav+to+mp3>`_

.. [2] `Converting WAV Files to High Quality MP3 Using Linux terminal - Stack Overflow <http://stackoverflow.com/questions/11216445/converting-wav-files-to-high-quality-mp3-using-linux-terminal>`_

.. [3] `Why would I choose Libav over FFmpeg, or is there even a difference? - Super User <http://superuser.com/questions/507386/why-would-i-choose-libav-over-ffmpeg-or-is-there-even-a-difference>`_

.. [4] `software installation - Is FFmpeg missing from the official repositories in 14.04? - Ask Ubuntu <http://askubuntu.com/questions/432542/is-ffmpeg-missing-from-the-official-repositories-in-14-04>`_

.. _bash: https://www.google.com/search?q=bash
.. _ffmpeg: https://www.google.com/search?q=ffmpeg

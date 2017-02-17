[Bash] Move Large Number of Files via tar Command
#################################################

:date: 2016-04-29T23:20+08:00
:tags: Bash, Commandline, tar command
:category: Bash
:summary: Move large number of files via tar_ command under Bash_. This is much
          quicker than mv_ command if there are a lot of small files.
:adsu: yes


Move large number of files via tar_ command under Bash_. This is much quicker
than mv_ command if there are a lot of small files.

.. code-block:: sh

  $ tar cvf foo.tar /path/to/your_directory
  $ tar xvf foo.tar -C /path/to/your_new_location
  $ rm foo.tar

----

.. adsu:: 2

References:

.. [1] | `tar directory without compression - Google search <https://www.google.com/search?q=tar+directory+without+compression>`_
       | `linux - How can I use tar command to group files without compression? - Super User <http://superuser.com/questions/529926/how-can-i-use-tar-command-to-group-files-without-compression>`_

.. [2] | `untar tar file - Google search <https://www.google.com/search?q=untar+tar+file>`_
       | `HowTo : Extract (untar) [tar], [tar.gz] and [tar.bz2] Files | Linux Commands <http://www.shellhacks.com/en/HowTo-Extract-untar-tar-targz-and-tarbz2-Files>`_


.. _Bash: https://www.google.com/search?q=Bash
.. _tar: http://linux.die.net/man/1/tar
.. _mv: http://linux.die.net/man/1/mv

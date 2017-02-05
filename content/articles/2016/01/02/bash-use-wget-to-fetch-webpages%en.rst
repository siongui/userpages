[Bash] Use wget to Fetch Webpages
#################################

:date: 2016-01-02T03:41+08:00
:tags: Bash, Commandline, String Manipulation
:category: Bash
:summary: Write a bash script to use wget to fetch webpages.
:adsu: yes


Question
++++++++

  Fetch webpages ``http://www.aia.or.th/prayerXX.htm``, where XX are numbers
  from 00 to 99.

Answer
++++++

.. show_github_file:: siongui userpages content/code/bash/wget/wget-webpages.sh
.. adsu:: 2

GNU wget_ is used to fetch the webpages. I tried to run the script:

.. code-block:: bash

  $ sh wget-webpages.sh

but did not work. Need to run with the following command:

.. code-block:: bash

  $ bash wget-webpages.sh

My development environment is *Ubuntu Linux 15.10*.

.. adsu:: 3

----

References:

.. [1] Google search: `bash print number with leading zeros <https://www.google.com/search?q=bash+print+number+with+leading+zeros>`_

.. [2] `GNU Wget Manual <https://www.gnu.org/software/wget/manual/wget.html>`_
       (or ``$ man wget`` in command line in Linux console)

.. [3] `sleep(3): sleep for specified number of seconds - Linux man page <http://linux.die.net/man/3/sleep>`_
       (or ``$ man sleep`` in command line in Linux console)


.. _wget: https://www.gnu.org/software/wget/

[Python] mkdir -p
#################

:date: 2016-02-21T23:22+08:00
:tags: Python, Bash, Commandline
:category: Python
:summary: `mkdir -p`_ command in Python_
:adsu: yes


`mkdir -p`_ command in Python_

.. code-block:: python

  import os

  def mkdirp(dirpath):
      if not os.path.exists(dirpath):
          os.makedirs(dirpath)


----

Tested on: ``Ubuntu Linux 15.10``, ``Python 2.7.10``.

----

References:

.. [1] `python mkdir p <https://www.google.com/search?q=python+mkdir+p>`_

.. [2] `mkdir -p functionality in Python - Stack Overflow <http://stackoverflow.com/questions/600268/mkdir-p-functionality-in-python>`_

.. _Python: https://www.python.org/
.. _mkdir -p: http://linux.die.net/man/1/mkdir

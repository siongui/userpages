[Makefile] Echo Color Output
############################

:date: 2016-02-18T12:07+08:00
:tags: Bash, Makefile, Commandline
:category: Bash
:summary: echo_ colorful output in Makefile_.
:adsu: yes


echo_ colorful output in Makefile_. I tried several methods but they did not
work. Then I found [1]_, which works on my ``Ubuntu Linux 15.10``.

.. code-block:: bash

  @echo "\033[92mHello World\033[0m"

See [1]_ for other text color.

----

Tested on ``Ubuntu Linux 15.10``.

----

.. [1] `bash - How to change the output color of echo in Linux - Stack Overflow <http://stackoverflow.com/a/5947779>`_

.. [2] `colorful Makefile echo output · siongui/pali@587672e · GitHub <https://github.com/siongui/pali/commit/587672eb8729112b926d197550d3b1a0e0fb4448>`_

.. _echo: https://www.google.com/search?q=linux+echo
.. _Makefile: https://www.google.com/search?q=Makefile

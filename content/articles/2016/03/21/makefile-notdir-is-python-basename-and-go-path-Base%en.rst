Makefile notdir = Python basename / Golang path.Base
####################################################

:date: 2016-03-21T07:47+08:00
:tags: Bash, Makefile, Commandline, String Manipulation, Go, Golang, Python
:category: Bash
:summary: Makefile_'s notdir_ = `Python basename`_ / `Go path.Base`_.
:adsu: yes

Makefile_'s notdir_ = `Python basename`_ / `Go path.Base`_.

Example:

.. code-block:: bash

  URL=http://nanda.online-dhamma.net/Tipitaka/Sutta/Khuddaka/Dhammapada/DhP_Story016.htm
  UTF8HTML=/tmp/utf8-$(notdir ${URL})

  default:
  	@go run big5ToUTF8.go -input=$(notdir ${URL}) -output=${UTF8HTML}


In above Makefile_,

| **$(notdir ${URL})**      =      *DhP_Story016.htm*
| **UTF8HTML**              =      */tmp/utf8-DhP_Story016.htm*.

.. adsu:: 2

----

Tested on ``Ubuntu Linux 15.10``, `GNU make 4.0-8.2`_.

.. _Makefile: https://www.google.com/search?q=Makefile
.. _notdir: https://www.gnu.org/software/make/manual/html_node/File-Name-Functions.html
.. _Python basename: https://docs.python.org/2/library/os.path.html#os.path.basename
.. _Go path.Base: https://golang.org/pkg/path/#Base
.. _Python: https://www.python.org/
.. _GNU make 4.0-8.2: http://packages.ubuntu.com/wily/make

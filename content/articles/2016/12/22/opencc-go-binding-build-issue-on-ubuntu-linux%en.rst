Build Issue of OpenCC Go binding on Ubuntu Linux 16.10
######################################################

:date: 2016-12-22T23:07+08:00
:tags: Go, Golang, Conversion of Traditional and Simplified Chinese, OpenCC,
       apt command
:category: Go
:summary: Fail to build OpenCC_ Go_ binding on `Ubuntu Linux 16.10`_, and how I
          fix it.
:adsu: yes


When I upgrade Ubuntu Linux from 16.04 to 16.10, the OpenCC_ version jump from
0.4.3 to 1.0.4. It is ok to install OpenCC:

.. code-block:: bash

  $ sudo apt-get update
  $ apt-cache search opencc
  libopencc-dbg - simplified-traditional chinese conversion library - debug
  libopencc-dev - simplified-traditional Chinese conversion library - development
  libopencc2 - simplified-traditional Chinese conversion library - runtime
  libopencc2-data - simplified-traditional Chinese conversion library - data files
  opencc - simplified-traditional Chinese conversion tool
  $ sudo apt-get install opencc

Then I tried to install ``libopencc-dev`` to compile *OpenCC Go binding/wrapper*
[2]_, I got the following error message:

.. code-block:: bash

  $ sudo apt-get install libopencc-dev
  ...
  The following packages have unmet dependencies:
   libopencc-dev : Depends: libopencc2 (= 1.0.4-1ubuntu0.16.10.1) but it is not going to be installed

After some googling, I manually download and install ``libopencc-dev``:

.. code-block:: bash

  $ wget https://launchpad.net/ubuntu/+archive/primary/+files/libopencc-dev_1.0.4-1ubuntu0.16.10.1_amd64.deb
  $ sudo dpkg -i libopencc-dev_1.0.4-1ubuntu0.16.10.1_amd64.deb

And finally I successfully compile the *OpenCC Go binding/wrapper*:

.. code-block:: bash

  $ go get -u github.com/stevenyao/go-opencc

But after you finish to use the *OpenCC Go binding/wrapper*, remove
``libopencc-dev`` manually:

.. code-block:: bash

  $ sudo dpkg -r libopencc-dev

If you keep ``libopencc-dev``, you cannot use ``apt-get`` to update and upgrade
your system.

----

Tested on: ``Ubuntu Linux 16.10``, ``Go 1.7.4``, ``opencc 1.0.4-1ubuntu0.16.10.1``.

----

References:

.. [1] `開放中文轉換 Open Chinese Convert (OpenCC) <http://opencc.byvoid.com/>`_
       (`source code <https://github.com/BYVoid/OpenCC>`__,
       `online doc <http://byvoid.github.io/OpenCC/>`__)

.. [2] `stevenyao/go-opencc · GitHub <https://github.com/stevenyao/go-opencc>`_
       (OpenCC wrapper for Golang, |godoc1|)

.. [3] `[Golang] Conversion of Traditional and Simplified Chinese <{filename}../../01/03/go-conversion-of-traditional-and-simplified-chinese%en.rst>`_

.. [4] `opencc package : Ubuntu <https://launchpad.net/ubuntu/+source/opencc>`_

.. [5] `GitHub - sakilu/go-opencc-crawler: golang opencc <https://github.com/sakilu/go-opencc-crawler>`_

.. [6] `command line - Update the system from terminal - Ask Ubuntu <http://askubuntu.com/questions/462449/update-the-system-from-terminal>`_

       `dpkg - How do I get a list of installed files from a package? - Ask Ubuntu <http://askubuntu.com/questions/32507/how-do-i-get-a-list-of-installed-files-from-a-package>`_

       `Linux Halwa: How to force install .deb package <http://linuxhalwa.blogspot.com/2013/12/how-to-force-install-deb-package.html>`_

       `How to uninstall a .deb installed with dpkg? - Unix & Linux Stack Exchange <http://unix.stackexchange.com/questions/195794/how-to-uninstall-a-deb-installed-with-dpkg>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _OpenCC: http://opencc.byvoid.com/
.. _OpenCC repository on GitHub: https://github.com/BYVoid/OpenCC
.. _Ubuntu Linux 16.10: http://releases.ubuntu.com/16.10/

.. |godoc1| image:: https://godoc.org/github.com/stevenyao/go-opencc?status.png
   :target: https://godoc.org/github.com/stevenyao/go-opencc

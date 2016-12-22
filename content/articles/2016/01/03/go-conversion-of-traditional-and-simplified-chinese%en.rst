[Golang] Conversion of Traditional and Simplified Chinese
#########################################################

:date: 2016-01-03T01:22+08:00
:tags: Go, Golang, String Manipulation, OpenCC,
       Conversion of Traditional and Simplified Chinese
:category: Go
:summary: Conversion of Traditional and Simplified Chinese by OpenCC_ and Go_
          programming language.


OpenCC_ is a tool (both online and offline) for conversion Traditional and
Simplified Chinese. In this post, we will write a Go_ program to use OpenCC_ to
convert Simplified Chinese to Traditional Chinese.


Install OpenCC_
+++++++++++++++

See `OpenCC repository on GitHub`_ for installation. If you use Ubuntu Linux
15.10, you can search and install OpenCC_ by following command:

.. code-block:: bash

  $ apt-cache search opencc
  fcitx-modules - Flexible Input Method Framework - core modules
  libopencc-dbg - simplified-traditional chinese conversion library - debug
  libopencc-dev - simplified-traditional chinese conversion library - development
  libopencc1 - simplified-traditional chinese conversion library - runtime
  opencc - simplified-traditional chinese conversion tool
  python-opencc - simplified-traditional chinese conversion library - Python support
  $ sudo apt-get install opencc libopencc-dev


OpenCC_ wrapper for Golang_
+++++++++++++++++++++++++++

After some googling [2]_, I found an OpenCC_ wrapper [3]_ for Go_. I tried this
wrapper but it did not work. So I forked it and made some modifications to make
it work on my system. Assume Go_ is already install in your system, install the
my modified wrapper [4]_ by:

.. code-block:: bash

  $ go get github.com/siongui/go-opencc

Another problem I had is that README said the configurations are .json files. I
run ``locate opencc`` command:

.. code-block:: bash

  $ locate opencc
  /usr/lib/x86_64-linux-gnu/libopencc.so.1
  /usr/lib/x86_64-linux-gnu/libopencc.so.1.0.0
  /usr/lib/x86_64-linux-gnu/opencc
  /usr/lib/x86_64-linux-gnu/opencc/from_tw_phrases.txt
  /usr/lib/x86_64-linux-gnu/opencc/from_tw_variants.txt
  /usr/lib/x86_64-linux-gnu/opencc/mix2zhs.ini
  /usr/lib/x86_64-linux-gnu/opencc/mix2zht.ini
  /usr/lib/x86_64-linux-gnu/opencc/simp_to_trad_characters.ocd
  /usr/lib/x86_64-linux-gnu/opencc/simp_to_trad_phrases.ocd
  /usr/lib/x86_64-linux-gnu/opencc/to_cn_phrases.txt
  /usr/lib/x86_64-linux-gnu/opencc/to_tw_phrases.txt
  /usr/lib/x86_64-linux-gnu/opencc/to_tw_variants.txt
  /usr/lib/x86_64-linux-gnu/opencc/trad_to_simp_characters.ocd
  /usr/lib/x86_64-linux-gnu/opencc/trad_to_simp_phrases.ocd
  /usr/lib/x86_64-linux-gnu/opencc/zhs2zht.ini
  /usr/lib/x86_64-linux-gnu/opencc/zhs2zhtw_p.ini
  /usr/lib/x86_64-linux-gnu/opencc/zhs2zhtw_v.ini
  /usr/lib/x86_64-linux-gnu/opencc/zhs2zhtw_vp.ini
  /usr/lib/x86_64-linux-gnu/opencc/zht2zhs.ini
  /usr/lib/x86_64-linux-gnu/opencc/zht2zhtw_p.ini
  /usr/lib/x86_64-linux-gnu/opencc/zht2zhtw_v.ini
  /usr/lib/x86_64-linux-gnu/opencc/zht2zhtw_vp.ini
  /usr/lib/x86_64-linux-gnu/opencc/zhtw2zhcn_s.ini
  /usr/lib/x86_64-linux-gnu/opencc/zhtw2zhcn_t.ini
  /usr/lib/x86_64-linux-gnu/opencc/zhtw2zhs.ini
  /usr/lib/x86_64-linux-gnu/opencc/zhtw2zht.ini
  /usr/share/doc/libopencc1
  /usr/share/doc/libopencc1/changelog.Debian.gz
  /usr/share/doc/libopencc1/copyright
  /var/lib/dpkg/info/libopencc1:amd64.list
  /var/lib/dpkg/info/libopencc1:amd64.md5sums
  /var/lib/dpkg/info/libopencc1:amd64.postinst
  /var/lib/dpkg/info/libopencc1:amd64.postrm
  /var/lib/dpkg/info/libopencc1:amd64.shlibs
  /var/lib/dpkg/info/libopencc1:amd64.symbols

I saw no .json files, but saw a lot of .ini files. I used these .ini files as
configurations and it worked. I guess that maybe at some moment the author of
OpenCC_ changed the name of configurations.


Souce Code
++++++++++

.. show_github_file:: siongui userpages content/code/go-opencc/zhCN2zhTW.go

You can replace ``zhs2zhtw_vp.ini`` with other configurations according to your
needs. All configurations I found by ``locate opencc`` are:

.. code-block:: txt

  mix2zhs.ini
  mix2zht.ini
  zhs2zht.ini
  zhs2zhtw_p.ini
  zhs2zhtw_v.ini
  zhs2zhtw_vp.ini
  zht2zhs.ini
  zht2zhtw_p.ini
  zht2zhtw_v.ini
  zht2zhtw_vp.ini
  zhtw2zhcn_s.ini
  zhtw2zhcn_t.ini
  zhtw2zhs.ini
  zhtw2zht.ini


Test
++++

.. show_github_file:: siongui userpages content/code/go-opencc/zhCN2zhTW_test.go

Output of Test
``````````````

.. code-block:: txt

  === RUN   TestCN2TW
  --- PASS: TestCN2TW (0.02s)
          zhCN2zhTW_test.go:6: 中國滑鼠軟體列印機
  PASS


Tested on: ``Ubuntu Linux 15.10``, ``Go 1.5.2``, ``opencc 0.4.3-2build1``.

----

References:

.. [1] `開放中文轉換 Open Chinese Convert (OpenCC) <http://opencc.byvoid.com/>`_
       (`source code <https://github.com/BYVoid/OpenCC>`__)

.. [2] Google Search: `golang opencc <https://www.google.com/search?q=golang+opencc>`_

.. [3] `stevenyao/go-opencc · GitHub <https://github.com/stevenyao/go-opencc>`_
       (OpenCC wrapper for Golang, |godoc1|)

.. [4] `siongui/go-opencc · GitHub <https://github.com/siongui/go-opencc>`_
       (my modified OpenCC wrapper for Golang, |godoc2|)

.. [5] `[JavaScript] Conversion of Traditional and Simplified Chinese <{filename}../../../2012/10/03/javascript-conversion-of-traditional-and-simplified-chinese%en.rst>`_

.. [6] `[Python] Conversion of Traditional and Simplified Chinese <{filename}../../../2016/01/04/python-conversion-of-traditional-and-simplified-chinese%en.rst>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _OpenCC: http://opencc.byvoid.com/
.. _OpenCC repository on GitHub: https://github.com/BYVoid/OpenCC

.. |godoc1| image:: https://godoc.org/github.com/stevenyao/go-opencc?status.png
   :target: https://godoc.org/github.com/stevenyao/go-opencc

.. |godoc2| image:: https://godoc.org/github.com/siongui/go-opencc?status.png
   :target: https://godoc.org/github.com/siongui/go-opencc

[Python] Conversion of Traditional and Simplified Chinese
#########################################################

:date: 2016-01-04T00:49+08:00
:tags: Python, String Manipulation, OpenCC,
       Conversion of Traditional and Simplified Chinese
:category: Python
:summary: Conversion of Traditional and Simplified Chinese by OpenCC_ and
          Python_ programming language.


OpenCC_ is a tool (both online and offline) for conversion Traditional and
Simplified Chinese. In this post, we will write a Python_ program to use OpenCC_
to convert Simplified Chinese to Traditional Chinese.


Install OpenCC_
+++++++++++++++

See `OpenCC repository on GitHub`_ for installation. If you use Ubuntu Linux
15.10, you can install OpenCC_ by:

.. code-block:: bash

  $ sudo apt-get install opencc


OpenCC_ binding for Python_
+++++++++++++++++++++++++++

I found a lot of OpenCC_ bindings for Python_ ([3]_, [5]_, [6]_) We will use
pyOpenCC_ in [3]_ to convert Chinese. If you use Ubuntu Linux 15.10, you can
install `pyOpenCC <https://pypi.python.org/pypi/pyopencc>`__ by:

.. code-block:: bash

  # Install necessary header files for compilation
  $ sudo apt-get install python-dev libopencc-dev
  # Install pip. Ignore this step if you already installed pip
  $ sudo apt-get install python-pip
  # Install pyOpenCC
  $ sudo pip install pyopencc


Souce Code For Demo
+++++++++++++++++++

.. show_github_file:: siongui userpages content/code/python-opencc/convert-chinese.py

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


Output of Demo
``````````````

.. code-block:: txt

  中國滑鼠軟體列印機


Tested on: ``Ubuntu Linux 15.10``, ``Python 2.7.10``, ``opencc 0.4.3-2build1``,
``pyopencc-0.4.2.2``.

----

References:

.. [1] `開放中文轉換 Open Chinese Convert (OpenCC) <http://opencc.byvoid.com/>`_
       (`source code <https://github.com/BYVoid/OpenCC>`__)

.. [2] Google Search: `python opencc <https://www.google.com/search?q=python+opencc>`_

.. [3] `pyOpenCC <https://github.com/cute/pyopencc>`_
       (`PyPI <https://pypi.python.org/pypi/pyopencc>`__,
       `OpenCC Python binding <http://liguangming.com/opencc-python-binding>`__)

.. [4] `python-jianfan - A python library for translation between traditional and simplified chinese - Google Project Hosting <https://code.google.com/p/python-jianfan/>`_
       (`mirror <https://github.com/siongui/python-jianfan>`__)

.. [5] `A Python wrapper for Open Chinese Convert <https://bitbucket.org/victorlin/opencc_python>`_
       (`PyPI <https://pypi.python.org/pypi/opencc-python/>`__)

.. [6] `A ctypes-based OpenCC converter for Chinese <https://github.com/lepture/opencc-python>`_
       (`PyPI <https://pypi.python.org/pypi/OpenCC>`__)

.. [7] `Python 繁簡轉換套件　OpenCC 安裝及使用方法 (Ubuntu) <http://danceintech.blogspot.com/2015/01/python-opencc-ubuntu.html>`_

.. [8] `使用opencc in python繁體（正體）簡體中文轉換 <http://sushiwens.blogspot.com/2012/07/opencc-in-python.html>`_

.. [9] `[Golang] Conversion of Traditional and Simplified Chinese <{filename}../../../2016/01/03/go-conversion-of-traditional-and-simplified-chinese%en.rst>`_

.. [10] `[JavaScript] Conversion of Traditional and Simplified Chinese <{filename}../../../2012/10/03/javascript-conversion-of-traditional-and-simplified-chinese%en.rst>`_

.. [11] Google Search: `Python.h: No such file or directory <https://www.google.com/search?q=Python.h%3A+No+such+file+or+directory>`_


.. _Python: https://www.python.org/
.. _pyOpenCC: https://github.com/cute/pyopencc
.. _OpenCC: http://opencc.byvoid.com/
.. _OpenCC repository on GitHub: https://github.com/BYVoid/OpenCC

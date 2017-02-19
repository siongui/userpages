[Golang] Converter for Traditional and Simplified Chinese
#########################################################

:date: 2017-02-19T16:53+08:00
:tags: Go, Golang, String Manipulation,
       Conversion of Traditional and Simplified Chinese
:category: Go
:summary: gojianfan_ - converter for Traditional and Simplified Chinese in Go_
          programming language.
:og_image: http://www.unixstickers.com/image/cache/data/stickers/golang/golang.sh-600x600.png
:adsu: yes


gojianfan_ is a converter implemented in Golang_ for conversion of Traditional
and Simplified Chinese. It is a very primitive, and based on `python-jianfan`_.
If you need a advanced converter, please visit OpenCC_ (`GitHub repo`_) [2]_.


Implementation of gojianfan_
++++++++++++++++++++++++++++

The Chinese writing language consists of large number of characters. Normally
speaking, one character in Traditional Chinese can be mapped to one character in
Simplified Chinese, and vice versa (although this statement may not be true in
some cases). The most basic form of converter is to construct mappings between
characters in Traditional and Simplified Chinese, and then do the conversion
character by character according to the mapping. This is how gojianfan_ is
implemented. The idea is simple, and the most difficult part is to find the
mapping of character sets. I found the mapping in the source code of
`python-jianfan`_, and write the converter in Go_ according to the mapping in
`python-jianfan`_. The Go_ implementation is only a few lines of code and easy
to understand. Please visit gojianfan_ if you want to know more details.

.. adsu:: 2


Usgae
+++++

First install the package by:

.. code-block:: bash

  $ go get -u github.com/siongui/gojianfan

Example:

.. code-block:: go

  package main

  import (
  	"fmt"
  	"github.com/siongui/gojianfan"
  )

  func main() {
  	// Traditional Chinese to Simplified Chinese
  	fmt.Println(gojianfan.T2S("橋頭"))

  	// Simplified Chinese to Traditional Chinese
  	fmt.Println(gojianfan.S2T("桥头"))
  }

.. adsu:: 3

----

Tested on: ``Ubuntu Linux 16.10``, ``Go 1.8``.

----

References:

.. [1] `GitHub - siongui/gojianfan: Traditional and Simplified Chinese Conversion in Go <https://github.com/siongui/gojianfan>`_
.. [2] `[Golang] Conversion of Traditional and Simplified Chinese <{filename}../../../2016/01/03/go-conversion-of-traditional-and-simplified-chinese%en.rst>`_
.. [3] `[JavaScript] Conversion of Traditional and Simplified Chinese <{filename}../../../2012/10/03/javascript-conversion-of-traditional-and-simplified-chinese%en.rst>`_
.. [4] `[Python] Conversion of Traditional and Simplified Chinese <{filename}../../../2016/01/04/python-conversion-of-traditional-and-simplified-chinese%en.rst>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _OpenCC: http://opencc.byvoid.com/
.. _GitHub repo: https://github.com/BYVoid/OpenCC
.. _gojianfan: https://github.com/siongui/gojianfan
.. _python-jianfan: https://code.google.com/archive/p/python-jianfan/

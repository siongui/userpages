[Golang] Convert Traditional Chinese PO file to Simplified Chinese
##################################################################

:date: 2016-03-05T21:14+08:00
:tags: Go, Golang, String Manipulation, i18n, Locale, gettext,
       Conversion of Traditional and Simplified Chinese, OpenCC
:category: Go
:summary: Automatically convert `Traditional Chinese`_ (`zh_TW`_) PO_ file to
          `Simplified Chinese`_ (`zh_CN`_) by OpenCC_ and Go_ programming
          language.

Automatically convert `Traditional Chinese`_ (`zh_TW`_) PO_ file to
`Simplified Chinese`_ (`zh_CN`_) by OpenCC_ and `OpenCC Go binding`_.


Install Necessary Tool
++++++++++++++++++++++

.. code-block:: bash

  $ sudo apt-get install opencc libopencc-dev
  $ go get -u github.com/siongui/go-opencc


Source Code
+++++++++++

.. code-block:: go

  package main

  import (
  	"bufio"
  	"fmt"
  	"github.com/siongui/go-opencc"
  	"os"
  	"path/filepath"
  	"strings"
  )

  var ct2s = opencc.NewConverter("zht2zhs.ini")

  func File2lines(filePath string) []string {
  	f, err := os.Open(filePath)
  	if err != nil {
  		panic(err)
  	}
  	defer f.Close()

  	var lines []string
  	scanner := bufio.NewScanner(f)
  	for scanner.Scan() {
  		lines = append(lines, scanner.Text())
  	}
  	if err := scanner.Err(); err != nil {
  		fmt.Fprintln(os.Stderr, err)
  		panic(err)
  	}

  	return lines
  }

  func main() {
  	twPOPath := "locale/zh_TW/LC_MESSAGES/messages.po"
  	cnPOPath := "locale/zh_CN/LC_MESSAGES/messages.po"
  	defer ct2s.Close()

  	os.MkdirAll(filepath.Dir(cnPOPath), 0755)
  	fo, err := os.Create(cnPOPath)
  	if err != nil {
  		fmt.Fprintln(os.Stderr, err)
  		panic(err)
  	}
  	defer fo.Close()

  	for _, line := range File2lines(twPOPath) {
  		if strings.HasPrefix(line, "msgstr") {
  			fo.Write([]byte(ct2s.Convert(line)))
  		} else {
  			if strings.Contains(line, "zh_TW") {
  				fo.Write([]byte(strings.Replace(line, "zh_TW", "zh_CN", 1)))
  			} else {
  				fo.Write([]byte(line))
  			}
  		}
  		fo.Write([]byte("\n"))
  	}
  }

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.6``, ``opencc 0.4.3-2build1``,
``GitHub: siongui/go-opencc``.

----

References:

.. [1] `[Python] Automatically Convert Traditional Chinese PO file to Simplified Chinese <{filename}../../01/08/python-automatically-convert-zhtw-po-file-to-zhcn%en.rst>`_

.. [2] `create zh_CN PO from zh_TW · siongui/pali@365d46c · GitHub <https://github.com/siongui/pali/commit/365d46ca999b3431e664c72502a5ba8cba8bd901>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _OpenCC: http://opencc.byvoid.com/
.. _OpenCC Go binding: https://github.com/siongui/go-opencc
.. _PO: https://www.gnu.org/software/gettext/manual/html_node/PO-Files.html
.. _Traditional Chinese: https://en.wikipedia.org/wiki/Traditional_Chinese_characters
.. _Simplified Chinese: https://en.wikipedia.org/wiki/Simplified_Chinese_characters
.. _zh_TW: https://docs.oracle.com/cd/E19455-01/806-0169/6j9hsml3g/index.html
.. _zh_CN: https://docs.oracle.com/cd/E19683-01/806-6642/new-tbl-72/index.html

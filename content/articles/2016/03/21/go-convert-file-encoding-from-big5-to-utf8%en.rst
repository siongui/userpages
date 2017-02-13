[Golang] Convert File Encoding From Big5 to UTF-8
#################################################

:date: 2016-03-21T01:23+08:00
:tags: Go, Golang, Commandline, String Manipulation, iconv command,
       File Input/Output, Go flag Package
:category: Go
:summary: Convert file encoding from Big5_ to UTF-8_ via Go_ programming
          language.
:adsu: yes

Convert file encoding from Big5_ encoding to UTF-8_ via Golang.
Only for demonstration purpose. It will be much easier to use iconv_ command
to convert encoding of one file.

Install `Go iconv binding`_:

.. code-block:: bash

  $ go get -u github.com/djimenez/iconv-go

Source code:

.. code-block:: go

  package main

  import (
  	"flag"
  	"fmt"
  	iconv "github.com/djimenez/iconv-go"
  	"io"
  	"os"
  )

  func big5ToUTF8(path, outpath string) {
  	fmt.Println("Converting " + path + " from Big5 to UTF-8 ...")

  	f, err := os.Open(path)
  	if err != nil {
  		panic(err)
  	}
  	defer f.Close()

  	reader, err := iconv.NewReader(f, "big5", "utf-8")
  	if err != nil {
  		panic(err)
  	}

  	fo, err := os.Create(outpath)
  	if err != nil {
  		panic(err)
  	}
  	defer fo.Close()

  	io.Copy(fo, reader)
  }

  func main() {
  	pPath := flag.String("input", "", "Path of file to be processed")
  	pOutpath := flag.String("output", "", "Output path of processed file")
  	flag.Parse()
  	path := *pPath
  	outpath := *pOutpath
  	if path == "" || outpath == "" {
  		fmt.Fprintf(os.Stderr, "Error: empty path!\n")
  		return
  	}
  	big5ToUTF8(path, outpath)
  	fmt.Println(outpath + " saved")
  }

.. adsu:: 2

Command line usage:

.. code-block:: bash

  $ go run big5ToUTF8.go -input=big5.html -output=utf8.html

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.6``.

----

References:

.. [1] `big5 to utf-8 · twnanda/twnanda@5681c8b · GitHub <https://github.com/twnanda/twnanda/commit/5681c8b94c68da8cf15cc4fc91ac4401f80d6eb7>`_

.. [2] `[Bash] Convert Files in Directory From Big5 to UTF-8 <{filename}../18/bash-convert-files-in-directory-from-big5-to-utf8%en.rst>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Big5: https://en.wikipedia.org/wiki/Big5
.. _UTF-8: https://en.wikipedia.org/wiki/UTF-8
.. _iconv: http://linux.die.net/man/1/iconv
.. _Go iconv binding: https://github.com/djimenez/iconv-go

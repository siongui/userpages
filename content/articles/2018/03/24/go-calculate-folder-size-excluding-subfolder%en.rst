[Golang] Calculate Directory Size Excluding Sub-Directories
###########################################################

:date: 2018-03-24T23:39+08:00
:tags: Go, Golang, File Input/Output
:category: Go
:summary: Calcualte total size of files in folder, excluding sub-folders in Go.
:og_image: https://www.ostechnix.com/wp-content/uploads/2016/05/Deepin-Terminal_012-1.png
:adsu: yes


Calculate total size of files in directory, excluding sub-directories.
We use `ioutil.ReadDir`_ to read information of files in directory, and
calculate the sum of size of all files.

.. code-block:: go

  import (
  	"io/ioutil"
  	"os"
  )

  func calculateDirSize(dirpath string) (dirsize int64, err error) {
  	err = os.Chdir(dirpath)
  	if err != nil {
  		return
  	}
  	files, err := ioutil.ReadDir(".")
  	if err != nil {
  		return
  	}

  	for _, file := range files {
  		if file.Mode().IsRegular() {
  			dirsize += file.Size()
  		}
  	}
  	return
  }

.. adsu:: 2

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.10``

----

References:

.. [1] `ioutil - The Go Programming Language <https://golang.org/pkg/io/ioutil/#ReadDir>`_

.. _ioutil.ReadDir: https://golang.org/pkg/io/ioutil/#ReadDir

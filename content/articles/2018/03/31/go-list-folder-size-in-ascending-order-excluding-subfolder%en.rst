[Golang] List Directory Size in Ascending Order (Excluding Sub-Directories)
###########################################################################

:date: 2018-03-31T23:30+08:00
:tags: Go, Golang, Go sort
:category: Go
:summary: List folder size in ascending orfer, excluding sub-folders, in Go.
:og_image: https://www.ostechnix.com/wp-content/uploads/2016/05/Deepin-Terminal_012-1.png
:adsu: yes


List folder size in ascending orfer, excluding sub-folders, in Go.
We use `ioutil.ReadDir`_ to read information of files/directories in current
directory, excluding sub-directories. If you want to include sub-directories,
use `filepath.Walk`_ instead. After calculating directory size, use sort.Slice_
to sort the size in ascending order.

.. code-block:: go

  package main

  import (
  	"fmt"
  	"io/ioutil"
  	"os"
  	"sort"
  )

  type nameSize struct {
  	name string
  	size int64
  }

  func calculateDirSize() (dirsize int64) {
  	files, err := ioutil.ReadDir(".")
  	if err != nil {
  		fmt.Println(err)
  		return
  	}

  	for _, file := range files {
  		if file.Mode().IsRegular() {
  			dirsize += file.Size()
  		}
  	}
  	return
  }

  func main() {
  	// read all files/directories in current directory
  	dirs, err := ioutil.ReadDir(".")
  	if err != nil {
  		fmt.Println(err)
  		return
  	}

  	var nameSizeArray []nameSize
  	for _, dir := range dirs {
  		// if it is not dir, continue to next dir
  		if !dir.IsDir() {
  			continue
  		}

  		err = os.Chdir(dir.Name())
  		if err != nil {
  			fmt.Println(err)
  			return
  		}

  		size := calculateDirSize()
  		ns := nameSize{dir.Name(), size}
  		nameSizeArray = append(nameSizeArray, ns)

  		err = os.Chdir("..")
  		if err != nil {
  			fmt.Println(err)
  			return
  		}
  	}

  	// sort all directories in current directory according to size
  	sort.Slice(nameSizeArray, func(i, j int) bool {
  		return nameSizeArray[i].size < nameSizeArray[j].size
  	})

  	// print sorted result
  	for _, ns := range nameSizeArray {
  		fmt.Println(ns)
  	}
  }

If you want to use shell command to list directory size, see [2]_.

.. adsu:: 2

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.10.1``

----

References:

.. [1] `[Golang] Calculate Directory Size Excluding Sub-Directories <{filename}/articles/2018/03/24/go-calculate-folder-size-excluding-subfolder%en.rst>`_
.. [2] `[Bash] List Directory Size in Descending and Ascending Order <{filename}/articles/2018/03/22/bash-list-folder-size-in-descending-and-ascending-order%en.rst>`_

.. _ioutil.ReadDir: https://golang.org/pkg/io/ioutil/#ReadDir
.. _filepath.Walk: https://golang.org/pkg/path/filepath/#Walk
.. _sort.Slice: https://golang.org/pkg/sort/#Slice

[Golang] Find Last Modified File With Specific Name Prefix
##########################################################

:date: 2018-04-18T23:52+08:00
:tags: Go, Golang, Go time Package, List Files in Directory
:category: Go
:summary: Get last modified file, name of which starts with specific prefix,
          in Go.
:og_image: https://i.ytimg.com/vi/jTBtVrn-Cq8/hqdefault.jpg
:adsu: yes


This post shows you how to find the last modified file, name of which starts
with specific prefx in Go.

We use `ioutil.ReadDir`_ to read files in the directory, but not in
sub-directories. If you need to read also sub-directories, Use `filepath.Walk`_.

.. code-block:: go

  package main

  import (
  	"io/ioutil"
  	"os"
  	"strings"
  )

  func findLastFileStartsWith(dir, prefix string) (lastFile os.FileInfo, err error) {
  	files, err := ioutil.ReadDir(dir)
  	if err != nil {
  		return
  	}

  	for _, file := range files {
  		if !file.Mode().IsRegular() {
  			continue
  		}
  		if strings.HasPrefix(file.Name(), prefix) {
  			if lastFile == nil {
  				lastFile = file
  			} else {
  				if lastFile.ModTime().Before(file.ModTime()) {
  					lastFile = file
  				}
  			}
  		}
  	}

  	if lastFile == nil {
  		err = os.ErrNotExist
  		return
  	}
  	return
  }

  func main() {
  	file, err := findLastFileStartsWith(".", "myprefix")
  	if err != nil {
  		panic(err)
  	}
  	println(file.Name())
  }

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.10.1``

**References**

.. [1] `[Golang] Find Last Modified File Before Specific Time <{filename}/articles/2018/04/06/go-get-last-modified-file-before-specific-time%en.rst>`_
.. [2] `[Golang] Find Oldest Modified File in Directory <{filename}/articles/2018/04/11/go-get-oldest-modified-file-in-directory%en.rst>`_

.. _ioutil.ReadDir: https://golang.org/pkg/io/ioutil/#ReadDir
.. _filepath.Walk: https://golang.org/pkg/path/filepath/#Walk

[Golang] Find Last Modified File
################################

:date: 2018-04-06T23:47+08:00
:tags: Go, Golang, Go time Package, List Files in Directory
:category: Go
:summary: Get last modified file (before a specific time) in Go.
:og_image: https://i.ytimg.com/vi/jTBtVrn-Cq8/hqdefault.jpg
:adsu: yes


This post shows you how to find the last modified file before a specific time in
Go. We use `filepath.Walk`_ to access all files in a directory and find the last
modified file before a specific time.

.. code-block:: go

  package main

  import (
  	"fmt"
  	"path/filepath"
  	"time"
  )

  func FindLastModifiedFileBefore(dir string, t time.Time) (path string, info os.FileInfo, err error) {
  	isFirst := true
  	min := 0 * time.Second
  	err = filepath.Walk(dir, func(p string, i os.FileInfo, e error) error {
  		if e != nil {
  			return e
  		}

  		if !i.IsDir() && i.ModTime().Before(t) {
  			if isFirst {
  				isFirst = false
  				path = p
  				info = i
  				min = t.Sub(i.ModTime())
  			}
  			if diff := t.Sub(i.ModTime()); diff < min {
  				path = p
  				min = diff
  				info = i
  			}
  		}
  		return nil
  	})
  	return
  }

  func main() {
  	dir := "/path/to/your/dir"
  	path, info, err := FindLastModifiedFileBefore(dir, time.Now())
  	if err != nil {
  		panic(err)
  	}
  	fmt.Println(path)
  	fmt.Println(info)
  }

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.10.1``

**References**

.. [1] | `linux find last modified file - Google search <https://www.google.com/search?q=linux+find+last+modified+file>`_
       | `linux find last modified file - DuckDuckGo search <https://duckduckgo.com/?q=linux+find+last+modified+file>`_
       | `linux find last modified file - Ecosia search <https://www.ecosia.org/search?q=linux+find+last+modified+file>`_
       | `linux find last modified file - Qwant search <https://www.qwant.com/?q=linux+find+last+modified+file>`_
       | `linux find last modified file - Bing search <https://www.bing.com/search?q=linux+find+last+modified+file>`_
       | `linux find last modified file - Yahoo search <https://search.yahoo.com/search?p=linux+find+last+modified+file>`_
       | `linux find last modified file - Baidu search <https://www.baidu.com/s?wd=linux+find+last+modified+file>`_
       | `linux find last modified file - Yandex search <https://www.yandex.com/search/?text=linux+find+last+modified+file>`_

.. _filepath.Walk: https://golang.org/pkg/path/filepath/#Walk

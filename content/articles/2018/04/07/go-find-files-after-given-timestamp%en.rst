[Golang] Find Files After Given Time
####################################

:date: 2018-04-07T23:00+08:00
:tags: Go, Golang, Go time Package, List Files in Directory
:category: Go
:summary: Find files after a given timestamp in Go.
:og_image: https://i.ytimg.com/vi/jTBtVrn-Cq8/hqdefault.jpg
:adsu: yes


This post shows you how to find files after a given timestamp in Go.
We use `filepath.Walk`_ to access all files in a directory and use `Time.After`_
to check if the modified time of the file is after or not.

If you want to find files before given time, use `Time.Before`_ instead.

You can also use shell command to do this [1]_ [2]_, but I like to code in Go
because it is more readable.

.. code-block:: go

  package main

  import (
  	"fmt"
  	"os"
  	"path/filepath"
  	"time"
  )

  func FindFilesAfter(dir string, t time.Time) (paths []string, infos []os.FileInfo, err error) {
  	err = filepath.Walk(dir, func(p string, i os.FileInfo, e error) error {
  		if err != nil {
  			return err
  		}

  		if !i.IsDir() && i.ModTime().After(t) {
  			paths = append(paths, p)
  			infos = append(infos, i)
  		}
  		return nil
  	})
  	return
  }

  func checkFile(path string, info os.FileInfo) {
  	fmt.Println(path)
  	fmt.Println(info)
  }

  func main() {
  	dir := "path/to/your/dir"
  	t, err := time.Parse("2006-01-02T15:04:05-07:00", "2018-04-07T05:48:03+08:00")
  	if err != nil {
  		panic(err)
  	}
  	paths, infos, err := FindFilesAfter(dir, t)
  	if err != nil {
  		panic(err)
  	}
  	for i, _ := range paths {
  		checkFile(paths[i], infos[i])
  	}
  }

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.10.1``

**References**

.. [1] | `linux find files after timestamp - Google search <https://www.google.com/search?q=linux+find+files+after+timestamp>`_
       | `linux find files after timestamp - DuckDuckGo search <https://duckduckgo.com/?q=linux+find+files+after+timestamp>`_
       | `linux find files after timestamp - Ecosia search <https://www.ecosia.org/search?q=linux+find+files+after+timestamp>`_
       | `linux find files after timestamp - Qwant search <https://www.qwant.com/?q=linux+find+files+after+timestamp>`_
       | `linux find files after timestamp - Bing search <https://www.bing.com/search?q=linux+find+files+after+timestamp>`_
       | `linux find files after timestamp - Yahoo search <https://search.yahoo.com/search?p=linux+find+files+after+timestamp>`_
       | `linux find files after timestamp - Baidu search <https://www.baidu.com/s?wd=linux+find+files+after+timestamp>`_
       | `linux find files after timestamp - Yandex search <https://www.yandex.com/search/?text=linux+find+files+after+timestamp>`_

.. [2] | `linux find files after certain date - Google search <https://www.google.com/search?q=linux+find+files+after+certain+date>`_
       | `linux find files after certain date - DuckDuckGo search <https://duckduckgo.com/?q=linux+find+files+after+certain+date>`_
       | `linux find files after certain date - Ecosia search <https://www.ecosia.org/search?q=linux+find+files+after+certain+date>`_
       | `linux find files after certain date - Qwant search <https://www.qwant.com/?q=linux+find+files+after+certain+date>`_
       | `linux find files after certain date - Bing search <https://www.bing.com/search?q=linux+find+files+after+certain+date>`_
       | `linux find files after certain date - Yahoo search <https://search.yahoo.com/search?p=linux+find+files+after+certain+date>`_
       | `linux find files after certain date - Baidu search <https://www.baidu.com/s?wd=linux+find+files+after+certain+date>`_
       | `linux find files after certain date - Yandex search <https://www.yandex.com/search/?text=linux+find+files+after+certain+date>`_

.. _filepath.Walk: https://golang.org/pkg/path/filepath/#Walk
.. _Time.After: https://golang.org/pkg/time/#Time.After
.. _Time.Before: https://golang.org/pkg/time/#Time.Before

[Golang] Find Files Older Than One Day
######################################

:date: 2018-03-25T23:15+08:00
:tags: Go, Golang, Go time Package
:category: Go
:summary: Find files modified more than one day via Go standard *time* package.
:og_image: https://www.systemcodegeeks.com/wp-content/uploads/2016/02/10-Examples-of-find-comand-in-UNIX.png
:adsu: yes


Find files of which modified time is older than one day in Go.

1. Use ioutil.ReadDir_ or filepath.Walk_ to find files: if you want to find all
   files including sub-directories, use filepath.Walk_. If you just want files
   in a directory but not in sub-directories, use ioutil.ReadDir_. Here we use
   ioutil.ReadDir_ in our example.

2. use `time.Now()`_ to get current time and use `func (Time) Sub`_ to perform
   subtraction of current time and modified time of file. If the duration is
   greater than 24 hours, than it's more than one day.

The following code is the implementation of above idea:

.. code-block:: go

  import (
  	"io/ioutil"
  	"os"
  	"time"
  )

  func isOlderThanOneDay(t time.Time) bool {
  	return time.Now().Sub(t) > 24*time.Hour
  }

  func findFilesOlderThanOneDay(dir string) (files []os.FileInfo, err error) {
  	tmpfiles, err := ioutil.ReadDir(dir)
  	if err != nil {
  		return
  	}

  	for _, file := range tmpfiles {
  		if file.Mode().IsRegular() {
  			if isOlderThanOneDay(file.ModTime()) {
  				files = append(files, file)
  			}
  		}
  	}
  	return
  }

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.10``

----

References:

.. [1] | `file longer than one day - Google search <https://www.google.com/search?q=file+longer+than+one+day>`_
       | `file longer than one day - DuckDuckGo search <https://duckduckgo.com/?q=file+longer+than+one+day>`_
       | `file longer than one day - Ecosia search <https://www.ecosia.org/search?q=file+longer+than+one+day>`_
       | `file longer than one day - Qwant search <https://www.qwant.com/?q=file+longer+than+one+day>`_
       | `file longer than one day - Bing search <https://www.bing.com/search?q=file+longer+than+one+day>`_
       | `file longer than one day - Yahoo search <https://search.yahoo.com/search?p=file+longer+than+one+day>`_
       | `file longer than one day - Baidu search <https://www.baidu.com/s?wd=file+longer+than+one+day>`_
       | `file longer than one day - Yandex search <https://www.yandex.com/search/?text=file+longer+than+one+day>`_
.. [2] `How To Find And Delete Files Older Than X Days In Linux - OSTechNix <https://www.ostechnix.com/how-to-find-and-delete-files-older-than-x-days-in-linux/>`_
.. [3] `10 Examples of find command in Unix and Linux | System Code Geeks - 2018 <https://www.systemcodegeeks.com/shell-scripting/bash/10-examples-find-command-unix-linux/>`_
.. [4] `Mocking system time in tests : golang <https://old.reddit.com/r/golang/comments/a8yx3x/mocking_system_time_in_tests/>`_


.. _ioutil.ReadDir: https://golang.org/pkg/io/ioutil/#ReadDir
.. _filepath.Walk: https://golang.org/pkg/path/filepath/#Walk
.. _time.Now(): https://golang.org/pkg/time/#Now
.. _func (Time) Sub: https://golang.org/pkg/time/#Time.Sub

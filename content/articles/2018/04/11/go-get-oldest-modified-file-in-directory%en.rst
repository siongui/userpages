[Golang] Find Oldest Modified File in Directory
###############################################

:date: 2018-04-11T09:18+08:00
:tags: Go, Golang, Go time Package
:category: Go
:summary: Find oldest modified file in the directory, excluding sub-directories,
          in Go.
:og_image: https://i.ytimg.com/vi/jTBtVrn-Cq8/hqdefault.jpg
:adsu: yes


Find oldest modified file in the directory, excluding sub-directories, in Go.
We use `ioutil.ReadDir`_ to read information of files and find the oldest
modified time of file. `ioutil.ReadDir`_ does not read sub-directories. If you
need to read also sub-directories, use `filepath.Walk`_ instead.

.. code-block:: go

  import (
  	"io/ioutil"
  	"os"
  	"time"
  )

  func FindOldestFile(dir string) (oldestFile os.FileInfo, err error) {
  	files, err := ioutil.ReadDir(dir)
  	if err != nil {
  		return
  	}

  	oldestTime := time.Now()
  	for _, file := range files {
  		if file.Mode().IsRegular() && file.ModTime().Before(oldestTime) {
  			oldestFile = file
  			oldestTime = file.ModTime()
  		}
  	}

  	if oldestFile == nil {
  		err = os.ErrNotExist
  	}
  	return
  }

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.10.1``

**References**

.. [1] | `oldest modified time of file in linux - Google search <https://www.google.com/search?q=oldest+modified+time+of+file+in+linux>`_
       | `oldest modified time of file in linux - DuckDuckGo search <https://duckduckgo.com/?q=oldest+modified+time+of+file+in+linux>`_
       | `oldest modified time of file in linux - Ecosia search <https://www.ecosia.org/search?q=oldest+modified+time+of+file+in+linux>`_
       | `oldest modified time of file in linux - Qwant search <https://www.qwant.com/?q=oldest+modified+time+of+file+in+linux>`_
       | `oldest modified time of file in linux - Bing search <https://www.bing.com/search?q=oldest+modified+time+of+file+in+linux>`_
       | `oldest modified time of file in linux - Yahoo search <https://search.yahoo.com/search?p=oldest+modified+time+of+file+in+linux>`_
       | `oldest modified time of file in linux - Baidu search <https://www.baidu.com/s?wd=oldest+modified+time+of+file+in+linux>`_
       | `oldest modified time of file in linux - Yandex search <https://www.yandex.com/search/?text=oldest+modified+time+of+file+in+linux>`_

.. _ioutil.ReadDir: https://golang.org/pkg/io/ioutil/#ReadDir
.. _filepath.Walk: https://golang.org/pkg/path/filepath/#Walk

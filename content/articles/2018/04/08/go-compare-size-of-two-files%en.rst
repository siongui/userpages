[Golang] Compare the Size of Two Files
######################################

:date: 2018-04-08T06:27+08:00
:tags: Go, Golang
:category: Go
:summary: Compare if the size of two files is the same in Go.
:og_image: https://securingtomorrow.mcafee.com/wp-content/uploads/2018/02/20180214-Free-Ransomware-9.png
:adsu: yes


This post shows how to check if the size of two files is the same. Use
`os.Stat`_ [2]_ to read information (`os.FileInfo`_) of file and check if the
size is the same.

.. code-block:: go

  import (
  	"os"
  )

  func isFileSameSize(path1, path2 string) (sameSize bool, err error) {
  	f1, err := os.Stat(path1)
  	if err != nil {
  		return
  	}
  	f2, err := os.Stat(path2)
  	if err != nil {
  		return
  	}
  	sameSize = (f1.Size() == f2.Size())
  	return
  }

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.10``

**References**:

.. [1] | `linux compare file size - Google search <https://www.google.com/search?q=linux+compare+file+size>`_
       | `linux compare file size - DuckDuckGo search <https://duckduckgo.com/?q=linux+compare+file+size>`_
       | `linux compare file size - Ecosia search <https://www.ecosia.org/search?q=linux+compare+file+size>`_
       | `linux compare file size - Qwant search <https://www.qwant.com/?q=linux+compare+file+size>`_
       | `linux compare file size - Bing search <https://www.bing.com/search?q=linux+compare+file+size>`_
       | `linux compare file size - Yahoo search <https://search.yahoo.com/search?p=linux+compare+file+size>`_
       | `linux compare file size - Baidu search <https://www.baidu.com/s?wd=linux+compare+file+size>`_
       | `linux compare file size - Yandex search <https://www.yandex.com/search/?text=linux+compare+file+size>`_
.. [2] `How to get file length in Go? - Stack Overflow <https://stackoverflow.com/a/40337760>`_

.. _os.Stat: https://golang.org/pkg/os/#Stat
.. _os.FileInfo: https://golang.org/pkg/os/#FileInfo

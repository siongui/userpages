[Golang] Create Directory If Not Exist
######################################

:date: 2017-03-28T09:52+08:00
:tags: Go, Golang, Existence Detection
:category: Go
:summary: Create a directory if it does not exist. Otherwise do nothing.
:adsu: yes

Create a directory if it does not exist. Otherwise do nothing.

.. code-block:: go

  import "os"

  func CreateDirIfNotExist(dir string) {
  	if _, err := os.Stat(dir); os.IsNotExist(err) {
  		err = os.MkdirAll(dir, 0755)
  		if err != nil {
  			panic(err)
  		}
  	}
  }

`os.MkdirAll`_ is similar to `mkdir -p`_ in shell command, which also creates
parent directory if not exists.

Tested on: ``Go 1.8``, ``Ubuntu Linux 16.10``

.. adsu:: 2

----

References:

.. [1] | `golang create directory if not exists - Google search <https://www.google.com/search?q=golang+create+directory+if+not+exists>`_
       | `golang create directory if not exists - DuckDuckGo search <https://duckduckgo.com/?q=golang+create+directory+if+not+exists>`_
       | `golang create directory if not exists - Ecosia search <https://www.ecosia.org/search?q=golang+create+directory+if+not+exists>`_
       | `golang create directory if not exists - Qwant search <https://www.qwant.com/?q=golang+create+directory+if+not+exists>`_
       | `golang create directory if not exists - Bing search <https://www.bing.com/search?q=golang+create+directory+if+not+exists>`_
       | `golang create directory if not exists - Yahoo search <https://search.yahoo.com/search?p=golang+create+directory+if+not+exists>`_
       | `golang create directory if not exists - Baidu search <https://www.baidu.com/s?wd=golang+create+directory+if+not+exists>`_
       | `golang create directory if not exists - Yandex search <https://www.yandex.com/search/?text=golang+create+directory+if+not+exists>`_

.. [2] `[Makefile] Create Directory If Not Exist <{filename}../../../2016/01/30/makefile-create-directory-if-not-exist%en.rst>`_
.. [3] `Package write provides a way to atomically create or replace a file or symbolic link. : golang <https://old.reddit.com/r/golang/comments/9sq6x1/package_write_provides_a_way_to_atomically_create/>`_

.. _os.MkdirAll: https://golang.org/pkg/os/#MkdirAll
.. _mkdir -p: https://www.google.com/search?q=what+is+mkdir+-p+option

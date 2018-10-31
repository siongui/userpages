[Golang] Check if File, Directory, or Symlink Exist
###################################################

:date: 2017-03-27T22:44+08:00
:tags: Go, Golang, Existence Detection
:category: Go
:summary: Check if a file, directory, or symbolic link exists, or not exists,
          in Go programming language.
:adsu: yes

Check if a file, directory, or `symbolic link`_ exists in Golang:

.. code-block:: go

  import "os"

  func IsExist(path string) bool {
  	if _, err := os.Stat(path); err == nil {
  		// exist
  		return true
  	}
  	// not exist
  	return false
  }

If you want to check if NOT exist:

.. code-block:: go

  import "os"

  func IsNotExist(path string) bool {
  	if _, err := os.Stat(path); os.IsNotExist(err) {
  		// not exist
  		return true
  	}
  	// exist
  	return false
  }

.. adsu:: 2

----

References:

.. [1] | `go check if directory exists - Google search <https://www.google.com/search?q=go+check+if+directory+exists>`_
       | `go check if directory exists - DuckDuckGo search <https://duckduckgo.com/?q=go+check+if+directory+exists>`_
       | `go check if directory exists - Ecosia search <https://www.ecosia.org/search?q=go+check+if+directory+exists>`_
       | `go check if directory exists - Qwant search <https://www.qwant.com/?q=go+check+if+directory+exists>`_
       | `go check if directory exists - Bing search <https://www.bing.com/search?q=go+check+if+directory+exists>`_
       | `go check if directory exists - Yahoo search <https://search.yahoo.com/search?p=go+check+if+directory+exists>`_
       | `go check if directory exists - Baidu search <https://www.baidu.com/s?wd=go+check+if+directory+exists>`_
       | `go check if directory exists - Yandex search <https://www.yandex.com/search/?text=go+check+if+directory+exists>`_

.. [2] `go - How to check whether a file or directory denoted by a path exists in Golang? - Stack Overflow <http://stackoverflow.com/questions/10510691/how-to-check-whether-a-file-or-directory-denoted-by-a-path-exists-in-golang>`_
.. [3] `Package write provides a way to atomically create or replace a file or symbolic link. : golang <https://old.reddit.com/r/golang/comments/9sq6x1/package_write_provides_a_way_to_atomically_create/>`_

.. _symbolic link: https://www.google.com/search?q=symbolic+link

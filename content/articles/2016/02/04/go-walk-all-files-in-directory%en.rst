[Golang] Walk All Files in Directory
####################################

:date: 2016-02-04T04:47+08:00
:modified: 2018-05-10T23:35+08:00
:tags: Go, Golang, List Files in Directory
:category: Go
:summary: List all files in a directory in Go_. Similar to Python_ `os.walk`_.
:og_image: http://www.avajava.com/tutorials/files/how-do-i-recursively-display-all-files-and-directories-in-a-directory/how-do-i-recursively-display-all-files-and-directories-in-a-directory-01.gif
:adsu: yes

List all files in a directory in Golang_. Use filepath.Walk_ in Go standard
library, which is similar to Python_ `os.walk`_.

.. code-block:: go

  import (
  	"fmt"
  	"os"
  	"path/filepath"
  )

  func WalkAllFilesInDir(dir string) error {
  	return filepath.Walk(dir, func(path string, info os.FileInfo, e error) error {
  		if e != nil {
  			return e
  		}

  		// check if it is a regular file (not dir)
  		if info.Mode().IsRegular() {
  			fmt.Println("file name:", info.Name())
  			fmt.Println("file path:", path)
  		}
  		return nil
  	})
  }

If you do not want to walk into sub-directories, use ioutil.ReadDir_.

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 18.04``, ``Go 1.10.2``.

----

References:

.. [1] `golang list files in directory <https://www.google.com/search?q=golang+list+files+in+directory>`_

.. [2] `List directory in go - Stack Overflow <http://stackoverflow.com/questions/14668850/list-directory-in-go>`_

.. [3] `golang walk filesystem <https://www.google.com/search?q=golang+walk+filesystem>`_

.. [4] `golang walk directory <https://www.google.com/search?q=golang+walk+directory>`_

.. [5] `Python os.walk <https://docs.python.org/2/library/os.html#os.walk>`_

.. [6] `filepath - The Go Programming Language <https://golang.org/pkg/path/filepath/>`_
.. [7] | `A quick comparison between different Go file walk implementations : golang <https://www.reddit.com/r/golang/comments/83lwfs/a_quick_comparison_between_different_go_file_walk/>`_
       | `A quick comparison between different Go file walk implementations | Ben E. C. Boyter <http://www.boyter.org/2018/03/quick-comparison-go-file-walk-implementations/>`_
.. [8] `[Golang] Delete Zero Size Files in Directory <{filename}/articles/2018/05/10/go-delete-zero-size-file-in-folder%en.rst>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Python: https://www.python.org/
.. _os.walk: https://docs.python.org/2/library/os.html#os.walk
.. _filepath.Walk: https://golang.org/pkg/path/filepath/#Walk
.. _ioutil.ReadDir: https://golang.org/pkg/io/ioutil/#ReadDir

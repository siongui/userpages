[Golang] Delete Zero Size Files in Directory
############################################

:date: 2018-05-10T23:22+08:00
:tags: Go, Golang, List Files in Directory
:category: Go
:summary: Delete all files with zero size in a directory in Go. Sub-directories
          included.
:og_image: http://www.avajava.com/tutorials/files/how-do-i-recursively-display-all-files-and-directories-in-a-directory/how-do-i-recursively-display-all-files-and-directories-in-a-directory-01.gif
:adsu: yes


Find all files with size = 0 in the directory recursively. First we use
filepath.Walk_ to recursively find all files in the directory, sub-directories
included. Then we check the file size and if it is zero, call os.Remove_ to
delete the file.

.. code-block:: go

  import (
  	"fmt"
  	"os"
  	"path/filepath"
  )

  func DeleteZeroSizeFile(dir string) error {
  	return filepath.Walk(dir, func(path string, info os.FileInfo, e error) error {
  		if e != nil {
  			return e
  		}

  		if info.Mode().IsRegular() && info.Size() == 0 {
  			fmt.Println("Removing", path)
  			os.Remove(path)
  		}
  		return nil
  	})
  }

If you do not want to walk into sub-directories, use ioutil.ReadDir_.

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 18.04``, ``Go 1.10.2``.

References:

.. [1] `func Walk - filepath - The Go Programming Language <https://golang.org/pkg/path/filepath/#Walk>`_
.. [2] `func ReadDir - ioutil - The Go Programming Language <https://golang.org/pkg/io/ioutil/#ReadDir>`_

.. _filepath.Walk: https://golang.org/pkg/path/filepath/#Walk
.. _ioutil.ReadDir: https://golang.org/pkg/io/ioutil/#ReadDir
.. _os.Remove: https://golang.org/pkg/os/#Remove

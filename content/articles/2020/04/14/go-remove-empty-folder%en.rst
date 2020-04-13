[Golang] Remove Empty Directory
###############################

:date: 2020-04-14T01:29+08:00
:tags: Go, Golang
:category: Go
:summary: Given a directory, remove all empty sub-directories in Go.
:og_image: https://golang.org/lib/godoc/images/footer-gopher.jpg
:adsu: yes


Given a directory, remove all empty sub-directories in Go_. We use
filepath.Walk_ to find all directories and ioutil.ReadDir_ to check if a
directory is empty. Then use os.Remove_ to remove the directory if the directory
is empty.

.. code-block:: go

  package main
  
  import (
  	"flag"
  	"fmt"
  	"io/ioutil"
  	"os"
  	"path/filepath"
  )
  
  func processDir(path string, info os.FileInfo) {
  	files, err := ioutil.ReadDir(path)
  	if err != nil {
  		panic(err)
  	}
  
  	if len(files) != 0 {
  		return
  	}
  
  	err = os.Remove(path)
  	if err != nil {
  		panic(err)
  	}
  
  	fmt.Println(path, "removed!")
  }
  
  func main() {
  	root := flag.String("root", "Instagram", "dir of downloaded files")
  	flag.Parse()
  
  	err := filepath.Walk(*root, func(path string, info os.FileInfo, err error) error {
  		if err != nil {
  			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
  			return err
  		}
  		if info.IsDir() {
  			processDir(path, info)
  		}
  
  		return nil
  	})
  	if err != nil {
  		panic(err)
  		return
  	}
  }

----

References:

.. [1] `[Golang] Parse Command Line Arguments - String Variable <{filename}/articles/2016/12/21/go-parse-commandline-arguments-string-variable%en.rst>`_
.. [2] `[Golang] Walk All Files in Directory <{filename}/articles/2016/02/04/go-walk-all-files-in-directory%en.rst>`_

.. _Go: https://golang.org/
.. _filepath.Walk: https://golang.org/pkg/path/filepath/#Walk
.. _ioutil.ReadDir: https://golang.org/pkg/io/ioutil/#ReadDir
.. _os.Remove: https://golang.org/pkg/os/#Remove

[Golang] Split Contents of a Directory into Multiple Sub Directories
####################################################################

:date: 2018-04-03T23:18+08:00
:tags: Go, Golang
:category: Go
:summary: Split a folder with many files into multiple sub-folders in Go.
:og_image: https://docs.oracle.com/cd/E13222_01/wls/docs90/programming/wwimages/split_ear_overview.gif
:adsu: yes


Split contents of a directory into multiple sub-directory in Go. For example,
if you have a directory with 1,000 files, you split them into 4 sub-directories,
each with 300, 300, 300, 100 files inside. You can specify the maximum number of
files in each subdirectories.

You can also do this via shell command. See [2]_.

.. code-block:: go

  package main

  import (
  	"flag"
  	"fmt"
  	"io/ioutil"
  	"os"
  	"path"
  	"strconv"
  )

  var maxFilesPerDir = 1500

  func getTargetDir(groupCount int, srcDir string) string {
  	return path.Join(srcDir, "dir"+strconv.Itoa(groupCount))
  }

  func createDirIfNotExist(dir string) {
  	if _, err := os.Stat(dir); os.IsNotExist(err) {
  		err = os.MkdirAll(dir, 0755)
  		if err != nil {
  			panic(err)
  		}
  	}
  }

  func moveFilesToDir(filesGroup [][]os.FileInfo, srcDir string) {
  	for i, group := range filesGroup {
  		targetDir := getTargetDir(i, srcDir)
  		createDirIfNotExist(targetDir)
  		for _, file := range group {
  			filepath := path.Join(srcDir, file.Name())
  			targetFilepath := path.Join(targetDir, file.Name())
  			fmt.Println("move", filepath, "to", targetFilepath)
  			os.Rename(filepath, targetFilepath)
  		}
  	}
  }

  func makeFilesGroup(srcDir string) [][]os.FileInfo {
  	files, err := ioutil.ReadDir(srcDir)
  	if err != nil {
  		panic(err)
  	}

  	count := 0
  	var filesGroup [][]os.FileInfo
  	filesGroup = append(filesGroup, []os.FileInfo{})
  	groupCount := 0
  	for _, file := range files {
  		if file.Mode().IsRegular() {
  			filesGroup[groupCount] = append(filesGroup[groupCount], file)
  			count++
  			if count == maxFilesPerDir {
  				count = 0
  				filesGroup = append(filesGroup, []os.FileInfo{})
  				groupCount++
  			}
  		}
  	}

  	return filesGroup
  }

  func main() {
  	srcDir := flag.String("src", "", "source dir to split")
  	max := flag.Int("max", 1500, "max number of files in split dir")
  	flag.Parse()

  	maxFilesPerDir = *max
  	filesGroup := makeFilesGroup(*srcDir)
  	moveFilesToDir(filesGroup, *srcDir)
  }

Usage: *max* means the number of max number of files in sub directory, and *src*
means the path to the directory to split

.. code-block:: bash

  $ go run splitdir.go -max=500 -src=/path/to/dir/

.. adsu:: 2

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.10.1``

----

References:

.. [1] | `split directory - Google search <https://www.google.com/search?q=split+directory>`_
       | `split directory - DuckDuckGo search <https://duckduckgo.com/?q=split+directory>`_
       | `split directory - Ecosia search <https://www.ecosia.org/search?q=split+directory>`_
       | `split directory - Qwant search <https://www.qwant.com/?q=split+directory>`_
       | `split directory - Bing search <https://www.bing.com/search?q=split+directory>`_
       | `split directory - Yahoo search <https://search.yahoo.com/search?p=split+directory>`_
       | `split directory - Baidu search <https://www.baidu.com/s?wd=split+directory>`_
       | `split directory - Yandex search <https://www.yandex.com/search/?text=split+directory>`_
.. [2] `command line - Split contents of a directory into multiple sub directories - Ask Ubuntu <https://askubuntu.com/questions/584724/split-contents-of-a-directory-into-multiple-sub-directories>`_

.. _Go Playground: https://play.golang.org/

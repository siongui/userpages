Golang Template Parse Directory
###############################

:date: 2017-02-13T21:15+08:00
:tags: Go, Golang, Golang template, List Files in Directory
:category: Go
:summary: Complement to Go_ `text/template`_ and `html/template`_ package -
          Parse all files in directory (including sub-directories).
:og_image: http://www.unixstickers.com/image/cache/data/stickers/golang/Go-brown-side.sh-600x600.png
:adsu: yes


There are ParseFiles_ and ParseGlob_ in Go_ `text/template`_ and
`html/template`_ package, but sometimes we want to parse all files in a
directory tree, so I write a *ParseDirectory* function to do this.

The code consists of two part:

1. Use `filepath.Walk`_ to walk the file tree in the directory, and return
   all file paths in the directory.

.. code-block:: go

  // Recursively get all file paths in directory, including sub-directories.
  func GetAllFilePathsInDirectory(dirpath string) ([]string, error) {
  	var paths []string
  	err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
  		if err != nil {
  			return err
  		}
  		if !info.IsDir() {
  			paths = append(paths, path)
  		}
  		return nil
  	})
  	if err != nil {
  		return nil, err
  	}

  	return paths, nil
  }

.. adsu:: 2

2. Pass the file paths to ParseFiles_ function in template package, and the
   parsed template will be returned.

.. code-block:: go

  // Recursively parse all files in directory, including sub-directories.
  func ParseDirectory(dirpath string) (*template.Template, error) {
  	paths, err := GetAllFilePathsInDirectory(dirpath)
  	if err != nil {
  		return nil, err
  	}
  	return template.ParseFiles(paths...)
  }

.. adsu:: 3

Combine the above code, now we have *ParseDirectory* function, which takes
directory path as argument and returns the parsed template.

----

Tested on:

- ``Ubuntu Linux 16.10``
- ``Go 1.7.5``

----

References:

.. [1] `[Golang] Walk All Files in Directory <{filename}../../../2016/02/04/go-walk-all-files-in-directory%en.rst>`_
.. [2] | `golang arguments dot - Google search <https://www.google.com/search?q=golang+arguments+dot>`_
       | `golang arguments dot - DuckDuckGo search <https://duckduckgo.com/?q=golang+arguments+dot>`_
       | `golang arguments dot - Ecosia search <https://www.ecosia.org/search?q=golang+arguments+dot>`_
       | `golang arguments dot - Bing search <https://www.bing.com/search?q=golang+arguments+dot>`_
       | `golang arguments dot - Yahoo search <https://search.yahoo.com/search?p=golang+arguments+dot>`_
       | `golang arguments dot - Baidu search <https://www.baidu.com/s?wd=golang+arguments+dot>`_
       | `golang arguments dot - Yandex search <https://www.yandex.com/search/?text=golang+arguments+dot>`_
.. adsu:: 4
.. [3] `GitHub - siongui/gotemplateutil: utility for Go template <https://github.com/siongui/gotemplateutil>`_

.. _Go: https://golang.org/
.. _html/template: https://golang.org/pkg/html/template/
.. _text/template: https://golang.org/pkg/text/template/
.. _filepath.Walk: https://golang.org/pkg/path/filepath/#Walk
.. _ParseFiles: https://golang.org/pkg/text/template/#ParseFiles
.. _ParseGlob: https://golang.org/pkg/text/template/#ParseGlob

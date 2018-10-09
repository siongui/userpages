[Golang] Download File From URL
###############################

:date: 2018-10-10T04:41+08:00
:tags: Go, Golang, HTTP GET, File Input/Output, Go net/http
:category: Go
:summary: Download and save files (image, pdf, etc.) from given URL in Go.
:og_image: http://www.tricksofit.com/wp-content/uploads/2015/10/Download-File-from-URL.png
:adsu: yes


Download and save files (image, pdf, etc.) from given URL in Go.

.. code-block:: go

  import (
  	"fmt"
  	"io"
  	"net/http"
  	"os"
  	"path"
  )

  func download(url string) (err error) {
  	filename := path.Base(url)
  	fmt.Println("Downloading ", url, " to ", filename)

  	resp, err := http.Get(url)
  	if err != nil {
  		return
  	}
  	defer resp.Body.Close()

  	f, err := os.Create(filename)
  	if err != nil {
  		return
  	}
  	defer f.Close()

  	_, err = io.Copy(f, resp.Body)
  	return
  }

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 18.04``, ``Go 1.10.1``.

----

References:

.. [1] `[Golang] Download HTML From URL <{filename}/articles/2016/03/19/go-download-html-from-url%en.rst>`_
.. [2] `save image from url at DuckDuckGo <https://duckduckgo.com/?q=save+image+from+url>`_
.. [3] `download file from url - Google Search <https://www.google.com/search?q=download+file+from+url>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/

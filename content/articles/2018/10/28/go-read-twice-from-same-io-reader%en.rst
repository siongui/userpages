[Golang] Read Twice From the Same io.Reader
###########################################

:date: 2018-10-28T01:20+08:00
:tags: Go, Golang
:category: Go
:summary: Read twice or multiple times from the same *io.Reader* in *Go*
:og_image: https://vsoch.github.io/assets/images/posts/learning-go/gophercises_jumping.gif
:adsu: yes


This post shows how to solve the problem when you have to read twice or multiple
times from the same reader (io.Reader_) in Go_.

Yesterday whem I was trying to write the code [1]_, I needed to read twice from
the resp.Body_ reader (HTTP response body), but I always got empty result after
function return. Soon I realized that the reader cannot be read more than once.
Otherwise nothing is read.

I made some DuckDuckGo search, and found that io.TeeReader_ may be the solution
to my problem. I tried to use io.TeeReader_ in my code, but still did not get
correct result I wanted.

After some trial and error, finally I came up with my own solution: First read
all bytes from the reader using ioutil.ReadAll_, then use bytes.NewReader_ to
create as many readers as we need.

The following is the example of the my own solution:

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/0jd_-DoQJ0f>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"bytes"
  	"fmt"
  	"io/ioutil"
  	"strings"
  )

  func main() {
  	originalReader := strings.NewReader("test string in test reader")

  	b, err := ioutil.ReadAll(originalReader)
  	if err != nil {
  		panic(err)
  	}

  	reader1 := bytes.NewReader(b)
  	b1, err := ioutil.ReadAll(reader1)
  	if err != nil {
  		panic(err)
  	}
  	fmt.Println(string(b1))

  	reader2 := bytes.NewReader(b)
  	b2, err := ioutil.ReadAll(reader2)
  	if err != nil {
  		panic(err)
  	}
  	fmt.Println(string(b2))
  }

For realistic example, see my yesterday's post [1]_.

----

Tested on:

- ``Ubuntu 18.04``, ``Go 1.11.1``
- `Go Playground`_

----

.. adsu:: 2

References:

.. [1] `[Golang] Auto-Detect and Convert Encoding of HTML to UTF-8 <{filename}/articles/2018/10/27/auto-detect-and-convert-html-encoding-to-utf8-in-go%en.rst>`_
.. [2] `golang io.reader read twice at DuckDuckGo <https://duckduckgo.com/?q=golang+io.reader+read+twice>`_
.. [3] `How to write a proper tar archive in golang? : golang <https://old.reddit.com/r/golang/comments/a2hofj/how_to_write_a_proper_tar_archive_in_golang/>`_
.. [4] `Speeding up Go Script to do text search : golang <https://old.reddit.com/r/golang/comments/a9960c/speeding_up_go_script_to_do_text_search/>`_

.. _Go: https://golang.org/
.. _io.Reader: https://golang.org/pkg/io/#Reader
.. _resp.Body: https://golang.org/pkg/net/http/#Response
.. _io.TeeReader: https://golang.org/pkg/io/#TeeReader
.. _ioutil.ReadAll: https://golang.org/pkg/io/ioutil/#ReadAll
.. _bytes.NewReader: https://golang.org/pkg/bytes/#NewReader
.. _Go Playground: https://play.golang.org/

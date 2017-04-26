Embed Data in Front-end Go Code
###############################

:date: 2017-04-08T23:21+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Web application, web,
       Full-Stack Go, Full-Stack Golang, Embed File in Go Code
:category: GopherJS
:summary: Embed data/assets/resources/binary in Go code, and compiled by
          GopherJS to a single JavaScript file which runs on the front-end
          browser.
:og_image: https://newrelic.com/assets/pages/golang/go-mascot.svg
:adsu: yes

Embed data/assets/resources/binary in Go_ code, and compiled by GopherJS_ to a
single JavaScript file which runs on the front-end browser.

I read interesting discussions on Reddit [2]_ and found that there are several
packages which let developers embed data/assets/resources/binary in the Go code.
Pakcages such as fileb0x_, go-bindata_, and go.rice_ are mentioned in the Reddit
post, and I made some survey and decided to give fileb0x_ a try.

go-bindata_ has more stars and popular on GitHub. The reason why I do not choose
go-bindata is that the way that go-bindata embed data in the code is very
similar to the way I do in my current project. I use the template package in Go
standard library to generate and embed data in Go code. fileb0x_, however, embed
the data in a separate package that we can *import* and use in our code, which
makes the code more clean, readable, and easy to understand. This is the main
reason I choose *fileb0x* instead of *go-bindata*.

Now I will show how I embed data in code that runs on front-end browser. I have
the following data of `succinct trie`_ to be embedded in the code:

.. code-block:: txt

  108K	website/rd.txt
  4.0K	website/strie_node_count.txt
  1.5M	website/strie.txt

Install fileb0x_ first:

.. code-block:: bash

  $ go get -u github.com/UnnoTed/fileb0x

.. adsu:: 2

Next, create a config file to tell *fileb0x* where to find the data and build a
Go package which has the data embedded in the code. The name of output Go
package is *paliDataVFS*. For more detail of config file, see
`yaml config example`_.

.. code-block:: yaml

  pkg: paliDataVFS
  dest: "github/repo/paliDataVFS/"

  fmt: false

  compression:
    compress: false
    method: ""
    keep: false

  clean: false
  output: "ab0x.go"
  unexporTed: false
  spread: false
  debug: false

  custom:
    - files:
      - "./website/rd.txt"
      - "./website/strie_node_count.txt"
      - "./website/strie.txt"

      base: "./website/"
      prefix: "public/"

Save the config as ``vfsb0x.yaml``, and run *fileb0x* to create Go package:

.. code-block:: bash

  $ fileb0x vfsb0x.yaml

The output will be ``github/repo/paliDataVFS/ab0x.go``, which contains the data
in it. Then I create a repo in GitHub, which is
``github.com/siongui/paliDataVFS``, and commit the ``ab0x.go`` to the repo.

.. adsu:: 3

Install the package created by *fileb0x*:

.. code-block:: bash

  $ go get -u github.com/siongui/paliDataVFS

Now we can access the data in our front-end Go code as follows:

.. code-block:: go

  import (
  	"github.com/siongui/paliDataVFS"
  	"strconv"
  )

  func GetTrieDataFromVFS() (succinctTrieDataBlob string, rankDirectoryDataBlob string, succinctTrieNodeCount uint) {
  	tmp, _ := paliDataVFS.ReadFile("public/strie.txt")
  	succinctTrieDataBlob = string(tmp)
  	tmp, _ = paliDataVFS.ReadFile("public/rd.txt")
  	rankDirectoryDataBlob = string(tmp)
  	tmp, _ = paliDataVFS.ReadFile("public/strie_node_count.txt")
  	tmp2, _ := strconv.ParseUint(string(tmp), 10, 64)
  	succinctTrieNodeCount = uint(tmp2)
  	return
  }

After coding finished, use GopherJS_ to compile the code to JavaScript!
It works!!!

----

Tested on: ``Ubuntu Linux 16.10``, ``Go 1.8.1``,
``Chromium Version 57.0.2987.98 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``.

----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)
.. [2] | `Is including assets (with a tool like go-bindata) an anti-pattern? : golang <https://www.reddit.com/r/golang/comments/60166q/is_including_assets_with_a_tool_like_gobindata_an/>`_
       | `How to build Go plugin with data inside : golang <https://www.reddit.com/r/golang/comments/63f3ag/how_to_build_go_plugin_with_data_inside/>`_
       | `golang - compile static files in app? : golang <https://www.reddit.com/r/golang/comments/66uewv/golang_compile_static_files_in_app/>`_
.. adsu:: 4
.. [3] `GitHub - UnnoTed/fileb0x: simple customizable tool to embed files in go <https://github.com/UnnoTed/fileb0x>`_
.. [4] `GitHub - jteeuwen/go-bindata: A small utility which generates Go code from any file. Useful for embedding binary data in a Go program. <https://github.com/jteeuwen/go-bindata>`_
.. [5] `GitHub - GeertJohan/go.rice: go.rice is a Go package that makes working with resources such as html,js,css,images,templates, etc very easy. <https://github.com/GeertJohan/go.rice>`_
.. [6] `GitHub - inconshreveable/go-update: Build self-updating Golang programs <https://github.com/inconshreveable/go-update>`_
.. [7] `GitHub - shurcooL/vfsgen: Takes an input http.FileSystem (likely at go generate time) and generates Go code that statically implements it. <https://github.com/shurcooL/vfsgen>`_
.. [8] | `embed data in go - Google search <https://www.google.com/search?q=embed+data+in+go>`_
       | `embed data in go - DuckDuckGo search <https://duckduckgo.com/?q=embed+data+in+go>`_
       | `embed data in go - Ecosia search <https://www.ecosia.org/search?q=embed+data+in+go>`_
       | `embed data in go - Qwant search <https://www.qwant.com/?q=embed+data+in+go>`_
       | `embed data in go - Bing search <https://www.bing.com/search?q=embed+data+in+go>`_
       | `embed data in go - Yahoo search <https://search.yahoo.com/search?p=embed+data+in+go>`_
       | `embed data in go - Baidu search <https://www.baidu.com/s?wd=embed+data+in+go>`_
       | `embed data in go - Yandex search <https://www.yandex.com/search/?text=embed+data+in+go>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _GopherJS: http://www.gopherjs.org/
.. _fileb0x: https://github.com/UnnoTed/fileb0x
.. _go-bindata: https://github.com/jteeuwen/go-bindata
.. _go.rice: https://github.com/GeertJohan/go.rice
.. _succinct trie: https://github.com/siongui/go-succinct-data-structure-trie
.. _yaml config example: https://github.com/UnnoTed/fileb0x/blob/master/_example/simple/b0x.yaml
.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js

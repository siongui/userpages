[Golang] Filename Globbing Summary
##################################

:date: 2017-04-25T17:03+08:00
:tags: Go, Golang, List Files in Directory
:category: Go
:summary: Test how filename globbing works in Go standard library.
:adsu: yes

Filename globbing are found in several places in Go standard library.

- filepath.Glob_
- filepath.Match_
- path.Match_
- template.ParseGlob_

I want to how the it works, so I use a small example to test. I have the two txt
files as follows:

.. code-block:: txt

  testdir/a.txt
  textdir/subdir/b.txt

Use ``testdir/*.txt`` as pattern in filepath.Glob_:

.. code-block:: go

  package main

  import (
  	"path/filepath"
  	"fmt"
  )

  func main() {
  	matches, err := filepath.Glob("testdir/*.txt")
  	if err != nil {
  		panic(err)
  	}

  	fmt.Println(matches)
  }

The output:

.. code-block:: txt

  [testdir/a.txt]

As expected, ``*`` does not match sub-directory. How about ``**``?
Modify the pattern from ``testdir/*.txt`` to ``testdir/**.txt``, and the output
is:

.. code-block:: txt

  [testdir/a.txt]

The output is the same. ``**`` does not match zero or more directories [3]_.

If sub-directory depth is known, we can still match file in sub-directory. For
example, modify the pattern from ``testdir/*.txt`` to ``testdir/*/*.txt`` [4]_.
The output is:

.. code-block:: txt

  [testdir/subdir/b.txt]

The file in sub-directory is matched.

.. adsu:: 2

----

My summary

- ``*`` does not match sub-directories.
- ``**`` is not supported, i.e., does not match zero or more directories.
- If depth of sub-directories is known, sub-directories can be matched with
  workaround.
- There are third-party packages which provides more complete glob features
  [5]_.

----

Tested on:

- Ubuntu Linux 17.04
- Go 1.8.1

----

References:

.. [1] | `glob vs regex - Google search <https://www.google.com/search?q=glob+vs+regex>`_
       | `glob vs regex - DuckDuckGo search <https://duckduckgo.com/?q=glob+vs+regex>`_
       | `glob vs regex - Ecosia search <https://www.ecosia.org/search?q=glob+vs+regex>`_
       | `glob vs regex - Qwant search <https://www.qwant.com/?q=glob+vs+regex>`_
       | `glob vs regex - Bing search <https://www.bing.com/search?q=glob+vs+regex>`_
       | `glob vs regex - Yahoo search <https://search.yahoo.com/search?p=glob+vs+regex>`_
       | `glob vs regex - Baidu search <https://www.baidu.com/s?wd=glob+vs+regex>`_
       | `glob vs regex - Yandex search <https://www.yandex.com/search/?text=glob+vs+regex>`_

.. [2] | `golang glob match - Google search <https://www.google.com/search?q=golang+glob+match>`_
       | `golang glob match - DuckDuckGo search <https://duckduckgo.com/?q=golang+glob+match>`_
       | `golang glob match - Ecosia search <https://www.ecosia.org/search?q=golang+glob+match>`_
       | `golang glob match - Qwant search <https://www.qwant.com/?q=golang+glob+match>`_
       | `golang glob match - Bing search <https://www.bing.com/search?q=golang+glob+match>`_
       | `golang glob match - Yahoo search <https://search.yahoo.com/search?p=golang+glob+match>`_
       | `golang glob match - Baidu search <https://www.baidu.com/s?wd=golang+glob+match>`_
       | `golang glob match - Yandex search <https://www.yandex.com/search/?text=golang+glob+match>`_

.. [3] `path/filepath: Glob should support \`**\` for zero or more directories · Issue #11862 · golang/go · GitHub <https://github.com/golang/go/issues/11862>`_
.. [4] `filebeat wildcard for directories · Issue #2084 · elastic/beats · GitHub <https://github.com/elastic/beats/issues/2084#issuecomment-252105586>`_
.. [5] | `glob - Go libraries and apps <https://golanglibs.com/top?q=glob>`_
       | `GitHub - gobwas/glob: Go glob <https://github.com/gobwas/glob>`_
.. [6] `Wildcards - GNU/Linux Command-Line Tools Summary <http://tldp.org/LDP/GNU-Linux-Tools-Summary/html/x11655.htm>`_
.. [7] `[Golang] Walk All Files in Directory <{filename}../../../2016/02/04/go-walk-all-files-in-directory%en.rst>`_
.. [8] | `Glob for go. Works much faster than regexp on equivalent patterns. : golang <https://www.reddit.com/r/golang/comments/41ulfq/glob_for_go_works_much_faster_than_regexp_on/>`_
       | `Glob Matching Can Be Simple and Fast Too | Hacker News <https://news.ycombinator.com/item?id=14184528>`_
       | `research!rsc: Glob Matching Can Be Simple And Fast Too : golang <https://www.reddit.com/r/golang/comments/67by6g/researchrsc_glob_matching_can_be_simple_and_fast/>`_
       | `research!rsc: Glob Matching Can Be Simple And Fast Too <https://research.swtch.com/glob>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _filepath.Glob: https://golang.org/pkg/path/filepath/#Glob
.. _filepath.Match: https://golang.org/pkg/path/filepath/#Match
.. _path.Match: https://golang.org/pkg/path/#Match
.. _template.ParseGlob: https://golang.org/pkg/text/template/#ParseGlob

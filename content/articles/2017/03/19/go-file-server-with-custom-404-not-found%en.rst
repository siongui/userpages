[Golang] File Server With Custom 404 Not Found
##############################################

:date: 2017-03-19T04:21+08:00
:tags: Go, Golang, Go net/http, HTTP 404 Not Found
:category: Go
:summary: Implement static file server with custom HTTP 404 Not Found handler
          via Go standard *net/http* package.
:og_image: https://i.ytimg.com/vi/z74PQXHnzjk/hqdefault.jpg
:adsu: yes


Implementation of static file server with custom `HTTP 404 Not Found`_ handler
via Go standard `net/http`_ package.

The built-in FileServer_ method in standard *net/http* package can serve staic
files in the file syatem, but does not allow custom 404 Not Found handler, so I
implement my own FileServer method with custom 404 handler.

The idea of implementation of FileServer with custom 404 handler comes from the
implementation of StripPrefix_ method in the same *net/http* package. You can
read the code of *StripPrefix* method for more details.

The key idea of my implementation:

- Accept the same parameters (FileSystem_) as *FileServer* method, which is
  return value of *http.Dir* in normal case.
- Use the *FileSystem.Open* and *os.IsNotExist* methods to check if the file
  exists. If not, call custom 404 handler. Otherwise serve the file with normal
  *FileServer* handler.

.. adsu:: 2

**Source Code**:

.. show_github_file:: siongui userpages content/code/go/custom-404/server.go
.. adsu:: 3

Not sure whether this implementation is robust or not, but it works well in my
application.

----

Tested on:

- ``Ubuntu Linux 16.10``
- ``Go 1.8``

----

References:

.. [1] | `golang custom 404 - Google search <https://www.google.com/search?q=golang+custom+404>`_
       | `golang custom 404 - DuckDuckGo search <https://duckduckgo.com/?q=golang+custom+404>`_
       | `golang custom 404 - Ecosia search <https://www.ecosia.org/search?q=golang+custom+404>`_
       | `golang custom 404 - Qwant search <https://www.qwant.com/?q=golang+custom+404>`_
       | `golang custom 404 - Bing search <https://www.bing.com/search?q=golang+custom+404>`_
       | `golang custom 404 - Yahoo search <https://search.yahoo.com/search?p=golang+custom+404>`_
       | `golang custom 404 - Baidu search <https://www.baidu.com/s?wd=golang+custom+404>`_
       | `golang custom 404 - Yandex search <https://www.yandex.com/search/?text=golang+custom+404>`_

.. [2] | `golang file server 404 - Google search <https://www.google.com/search?q=golang+file+server+404>`_
       | `golang file server 404 - DuckDuckGo search <https://duckduckgo.com/?q=golang+file+server+404>`_
       | `golang file server 404 - Ecosia search <https://www.ecosia.org/search?q=golang+file+server+404>`_
       | `golang file server 404 - Qwant search <https://www.qwant.com/?q=golang+file+server+404>`_
       | `golang file server 404 - Bing search <https://www.bing.com/search?q=golang+file+server+404>`_
       | `golang file server 404 - Yahoo search <https://search.yahoo.com/search?p=golang+file+server+404>`_
       | `golang file server 404 - Baidu search <https://www.baidu.com/s?wd=golang+file+server+404>`_
       | `golang file server 404 - Yandex search <https://www.yandex.com/search/?text=golang+file+server+404>`_

.. [3] `func StripPrefix - net/http - The Go Programming Language <https://golang.org/pkg/net/http/#StripPrefix>`_
.. [4] `file server with custom 404 · siongui/pali@e872a78 · GitHub <https://github.com/siongui/pali/commit/e872a787647a3e0d7294c75d4ce28d6f9988b6ce>`_
.. [5] `How to serve static files with custom NotFound handler : golang <https://www.reddit.com/r/golang/comments/64230v/how_to_serve_static_files_with_custom_notfound/>`_
       `How to serve static files with custom NotFound handler - Blog from kefaise.com <http://kefblog.com/2017-04-07/How-to-serve-static-files-with-custom-not-found-handler>`_

.. _HTTP 404 not found: https://www.google.com/search?q=HTTP+404+not+found
.. _net/http: https://golang.org/pkg/net/http/
.. _FileServer: https://golang.org/pkg/net/http/#FileServer
.. _StripPrefix: https://golang.org/pkg/net/http/#StripPrefix
.. _FileSystem: https://golang.org/pkg/net/http/#FileSystem

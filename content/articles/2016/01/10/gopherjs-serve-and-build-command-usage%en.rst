[Golang] GopherJS serve and build Command Usage
###############################################

:date: 2016-01-10T19:14+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Commandline, Continuous Integration
:category: GopherJS
:summary: Example of how to use *serve* and *build* command of GopherJS_.
:adsu: yes


GopherJS provide command line tools but there is no enough documentation of the
command. I gave the *serve* and *build* command a try and give a summary of the
usage.

First we see the *help* of *gopherjs* command line tool:

.. code-block:: bash

  $ gopherjs help
  GopherJS is a tool for compiling Go source code to JavaScript.

  Usage:
    gopherjs [command]

  Available Commands:
    build       compile packages and dependencies
    get         download and install packages and dependencies
    install     compile and install packages and dependencies
    run         compile and run Go program
    test        test packages
    tool        run specified go tool
    serve       compile on-the-fly and serve

  Flags:
        --color         colored output (default true)
    -m, --minify        minify generated code
    -q, --quiet         suppress non-fatal warnings
        --tags string   a list of build tags to consider satisfied during the build
    -v, --verbose       print the names of packages as they are compiled
    -w, --watch         watch for changes to the source files

  Use "gopherjs [command] --help" for more information about a command.

.. adsu:: 2

Build Command Usage
+++++++++++++++++++

We will try *build* command first. See the help of *build* command first:

.. code-block:: bash

  $ gopherjs build --help
  compile packages and dependencies

  Usage:
    gopherjs build [packages] [flags]

  Flags:
        --color           colored output (default true)
    -m, --minify          minify generated code
    -o, --output string   output file
    -q, --quiet           suppress non-fatal warnings
        --tags string     a list of build tags to consider satisfied during the build
    -v, --verbose         print the names of packages as they are compiled
    -w, --watch           watch for changes to the source files

Now I have a Go_ program named ``dom.js``, put under ``.`` directory. Compile
this file to JavaScript by:

.. code-block:: bash

  $ gopherjs build dom.go -o dom.js

The JavaScript output file will be ``dom.js``, put under ``.`` directory.

Next we try the **-w** flag of *build* command:

.. code-block:: bash

  $ gopherjs build dom.go -w -o dom.js

The GopherJS command line tool will watch the directory, and if you make any
changes of ``dom.go``, it will automatically recompile the Go_ code to
JavaScript. This feature is great for daily development.

Assume *GOPATH* is set, and you put the ``dom.go`` under ``$(GOPATH)/src/demo``
directory, you can tell GopherJS command line tool to look for the *demo*
directory by:

.. code-block:: bash

  $ gopherjs build demo -w -o src/demo/dom.js

The command line tool will compile the ``dom.go`` under ``$(GOPATH)/src/demo``
and the JavaScript output file will be ``$(GOPATH)/src/demo/dom.js``.

.. adsu:: 3

Serve Command Usage
+++++++++++++++++++

See the *help* of *serve* command first:

.. code-block:: bash

  $ gopherjs serve --help
  compile on-the-fly and serve

  Usage:
    gopherjs serve [flags]

  Flags:
        --color         colored output (default true)
        --http string   HTTP bind address to serve (default ":8080")
    -m, --minify        minify generated code
    -q, --quiet         suppress non-fatal warnings
        --tags string   a list of build tags to consider satisfied during the build
    -v, --verbose       print the names of packages as they are compiled

  Global Flags:
    -w, --watch   watch for changes to the source files

Run the command without any flag:

.. code-block:: bash

  $ gopherjs serve

The GopherJS command line tool will serve ``$(GOPATH)/src`` directory by
default. It looks like there is no way to change the serving directory. Open
your browser at ``http://localhost:8080`` to visit the webpage.

I like to visit the webpage at ``http://localhost:8000``. Change the port by:

.. code-block:: bash

  $ gopherjs serve --http ":8000"

I try to run the *serve* command with **-w** flag but it looks like the command
line tool did not watch the changes and recompile for me. So I guess the **-w**
flag is useless combined with *serve* command.

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.5.2``.

----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] `Getting Started with GopherJS <https://www.hakkalabs.co/articles/getting-started-gopherjs>`_

.. [3] `GopherJSの紹介 - GolangRdyJp <http://golang.rdy.jp/2015/10/15/gopherjs/>`_

.. [4] `albrow/gopherjs-live · GitHub <https://github.com/albrow/gopherjs-live>`_
       (Automatic watching and recompiling for gopherjs)

.. [5] `ajhager/srvi · GitHub <https://github.com/ajhager/srvi>`_
       (Quickly build, serve, run, and refresh your GopherJS programs)

.. [6] `cmd/gopherjs_serve_html at master · shurcooL/cmd · GitHub <https://github.com/shurcooL/cmd/tree/master/gopherjs_serve_html>`_

.. [7] `Add "gopherjs serve" command · Issue #121 · gopherjs/gopherjs · GitHub <https://github.com/gopherjs/gopherjs/issues/121>`_

.. [8] `It's easy to get an infinite loop with the watch flag · Issue #212 · gopherjs/gopherjs · GitHub <https://github.com/gopherjs/gopherjs/issues/212>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _GopherJS: http://www.gopherjs.org/

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js

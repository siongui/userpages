[Golang] Compile SASS/SCSS files to CSS via libsass
###################################################

:date: 2016-01-28T05:30+08:00
:tags: CSS, SASS, SCSS, Go, Golang
:category: Go
:summary: Write a Go_ program to compile your SASS_/SCSS_ (CSS_ extension
          language) files to CSS via libsass_ and `Go wrapper for libsass`_.
:adsu: yes


Write a Golang_ program to compile your SASS_/SCSS_ (CSS_ extension language)
files to CSS via libsass_ and `Go wrapper for libsass`_.

Download Makefile_ for demo (remember to modify GOROOT and GOPATH in Makefile):

.. show_github_file:: siongui userpages content/code/go-sass/Makefile

Install `go-libsass`_ (You do not need to install libsass_):

.. code-block:: bash

  # cd to the directory where you put Makefile
  $ make install

Download demo SCSS_ files:

.. code-block:: bash

  # cd to the directory where you put Makefile
  $ make download

.. adsu:: 2

Now we write a Go_ program to compile SCSS files to CSS (Put this Go file in the
same directory as Makefile):

.. show_github_file:: siongui userpages content/code/go-sass/testsass.go

Run the Go_ program:

.. code-block:: bash

  # cd to the directory where you put Makefile
  $ make

The output file will be *style.css* in the same directory.

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.5.3``.

----

Reference:

.. [1] `wellington/go-libsass: Go wrapper for libsass, the only Sass 3.3 compiler for Go <https://github.com/wellington/go-libsass>`_

.. [2] `golang os write file <https://www.google.com/search?q=golang+os+write+file>`_

.. [3] `go - How to read/write from/to file? - Stack Overflow <http://stackoverflow.com/questions/1821811/how-to-read-write-from-to-file>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _libsass: https://github.com/sass/libsass
.. _SASS: http://sass-lang.com/
.. _SCSS: http://sass-lang.com/documentation/file.SCSS_FOR_SASS_USERS.html
.. _CSS: https://www.google.com/search?q=css
.. _Go wrapper for libsass: https://github.com/wellington/go-libsass
.. _go-libsass: https://github.com/wellington/go-libsass
.. _Makefile: https://www.google.com/search?q=makefile

[Makefile] Check if Run on Travis CI
####################################

:date: 2017-06-13T23:06+08:00
:tags: Bash, Makefile, Commandline
:category: Bash
:summary: Check if Makefile runs on Travis CI or not.
:adsu: yes


I want my Makefile_ to be used both on local development and Travis CI build.
The programming language is Go, and I need to export *GOROOT*, *GOPATH*, and
*PATH* during local devlopement. The build environment on Travis CI has these
environment variables already set, so I want to check if the running environment
is Travis CI or not. If it is, do not export these variables.

I searched the Travis CI environment variables [1]_, and found that *TRAVIS* env
is set to *true* in Travis CI build. We can export variables only on Travis CI
build as follows:

.. code-block:: makefile

  ifndef TRAVIS
  	export GOROOT=$(realpath ../go)
  	export GOPATH=$(realpath .)
  	export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)
  endif

*ifndef* means *if not defined*, so if *TRAVIS* is not defined, which means it's
local development environment, we will export *GOROOT*, *GOPATH*, and *PATH*.

----

Tested on `Travis CI`, `Ubuntu Linux 17.04`_

----

References:

.. [1] `Environment Variables - Travis CI <https://docs.travis-ci.com/user/environment-variables/>`_

.. _Makefile: https://www.google.com/search?q=Makefile
.. _Ubuntu Linux 17.04: http://releases.ubuntu.com/17.04/
.. _GNU make 4.1-9: https://www.gnu.org/software/make/

===========================
Dart Toggle Sidebar Example
===========================

Inspired by `Octopress <http://octopress.org/>`_, create toggle-able
sidebar with `Dart <https://www.dartlang.org/>`_.

Setup Development Environment
=============================

1. Development environment is Ubuntu 13.10.

2. Download `Dart Zip file <https://www.dartlang.org/>`_.

3. Modify the first line in Makefile to point to location
   of your unzipped Dart directory. For example, your unzipped
   Dart dir is located at ``../dart``, you should set the first
   line like this:

   .. code-block:: Makefile

     DART_DIR=../dart

4. Run the demo:

   .. code-block:: bash

     $ make

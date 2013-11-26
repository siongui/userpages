Dart Development Using Vim
##########################

:date: 2013-11-27 02:22
:tags: Dart
:category: Dart
:author: Siong-Ui Te
:summary: Develop Dart application using Vim Editor

1. Use `Vundle <https://github.com/gmarik/vundle>`_ to install
`Dart plugin for VIM <https://github.com/dart-lang/dart-vim-plugin>`_.

2. Check `dart: The standalone Dart VM <https://www.dartlang.org/docs/dart-up-and-running/contents/ch04-tools-dart-vm.html>`_
to see how to run Dart command-line (server-side) apps.

3. Check `Dartium: Chromium with the Dart VM <https://www.dartlang.org/docs/dart-up-and-running/contents/ch04-tools-dartium.html>`_
to see how to run Dart web apps. For Ubuntu 13.10, missing libudev.so.0 occurs
while Chromium startup. Check `here <http://askubuntu.com/questions/369310/how-to-fix-missing-libudev-so-0-for-chrome-to-start-again>`_
for solution.

The following is the reference Makefile for automating the process 
of running apps.

.. code-file:: Makefile


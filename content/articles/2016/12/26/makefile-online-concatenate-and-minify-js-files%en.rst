[Makefile] Online Concatenate and Compress JavaScript Files
###########################################################

:date: 2016-12-26T21:33+08:00
:tags: Bash, JavaScript, Commandline, HTTP POST, Google Closure,
       Minify HTML/CSS/JavaScript, curl, cat command, Makefile
:category: Bash
:summary: Concatenate and minify JavaScript_ files via Makefile_, curl_, and
          online `Google Closure Compiler`_.
:adsu: yes


Concatenate and minify/compress JavaScript_ files via Makefile_, curl_, and
online `Google Closure Compiler`_.

Put this Makefile_ and js files in the same directory and then type ``make``.
After finish, the minified/compressed js file will be named ``app.min.js`` in
the same directory.

.. show_github_file:: siongui userpages content/code/make/google-closure-minify-js/Makefile

----

Tested on: ``Ubuntu Linux 16.10``.

----

References:

.. [1] `[Bash] Online Concatenate and Compress JavaScript Files <{filename}../25/bash-online-concatenate-and-minify-js-files%en.rst>`_

.. [2] `[Makefile] Concetenate and Minify CSS via sed and tr Command <{filename}../../03/11/makefile-concetenate-and-minify-css-via-sed-and-tr%en.rst>`_


.. _JavaScript: https://www.google.com/search?q=javascript
.. _Makefile: https://www.google.com/search?q=Makefile
.. _curl: https://www.google.com/search?q=curl
.. _Google Closure Compiler: https://developers.google.com/closure/compiler/

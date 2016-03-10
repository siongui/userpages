[Makefile] Concetenate and Minify CSS via sed and tr Command
############################################################

:date: 2016-03-11T04:52+08:00
:tags: Bash, Makefile, Commandline, sed, tr command, String Manipulation,
       Regular Expression, remove trailing newline, remove carriage return,
       Minify HTML/CSS/JavaScript
:category: Bash
:summary: Concatenate and Minify CSS via sed_ and tr_ command in Makefile_.


Concatenate and Minify CSS via sed_ and tr_ command in Makefile_.

.. code-block:: bash

  SOURCE_CSS_FILES=$(filter-out %.min.css, $(wildcard *.css))
  MINIFIED_CSS=app.min.css

  minify: concat
  	@# remove css comments
  	sed -r ':a; s%(.*)/\*.*\*/%\1%; ta; /\/\*/ !b; N; ba' -i $(MINIFIED_CSS)
  	@# remove leading spaces and tabs
  	sed 's/^\s*//' -i $(MINIFIED_CSS)
  	@# remove trailing spaces, tabs, and newline
  	sed 's/\s*$$//' -i $(MINIFIED_CSS)
  	@# remove newline
  	tr --delete '\n' < $(MINIFIED_CSS) > tmp.css
  	mv tmp.css $(MINIFIED_CSS)

  concat:
  	cat $(SOURCE_CSS_FILES) > $(MINIFIED_CSS)

----

Tested on ``Ubuntu Linux 15.10``.

----

References:

.. [1] `GNU make: Text Functions <https://www.gnu.org/software/make/manual/html_node/Text-Functions.html>`_

.. [2] `GNU make: Wildcard Function <https://www.gnu.org/software/make/manual/html_node/Wildcard-Function.html>`_

.. [3] `sed remove comments c <https://www.google.com/search?q=sed+remove+comments+c>`_

       `sed - Remove multi-line comments - Stack Overflow <http://stackoverflow.com/questions/13061785/remove-multi-line-comments>`_

.. [4] `unix - how to remove leading whitespace from each line in a file? - Stack Overflow <http://stackoverflow.com/questions/2310605/how-to-remove-leading-whitespace-from-each-line-in-a-file>`_

.. [5] `bash - How to use sed in a Makefile - Stack Overflow <http://stackoverflow.com/questions/3140974/how-to-use-sed-in-a-makefile>`_

.. [6] `unix - How can I replace a newline (\n) using sed? - Stack Overflow <http://stackoverflow.com/questions/1251999/how-can-i-replace-a-newline-n-using-sed>`_

.. [7] `makefile instead of grunt <https://www.google.com/search?q=makefile+instead+of+grunt>`_

       `What's in a Build Tool? (lihaoyi.com) <http://www.lihaoyi.com/post/WhatsinaBuildTool.html>`_
       (`HN discussions <https://news.ycombinator.com/item?id=11222967>`__)

       `ocaml-9p/Makefile at master · mirage/ocaml-9p · GitHub <https://github.com/mirage/ocaml-9p/blob/master/Makefile>`_

       `rappel/Makefile at master · yrp604/rappel · GitHub <https://github.com/yrp604/rappel/blob/master/Makefile>`_

       `In defense of Unix (leancrew.com) <http://leancrew.com/all-this/2016/03/in-defense-of-unix/>`_
       (`HN discussions <https://news.ycombinator.com/item?id=11229025>`__)

.. [8] `makefile check if symlink exists <https://www.google.com/search?q=makefile+check+if+symlink+exists>`_

.. [9] `makefile concatenate files <https://www.google.com/search?q=makefile+concatenate+files>`_

       `javascript - Makefile to combine js files and make a compressed version - Stack Overflow <http://stackoverflow.com/questions/4413903/makefile-to-combine-js-files-and-make-a-compressed-version>`_

       `build - Is there a way to exclude certain source files or folders from a makefile? - Stack Overflow <http://stackoverflow.com/questions/1531318/is-there-a-way-to-exclude-certain-source-files-or-folders-from-a-makefile>`_

.. [10] `concatenate and minify css via make/sed/tr · siongui/pali@188d66f · GitHub <https://github.com/siongui/pali/commit/188d66f704552b9c6e6fa5f0a7bb79d4b8b77524>`_

.. [11] `tr replace file <https://www.google.com/search?q=tr+replace+file>`_

.. _Makefile: https://www.google.com/search?q=Makefile
.. _sed: http://www.grymoire.com/Unix/Sed.html
.. _tr: http://www.linfo.org/tr.html

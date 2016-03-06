Google Closure Library Editor with Equation Editor Dialog
#########################################################

:date: 2012-07-08 20:57
:modified: 2015-04-06 21:16
:tags: JavaScript, Web application, web, html, math symbol, online latex editor,
       LaTeX, Google Closure
:category: JavaScript
:summary: Rich text editor with math equation dialog using
          `Google Closure Library`_.


Recently I have an application which needs a HTML editor with LaTex_ equation
support. There are many `WYSIWYG (WhatYou See Is What You Get)`_ editors
available on Internet, and finally I decide to use `Google Closure Library`_
built-in rich text editor (see [1]_).

The default closure works perfectly, and there are also demos of equation editor
dialog. But there are two problem:

  1. It is not trivial to embed the equation editor dialog into default rich
     text editor, and lack of documentation on this part. (actually it's trivial
     to embed assumed that you are familiar with the code structure of closure
     library, but I am not in the beginning.)

  2. The preview feature of math equation in equation editor dialog is broken.
     We have to find some way to fix this.

To fix the issue #1, I read some material online (References [2]_ ~ [7]_) and
trace the `online source code`_ and api_. To fix the issue #2, I read the
references [9]_ ~ [11]_, and I decide to use the solution provided by [11]_. The
following is my patch for the closure library (under `closure/goog/`_, revision
2021_):

.. show_github_file:: siongui userpages content/articles/2012/07/08/patch.diff

As you can see, it's not so difficult to customize the library, but it takes
time to be familiar with the code structure. I also patch default `editor demo`_
to make it include equation editor dialog for reference. I hope this post will
help those who need a rich text editor with math equation.

----

References:

.. [1] `Introducing the Closure Library Editor - Closure Tools Blog <http://closuretools.blogspot.com/2010/07/introducing-closure-library-editor.html>`_

.. [2] Google Search: `google closure editor plugin <https://www.google.com/search?q=google+closure+editor+plugin>`_

.. [3] `Write a Google Closure Editor Plugin <http://www.slideshare.net/yinhm/plugin-6345064>`_

.. [4] `yinhm/google-closure-editor-image · GitHub <https://github.com/yinhm/google-closure-editor-image>`_

.. [5] `The Road to HTML 5: contentEditable <https://blog.whatwg.org/the-road-to-html-5-contenteditable>`_

.. [6] `Rich HTML editing in the browser: part 1 <http://dev.opera.com/articles/view/rich-html-editing-in-the-browser-part-1/>`_
       (`backup <https://github.com/operasoftware/devopera-static-backup/tree/master/http/dev.opera.com/articles/view/rich-html-editing-in-the-browser-part-1>`__)

.. [7] `Rich HTML editing in the browser: part 2 <http://dev.opera.com/articles/view/rich-html-editing-in-the-browser-part-2/>`_
       (`backup <https://github.com/operasoftware/devopera-static-backup/tree/master/http/dev.opera.com/articles/view/rich-html-editing-in-the-browser-part-2>`__)

.. [8] `Updated richtext editor to include HTML source editing <https://code.google.com/p/cruise-control-for-app-engine/source/detail?r=32>`_

.. [9] `jqMath - Put Math on the Web <http://mathscribe.com/author/jqmath.html>`_

.. [10] `javascript - Embeddable WYSIWYG equation editor - Stack Overflow <http://stackoverflow.com/questions/7433540/embeddable-wysiwyg-equation-editor>`_

.. [11] `LaTeX Equations in HTML | CodeCogs Equation Editor <http://www.codecogs.com/latex/integration/htmlequations.php>`_


.. _LaTeX: http://en.wikipedia.org/wiki/LaTeX

.. _WYSIWYG (WhatYou See Is What You Get): http://en.wikipedia.org/wiki/WYSIWYG

.. _Google Closure Library: https://developers.google.com/closure/library/

.. _online source code: https://code.google.com/p/closure-library/

.. _api: http://docs.closure-library.googlecode.com/git/index.html

.. _closure/goog/: https://github.com/google/closure-library/tree/master/closure/goog

.. _2021: https://code.google.com/p/closure-library/source/detail?r=2021

.. _editor demo: https://github.com/google/closure-library/blob/master/closure/goog/demos/editor/editor.html

.. _Google Closure Library: https://developers.google.com/closure/library/

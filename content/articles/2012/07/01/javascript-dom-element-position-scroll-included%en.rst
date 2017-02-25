JavaScript DOM Element Position (Scroll Position Included)
##########################################################

:date: 2012-07-01 14:13
:modified: 2015-02-18 21:05
:tags: JavaScript, element offset, element position, DOM
:category: JavaScript
:summary: Get DOM_ element position (including scroll position)
          using JavaScript_.
:og_image: http://www.javatpoint.com/images/javascript/javascript_logo.png
:adsu: yes


To detect the position of a DOM_ element (for example, HTML_ *div* tag or *p*
tag) in JavaScript_ is an interesting topic. In [1]_, [2]_, [3]_, and [4]_,
there are discussion and code for detecting the position of a DOM_ elemenet. In
[5]_, [6]_, [7]_, there are discussion and code for detecting the scroll
position of a DOM_ element. In this post, sample code for detecting the DOM_
element position, including scroll position, will be shown.

.. rubric:: `Demo <{filename}position.html>`_
   :class: align-center

Source Code for Demo (*HTML*):

.. show_github_file:: siongui userpages content/articles/2012/07/01/position.html
.. adsu:: 2

Source Code for Demo (*JavaScript*):

.. show_github_file:: siongui userpages content/articles/2012/07/01/position.js
.. adsu:: 3

The sample code above, however, is not complete. The code will not work if
`CSS margin property`_ and absolute `display property`_ of a DOM element is set.
For example, if the website designer uses the semi-fluid layout mentioned in
this `holy grail tutorial`_, the code above will not work because the margin and
relative position is not taken into consideration in the sample code. The
position function provided by jQuery_ in [11]_ will still work under such
circumstance. I tried to trace the source code of jQuery_ in [10]_, but still
cannot figure out how to do it. Maybe I will be back to this issue some time.

----

References:

.. [1] `getBoundingClientRect method - Dottoro Web Reference <http://help.dottoro.com/ljvmcrrn.php>`_

.. [2] `Retrieve the position (X,Y) of an HTML element <http://stackoverflow.com/questions/442404/retrieve-the-position-x-y-of-an-html-element>`_

.. [3] `offsetLeft property JavaScript - Dottoro Web Reference <http://help.dottoro.com/ljajgavt.php>`_

.. [4] `Get real offset left/top of element with inline offsetparent <http://stackoverflow.com/questions/5459894/get-real-offset-left-top-of-element-with-inline-offsetparent>`_

|

.. [5] `JavaScript tutorial - Window size and scrolling <http://www.howtocreate.co.uk/tutorials/javascript/browserwindow>`_

.. [6] `scrollLeft property - Dottoro Web Reference <http://help.dottoro.com/ljcjgrml.php>`_

.. [7] `How to get iframe scroll position in IE using Java Script? <http://stackoverflow.com/questions/2347491/how-to-get-iframe-scroll-position-in-ie-using-java-script>`_

|

.. [8] `JavaScript tutorial - Browser specific referencing <http://www.howtocreate.co.uk/tutorials/javascript/browserspecific>`_

.. [9] `W3C DOM Compatibility - CSS Object Model View <http://www.quirksmode.org/dom/w3c_cssom.html>`_

.. [10] `source code of jQuery .offset <https://github.com/jquery/jquery/blob/1.5.1/src/offset.js>`_

.. [11] `.position() – jQuery API <http://api.jquery.com/position/>`_

.. [12] `.offset() – jQuery API <http://api.jquery.com/offset/>`_

.. _CSS margin property: http://www.w3schools.com/css/css_margin.asp
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _display property: http://www.w3schools.com/cssref/pr_class_display.asp
.. _DOM: https://www.google.com/search?q=DOM
.. _holy grail tutorial: http://alistapart.com/article/holygrail
.. _HTML: https://www.google.com/search?q=HTML
.. _jQuery: http://jquery.com/

Hide Div When Clicked Outside It
################################

:tags: JavaScript
:category: JavaScript
:summary: Hide div element when clicked outside it, in vanilla JavaScript way. No jQuery or libraries are used.

The `search result`_ of the tutorial for the feature
"hide div element when clicked outside the div" is mostly about using jQuery_.
However I want to keep my code pure without dependency (i.e., only vanilla
JavaScript), so I tried to find a solution without using jQuey. Then I found
[1]_, but it seems not working well for my code. As a result, I re-write the
code and come up with a working code.

Please first see:

.. rubric:: `Demo <{filename}vanilla-javascript-dropdown-menu-example.html>`_
   :class: align-center

Source Code for Demo (*html*):

.. show_github_file:: siongui userpages content/articles/2015/02/13/vanilla-javascript-dropdown-menu-example.html

.. note::

  In line 10, `inline css`_ :code:`display: none;` is added to the *div* element.

  Why not put :code:`display: none;` in css file? Because if we put it in the
  css file, :code:`document.getElementById("menuDiv-dropdown").style.display`
  will return *null* instead of *none*, which will cause problem in our
  following JavaScript code.

Source Code for Demo (*JavaScript*):

.. show_github_file:: siongui userpages content/articles/2015/02/13/vanilla-javascript-dropdown-menu-example.js

Source Code for Demo (*css*):

.. show_github_file:: siongui userpages content/articles/2015/02/13/style.css

----

References:

.. [1] `hide popup div when clicking outside the div <http://www.webdeveloper.com/forum/showthread.php?t=98973>`_

.. [2] `Javascript Drop Down Menu <{filename}javascript-dropdown-menu%en.rst>`_

.. _search result: https://www.google.com/search?aq=f&gcx=w&sourceid=chrome&ie=UTF-8&q=Hide+Div+When+Clicking+Outside+the+Div

.. _jQuery: http://jquery.com/

.. _inline css: http://www.w3schools.com/css/css_howto.asp

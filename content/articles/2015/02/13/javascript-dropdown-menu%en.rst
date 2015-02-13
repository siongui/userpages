Javascript Drop Down Menu
#########################

:tags: JavaScript, dropdown menu
:category: JavaScript
:summary: Dropdown menu using vanilla JavaScript. No jQuery or libraries are used.

This post gives an example of JavaScript dropdown menu. jQuery is not used
because I want to keep it pure without dependency.

Please first see:

.. rubric:: `Demo <{filename}vanilla-javascript-dropdown-menu-example.html>`_
   :class: align-center

Source Code for Demo (*html*):

.. show_github_file:: siongui userpages content/articles/2015/02/13/vanilla-javascript-dropdown-menu-example.html

Note: In line 10, *inline css* ``display: none;`` is added to the *div* element.
Why not put ``display: none;`` in css file? Because if we put it in the css
file, ``document.getElementById("menuDiv-dropdown").style.display`` will return
*null* instead of *none*, which will cause problem in our following JavaScript
code.

Source Code for Demo (*JavaScript*):

.. show_github_file:: siongui userpages content/articles/2015/02/13/vanilla-javascript-dropdown-menu-example.js

Source Code for Demo (*css*):

.. show_github_file:: siongui userpages content/articles/2015/02/13/style.css

----

References:

.. [1] `Hide Div When Clicked Outside It <{filename}hide-div-when-clicked-outside-it%en.rst>`_

.. [2] `Retrieve the position (X,Y) of an HTML element <http://stackoverflow.com/questions/442404/retrieve-the-position-x-y-of-an-html-element>`_

.. [3] `hide popup div when clicking outside the div <http://www.webdeveloper.com/forum/showthread.php?t=98973>`_


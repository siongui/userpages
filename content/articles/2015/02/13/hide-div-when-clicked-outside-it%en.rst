Hide Div When Clicked Outside It
################################

:modified: 2017-02-02T09:57+08:00
:tags: JavaScript
:category: JavaScript
:summary: Hide div element when clicked outside it, in `vanilla JavaScript`_
          way. No jQuery_ or libraries are used.
:og_image: http://www.javatpoint.com/images/javascript/javascript_logo.png
:adsu: yes

The `search result`_ of the tutorial for the feature
"hide div element when clicked outside the div" is mostly about using jQuery_.
However I want to keep my code pure without dependency (i.e., only
`vanilla JavaScript`_), so I tried to find a solution without using jQuey.
Then I found [1]_, but it seems not working well for my code. As a result,
I re-write the code and come up with a working code.

Please first see:

.. rubric:: `Demo <{filename}/code/javascript/dropdown-menu/vanilla-javascript-dropdown-menu-example.html>`_
   :class: align-center

Source Code for Demo (*html*):

.. adsu:: 2
.. show_github_file:: siongui userpages content/code/javascript/dropdown-menu/vanilla-javascript-dropdown-menu-example.html

Source Code for Demo (*JavaScript*):

.. show_github_file:: siongui userpages content/code/javascript/dropdown-menu/vanilla-javascript-dropdown-menu-example.js
.. adsu:: 3

Source Code for Demo (*CSS*):

.. show_github_file:: siongui userpages content/code/javascript/dropdown-menu/style.css

----

References:

.. [1] `hide popup div when clicking outside the div <http://www.webdeveloper.com/forum/showthread.php?t=98973>`_

.. [2] `Javascript Drop Down Menu <{filename}javascript-dropdown-menu%en.rst>`_

.. [3] `[Golang] GopherJS DOM Example - Dropdown Menu <{filename}../../../2016/01/16/gopherjs-dom-example-dropdown-menu%en.rst>`_


.. _search result: https://www.google.com/search?q=Hide+Div+When+Clicking+Outside+the+Div
.. _vanilla JavaScript: https://www.google.com/search?q=vanilla+JavaScript
.. _jQuery: http://jquery.com/

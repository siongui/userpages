Fixed Menu Bar on Top of Website
################################

:date: 2012-06-02 22:36
:tags: CSS
:category: CSS
:summary: Put your menu bar on top of the website, no matter how users scroll.
:adsu: yes


When we log in Gmail_ or Facebook_, there is a fixed menu bar on top of the web
page. How is this being done? The answer comes to my mind after I read the
tutorial in [1]_.

The trick is to use `CSS property`_ code :code:`position:fixed` in your HTML_
element. Please refer to the following demo and sample code:

.. rubric:: `Demo <{filename}fixed.html>`_
      :class: align-center

Source Code for Demo (*html*):

.. show_github_file:: siongui userpages content/articles/2012/06/02/fixed.html
.. adsu:: 2

Source Code for Demo (*css*):

.. show_github_file:: siongui userpages content/articles/2012/06/02/style.css
.. adsu:: 3

With the above code, no matter how you scroll the web page, the menu bar is
always on top, which is the same as the menu bar on Gmail_ or Facebook_.

----

References:

.. [1] `CSS Floating Menu <http://www.quackit.com/css/codes/css_floating_menu.cfm>`_
.. [2] `CSS Float <http://www.quackit.com/css/tutorial/css_float.cfm>`_
.. [3] `CSS position <http://www.quackit.com/css/properties/css_position.cfm>`_

.. _Gmail: https://mail.google.com/
.. _Facebook: https://www.facebook.com/
.. _HTML: https://www.google.com/search?q=HTML
.. _CSS property: https://www.google.com/search?q=CSS+property

[Javascript] Width Percentage to Pixel
######################################

:date: 2012-09-23 02:42
:modified: 2015-02-26 06:24
:tags: html, JavaScript, DOM
:category: JavaScript
:summary: 1% width of browser window equal to how many pixels?
:adsu: yes


Here is an interesting and sometimes practical question:

*How do I know the 1% width of browser window equal to how many pixels?*

The answer is easy: add a *div* element right after the *body* tag, make it 100%
wide, and get the offsetWidth_ property of the *div*. From the value of
*offsetWidth* property, we will know 100% width equal to how many pixels, and
hence we can know 1% equal to how many pixels.

.. rubric:: `Demo <{filename}percentage2pixel.html>`_
      :class: align-center

Source Code for Demo:

.. adsu:: 2
.. show_github_file:: siongui userpages content/articles/2012/09/23/percentage2pixel.html
.. adsu:: 3


.. _offsetWidth: https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/offsetWidth

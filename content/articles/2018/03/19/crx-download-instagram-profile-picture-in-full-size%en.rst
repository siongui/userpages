[Chrome Extension] Download Instagram Profile Picture in Full Size
##################################################################

:date: 2018-03-19T23:26+08:00
:tags: JavaScript, String Manipulation, Instagram, Regular Expression,
       Chrome Extension
:category: JavaScript
:summary: This Chrome extension helps you download Instagram profile picture in
          full size.
:og_image: https://images.vexels.com/media/users/3/147102/isolated/preview/082213cb0f9eabb7e6715f59ef7d322a-instagram-profile-icon-by-vexels.png
:adsu: yes


This Chrome extension helps you download Instagram_ profile picture in full
size. When you navigate to profile page of Instagram user, move your mouse to
the upper-right of the profile picture, a small gray square will show and click
it. The full size picture will be downloaded.

The trick to get full size profile picture comes from
`the comment of SO question`_ [1]_. If you do not want to install extension, you
can copy the URL of the profile picture and paste the URL in the
`web page to get full size URL`_, the URL of full size picture will be returned.

.. show_github_file:: siongui userpages content/code/javascript/ig-profile-pic/manifest.json
.. show_github_file:: siongui userpages content/code/javascript/ig-profile-pic/eventPage.js
.. show_github_file:: siongui userpages content/code/javascript/ig-profile-pic/fullpic.js

----

.. adsu:: 2

References:

.. [1] `How to view instagram profile picture in full-size? - Stack Overflow <https://stackoverflow.com/questions/48468144/how-to-view-instagram-profile-picture-in-full-size>`_
.. [2] `[Chrome Extension] Show Instagram Id on User Page <{filename}/articles/2018/02/22/crx-show-ig-id-on-user-page%en.rst>`_

.. _Instagram: https://www.instagram.com/
.. _the comment of SO question: https://stackoverflow.com/questions/48468144/how-to-view-instagram-profile-picture-in-full-size#comment85451994_48468144
.. _web page to get full size URL: {filename}/articles/2018/03/18/get-instagram-profile-picture-in-full-size%en.rst

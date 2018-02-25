[Chrome Extension] Show Instagram Mutual Followers on User Web Profile
######################################################################

:date: 2018-02-26T07:39+08:00
:tags: JavaScript, Chrome Extension, Web Scrape, Instagram
:category: JavaScript
:summary: A Chrome extension to help you show Instagram mutual followers on user
          profile page.
:og_image: https://techuntold-techuntold.netdna-ssl.com/wp-content/uploads/2017/08/How-to-See-Mutual-Friends-on-Instagram.png
:adsu: yes


When you use mobile Instagram_ app and visit user profile, there will be mutual
followers like

::

  Followed by userA, userB, userC + 2 more

This feature is missing on Instagram web interface, and this `Chrome extension`_
help you show mutual followers on the Instagram user web profile page.

`manifest json`_:

.. show_github_file:: siongui userpages content/code/javascript/ig-mutual-follower/manifest.json

**background.js**:

.. show_github_file:: siongui userpages content/code/javascript/ig-mutual-follower/eventPage.js
.. adsu:: 2

**showinfo.js**:

.. show_github_file:: siongui userpages content/code/javascript/ig-mutual-follower/showinfo.js

----

.. adsu:: 3

References:

.. [1] `GitHub - siongui/igidcrx: Get Instagram Id via Chrome Extension <https://github.com/siongui/igidcrx>`_

.. _Chrome extension: https://www.google.com/search?q=Chrome+Extension
.. _Instagram: https://www.instagram.com/
.. _manifest json: https://developer.chrome.com/extensions/manifest

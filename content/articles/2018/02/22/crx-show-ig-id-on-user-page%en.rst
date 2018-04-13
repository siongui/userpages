[Chrome Extension] Show Instagram Id on User Page
#################################################

:date: 2018-02-22T06:16+08:00
:modified: 2018-04-13T23:49+08:00
:tags: JavaScript, Chrome Extension, Web Scrape, String Manipulation, Instagram,
       Regular Expression
:category: Chrome Extension
:summary: A Chrome extension to help you show Instagram id on the user page.
:og_image: https://www.otzberg.net/iguserid/instagram_user_id_chrome.png
:adsu: yes

[Update]: Now Instagram bans access to */?__a=1* URL, so this extension cannot
work now. See my new way to do this [2]_.

A `Chrome extension`_ to help you show Instagram_ id on the user page.

`manifest json`_:

.. show_github_file:: siongui userpages content/code/javascript/ig-id-xhr/manifest.json

**background.js**:

.. show_github_file:: siongui userpages content/code/javascript/ig-id-xhr/eventPage.js

**showid.js**:

.. show_github_file:: siongui userpages content/code/javascript/ig-id-xhr/showid.js

----

.. adsu:: 2

References:

.. [1] `GitHub - siongui/fbidcrx: Try to find Facebook Id via Chrome Extension <https://github.com/siongui/fbidcrx>`_
.. [2] `[Chrome Extension] Get Instagram User Information From HTML Source <{filename}/articles/2018/04/13/crx-get-instagram-user-information%en.rst>`_

.. _Chrome extension: https://www.google.com/search?q=Chrome+Extension
.. _Instagram: https://www.instagram.com/
.. _manifest json: https://developer.chrome.com/extensions/manifest

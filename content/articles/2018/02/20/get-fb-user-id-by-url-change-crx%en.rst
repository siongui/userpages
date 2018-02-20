Try to Get Facebook User Id by URL Change and Chrome Extension
##############################################################

:date: 2018-02-20T10:16+08:00
:tags: JavaScript, Chrome Extension, Web Scrape, String Manipulation,
       Regular Expression
:category: JavaScript
:summary: A Chrome extension to help you try to get Facebook user id by URL
          change.
:og_image: http://2.bp.blogspot.com/_nZM0Wsesudk/TFb0TTsMBEI/AAAAAAAAAQk/5Ooee76P1eo/s1600/SS-USER%2BID-s.jpg
:adsu: yes


Sometimes Facebook URL will change from:

::

  https://www.facebook.com/profile.php?id=12345678

To

::

  https://www.facebook.com/my.user.name


From the URL change, we know the id of *my.user.name* is *12345678*.

So I write a `Chrome extension`_ to detect URL change in *facebook.com* and help
you get the id of users by URL change.

You can `load extension`_ in developer mode. When the URL change is detected,
you can see the user *name : id* pair in the `extension console`_ (not Chrome
DevTools console).

`manifest json`_:

.. show_github_file:: siongui userpages content/code/javascript/fb-url-change/manifest.json

**background.js**:

.. show_github_file:: siongui userpages content/code/javascript/fb-url-change/eventPage.js

----

.. adsu:: 2

References:

.. [1] `javascript - How to listen for url change with Chrome Extension - Stack Overflow <https://stackoverflow.com/questions/34957319/how-to-listen-for-url-change-with-chrome-extension>`_
.. [2] `debugging - Where to read console messages from background.js in a Chrome extension? - Stack Overflow <https://stackoverflow.com/questions/10257301/where-to-read-console-messages-from-background-js-in-a-chrome-extension>`_
.. [3] | `String.prototype.match() - JavaScript | MDN <https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/match>`_
       | `javascript - Regex - Match whole string - Stack Overflow <https://stackoverflow.com/questions/6298566/regex-match-whole-string>`_
       | `Online regex tester and debugger: PHP, PCRE, Python, Golang and JavaScript <https://regex101.com/>`_

.. _Chrome extension: https://www.google.com/search?q=Chrome+Extension
.. _manifest json: https://developer.chrome.com/extensions/manifest
.. _load extension: https://developer.chrome.com/extensions/getstarted#unpacked
.. _extension console: https://stackoverflow.com/a/10258029

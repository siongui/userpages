[Chrome Extension] Try to Print Facebook Id Next to User Name
#############################################################

:date: 2018-02-21T21:16+08:00
:tags: JavaScript, Chrome Extension, Web Scrape, String Manipulation,
       Regular Expression
:category: JavaScript
:summary: A Chrome extension to help you try to get Facebook user id by URL
          change and print the id next to name if found.
:og_image: http://2.bp.blogspot.com/_nZM0Wsesudk/TFb0TTsMBEI/AAAAAAAAAQk/5Ooee76P1eo/s1600/SS-USER%2BID-s.jpg
:adsu: yes


Continue `the work I have done`_ yesterday, instead of print in the console, the
id will be printed next to the name of the user if the id is found.
The following is the code:

`manifest json`_:

.. show_github_file:: siongui userpages content/code/javascript/fb-url-change-msg/manifest.json

**background.js**:

.. show_github_file:: siongui userpages content/code/javascript/fb-url-change-msg/eventPage.js

**showid.js**:

.. show_github_file:: siongui userpages content/code/javascript/fb-url-change-msg/showid.js

----

.. adsu:: 2

References:

.. [1] `Message Passing - Google Chrome <https://developer.chrome.com/apps/messaging>`_
.. [2] `javascript - How to wait for another JS to load to proceed operation? - Stack Overflow <https://stackoverflow.com/questions/8618464/how-to-wait-for-another-js-to-load-to-proceed-operation>`_
.. [3] `Window setInterval() Method <https://www.w3schools.com/jsref/met_win_setinterval.asp>`_

.. _Chrome extension: https://www.google.com/search?q=Chrome+Extension
.. _manifest json: https://developer.chrome.com/extensions/manifest
.. _load extension: https://developer.chrome.com/extensions/getstarted#unpacked
.. _extension console: https://stackoverflow.com/a/10258029
.. _the work I have done: {filename}/articles/2018/02/20/get-fb-user-id-by-url-change-crx%en.rst

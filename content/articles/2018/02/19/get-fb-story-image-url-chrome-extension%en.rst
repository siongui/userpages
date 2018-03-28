Get Facebook Stories Image URL via Chrome Extension
###################################################

:date: 2018-02-19T07:38+08:00
:tags: JavaScript, Chrome Extension, Web Scrape
:category: Chrome Extension
:summary: A naive Chrome extension to help you get URL of images of Facebook
          Stories.
:og_image: https://blog.bufferapp.com/wp-content/uploads/2017/01/facebook-stories-24-hours.png
:adsu: yes


A naive `Chrome extension`_ to help you get URL of images of Facebook Stories,
so you don't have to right-click and open Chrome DevTools to check the URL.

You can `load extension`_ in developer mode. When the image of stories is
playing, click the icon of this extension and the URL of the image will be
logged in the console of Chrome DevTools.

`manifest json`_:

.. show_github_file:: siongui userpages content/code/javascript/fb-story-img-url/manifest.json

**background.js**:

.. show_github_file:: siongui userpages content/code/javascript/fb-story-img-url/background.js

**inject.js**:

.. show_github_file:: siongui userpages content/code/javascript/fb-story-img-url/inject.js

----

.. adsu:: 2

References:

.. [1] | `chrome extension execute script on click - Google search <https://www.google.com/search?q=chrome+extension+execute+script+on+click>`_
       | `chrome extension execute script on click - DuckDuckGo search <https://duckduckgo.com/?q=chrome+extension+execute+script+on+click>`_
       | `chrome extension execute script on click - Ecosia search <https://www.ecosia.org/search?q=chrome+extension+execute+script+on+click>`_
       | `chrome extension execute script on click - Qwant search <https://www.qwant.com/?q=chrome+extension+execute+script+on+click>`_
       | `chrome extension execute script on click - Bing search <https://www.bing.com/search?q=chrome+extension+execute+script+on+click>`_
       | `chrome extension execute script on click - Yahoo search <https://search.yahoo.com/search?p=chrome+extension+execute+script+on+click>`_
       | `chrome extension execute script on click - Baidu search <https://www.baidu.com/s?wd=chrome+extension+execute+script+on+click>`_
       | `chrome extension execute script on click - Yandex search <https://www.yandex.com/search/?text=chrome+extension+execute+script+on+click>`_

.. _Chrome extension: https://www.google.com/search?q=Chrome+Extension
.. _DOM manipulation: https://www.google.com/search?q=DOM+manipulation
.. _manifest json: https://developer.chrome.com/extensions/manifest
.. _tabs permission: https://developer.chrome.com/extensions/declare_permissions
.. _document.write: https://developer.mozilla.org/en-US/docs/Web/API/Document/write
.. _load extension: https://developer.chrome.com/extensions/getstarted#unpacked

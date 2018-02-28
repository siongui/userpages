[Chrome Extension] Not Being Logged in in Incognito Mode
########################################################

:date: 2018-02-28T21:21+08:00
:tags: Chrome Extension
:category: Chrome Extension
:summary: Use "*incognito*" manifest key "*split*" in Chrome Extension to leave
          logged in status.
:og_image: https://browsernative.com/wp-content/uploads/allow-chrome-extension-incognito.png
:adsu: yes


When I was writing `Chrome extension`_, sometimes I need to open the tab in
incognito mode. But something puzzled me that when I open the tab in incognito
mode, the Chrome extension still keeps the login status. This is obviously not
something that I expected because when you are logged into a website and open a
new tab of the same site in incognito mode, you are not logged in in the
incognito mode, but the extension still keeps logged in!

After some googling, I found the incognito_ key in `manifest json`_ of Chrome
extension. If you do not specify in the *manifest.json*, the default is
*spanning*, which means you make HTTP request in the extension, you have still
have the same cookie information of being logged in.

The description of incognito_ key on developer website does not say very clearly
about this. Only the following sentences in the final:

  As a rule of thumb, if your extension or app needs to load a tab in an
  incognito browser, use split incognito behavior. If your extension or app
  needs to be logged into a remote server use spanning incognito behavior.

So if you want your extension to make HTTP request without being logged in in
incognito mode, use "*incognito*" manifest key "*split*" in your
*manifest.json*.

----

.. adsu:: 2

References:

.. [1] `GitHub - siongui/igidcrx: Get Instagram Id via Chrome Extension <https://github.com/siongui/igidcrx>`_

.. _Chrome extension: https://www.google.com/search?q=Chrome+Extension
.. _manifest json: https://developer.chrome.com/extensions/manifest
.. _incognito: https://developer.chrome.com/extensions/manifest/incognito

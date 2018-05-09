[Golang] Get Instagram Highlight Stories of Specific User
#########################################################

:date: 2018-02-09T06:21+08:00
:modified: 2018-05-09T23:47+08:00
:tags: Go, Golang, Web Scrape, Go net/http, Instagram, HTTP cookie
:category: Go
:summary: Get links of Instagram highlight stories of a specific user via Go
          programming language. Use only Go standard library and no third-party
          packages needed.
:og_image: https://tctechcrunch2011.files.wordpress.com/2017/12/instagram-stories-highlights-archive.png
:adsu: yes

**[Update]**: the method in this post can only get partial story highlights now.
Please see my `another new post`_ for how to get the rest of story highlights.

Interesting small program to get the JSON format data of highlight stories of a
specific Instagram_ user.

When I used Chrome extension `InstaG Downloader`_ version 2.1.0, I found it will
show highlight stories of users when you visit the page of the user. I am
curious how it works so I installed another extension `Network Sniffer`_ to see
the API URL of highlight stories, and found that API URL is

::

  https://i.instagram.com/api/v1/highlights/{{USERID}}/highlights_tray/

The procedure of reading JSON data is the same as the original story API [1]_.
Please read the post if you are not familiar with the API.

In this program only Go standard library is used, no third-party packages.

To access the Instagram API via local Go program, you need to login Instagram_
and get the following information from your browser:

- ``ds_user_id``
- ``sessionid``
- ``csrftoken``

Please see this `SO answer`_ to get above values on Chrome browser.

Moreover, you need to know the user id of the specific user, please read my
previous post to get the id of the user. [3]_

After you get the values, you can get the JSON response from the following code:

.. show_github_file:: siongui userpages content/code/go/ig-highlight-story/config.go
.. show_github_file:: siongui userpages content/code/go/ig-highlight-story/userhighlightstories.go
.. adsu:: 2

**Example**:

.. show_github_file:: siongui userpages content/code/go/ig-highlight-story/userhighlightstories_test.go
.. adsu:: 3

How to parse the JSON data to get URL and timestamp of stories, please see my
GitHub repo [2]_.

----

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.9.4``.

----

References:

.. [1] `Chrome IG Story — Bribing the Instagram Story API with cookies 🍪🍪🍪 <https://medium.com/@calialec/chrome-ig-story-bribing-the-instagram-story-api-with-cookies-c813e6dff911>`_
.. [2] `GitHub - siongui/goigstorylink: Get Links (URL) of Instagram Stories in Go <https://github.com/siongui/goigstorylink>`_
.. [3] `[Golang] Get Instagram User ID <{filename}/articles/2018/02/04/go-get-instagram-user-id%en.rst>`_

.. _Instagram: https://www.instagram.com/
.. _SO answer: https://stackoverflow.com/a/44773079
.. _InstaG Downloader: https://chrome.google.com/webstore/detail/instag-downloader/jnkdcmgmnegofdddphijckfagibepdlb
.. _Network Sniffer: https://chrome.google.com/webstore/detail/network-sniffer/coblekblkacfilmgdghecpekhadldjfj
.. _another new post: {filename}/articles/2018/05/09/go-get-all-story-highlights-of-instagram-user%en.rst

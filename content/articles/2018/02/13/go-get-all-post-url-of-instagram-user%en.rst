[Golang] Get Url of All Posts of Instagram User
###############################################

:date: 2018-02-13T09:16+08:00
:tags: Go, Golang, Web Scrape, Go net/http, Instagram
:category: Go
:summary: Get URL of all posts of a specific Instagram user via Go programming
          language. Use only Go standard library and no third-party packages
          needed.
:og_image: http://media.idownloadblog.com/wp-content/uploads/2017/06/Instagram-for-iOS-how-to-archive-post-iPhone-screenshot-001.jpg
:adsu: yes

Interesting small program to get URL of all posts of a specific Instagram_ user.

In this program only Go standard library is used, no third-party packages.

To access the Instagram API via local Go program, you need to login Instagram_
and get the following information from your browser:

- ``ds_user_id``
- ``sessionid``
- ``csrftoken``

.. image:: https://i.stack.imgur.com/psJLZ.png
   :align: center
   :alt: ds_user_id sessionid csrftoken

Please see this `SO answer`_ to get above values on Chrome browser.

In each HTTP request, Instagram API returns only one page (12 posts), and also a
token (*end_cursor* in this case) to get next page. The token is embedded in the
query string of next HTTP request, and the name is *max_id*. Please read the
code for more details.

.. show_github_file:: siongui userpages content/code/go/ig-all-posts/userinfo.go
.. adsu:: 2

**Example**:

.. show_github_file:: siongui userpages content/code/go/ig-all-posts/userinfo_test.go
.. adsu:: 3

The full code is also available on my GitHub repo [1]_.

----

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.9.4``.

----

References:

.. [1] `GitHub - siongui/goigmedia: Get links of Instagram user media (photos and videos) in Go. <https://github.com/siongui/goigmedia>`_

.. _Instagram: https://www.instagram.com/
.. _SO answer: https://stackoverflow.com/a/44773079

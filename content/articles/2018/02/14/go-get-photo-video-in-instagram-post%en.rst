[Golang] Get Photos and Videos in Instagram Post
################################################

:date: 2018-02-14T07:43+08:00
:tags: Go, Golang, Web Scrape, Go net/http, Instagram, HTTP cookie
:category: Go
:summary: Get URL of photos and videos in Instagram post via Go programming
          language. Use only Go standard library and no third-party packages
          needed.
:og_image: http://media.idownloadblog.com/wp-content/uploads/2017/06/Instagram-for-iOS-how-to-archive-post-iPhone-screenshot-001.jpg
:adsu: yes

Interesting small program to get URL of all photos and videos in Instagram_
post.

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

Given the URL of the post as follows:

::

  https://www.instagram.com/p/BfJzG64BZVY/

The *code* of the post is **BfJzG64BZVY**. We will use the code as one of the
arguments in our func call.

.. show_github_file:: siongui userpages content/code/go/ig-post-media/post.go
.. adsu:: 2

**Example**:

.. show_github_file:: siongui userpages content/code/go/ig-post-media/post_test.go
.. adsu:: 3

The full code is also available on my GitHub repo [1]_.

----

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.9.4``.

----

References:

.. [1] `GitHub - siongui/goigmedia: Get links of Instagram user media (photos and videos) in Go. <https://github.com/siongui/goigmedia>`_

.. _Instagram: https://www.instagram.com/
.. _SO answer: https://stackoverflow.com/a/44773079

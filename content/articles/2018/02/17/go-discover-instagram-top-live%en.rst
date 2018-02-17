[Golang] Discover Instagram Top Live
####################################

:date: 2018-02-17T21:46+08:00
:tags: Go, Golang, Web Scrape, Go net/http, Instagram
:category: Go
:summary: Access Instagram API to discover Top lives in Go programming language.
          Use only Go standard library and no third-party packages needed.
:og_image: https://simplymeasured.com/wp-content/uploads/2016/12/IMG_0556-576x1024.png
:adsu: yes

Interesting small program to discover Instagram_ top lives on local machines in
Go.

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

.. show_github_file:: siongui userpages content/code/go/ig-toplive/toplive.go
.. adsu:: 2

**Example**:

.. show_github_file:: siongui userpages content/code/go/ig-toplive/toplive_test.go
.. adsu:: 3

The full code is also available on my GitHub repo [1]_.

----

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.9.4``.

----

References:

.. [1] `GitHub - siongui/instago: Get photos, videos, stories, following and followers of Instagram <https://github.com/siongui/instago>`_

.. _Instagram: https://www.instagram.com/
.. _SO answer: https://stackoverflow.com/a/44773079

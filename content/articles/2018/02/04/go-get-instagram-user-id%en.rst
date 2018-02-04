[Golang] Get Instagram User ID
##############################

:date: 2018-02-04T08:52+08:00
:tags: Go, Golang, Web Scrape, Go net/http, Go strings Package, JSON
:category: Go
:summary: Given user name, get Instagram user id via Go programming language.
          Use only Go standard library and no third-party packages needed.
:og_image: http://zentools2.joomlabamboo.com/images/documentation/instagram/get-user-id.png
:adsu: yes

Interesting small program to get Instagram_ user id, assume that user name is
given. The method of getting id is found in the answer of Stack Overflow
question [1]_. No login or cookie required.
We use only Go standard library to get id. No third-party packages are used.

The following is complete source code:

.. show_github_file:: siongui userpages content/code/go/ig-user-id/userinfo.go

If you want, you can take a look at the JSON response and get more information
of the user.

.. adsu:: 2

**Usage**:

.. show_github_file:: siongui userpages content/code/go/ig-user-id/userinfo_test.go

**Output**:

.. code-block:: txt

  === RUN   TestGetUserInfo
  --- PASS: TestGetUserInfo (0.74s)
  	userinfo_test.go:13: 25025320
  	userinfo_test.go:14: Discovering — and telling — stories from around the world. Curated by Instagram’s community team.
  PASS

.. adsu:: 3

----

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.9.3``.

----

References:

.. [1] `Instagram API -Get the userId - Stack Overflow <https://stackoverflow.com/a/44773079>`_

.. _Instagram: https://www.instagram.com/

.. |godoc| image:: https://godoc.org/github.com/PuerkitoBio/goquery?status.png
   :target: https://godoc.org/github.com/PuerkitoBio/goquery

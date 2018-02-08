[Golang] Get List of Instagram Stories
######################################

:date: 2018-02-05T02:13+08:00
:tags: Go, Golang, Web Scrape, Go net/http, Instagram
:category: Go
:summary: Get the list of following Instagram stories via Go programming
          language. Use only Go standard library and no third-party packages
          needed.
:og_image: http://s1.r29static.com//bin/entry/65b/720x420/1800328/image.png
:adsu: yes

Interesting small program to get the JSON format data of the list of your
following Instagram_ stories.
The whole picture of getting the JSON comes from the post of Chrome IG Story
[1]_. Please read the post first.
In this program only Go standard library is used, no third-party packages.

To access the Instagram API via local Go program, you need to login Instagram_
and get the following information from your browser:

- ``ds_user_id``
- ``sessionid``
- ``csrftoken``

Please see this `SO answer`_ to get above values on Chrome browser.

After you get the values, modify the following code to fill the values:

.. show_github_file:: siongui userpages content/code/go/ig-stories-list/config.go
.. adsu:: 2

After the values are set in the code, the following code can get JSON data of
the list of your following Instagram stories:

.. show_github_file:: siongui userpages content/code/go/ig-stories-list/allstories.go
.. adsu:: 3

**Usage**:

.. show_github_file:: siongui userpages content/code/go/ig-stories-list/allstories_test.go
.. adsu:: 4

----

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.9.3``.

----

References:

.. [1] `Chrome IG Story — Bribing the Instagram Story API with cookies 🍪🍪🍪 <https://medium.com/@calialec/chrome-ig-story-bribing-the-instagram-story-api-with-cookies-c813e6dff911>`_

.. _Instagram: https://www.instagram.com/
.. _SO answer: https://stackoverflow.com/a/44773079

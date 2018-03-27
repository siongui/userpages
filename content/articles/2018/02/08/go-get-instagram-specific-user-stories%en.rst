[Golang] Get Instagram Stories of Specific User
###############################################

:date: 2018-02-08T06:00+08:00
:tags: Go, Golang, Web Scrape, Go net/http, Instagram, HTTP cookie
:category: Go
:summary: Get links of Instagram stories of a specific user via Go programming
          language. Use only Go standard library and no third-party packages
          needed.
:og_image: http://s1.r29static.com//bin/entry/65b/720x420/1800328/image.png
:adsu: yes

Interesting small program to get the JSON format data of stories of a specific
Instagram_ user.
The whole picture of getting the JSON comes from the post of Chrome IG Story
[1]_. Please read the post first.
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

.. show_github_file:: siongui userpages content/code/go/ig-user-stories/userstories.go
.. adsu:: 2

How to parse the JSON data to get URL and timestamp of stories, please see my
GitHub repo [2]_.

----

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.9.3``.

----

References:

.. [1] `Chrome IG Story — Bribing the Instagram Story API with cookies 🍪🍪🍪 <https://medium.com/@calialec/chrome-ig-story-bribing-the-instagram-story-api-with-cookies-c813e6dff911>`_
.. [2] `GitHub - siongui/goigstorylink: Get Links (URL) of Instagram Stories in Go <https://github.com/siongui/goigstorylink>`_
.. [3] `[Golang] Get Instagram User ID <{filename}/articles/2018/02/04/go-get-instagram-user-id%en.rst>`_

.. _Instagram: https://www.instagram.com/
.. _SO answer: https://stackoverflow.com/a/44773079

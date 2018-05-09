[Golang] Get All Story Highlights of Specific Instagram User
############################################################

:date: 2018-05-09T23:42+08:00
:tags: Go, Golang, Web Scrape, Go net/http, Instagram, HTTP cookie
:category: Go
:summary: Get all links of story highlights of a specific Instagram user in Go.
          Use only Go standard library and no third-party packages needed.
:og_image: https://tctechcrunch2011.files.wordpress.com/2017/12/instagram-stories-highlights-archive.png
:adsu: yes


In my `previous post`_ three months ago, we can get all story highlights from
the following single API endpoint:

::

  https://i.instagram.com/api/v1/highlights/{{USERID}}/highlights_tray/

But Instagram_ changes API again, and the above endpoint returns only partial
links now. So I mades some searches [1]_ today and found [2]_ [3]_ that we can
get the rest of links from another API endpoint:

::

  https://i.instagram.com/api/v1/feed/reels_media/

The first endpoint returns an array of highlight trays, and each of the
highlight tray consists of stories of a specific title. For users who has many
highlight trays, only first few trays have links of stories. The rest of the
trays are empty.

The get the content of empty highlight trays, we need to supply the id of the
highlight trays to the second API endpoint one by one, and the endpoint will
return the content of the highlight tray.

The above description is the whole picture of how to get all links of story
highlights. For implementation details, please see my commit [4]_ and instago_
repo.

----

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.10.2``.

----

References:

.. [1] | `instagram story highlight github - Google search <https://www.google.com/search?q=instagram+story+highlight+github>`_
       | `instagram story highlight github - DuckDuckGo search <https://duckduckgo.com/?q=instagram+story+highlight+github>`_
       | `instagram story highlight github - Ecosia search <https://www.ecosia.org/search?q=instagram+story+highlight+github>`_
       | `instagram story highlight github - Qwant search <https://www.qwant.com/?q=instagram+story+highlight+github>`_
       | `instagram story highlight github - Bing search <https://www.bing.com/search?q=instagram+story+highlight+github>`_
       | `instagram story highlight github - Yahoo search <https://search.yahoo.com/search?p=instagram+story+highlight+github>`_
       | `instagram story highlight github - Baidu search <https://www.baidu.com/s?wd=instagram+story+highlight+github>`_
       | `instagram story highlight github - Yandex search <https://www.yandex.com/search/?text=instagram+story+highlight+github>`_
.. [2] `Instagram-API/Highlight.php at master · mgp25/Instagram-API · GitHub <https://github.com/mgp25/Instagram-API/blob/master/src/Request/Highlight.php>`_
.. [3] `Instagram-API/Story.php at master · mgp25/Instagram-API · GitHub <https://github.com/mgp25/Instagram-API/blob/master/src/Request/Story.php>`_
.. [4] `get all story highlights of user · siongui/instago@f2a9f33 · GitHub <https://github.com/siongui/instago/commit/f2a9f3385d998155a71c5b779ab73e5128be8599>`_

.. _Instagram: https://www.instagram.com/
.. _previous post: {filename}/articles/2018/02/09/go-get-instagram-specific-user-highlight-stories%en.rst
.. _instago: https://github.com/siongui/instago

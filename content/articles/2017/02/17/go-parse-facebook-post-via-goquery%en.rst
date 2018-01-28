[Golang] Web Scrape Facebook Post via goquery
#############################################

:date: 2017-02-17T21:53+08:00
:tags: Go, Golang, goquery, Web Scrape, html, DOM, Go strconv Package,
       Go strings Package, CSS selectors
:category: Go
:summary:  Fetch and parse a public post on Facebook_ and extract data via
           goquery_ in Go_ programming language.
:og_image: https://jonathanmh.com/wp-content/uploads/2016/10/web-scraping-web-crawling-golang-go-goquery-825x371.png
:adsu: yes


Introduction
++++++++++++

When it comes to web scraping, people usually thinks of Python_. There are a lot
of powerful libraries and tools for web scraping and html parsing in Python_
ecosystem, but Go_ is catching up too. Here we will show how to extract data
from a public Facebook_ post via Golang_ and goquery_, which is similar to
JavaScript_ jQuery_ for html parsing.

We will us this `Facebook post`_ as example and shows how to extract data from
it. Remember that Facebook serves different HTML contents depending on whether
users are logged in or not. Here we fetch and parse only public posts, so login
is not required. If you view source of the `Facebook post`_, remember to logout
all Facebook account, so that you will not be overwhelmed by too much HTML code.

The following are steps for web scraping:

  1. Get the part of HTML which contains the `Facebook post`_.
  2. Get timestamp of the post
  3. Get profile link of the post
  4. Get image url of the post
  5. Get content of the post

.. adsu:: 2

Extract HTML String of Post
+++++++++++++++++++++++++++

First we need to get the part of HTML which contains the `Facebook post`_. After
checking the source of the post, I found that the part of HTML containing the
post looks like the following without login:

.. code-block:: html

  <div class="hidden_elem"><code id="u_0_f"><!-- {{POST_HTML}} --></code></div>

It's embedded as a comment and used by JavaScript_. The following code can
extract *{{POST_HTML}}*:

.. show_github_file:: siongui userpages content/code/go/scrape-facebook-post/fb.go

We can use goquery_ Find_ function with `CSS selector`_
**div.hidden_elem > code**. There will be two elements match the above selector,
the post is in first matched element, so we use First_ function to get the first
matched element. Retrieve the innerHTML_ of the first matched element by Html_
function, and remove the leading and trailing arrows of comments, and return the
HTML string of the post.

.. adsu:: 3

Post Timestamp
++++++++++++++

Now given the HTML string of post, next step is to get the timestamp of the
post. Again we look at the HTML string of the post, we find that the time is
embedded in the following HTML element:

.. code-block:: html

  <abbr title="Wednesday, February 15, 2017 at 7:00am" data-utime="1487113202" data-shorten="1" class="_5ptz"><span class="timestampContent">Yesterday at 7:00am</span></abbr>

We can use the following code to extract the utime and convert it to
human-readable form:

.. show_github_file:: siongui userpages content/code/go/scrape-facebook-post/time.go

See `my post for parsing Unix time`_ [4]_ for more details.

.. adsu:: 4

Post Profile Link
+++++++++++++++++

Get the name and url of the user of the post.

.. show_github_file:: siongui userpages content/code/go/scrape-facebook-post/profilelink.go

The logic in above code is the same. Just find the element which contains the
data you are looking for, and use correct `CSS selector`_ to get the element we
need.

.. adsu:: 5

Post Image
++++++++++

Retrieve the URL of the image of the post:

.. show_github_file:: siongui userpages content/code/go/scrape-facebook-post/image.go

Post Content
++++++++++++

Get the content of the post:

.. show_github_file:: siongui userpages content/code/go/scrape-facebook-post/content.go

Summary
+++++++

Use all the above code to extract data from post:

.. show_github_file:: siongui userpages content/code/go/scrape-facebook-post/fb_test.go

----

Tested on: ``Ubuntu Linux 16.10``, ``Go 1.8``.

----

.. adsu:: 6

References:

.. [1] `GitHub - PuerkitoBio/goquery: A little like that j-thing, only in Go. <https://github.com/PuerkitoBio/goquery>`_ |godoc|
.. [2] `Tips and tricks · PuerkitoBio/goquery Wiki · GitHub <https://github.com/PuerkitoBio/goquery/wiki/Tips-and-tricks>`_
.. [3] `goquery querySelector <{filename}../15/goquery-querySelector-golang%en.rst>`_
.. [4] `[Golang] Parse Unix Time (utime) Example <{filename}../16/go-parse-utime-timestamp%en.rst>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Python: https://www.python.org/
.. _goquery: https://github.com/PuerkitoBio/goquery
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _jQuery: https://jquery.com/
.. _Facebook: https://www.facebook.com/
.. _view HTML source code: view-source:https://www.facebook.com/jayasaro.panyaprateep.org/posts/1095007907274561:0
.. _Facebook post: https://www.facebook.com/jayasaro.panyaprateep.org/posts/1095007907274561:0
.. _CSS selector: https://www.google.com/search?q=CSS+selector
.. _Find: https://godoc.org/github.com/PuerkitoBio/goquery#Selection.Find
.. _First: https://godoc.org/github.com/PuerkitoBio/goquery#Selection.First
.. _innerHTML: https://developer.mozilla.org/en-US/docs/Web/API/Element/innerHTML
.. _Html: https://godoc.org/github.com/PuerkitoBio/goquery#Selection.Html
.. _*Selection: https://godoc.org/github.com/PuerkitoBio/goquery#Selection
.. _my post for parsing Unix time: {filename}../16/go-parse-utime-timestamp%en.rst

.. |godoc| image:: https://godoc.org/github.com/PuerkitoBio/goquery?status.png
   :target: https://godoc.org/github.com/PuerkitoBio/goquery

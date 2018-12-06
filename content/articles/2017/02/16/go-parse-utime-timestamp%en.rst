[Golang] Parse Unix Time (utime) Example
########################################

:date: 2017-02-16T09:56+08:00
:tags: Go, Golang, Web Scrape, Go time Package, Go strconv Package
:category: Go
:summary: Parse the string of `Unix time`_ (also known as *POSIX time* or
          *Epoch time*) in Go_ programmming language.
:og_image: http://www.unixstickers.com/image/cache/data/stickers/golang/Go-brown-side.sh-600x600.png
:adsu: yes


Introduction
++++++++++++

When I tried to get the timestamp of the `Facebook post`_, I found that the
timestamp is embedded in the *data-utime* attribute of abbr_ element:

.. code-block:: html

  <abbr title="Wednesday, February 15, 2017 at 7:00am" data-utime="1487113202" data-shorten="1" class="_5ptz"><span class="timestampContent">Yesterday at 7:00am</span></abbr>

The string *1487113202* looks familiar, so I did some googling [1]_ [7]_ and
found that it represents `Unix time`_, seconds and nanoseconds that have elapsed
since January 1, 1970 UTC.

In this post, we will show how to parse the string of `Unix time`_ (also known
as *POSIX time* or *Epoch time*) in Go_ programmming language.

.. adsu:: 2

Solution
++++++++

Steps:

1. Use `strconv.ParseInt`_ to convert the string to integer.

2. Use `time.Unix`_ to convert the integer to `time.Time`_ type,
   which represents an instant in time with nanosecond precision.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/E5eVo99fWO>`_
      :class: align-center

.. code-block:: go

  import (
  	"strconv"
  	"time"
  )

  func ParseTimeStamp(utime string) (string, error) {
  	i, err := strconv.ParseInt(utime, 10, 64)
  	if err != nil {
  		return "", err
  	}
  	t := time.Unix(i, 0)
  	return t.Format(time.UnixDate), nil
  }

.. adsu:: 3

Result returned from running code on `Go Playground`_:

.. code-block:: txt

  Tue Feb 14 23:00:02 UTC 2017

----

Tested on:

- ``Ubuntu Linux 16.10``, ``Go 1.7.5``
- The `Go Playground`_

----

References
++++++++++

.. [1] | `golang parse timestamp - Google search <https://www.google.com/search?q=golang+parse+timestamp>`_
       | `golang parse timestamp - DuckDuckGo search <https://duckduckgo.com/?q=golang+parse+timestamp>`_
       | `golang parse timestamp - Ecosia search <https://www.ecosia.org/search?q=golang+parse+timestamp>`_
       | `golang parse timestamp - Bing search <https://www.bing.com/search?q=golang+parse+timestamp>`_
       | `golang parse timestamp - Yahoo search <https://search.yahoo.com/search?p=golang+parse+timestamp>`_
       | `golang parse timestamp - Baidu search <https://www.baidu.com/s?wd=golang+parse+timestamp>`_
       | `golang parse timestamp - Yandex search <https://www.yandex.com/search/?text=golang+parse+timestamp>`_
.. [2] `date - How to parse unix timestamp in golang - Stack Overflow <http://stackoverflow.com/questions/24987131/how-to-parse-unix-timestamp-in-golang>`_
.. [3] | `time - The Go Programming Language <https://golang.org/pkg/time/>`_
       | `strconv - The Go Programming Language <https://golang.org/pkg/strconv/>`_
.. [4] | `html utime - Google search <https://www.google.com/search?q=html+utime>`_
       | `html utime - DuckDuckGo search <https://duckduckgo.com/?q=html+utime>`_
       | `html utime - Ecosia search <https://www.ecosia.org/search?q=html+utime>`_
       | `html utime - Bing search <https://www.bing.com/search?q=html+utime>`_
       | `html utime - Yahoo search <https://search.yahoo.com/search?p=html+utime>`_
       | `html utime - Baidu search <https://www.baidu.com/s?wd=html+utime>`_
       | `html utime - Yandex search <https://www.yandex.com/search/?text=html+utime>`_
.. [5] `html - How to use <abbr> to show dates (example from Facebook) - Stack Overflow <http://stackoverflow.com/questions/31678736/how-to-use-abbr-to-show-dates-example-from-facebook>`_
.. adsu:: 4
.. [6] | `abbr utime - Google search <https://www.google.com/search?q=abbr+utime>`_
       | `abbr utime - DuckDuckGo search <https://duckduckgo.com/?q=abbr+utime>`_
       | `abbr utime - Ecosia search <https://www.ecosia.org/search?q=abbr+utime>`_
       | `abbr utime - Bing search <https://www.bing.com/search?q=abbr+utime>`_
       | `abbr utime - Yahoo search <https://search.yahoo.com/search?p=abbr+utime>`_
       | `abbr utime - Baidu search <https://www.baidu.com/s?wd=abbr+utime>`_
       | `abbr utime - Yandex search <https://www.yandex.com/search/?text=abbr+utime>`_
.. [7] | `time since epoch - Google search <https://www.google.com/search?q=time+since+epoch>`_
       | `time since epoch - DuckDuckGo search <https://duckduckgo.com/?q=time+since+epoch>`_
       | `time since epoch - Ecosia search <https://www.ecosia.org/search?q=time+since+epoch>`_
       | `time since epoch - Bing search <https://www.bing.com/search?q=time+since+epoch>`_
       | `time since epoch - Yahoo search <https://search.yahoo.com/search?p=time+since+epoch>`_
       | `time since epoch - Baidu search <https://www.baidu.com/s?wd=time+since+epoch>`_
       | `time since epoch - Yandex search <https://www.yandex.com/search/?text=time+since+epoch>`_
.. [8] `Unix time - Wikipedia <https://en.wikipedia.org/wiki/Unix_time>`_
.. [9] | `Golang Date Parsing : golang <https://www.reddit.com/r/golang/comments/7w8pwn/golang_date_parsing/>`_
       | `Fucking Go Date Format <http://fuckinggodateformat.com/>`_
.. [10] `In case you want to program a microwave oven in go - the time package formats got you covered : golang <https://old.reddit.com/r/golang/comments/a3m8o2/in_case_you_want_to_program_a_microwave_oven_in/>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
.. _Facebook post: https://www.facebook.com/jayasaro.panyaprateep.org/photos/a.318290164946343.68815.318196051622421/1119567364818615/?type=3
.. _abbr: https://www.google.com/search?q=html+abbr
.. _Unix time: https://en.wikipedia.org/wiki/Unix_time
.. _strconv.ParseInt: https://golang.org/pkg/strconv/#ParseInt
.. _time.Unix: https://golang.org/pkg/time/#Unix
.. _time.Time: https://golang.org/pkg/time/#Time

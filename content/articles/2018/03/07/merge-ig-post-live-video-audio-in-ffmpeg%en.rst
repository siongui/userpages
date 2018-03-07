Merge Instagram Post Live Video and Audio in ffmpeg
###################################################

:date: 2018-03-07T23:12+08:00
:tags: Bash, Commandline, Instagram, Ubuntu Linux, ffmpeg
:category: Bash
:summary: Use ffmpeg to merge Instagram video and audio of live replay shared to
          stories.
:adsu: yes


In my previous post, we get the `feed of Instagram post live`_ video and
`parse dash manifest`_ of post live video shared to stories. We can get two urls
in the dash manifest. One is *video* of the post live replay and the other is
*audio*. We need to find some way to merge the two *mp4* file. After some Google
searches [1]_ [2]_, I got the answer. We can use ffmpeg_ to help us:

.. code-block:: bash

  $ ffmpeg -i video.mp4 -i audio.mp4 -c:v copy -c:a copy merged.mp4

The speed of merge is so fast that I got the merged mp4 instantly after press
enter key. Really amazing!

Tested on: `Ubuntu 17.10`

.. adsu:: 2

----

References:

.. [1] | `instagram_private_api_extensions/replay.py at master · ping/instagram_private_api_extensions · GitHub <https://github.com/ping/instagram_private_api_extensions/blob/master/instagram_private_api_extensions/replay.py>`_
       |
       | `ffmpeg instagram live - Google search <https://www.google.com/search?q=ffmpeg+instagram+live>`_
       | `ffmpeg instagram live - DuckDuckGo search <https://duckduckgo.com/?q=ffmpeg+instagram+live>`_
       | `ffmpeg instagram live - Ecosia search <https://www.ecosia.org/search?q=ffmpeg+instagram+live>`_
       | `ffmpeg instagram live - Qwant search <https://www.qwant.com/?q=ffmpeg+instagram+live>`_
       | `ffmpeg instagram live - Bing search <https://www.bing.com/search?q=ffmpeg+instagram+live>`_
       | `ffmpeg instagram live - Yahoo search <https://search.yahoo.com/search?p=ffmpeg+instagram+live>`_
       | `ffmpeg instagram live - Baidu search <https://www.baidu.com/s?wd=ffmpeg+instagram+live>`_
       | `ffmpeg instagram live - Yandex search <https://www.yandex.com/search/?text=ffmpeg+instagram+live>`_

.. [2] | `How to merge audio and video file in ffmpeg - Super User <https://superuser.com/a/277667>`_
       |
       | `ffmpeg mix audio and video - Google search <https://www.google.com/search?q=ffmpeg+mix+audio+and+video>`_
       | `ffmpeg mix audio and video - DuckDuckGo search <https://duckduckgo.com/?q=ffmpeg+mix+audio+and+video>`_
       | `ffmpeg mix audio and video - Ecosia search <https://www.ecosia.org/search?q=ffmpeg+mix+audio+and+video>`_
       | `ffmpeg mix audio and video - Qwant search <https://www.qwant.com/?q=ffmpeg+mix+audio+and+video>`_
       | `ffmpeg mix audio and video - Bing search <https://www.bing.com/search?q=ffmpeg+mix+audio+and+video>`_
       | `ffmpeg mix audio and video - Yahoo search <https://search.yahoo.com/search?p=ffmpeg+mix+audio+and+video>`_
       | `ffmpeg mix audio and video - Baidu search <https://www.baidu.com/s?wd=ffmpeg+mix+audio+and+video>`_
       | `ffmpeg mix audio and video - Yandex search <https://www.yandex.com/search/?text=ffmpeg+mix+audio+and+video>`_

.. _feed of Instagram post live: {filename}/articles/2018/02/18/trick-to-get-instagram-live-video-replay%en.rst
.. _parse dash manifest: {filename}/articles/2018/03/05/go-parse-dash-manifest-in-ig-post-live-broadcast%en.rst
.. _ffmpeg: https://www.google.com/search?q=ffmpeg

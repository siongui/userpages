[Pelican] Build Offline Copy of Website
#######################################

:date: 2019-01-01T23:19+08:00
:tags: Bash, Pelican, wget, Commandline
:category: Bash
:summary: Build an offline copy of a website made by *Pelican* static site
          generator.
:og_image: https://upload.wikimedia.org/wikipedia/commons/6/60/Wget_1.13.4.png
:adsu: yes


My friend asked me whether it is possible to use Pelican_ static site generator
to build a offline website which can be viewd on local machines without internet
access. After some searches [1]_, it seems impossible to do it. Suddenly I found
that we asked the wrong question because the our purpose is totally nothing to
do with Pelican! The question should be: is it possible to get an offline copy
of an online website?

The answer is *YES*! wget_ can do this for static sites without any problem.
After some searches [2]_, I found the solution [3]_ is exactly what we need. So
sometimes asking questions from different ways can lead to unexpected easy
solution!

Tested on: ``Ubuntu Linux 18.04``.

----

References:

.. [1] | `pelican standalone website - Google search <https://www.google.com/search?q=pelican+standalone+website>`_
       | `pelican standalone website - DuckDuckGo search <https://duckduckgo.com/?q=pelican+standalone+website>`_
       | `pelican standalone website - Ecosia search <https://www.ecosia.org/search?q=pelican+standalone+website>`_
       | `pelican standalone website - Qwant search <https://www.qwant.com/?q=pelican+standalone+website>`_
       | `pelican standalone website - Bing search <https://www.bing.com/search?q=pelican+standalone+website>`_
       | `pelican standalone website - Yahoo search <https://search.yahoo.com/search?p=pelican+standalone+website>`_
       | `pelican standalone website - Baidu search <https://www.baidu.com/s?wd=pelican+standalone+website>`_
       | `pelican standalone website - Yandex search <https://www.yandex.com/search/?text=pelican+standalone+website>`_

.. [2] | `wget standalone website - Google search <https://www.google.com/search?q=wget+standalone+website>`_
       | `wget standalone website - DuckDuckGo search <https://duckduckgo.com/?q=wget+standalone+website>`_
       | `wget standalone website - Ecosia search <https://www.ecosia.org/search?q=wget+standalone+website>`_
       | `wget standalone website - Qwant search <https://www.qwant.com/?q=wget+standalone+website>`_
       | `wget standalone website - Bing search <https://www.bing.com/search?q=wget+standalone+website>`_
       | `wget standalone website - Yahoo search <https://search.yahoo.com/search?p=wget+standalone+website>`_
       | `wget standalone website - Baidu search <https://www.baidu.com/s?wd=wget+standalone+website>`_
       | `wget standalone website - Yandex search <https://www.yandex.com/search/?text=wget+standalone+website>`_
.. adsu:: 2
.. [3] `Make Offline Mirror of a Site using \`wget\` | Guy Rutenberg <https://www.guyrutenberg.com/2014/05/02/make-offline-mirror-of-a-site-using-wget/>`_

.. _Pelican: https://github.com/getpelican/pelican
.. _wget: https://en.wikipedia.org/wiki/Wget

[sed] Replace String in Directory
#################################

:date: 2017-01-11T23:47+08:00
:tags: sed, find command, Bash, Commandline, String Manipulation,
       List Files in Directory, Regular Expression
:category: Bash
:summary: Replace specific pattern of strings in HTML files under directory via
          sed_.
:adsu: yes


Question:

  Replace strings of the pattern ``gettext('Archives', DEFAULT_LANG)`` with
  ``'Archives'|gettext(DEFAULT_LANG)`` in HTML files under ``template/``
  directory.

Answer(inspired by [2]_, [4]_, and [5]_):

  .. code-block:: bash

    $ cd template/
    $ find -type f -name '*.html' | xargs sed -i "s/gettext('\(.*\)', /'\1'|gettext(/g"

----

References:

.. [1] `sed process multiple files - Google search <https://www.google.com/search?q=sed+process+multiple+files>`_

       `sed process multiple files - DuckDuckGo search <https://duckduckgo.com/?q=sed+process+multiple+files>`_

       `sed process multiple files - Bing search <https://www.bing.com/search?q=sed+process+multiple+files>`_

       `sed process multiple files - Yahoo search <https://search.yahoo.com/search?p=sed+process+multiple+files>`_

       `sed process multiple files - Baidu search <https://www.baidu.com/s?wd=sed+process+multiple+files>`_

       `sed process multiple files - Yandex search <https://www.yandex.com/search/?text=sed+process+multiple+files>`_

.. [2] `sed - Change multiple files - Stack Overflow <http://stackoverflow.com/a/30717770>`_

.. [3] `sed single quote - Google search <https://www.google.com/search?q=sed+single+quote>`_

       `sed single quote - DuckDuckGo search <https://duckduckgo.com/?q=sed+single+quote>`_

       `sed single quote - Bing search <https://www.bing.com/search?q=sed+single+quote>`_

       `sed single quote - Yahoo search <https://search.yahoo.com/search?p=sed+single+quote>`_

       `sed single quote - Baidu search <https://www.baidu.com/s?wd=sed+single+quote>`_

       `sed single quote - Yandex search <https://www.yandex.com/search/?text=sed+single+quote>`_

.. [4] `escaping - How to escape single quote in sed? - Stack Overflow <http://stackoverflow.com/a/24509931>`_

.. [5] `xgettext Extract Translatable Strings From Golang html/template <{filename}../../../2016/01/19/xgettext-extract-translatable-string-from-go-html-template%en.rst>`_

.. _sed: https://www.google.com/search?q=sed

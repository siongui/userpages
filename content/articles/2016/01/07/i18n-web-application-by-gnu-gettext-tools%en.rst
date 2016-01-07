Internationalization (i18n) of Web Application by GNU gettext Tools
###################################################################

:date: 2016-01-07T02:28+08:00
:tags: Python, JavaScript, Go, Golang, Web application, web, i18n, Google App Engine, web.py, Locale, sed, gettext
:category: Web Development
:summary: Support multiple languages in your (web) applications, by GNU gettext_
          tools, which include *xgettext*, *msginit*, *msgmerge*, and *msgfmt*.


GNU gettext_ tools are great for creating i18n (web) applications. In this post,
we will show how to use gettext_ tools, which include *xgettext*, *msginit*,
*msgmerge*, and *msgfmt* to create i18n applications.

My work, `Pāḷi Dictionary`_, is a example of i18n web application. I will show
how to use gettext_ tools by how I use gettext_ tools in `Pāḷi Dictionary`_.


Mark strings for i18n_
++++++++++++++++++++++

First, mark strings you want to tranlste by ``_("string")``. For example,

.. code-block:: html

  <title>Hello World</title>

after mark

.. code-block:: html

  <title>_("Hello World")</title>


Extract translatable strings
++++++++++++++++++++++++++++

Next, we extract the marked strings in previous step by xgettext_, which
extracts translatable strings from given input files. Assume the marked strings
are located in *html* files under ``app/html/`` directory. Run xgettext_ by:

.. code-block:: bash

  $ xgettext --no-wrap --from-code=UTF-8 --keyword=_ --output=locale/messages.pot `find app/html/ -name "*.html"`

A file named *messages.pot* will be created under ``locale/`` directory. We will
put all i18n-related stuff in ``locale/`` directory.

Next, we will set *charset* in *messages.pot*. You can open your editor to edit
by hand. I use sed_ to automate the task:

.. code-block:: bash

  $ sed -i "s/charset=CHARSET/charset=utf-8/g" locale/messages.pot


Generate translation file for each supported language
+++++++++++++++++++++++++++++++++++++++++++++++++++++

Support we want to support three languages (*English*, *Traditional Chinese*,
and *Vietnamese*). We use **en_US** locale_ for *English*, **zh_TW** for
*Traditional Chinese*, and **vi_VN** for *Vietnamese*. msginit_ is used to
generate translation file for each supported language:

.. code-block:: bash

  $ msginit --no-wrap --no-translator --input=locale/messages.pot --locale=en_US -o locale/en_US/LC_MESSAGES/messages.po
  $ msginit --no-wrap --no-translator --input=locale/messages.pot --locale=zh_TW -o locale/zh_TW/LC_MESSAGES/messages.po
  $ msginit --no-wrap --no-translator --input=locale/messages.pot --locale=vi_VN -o locale/vi_VN/LC_MESSAGES/messages.po

the input is the *messages.pot* in previous step, and three **PO** files are
generated and put in respective directories. Then edit these **PO** files to
translate the strings in the files.


Update translation file for each supported language
+++++++++++++++++++++++++++++++++++++++++++++++++++

Everytime you add or delete translatable strings in original *html* files,
msgmerge_ is used to help you update the **PO** files for each locale.
Re-generate **messages.pot** again by xgettext_ and then run msgmerge_:

.. code-block:: bash

  $ msgmerge --no-wrap --backup=none --update locale/en_US/LC_MESSAGES/messages.po locale/messages.pot
  $ msgmerge --no-wrap --backup=none --update locale/zh_TW/LC_MESSAGES/messages.po locale/messages.pot
  $ msgmerge --no-wrap --backup=none --update locale/vi_VN/LC_MESSAGES/messages.po locale/messages.pot

After the update, you maybe need to edit the **PO** files to translate the newly
added strings.


Generate MO_ file for run-time use of web application
+++++++++++++++++++++++++++++++++++++++++++++++++++++

During the run-time of i18n application, the **POT** or **PO** files are not
used. Instead we will generate **MO** files from **PO** files in previous step
for run-time application use. **MO** files are binary message catalog. We can
generate **MO** files by msgfmt_:

.. code-block:: bash

  msgfmt locale/en_US/LC_MESSAGES/messages.po -o locale/en_US/LC_MESSAGES/messages.mo
  msgfmt locale/zh_TW/LC_MESSAGES/messages.po -o locale/zh_TW/LC_MESSAGES/messages.mo
  msgfmt locale/vi_VN/LC_MESSAGES/messages.po -o locale/vi_VN/LC_MESSAGES/messages.mo

These **MO** files are the files we really need in our applications during
run-time.


Use MO_ file in your application
++++++++++++++++++++++++++++++++

The use of **MO** files are supported in different programming languages, such
as Python_ or Go_. I will write another posts to show how to use **MO** files
during run-time.

For Go_ to use gettext_
```````````````````````

Please read [7]_ to see how to use the **PO** and **MO** file in your Go_
application.


----

References:

.. [1] `gettext - GNU Project - Free Software Foundation (FSF) <https://www.gnu.org/software/gettext/>`_

.. [2] `Internationalize a Python application - maemo.org wiki <http://wiki.maemo.org/Internationalize_a_Python_application>`_

.. [3] `Python localization made easy «  Supernifty – nifty stuff <http://www.supernifty.org/blog/2011/09/16/python-localization-made-easy/>`_

.. [4] `localization - I18n strategies for Go with App Engine - Stack Overflow <http://stackoverflow.com/questions/14124630/i18n-strategies-for-go-with-app-engine>`_

.. [5] `Table of locales - MoodleDocs <https://docs.moodle.org/dev/Table_of_locales>`_

.. [6] `default i18n config of webapp2 <http://webapp-improved.appspot.com/api/webapp2_extras/i18n.html#webapp2_extras.i18n.default_config>`_
       (default locale dir of webapp2 i18n is $PROJECT_DIR/locale,
       and default domain of webapp2 i18n is 'messages')

.. [7] `[Golang] Internationalization (i18n) of Go Application by GNU gettext Tools <{filename}../08/golang-i18n-go-application-by-gnu-gettext%en.rst>`_


.. _Pāḷi Dictionary: https://palidictionary.appspot.com/
.. _gettext: https://www.gnu.org/software/gettext/
.. _i18n: https://en.wikipedia.org/wiki/Internationalization_and_localization
.. _xgettext: https://www.gnu.org/software/gettext/manual/html_node/xgettext-Invocation.html
.. _sed: http://www.grymoire.com/Unix/Sed.html
.. _locale: https://en.wikipedia.org/wiki/Locale
.. _msginit: https://www.gnu.org/software/gettext/manual/html_node/msginit-Invocation.html
.. _msgmerge: https://www.gnu.org/software/gettext/manual/html_node/msgmerge-Invocation.html
.. _msgfmt: https://www.gnu.org/software/gettext/manual/html_node/msgfmt-Invocation.html
.. _Python: https://www.python.org/
.. _Go: https://golang.org/
.. _MO: https://www.gnu.org/software/gettext/manual/html_node/MO-Files.html

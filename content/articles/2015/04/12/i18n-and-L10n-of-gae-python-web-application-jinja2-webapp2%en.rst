Internationalization and Localization of Google App Engine Python Webapp Using webapp2 and Jinja2
#################################################################################################

:date: 2015-04-12 05:30
:tags: Python, Web application, web, Google App Engine, Locale, Jinja2, i18n, gettext
:category: Python
:summary: i18n and L10n of Google App Engine Python web application, with
          webapp2, Jinja2, Babel, and pytz.


This post shows how to use webapp2_ web framework and Jinja2_ templating system
to achieve i18n of `Google App Engine Python`_ webapp with Babel_ and pytz_ (or
pytz-appengine_, gae-pytz_).


Install Babel_
++++++++++++++

.. code-block:: bash

  # install babel
  $ sudo pip install babel

  # download babel 1.3
  $ wget https://github.com/mitsuhiko/babel/archive/1.3.zip

  # unzip and enter babel source directory
  $ unzip 1.3.zip
  $ cd babel-1.3/

  # create babel.zip for zipimport
  $ zip -r babel.zip babel/*

  # move babel.zip to your project root directory


Install pytz_ (use gae-pytz_ in this example)
+++++++++++++++++++++++++++++++++++++++++++++

.. code-block:: bash

  # download gaepytz 2011h
  $ wget https://pypi.python.org/packages/source/g/gaepytz/gaepytz-2011h.tar.gz#md5=b7abe173cd98b417fab3e91c1498cdd2

  # unzip and enter gaepytz source directory
  $ tar xvzf gaepytz-2011h.tar.gz
  $ cd gaepytz-2011h/

  # create pytz.zip for zipimport
  $ zip -r pytz.zip pytz/*

  # move pytz.zip to your project root directory


Sample GAE project
++++++++++++++++++

Download the following files to your project root directory:

*i18n.py* (GAE Python webapp): The webapp determine the locale_ by the
`query string`_ in URL.

.. show_github_file:: siongui userpages content/code/i18n-gae-python/i18n.py

*index.html* (HTML template): Wrap the *string* you want to tranlsate in
*{{ _("string") }}*. See [18]_ for more syntax you can use.

.. show_github_file:: siongui userpages content/code/i18n-gae-python/index.html

*app.yaml* (GAE Python config)

.. show_github_file:: siongui userpages content/code/i18n-gae-python/app.yaml

*babel.cfg* (config for Babel_ extraction): This *babel.cfg* tells babel to
extract all translations from all HTML files in your webapp and the encoding of
HTML files are *utf-8*. For more information of the config file, refer to [16]_
and [17]_.

.. show_github_file:: siongui userpages content/code/i18n-gae-python/babel.cfg


Extract and compile translations
++++++++++++++++++++++++++++++++

By default, webapp2_ looks for *pot* and *po* files in *locale* directory under
your project root directory, so first create a directory named *locale*:

.. code-block:: bash

  # in your project root directory:
  $ mkdir locale

Then extract all translations (create *pot* file).

.. code-block:: bash

  # in your project root directory:
  $ pybabel extract -F ./babel.cfg -o ./locale/messages.pot ./

The *pot* file looks like:

.. code-block:: txt

  # Translations template for PROJECT.
  # Copyright (C) 2015 ORGANIZATION
  # This file is distributed under the same license as the PROJECT project.
  # FIRST AUTHOR <EMAIL@ADDRESS>, 2015.
  #
  #, fuzzy
  msgid ""
  msgstr ""
  "Project-Id-Version: PROJECT VERSION\n"
  "Report-Msgid-Bugs-To: EMAIL@ADDRESS\n"
  "POT-Creation-Date: 2015-04-12 03:32+0800\n"
  "PO-Revision-Date: YEAR-MO-DA HO:MI+ZONE\n"
  "Last-Translator: FULL NAME <EMAIL@ADDRESS>\n"
  "Language-Team: LANGUAGE <LL@li.org>\n"
  "MIME-Version: 1.0\n"
  "Content-Type: text/plain; charset=utf-8\n"
  "Content-Transfer-Encoding: 8bit\n"
  "Generated-By: Babel 1.3\n"

  #: index.html:8
  msgid "home"
  msgstr ""

  #: index.html:10
  msgid "about"
  msgstr ""

  #: index.html:12
  msgid "'link'"
  msgstr ""


Then initialize the directory for each locale_ that your webapp will support.
*en_US* and *zh_TW* are supported in our example. See [19]_ for table of
locales.

.. code-block:: bash

  # in your project root directory:
  $ pybabel init -l en_US -d ./locale -i ./locale/messages.pot
  $ pybabel init -l zh_TW -d ./locale -i ./locale/messages.pot

Two *po* files (``locale/en_US/LC_MESSAGES/messages.po`` and
``locale/zh_TW/LC_MESSAGES/messages.po``) are created. You do not need to do
anything with the *en_US* po file because English is default language.
Translate only non-default-language *po* files. In our exmaple, the *zh_TW* *po*
file after translation looks like:

.. code-block:: txt

  # Chinese (Traditional, Taiwan) translations for PROJECT.
  # Copyright (C) 2015 ORGANIZATION
  # This file is distributed under the same license as the PROJECT project.
  # FIRST AUTHOR <EMAIL@ADDRESS>, 2015.
  #
  #, fuzzy
  msgid ""
  msgstr ""
  "Project-Id-Version: PROJECT VERSION\n"
  "Report-Msgid-Bugs-To: EMAIL@ADDRESS\n"
  "POT-Creation-Date: 2015-04-12 03:32+0800\n"
  "PO-Revision-Date: 2015-04-12 03:35+0800\n"
  "Last-Translator: FULL NAME <EMAIL@ADDRESS>\n"
  "Language-Team: zh_Hant_TW <LL@li.org>\n"
  "Plural-Forms: nplurals=2; plural=(n != 1)\n"
  "MIME-Version: 1.0\n"
  "Content-Type: text/plain; charset=utf-8\n"
  "Content-Transfer-Encoding: 8bit\n"
  "Generated-By: Babel 1.3\n"

  #: index.html:8
  msgid "home"
  msgstr "首頁"

  #: index.html:10
  msgid "about"
  msgstr "關於"

  #: index.html:12
  msgid "'link'"
  msgstr "'連結'"

After all translations done, compile *po* file with the following command:

.. code-block:: bash

  # in your project root directory:
  $ pybabel compile -f -d ./locale

Now we can run this GAE Python webapp, and then open the browser with URL:

  http://localhost:8080/

You will see the webpage in default language. Then open the browser with URL:

  http://localhost:8080/?locale=zh_TW

You will see the webpage in Traditional Chinese.


Update translations
+++++++++++++++++++

When the strings to be translated change, re-create *pot* file:

.. code-block:: bash

  # in your project root directory:
  $ pybabel extract -F ./babel.cfg -o ./locale/messages.pot ./

Then update each *locale*:

.. code-block:: bash

  # in your project root directory:
  $ pybabel update -l en_US -d ./locale/ -i ./locale/messages.pot
  $ pybabel update -l zh_TW -d ./locale/ -i ./locale/messages.pot

Again, translate the strings in each *po* file, and then compile again:

.. code-block:: bash

  # in your project root directory:
  $ pybabel compile -f -d ./locale

----

References:

.. [1] `Internationalization and localization with webapp2 <http://webapp-improved.appspot.com/tutorials/i18n.html>`_

.. [2] `python - How to enable {% trans %} tag for jinja templates? - Stack Overflow <http://stackoverflow.com/questions/8471455/how-to-enable-trans-tag-for-jinja-templates>`_

.. [3] `I18N support · Issue #92 · getpelican/pelican · GitHub <https://github.com/getpelican/pelican/issues/92>`_

.. [4] `python - i18n with jinja2 + GAE - Stack Overflow <http://stackoverflow.com/questions/7961800/i18n-with-jinja2-gae>`_

.. [5] `Enable jinja2 and i18n translations on Google AppEngine | Mikhail Shilkov <http://mikhail.io/2012/07/26/enable-jinja2-and-i18n-translations-on-google-appengine/>`_

.. [6] `How to use i18n from webapp2_extras? - Google Groups <https://groups.google.com/d/topic/google-appengine-python/RhXxIOfnfm0>`_

.. [7] `google app engine - Internationalization with python gae, babel and i18n. Can't output the correct string - Stack Overflow <http://stackoverflow.com/questions/14414960/internationalization-with-python-gae-babel-and-i18n-cant-output-the-correct-s>`_

.. [8] `Internationalization and localization - Wikipedia, the free encyclopedia <http://en.wikipedia.org/wiki/Internationalization_and_localization>`_

.. [9] `python - How to import modules in Google App Engine? - Stack Overflow <http://stackoverflow.com/questions/2710861/how-to-import-modules-in-google-app-engine>`_

.. [10] `Max number of files and blobs is 1000 - Google Code <https://code.google.com/p/googleappengine/issues/detail?id=161>`_

.. [11] `Moon blue diary: Using zipped pytz on GAE <http://takashi-matsuo.blogspot.com/2008/07/using-zipped-pytz-on-gae.html>`_

.. [12] `Moon blue diary: Using the newest zipped pytz on GAE <http://takashi-matsuo.blogspot.com/2008/07/using-newest-zipped-pytz-on-gae.html>`_

.. [13] `brianmhunt/pytz-appengine · GitHub <https://github.com/brianmhunt/pytz-appengine>`_

.. [14] `Babel (old site) <http://babel.edgewall.org/>`_

.. [15] `gaepytz on Python Package Index <https://pypi.python.org/pypi/gaepytz>`_

.. [16] `Babel Integration - Jinja2 Documentation <http://jinja.pocoo.org/docs/dev/integration/#babel-integration>`_

.. [17] `Extraction Method Mapping and Configuration - Working with Message Catalogs - Babel 1.0 documentation <http://babel.pocoo.org/docs/messages/#extraction-method-mapping-and-configuration>`_

.. [18] `i18n - Template Designer Documentation - Jinja2 Documentation <http://jinja.pocoo.org/docs/dev/templates/#i18n>`_

.. [19] `Table of locales - MoodleDocs <https://docs.moodle.org/dev/Table_of_locales>`_



.. _Google App Engine Python: https://cloud.google.com/appengine/docs/python/

.. _pytz: http://pytz.sourceforge.net/

.. _gae-pytz: https://code.google.com/p/gae-pytz/

.. _webapp2: https://webapp-improved.appspot.com/

.. _Jinja2: http://jinja.pocoo.org/docs/dev/

.. _Babel: http://babel.pocoo.org/

.. _pytz-appengine: https://github.com/brianmhunt/pytz-appengine

.. _query string: http://en.wikipedia.org/wiki/Query_string

.. _locale: http://en.wikipedia.org/wiki/Locale

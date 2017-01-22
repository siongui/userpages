[Pelican] Localize Theme via Jinja2 Custom Filter
#################################################

:date: 2017-01-12T22:06+08:00
:tags: Python, Pelican, Jinja2, i18n, gettext, html
:category: Python
:summary: Pelican_ static site generator - Localize theme_ with i18n_subsites_
          plugin via `custom Jinja2 filter`_. Implementation of gettext_-like
          filter.
:adsu: yes


Introduction
++++++++++++

Pelican_ static site generator and i18n_subsites_ plugin can help you build
website or blog which supports multiple languages. When generating websites for
each supported language, some strings in the theme_ need to be translated
according to `default language`_ while i18n_subsites_ plugin generates sub-site
for each supported language. There are several ways to localize theme_:

1. Use Jinja2_ i18n extension in GNU gettext_ way as [2]_:

   This way is good for big projects, but for small projects it is too tedious.

2. Implement a gettext_ macro_ in theme_ as [3]_:

   This way needs to put the code in the HTML theme_, which is not good in terms
   of overall architecture.

Instead, this post shows another way to implement gettext_-like filter in
``pelicanconf.py`` and ust it to localize theme_ (translate strings in the
theme_).


Solution
++++++++

First implement a custom filter named *gettext* in ``pelicanconf.py`` as
follows. Note that strings to be translated and translations are included in
this filter, in this case ``en`` is default locale, ``zh`` and ``th`` are
supported locales except ``en``.

.. code-block:: python

  # custom Jinja2 filter for localizing theme
  def gettext(string, lang):
      if lang == "en":
          return string
      elif lang == "zh":
          if string == "Archives": return "歸檔"
          elif string == "Categories": return "分類"
          elif string == "Category": return "分類"
          elif string == "Authors": return "作者"
          elif string == "Author": return "作者"
          elif string == "Tags": return "標籤"
          elif string == "Updated": return "更新"
          elif string == "Translation(s)": return "翻譯"
          elif string == "Edit on Github": return "在Github上編輯"
          else: return string
      elif lang == "th":
          if string == "Archives": return "สารบรรณ"
          elif string == "Categories": return "ประเภท"
          elif string == "Category": return "ประเภท"
          elif string == "Authors": return "ผู้เขียน"
          elif string == "Author": return "ผู้เขียน"
          elif string == "Tags": return "แท็ก"
          elif string == "Updated": return "การปรับปรุง"
          elif string == "Translation(s)": return "การแปล"
          elif string == "Edit on Github": return "แก้ไขที่ Github"
          else: return string
      else:
          return string

  JINJA_FILTERS = {
      "gettext": gettext,
  }


Strings to be localized in the theme_ are marked as follows, and i18n_subsites_
plugin will replace strings according to `DEFAULT_LANG setting`_ while
generating subsites of each language.

.. code-block:: html

  <!-- mark all the strings to be localized ... -->
  <span>{{ 'Archives'|gettext(DEFAULT_LANG) }}</span>
  <span>{{ 'Updated'|gettext(DEFAULT_LANG) }}</span>
  <span>{{ 'Author'|gettext(DEFAULT_LANG) }}</span>
  <!-- ... -->

For full working example, see [1]_.

----

Tested on: ``Ubuntu Linux 16.10``, ``Python 2.7.12+``, `Pelican 3.7.0`_.

----

References:

.. [1] `use Jinja2 custom filter to localize theme · siongui/pelican-template@b7dcda2 · GitHub <https://github.com/siongui/pelican-template/commit/b7dcda254f4b1b6a8856679b24a4bdaed7de97e5>`_

.. [2] `Localizing themes with Jinja2 - i18n_subsites - pelican-plugins <https://github.com/getpelican/pelican-plugins/blob/master/i18n_subsites/localizing_using_jinja2.rst>`_

.. [3] `[Pelican] Translate String According to Default Language in Theme <{filename}../07/pelican-translate-string-according-to-default-language-in-theme%en.rst>`_

.. [4] `[Pelican] Get Single Page or Article by slug Metadata in Theme <{filename}../08/pelican-get-single-page-or-article-by-slug-metadata-in-theme%en.rst>`_

.. _Python: https://www.python.org/
.. _gettext: https://www.google.com/search?q=gettext
.. _Pelican: http://blog.getpelican.com/
.. _Pelican 3.7.0: http://docs.getpelican.com/en/3.7.0/
.. _i18n_subsites: https://github.com/getpelican/pelican-plugins/tree/master/i18n_subsites
.. _theme: http://docs.getpelican.com/en/latest/themes.html
.. _macro: http://jinja.pocoo.org/docs/dev/templates/#macros
.. _default language: http://docs.getpelican.com/en/latest/settings.html#translations
.. _settings: http://docs.getpelican.com/en/latest/settings.html
.. _DEFAULT_LANG setting: http://docs.getpelican.com/en/latest/settings.html#translations
.. _custom Jinja2 filter: http://jinja.pocoo.org/docs/latest/api/#custom-filters
.. _pelicanconf.py: http://docs.getpelican.com/en/latest/settings.html
.. _Jinja2: https://www.google.com/search?q=jinja2

[Pelican] Translate String According to Default Language in Theme
#################################################################

:date: 2017-01-07T20:49+08:00
:tags: Python, Pelican, Jinja2, i18n, gettext, html
:category: Python
:summary: Pelican_ static site generator - Support multilingual in Pelican
          theme_. Translate strings according to `default language`_ of
          settings_ in the theme. (implement macro_ gettext_)
:adsu: yes


Pelican_ static site generator and i18n_subsites_ plugin can help you build
website or blog which supports multiple languages. When generating websites for
each supported language, some strings in the theme_ need to be translated
according to `default language`_ while i18n_subsites_ plugin generates sub-site
for each supported language. This post shows how to implement a gettext_ macro_
to achieve this goal.

Show by example. Assume three languages ``en``, ``zh``, and ``th`` are
supported. Strings to be translated in the theme_ are *Archives*, *Authors*, and
*Categories*.

Then put the following code, which includes strings to be translated and
translation in the theme_. (save as ``layout/includes/i18n.html`` in the theme_
directory in this case):

.. code-block:: html

  {% macro gettext(string, lang) -%}
    {%- if lang == 'en' -%}
      {{ string }}
    {%- elif lang == 'zh' -%}
      {%- if string == 'Archives' -%} 歸檔
      {%- elif string == 'Categories' -%} 分類
      {%- elif string == 'Authors' -%} 作者
      {%- else -%}
        {{ string }}
      {%- endif -%}
    {%- elif lang == 'th' -%}
      {%- if string == 'Archives' -%} สารบรรณ
      {%- elif string == 'Categories' -%} ประเภท
      {%- elif string == 'Authors' -%} ผู้เขียน
      {%- else -%}
        {{ string }}
      {%- endif -%}
    {%- else -%}
      {{ string }}
    {%- endif -%}
  {%- endmacro %}


Next, in the layout, which is the template that every other template is
inherited from (``layout/layout.html`` in this case), import the above gettext
macro_ in the beginning of the layout:

.. code-block:: html

  {%- from 'layout/includes/i18n.html' import gettext -%}

Finally, strings in the theme_ are translated like the following:

.. code-block:: html

  <span>{{ gettext('Archives', DEFAULT_LANG) }}</span>
  <span>{{ gettext('Categories', DEFAULT_LANG) }}</span>
  <span>{{ gettext('Authors', DEFAULT_LANG) }}</span>


For full working example, see [1]_.

For alternatives to localize theme_, see [2]_ and [3]_.

----

Tested on: ``Ubuntu Linux 15.10``, ``Python 2.7.10``, `Pelican 3.6.3`_.

----

References:

.. [1] `GitHub - siongui/pelican-template <https://github.com/siongui/pelican-template>`_

.. [2] `Localizing themes with Jinja2 - i18n_subsites - pelican-plugins <https://github.com/getpelican/pelican-plugins/blob/master/i18n_subsites/localizing_using_jinja2.rst>`_

.. [3] `[Pelican] Localize Theme via Jinja2 Custom Filter <{filename}../12/pelican-localize-theme-via-jinja2-custom-filter%en.rst>`_


.. _Python: https://www.python.org/
.. _gettext: https://www.google.com/search?q=gettext
.. _Pelican: http://blog.getpelican.com/
.. _Pelican 3.6.3: http://docs.getpelican.com/en/3.6.3/
.. _i18n_subsites: https://github.com/getpelican/pelican-plugins/tree/master/i18n_subsites
.. _theme: http://docs.getpelican.com/en/latest/themes.html
.. _macro: http://jinja.pocoo.org/docs/dev/templates/#macros
.. _default language: http://docs.getpelican.com/en/latest/settings.html#translations
.. _settings: http://docs.getpelican.com/en/latest/settings.html

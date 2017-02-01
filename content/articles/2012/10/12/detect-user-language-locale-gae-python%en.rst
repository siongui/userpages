Detect User Language (Locale) on Google App Engine Python
#########################################################

:date: 2012-10-12 02:30
:modified: 2015-02-21 07:48
:tags: Python, Locale, HTTP Header, Google App Engine, Accept-Language
:category: Python
:summary: Detect user locale/language from HTTP request header on Google App
          Engine for Python.
:adsu: yes


The best way to guess/detect/determine the locale/language of users is probably
by using the Accept-Language_ field in HTTP request header (see reference [1]_
for detail). In this post, we will show how to detect/determine user locale
according to Accept-Language_ field in HTTP request header on `Google App Engine
for Python`_.

Assume that the website we own support three locales: *de (German)*, *en_US
(English, United States)*, and *zh_TW (Traditional Chinese)*. If these three
locales matches the locales in Accept-Language_ field, the website will serve
the web pages with the locale of higher preference. If no matches, the website
will serve web page with *en_US* locale (default locale).

The first thing to do to detect/determine user locale is parsing
Accept-Language_ field. In previous post "*[Python] Parse Accept-Language in
HTTP Request Header*" [2]_, this topic is discussed. Please refer to the post
for detail.

Next, the parsed data will be compared with supported locales of the website.
The following code snippet demonstrates how to do this:

.. adsu:: 2

Detect Locales on Google App Engine Python
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

.. code-block:: python

  def parseAcceptLanguage(acceptLanguage):
    languages = acceptLanguage.split(",")
    locale_q_pairs = []

    for language in languages:
      if language.split(";")[0] == language:
        # no q => q = 1
        locale_q_pairs.append((language.strip(), "1"))
      else:
        locale = language.split(";")[0].strip()
        q = language.split(";")[1].split("=")[1]
        locale_q_pairs.append((locale, q))

    return locale_q_pairs


  def detectLocale(acceptLanguage):
    defaultLocale = 'en_US'
    supportedLocales = ['de', 'en_US', 'zh_TW']

    locale_q_pairs = parseAcceptLanguage(acceptLanguage)
    for pair in locale_q_pairs:
      for locale in supportedLocales:
        # pair[0] is locale, pair[1] is q value
        if pair[0].replace('-', '_').lower().startswith(locale.lower()):
          return locale

    return defaultLocale

The *parseAcceptLanguage* function is the same as the function in previous post
[2]_ for parsing Accept-Language_ field. The *detectLocale* function compares
supported locales of the website with locales in Accept-Language_ field. Note
that the default locale is set to *en_US*. You can change the *defaultLocale*
and *supportedLocales* variables to fit the locales design in your website.

.. adsu:: 3

Example Usage
~~~~~~~~~~~~~

If you are using webapp2_, pass
:code:`self.request.headers.get('accept_language')`, the variable which contains
Accept-Language_ field on `Google App Engine for Python`_, to *detectLocale*
function and the locale to be served will be returned. The following example
will give you idea of what it looks like:

.. code-block:: python

  >>> print(parseAcceptLanguage('da, en-gb;q=0.8, en;q=0.7'))
  >>> print(detectLocale('da, en-gb;q=0.8, en;q=0.7'))
  [('da', '1'), ('en-gb', '0.8'), ('en', '0.7')]
  en_US

  >>> print(parseAcceptLanguage('zh, en-us; q=0.8, en; q=0.6'))
  >>> print(detectLocale('zh, en-us; q=0.8, en; q=0.6'))
  [('zh', '1'), ('en-us', '0.8'), ('en', '0.6')]
  en_US

  >>> print(parseAcceptLanguage('zh-tw, en-us; q=0.8, en; q=0.6'))
  >>> print(detectLocale('zh-tw, en-us; q=0.8, en; q=0.6'))
  [('zh-tw', '1'), ('en-us', '0.8'), ('en', '0.6')]
  zh_TW

  >>> print(parseAcceptLanguage('es-mx, es, en'))
  >>> print(detectLocale('es-mx, es, en'))
  [('es-mx', '1'), ('es', '1'), ('en', '1')]
  en_US

  >>> print(parseAcceptLanguage('de; q=1.0, en; q=0.5'))
  >>> print(detectLocale('de; q=1.0, en; q=0.5'))
  [('de', '1.0'), ('en', '0.5')]
  de

----

References:

.. [1] `Accept-Language used for locale setting <http://www.w3.org/International/questions/qa-accept-lang-locales.en.php>`_

.. [2] `[Python] Parse Accept-Language in HTTP Request Header <{filename}../11/python-parse-accept-language-in-http-request-header%en.rst>`_

.. _Accept-Language: http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html

.. _Google App Engine for Python: https://cloud.google.com/appengine/docs/python/

.. _webapp2: https://cloud.google.com/appengine/docs/python/gettingstartedpython27/usingwebapp

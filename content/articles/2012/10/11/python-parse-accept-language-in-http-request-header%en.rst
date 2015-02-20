[Python] Parse Accept-Language in HTTP Request Header
#####################################################

:date: 2012-10-11 02:36
:modified: 2015-02-21 07:10
:tags: Python, Locale, Google App Engine, HTTP Header, Accept-Language
:category: Python
:summary: Parse Accept-Language in HTTP Request Header in Python.


The Accept-Language_ field in `HTTP Request header`_ is important for
`determining user locale`_. In this post, we will show how to parse
*Accept-Language* field using Python. The following *parseAcceptLanguage*
function takes *Accept-Language* field as argument, and returns a Python list
which contains pairs of (locale, q). The order of (locale, q) pairs are
determined by locale preference of the user (highest preference first).

Python Parse Accept-Language
~~~~~~~~~~~~~~~~~~~~~~~~~~~~

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

Example Usage and Output
~~~~~~~~~~~~~~~~~~~~~~~~

.. code-block:: python

  >>> print(parseAcceptLanguage('da, en-gb;q=0.8, en;q=0.7'))
  [('da', '1'), ('en-gb', '0.8'), ('en', '0.7')]

  >>> print(parseAcceptLanguage('zh, en-us; q=0.8, en; q=0.6'))
  [('zh', '1'), ('en-us', '0.8'), ('en', '0.6')]

  >>> print(parseAcceptLanguage('es-mx, es, en'))
  [('es-mx', '1'), ('es', '1'), ('en', '1')]

  >>> print(parseAcceptLanguage('de; q=1.0, en; q=0.5'))
  [('de', '1.0'), ('en', '0.5')]

To know how to detect user *locale/language* according to *Accept-Language*
field in HTTP request header, please refer to the post "*Detect User Language
(Locale) on Google App Engine Python*" [4]_.

----

References:

.. [1] `List of HTTP header fields <http://en.wikipedia.org/wiki/List_of_HTTP_header_fields>`_

.. [2] `HTTP/1.1: Header Field Definitions <http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html>`_

.. [3] `Accept-Language used for locale setting <http://www.w3.org/International/questions/qa-accept-lang-locales.en.php>`_

.. [4] `Detect User Language (Locale) on Google App Engine Python <{filename}../12/detect-user-language-locale-gae-python%en.rst>`_

.. _Accept-Language: http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html

.. _HTTP Request Header: http://en.wikipedia.org/wiki/List_of_HTTP_header_fields

.. _determining user locale: http://www.w3.org/International/questions/qa-accept-lang-locales.en.php

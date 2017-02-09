[JavaScript] Parse Accept-Language in HTTP Request Header
#########################################################

:date: 2016-01-22T19:20+08:00
:tags: JavaScript, String Manipulation, HTTP Header, Accept-Language, Locale
:category: JavaScript
:summary: Parse `Accept-Language`_ in `HTTP Request Header`_ in JavaScript_ Programming Language.
:adsu: yes


This post is the JavaScript_ version of my previous posts [1]_ and [3]_. This
JavaScript_ example parses `Accept-Language`_ string in `HTTP Request Header`_
and output (languageTag, quality) pairs:

.. code-block:: javascript

    /**
     * Parse HTTP accept-language header of the user browser.
     *
     * @param {string} hdr The string of accpet-language header
     * @return {Array} Array of language-quality pairs
     */
    function getParsedAcceptLangs(hdr) {
      var pairs = hdr.split(',');
      var result = [];
      for (var i=0; i < pairs.length; i++) {
        var pair = pairs[i].split(';');
        if (pair.length == 1) result.push( [pair[0], '1'] );
        else result.push( [pair[0], pair[1].split('=')[1] ] );
      }
      return result;
    }

You can use this information to `determine user locale`_.

.. adsu:: 2

----

References:

.. [1] `[Python] Parse Accept-Language in HTTP Request Header <{filename}../../../2012/10/11/python-parse-accept-language-in-http-request-header%en.rst>`_

.. [2] `Detect User Language (Locale) on Google App Engine Python <{filename}../../../2012/10/12/detect-user-language-locale-gae-python%en.rst>`_

.. [3] `[Golang] Parse Accept-Language in HTTP Request Header <{filename}../../../2015/02/22/go-parse-accept-language%en.rst>`_

.. [4] `List of HTTP header fields <http://en.wikipedia.org/wiki/List_of_HTTP_header_fields>`_

.. [5] `HTTP/1.1: Header Field Definitions <http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html>`_

.. [6] `Accept-Language used for locale setting <http://www.w3.org/International/questions/qa-accept-lang-locales.en.php>`_
.. adsu:: 3
.. [7] `localization - JavaScript for detecting browser language preference - Stack Overflow <http://stackoverflow.com/questions/1043339/javascript-for-detecting-browser-language-preference>`_

.. [8] `NavigatorLanguage.languages - Web APIs | MDN <https://developer.mozilla.org/en-US/docs/Web/API/NavigatorLanguage/languages>`_

.. [9] `Accessing the web page's HTTP Headers in JavaScript - Stack Overflow <http://stackoverflow.com/questions/220231/accessing-the-web-pages-http-headers-in-javascript>`_

.. [10] `How do I access the HTTP request header fields via JavaScript? - Stack Overflow <http://stackoverflow.com/questions/220149/how-do-i-access-the-http-request-header-fields-via-javascript>`_

.. [11] `XML DOM - HttpRequest object <http://www.w3schools.com/xml/dom_http.asp>`_


.. _JavaScript: http://www.w3schools.com/js/

.. _Accept-Language: http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html

.. _HTTP Request Header: http://en.wikipedia.org/wiki/List_of_HTTP_header_fields

.. _determine user locale: http://www.w3.org/International/questions/qa-accept-lang-locales.en.php

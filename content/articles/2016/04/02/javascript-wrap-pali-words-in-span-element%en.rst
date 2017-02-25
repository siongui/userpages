[JavaScript] Wrap Pāli Words in Span Element
############################################

:date: 2016-04-02T21:23+08:00
:tags: JavaScript, html, Regular Expression, String Manipulation
:category: JavaScript
:summary: Wrap `Pāli`_ words in span_ element via `regular expression`_ and
          JavaScript_.
:og_image: http://www.javatpoint.com/images/javascript/javascript_logo.png
:adsu: yes


Wrap `Pāli`_ words in span_ element via `regular expression`_ and JavaScript_.

.. code-block:: javascript

  function markInSpan(htmlString) {
    return string.replace(/[AaBbCcDdEeGgHhIiJjKkLlMmNnOoPpRrSsTtUuVvYyĀāĪīŪūṀṁṂṃŊŋṆṇṄṅÑñṬṭḌḍḶḷ]+/g, '<span>$&</span>');
  }

.. adsu:: 2

----

References:

.. [1] `javascript string replace <https://www.google.com/search?q=javascript+string+replace>`_


.. _JavaScript: https://www.google.com/search?q=javascript
.. _Pāli: https://en.wikipedia.org/wiki/Pali
.. _span: http://www.w3schools.com/tags/tag_span.asp
.. _regular expression: https://www.google.com/search?q=regular+expression

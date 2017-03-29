Pure CSS Toggle (Show/Hide) HTML Element
########################################

:date: 2017-02-27T02:23+08:00
:tags: Pure CSS Toggle, toggle, toggleable, html, CSS
:category: CSS
:summary: CSS_ only toggle (show/hide) HTML_ `DOM element`_ without JavaScript_.
:og_image: https://www.sololearn.com/Icons/Courses/1023.png
:adsu: yes


.. role:: raw-html(raw)
   :format: html

CSS_ only toggle (show/hide) HTML_ `DOM element`_ without JavaScript_.
See demo first. Please click the following :raw-html:`&#8803;` to toggle:

.. raw:: html

  <style>
  label[for=element-toggle] {
    cursor: pointer;
  }
  #element-toggle {
    display: none;
  }
  #element-toggle:not(:checked) ~ #toggled-element {
    display: none;
  }
  </style>

  <div style="padding: 2em;">
  <label for="element-toggle">&#8803;</label>
  <input id="element-toggle" type="checkbox" />
  <ul id="toggled-element">
    <li><a href="https://www.google.com/">Google</a></li>
    <li><a href="https://duckduckgo.com/">DuckDuckGo</a></li>
    <li><a href="https://www.ecosia.org/">Ecosia</a></li>
  </ul>
  </div>


There are three HTML elements in the demo:

.. code-block:: html

  <label for="element-toggle">&#8803;</label>
  <input id="element-toggle" type="checkbox" />
  <ul id="toggled-element">
    <li><a href="https://www.google.com/">Google</a></li>
    <li><a href="https://duckduckgo.com/">DuckDuckGo</a></li>
    <li><a href="https://www.ecosia.org/">Ecosia</a></li>
  </ul>

1. A visible HTML label_ element
2. A invisible HTML input_ checkbox element, referenced by the above label_
   element via for_ attribute.
3. The element to be toggled. This could be any element, such as *div* or
   *span*, and it is *ul* element in this demo.

.. adsu:: 2

When users click on the visible label element, the invisible input checkbox
element becomes checked/unchecked by the click. Then the following CSS rules are
used to show or hide the element to be toggled according to whether the checkbox
is checked or not.

.. code-block:: css

  label[for=element-toggle] {
    cursor: pointer;
  }
  #element-toggle {
    display: none;
  }
  #element-toggle:not(:checked) ~ #toggled-element {
    display: none;
  }

There are three rules in above CSS_ code. The last rule is the most important
part in this trick. It says if the element with *id=element-toggle* (i.e., the
input checkbox) is not checked, the adjacent element with *id=toggled-element*
(*ul* element in this case) is hidden. If the input checkbox is checked, this
rule is not applied to the *ul* element, which makes the *ul* element visible.
The last CSS rule shows or hides the element to be toggled according to whether
the checkbox is checked or not.

.. adsu:: 3

This technique can be used to toggle nav menu in mobile screen [7]_, and there
is no JavaScript_ code to slow down the speed of the website.

----

Tested on:

- ``Chromium Version 55.0.2883.87 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``

----

References:

.. [1] | `css only nav toggle - Google search <https://www.google.com/search?q=css+only+nav+toggle>`_
       | `css only nav toggle - DuckDuckGo search <https://duckduckgo.com/?q=css+only+nav+toggle>`_
       | `css only nav toggle - Ecosia search <https://www.ecosia.org/search?q=css+only+nav+toggle>`_
       | `css only nav toggle - Qwant search <https://www.qwant.com/?q=css+only+nav+toggle>`_
       | `css only nav toggle - Bing search <https://www.bing.com/search?q=css+only+nav+toggle>`_
       | `css only nav toggle - Yahoo search <https://search.yahoo.com/search?p=css+only+nav+toggle>`_
       | `css only nav toggle - Baidu search <https://www.baidu.com/s?wd=css+only+nav+toggle>`_
       | `css only nav toggle - Yandex search <https://www.yandex.com/search/?text=css+only+nav+toggle>`_

.. [2] `CSS only menu toggle - no JavaScript required <http://www.outofscope.com/css-only-menu-toggle-no-javascript-required/>`_
.. adsu:: 4
.. [3] | `css only toggle element - Google search <https://www.google.com/search?q=css+only+toggle+element>`_
       | `css only toggle element - DuckDuckGo search <https://duckduckgo.com/?q=css+only+toggle+element>`_
       | `css only toggle element - Ecosia search <https://www.ecosia.org/search?q=css+only+toggle+element>`_
       | `css only toggle element - Qwant search <https://www.qwant.com/?q=css+only+toggle+element>`_
       | `css only toggle element - Bing search <https://www.bing.com/search?q=css+only+toggle+element>`_
       | `css only toggle element - Yahoo search <https://search.yahoo.com/search?p=css+only+toggle+element>`_
       | `css only toggle element - Baidu search <https://www.baidu.com/s?wd=css+only+toggle+element>`_
       | `css only toggle element - Yandex search <https://www.yandex.com/search/?text=css+only+toggle+element>`_

.. [4] | `Creating a toggle feature using :checked pseudo class. | Clearleft <http://clearleft.com/thinks/creatingatogglefeatureusingcheckedpseudoclass/>`_
       | `:checked | CSS-Tricks <https://css-tricks.com/almanac/selectors/c/checked/>`_

.. [5] | `HTML label - Google search <https://www.google.com/search?q=HTML+label>`_
       | `HTML label - DuckDuckGo search <https://duckduckgo.com/?q=HTML+label>`_
       | `HTML label - Ecosia search <https://www.ecosia.org/search?q=HTML+label>`_
       | `HTML label - Qwant search <https://www.qwant.com/?q=HTML+label>`_
       | `HTML label - Bing search <https://www.bing.com/search?q=HTML+label>`_
       | `HTML label - Yahoo search <https://search.yahoo.com/search?p=HTML+label>`_
       | `HTML label - Baidu search <https://www.baidu.com/s?wd=HTML+label>`_
       | `HTML label - Yandex search <https://www.yandex.com/search/?text=HTML+label>`_

.. [6] | `html special characters - Google search <https://www.google.com/search?q=html+special+characters>`_
       | `HTML Special Characters - Quackit Tutorials <http://www.quackit.com/html/html_special_characters.cfm>`_
       | `List of Unicode Characters - Quackit Tutorials <http://www.quackit.com/character_sets/unicode/>`_
       | `Unicode 9.0 Characters: Mathematical Operators - Quackit Tutorials <http://www.quackit.com/character_sets/unicode/versions/unicode_9.0.0/mathematical_operators_unicode_character_codes.cfm>`_
       | `Unicode Name: STRICTLY EQUIVALENT TO - Quackit Tutorials <http://www.quackit.com/html/html_editors/scratchpad/?app=charset_ref&hexadecimal=02263&decimal=8803&unicodeName=STRICTLY_EQUIVALENT_TO>`_
.. adsu:: 5
.. [7] `CSS Only Bulma Responsive Nav Bar (Navigation Bar) <{filename}../22/css-only-bulma-responsive-navbar%en.rst>`_

.. _HTML: https://www.google.com/search?q=HTML
.. _CSS: https://www.google.com/search?q=CSS
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _DOM element: https://www.google.com/search?q=DOM+element
.. _label: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label
.. _input: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/checkbox
.. _for: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label#Using_the_for_attribute

純CSS可切換選單
###############

:date: 2017-02-21T23:52+08:00
:tags: CSS, Web開發
:category: CSS
:summary: 純 CSS_ 可切換(toggle)選單(menu)，不需要任何 JavaScript_
:og_image: http://cssmenumaker.com/sites/default/files/styles/new_menu_thumb/public/menu/90/css_menu_thumb.png
:adsu: yes

.. role:: raw-html(raw)
   :format: html

不囉唆，先看範例demo，請點擊下方的 :raw-html:`&#8803;`

.. raw:: html

  <style>
    label {
      cursor: pointer;
    }
    #menu-toggle-demo {
      display: none;
    }
    #menu {
      display: none;
    }
    #menu-toggle-demo:checked + #menu {
      display: block;
    }
  </style>

  <div style="padding: 2em;">
  <label for="menu-toggle-demo">&#8803;</label>
  <input type="checkbox" id="menu-toggle-demo"/>
  <ul id="menu">
    <li><a href="https://www.google.com/">Google</a></li>
    <li><a href="https://duckduckgo.com/">DuckDuckGo</a></li>
    <li><a href="https://www.ecosia.org/">Ecosia</a></li>
  </ul>
  </div>

點擊一次會顯示選單，再點擊一次會隱藏選單，反覆點擊就會反覆地顯示隱藏。
此技巧來自 [2]_ ，我稍微修改使之感覺更順眼。
這個技巧神奇就在於它是純 CSS_ 實現，不需要任何的 JavaScript_
程式碼來控制，可以讓網頁載入速度更快，減低網站複雜度，
整個實作的程式碼不超過20行（程式碼請見下方）
如果選單設計不複雜的話，這是一個很不錯的實作方式，
我很喜歡這樣的技巧。

.. adsu:: 2

**HTML** 程式碼：

.. code-block:: html

  <label for="menu-toggle">&#8803;</label>
  <input type="checkbox" id="menu-toggle"/>
  <ul id="menu">
    <li><a href="https://www.google.com/">Google</a></li>
    <li><a href="https://duckduckgo.com/">DuckDuckGo</a></li>
    <li><a href="https://www.ecosia.org/">Ecosia</a></li>
  </ul>

**CSS** 程式碼：

.. code-block:: css

  label {
    cursor: pointer;
  }
  #menu-toggle {
    display: none;
  }
  #menu {
    display: none;
  }
  #menu-toggle:checked + #menu {
    display: block;
  }

.. adsu:: 3

程式不多做解釋了，有興趣了解細節請自行參閱下方的參考。

----

參考：

.. [1] | `css only nav toggle - Google search <https://www.google.com/search?q=css+only+nav+toggle>`_
       | `css only nav toggle - DuckDuckGo search <https://duckduckgo.com/?q=css+only+nav+toggle>`_
       | `css only nav toggle - Ecosia search <https://www.ecosia.org/search?q=css+only+nav+toggle>`_
       | `css only nav toggle - Bing search <https://www.bing.com/search?q=css+only+nav+toggle>`_
       | `css only nav toggle - Yahoo search <https://search.yahoo.com/search?p=css+only+nav+toggle>`_
       | `css only nav toggle - Baidu search <https://www.baidu.com/s?wd=css+only+nav+toggle>`_
       | `css only nav toggle - Yandex search <https://www.yandex.com/search/?text=css+only+nav+toggle>`_

.. [2] `CSS only menu toggle - no JavaScript required <http://www.outofscope.com/css-only-menu-toggle-no-javascript-required/>`_

.. [3] | `toggle menu 中文 - Google search <https://www.google.com/search?q=toggle+menu+%E4%B8%AD%E6%96%87>`_
       | `toggle menu 中文 - DuckDuckGo search <https://duckduckgo.com/?q=toggle+menu+%E4%B8%AD%E6%96%87>`_
       | `toggle menu 中文 - Ecosia search <https://www.ecosia.org/search?q=toggle+menu+%E4%B8%AD%E6%96%87>`_
       | `toggle menu 中文 - Bing search <https://www.bing.com/search?q=toggle+menu+%E4%B8%AD%E6%96%87>`_
       | `toggle menu 中文 - Yahoo search <https://search.yahoo.com/search?p=toggle+menu+%E4%B8%AD%E6%96%87>`_
       | `toggle menu 中文 - Baidu search <https://www.baidu.com/s?wd=toggle+menu+%E4%B8%AD%E6%96%87>`_
       | `toggle menu 中文 - Yandex search <https://www.yandex.com/search/?text=toggle+menu+%E4%B8%AD%E6%96%87>`_
.. adsu:: 4
.. [4] | `html special characters - Google search <https://www.google.com/search?q=html+special+characters>`_
       | `HTML Special Characters - Quackit Tutorials <http://www.quackit.com/html/html_special_characters.cfm>`_
       | `List of Unicode Characters - Quackit Tutorials <http://www.quackit.com/character_sets/unicode/>`_
       | `Unicode 9.0 Characters: Mathematical Operators - Quackit Tutorials <http://www.quackit.com/character_sets/unicode/versions/unicode_9.0.0/mathematical_operators_unicode_character_codes.cfm>`_
       | `Unicode Name: STRICTLY EQUIVALENT TO - Quackit Tutorials <http://www.quackit.com/html/html_editors/scratchpad/?app=charset_ref&hexadecimal=02263&decimal=8803&unicodeName=STRICTLY_EQUIVALENT_TO>`_

.. _CSS: https://www.google.com/search?q=CSS
.. _JavaScript: https://www.google.com/search?q=JavaScript

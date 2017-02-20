純CSS提示框(tooltip, 工具提示) - MDN網頁範例
############################################

:date: 2017-02-20T23:35+08:00
:tags: CSS, Web開發
:category: CSS
:summary: 純 CSS_ tooltip_ (提示框，工具提示)，截取自MDN的範例。
:og_image: http://www.javatpoint.com/images/javascript/javascript_logo.png
:adsu: yes


不囉唆，先看範例demo，把滑鼠移到藍色底線字上：

.. raw:: html

  <style>
    span[data-descr] {
      position: relative;
      text-decoration: underline;
      color: blue;
      cursor: help;
    }
    span[data-descr]:hover::after {
      content: attr(data-descr);
      position: absolute;
      left: 0;
      top: 1.25em;
      background-color: yellow;
      border: 1px gray solid;
      border-radius: 3px;
      padding: 3px;
    }
  </style>
  <div style="text-align: center; padding: 2em;">
  這是一個<span data-descr="大家好，我是提示">tooltip提示框</span>範例
  </div>

我把MDN網頁範例 [1]_ 改寫的更簡單一點。
程式原始碼如下：

**HTML程式碼** ：

.. code-block:: html

  這是一個<span data-descr="大家好，我是提示">tooltip提示框</span>範例

.. adsu:: 2

**CSS程式碼** ：

.. code-block:: css

  span[data-descr] {
      position: relative;
      text-decoration: underline;
      color: blue;
      cursor: help;
  }
  span[data-descr]:hover::after {
      content: attr(data-descr);
      position: absolute;
      left: 0;
      top: 1.25em;
      background-color: yellow;
      border: 1px gray solid;
      border-radius: 3px;
      padding: 3px;
  }

.. adsu:: 3

簡單解釋，就是利用 `::after`_ 偽元素(pseudo-element_) 結合CSS的 `attr()`_ 表達式
來製造一個tooltip提示框。
有興趣的話，順著我提供的連結(英文)往下看，不然就把程式碼拿去照著改一改就好！
希望對大家有幫助！

----

參考：

.. [1] `::after (:after) - CSS | MDN <https://developer.mozilla.org/en-US/docs/Web/CSS/::after#Tooltips>`_
.. [2] | `tooltip 中文 - Google search <https://www.google.com/search?q=tooltip+%E4%B8%AD%E6%96%87>`_
       | `tooltip 中文 - DuckDuckGo search <https://duckduckgo.com/?q=tooltip+%E4%B8%AD%E6%96%87>`_
       | `tooltip 中文 - Ecosia search <https://www.ecosia.org/search?q=tooltip+%E4%B8%AD%E6%96%87>`_
       | `tooltip 中文 - Bing search <https://www.bing.com/search?q=tooltip+%E4%B8%AD%E6%96%87>`_
       | `tooltip 中文 - Yahoo search <https://search.yahoo.com/search?p=tooltip+%E4%B8%AD%E6%96%87>`_
       | `tooltip 中文 - Baidu search <https://www.baidu.com/s?wd=tooltip+%E4%B8%AD%E6%96%87>`_
       | `tooltip 中文 - Yandex search <https://www.yandex.com/search/?text=tooltip+%E4%B8%AD%E6%96%87>`_
.. adsu:: 4
.. [3] | `pure css tooltip - Google search <https://www.google.com/search?q=pure+css+tooltip>`_
       | `pure css tooltip - DuckDuckGo search <https://duckduckgo.com/?q=pure+css+tooltip>`_
       | `pure css tooltip - Ecosia search <https://www.ecosia.org/search?q=pure+css+tooltip>`_
       | `pure css tooltip - Bing search <https://www.bing.com/search?q=pure+css+tooltip>`_
       | `pure css tooltip - Yahoo search <https://search.yahoo.com/search?p=pure+css+tooltip>`_
       | `pure css tooltip - Baidu search <https://www.baidu.com/s?wd=pure+css+tooltip>`_
       | `pure css tooltip - Yandex search <https://www.yandex.com/search/?text=pure+css+tooltip>`_

.. _CSS: https://www.google.com/search?q=CSS
.. _tooltip: https://www.google.com/search?q=tooltip
.. _pseudo-element: https://developer.mozilla.org/en-US/docs/Web/CSS/Pseudo-elements
.. _\:\:after: https://developer.mozilla.org/en-US/docs/Web/CSS/::after
.. _attr(): https://developer.mozilla.org/en-US/docs/Web/CSS/attr

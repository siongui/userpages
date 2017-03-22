純CSS將HTML的Table的行列互換
############################

:date: 2017-03-23T03:05+08:00
:tags: CSS, Web開發
:category: CSS
:summary: 純 CSS_ 解決方案：將HTML Table的行(row)以及列(column)對調
          ，不需要任何 JavaScript_ 程式碼。
:og_image: https://www.webcodegeeks.com/wp-content/uploads/2015/03/css-table11.jpg
:adsu: yes

純 CSS_ 解決方案：將HTML Table的行(row)以及列(column)對調，不需要任何
JavaScript_ 程式碼。

先看demo，原本的Table是：

.. raw:: html

  <style>
  th, td { border: 1px solid black; padding: 5px; }

  .table-transposed tr {
    display: block;
    float: left;
  }
  .table-transposed th, .table-transposed td {
    display: block;
  }
  </style>

  <div style="padding: 2em; background-color: Azure; padding: 5px;">

  <table>
  <tr><th>Type</th><th>For</th></tr>
  <tr><td>aa</td><td>ā</td></tr>
  <tr><td>ii</td><td>ī</td></tr>
  <tr><td>uu</td><td>ū</td></tr>
  <tr><td>"n</td><td>ṅ</td></tr>
  <tr><td>.m</td><td>ṃ</td></tr>
  <tr><td>~n</td><td>ñ</td></tr>
  <tr><td>.t</td><td>ṭ</td></tr>
  <tr><td>.d</td><td>ḍ</td></tr>
  <tr><td>.n</td><td>ṇ</td></tr>
  <tr><td>.l</td><td>ḷ</td></tr>
  </table>

  </div>

利用CSS行列互換後變成：

.. raw:: html

  <div style="padding: 2em; background-color: Azure; padding: 5px;">

  <table class="table-transposed">
  <tr><th>Type</th><th>For</th></tr>
  <tr><td>aa</td><td>ā</td></tr>
  <tr><td>ii</td><td>ī</td></tr>
  <tr><td>uu</td><td>ū</td></tr>
  <tr><td>"n</td><td>ṅ</td></tr>
  <tr><td>.m</td><td>ṃ</td></tr>
  <tr><td>~n</td><td>ñ</td></tr>
  <tr><td>.t</td><td>ṭ</td></tr>
  <tr><td>.d</td><td>ḍ</td></tr>
  <tr><td>.n</td><td>ṇ</td></tr>
  <tr><td>.l</td><td>ḷ</td></tr>
  </table>

  </div>

.. adsu:: 2

解法來自 [3]_ ， CSS_ 程式碼如下：

.. code-block:: css

  tr { display: block; float: left; }
  th, td { display: block; }

----

參考：

.. [1] | `html table swap rows columns - Google search <https://www.google.com/search?q=html+table+swap+rows+columns>`_
       | `html table swap rows columns - DuckDuckGo search <https://duckduckgo.com/?q=html+table+swap+rows+columns>`_
       | `html table swap rows columns - Ecosia search <https://www.ecosia.org/search?q=html+table+swap+rows+columns>`_
       | `html table swap rows columns - Qwant search <https://www.qwant.com/?q=html+table+swap+rows+columns>`_
       | `html table swap rows columns - Bing search <https://www.bing.com/search?q=html+table+swap+rows+columns>`_
       | `html table swap rows columns - Yahoo search <https://search.yahoo.com/search?p=html+table+swap+rows+columns>`_
       | `html table swap rows columns - Baidu search <https://www.baidu.com/s?wd=html+table+swap+rows+columns>`_
       | `html table swap rows columns - Yandex search <https://www.yandex.com/search/?text=html+table+swap+rows+columns>`_

.. [2] | `transpose html table css - Google search <https://www.google.com/search?q=transpose+html+table+css>`_
       | `transpose html table css - DuckDuckGo search <https://duckduckgo.com/?q=transpose+html+table+css>`_
       | `transpose html table css - Ecosia search <https://www.ecosia.org/search?q=transpose+html+table+css>`_
       | `transpose html table css - Qwant search <https://www.qwant.com/?q=transpose+html+table+css>`_
       | `transpose html table css - Bing search <https://www.bing.com/search?q=transpose+html+table+css>`_
       | `transpose html table css - Yahoo search <https://search.yahoo.com/search?p=transpose+html+table+css>`_
       | `transpose html table css - Baidu search <https://www.baidu.com/s?wd=transpose+html+table+css>`_
       | `transpose html table css - Yandex search <https://www.yandex.com/search/?text=transpose+html+table+css>`_
.. adsu:: 3
.. [3] `javascript - HTML Table with vertical rows - Stack Overflow <http://stackoverflow.com/a/16919439>`_

.. _CSS: https://www.google.com/search?q=CSS
.. _JavaScript: https://www.google.com/search?q=JavaScript

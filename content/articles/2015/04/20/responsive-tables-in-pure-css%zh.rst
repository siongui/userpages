純CSS響應式Table
################

:date: 2015-04-20 03:31
:tags: 響應式網頁設計(Responsive Web Design), Web開發, CSS
:category: 響應式網頁設計(Responsive Web Design)
:summary: 利用HTML data屬性來達成純CSS響應式Table
:adsu: yes

在 [1]_ 裡頭用一個有趣的手法來達成純CSS的響應式 Table_.
把 *th* 欄位裡的資料，利用 `HTML data屬性`_ 藏在每筆 *td* 旁，
若是大螢幕的裝置則正常顯示 *th* 裡的資料，若是小螢幕裝置則把
*th* 隱藏起來，改成顯示藏在 *td* 旁的資料。

參考範例碼（改寫自 [1]_ ）：

*HTML*

.. code-block:: html

  <table>
    <thead>
      <tr>
        <th>標頭1</th>
        <th>標頭2</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td data-header="標頭1">標頭1之資料1a</td>
        <td data-header="標頭2">標頭2之資料2a</td>
      </tr>
      <tr>
        <td data-header="標頭1">標頭1之資料1b</td>
        <td data-header="標頭2">標頭2之資料2b</td>
      </tr>
    </tbody>
  </table>

*CSS*

.. code-block:: css

  @media screen and (max-width: 767px)
    table td:before {
      content: attr(data-header);
      float: left;
      text-transform: uppercase;
      font-weight: bold;
    }
  }

----

參考：

.. [1] `Responsive Tables in Pure CSS - LivingSocial's Technology Blog <https://techblog.livingsocial.com/blog/2015/04/06/responsive-tables-in-pure-css/>`_
       (`Hacker News討論 <https://news.ycombinator.com/item?id=9328684>`__)

.. [2] `html - How to make responsive table - Stack Overflow <http://stackoverflow.com/questions/18436864/how-to-make-responsive-table>`_

.. [3] Google搜尋 `Responsive Tables in Pure CSS <https://www.google.com/search?q=Responsive+Tables+in+Pure+CSS>`__

.. [4] DuckDuckGo搜尋 `Responsive Tables in Pure CSS <https://duckduckgo.com/?q=Responsive+Tables+in+Pure+CSS>`__


.. _HTML data屬性: http://www.w3schools.com/tags/att_global_data.asp
.. _Table: http://www.w3schools.com/tags/tag_table.asp

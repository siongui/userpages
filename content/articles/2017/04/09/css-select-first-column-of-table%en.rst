[CSS] Select First Column of Table
##################################

:date: 2017-04-09T14:27+08:00
:tags: CSS, html, SCSS
:category: CSS
:summary: Select and style the first column (or n-th column) of table
          via CSS *:nth-of-type()* selector
:og_image: https://i.stack.imgur.com/pRYEO.png
:adsu: yes


See the demo below. The first column is selected and styled.

.. raw:: html

  <style>
  table.first-column > tbody > tr > td:nth-of-type(1) {
    font-size: larger;
    color: red;
  }
  </style>
  <div style="background-color: Azure; padding: 5px;">
  <table class="first-column">
    <tr>
      <td>(1,1)</td>
      <td>(1,2)</td>
      <td>(1,3)</td>
      <td>(1,4)</td>
    </tr>
    <tr>
      <td>(2,1)</td>
      <td>(2,2)</td>
      <td>(2,3)</td>
      <td>(2,4)</td>
    </tr>
    <tr>
      <td>(3,1)</td>
      <td>(3,2)</td>
      <td>(3,3)</td>
      <td>(3,4)</td>
    </tr>
  </table>
  </div>

**HTML**:

.. code-block:: html

  <table class="first-column">
    <tr>
      <td>(1,1)</td>
      <td>(1,2)</td>
      <td>(1,3)</td>
      <td>(1,4)</td>
    </tr>
    <tr>
      <td>(2,1)</td>
      <td>(2,2)</td>
      <td>(2,3)</td>
      <td>(2,4)</td>
    </tr>
    <tr>
      <td>(3,1)</td>
      <td>(3,2)</td>
      <td>(3,3)</td>
      <td>(3,4)</td>
    </tr>
  </table>

**CSS**:

.. code-block:: css

  table.first-column > tbody > tr > td:nth-of-type(1) {
    font-size: larger;
    color: red;
  }

We use `nth-of-type`_ CSS selector to select the first *td* of rach row. If you
want to select n-th column, change the number from 1 to n.

.. adsu:: 2

----

Tested on:
``Chromium Version 56.0.2924.76 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``

----

References
++++++++++

.. [1] | `css select first column of table - Google search <https://www.google.com/search?q=css+select+first+column+of+table>`_
       | `css select first column of table - DuckDuckGo search <https://duckduckgo.com/?q=css+select+first+column+of+table>`_
       | `css select first column of table - Ecosia search <https://www.ecosia.org/search?q=css+select+first+column+of+table>`_
       | `css select first column of table - Qwant search <https://www.qwant.com/?q=css+select+first+column+of+table>`_
       | `css select first column of table - Bing search <https://www.bing.com/search?q=css+select+first+column+of+table>`_
       | `css select first column of table - Yahoo search <https://search.yahoo.com/search?p=css+select+first+column+of+table>`_
       | `css select first column of table - Baidu search <https://www.baidu.com/s?wd=css+select+first+column+of+table>`_
       | `css select first column of table - Yandex search <https://www.yandex.com/search/?text=css+select+first+column+of+table>`_
.. [2] `CSS3 :nth-child() Selector - W3Schools <https://www.w3schools.com/cssref/sel_nth-child.asp>`_

.. _nth-of-type: https://www.google.com/search?q=nth-of-type
.. _nth-child: https://www.google.com/search?q=nth-child

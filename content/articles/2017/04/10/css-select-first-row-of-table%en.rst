[CSS] Select First Row of Table
###############################

:date: 2017-04-10T21:40+08:00
:tags: CSS, html, SCSS
:category: CSS
:summary: Select and style the first row (or n-th row) of table
          via CSS *:nth-of-type()* selector
:og_image: https://i.stack.imgur.com/pRYEO.png
:adsu: yes


See the demo below. The first row is selected and styled.

.. raw:: html

  <style>
  table.first-row > tbody > tr:nth-of-type(1) {
    font-size: larger;
    color: red;
  }
  </style>
  <div style="background-color: Azure; padding: 5px;">
  <table class="first-row">
   <tbody>
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
   </tbody>
  </table>
  </div>

**HTML**:

.. code-block:: html

  <table class="first-row">
   <tbody>
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
   </tbody>
  </table>

**CSS**:

.. code-block:: css

  table.first-row > tbody > tr:nth-of-type(1) {
    font-size: larger;
    color: red;
  }

We use `nth-of-type`_ CSS selector to select the first *tr* of the table. If you
want to select n-th row, change the number from 1 to n.

.. adsu:: 2

You may also be interested in
`[CSS] Select First Column of Table <{filename}../09/css-select-first-column-of-table%en.rst>`_

----

Tested on:
``Chromium Version 56.0.2924.76 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``

----

References
++++++++++

.. [1] | `css select first row of table - Google search <https://www.google.com/search?q=css+select+first+row+of+table>`_
       | `css select first row of table - DuckDuckGo search <https://duckduckgo.com/?q=css+select+first+row+of+table>`_
       | `css select first row of table - Ecosia search <https://www.ecosia.org/search?q=css+select+first+row+of+table>`_
       | `css select first row of table - Qwant search <https://www.qwant.com/?q=css+select+first+row+of+table>`_
       | `css select first row of table - Bing search <https://www.bing.com/search?q=css+select+first+row+of+table>`_
       | `css select first row of table - Yahoo search <https://search.yahoo.com/search?p=css+select+first+row+of+table>`_
       | `css select first row of table - Baidu search <https://www.baidu.com/s?wd=css+select+first+row+of+table>`_
       | `css select first row of table - Yandex search <https://www.yandex.com/search/?text=css+select+first+row+of+table>`_
.. [2] `CSS3 :nth-child() Selector - W3Schools <https://www.w3schools.com/cssref/sel_nth-child.asp>`_

.. _nth-of-type: https://www.google.com/search?q=nth-of-type
.. _nth-child: https://www.google.com/search?q=nth-child

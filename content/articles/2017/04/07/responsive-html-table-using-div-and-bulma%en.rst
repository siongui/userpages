Responsive HTML Table using Div and Bulma
#########################################

:date: 2017-04-07T23:26+08:00
:tags: CSS, Responsive Web Design, html, Bulma
:category: CSS
:summary: Make a responsive HTML table using *div* element and *Bulma* CSS
          framework.
:adsu: yes


Making a HTML table with normal *table* element is very difficult to make the
table responsive. So I use the `responsive columns`_ of Bulma_ framework to
make divs look like table.

Assume we have the following table:

`View on CodePen <http://codepen.io/anon/pen/ZedpyX>`__

.. code-block:: html

  <table>
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
  </table>

We can make a similar looking and responsive table by the following div
elements, and CSS classes of `responsive columns`_ provided by Bulma_:

.. code-block:: html

  <div class="div-table">
    <div class="columns is-gapless">
      <div class="column">(1,1)</div>
      <div class="column">(1,2)</div>
      <div class="column">(1,3)</div>
      <div class="column">(1,4)</div>
    </div>
    <div class="columns is-gapless">
      <div class="column">(2,1)</div>
      <div class="column">(2,2)</div>
      <div class="column">(2,3)</div>
      <div class="column">(2,4)</div>
    </div>
  </div>

`View on CodePen <http://codepen.io/anon/pen/PprGKw>`__

Bulma add ``margin-bottom`` between columns, we can use the following SCSS_ rule
to remove the ``margin-bottom``:

.. code-block:: scss

  .div-table {
    div.columns.is-gapless {
      margin: 0;
    }
  }

----

Tested on:
  - ``Chromium Version 56.0.2924.76 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``
  - ``Bulma 0.4.0``

----

References
++++++++++

.. [1] | `responsive table div - Google search <https://www.google.com/search?q=responsive+table+div>`_
       | `responsive table div - DuckDuckGo search <https://duckduckgo.com/?q=responsive+table+div>`_
       | `responsive table div - Ecosia search <https://www.ecosia.org/search?q=responsive+table+div>`_
       | `responsive table div - Qwant search <https://www.qwant.com/?q=responsive+table+div>`_
       | `responsive table div - Bing search <https://www.bing.com/search?q=responsive+table+div>`_
       | `responsive table div - Yahoo search <https://search.yahoo.com/search?p=responsive+table+div>`_
       | `responsive table div - Baidu search <https://www.baidu.com/s?wd=responsive+table+div>`_
       | `responsive table div - Yandex search <https://www.yandex.com/search/?text=responsive+table+div>`_

.. [2] `html - Why are people making tables with divs? - Software Engineering Stack Exchange <http://softwareengineering.stackexchange.com/questions/277778/why-are-people-making-tables-with-divs>`_

.. [3] | `responsive html table - Google search <https://www.google.com/search?q=responsive+html+table>`_
       | `responsive html table - DuckDuckGo search <https://duckduckgo.com/?q=responsive+html+table>`_
       | `responsive html table - Ecosia search <https://www.ecosia.org/search?q=responsive+html+table>`_
       | `responsive html table - Qwant search <https://www.qwant.com/?q=responsive+html+table>`_
       | `responsive html table - Bing search <https://www.bing.com/search?q=responsive+html+table>`_
       | `responsive html table - Yahoo search <https://search.yahoo.com/search?p=responsive+html+table>`_
       | `responsive html table - Baidu search <https://www.baidu.com/s?wd=responsive+html+table>`_
       | `responsive html table - Yandex search <https://www.yandex.com/search/?text=responsive+html+table>`_

.. _Bulma: http://bulma.io/
.. _SCSS: https://www.google.com/search?q=SCSS
.. _responsive columns: http://bulma.io/documentation/grid/columns/

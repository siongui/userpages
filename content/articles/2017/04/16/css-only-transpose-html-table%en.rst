CSS Invert (Transpose) Rows and Columns of HTML Table
#####################################################

:date: 2017-04-16T08:51+08:00
:tags: CSS, html, Pure CSS Toggle, toggle, toggleable
:category: CSS
:summary: Pure CSS solution for inverting (transposing) rows and columns of a
          HTML Table, and provide an option to toggle (original/inverted) the
          HTML table as well.
:og_image: https://www.webcodegeeks.com/wp-content/uploads/2015/03/css-table11.jpg
:adsu: yes


See demo first. Click the red **Transpose** text several times to see the table
transposed back and forth:

.. raw:: html

  <style>

  label[for=element-toggle] {
    cursor: pointer;
    color: red;
  }
  #element-toggle {
    display: none;
  }

  #element-toggle:checked ~ #toggled-element tr {
    display: block;
    float: left;
  }
  #element-toggle:checked ~ #toggled-element th,
  #element-toggle:checked ~ #toggled-element td {
    display: block;
  }

  </style>

  <div style="padding: 2em; background-color: Azure; padding: 5px;">

  <label for="element-toggle">Transpose</label>
  <input id="element-toggle" type="checkbox" />
  <br><br>

  <table id="toggled-element">
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


To invert (transpose) the rows and columns of an HTML table, we can use the
following CSS rules [2]_:

.. code-block:: css

  tr { display: block; float: left; }
  th, td { display: block; }

To toggle (original/transposed) the HTML table, we can use the technique
described in [3]_ and [4]_.

.. adsu:: 2

The complete source code for above demo:

**HTML**

.. code-block:: html

  <label for="element-toggle">Transpose</label>
  <input id="element-toggle" type="checkbox" />

  <table id="toggled-element">
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

.. adsu:: 3

**CSS**

.. code-block:: css

  label[for=element-toggle] {
    cursor: pointer;
    color: red;
  }
  #element-toggle {
    display: none;
  }

  #element-toggle:checked ~ #toggled-element tr {
    display: block;
    float: left;
  }
  #element-toggle:checked ~ #toggled-element th,
  #element-toggle:checked ~ #toggled-element td {
    display: block;
  }

----

References:

.. [1] | `css transpose table - Google search <https://www.google.com/search?q=css+transpose+table>`_
       | `css transpose table - DuckDuckGo search <https://duckduckgo.com/?q=css+transpose+table>`_
       | `css transpose table - Ecosia search <https://www.ecosia.org/search?q=css+transpose+table>`_
       | `css transpose table - Qwant search <https://www.qwant.com/?q=css+transpose+table>`_
       | `css transpose table - Bing search <https://www.bing.com/search?q=css+transpose+table>`_
       | `css transpose table - Yahoo search <https://search.yahoo.com/search?p=css+transpose+table>`_
       | `css transpose table - Baidu search <https://www.baidu.com/s?wd=css+transpose+table>`_
       | `css transpose table - Yandex search <https://www.yandex.com/search/?text=css+transpose+table>`_
.. [2] `javascript - HTML Table with vertical rows - Stack Overflow <http://stackoverflow.com/a/16919439>`_
.. adsu:: 4
.. [3] `Pure CSS Toggle Centered Element Width <{filename}../../03/29/css-only-toggle-centered-element-width%en.rst>`_
.. [4] `Pure CSS Toggle (Show/Hide) HTML Element <{filename}../../02/27/css-only-toggle-dom-element%en.rst>`_

.. _CSS: https://www.google.com/search?q=CSS
.. _JavaScript: https://www.google.com/search?q=JavaScript

Equal Height Components in Bulma Columns
########################################

:date: 2018-01-25T08:10+08:00
:tags: CSS, html, Responsive Web Design, Bulma
:category: CSS
:summary: Make Bulma components (message, card, etc.) to be of equal height in
          columns.
:og_image: https://cloud.githubusercontent.com/assets/18059922/25303180/ed993e0a-274d-11e7-887e-5425109adb65.png
:adsu: yes

.. raw:: html

  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.6.2/css/bulma.min.css">

Look at the following example. We have three message_ components in columns_ of
Bulma (version 0.6.2).

.. raw:: html

  <br>

  <div class="columns">
    <div class="column">
  
  <article class="message">
    <div class="message-header">
      <p>Header of Column 1</p>
    </div>
    <div class="message-body">
      Body of Column 1
    </div>
  </article>
  
    </div>
    <div class="column">
  
  <article class="message">
    <div class="message-header">
      <p>Header of Column 2</p>
    </div>
    <div class="message-body">
      Body of Column 3
    </div>
  </article>
  
    </div>
    <div class="column">
  
  <article class="message">
    <div class="message-header">
      <p>Header of Column 2</p>
    </div>
    <div class="message-body">
      Body of<br><br> Column 3
    </div>
  </article>
  
    </div>
  </div>
  <br>

The column 3 is higher than the other two columns on wide screens, and it looks
ugly. We want to make the three messages to be of equal height. (You should not
view this page on mobile devices because the columns will be stacked)

The solution [2]_ is to add ``height: 100%;`` to your components. In this case,
we add the following rule in CSS:

.. raw:: html

  <br>

.. code-block:: css

  .message {
    height: 100%;
  }

After adding ``height: 100%;``, The above example will look like:

.. raw:: html

  <style>
  .message.after {
    height: 100%;
  }
  .message-body {
    border: 0;
  }
  </style>
  <br>

  <div class="columns">
    <div class="column">
  
  <article class="message after">
    <div class="message-header">
      <p>Header of Column 1</p>
    </div>
    <div class="message-body">
      Body of Column 1
    </div>
  </article>
  
    </div>
    <div class="column">
  
  <article class="message after">
    <div class="message-header">
      <p>Header of Column 2</p>
    </div>
    <div class="message-body">
      Body of Column 3
    </div>
  </article>
  
    </div>
    <div class="column">
  
  <article class="message after">
    <div class="message-header">
      <p>Header of Column 2</p>
    </div>
    <div class="message-body">
      Body of<br><br> Column 3
    </div>
  </article>
  
    </div>
  </div>
  <br>

Now it looks better and should be the result of most designers want.

.. adsu:: 2

----

References:

.. [1] | `bulma columns equal height - Google search <https://www.google.com/search?q=bulma+columns+equal+height>`_
       | `bulma columns equal height - DuckDuckGo search <https://duckduckgo.com/?q=bulma+columns+equal+height>`_
       | `bulma columns equal height - Ecosia search <https://www.ecosia.org/search?q=bulma+columns+equal+height>`_
       | `bulma columns equal height - Qwant search <https://www.qwant.com/?q=bulma+columns+equal+height>`_
       | `bulma columns equal height - Bing search <https://www.bing.com/search?q=bulma+columns+equal+height>`_
       | `bulma columns equal height - Yahoo search <https://search.yahoo.com/search?p=bulma+columns+equal+height>`_
       | `bulma columns equal height - Baidu search <https://www.baidu.com/s?wd=bulma+columns+equal+height>`_
       | `bulma columns equal height - Yandex search <https://www.yandex.com/search/?text=bulma+columns+equal+height>`_

.. [2] `Equal height of card components · Issue #218 · jgthms/bulma · GitHub <https://github.com/jgthms/bulma/issues/218>`_
.. [3] `[Question] How to set columns height equal to the longest column? · Issue #696 · jgthms/bulma · GitHub <https://github.com/jgthms/bulma/issues/696>`_
.. [4] `css - How to set column height equal to longest column in Bulma (flexbox)? - Stack Overflow <https://stackoverflow.com/questions/43564132/how-to-set-column-height-equal-to-longest-column-in-bulma-flexbox>`_
.. [5] | `bulma is-full height - Google search <https://www.google.com/search?q=bulma+is-full+height>`_
       | `bulma is-full height - DuckDuckGo search <https://duckduckgo.com/?q=bulma+is-full+height>`_
       | `bulma is-full height - Ecosia search <https://www.ecosia.org/search?q=bulma+is-full+height>`_
       | `bulma is-full height - Qwant search <https://www.qwant.com/?q=bulma+is-full+height>`_
       | `bulma is-full height - Bing search <https://www.bing.com/search?q=bulma+is-full+height>`_
       | `bulma is-full height - Yahoo search <https://search.yahoo.com/search?p=bulma+is-full+height>`_
       | `bulma is-full height - Baidu search <https://www.baidu.com/s?wd=bulma+is-full+height>`_
       | `bulma is-full height - Yandex search <https://www.yandex.com/search/?text=bulma+is-full+height>`_

.. _Bulma: https://bulma.io/
.. _columns: https://bulma.io/documentation/columns/basics/
.. _message: https://bulma.io/documentation/components/message/

Pure CSS Bootstrap Accordion
############################

:date: 2018-09-29T20:39+08:00
:tags: Pure CSS Toggle, toggle, toggleable, html, CSS, Bootstrap,
       Accordion (Collapsible Content)
:category: CSS
:summary: CSS only toggle *Bootstrap* accordion (collapsible content).
          No *JavaScript* required.
:og_image: http://www.jquery-az.com/wp-content/uploads/2016/08/50.0_1-bootstrap-accordion.jpg
:adsu: yes

CSS only toggle Bootstrap_ accordion_. No *JavaScript* required.
Click the following texts of items to see demo.

.. raw:: html

  <style>
  input[name="accordion-select"], .panel-block {
    display: none;
  }

  /* the magic */
  #panel-1:checked ~ .accordion > .card > #collapseOne,
  #panel-2:checked ~ .accordion > .card > #collapseTwo,
  #panel-3:checked ~ .accordion > .card > #collapseThree {
    display: block;
  }
  </style>

  <!-- you can put the following line in your html head -->
  <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">

  <input type="radio" id="panel-1" name="accordion-select">
  <input type="radio" id="panel-2" name="accordion-select">
  <input type="radio" id="panel-3" name="accordion-select">

  <div class="accordion" id="accordionExample">
    <div class="card">
      <div class="card-header" id="headingOne">
        <h5 class="mb-0">
          <button class="btn btn-link" type="button" data-toggle="collapse" data-target="#collapseOne" aria-expanded="true" aria-controls="collapseOne">
            <label for="panel-1">Collapsible Group Item #1</label>
          </button>
        </h5>
      </div>

      <div id="collapseOne" class="collapse" aria-labelledby="headingOne" data-parent="#accordionExample">
        <div class="card-body">
          Anim pariatur cliche reprehenderit, enim eiusmod high life accusamus terry richardson ad squid. 3 wolf moon officia aute, non cupidatat skateboard dolor brunch. Food truck quinoa nesciunt laborum eiusmod. Brunch 3 wolf moon tempor, sunt aliqua put a bird on it squid single-origin coffee nulla assumenda shoreditch et. Nihil anim keffiyeh helvetica, craft beer labore wes anderson cred nesciunt sapiente ea proident. Ad vegan excepteur butcher vice lomo. Leggings occaecat craft beer farm-to-table, raw denim aesthetic synth nesciunt you probably haven't heard of them accusamus labore sustainable VHS.
        </div>
      </div>
    </div>
    <div class="card">
      <div class="card-header" id="headingTwo">
        <h5 class="mb-0">
          <button class="btn btn-link collapsed" type="button" data-toggle="collapse" data-target="#collapseTwo" aria-expanded="false" aria-controls="collapseTwo">
            <label for="panel-2">Collapsible Group Item #2</label>
          </button>
        </h5>
      </div>
      <div id="collapseTwo" class="collapse" aria-labelledby="headingTwo" data-parent="#accordionExample">
        <div class="card-body">
          Anim pariatur cliche reprehenderit, enim eiusmod high life accusamus terry richardson ad squid. 3 wolf moon officia aute, non cupidatat skateboard dolor brunch. Food truck quinoa nesciunt laborum eiusmod. Brunch 3 wolf moon tempor, sunt aliqua put a bird on it squid single-origin coffee nulla assumenda shoreditch et. Nihil anim keffiyeh helvetica, craft beer labore wes anderson cred nesciunt sapiente ea proident. Ad vegan excepteur butcher vice lomo. Leggings occaecat craft beer farm-to-table, raw denim aesthetic synth nesciunt you probably haven't heard of them accusamus labore sustainable VHS.
        </div>
      </div>
    </div>
    <div class="card">
      <div class="card-header" id="headingThree">
        <h5 class="mb-0">
          <button class="btn btn-link collapsed" type="button" data-toggle="collapse" data-target="#collapseThree" aria-expanded="false" aria-controls="collapseThree">
            <label for="panel-3">Collapsible Group Item #3</label>
          </button>
        </h5>
      </div>
      <div id="collapseThree" class="collapse" aria-labelledby="headingThree" data-parent="#accordionExample">
        <div class="card-body">
          Anim pariatur cliche reprehenderit, enim eiusmod high life accusamus terry richardson ad squid. 3 wolf moon officia aute, non cupidatat skateboard dolor brunch. Food truck quinoa nesciunt laborum eiusmod. Brunch 3 wolf moon tempor, sunt aliqua put a bird on it squid single-origin coffee nulla assumenda shoreditch et. Nihil anim keffiyeh helvetica, craft beer labore wes anderson cred nesciunt sapiente ea proident. Ad vegan excepteur butcher vice lomo. Leggings occaecat craft beer farm-to-table, raw denim aesthetic synth nesciunt you probably haven't heard of them accusamus labore sustainable VHS.
        </div>
      </div>
    </div>
  </div>

The basic technique is the same as *Pure CSS Tab Panel* [3]_. We need:

1 - Visible HTML label_ elements, which is the header of card_.

2 - Invisible HTML *input* *radio* elements, referenced by the *label* elements.
The value of *name* attribute of these elements must be the same.

3 - card body (content), invisible in the beginning.

When users click on the visible label elements, which are the headers of cards,
the corresponding input radio element is checked. We use CSS3 **:checked**
selector to select corresponding card body (content), and make the selected body
content visible.

.. adsu:: 2

The complete source code is as follows:

**HTML**:

.. code-block:: html

  <!-- you can put the following line in your html head -->
  <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">

  <input type="radio" id="panel-1" name="accordion-select">
  <input type="radio" id="panel-2" name="accordion-select">
  <input type="radio" id="panel-3" name="accordion-select">

  <div class="accordion" id="accordionExample">
    <div class="card">
      <div class="card-header" id="headingOne">
        <h5 class="mb-0">
          <button class="btn btn-link" type="button" data-toggle="collapse" data-target="#collapseOne" aria-expanded="true" aria-controls="collapseOne">
            <label for="panel-1">Collapsible Group Item #1</label>
          </button>
        </h5>
      </div>

      <div id="collapseOne" class="collapse" aria-labelledby="headingOne" data-parent="#accordionExample">
        <div class="card-body">
          Anim pariatur cliche reprehenderit, enim eiusmod high life accusamus terry richardson ad squid. 3 wolf moon officia aute, non cupidatat skateboard dolor brunch. Food truck quinoa nesciunt laborum eiusmod. Brunch 3 wolf moon tempor, sunt aliqua put a bird on it squid single-origin coffee nulla assumenda shoreditch et. Nihil anim keffiyeh helvetica, craft beer labore wes anderson cred nesciunt sapiente ea proident. Ad vegan excepteur butcher vice lomo. Leggings occaecat craft beer farm-to-table, raw denim aesthetic synth nesciunt you probably haven't heard of them accusamus labore sustainable VHS.
        </div>
      </div>
    </div>
    <div class="card">
      <div class="card-header" id="headingTwo">
        <h5 class="mb-0">
          <button class="btn btn-link collapsed" type="button" data-toggle="collapse" data-target="#collapseTwo" aria-expanded="false" aria-controls="collapseTwo">
            <label for="panel-2">Collapsible Group Item #2</label>
          </button>
        </h5>
      </div>
      <div id="collapseTwo" class="collapse" aria-labelledby="headingTwo" data-parent="#accordionExample">
        <div class="card-body">
          Anim pariatur cliche reprehenderit, enim eiusmod high life accusamus terry richardson ad squid. 3 wolf moon officia aute, non cupidatat skateboard dolor brunch. Food truck quinoa nesciunt laborum eiusmod. Brunch 3 wolf moon tempor, sunt aliqua put a bird on it squid single-origin coffee nulla assumenda shoreditch et. Nihil anim keffiyeh helvetica, craft beer labore wes anderson cred nesciunt sapiente ea proident. Ad vegan excepteur butcher vice lomo. Leggings occaecat craft beer farm-to-table, raw denim aesthetic synth nesciunt you probably haven't heard of them accusamus labore sustainable VHS.
        </div>
      </div>
    </div>
    <div class="card">
      <div class="card-header" id="headingThree">
        <h5 class="mb-0">
          <button class="btn btn-link collapsed" type="button" data-toggle="collapse" data-target="#collapseThree" aria-expanded="false" aria-controls="collapseThree">
            <label for="panel-3">Collapsible Group Item #3</label>
          </button>
        </h5>
      </div>
      <div id="collapseThree" class="collapse" aria-labelledby="headingThree" data-parent="#accordionExample">
        <div class="card-body">
          Anim pariatur cliche reprehenderit, enim eiusmod high life accusamus terry richardson ad squid. 3 wolf moon officia aute, non cupidatat skateboard dolor brunch. Food truck quinoa nesciunt laborum eiusmod. Brunch 3 wolf moon tempor, sunt aliqua put a bird on it squid single-origin coffee nulla assumenda shoreditch et. Nihil anim keffiyeh helvetica, craft beer labore wes anderson cred nesciunt sapiente ea proident. Ad vegan excepteur butcher vice lomo. Leggings occaecat craft beer farm-to-table, raw denim aesthetic synth nesciunt you probably haven't heard of them accusamus labore sustainable VHS.
        </div>
      </div>
    </div>
  </div>

.. adsu:: 3

**CSS**:

.. code-block:: css

  input[name="accordion-select"], .panel-block {
    display: none;
  }

  /* the magic */
  #panel-1:checked ~ .accordion > .card > #collapseOne,
  #panel-2:checked ~ .accordion > .card > #collapseTwo,
  #panel-3:checked ~ .accordion > .card > #collapseThree {
    display: block;
  }

The magic is in last rule. We use **:checked** and general sibling selector (~)
to make the user-selected card body content visible.

----

Tested on:

- ``Chromium 69.0.3497.81 on Ubuntu 18.04 (64-bit)``
- ``Bootstrap 4.1.3``

----

.. adsu:: 4

References:

.. [1] `Pure CSS Toggle (Show/Hide) HTML Element <{filename}/articles/2017/02/27/css-only-toggle-dom-element%en.rst>`_
.. [2] `Pure CSS Bootstrap Modal <{filename}/articles/2018/09/25/css-only-toggle-bootstrap-modal%en.rst>`_
.. [3] `Pure CSS Tab Panel <{filename}/articles/2017/05/21/css-only-tab-panel%en.rst>`_
.. [4] `Pure CSS Accordion (Collapsible Content) <{filename}/articles/2017/05/23/css-only-accordion-collapsible-content%en.rst>`_
.. [5] `[Vue.js] Bulma Accordion (Collapsible Content) <{filename}/articles/2018/01/28/vuejs-bulma-accordion-collapsible-content%en.rst>`_
.. [6] `Pure CSS Bulma Accordion (Collapsible Content) <{filename}/articles/2018/01/29/css-only-bulma-accordion-collapsible-content%en.rst>`_
.. [7] `Pure CSS Semantic UI Standard Accordion <{filename}/articles/2018/02/07/css-only-semantic-ui-standard-accordion%en.rst>`_

.. _label: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label
.. _for: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label#Using_the_for_attribute
.. _Bootstrap: https://getbootstrap.com/
.. _accordion: https://getbootstrap.com/docs/4.1/components/collapse/#accordion-example
.. _card: https://getbootstrap.com/docs/4.1/components/card/

Pure CSS Semantic UI Standard Accordion
#######################################

:date: 2018-02-07T07:44+08:00
:tags: Pure CSS Toggle, toggle, toggleable, html, CSS, Semantic UI,
       Accordion (Collapsible Content)
:category: CSS
:summary: CSS only *Semantic UI* standard accordion. No *JavaScript* required.
:og_image: https://ckeditor.com/cke4/sites/default/files/styles/large/public/suiaccordion/q4_1.JPG
:adsu: yes

CSS only toggle `Semantic UI`_ standard accordion_. No *JavaScript* required.
The following demo is not totally the same as the standard accordion in official
website, but it's fun to do experiment without JavaScript.

.. raw:: html

  <style>
  input[name="accordion-select"] {
    display: none;
  }
  label.title {
    cursor: pointer;
    display: block;
  }

  /* the magic */
  #content1:checked ~ .accordion > .body-1,
  #content2:checked ~ .accordion > .body-2,
  #content3:checked ~ .accordion > .body-3 {
    display: block;
  }
  .ui.accordion > .content {
    border: 0;
  }
  </style>

  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/semantic-ui@2.2.14/dist/semantic.min.css">
  <div style="margin: 2rem;">

  <input type="radio" id="content1" name="accordion-select">
  <input type="radio" id="content2" name="accordion-select">
  <input type="radio" id="content3" name="accordion-select">
  <div class="ui accordion">
    <label class="title" for="content1">
      <i class="dropdown icon"></i>
      What is a dog?
    </label>
    <div class="content body-1">
      <p>A dog is a type of domesticated animal. Known for its loyalty and faithfulness, it can be found as a welcome guest in many households across the world.</p>
    </div>
    <label class="title" for="content2">
      <i class="dropdown icon"></i>
      What kinds of dogs are there?
    </label>
    <div class="content body-2">
      <p>There are many breeds of dogs. Each breed varies in size and temperament. Owners often select a breed of dog that they find to be compatible with their own lifestyle and desires from a companion.</p>
    </div>
    <label class="title" for="content3">
      <i class="dropdown icon"></i>
      How do you acquire a dog?
    </label>
    <div class="content body-3">
      <p>Three common ways for a prospective owner to acquire a dog is from pet shops, private owners, or shelters.</p>
      <p>A pet shop may be the most convenient way to buy a dog. Buying a dog from a private owner allows you to assess the pedigree and upbringing of your dog before choosing to take it home. Lastly, finding your dog from a shelter, helps give a good home to a dog who may not find one so readily.</p>
    </div>
  </div>

  </div>


The above demo apply the technique of Bulma accordion [1]_ to Semantic UI
standard accordion.

The following is the source code for above demo.

**HTML**:

.. code-block:: html

  <input type="radio" id="content1" name="accordion-select">
  <input type="radio" id="content2" name="accordion-select">
  <input type="radio" id="content3" name="accordion-select">
  <div class="ui accordion">
    <label class="title" for="content1">
      <i class="dropdown icon"></i>
      What is a dog?
    </label>
    <div class="content body-1">
      <p>A dog is a type of domesticated animal. Known for its loyalty and faithfulness, it can be found as a welcome guest in many households across the world.</p>
    </div>
    <label class="title" for="content2">
      <i class="dropdown icon"></i>
      What kinds of dogs are there?
    </label>
    <div class="content body-2">
      <p>There are many breeds of dogs. Each breed varies in size and temperament. Owners often select a breed of dog that they find to be compatible with their own lifestyle and desires from a companion.</p>
    </div>
    <label class="title" for="content3">
      <i class="dropdown icon"></i>
      How do you acquire a dog?
    </label>
    <div class="content body-3">
      <p>Three common ways for a prospective owner to acquire a dog is from pet shops, private owners, or shelters.</p>
      <p>A pet shop may be the most convenient way to buy a dog. Buying a dog from a private owner allows you to assess the pedigree and upbringing of your dog before choosing to take it home. Lastly, finding your dog from a shelter, helps give a good home to a dog who may not find one so readily.</p>
    </div>
  </div>

.. adsu:: 2

**CSS**:

.. code-block:: css

  input[name="accordion-select"] {
    display: none;
  }
  label.title {
    cursor: pointer;
    display: block;
  }

  /* the magic */
  #content1:checked ~ .accordion > .body-1,
  #content2:checked ~ .accordion > .body-2,
  #content3:checked ~ .accordion > .body-3 {
    display: block;
  }

The magic is in last rule. We use **:checked** and general sibling selector (~)
to make the user-selected content visible.

----

Tested on:

- ``Chromium 64.0.3282.119 on Ubuntu 17.10 (64-bit)``
- ``Semantic UI 2.2.14``

----

.. adsu:: 3

References:

.. [1] `Pure CSS Bulma Accordion (Collapsible Content) <{filename}/articles/2018/01/29/css-only-bulma-accordion-collapsible-content%en.rst>`_

.. _label: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label
.. _input checkbox: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/checkbox
.. _for: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label#Using_the_for_attribute
.. _Semantic UI: https://semantic-ui.com/
.. _accordion: https://semantic-ui.com/modules/accordion.html

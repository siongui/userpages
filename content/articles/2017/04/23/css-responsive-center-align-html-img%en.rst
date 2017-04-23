[CSS] Responsive Center Align Image
###################################

:date: 2017-04-23T21:22+08:00
:tags: CSS, Responsive Web Design, html
:category: CSS
:summary: CSS rule for HTML *img* element to be both responsive and centered.
:og_image: https://www.campaignmonitor.com/assets/uploads/2014/07/centering_before.png
:adsu: yes

CSS_ rule for HTML img_ element to be both responsive and centered horizontally.

**CSS**:

.. code-block:: css

  img.responsive-image-align-center {
      display: block;
      margin-left: auto;
      margin-right: auto;
      max-width: 100%;
      height: auto;
  }

**Demo**:

.. raw:: html

  <style>
  img.responsive-image-align-center {
      display: block;
      margin-left: auto;
      margin-right: auto;
      max-width: 100%;
      height: auto;
  }
  </style>
  <img class="responsive-image-align-center" src="http://lh6.googleusercontent.com/-_30QEfBiX_c/V4jDdy6klfI/AAAAAAAACjY/sk0jUePFaGQdcmCqJ-A6reZOQvDfLxVvwCJkC/s640/photo.jpg">

.. adsu:: 2

| To know how to make image responsive, see [1]_.
| To know how to make image centered, see [2]_.

Tested on:
``Chromium Version 57.0.2987.98 Built on Ubuntu , running on Ubuntu 17.04 (64-bit)``

----

References
++++++++++

.. [1] | `image responsive - Google search <https://www.google.com/search?q=image+responsive>`_
       | `image responsive - DuckDuckGo search <https://duckduckgo.com/?q=image+responsive>`_
       | `image responsive - Ecosia search <https://www.ecosia.org/search?q=image+responsive>`_
       | `image responsive - Qwant search <https://www.qwant.com/?q=image+responsive>`_
       | `image responsive - Bing search <https://www.bing.com/search?q=image+responsive>`_
       | `image responsive - Yahoo search <https://search.yahoo.com/search?p=image+responsive>`_
       | `image responsive - Baidu search <https://www.baidu.com/s?wd=image+responsive>`_
       | `image responsive - Yandex search <https://www.yandex.com/search/?text=image+responsive>`_
       | `Responsive Web Design Images <https://www.w3schools.com/css/css_rwd_images.asp>`_

.. [2] | `img align center - Google search <https://www.google.com/search?q=img+align+center>`_
       | `img align center - DuckDuckGo search <https://duckduckgo.com/?q=img+align+center>`_
       | `img align center - Ecosia search <https://www.ecosia.org/search?q=img+align+center>`_
       | `img align center - Qwant search <https://www.qwant.com/?q=img+align+center>`_
       | `img align center - Bing search <https://www.bing.com/search?q=img+align+center>`_
       | `img align center - Yahoo search <https://search.yahoo.com/search?p=img+align+center>`_
       | `img align center - Baidu search <https://www.baidu.com/s?wd=img+align+center>`_
       | `img align center - Yandex search <https://www.yandex.com/search/?text=img+align+center>`_
       | `html - Center image using text-align center? - Stack Overflow <http://stackoverflow.com/questions/7055393/center-image-using-text-align-center>`_
       | `CSS Layout - Horizontal & Vertical Align <https://www.w3schools.com/css/css_align.asp>`_

.. [3] | `html img default display - Google search <https://www.google.com/search?q=html+img+default+display>`_
       | `html img default display - DuckDuckGo search <https://duckduckgo.com/?q=html+img+default+display>`_
       | `html img default display - Ecosia search <https://www.ecosia.org/search?q=html+img+default+display>`_
       | `html img default display - Qwant search <https://www.qwant.com/?q=html+img+default+display>`_
       | `html img default display - Bing search <https://www.bing.com/search?q=html+img+default+display>`_
       | `html img default display - Yahoo search <https://search.yahoo.com/search?p=html+img+default+display>`_
       | `html img default display - Baidu search <https://www.baidu.com/s?wd=html+img+default+display>`_
       | `html img default display - Yandex search <https://www.yandex.com/search/?text=html+img+default+display>`_
       | `html - Is <img> element block level or inline level? - Stack Overflow <http://stackoverflow.com/questions/2402761/is-img-element-block-level-or-inline-level>`_

.. _CSS: https://www.google.com/search?q=CSS
.. _img: https://www.w3schools.com/tags/tag_img.asp

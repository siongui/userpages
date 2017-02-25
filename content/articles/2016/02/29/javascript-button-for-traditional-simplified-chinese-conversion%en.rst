[JavaScript] Button For Traditional/Simplified Chinese Conversion on Website
############################################################################

:date: 2016-02-29T01:39+08:00
:tags: JavaScript, SCSS, CSS, html, String Manipulation,
       Conversion of Traditional and Simplified Chinese
:category: JavaScript
:summary: Conversion Button of Traditional/Simplified Chinese on Website via
          JavaScript_.
:og_image: http://www.javatpoint.com/images/javascript/javascript_logo.png
:adsu: yes


Conversion Button of Traditional/Simplified Chinese on Website via JavaScript_.

First we implement button-like element by HTML and SCSS_:
(Assume default language of website is Traditional Chinese)

*HTML*

.. code-block:: html

  <div id="zh-convert" class="zh-convert" data-zh="tw">简体</div>

*SCSS*

.. code-block:: scss

  .zh-convert {
    &:hover {cursor: pointer;}
    float: right;
    margin-right: 3px;

    padding: 5px 5px;
    background: #fff;
    border-radius: 4px;
  }

.. adsu:: 2

We need `New Tong Wen Tang <http://tongwen.openfoundry.org/>`_
(`GitHub <https://github.com/softcup/New-Tongwentang-for-Web>`__) library to do
the conversion for us. Insert the following line in your HTML right before the
end of *body* tag.

.. code-block:: html

  <script language="JavaScript" src="http://tongwen.openfoundry.org/src/web/tongwen_core.js"></script>
  <script language="JavaScript" src="http://tongwen.openfoundry.org/src/web/tongwen_table_s2t.js"></script>
  <script language="JavaScript" src="http://tongwen.openfoundry.org/src/web/tongwen_table_t2s.js"></script>
  <script language="JavaScript" src="http://tongwen.openfoundry.org/src/web/tongwen_table_ps2t.js"></script>
  <script language="JavaScript" src="http://tongwen.openfoundry.org/src/web/tongwen_table_pt2s.js"></script>

When users click on ``简体``, the whole website will be converted to Simplified
Chinese.
When users click on ``繁體``, the whole website will be converted to Traditional
Chinese.

*JavaScript*

.. code-block:: javascript

  var zhconvert = document.getElementById("zh-convert");
  zhconvert.onclick = function() {
    if (zhconvert.dataset.zh == "tw") {
      TongWen.trans2Simp(document);
      zhconvert.dataset.zh = "cn";
      setTimeout(function() {zhconvert.textContent = "繁體";}, 500);
    } else {
      TongWen.trans2Trad(document);
      zhconvert.dataset.zh = "tw";
      setTimeout(function() {zhconvert.textContent = "简体";}, 500);
    }
  }

.. adsu:: 3

----

References:

.. [1] `zh_TW to/from zh_CN · twnanda/twnanda@6897761 · GitHub <https://github.com/twnanda/twnanda/commit/689776194597d62cf531d9556f97a958afb0496c>`_


.. _JavaScript: https://www.google.com/search?q=javascript
.. _SCSS: https://www.google.com/search?q=scss

[JavaScript] Font Size Larger/Smaller
#####################################

:date: 2016-02-27T19:58+08:00
:tags: JavaScript, SCSS, CSS, html
:category: JavaScript
:summary: Make font size of your website larger/smaller via JavaScript_.
:adsu: yes


Make font size of your website larger/smaller via JavaScript_.

First we implement button-like element by HTML and SCSS_:

*HTML*

.. code-block:: html

  <div class="font-adjust">
    <span id="font-larger">A+</span>
    <span id="font-smaller">A-</span>
  </div>

*SCSS*

.. code-block:: scss

  .font-adjust {
    float: right;
    margin-right: 3px;

    span {
      &:hover {cursor: pointer;}
      padding: 5px 5px;
      background: #fff;
      border-radius: 4px;
    }
  }

When users click on ``A+``/``A-``, make larger/smaller font size of div element
whose class is *.main-content*:

*JavaScript*

.. code-block:: javascript

  var fl = document.getElementById("font-larger");
  fl.onclick = function() {
    var contents = document.querySelectorAll(".main-content");
    for (var i = 0; i < contents.length; i++) {
      var content = contents[i];
      if (content.style.fontSize == "") {
        content.style.fontSize = "1.0em";
      }
      content.style.fontSize = (parseFloat(content.style.fontSize) + 0.25) + "em";
    }
  }

  var fs = document.getElementById("font-smaller");
  fs.onclick = function() {
    var contents = document.querySelectorAll(".main-content");
    for (var i = 0; i < contents.length; i++) {
      var content = contents[i];
      if (content.style.fontSize == "") {
        content.style.fontSize = "1.0em";
      }
      content.style.fontSize = (parseFloat(content.style.fontSize) - 0.25) + "em";
    }
  }


----

References:

.. [1] `font size larger/smaller · twnanda/twnanda@7233501 · GitHub <https://github.com/twnanda/twnanda/commit/723350195ba39c2ecfa2f303e7a4cac2af5bec37>`_

.. [2] `javascript font size increase/decrease <https://www.google.com/search?q=javascript+font+size+increase%2Fdecrease>`_

.. [3] `Change Text Size On Click With JavaScript <https://davidwalsh.name/change-text-size-onclick-with-javascript>`_


.. _JavaScript: https://www.google.com/search?q=javascript
.. _SCSS: https://www.google.com/search?q=scss

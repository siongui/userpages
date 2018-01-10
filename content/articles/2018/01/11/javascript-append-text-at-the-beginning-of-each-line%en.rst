[JavaScript] Append Text at the Beginning of Each Line
######################################################

:date: 2018-01-11T04:50+08:00
:tags: JavaScript, Copy to Clipboard, Online Toolkit, String Manipulation,
       Regular Expression
:category: JavaScript
:summary: Append text at the beginning of each line via JavaScript.
:adsu: yes


.. raw:: html

  <style>
  textarea {
    box-sizing: border-box;
    width: 100%;
  }
  </style>

.. raw:: html

   <input type="text" id="text" tabindex="1" placeholder="text to be prepended">
   <button type="button" id="prepend">Prepend</button><br><br>
   Paste your lines below:<br>
   <textarea id="origin" rows="10"></textarea><br><br>
   Lines after prepended:<br>
   <textarea id="after" rows="10"></textarea><br>
   <button type="button" id="copy" tabindex="2">Copy textarea to clipboard</button>


.. raw:: html

  <script>
    var inputElm = document.getElementById("text");
    var buttonElm = document.getElementById("prepend");
    var oriTextareaElm = document.getElementById("origin");
    var afterTextareaElm = document.getElementById("after");
    var copyElm = document.getElementById("copy");
    inputElm.onkeyup = function(event) {
      // press enter
      if (event.keyCode == 13) {
        prependToEachLine();
      }
    }
    buttonElm.onclick = function(event) { prependToEachLine(); }
    copyElm.onclick = function(event) {
      afterTextareaElm.select();
      var isSuccessful = document.execCommand('copy');
      if (isSuccessful) {
        afterTextareaElm.value = "Copy OK";
      } else {
        afterTextareaElm.value = "Copy Fail";
      }
    }

    function prependToEachLine() {
      afterTextareaElm.value = oriTextareaElm.value.replace(/^/gm, inputElm.value);
    }

  </script>

.. adsu:: 2

The first *textarea* is the lines to be prepended. The second *textarea* is the
lines after prepended. The JavaScript code to append text at the beginning of
each line is:

.. code-block:: javascript

  afterTextareaElm.value = oriTextareaElm.value.replace(/^/gm, inputElm.value);

- ``^`` means beginning of line.
- ``g`` means global match
- ``m`` means multiline matching

If you are not familiar with JavaScript regular expression, see [3]_.

----

References:

.. [1] `Search Links of Major Search Engines <{filename}../../../2016/04/03/search-links-of-major-search-engines%en.rst>`_
.. [2] `[sed] Append Text at the Beginning of Each Line <{filename}../../../2016/03/15/sed-append-text-at-the-beginning-of-each-line%en.rst>`_
.. [3] | `Regular Expressions - JavaScript | MDN <https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Regular_Expressions>`_
       | `JavaScript RegExp Object <https://www.w3schools.com/js/js_regexp.asp>`_
.. [4] `css - HTML Textarea isn't responsive - Stack Overflow <https://stackoverflow.com/questions/39068128/html-textarea-isnt-responsive>`_

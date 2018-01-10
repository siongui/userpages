Search Taobao.com
#################

:date: 2016-05-06T22:36+08:00
:tags: JavaScript, Copy to Clipboard, Online Toolkit
:category: JavaScript
:summary: Search Taobao.com_.
:adsu: yes


.. raw:: html

   Search <a href="https://www.taobao.com/" target="_blank">Taobao.com</a>:
   <input type="text" id="search" tabindex="1" placeholder="Input Search String">
   <button type="button" id="getlinks">Get Links</button><br>
   <input type="checkbox" id="pricel2h"> Price: Low to High<br>
   <hr>
   <div id="searchlinks"></div>
   <hr>
   <textarea id="links" rows="8" cols="80"></textarea><br>
   <button type="button" id="copy" tabindex="2">Copy textarea to clipboard</button>


.. raw:: html

  <script>
    var inputElm = document.getElementById("search");
    var buttonElm = document.getElementById("getlinks");
    var divElm = document.getElementById("searchlinks");
    var textareaElm = document.getElementById("links");
    var copyElm = document.getElementById("copy");
    var pricel2hElm = document.getElementById("pricel2h");
    inputElm.onkeyup = function(event) {
      if (event.keyCode == 13) {
        processSearchInput();
      }
    }
    buttonElm.onclick = function(event) { processSearchInput(); }
    copyElm.onclick = function(event) {
      textareaElm.select();
      var isSuccessful = document.execCommand('copy');
      if (isSuccessful) {
        textareaElm.value = "Copy OK";
      } else {
        textareaElm.value = "Copy Fail";
      }
    }

    function processSearchInput() {
      var search_string = inputElm.value.trim().replace(/\s+/g, " ");
      var qstring = encodeURI(inputElm.value.trim().replace(/\s+/g, "+"));

      textareaElm.value = ".. [1] `TERM_Taobao search <TAOBAO>`_\n".replace("TAOBAO", TaobaoURL(qstring)).replace(/TERM/g, search_string);

      divElm.innerHTML = '<a target="_blank" href="TAOBAO">TERM_Taobao search</a><br>'.replace("TAOBAO", TaobaoURL(qstring)).replace(/TERM/g, search_string);
    }

    function TaobaoURL(qstring) {
      var turl = "https://s.taobao.com/search?q=" + qstring;
      if (pricel2hElm.checked) {
        turl += "&sort=price-asc";
      }
      return turl;
    }
  </script>

----

.. adsu:: 2

References:

.. [1] `Search Links of Major Search Engines <{filename}../../04/03/search-links-of-major-search-engines%en.rst>`_

.. [2] `Search PTT BBS <{filename}../../04/07/search-ptt-bbs%en.rst>`_


.. _Taobao.com: https://www.taobao.com/

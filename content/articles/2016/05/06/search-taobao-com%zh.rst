搜尋 淘寶網
###########

:date: 2016-05-06T22:36+08:00
:tags: Web開發, 字串處理
:category: 網頁小工具
:summary: 搜尋 `淘寶網`_.
:adsu: yes


.. raw:: html

   搜尋 <a href="https://www.taobao.com/" target="_blank">淘寶網</a>:
   <input type="text" id="search" tabindex="1" placeholder="輸入搜尋字串">
   <button type="button" id="getlinks">獲取連結</button><br>
   <input type="checkbox" id="pricel2h"> 價格從低到高<br>
   <hr>
   <div id="searchlinks"></div>
   <hr>
   <textarea id="links" rows="8" cols="80"></textarea><br>
   <button type="button" id="copy" tabindex="2">複製到剪貼簿</button>


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

      textareaElm.value = ".. [1] `TERM_淘寶搜尋 <TAOBAO>`_\n".replace("TAOBAO", TaobaoURL(qstring)).replace(/TERM/g, search_string);

      divElm.innerHTML = '<a target="_blank" href="TAOBAO">TERM_淘寶搜尋</a><br>'.replace("TAOBAO", TaobaoURL(qstring)).replace(/TERM/g, search_string);
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

.. [2] `PTT搜尋 <{filename}../../04/07/search-ptt-bbs%zh.rst>`_


.. _淘寶網: https://www.taobao.com/

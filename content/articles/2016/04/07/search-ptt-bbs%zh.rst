PTT搜尋
#######

:date: 2016-04-07T06:42+08:00
:tags: Web開發, 字串處理
:category: 網頁小工具
:summary: 搜尋 PTT_ BBS_.
:adsu: yes


.. raw:: html

   搜尋 <a href="https://www.ptt.cc/" target="_blank">PTT</a> ：
   <input type="text" id="search" tabindex="1" placeholder="搜尋字串">
   <button type="button" id="getlinks">獲取連結</button><br>
   <hr>
   <div id="searchlinks"></div>
   <hr>
   <textarea id="links" rows="8" cols="80"></textarea><br>
   <button type="button" id="copy" tabindex="2">複製到剪貼簿</button>

.. adsu:: 2

.. raw:: html

  <script>
    var inputElm = document.getElementById("search");
    var buttonElm = document.getElementById("getlinks");
    var divElm = document.getElementById("searchlinks");
    var textareaElm = document.getElementById("links");
    var copyElm = document.getElementById("copy");
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

      textareaElm.value = ".. [1] `TERM - Google 搜尋 <GOOG>`_\n\n       `TERM - DuckDuckGo 搜尋 <DUCK>`_\n\n       `TERM - Bing 搜尋 <BING>`_\n\n       `TERM - Yahoo 搜尋 <YAHOO>`_\n\n       `TERM - Baidu 搜尋 <BAIDU>`_\n\n       `TERM - Yandex 搜尋 <YANDEX>`_\n".replace("GOOG", GoogleURL(qstring)).replace("DUCK", DuckDuckGoURL(qstring)).replace("BING", BingURL(qstring)).replace("YAHOO", YahooURL(qstring)).replace("BAIDU", BaiduURL(qstring)).replace("YANDEX", YandexURL(qstring)).replace(/TERM/g, search_string);

      divElm.innerHTML = '<a target="_blank" href="GOOG">TERM - Google 搜尋</a><br><a target="_blank" href="DUCK">TERM - DuckDuckGo 搜尋</a><br><a target="_blank" href="BING">TERM - Bing 搜尋</a><br><a target="_blank" href="YAHOO">TERM - Yahoo 搜尋</a><br><a target="_blank" href="BAIDU">TERM - Baidu 搜尋</a><br><a target="_blank" href="YANDEX">TERM - Yandex 搜尋</a><br>'.replace("GOOG", GoogleURL(qstring)).replace("DUCK", DuckDuckGoURL(qstring)).replace("BING", BingURL(qstring)).replace("YAHOO", YahooURL(qstring)).replace("BAIDU", BaiduURL(qstring)).replace("YANDEX", YandexURL(qstring)).replace(/TERM/g, search_string);
    }

    function GoogleURL(qstring) {
      return "https://www.google.com/search?q=" + qstring + "+site%3Aptt.cc";
    }
    function BingURL(qstring) {
      return "https://www.bing.com/search?q=" + qstring + "+site%3Aptt.cc";
    }
    function YahooURL(qstring) {
      return "https://search.yahoo.com/search?p=" + qstring + "+site%3Aptt.cc";
    }
    function DuckDuckGoURL(qstring) {
      return "https://duckduckgo.com/?q=" + qstring + "+site%3Aptt.cc";
    }
    function BaiduURL(qstring) {
      return "https://www.baidu.com/s?wd=" + qstring + "+site%3Aptt.cc";
    }
    function YandexURL(qstring) {
      return "https://www.yandex.com/search/?text=" + qstring + "+site%3Aptt.cc";
    }
  </script>

----

References:

.. [1] `javascript open link in new tab - Google search <https://www.google.com/search?q=javascript+open+link+in+new+tab>`_

       `javascript open link in new tab - DuckDuckGo search <https://duckduckgo.com/?q=javascript+open+link+in+new+tab>`_

       `javascript open link in new tab - Bing search <https://www.bing.com/search?q=javascript+open+link+in+new+tab>`_

       `javascript open link in new tab - Yahoo search <https://search.yahoo.com/search?p=javascript+open+link+in+new+tab>`_

       `javascript open link in new tab - Baidu search <https://www.baidu.com/s?wd=javascript+open+link+in+new+tab>`_

       `javascript open link in new tab - Yandex search <https://www.yandex.com/search/?text=javascript+open+link+in+new+tab>`_

       `Open a URL in a new tab (and not a new window) using JavaScript - Stack Overflow <http://stackoverflow.com/questions/4907843/open-a-url-in-a-new-tab-and-not-a-new-window-using-javascript>`_

.. [2] `Search Links of Major Search Engines <{filename}../03/search-links-of-major-search-engines%en.rst>`_

.. [3] `搜尋 淘寶網 <{filename}../../05/06/search-taobao-com%zh.rst>`_


.. _PTT: https://www.ptt.cc/
.. _BBS: https://en.wikipedia.org/wiki/Bulletin_board_system

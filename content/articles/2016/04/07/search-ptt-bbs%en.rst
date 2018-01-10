Search PTT BBS
##############

:date: 2016-04-07T06:33+08:00
:tags: JavaScript, Copy to Clipboard, Online Toolkit
:category: JavaScript
:summary: Search PTT_ BBS_.
:adsu: yes


.. raw:: html

   Search <a href="https://www.ptt.cc/" target="_blank">PTT</a>:
   <input type="text" id="search" tabindex="1" placeholder="Input Search String">
   <button type="button" id="getlinks">Get Links</button><br>
   <hr>
   <div id="searchlinks"></div>
   <hr>
   <textarea id="links" rows="8" cols="80"></textarea><br>
   <button type="button" id="copy" tabindex="2">Copy textarea to clipboard</button>

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

      textareaElm.value = ".. [1] `TERM - Google search <GOOG>`_\n\n       `TERM - DuckDuckGo search <DUCK>`_\n\n       `TERM - Bing search <BING>`_\n\n       `TERM - Yahoo search <YAHOO>`_\n\n       `TERM - Baidu search <BAIDU>`_\n\n       `TERM - Yandex search <YANDEX>`_\n".replace("GOOG", GoogleURL(qstring)).replace("DUCK", DuckDuckGoURL(qstring)).replace("BING", BingURL(qstring)).replace("YAHOO", YahooURL(qstring)).replace("BAIDU", BaiduURL(qstring)).replace("YANDEX", YandexURL(qstring)).replace(/TERM/g, search_string);

      divElm.innerHTML = '<a target="_blank" href="GOOG">TERM - Google Search</a><br><a target="_blank" href="DUCK">TERM - DuckDuckGo Search</a><br><a target="_blank" href="BING">TERM - Bing Search</a><br><a target="_blank" href="YAHOO">TERM - Yahoo Search</a><br><a target="_blank" href="BAIDU">TERM - Baidu Search</a><br><a target="_blank" href="YANDEX">TERM - Yandex Search</a><br>'.replace("GOOG", GoogleURL(qstring)).replace("DUCK", DuckDuckGoURL(qstring)).replace("BING", BingURL(qstring)).replace("YAHOO", YahooURL(qstring)).replace("BAIDU", BaiduURL(qstring)).replace("YANDEX", YandexURL(qstring)).replace(/TERM/g, search_string);
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

.. [3] `Search Taobao.com <{filename}../../05/06/search-taobao-com%en.rst>`_


.. _PTT: https://www.ptt.cc/
.. _BBS: https://en.wikipedia.org/wiki/Bulletin_board_system

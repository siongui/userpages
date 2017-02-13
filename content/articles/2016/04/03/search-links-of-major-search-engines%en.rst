Search Links of Major Search Engines
####################################

:date: 2016-04-03T21:51+08:00
:modified: 2017-02-07T03:58+08:00
:tags: JavaScript, reStructuredText, Copy to Clipboard
:category: JavaScript
:summary: Search links of major search engines - Google_, DuckDuckGo_, Bing_,
          Yahoo_, Baidu_, Yandex_, Ecosia_.
:adsu: yes


.. raw:: html

   <input type="text" id="search" tabindex="1" placeholder="Input Search String">
   <button type="button" id="getlinks">Get Links</button><br>
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

      textareaElm.value = ".. [1] | `TERM - Google search <GOOG>`_\n       | `TERM - DuckDuckGo search <DUCK>`_\n       | `TERM - Ecosia search <ECOS>`_\n       | `TERM - Bing search <BING>`_\n       | `TERM - Yahoo search <YAHOO>`_\n       | `TERM - Baidu search <BAIDU>`_\n       | `TERM - Yandex search <YANDEX>`_\n".replace("GOOG", GoogleURL(qstring)).replace("DUCK", DuckDuckGoURL(qstring)).replace("BING", BingURL(qstring)).replace("YAHOO", YahooURL(qstring)).replace("BAIDU", BaiduURL(qstring)).replace("YANDEX", YandexURL(qstring)).replace("ECOS", EcosiaURL(qstring)).replace(/TERM/g, search_string);

      divElm.innerHTML = '<a target="_blank" href="GOOG">TERM - Google Search</a><br><a target="_blank" href="DUCK">TERM - DuckDuckGo Search</a><br><a target="_blank" href="ECOS">TERM - Ecosia Search</a><br><a target="_blank" href="BING">TERM - Bing Search</a><br><a target="_blank" href="YAHOO">TERM - Yahoo Search</a><br><a target="_blank" href="BAIDU">TERM - Baidu Search</a><br><a target="_blank" href="YANDEX">TERM - Yandex Search</a><br>'.replace("GOOG", GoogleURL(qstring)).replace("DUCK", DuckDuckGoURL(qstring)).replace("BING", BingURL(qstring)).replace("YAHOO", YahooURL(qstring)).replace("BAIDU", BaiduURL(qstring)).replace("YANDEX", YandexURL(qstring)).replace("ECOS", EcosiaURL(qstring)).replace(/TERM/g, search_string);
    }

    function GoogleURL(qstring) {
      return "https://www.google.com/search?q=" + qstring;
    }
    function BingURL(qstring) {
      return "https://www.bing.com/search?q=" + qstring;
    }
    function YahooURL(qstring) {
      return "https://search.yahoo.com/search?p=" + qstring;
    }
    function DuckDuckGoURL(qstring) {
      return "https://duckduckgo.com/?q=" + qstring;
    }
    function BaiduURL(qstring) {
      return "https://www.baidu.com/s?wd=" + qstring;
    }
    function YandexURL(qstring) {
      return "https://www.yandex.com/search/?text=" + qstring;
    }
    function EcosiaURL(qstring) {
      return "https://www.ecosia.org/search?q=" + qstring;
    }
  </script>

.. adsu:: 2

----

References:

.. [1] `HTML textarea tag - W3Schools <http://www.w3schools.com/tags/tag_textarea.asp>`_

.. [2] `JavaScript String trim() Method - W3Schools <http://www.w3schools.com/jsref/jsref_trim_string.asp>`_

.. [3] `JavaScript String replace() Method - W3Schools <http://www.w3schools.com/jsref/jsref_replace.asp>`_

.. [4] `Trigger a button click with JavaScript on the Enter key in a text box - Stack Overflow <http://stackoverflow.com/questions/155188/trigger-a-button-click-with-javascript-on-the-enter-key-in-a-text-box>`_

.. [5] | `JavaScript encodeURIComponent() Function - W3Schools <http://www.w3schools.com/jsref/jsref_encodeuricomponent.asp>`_
       | `JavaScript encodeURI() Function - W3Schools <http://www.w3schools.com/jsref/jsref_encodeuri.asp>`_
       | `javascript - When are you supposed to use escape instead of encodeURI / encodeURIComponent? - Stack Overflow <http://stackoverflow.com/questions/75980/when-are-you-supposed-to-use-escape-instead-of-encodeuri-encodeuricomponent>`_
.. adsu:: 3
.. [6] | `javascript copy to clipboard <https://www.google.com/search?q=javascript+copy+to+clipboard>`_
       | `How do I copy to the clipboard in JavaScript? - Stack Overflow <http://stackoverflow.com/questions/400212/how-do-i-copy-to-the-clipboard-in-javascript>`_

.. [7] `Search PTT BBS <{filename}../07/search-ptt-bbs%en.rst>`_

.. [8] `Search Taobao.com <{filename}../../05/06/search-taobao-com%en.rst>`_


.. _Google: https://www.google.com/
.. _DuckDuckGo: https://duckduckgo.com/
.. _Bing: https://www.bing.com/
.. _Yahoo: https://search.yahoo.com/
.. _Baidu: https://www.baidu.com/
.. _Yandex: https://www.yandex.com/
.. _Ecosia: https://www.ecosia.org/

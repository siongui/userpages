[JavaScript] Test Text-to-Speech in Web Speech API
##################################################

:date: 2017-04-22T10:49+08:00
:tags: html, JavaScript, Web Speech API
:category: JavaScript
:summary: Test Text-to-Speech feature in Web Speech API. Not working in Chromium
          57, but OK in Firefox 53.
:og_image: https://sfault-image.b0.upaiyun.com/308/919/3089192407-5714f6f1a0abe
:adsu: yes

Test Text-to-Speech (TTS) feature in `Web Speech API`_. Not working in Chromium
57, but working in Firefox 53 (both running on Ubuntu Linux 17.04).

.. raw:: html

  <div style="padding: 2em; background-color: Azure; padding: 5px;">

  <input type="text" value="hello world" id="text">
  <button type="button" id="tts">Text to Speech</button>

  </div>
  <script>

  var input = document.querySelector("#text");
  var tts = document.querySelector("#tts");

  tts.addEventListener("click", function(e) {
    var text = new SpeechSynthesisUtterance(input.value);
    window.speechSynthesis.speak(text);
  });

  </script>

**HTML**:

.. code-block:: html

  <input type="text" value="hello world" id="text">
  <button type="button" id="tts">Text to Speech</button>

**JavaScript**:

.. code-block:: javascript

  var input = document.querySelector("#text");
  var tts = document.querySelector("#tts");

  tts.addEventListener("click", function(e) {
    var text = new SpeechSynthesisUtterance(input.value);
    window.speechSynthesis.speak(text);
  });

.. adsu:: 2

----

References:

.. [1] `Using Google Text-To-Speech in Javascript - Stack Overflow <http://stackoverflow.com/questions/15653145/using-google-text-to-speech-in-javascript>`_
.. [2] `Web apps that talk - Introduction to the Speech Synthesis API  |  Web  |  Google Developers <https://developers.google.com/web/updates/2014/01/Web-apps-that-talk-Introduction-to-the-Speech-Synthesis-API>`_
.. [3] `javascript - Web speech api not working currently in chromium / electron / nw js? - Stack Overflow <http://stackoverflow.com/questions/36052774/web-speech-api-not-working-currently-in-chromium-electron-nw-js>`_
.. [4] | `使用 JavaScript 进行单词发音 Use JavaScript to Speech Your Text  - 静逸秋水の专栏 - SegmentFault <https://segmentfault.com/a/1190000004963610>`_
       | `使用 JavaScript 进行单词发音 - WEB前端 - 伯乐在线 <http://web.jobbole.com/91132/>`_

.. [5] | `Web Speech API - Google search <https://www.google.com/search?q=Web+Speech+API>`_
       | `Web Speech API - DuckDuckGo search <https://duckduckgo.com/?q=Web+Speech+API>`_
       | `Web Speech API - Ecosia search <https://www.ecosia.org/search?q=Web+Speech+API>`_
       | `Web Speech API - Qwant search <https://www.qwant.com/?q=Web+Speech+API>`_
       | `Web Speech API - Bing search <https://www.bing.com/search?q=Web+Speech+API>`_
       | `Web Speech API - Yahoo search <https://search.yahoo.com/search?p=Web+Speech+API>`_
       | `Web Speech API - Baidu search <https://www.baidu.com/s?wd=Web+Speech+API>`_
       | `Web Speech API - Yandex search <https://www.yandex.com/search/?text=Web+Speech+API>`_

.. _Web Speech API: https://www.google.com/search?q=Web+Speech+API

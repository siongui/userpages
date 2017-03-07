[JavaScript] Type Text Effect
#############################

:date: 2017-03-07T22:58+08:00
:tags: JavaScript, DOM, html
:category: JavaScript
:summary: Type Text Effect by JavaScript_.
:og_image: http://jquery-plugins.net/image/plugin/theaterjs-javascript-typing-effect-plugin.png
:adsu: yes


Demo
++++

.. raw:: html

  <span data-typedtext="May I be happy. May I be free from stress & pain. May I be free from animosity. May I be free from oppression. May I be free from trouble. May I look after myself with ease."></span>
  <br>
  <span data-typedtext="願我得喜樂。願我離憂苦。願我不受敵意。願我不受壓迫。願我免遭困難。願我輕鬆照顧自己。"></span>
  <br>
  <span data-typedtext="ขอให้ข้าพเจ้าจงเป็นผู้ถึงสุข จงเป็นผู้ไร้ทุกข์ จงเป็นผู้ไม่มีเวร จงเป็นผู้ไม่เบียดเบียนซึ่งกันและกัน จงเป็นผู้ไม่มีทุกข์ จงรักษาตนอยู่เป็นสุขเถิด"></span>

  <script>
  'use strict';

  var speed = 80;

  function PlayText() {
    var spans = document.querySelectorAll("span[data-typedtext]");
    for (var i = 0; i < spans.length; ++i) {
      var span = spans[i];
      for (var j = 0; j < span.dataset.typedtext.length; ++j) {
        setTimeout(function(s, l) {
          s.textContent += l;
        }, speed*j, span, span.dataset.typedtext[j]);
      }
    }
  }

  document.addEventListener("DOMContentLoaded", function(event) {
    PlayText();
  });
  </script>

Yesterday I saw the typing effect of the wesbite [3]_, and felt it's really
amazing and impressing, so I thought that it may not be too difficult to
implement the effect in JavaScript. And it is indeed not difficult to implement,
the following is my code.

Source Code
+++++++++++

**HTML**: Put the texts to be typed in *data-typedtext* attribute of *span*
element.

.. code-block:: html

  <span data-typedtext="May I be happy. May I be free from stress & pain. May I be free from animosity. May I be free from oppression. May I be free from trouble. May I look after myself with ease."></span>
  <span data-typedtext="願我得喜樂。願我離憂苦。願我不受敵意。願我不受壓迫。願我免遭困難。願我輕鬆照顧自己。"></span>
  <span data-typedtext="ขอให้ข้าพเจ้าจงเป็นผู้ถึงสุข จงเป็นผู้ไร้ทุกข์ จงเป็นผู้ไม่มีเวร จงเป็นผู้ไม่เบียดเบียนซึ่งกันและกัน จงเป็นผู้ไม่มีทุกข์ จงรักษาตนอยู่เป็นสุขเถิด"></span>

.. adsu:: 2

**JavaScript**:

.. code-block:: javascript

  'use strict';

  var speed = 80;

  function PlayText() {
    var spans = document.querySelectorAll("span[data-typedtext]");
    for (var i = 0; i < spans.length; ++i) {
      var span = spans[i];
      for (var j = 0; j < span.dataset.typedtext.length; ++j) {
        setTimeout(function(s, l) {
          s.textContent += l;
        }, speed*j, span, span.dataset.typedtext[j]);
      }
    }
  }

  document.addEventListener("DOMContentLoaded", function(event) {
    PlayText();
  });

The key point in the above is that we use setTimeout_ to append the *j-th*
letter of the texts in *data-typedtext* attribute to the *textContent* of *span*
element after *j\*speed* milliseconds.

.. adsu:: 3

----

Tested on:
``Chromium Version 56.0.2924.76 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``

----

References:

.. [1] | `personal blog examples - Google search <https://www.google.com/search?q=personal+blog+examples>`_
       | `personal blog examples - DuckDuckGo search <https://duckduckgo.com/?q=personal+blog+examples>`_
       | `personal blog examples - Ecosia search <https://www.ecosia.org/search?q=personal+blog+examples>`_
       | `personal blog examples - Qwant search <https://www.qwant.com/?q=personal+blog+examples>`_
       | `personal blog examples - Bing search <https://www.bing.com/search?q=personal+blog+examples>`_
       | `personal blog examples - Yahoo search <https://search.yahoo.com/search?p=personal+blog+examples>`_
       | `personal blog examples - Baidu search <https://www.baidu.com/s?wd=personal+blog+examples>`_
       | `personal blog examples - Yandex search <https://www.yandex.com/search/?text=personal+blog+examples>`_

.. [2] `18 of the Best Personal Websites We've Ever Seen <https://blog.hubspot.com/marketing/best-personal-websites>`_
.. [3] `STRML: Projects and Work <http://strml.net/>`_

.. [4] | `text animation css - Google search <https://www.google.com/search?q=text+animation+css>`_
       | `text animation css - DuckDuckGo search <https://duckduckgo.com/?q=text+animation+css>`_
       | `text animation css - Ecosia search <https://www.ecosia.org/search?q=text+animation+css>`_
       | `text animation css - Qwant search <https://www.qwant.com/?q=text+animation+css>`_
       | `text animation css - Bing search <https://www.bing.com/search?q=text+animation+css>`_
       | `text animation css - Yahoo search <https://search.yahoo.com/search?p=text+animation+css>`_
       | `text animation css - Baidu search <https://www.baidu.com/s?wd=text+animation+css>`_
       | `text animation css - Yandex search <https://www.yandex.com/search/?text=text+animation+css>`_

.. [5] | `javascript type text effect - Google search <https://www.google.com/search?q=javascript+type+text+effect>`_
       | `javascript type text effect - DuckDuckGo search <https://duckduckgo.com/?q=javascript+type+text+effect>`_
       | `javascript type text effect - Ecosia search <https://www.ecosia.org/search?q=javascript+type+text+effect>`_
       | `javascript type text effect - Qwant search <https://www.qwant.com/?q=javascript+type+text+effect>`_
       | `javascript type text effect - Bing search <https://www.bing.com/search?q=javascript+type+text+effect>`_
       | `javascript type text effect - Yahoo search <https://search.yahoo.com/search?p=javascript+type+text+effect>`_
       | `javascript type text effect - Baidu search <https://www.baidu.com/s?wd=javascript+type+text+effect>`_
       | `javascript type text effect - Yandex search <https://www.yandex.com/search/?text=javascript+type+text+effect>`_

.. [6] | `javascript content loaded - Google search <https://www.google.com/search?q=javascript+content+loaded>`_
       | `javascript content loaded - DuckDuckGo search <https://duckduckgo.com/?q=javascript+content+loaded>`_
       | `javascript content loaded - Ecosia search <https://www.ecosia.org/search?q=javascript+content+loaded>`_
       | `javascript content loaded - Qwant search <https://www.qwant.com/?q=javascript+content+loaded>`_
       | `javascript content loaded - Bing search <https://www.bing.com/search?q=javascript+content+loaded>`_
       | `javascript content loaded - Yahoo search <https://search.yahoo.com/search?p=javascript+content+loaded>`_
       | `javascript content loaded - Baidu search <https://www.baidu.com/s?wd=javascript+content+loaded>`_
       | `javascript content loaded - Yandex search <https://www.yandex.com/search/?text=javascript+content+loaded>`_

.. [7] | `javascript sleep - Google search <https://www.google.com/search?q=javascript+sleep>`_
       | `javascript sleep - DuckDuckGo search <https://duckduckgo.com/?q=javascript+sleep>`_
       | `javascript sleep - Ecosia search <https://www.ecosia.org/search?q=javascript+sleep>`_
       | `javascript sleep - Qwant search <https://www.qwant.com/?q=javascript+sleep>`_
       | `javascript sleep - Bing search <https://www.bing.com/search?q=javascript+sleep>`_
       | `javascript sleep - Yahoo search <https://search.yahoo.com/search?p=javascript+sleep>`_
       | `javascript sleep - Baidu search <https://www.baidu.com/s?wd=javascript+sleep>`_
       | `javascript sleep - Yandex search <https://www.yandex.com/search/?text=javascript+sleep>`_

.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _setTimeout: https://www.w3schools.com/jsref/met_win_settimeout.asp

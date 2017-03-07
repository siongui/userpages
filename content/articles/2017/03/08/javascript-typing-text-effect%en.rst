[JavaScript] Typing Text Effect
###############################

:date: 2017-03-08T03:00+08:00
:tags: JavaScript, DOM, html, Typing Text Effect
:category: JavaScript
:summary: Sequential typing text effect by JavaScript_.
:og_image: http://jquery-plugins.net/image/plugin/theaterjs-javascript-typing-effect-plugin.png
:adsu: yes


Demo
++++

.. raw:: html

  <div style="background-color: Azure; padding: 5px;">
  <span data-typedtext="May I be happy. May I be free from stress & pain. May I be free from animosity. May I be free from oppression. May I be free from trouble. May I look after myself with ease."></span>
  <br><br>
  <span data-typedtext="願我得喜樂。願我離憂苦。願我不受敵意。願我不受壓迫。願我免遭困難。願我輕鬆照顧自己。"></span>
  <br><br>
  <span data-typedtext="ขอให้ข้าพเจ้าจงเป็นผู้ถึงสุข จงเป็นผู้ไร้ทุกข์ จงเป็นผู้ไม่มีเวร จงเป็นผู้ไม่เบียดเบียนซึ่งกันและกัน จงเป็นผู้ไม่มีทุกข์ จงรักษาตนอยู่เป็นสุขเถิด"></span>
  </div>

  <script>
  'use strict';

  var speed = 15;

  function Playing(elements, elementIndex, textIndex) {
    if (elementIndex == elements.length) {
      return;
    }
    var element = elements[elementIndex];
    if (textIndex == element.dataset.typedtext.length) {
      setTimeout(Playing, speed, elements, elementIndex + 1, 0);
    } else {
      element.textContent += element.dataset.typedtext[textIndex];
      setTimeout(Playing, speed, elements, elementIndex, textIndex + 1);
    }
  }

  function PlaySequential() {
    var spans = document.querySelectorAll("span[data-typedtext]");
    Playing(spans, 0, 0);
  }

  document.addEventListener("DOMContentLoaded", function(event) {
    PlaySequential();
  });
  </script>

The three lines in demo are being typed sequentially. If you want the three
lines being typed parallel, see `my another post`_ [1]_.

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

  var speed = 15;

  function Playing(elements, elementIndex, textIndex) {
    if (elementIndex == elements.length) {
      return;
    }
    var element = elements[elementIndex];
    if (textIndex == element.dataset.typedtext.length) {
      setTimeout(Playing, speed, elements, elementIndex + 1, 0);
    } else {
      element.textContent += element.dataset.typedtext[textIndex];
      setTimeout(Playing, speed, elements, elementIndex, textIndex + 1);
    }
  }

  function PlaySequential() {
    var spans = document.querySelectorAll("span[data-typedtext]");
    Playing(spans, 0, 0);
  }

  document.addEventListener("DOMContentLoaded", function(event) {
    PlaySequential();
  });

The key point in the above code is that the *Playing* function calls itself
repeatedly by setTimeout_, and in each function call, only one character or
letter is appended to the *textContent* of the element.

.. adsu:: 3

----

Tested on:
``Chromium Version 56.0.2924.76 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``

----

References:

.. [1] `[JavaScript] Type Text Effect <{filename}../07/javascript-type-text-effect%en.rst>`_
.. [2] | `typing text effect css - Google search <https://www.google.com/search?q=typing+text+effect+css>`_
       | `typing text effect css - DuckDuckGo search <https://duckduckgo.com/?q=typing+text+effect+css>`_
       | `typing text effect css - Ecosia search <https://www.ecosia.org/search?q=typing+text+effect+css>`_
       | `typing text effect css - Qwant search <https://www.qwant.com/?q=typing+text+effect+css>`_
       | `typing text effect css - Bing search <https://www.bing.com/search?q=typing+text+effect+css>`_
       | `typing text effect css - Yahoo search <https://search.yahoo.com/search?p=typing+text+effect+css>`_
       | `typing text effect css - Baidu search <https://www.baidu.com/s?wd=typing+text+effect+css>`_
       | `typing text effect css - Yandex search <https://www.yandex.com/search/?text=typing+text+effect+css>`_

.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _setTimeout: https://www.w3schools.com/jsref/met_win_settimeout.asp
.. _my another post: {filename}../07/javascript-type-text-effect%en.rst

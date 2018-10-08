JavaScript 鍵盤事件 (方向鍵範例)
################################

:date: 2018-10-08T22:24+08:00
:tags: JavaScript, 鍵盤事件, 文檔對象模型(DOM)
:category: JavaScript
:summary: 利用 *JavaScript* 偵測 `方向鍵按鍵`_ 。
:og_image: http://www.javatpoint.com/images/javascript/javascript_logo.png
:adsu: yes


本文給出一個利用 JavaScript_ 來偵測 `方向鍵`_ 的範例。
至於一般的 `鍵盤事件`_ ，只要把此範例稍微修改延伸即可達成
（請順便閱讀下方參考）。先試看看下方demo：

.. raw:: html

  <div style="background-color: Azure; padding: 1rem;">

  Press Arrow Key or Any Other Key:<br><br>
  <div id="info"></div>

  </div>

.. raw:: html

  <script>
  var Key = {
    LEFT:   37,
    UP:     38,
    RIGHT:  39,
    DOWN:   40
  };

  /**
   * old IE: attachEvent
   * Firefox, Chrome, or modern browsers: addEventListener
   */
  function _addEventListener(evt, element, fn) {
    if (window.addEventListener) {
      element.addEventListener(evt, fn, false);
    }
    else {
      element.attachEvent('on'+evt, fn);
    }
  }

  function handleKeyboardEvent(evt) {
    if (!evt) {evt = window.event;} // for old IE compatible
    var keycode = evt.keyCode || evt.which; // also for cross-browser compatible

    var info = document.getElementById("info");
    switch (keycode) {
      case Key.LEFT:
        info.innerHTML += "&larr; ";
        break;
      case Key.UP:
        info.innerHTML += "&uarr; ";
        break;
      case Key.RIGHT:
        info.innerHTML += "&rarr; ";
        break;
      case Key.DOWN:
        info.innerHTML += "&darr; ";
        break;
      default:
        info.innerHTML += "SOMEKEY ";
    }
  }

  _addEventListener('keydown', document, handleKeyboardEvent);
  </script>

Demo原始碼 (*HTML*):

.. code-block:: html

  Press Arrow Key or Any Other Key:<br><br>
  <div id="info"></div>

.. adsu:: 2

Demo原始碼 (*JavaScript*):

.. code-block:: javascript

  var Key = {
    LEFT:   37,
    UP:     38,
    RIGHT:  39,
    DOWN:   40
  };

  /**
   * old IE: attachEvent
   * Firefox, Chrome, or modern browsers: addEventListener
   */
  function _addEventListener(evt, element, fn) {
    if (window.addEventListener) {
      element.addEventListener(evt, fn, false);
    }
    else {
      element.attachEvent('on'+evt, fn);
    }
  }

  function handleKeyboardEvent(evt) {
    if (!evt) {evt = window.event;} // for old IE compatible
    var keycode = evt.keyCode || evt.which; // also for cross-browser compatible

    var info = document.getElementById("info");
    switch (keycode) {
      case Key.LEFT:
        info.innerHTML += "&larr; ";
        break;
      case Key.UP:
        info.innerHTML += "&uarr; ";
        break;
      case Key.RIGHT:
        info.innerHTML += "&rarr; ";
        break;
      case Key.DOWN:
        info.innerHTML += "&darr; ";
        break;
      default:
        info.innerHTML += "SOMEKEY ";
    }
  }

  _addEventListener('keydown', document, handleKeyboardEvent);

跟鍵盤相關的事件有三個： *onkeydown*, *onkeypress*, *onkeyup* 。
偵測方向鍵，請用 onkeydown_ (參看 [2]_).

根據 MDN_ 的 KeyboardEvent_ 文檔，已經不贊成使用 event.keyCode_ 。
應使用 event.key_ 。如果您想要支援非常舊的瀏覽器，請使用本文的範例。
否則請參看 [5]_ 使用 event.key_ 的新範例。

.. adsu:: 3

----

**參考**

.. [1] `Keyboard events | JavaScript Tutorial <https://javascript.info/tutorial/keyboard-events>`_
.. [2] `keyboard events - Detecting arrow key presses in JavaScript - Stack Overflow <https://stackoverflow.com/questions/5597060/detecting-arrow-key-presses-in-javascript>`_
.. [3] `JavaScript - Detecting keystrokes - QuirksMode <https://www.quirksmode.org/js/keys.html>`_
.. [4] `JavaScript Madness: Keyboard Events <http://unixpapa.com/js/key.html>`_
.. [5] `JavaScript Arrow Key Example via event.key in Keyboard Event <{filename}../../../2017/02/14/javascript-arrow-key-example-via-event-key%en.rst>`_

.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _方向鍵按鍵: https://www.google.com/search?q=arrow+keystrokes
.. _方向鍵: https://www.google.com/search?q=arrow+keys
.. _鍵盤事件: https://www.google.com/search?q=keyboard+event
.. _onkeydown: http://www.w3schools.com/jsref/event_onkeydown.asp
.. _KeyboardEvent: https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent
.. _MDN: https://developer.mozilla.org/
.. _event.key: https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key
.. _event.keyCode: https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/keyCode

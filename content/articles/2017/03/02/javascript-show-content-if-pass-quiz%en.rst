[JavaScript] Show Content If Pass Quiz
######################################

:date: 2017-03-02T04:03+08:00
:tags: JavaScript, DOM, html, Keyboard Event
:category: JavaScript
:summary: Show youtube video if users pass the quiz.
:og_image: http://www.javatpoint.com/images/javascript/javascript_logo.png
:adsu: yes


Show youtube video if users pass the quiz. See demo first. Please answer the
following quiz.

.. raw:: html

  <div style="background-color: Azure; padding: 1em;">

  <div>13 + 51 = ?</div>
  Input answer and press enter: <input id="ans" type="text"><br><br>
  <div id="info"></div>

  <script>
  var ans = document.getElementById("ans");
  var info = document.getElementById("info");
  ans.addEventListener("keydown", function(e) {
    if (e.key == "Enter") {
      if (e.target.value.trim() == "64") {
        // correct answer, show youtube video
        info.innerHTML = '<div style="position:relative;height:0;padding-bottom:75.0%"><iframe src="https://www.youtube.com/embed/UOTKWp-5bKI?ecver=2" width="480" height="360" frameborder="0" style="position:absolute;width:100%;height:100%;left:0" allowfullscreen></iframe></div>';
      } else {
        info.innerHTML = "wrong answer!";
      }
    }
  });
  </script>
  </div>

HTML_:

.. code-block:: html

  <div>13 + 51 = ?</div>
  Input answer and press enter: <input id="ans" type="text"><br><br>
  <div id="info"></div>

JavaScript_:

.. code-block:: javascript

  var ans = document.getElementById("ans");
  var info = document.getElementById("info");
  ans.addEventListener("keydown", function(e) {
    if (e.key == "Enter") {
      if (e.target.value.trim() == "64") {
        // correct answer, show youtube video
        info.innerHTML = '<div style="position:relative;height:0;padding-bottom:75.0%"><iframe src="https://www.youtube.com/embed/UOTKWp-5bKI?ecver=2" width="480" height="360" frameborder="0" style="position:absolute;width:100%;height:100%;left:0" allowfullscreen></iframe></div>';
      } else {
        info.innerHTML = "wrong answer!";
      }
    }
  });

Register a keyboard event listener to the input element where users input the
answer. If the user's answer is correct, embed iframe of the youtube video.
Otherwise tell the user that the answer is wrong.

.. adsu:: 2

----

Tested on:

- ``Chromium Version 55.0.2883.87 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``

----

References:

.. [1] `KeyboardEvent.key - Web APIs | MDN <https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key>`_

.. _HTML: https://www.google.com/search?q=HTML
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _DOM element: https://www.google.com/search?q=DOM+element
.. _classList: https://www.w3schools.com/jsref/prop_element_classlist.asp
.. _toggle: https://developer.mozilla.org/en/docs/Web/API/Element/classList

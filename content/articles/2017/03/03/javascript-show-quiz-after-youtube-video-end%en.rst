[JavaScript] Show Quiz After End of YouTube Video
#################################################

:date: 2017-03-03T00:09+08:00
:tags: JavaScript, DOM, html, YouTube Player API, CSS
:category: JavaScript
:summary: Quiz will appear at the end of YouTube video.
:og_image: http://www.javatpoint.com/images/javascript/javascript_logo.png
:adsu: yes


Demo
++++

Show quiz after the end of YouTube video. The quiz will appear at the end of
video. See demo first.

.. raw:: html

  <style>
  .invisible { display: none; }
  #quiz {color: red; font-size: larger; margin-top: 1.25em;}
  </style>
  <div style="background-color: Azure; padding: 1em;">
    <div id="player"></div>
    <div id="quiz" class="invisible">Quiz: What is the name of the video?</div>
  </div>
  <script>

  var quiz = document.getElementById("quiz");

  var tag = document.createElement('script');
  tag.src = "https://www.youtube.com/iframe_api";
  var firstScriptTag = document.getElementsByTagName('script')[0];
  firstScriptTag.parentNode.insertBefore(tag, firstScriptTag);

  var player;
  function onYouTubeIframeAPIReady() {
    player = new YT.Player('player', {
      height: '390',
      width: '640',
      videoId: 'hS5CfP8n_js',
      events: {
        'onReady': onPlayerReady,
        'onStateChange': onPlayerStateChange
      }
    });
  }

  function onPlayerReady(event) {
    event.target.playVideo();
  }

  function onPlayerStateChange(event) {
    if (event.data == YT.PlayerState.ENDED) {
      quiz.classList.remove("invisible");
    }
  }

  </script>


.. adsu:: 2

Source Code
+++++++++++

HTML_: Note that put the HTML code before JavaScript_ code, otherwise
getElementById will return null.

.. code-block:: html

  <div id="player"></div>
  <div id="quiz" class="invisible">Quiz: What is the name of the video?</div>

CSS_: the class ``.invisible`` is used to hide the quiz initially. After the end
of YouTube video, this class will be removed from the *div* of quiz.

.. code-block:: css

  .invisible { display: none; }
  #quiz {color: red; font-size: larger; margin-top: 1.25em;}

JavaScript_: Basically the same as the example of YouTube API reference [1]_,
except that check if the video is ended in *onPlayerStateChange*, and show the
quiz if the video is ended.

.. code-block:: javascript

  var quiz = document.getElementById("quiz");

  var tag = document.createElement('script');
  tag.src = "https://www.youtube.com/iframe_api";
  var firstScriptTag = document.getElementsByTagName('script')[0];
  firstScriptTag.parentNode.insertBefore(tag, firstScriptTag);

  var player;
  function onYouTubeIframeAPIReady() {
    player = new YT.Player('player', {
      height: '390',
      width: '640',
      videoId: 'hS5CfP8n_js',
      events: {
        'onReady': onPlayerReady,
        'onStateChange': onPlayerStateChange
      }
    });
  }

  function onPlayerReady(event) {
    event.target.playVideo();
  }

  function onPlayerStateChange(event) {
    if (event.data == YT.PlayerState.ENDED) {
      quiz.classList.remove("invisible");
    }
  }

.. adsu:: 3

You can force the user to watch the full YouTube video by the trick [2]_. But
this is not suggested because it is too annoying.

----

Tested on:

- ``Chromium Version 55.0.2883.87 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``

----

References:

.. [1] `YouTube Player API Reference for iframe Embeds | YouTube IFrame Player API | Google Developers <https://developers.google.com/youtube/iframe_api_reference>`_
.. [2] `Force a user to watch the whole video - Stack Overflow <http://stackoverflow.com/questions/29082759/force-a-user-to-watch-the-whole-video>`_

.. _HTML: https://www.google.com/search?q=HTML
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _CSS: https://www.google.com/search?q=CSS

[JavaScript] Toggle (Play/Pause) Sound on Click Event of DOM Element
####################################################################

:date: 2012-10-12 06:25
:modified: 2015-02-25 12:10
:tags: html, HTML5 audio, JavaScript
:category: JavaScript
:summary: Toggle sound on click event of DOM element.
:og_image: http://www.javatpoint.com/images/javascript/javascript_logo.png
:adsu: yes


This post shows how to toggle (play/pause) sound by JavaScript manipulation of
onclick event of a DOM element. Browser support of HTML5 audio_ tag is a must in
this example.

.. rubric:: `Demo <{filename}/code/javascript/play-toggle-sound/toggle.html>`_
      :class: align-center

Source Code for Demo:

.. show_github_file:: siongui userpages content/code/javascript/play-toggle-sound/toggle.html

.. adsu:: 2

When you click on the *button* element, the sound will be played. If the element
is clicked again, the sound will be paused. To toggle sound like this, a HTML5
audio_ element is embedded in the HTML document, and not displayed on screen.
Every time the *button* element is clicked, the *toggleSound* function will be
executed. The *toggleSound* function checks if the *audio* element is paused_.
If the audio element is paused, call `play()`_ to play sound. Otherwise call
`pause()`_ to stop playing.

----

.. adsu:: 3

References:

.. [1] `HTML Audio/Video DOM Reference <http://www.w3schools.com/tags/ref_av_dom.asp>`_

.. [2] `HTML5 Audio <http://www.w3schools.com/html/html5_audio.asp>`_

.. [3] `Introduction to the HTML5 audio tag javascript manipulation <http://www.position-absolute.com/articles/introduction-to-the-html5-audio-tag-javascript-manipulation/>`_

.. [4] `[JavaScript] Play Sound on Click Event of DOM Element <{filename}../08/javascript-play-sound-onclick%en.rst>`_

.. [5] `[Golang] GopherJS DOM Example - Play Sound on Click Event <{filename}../../../2016/01/15/gopherjs-dom-example-play-sound-onclick-event%en.rst>`_

.. [6] `[Golang] GopherJS DOM Example - Toggle (Play/Pause) Sound on Click Event <{filename}../../../2016/01/15/gopherjs-dom-example-toggle-sound-onclick-event%en.rst>`_


.. _audio: http://www.w3schools.com/tags/tag_audio.asp

.. _paused: http://www.w3schools.com/tags/av_prop_paused.asp

.. _play(): http://www.w3schools.com/tags/av_met_play.asp

.. _pause(): http://www.w3schools.com/tags/av_met_pause.asp

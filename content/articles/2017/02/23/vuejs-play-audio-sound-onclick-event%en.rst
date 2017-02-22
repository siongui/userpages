[Vue.js] Play Audio Sound on Click Event of DOM Element
#######################################################

:date: 2017-02-23T02:51+08:00
:tags: Vue.js, JavaScript, DOM, HTML5 audio, html
:category: Vue.js
:summary: Play audio sound if users click on `DOM Element`_ in Vue.js_.
:og_image: https://vuejs.org/images/logo.png
:adsu: yes


Play audio sound if uers click on `DOM Element`_ in Vue.js_/JavaScript_.

.. rubric:: `Demo <{filename}/code/vuejs/play-sound/index.html>`_
   :class: align-center

The key points of the following source code:
  - Use v-on_ directive to attach event listener of button_ onClick_ event.
  - Use ref_ attribute to register a reference to audio_ element. We will later
    access the audio_ element by the unique reference name *audioElm* in the
    event listener.
  - If users click on the button, call `play()`_ method of the audio element to
    play the sound.

.. adsu:: 2
.. show_github_file:: siongui userpages content/code/vuejs/play-sound/index.html
.. adsu:: 3
.. show_github_file:: siongui userpages content/code/vuejs/play-sound/app.js
.. adsu:: 4

----

Tested on:
  - ``Chromium Version 55.0.2883.87 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``
  - ``Vue.js 2.1.10``.

----

References:

.. [1] | `vuejs reference element - Google search <https://www.google.com/search?q=vuejs+reference+element>`_
       | `Special Attributes - ref - Vue.js <https://vuejs.org/v2/api/#ref>`_
.. [2] `Event Handling — Vue.js <https://vuejs.org/v2/guide/events.html>`_
.. adsu:: 5
.. [3] | `<audio> - HTML | MDN <https://developer.mozilla.org/en-US/docs/Web/HTML/Element/audio>`_
       | `HTML Audio/Video DOM Reference <https://www.w3schools.com/TAGs/ref_av_dom.asp>`_

.. _HTML: https://www.google.com/search?q=HTML
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _onClick: https://www.w3schools.com/jsref/event_onclick.asp
.. _button: https://www.w3schools.com/tags/tag_button.asp
.. _audio: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/audio
.. _DOM element: https://www.google.com/search?q=DOM+element
.. _Vue.js: https://vuejs.org/
.. _v-on: https://vuejs.org/v2/api/#v-on
.. _ref: https://vuejs.org/v2/api/#ref
.. _play(): https://www.w3schools.com/tags/av_met_play.asp

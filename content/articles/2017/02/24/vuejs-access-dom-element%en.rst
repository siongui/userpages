[Vue.js] Access (Get/Reference) DOM Element
###########################################

:date: 2017-02-24T03:09+08:00
:tags: Vue.js, JavaScript, DOM, html
:category: Vue.js
:summary: How to access (get/reference) `DOM Element`_ in Vue.js_?
:og_image: https://vuejs.org/images/logo.png
:adsu: yes


How to access (get/reference) `DOM Element`_ in Vue.js_?

Assume we have the following audio_ element in HTML code:

.. code-block:: html

  <audio src="Wat_Metta_Buddha_Qualities.mp3"></audio>

We can give id to the audio element in HTML:

.. code-block:: html

  <audio id="audioElm" src="Wat_Metta_Buddha_Qualities.mp3"></audio>

And then access the audio element via getElementById_ in JavaScript:

.. code-block:: javascript

  var myAudio = document.getElementById("audioElm");
  // do something with myAudio

But this seems not an elegant and canonical way to reference a DOM element in
Vue.js_. So I tried some googling [1]_, and found no working answer for my Vue
version 2.1.10.

Then I read the documentation of Vue.js_, and found the ref_ attribute, which is
exactly what I need.

.. adsu:: 2

The solution [2]_ is to give the ref_ name to the audio element in HTML:

.. code-block:: html

  <audio ref="audioElm" src="Wat_Metta_Buddha_Qualities.mp3"></audio>

Inside the `Vue instance`_, access the `DOM element`_ by ``this.$refs.audioElm``
(this_ here is the Vue instance). For example, you can play the audio in
JavaScript code by:

.. code-block:: javascript

  this.$refs.audioElm.play();

Or stop playing by:

.. code-block:: javascript

  this.$refs.audioElm.pause();

The instance property `$refs`_ holds the DOM elements that have ref_ registered.

If you need complete examples of getting DOM elements in Vue.js_, see my posts
of `play audio`_ [3]_ and `toggle audio`_ [4]_.

----

Tested on:
  - ``Chromium Version 55.0.2883.87 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``
  - ``Vue.js 2.1.10``.

----

.. adsu:: 3

References:

.. [1] `vuejs reference element - Google search <https://www.google.com/search?q=vuejs+reference+element>`_
.. [2] | `Special Attributes - ref - Vue.js <https://vuejs.org/v2/api/#ref>`_
       | `Instance Properties - vm.$refs - Vue.js <https://vuejs.org/v2/api/#vm-refs>`_
.. [3] `[Vue.js] Play Audio Sound on Click Event of DOM Element <{filename}../23/vuejs-play-audio-sound-onclick-event%en.rst>`_
.. [4] `[Vue.js] Toggle (Play/Pause) Audio Sound on Click Event of DOM Element <{filename}vuejs-toggle-audio-sound-onclick-event%en.rst>`_

.. _HTML: https://www.google.com/search?q=HTML
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _audio: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/audio
.. _getElementById: https://www.google.com/search?q=getElementById
.. _DOM element: https://www.google.com/search?q=DOM+element
.. _Vue.js: https://vuejs.org/
.. _ref: https://vuejs.org/v2/api/#ref
.. _Vue instance: https://vuejs.org/v2/guide/instance.html
.. _this: https://www.google.com/search?q=javascript+this
.. _$refs: https://vuejs.org/v2/api/#vm-refs
.. _play audio: {filename}../23/vuejs-play-audio-sound-onclick-event%en.rst
.. _toggle audio: {filename}vuejs-toggle-audio-sound-onclick-event%en.rst

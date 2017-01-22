Load Dart Script if Dartium, Otherwise Load JavaScript in Browser
#################################################################

:tags: JavaScript, Dart, html
:category: Dart
:summary: Load Dart script while development or the browser supports dart. Otherwise load JavaScript as usual.
:adsu: yes

Assume that we write a dart script called *app.dart*, and the corresponding
compiled JavaScript file is called *app.js*. We want the browser to load
*app.dart* if the browser supports dart, otherwise load *app.js* for
compatibility. How do we do?

Put the following code in your html before the end of *body* tag.

.. code-block:: javascript

  <script type='text/javascript'>
    var script = document.createElement('script');
    if (navigator.userAgent.indexOf('(Dart)') === -1) {
      // Browser doesn't support Dart
      script.setAttribute("type","text/javascript");
      script.setAttribute("src", "app.js");
    } else {
      script.setAttribute("type","application/dart");
      script.setAttribute("src", "app.dart");
    }
    document.getElementsByTagName("head")[0].appendChild(script);
  </script>

----

References:

.. [1] `Dynamically loading an external JavaScript or CSS file <http://www.javascriptkit.com/javatutors/loadjavascriptcss.shtml>`_

.. [2] Detect language unavailability - `Synonyms - Dart, JavaScript, C#, Python <https://www.dartlang.org/docs/synonyms/>`_

Add Social Buttons (Google+, Facebook, Twitter) on Website
##########################################################

:date: 2012-09-29 14:51
:modified: 2015-03-13 04:41
:tags: Social Button
:category: Web
:summary: Use social buttons to make your website more popular.
:adsu: yes


There are already too many posts on social buttons, and I think the best
references is the official sites of social buttons (See References_ for social
buttons of *Google+*, *Facebook*, and *Twitter*). The following is the demo of
social buttons of this blog:

----

.. raw:: html

  <div class="floating-buttons">
    <div class="g-plusone" data-href="https://siongui.github.io/" data-size="tall"></div>
    <div class="fb-like" data-href="https://siongui.github.io/" data-layout="box_count" data-send="true" data-show-faces="true"></div>
    <div>
      <a class="twitter-share-button" data-count="vertical" data-url="https://siongui.github.io/" href="https://twitter.com/share">Tweet</a>
    </div>
  </div>
  <script type="text/javascript">
    (function() {
      var po = document.createElement('script'); po.type = 'text/javascript'; po.async = true;
      po.src = 'https://apis.google.com/js/plusone.js';
      var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(po, s);
    })();
  </script>

  <br />
  <div id="fb-root">
  </div>
  <script>(function(d, s, id) {
    var js, fjs = d.getElementsByTagName(s)[0];
    if (d.getElementById(id)) return;
    js = d.createElement(s); js.id = id;
    js.src = "//connect.facebook.net/en_US/all.js#xfbml=1";
    fjs.parentNode.insertBefore(js, fjs);
  }(document, 'script', 'facebook-jssdk'));</script>

  <script>!function(d,s,id){var js,fjs=d.getElementsByTagName(s)[0];if(!d.getElementById(id)){js=d.createElement(s);js.id=id;js.src="//platform.twitter.com/widgets.js";fjs.parentNode.insertBefore(js,fjs);}}(document,"script","twitter-wjs");</script>

  <style>
  div.floating-buttons {
    padding-top: 2em;
  }

  div.floating-buttons > div {
    display: block;
    margin-top: 10px;
  }
  </style>

----

The source code of the above demo:

.. code-block:: html
  :linenos: table

  <div class="floating-buttons">
    <div class="g-plusone" data-href="https://siongui.github.io/" data-size="tall"></div>
    <div class="fb-like" data-href="https://siongui.github.io/" data-layout="box_count" data-send="true" data-show-faces="true"></div>
    <div>
      <a class="twitter-share-button" data-count="vertical" data-url="https://siongui.github.io/" href="https://twitter.com/share">Tweet</a>
    </div>
  </div>

If you want to use the above code, remember to the change the URL to your
website in line 2, 3, and 5. And then put them inside the *body* tag of HTML
document.

.. code-block:: html

  <!-- Place this tag after the last +1 button tag. -->
  <script type="text/javascript">
    (function() {
      var po = document.createElement('script'); po.type = 'text/javascript'; po.async = true;
      po.src = 'https://apis.google.com/js/plusone.js';
      var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(po, s);
    })();
  </script>

  <br />
  <div id="fb-root">
  </div>
  <script>(function(d, s, id) {
    var js, fjs = d.getElementsByTagName(s)[0];
    if (d.getElementById(id)) return;
    js = d.createElement(s); js.id = id;
    js.src = "//connect.facebook.net/en_US/all.js#xfbml=1";
    fjs.parentNode.insertBefore(js, fjs);
  }(document, 'script', 'facebook-jssdk'));</script>

  <script>!function(d,s,id){var js,fjs=d.getElementsByTagName(s)[0];if(!d.getElementById(id)){js=d.createElement(s);js.id=id;js.src="//platform.twitter.com/widgets.js";fjs.parentNode.insertBefore(js,fjs);}}(document,"script","twitter-wjs");</script>

You can put above code right before the end of the *body* tag of HTML document.

.. code-block:: css

  div.floating-buttons {
    padding-top: 2em;
  }

  div.floating-buttons > div {
    display: block;
    margin-top: 10px;
  }

You can put the above code in your CSS file.

To have more customization of the social buttons, please see References_ below.

----

References
++++++++++

.. [1] `+1 Button - Google+ Platform — Google Developers <https://developers.google.com/+/web/+1button/>`_

.. [2] `Facebook Like Button <https://developers.facebook.com/docs/plugins/like-button>`_

.. [3] `Twitter Buttons | About <https://about.twitter.com/resources/buttons>`_ (Basic Version)

.. [4] `Tweet Button | Twitter Developers <https://dev.twitter.com/web/tweet-button>`_ (Advanced Version)

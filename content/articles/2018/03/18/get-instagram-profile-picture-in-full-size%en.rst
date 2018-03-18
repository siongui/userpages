Get Instagram Profile Picture in Full Size
##########################################

:date: 2018-03-18T23:00+08:00
:tags: JavaScript, String Manipulation, Instagram, Regular Expression
:category: JavaScript
:summary: Given the url of Instagram profile picture, return the url of profile
          picture in full size.
:og_image: https://images.vexels.com/media/users/3/147102/isolated/preview/082213cb0f9eabb7e6715f59ef7d322a-instagram-profile-icon-by-vexels.png
:adsu: yes


Paste URL of Instagram_ profile picture and get full-size URL of it.
For example, you can try to paste the following URL:

  *https://instagram.fkhh1-2.fna.fbcdn.net/vp/dd4c00ec38e64a8cc8340771e555ea62/5B3BD6AB/t51.2885-19/s320x320/14719833_310540259320655_1605122788543168512_a.jpg*

.. raw:: html

  <button id="get" type="button">Get Full-Size URL</button>
  <input id="oriUrl" type="text" name="profile_pic_url"><br>
  <div id="result"></div>

  <style>
  #result {
    margin-top: .5rem;
    margin-bottom: .5rem;
  }
  </style>
  <script>
  function getFullSizeProfilePicUrl(picurl) {
    var fullsizeurl = picurl.replace("/vp/", "/");
    fullsizeurl = fullsizeurl.replace("t51.2885-19", "t51.2885-15");
    fullsizeurl = fullsizeurl.replace(/\/s\d+x\d+\//, "/sh0.08/e35/");
    return fullsizeurl;
  }

  var btn = document.querySelector("#get");
  var result = document.querySelector("#result");

  btn.addEventListener("click", function(e) {
    var fullUrl = getFullSizeProfilePicUrl(document.querySelector("#oriUrl").value);

    var link = document.createElement("a");
    link.setAttribute("href", fullUrl);
    link.setAttribute("target", "_blank");
    link.appendChild(document.createTextNode("View Full Size Profile Pic"));

    result.innerHTML = "";
    result.appendChild(link);
  });
  </script>

The trick comes from `the comment of SO question`_ [3]_. The following is the
implementation of the trick in JavaScript:

.. code-block:: javascript

  function getFullSizeProfilePicUrl(picurl) {
    var fullsizeurl = picurl.replace("/vp/", "/");
    fullsizeurl = fullsizeurl.replace("t51.2885-19", "t51.2885-15");
    fullsizeurl = fullsizeurl.replace(/\/s\d+x\d+\//, "/sh0.08/e35/");
    return fullsizeurl;
  }

----

.. adsu:: 2

References:

.. [1] `download instagram profile picture - Google search <https://www.google.com/search?q=download+instagram+profile+picture>`_
.. [2] `How to download my Instagram profile picture - Quora <https://www.quora.com/How-can-I-download-my-Instagram-profile-picture>`_
.. [3] `How to view instagram profile picture in full-size? - Stack Overflow <https://stackoverflow.com/questions/48468144/how-to-view-instagram-profile-picture-in-full-size>`_

.. _Instagram: https://www.instagram.com/
.. _the comment of SO question: https://stackoverflow.com/questions/48468144/how-to-view-instagram-profile-picture-in-full-size#comment85451994_48468144

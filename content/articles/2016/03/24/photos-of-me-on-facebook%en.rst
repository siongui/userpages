Photos of Me on Facebook
########################

:date: 2016-03-24T02:17+08:00
:tags: JavaScript
:category: JavaScript
:summary: Find photos of me on Facebook_.
:adsu: yes


1. log into Facebook_.

2. .. raw:: html

     <p><a class="reference external" href="http://findmyfbid.com/" target="_blank">Find my Facebook ID</a>.</p>

3. .. raw:: html

     <button type="button" id="getlinks">Get Links</button>
     <input type="text" id="myfbid" placeholder="my Facebook ID">
     <div id="links"></div>

.. raw:: html

  <script>

  var linksElm = document.getElementById("links")

  function photosURL(id) {
    return "https://www.facebook.com/search/" + id + "/photos";
  }

  function videosURL(id) {
    return "https://www.facebook.com/search/" + id + "/videos";
  }

  function photosofURL(id) {
    return "https://www.facebook.com/search/" + id + "/photos-of";
  }

  function videosofURL(id) {
    return "https://www.facebook.com/search/" + id + "/videos-of";
  }

  function photoslikedURL(id) {
    return "https://www.facebook.com/search/" + id + "/photos-liked";
  }

  function photoscommentedURL(id) {
    return "https://www.facebook.com/search/" + id + "/photos-commented";
  }

  document.getElementById("getlinks").onclick = function() {
    var id = document.getElementById("myfbid").value.trim();

    while (linksElm.hasChildNodes()) {
      linksElm.removeChild(linksElm.lastChild);
    }

    linksElm.innerHTML = '<br> \
      <div><a href="PHOTOS" target="_blank">me\'s photos</a></div><br> \
      <div><a href="VIDEOS" target="_blank">me\'s videos</a></div><br> \
      <div><a href="PHOTOSOF" target="_blank">Photos of me</a></div><br> \
      <div><a href="VIDEOSOF" target="_blank">Videos of me</a></div><br> \
      <div><a href="PHOTOSLIKED" target="_blank">Photos liked by me</a></div><br> \
      <div><a href="PHOTOSCOMMENTED" target="_blank">Photos commented on by me</a></div><br> \
    '.replace("PHOTOSOF", photosofURL(id)).replace("PHOTOSLIKED", photoslikedURL(id)).replace("PHOTOSCOMMENTED", photoscommentedURL(id)).replace("PHOTOS", photosURL(id)).replace("VIDEOS", videosURL(id)).replace("VIDEOSOF", videosofURL(id));
    console.log(photosofURL(id));
  }
  </script>


.. _Facebook: https://www.facebook.com/
.. _Find my Facebook ID: http://findmyfbid.com/

Photos of Me on Facebook
########################

:date: 2016-03-24T02:17+08:00
:tags: JavaScript
:category: JavaScript
:summary: Find photos of me on Facebook_.


1. log into Facebook_.

2. .. raw:: html

     <a class="reference external" href="http://findmyfbid.com/" target="_blank">Find my Facebook ID</a>.
     <br>

3. .. raw:: html

     <button type="button" id="getlinks">Get Links</button>
     <input type="text" id="myfbid" placeholder="my Facebook ID">
     <div id="links"></div>

.. raw:: html

  <script>

  var linksElm = document.getElementById("links")

  function photosofURL(id) {
    return "https://www.facebook.com/search/" + id + "/photos-of";
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
      <div><a href="PHOTOSOF" target="_blank">Photos of me</a></div><br> \
      <div><a href="PHOTOSLIKED" target="_blank">Photos liked by me</a></div><br> \
      <div><a href="PHOTOSCOMMENTED" target="_blank">Photos commented on by me</a></div><br> \
    '.replace("PHOTOSOF", photosofURL(id)).replace("PHOTOSLIKED", photoslikedURL(id)).replace("PHOTOSCOMMENTED", photoscommentedURL(id))
    console.log(photosofURL(id));
  }
  </script>


.. _Facebook: https://www.facebook.com/
.. _Find my Facebook ID: http://findmyfbid.com/

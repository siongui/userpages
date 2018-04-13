[Chrome Extension] Get Instagram User Information From HTML Source
##################################################################

:date: 2018-04-13T23:46+08:00
:tags: JavaScript, Chrome Extension, JSON, Instagram, Web Scrape
:category: Chrome Extension
:summary: Chrome extension to help you get user information on user profile
          page. Also show id and invisible link to profile picture on the
          profile page.
:og_image: https://www.otzberg.net/iguserid/instagram_user_id_chrome.png
:adsu: yes


Instagram now bans access to the URL:

  *https://www.instagram.com/<username>/?__a=1*

So now we cannot get user information, such as id, profile pic url, etc. by XHR
to the URL. But I found there is another place to get the user information, that
is, from HTML source of user profile page you are visiting. So I modified my
previous Chrome extension to adapt to the change.

This new Chrome extension gets the JSON format variable *window._sharedData*
from the HTML source and show id and an invisible link to user profile pic. When
you visit the user profile page, press F5 to refresh the page to make this
extension work.

**manifest.json**:

.. code-block:: json

  {
    "manifest_version": 2,

    "name": "showinfo",
    "description" : "Show Instagram Id and invisible profile pic link on user page",
    "version": "0.1",

    "incognito": "split",

    "content_scripts": [
      {
        "matches": ["https://www.instagram.com/*"],
        "js": ["showinfo.js"]
      }
    ],

    "permissions": [
      "https://www.instagram.com/*",
      "tabs"
    ]
  }

**showinfo.js**:

.. code-block:: javascript

  function showId(elm, user) {
    // chech if id already exist
    if (elm.querySelector("a._idofuser") != null) {
      return;
    }

    var span_l8ji8 = elm.querySelector("span._l8ji8");
    if (span_l8ji8 == null) {
      span_l8ji8 = document.createElement("span");
      span_l8ji8.setAttribute("class", "_l8ji8");
      elm.appendChild(span_l8ji8);
    }

    var id = user["id"];

    var link = document.createElement("a");
    link.setAttribute("target", "_blank");
    link.setAttribute("class", "_q3gn4 _idofuser");
    link.appendChild(document.createTextNode(id));

    span_l8ji8.appendChild(document.createElement("br"));
    span_l8ji8.appendChild(document.createTextNode("ID: "));
    span_l8ji8.appendChild(link);
  }

  function addProfilePicDownloadLink(user) {
    var div_b0acm = document.querySelector("div._b0acm");
    if (div_b0acm == null) {
      console.log("no profile pic?");
      return;
    }

    var div = document.createElement("div");
    div.setAttribute("style", "z-index: 55; height: 40px; width: 46px; position: absolute; right: 10px; top: 8px; display: inline-block;");

    var picurl = user["profile_pic_url_hd"];
    var username = user["username"];
    var link = document.createElement("a");
    link.setAttribute("href", picurl);
    link.setAttribute("target", "_blank");
    link.setAttribute("style", "height: 40px; width: 40px; display: inline-block;");

    // show on hover
    link.addEventListener("mouseenter", function(e) {
      e.target.style.background = "#999";
    });
    link.addEventListener("mouseleave", function(e) {
      e.target.style.background = "";
    });

    // filename will not change to username
    // because href is on the different domain
    // https://stackoverflow.com/a/10049353
    link.setAttribute("download", username);

    div.appendChild(link);
    div_b0acm.appendChild(div);
  }

  function findSharedData(elm) {
    if (elm.nodeType == Node.ELEMENT_NODE || elm.nodeType == Node.DOCUMENT_NODE) {
      for (var i=0; i < elm.childNodes.length; i++) {
        // recursively call self
        var result = findSharedData(elm.childNodes[i]);
        if (result != null) {
          return result;
        }
      }
    }

    if (elm.nodeType == Node.TEXT_NODE) {
      if (elm.nodeValue.startsWith("window._sharedData = ")) {
        var jsonString = elm.nodeValue.replace("window._sharedData = ", "");
        jsonString = jsonString.slice(0, -1);
        return JSON.parse(jsonString);
      }
      return null;
    }

    return null;
  }

  var sdata = findSharedData(document);
  if (sdata != null && sdata["entry_data"].hasOwnProperty("ProfilePage")) {
    var user = sdata["entry_data"]["ProfilePage"][0]["graphql"]["user"];

    // wait page to be loaded
    var timerId = setInterval(function() {
      var n = document.querySelector("div._tb97a");
      if (n != null) {
        clearInterval(timerId);
        showId(n, user);
        addProfilePicDownloadLink(user);
      }
    }, 500);
  }

----

.. adsu:: 2

References:

.. [1] `[Golang] Get Instagram User Information From HTML Source <{filename}/articles/2018/04/12/go-get-instagram-user-information-from-html-source%en.rst>`_

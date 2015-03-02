/**
 * Cross-browser HTTP GET request
 * @param {string} url The url of HTTP GET request
 * @param {function} callback The callback function if HTTP GET succeeds
 * @param {function} failCallback The callback function if HTTP GET fails
 */
HTTPGET = function(url, callback, failCallback) {
  var xmlhttp;

  if (window.XMLHttpRequest) {
    // code for IE7+, Firefox, Chrome, Opera, Safari
    xmlhttp=new XMLHttpRequest();
  } else {
    // code for IE6, IE5
    xmlhttp=new ActiveXObject("Microsoft.XMLHTTP");
  }

  xmlhttp.onreadystatechange = function() {
    if (xmlhttp.readyState == 4) {
      if (xmlhttp.status == 200)
        callback(xmlhttp.responseText, url);
      else
        failCallback(url);
    }
  };

  xmlhttp.open("GET", url, true);
  xmlhttp.send();
};

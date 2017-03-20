/**
 * Cross-Browser Cross-Domain XMLHttpRequest (XDomainRequest in IE)
 *
 * @param {string} url The url of HTTP GET (AJAX) request.
 * @param {function} callback The callback function if the request succeeds.
 * @param {function} failCallback The callback function if the request fails.
 */
AJAXRequest = function(url, callback, failCallback) {
  var xmlhttp = new XMLHttpRequest();

  // @see http://blogs.msdn.com/b/ie/archive/2012/02/09/cors-for-xhr-in-ie10.aspx
  // @see http://bionicspirit.com/blog/2011/03/24/cross-domain-requests.html
  // @see http://msdn.microsoft.com/en-us/library/ie/cc288060(v=vs.85).aspx
  if ("withCredentials" in xmlhttp) {
    // for Chrome, Firefox, Opera
    xmlhttp.onreadystatechange = function() {
      if (xmlhttp.readyState == 4) {
        if (xmlhttp.status == 200 || xmlhttp.status == 304) {
          callback(xmlhttp.responseText);
        } else {
          setTimeout(failCallback, 0);
        }
      }
    }

    xmlhttp.open("GET", url, true);
    xmlhttp.send();
  } else {
    // for IE
    var xdr = new XDomainRequest();
    xdr.onerror = function(){setTimeout(failCallback, 0);};
    xdr.ontimeout = function(){setTimeout(failCallback, 0);};
    xdr.onload = function() {
      callback(xdr.responseText);
    };

    xdr.open("get", url);
    xdr.send();
  }
};

/**
 * Callback function of AJAX request if the request succeeds.
 */
callback = function(responseText) {
  // write your own handler here.
  alert('result from https://golden-operator-130720.appspot.com/sukhada.json\n' + responseText);
};

/**
 * Callback function of AJAX request if the request fails.
 */
failCallback = function() {
  // write your own failure handler here.
  alert('AJAXRequest failure!');
};


document.getElementById('bt').onclick = function() {
  AJAXRequest('https://golden-operator-130720.appspot.com/sukhada.json', callback, failCallback);
};

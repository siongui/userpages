/**
 * Cross-Browser AJAX request (XMLHttpRequest)
 *
 * @param {string} url The url of HTTP GET (AJAX) request.
 * @param {function} callback The callback function if the request succeeds.
 * @param {function} failCallback The callback function if the request fails.
 */
AjaxRequest = function(url, callback, failCallback) {
  var xmlhttp;

  if (window.XMLHttpRequest)
    xmlhttp=new XMLHttpRequest();
  else
    xmlhttp=new ActiveXObject("Microsoft.XMLHTTP");

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

/**
 * Issue Multiple AJAX requests to get data, and a single callback is called
 * after all AJAX requests ar completed successfully.
 *
 * @param {Array} urls The urls of HTTP GET (AJAX) requests.
 * @param {function} callbackMulti The callback function to be run after all
 *                                 AJAX requests are completed successfully.
 *                                 The callbackMulti takes one argument, which
 *                                 is data object of url-responseText pairs.
 * @param {function} failCallbackMulti The callback function to be run if one of
 *                                     the AJAX requests fails.
 */
AjaxRequestsMulti = function(urls, callbackMulti, failCallbackMulti) {
  var isAllCallsCompleted = false;
  var isCallFailed = false;
  var data = {};

  for (var i=0; i<urls.length; i++) {
    var callback = function(responseText, url) {
      if (isCallFailed) return;

      data[url] = responseText;

      // get size of data
      var size = 0;
      for (var index in data) {
        if (data.hasOwnProperty(index))
          size ++;
      }

      if (size == urls.length)
        // all AJAX requests are completed successfully
        callbackMulti(data);
    };

    var failCallback = function(url) {
      isCallFailed = true;
      failCallbackMulti(url);
    };

    AjaxRequest(urls[i], callback, failCallback);
  }
};

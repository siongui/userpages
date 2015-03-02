/**
 * Cross-browser HTTP POST request
 * @param {string} url The url of HTTP POST request
 * @param {object} keyValuePairs The object which contains data of key-value
 *                 pair(s) to be POSTed. For example, object {'name': 'Bob',
                   'age': '21'} represents "name=Bob&age=21".
 * @param {function} callback The callback function if HTTP POST succeeds
 * @param {function} failCallback The callback function if HTTP POST fails
 */
HTTPPOST = function(url, keyValuePairs, callback, failCallback) {
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

  xmlhttp.open("POST", url, true);
  xmlhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");

  var kvpairs = '';
  for (var key in keyValuePairs) {
    if (kvpairs == '')
      kvpairs = encodeURIComponent(key) + '=' +
                encodeURIComponent(keyValuePairs[key]);
    else
      kvpairs = kvpairs + '&' + encodeURIComponent(key) + '=' +
                encodeURIComponent(keyValuePairs[key]);
  }

  xmlhttp.send(kvpairs);
};


function ajaxformpost() {
  var keyValuePairs = {
    'input1': document.getElementById('input1').value,
    'input2': document.getElementById('input2').value
  };

  var callback = function(responseText, url) {
    alert('responseText from url ' + url + ':\n'
          + responseText);
  };

  var failCallback = function(url) {
    // write your own handler for failure of HTTP POST
    alert('fail to post ' + url);
  }

  HTTPPOST("/post", keyValuePairs, callback, failCallback);
}

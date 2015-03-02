var url = '/demo/path?a=1&b=2';

var callback = function(responseText, url) {
  // write your own handler for success of HTTP GET
  alert('responseText from url ' + url + ':\n'
        + responseText);
};

var failCallback = function(url) {
  // write your own handler for failure of HTTP GET
  alert('fail to get ' + url);
};

HTTPGET(url, callback, failCallback);

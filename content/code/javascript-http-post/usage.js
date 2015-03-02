var url = '/demo/post';
var keyValuePairs = {'name': 'Bob', 'age': '21'};

var callback = function(responseText, url) {
  // write your own handler for success of HTTP POST
  alert('responseText from url ' + url + ':\n'
        + responseText);
};

var failCallback = function(url) {
  // write your own handler for failure of HTTP POST
  alert('fail to post ' + url);
};

HTTPPOST(url, keyValuePairs, callback, failCallback);

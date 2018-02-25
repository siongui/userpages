var reUrlUser = /^https:\/\/www\.instagram\.com\/([a-z0-9_.]+)\/$/;

function sendMsg(jsonData, url) {
  chrome.tabs.query({active: true, currentWindow: true}, function(tabs) {
    chrome.tabs.sendMessage(
      tabs[0].id,
      {jsonData: jsonData, url: url}
    );
  });
}

function getUserJsonData(url) {
  var jsonUrl = url + "?__a=1";
  var xhr = new XMLHttpRequest();
  xhr.open("GET", jsonUrl, true);
  xhr.onreadystatechange = function() {
    if (xhr.readyState == 4) {
      // JSON.parse does not evaluate the attacker's scripts.
      var resp = JSON.parse(xhr.responseText);
      sendMsg(resp, jsonUrl);
    }
  }
  xhr.send();
}

chrome.tabs.onUpdated.addListener(function(tabId, changeInfo, tab) {
  if (changeInfo.hasOwnProperty('url')) {

    // check if URL is user page
    var result = changeInfo.url.match(reUrlUser);
    if (result) {
      getUserJsonData(changeInfo.url);
    }
  }
});

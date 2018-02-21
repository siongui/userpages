var reUrlUser = /^https:\/\/www\.instagram\.com\/([a-z0-9_.]+)\/$/;

function sendMsg(id, url) {
  chrome.tabs.query({active: true, currentWindow: true}, function(tabs) {
    chrome.tabs.sendMessage(
      tabs[0].id,
      {id: id, url: url}
    );
  });
}

function getId(url) {
  var jsonUrl = url + "?__a=1";
  var xhr = new XMLHttpRequest();
  xhr.open("GET", jsonUrl, true);
  xhr.onreadystatechange = function() {
    if (xhr.readyState == 4) {
      // JSON.parse does not evaluate the attacker's scripts.
      var resp = JSON.parse(xhr.responseText);
      sendMsg(resp["user"]["id"], jsonUrl);
    }
  }
  xhr.send();
}

chrome.tabs.onUpdated.addListener(function(tabId, changeInfo, tab) {
  if (changeInfo.hasOwnProperty('url')) {

    // check if URL is user page
    var result = changeInfo.url.match(reUrlUser);
    if (result) {
      getId(changeInfo.url);
    }
  }
});

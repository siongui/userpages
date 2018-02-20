var reId = /^https:\/\/www\.facebook\.com\/profile\.php\?id=(\d+)$/;
var reUser = /^https:\/\/www\.facebook\.com\/([a-zA-Z0-9.]+)$/;
var lastUrlIsId = false;

chrome.tabs.onUpdated.addListener(function(tabId, changeInfo, tab) {
  if (changeInfo.hasOwnProperty('url')) {

    console.log(changeInfo.url);

    // check if URL consists of username
    var resultUser = changeInfo.url.match(reUser);
    if (resultUser) {
      //console.log(resultUser[1]); // username
      if (lastUrlIsId) {
        console.log(resultUser[1] + " : " + id);
      }
    }

    // check if URL consists of Id
    var resultId = changeInfo.url.match(reId);
    if (resultId) {
      //console.log(resultId[0]); // URL
      //console.log(resultId[1]); // id
      id = resultId[1];
      lastUrlIsId = true;
    } else {
      lastUrlIsId = false;
    }

  }
});

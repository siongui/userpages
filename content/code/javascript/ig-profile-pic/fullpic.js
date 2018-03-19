function getFullSizeProfilePicUrl(picurl) {
  var fullsizeurl = picurl.replace("/vp/", "/");
  fullsizeurl = fullsizeurl.replace("t51.2885-19", "t51.2885-15");
  fullsizeurl = fullsizeurl.replace(/\/s\d+x\d+\//, "/sh0.08/e35/");
  return fullsizeurl;
}

function addProfilePicDownloadLink(jsonData, url) {
  var div_b0acm = document.querySelector("div._b0acm");
  if (div_b0acm == null) {
    console.log("no profile pic?");
    return;
  }

  var div = document.createElement("div");
  div.setAttribute("style", "z-index: 55; height: 40px; width: 46px; position: absolute; right: 10px; top: 8px; display: inline-block;");

  var picurl = getFullSizeProfilePicUrl(jsonData["graphql"]["user"]["profile_pic_url_hd"]);
  var username = jsonData["graphql"]["user"]["username"];
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

chrome.runtime.onMessage.addListener(
  function(request, sender, sendResponse) {

    // wait page to be loaded
    var timerId = setInterval(function() {
      var n = document.querySelector("div._tb97a");
      if (n != null) {
        addProfilePicDownloadLink(request.jsonData, request.url);
        clearInterval(timerId);
      }
    }, 500);
  });

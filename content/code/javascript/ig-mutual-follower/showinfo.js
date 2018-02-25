function showMutualFollowers(elm, jsonData) {
  if (jsonData["user"]["mutual_followers"] == null) {
    return;
  }
  var followers = jsonData["user"]["mutual_followers"]["usernames"];
  var count = jsonData["user"]["mutual_followers"]["additional_count"];
  if (followers.length == 0) {
    return;
  }

  var div = document.createElement("div");
  div.setAttribute("style", "display: inline;");
  div.appendChild(document.createTextNode("Followed by: "));
  while (followers.length > 0) {
    var follower = followers.pop();
    var flink = document.createElement("a");
    flink.setAttribute("href", "https://www.instagram.com/"+follower+"/");
    //flink.setAttribute("target", "_blank");
    flink.setAttribute("class", "_fd86t");
    flink.appendChild(document.createTextNode(follower));
    div.appendChild(flink);
    if (followers.length != 0) {
      //div.appendChild(document.createTextNode(",　"));
      div.appendChild(document.createTextNode(", "));
    } else {
      if (count > 0) {
        div.appendChild(document.createTextNode(" + "+count+" more"));
      }
    }
  }
  elm.appendChild(div);
}

chrome.runtime.onMessage.addListener(
  function(request, sender, sendResponse) {

    // wait page to be loaded
    var timerId = setInterval(function() {
      var n = document.querySelector("section._o6mpc");
      if (n != null) {
        showMutualFollowers(n, request.jsonData);
        clearInterval(timerId);
      }
    }, 500);
  });

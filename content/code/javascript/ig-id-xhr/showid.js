function showId(elm, id, url) {
  var link = document.createElement("a");
  link.setAttribute("href", url);
  link.setAttribute("target", "_blank");
  link.setAttribute("class", "_fd86t");
  link.appendChild(document.createTextNode(id));

  elm.appendChild(document.createElement("br"));
  elm.appendChild(link);
}

chrome.runtime.onMessage.addListener(
  function(request, sender, sendResponse) {

    // wait page to be loaded
    var timerId = setInterval(function() {
      var n = document.querySelector("section._o6mpc");
      if (n != null) {
        showId(n, request.id, request.url);
        clearInterval(timerId);
      }
    }, 500);
  });

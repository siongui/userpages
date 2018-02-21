chrome.runtime.onMessage.addListener(
  function(request, sender, sendResponse) {
    //console.log(request.name);
    //console.log(request.id);

    // wait page to be loaded
    var timerId = setInterval(function() {
      var n = document.querySelector("#fb-timeline-cover-name");
      if (n != null) {
        var idNode = document.createTextNode(request.id);
        n.appendChild(idNode);
        clearInterval(timerId);
      }
    }, 500);
  });

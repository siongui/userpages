import 'dart:html';
// https://www.dartlang.org/articles/json-web-service/
import 'dart:convert';

void httpPostJSON(Event e) {
  ButtonElement elm = e.target as ButtonElement;

  HttpRequest request = new HttpRequest(); // create a new XHR

  // add an event handler that is called when the request finishes
  request.onReadyStateChange.listen((_) {
    if (request.readyState == HttpRequest.DONE &&
        (request.status == 200 || request.status == 0)) {
      // data saved OK.
      // http://stackoverflow.com/questions/12177970/how-do-i-add-arbitrary-html-to-an-element-in-dart
      querySelector("body").appendHtml(request.responseText);
    }
  });

  // POST the data to the server
  var url = "/post";
  request.open("POST", url, async: false);

  request.send(JSON.encode(elm.dataset)); // perform the async POST
}

void main() {
  querySelectorAll("button").forEach((Element elm) {
    elm.onClick.listen(httpPostJSON);
  });
}

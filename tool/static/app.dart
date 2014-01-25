import 'dart:html';
import 'dart:convert';


void handleHtml(String html) {
  print(html);
}

void sendUrl(Event e) {
  ButtonElement elm = e.target as ButtonElement;
  InputElement urlElm = query("#url") as InputElement;
  InputElement tagElm = query("#tag") as InputElement;
  TextAreaElement taElm = query("#info") as TextAreaElement;
  taElm.text = urlElm.value + '\n' + tagElm.value;

  HttpRequest request = new HttpRequest();

  // add an event handler that is called when the request finishes
  request.onReadyStateChange.listen((_) {
    if (request.readyState == HttpRequest.DONE &&
        (request.status == 200 || request.status == 0)) {
      // data saved OK.
      handleHtml(request.responseText);
    }
  });

  // POST the data to the server
  String url = "/url/";
  request.open("POST", url, async: false);
  Map<String, String> jsonData = {
    "Url": urlElm.value,
    "Tag": tagElm.value,
  };
  request.send(JSON.encode(jsonData));
}

void main() {
  query("#getBtn").onClick.listen(sendUrl);
}

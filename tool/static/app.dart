import 'dart:html';
import 'dart:convert';


void sendUrl(Event e) {
  ButtonElement elm = e.target as ButtonElement;
  InputElement urlElm = query("#url") as InputElement;
  InputElement tagElm = query("#tag") as InputElement;
  TextAreaElement taElm = query("#info") as TextAreaElement;
  taElm.text = urlElm.value + '\n' + tagElm.value;

  HttpRequest request = new HttpRequest();

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

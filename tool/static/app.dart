import 'dart:html';


void sendUrl(Event e) {
  ButtonElement elm = e.target as ButtonElement;
  InputElement urlElm = query("#url") as InputElement;
  TextAreaElement taElm = query("#info") as TextAreaElement;
  taElm.text = urlElm.value;
}

void main() {
  query("#getBtn").onClick.listen(sendUrl);
}

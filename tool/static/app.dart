import 'dart:html';


void sendUrl(Event e) {
  ButtonElement elm = e.target as ButtonElement;
  elm.text = "123";
}

void main() {
  query("#getBtn").onClick.listen(sendUrl);
}

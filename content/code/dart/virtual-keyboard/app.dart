import 'dart:html';

void main() {
  InputElement word = querySelector("#userinput");
  Element keypad = querySelector(".keypad");
  var letters = [['q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p'],
                 ['a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l'],
                 ['z', 'x', 'c', 'v', 'b', 'n', 'm'],
                 ['ā', 'ḍ', 'ī', 'ḷ', 'ṁ', 'ṃ', 'ñ', 'ṇ', 'ṭ', 'ū', 'ŋ', 'ṅ'] ];

  for (var row = 0; row < letters.length; row++) {
    var divElm = new Element.tag('div');
    for (var i = 0; i < letters[row].length; i++) {
      var inputElm = new InputElement();
      inputElm.type = 'button';
      inputElm.value = letters[row][i];
      inputElm.onClick.listen((e) => word.value += e.target.value);
      divElm.children.add(inputElm);
    }
    keypad.children.add(divElm);
  }
}

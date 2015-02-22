import 'dart:html';

// http://stackoverflow.com/questions/22747125/use-static-variable-in-function
String _lastInput = "";

String checkLastTwoCharacter(String lastChar, String newChar) {
  if (lastChar == "A" && newChar == "A") return "Ā";
  if (lastChar == "a" && newChar == "a") return "ā";
  if (lastChar == "I" && newChar == "I") return "Ī";
  if (lastChar == "i" && newChar == "i") return "ī";
  if (lastChar == "U" && newChar == "U") return "Ū";
  if (lastChar == "u" && newChar == "u") return "ū";
  if (lastChar == '"' && newChar == "N") return "Ṅ";
  if (lastChar == '"' && newChar == "n") return "ṅ";
  if (lastChar == "." && newChar == "M") return "Ṃ";
  if (lastChar == "." && newChar == "m") return "ṃ";
  if (lastChar == "~" && newChar == "N") return "Ñ";
  if (lastChar == "~" && newChar == "n") return "ñ";
  if (lastChar == "." && newChar == "T") return "Ṭ";
  if (lastChar == "." && newChar == "t") return "ṭ";
  if (lastChar == "." && newChar == "D") return "Ḍ";
  if (lastChar == "." && newChar == "d") return "ḍ";
  if (lastChar == "." && newChar == "N") return "Ṇ";
  if (lastChar == "." && newChar == "n") return "ṇ";
  if (lastChar == "." && newChar == "L") return "Ḷ";
  if (lastChar == "." && newChar == "l") return "ḷ";
  return "";
}

void paliIME(Event e) {
  InputElement elm = e.target;
  // https://api.dartlang.org/apidocs/channels/stable/dartdoc-viewer/dart:core.String
  if (_lastInput.isNotEmpty && elm.value.isNotEmpty) {
    String v = elm.value;
    if ( v.length == (_lastInput.length + 1) ) {
      String result = checkLastTwoCharacter(
                 _lastInput[_lastInput.length-1],
                 v[v.length-1]
               );
      if (result.isNotEmpty) {
        elm.value = v.substring(0, v.length-2) + result;
      }
    }
  }
  _lastInput = elm.value;
}

void main() {
  querySelector("#pali").focus();
  // http://stackoverflow.com/questions/14433156/respond-immediately-to-textarea-changes-in-dart
  querySelector("#pali").onKeyUp.listen(paliIME);
}
/*
TODO: Use selection
http://stackoverflow.com/questions/2897155/get-cursor-position-in-characters-within-a-text-input-field
http://stackoverflow.com/questions/22797294/textarea-cursor-position
http://stackoverflow.com/questions/21730134/caret-position-in-an-editable-div-dartlang
http://stackoverflow.com/questions/28477487/setting-and-getting-selectionrange-caret-position-in-contenteditable-div-element
*/

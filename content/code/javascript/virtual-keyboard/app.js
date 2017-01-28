'use strict';

function createKeypad() {
  var word = document.querySelector("#userinput");
  var keypad = document.querySelector(".keypad");

  const letters = [ ['q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p'],
                    ['a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l'],
                    ['z', 'x', 'c', 'v', 'b', 'n', 'm'],
                    ['ā', 'ḍ', 'ī', 'ḷ', 'ṁ', 'ṃ', 'ñ', 'ṇ', 'ṭ', 'ū', 'ŋ', 'ṅ'] ];

  for (var row = 0; row < letters.length; row++) {
    var divElm = document.createElement("div");
    for (var i = 0; i < letters[row].length; i++) {
      var inputElm = document.createElement("input");
      inputElm.type = "button";
      inputElm.value = letters[row][i];
      divElm.appendChild(inputElm);
      inputElm.addEventListener("click", function(e) {
        word.value += e.target.value;
      }, false);
    }
    keypad.appendChild(divElm);
  }
}

createKeypad();

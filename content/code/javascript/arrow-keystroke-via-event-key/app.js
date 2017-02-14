'use strict';

var Key = {
  LEFT:  "ArrowLeft",
  UP:    "ArrowUp",
  RIGHT: "ArrowRight",
  DOWN:  "ArrowDown"
};

function handleKeyboardEvent(event) {
  var info = document.getElementById("info");
  switch (event.key) {
    case Key.LEFT:
      info.value += "LEFT ";
      break;
    case Key.UP:
      info.value += "UP ";
      break;
    case Key.RIGHT:
      info.value += "RIGHT ";
      break;
    case Key.DOWN:
      info.value += "DOWN ";
      break;
    default:
      info.value += "SOMEKEY ";
  }
}

document.addEventListener("keydown", handleKeyboardEvent);

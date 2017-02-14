var Key = {
  LEFT:   37,
  UP:     38,
  RIGHT:  39,
  DOWN:   40
};

/**
 * old IE: attachEvent
 * Firefox, Chrome, or modern browsers: addEventListener
 */
function _addEventListener(evt, element, fn) {
  if (window.addEventListener) {
    element.addEventListener(evt, fn, false);
  }
  else {
    element.attachEvent('on'+evt, fn);
  }
}

function handleKeyboardEvent(evt) {
  if (!evt) {evt = window.event;} // for old IE compatible
  var keycode = evt.keyCode || evt.which; // also for cross-browser compatible

  var info = document.getElementById("info");
  switch (keycode) {
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

_addEventListener('keydown', document, handleKeyboardEvent);

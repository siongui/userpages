var Key = {
  LEFT:   37,
  UP:     38,
  RIGHT:  39,
  DOWN:   40
};

/* IE: attachEvent, Firefox & Chrome: addEventListener */
function _addEventListener(evt, element, fn) {
  if (window.addEventListener) {element.addEventListener(evt, fn, false);}
  else {element.attachEvent('on'+evt, fn);}
}

function onInputKeydown(evt) {
  if (!evt) {evt = window.event;} // for IE compatible
  var keycode = evt.keyCode || evt.which; // also for cross-browser compatible
  if (keycode == Key.LEFT) {document.getElementById("info").innerHTML += "LEFT ";}
  else if (keycode == Key.UP) {document.getElementById("info").innerHTML += "UP ";}
  else if (keycode == Key.RIGHT) {document.getElementById("info").innerHTML += "RIGHT ";}
  else if (keycode == Key.DOWN) {document.getElementById("info").innerHTML += "DOWN ";}
  else {document.getElementById("info").innerHTML += "SOMEKEY ";}
}

function addevt() {
  _addEventListener('keydown', document, onInputKeydown);
}

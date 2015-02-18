function getPosition( el ) {
  var body = document.documentElement || document.body;
  var scrollX = window.pageXOffset || body.scrollLeft;
  var scrollY = window.pageYOffset || body.scrollTop;
  _x = el.getBoundingClientRect().left + scrollX;
  _y = el.getBoundingClientRect().top + scrollY;
  return { top: _y, left: _x };
}

function onButtonClick(element) {
  alert("Poistion top: " + getPosition(element).top + "\nPosition left: " + getPosition(element).left);
}

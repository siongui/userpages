/* IE: attachEvent, Firefox & Chrome: addEventListener */
function _addEventListener(evt, element, fn) {
  if (window.addEventListener) {element.addEventListener(evt, fn, false);}
  else {element.attachEvent('on'+evt, fn);}
}

function _getParentElement(element, id) {
  /* sometimes mouse event return the child element of actual element we need, so we need to check parent element */
  /* Chrome and Firefox use parentNode, while Opera uses offsetParent */
  while(element.parentNode) {
    if( element.id == id ) {return element;}
    element = element.parentNode;
  }
  while(element.offsetParent) {
    if( element.id == id ) {return element;}
    element = element.offsetParent;
  }
  return null;
}

function onDiv1Click(evt) {
  if(!evt) {evt = window.event;}
  var target = evt.target || evt.srcElement;
  document.getElementById("info").innerHTML = target.id + " clicked";
  if (target.id != 'div1') {
    var ediv1 = _getParentElement(target, 'div1');
    document.getElementById("info").innerHTML += '<br />' + ediv1.id + ' element was gotten';
  }
}

function addevt() {
  _addEventListener('click', document.getElementById("div1"), onDiv1Click);
  document.getElementById("info").innerHTML = "only div1 onclick event registered";
}

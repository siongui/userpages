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

function onDiv1MouseOver(evt) {
  if(!evt) {evt = window.event;}
  var target = evt.target || evt.srcElement;
  document.getElementById("info").innerHTML = target.id + " mouseover";
  if (target.id != 'div1') {
    var ediv1 = _getParentElement(target, 'div1');
    document.getElementById("info").innerHTML += '<br />' + ediv1.id + ' element was gotten';
  }
}

function onDiv1MouseOut(evt) {
  if(!evt) {evt = window.event;}
  var target = evt.target || evt.srcElement;
  document.getElementById("info2").innerHTML = target.id + " mouseout";
  if (target.id != 'div1') {
    var ediv1 = _getParentElement(target, 'div1');
    document.getElementById("info2").innerHTML += '<br />' + ediv1.id + ' element was gotten';
  }
}


function addevt() {
  _addEventListener('mouseover', document.getElementById("div1"), onDiv1MouseOver);
  _addEventListener('mouseout', document.getElementById("div1"), onDiv1MouseOut);
  document.getElementById("info").innerHTML = "only div1 onmouseover event registered";
  document.getElementById("info2").innerHTML = "only div1 onmouseout event registered";
}

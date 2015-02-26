/**
 * Cross-browser addEventListener function.
 *
 * @param {DOM element} element The element to add event listener.
 * @param {string} evt The event to be listened.
 * @param {function} fn The callback function when event occurs.
 */
addEventListener_ = function(element, evt, fn) {
  if (window.addEventListener) {
    /* W3C compliant browser */
    element.addEventListener(evt, fn, false);
  } else {
    /* IE */
    element.attachEvent('on' + evt, fn);
  }
};

/**
 * check if element is targer or the parent of target
 * @param {DOM Element} target
 * @param {DOM Element} element
 * @return {boolean} Return true if element is target or the parent of target
 *                   else return false.
 */
checkParent = function(target, element) {
  // Chrome and Firefox use parentNode, while Opera use offsetParent
  while(target.parentNode) {
    if( target == element ) return true;
    target = target.parentNode;
  }
  while(target.offsetParent) {
    if( target == element ) return true;
    target = target.offsetParent;
  }
  return false;
};


/**
 * Cross-browser add mouse enter event listener. The mouse enter event only
 * fired if mouse enters the element, not fired if mouse enters child element(s)
 * of the registered element.
 *
 * @param {DOM element} element The element to add mouse enter event listener.
 * @param {function} fn The callback function when the mouse enter event occurs.
 */
addMouseEnterEventListener = function(element, fn) {
  var wrapper = function(e) {
    var evt = e || window.event;
    var targetElement = evt.target || evt.srcElement;

    // check if mouse moves inside the element, if yes, return.
    var relTarg = evt.relatedTarget || evt.fromElement;
    if (checkParent(relTarg, element)) return;

    setTimeout(fn, 0);
  };

  addEventListener_(element, 'mouseover', wrapper);
};


/**
 * Cross-browser add mouse leave event listener. The mouse leave event only
 * fired if mouse leaves the element, not fired if mouse leaves child element(s)
 * of the registered element.
 *
 * @param {DOM element} element The element to add mouse leave event listener.
 * @param {function} fn The callback function when the mouse leave event occurs.
 */
addMouseLeaveEventListener = function(element, fn) {
  var wrapper = function(e) {
    var evt = e || window.event;
    var targetElement = evt.target || evt.srcElement;

    // check if mouse moves inside the element, if yes, return.
    var relTarg = evt.relatedTarget || evt.toElement;
    if (checkParent(relTarg, element)) return;

    setTimeout(fn, 0);
  };

  addEventListener_(element, 'mouseout', wrapper);
};

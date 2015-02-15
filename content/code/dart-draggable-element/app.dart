// https://api.dartlang.org/apidocs/channels/stable/dartdoc-viewer/dart:html.Element
// https://api.dartlang.org/apidocs/channels/stable/dartdoc-viewer/dart:html.MouseEvent
// https://www.dartlang.org/articles/improving-the-dom/
// http://stackoverflow.com/questions/14476738/remove-event-listener-with-the-new-library
import 'dart:html';

void draggable(Element elm) {
  int startX, startY, initialMouseX, initialMouseY;
  var docMouseMoveSub, docMouseUpSub;
  elm.style.position = "absolute";

  void mousemove(e) {
    e.preventDefault();
    int dx = e.clientX - initialMouseX;
    int dy = e.clientY - initialMouseY;
    elm.style.top = (startY+dy).toString() + 'px';
    elm.style.left = (startX+dx).toString() + 'px';
  }

  void mouseup(e) {
    docMouseMoveSub.cancel();
    docMouseUpSub.cancel();
  }

  elm.onMouseDown.listen((e) {
    e.preventDefault();
    startX = elm.offsetLeft;
    startY = elm.offsetTop;
    initialMouseX = e.clientX;
    initialMouseY = e.clientY;
    docMouseMoveSub = document.onMouseMove.listen(mousemove);
    docMouseUpSub = document.onMouseUp.listen(mouseup);
  });
}

void main() {
  draggable(querySelector("#dragMe"));
}

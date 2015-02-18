import 'dart:html';

void arrowKeystrokes(Event e) {
  if(e is KeyboardEvent) {
    KeyboardEvent kevt = e as KeyboardEvent;

    switch(kevt.keyCode) {
      case KeyCode.LEFT:
        querySelector("#info").appendHtml("LEFT ");
        break;

      case KeyCode.UP:
        querySelector("#info").appendHtml("UP ");
        break;

      case KeyCode.DOWN:
        querySelector("#info").appendHtml("DOWN ");
        break;

      case KeyCode.RIGHT:
        querySelector("#info").appendHtml("RIGHT ");
        break;

      default:
        querySelector("#info").appendHtml("SOMEKEY ");
    }
  }
}

void main() {
  window.onKeyUp.listen(arrowKeystrokes);
}

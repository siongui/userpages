import 'dart:html';

DivElement toggle = querySelector('#menu-dropdown');
DivElement menu = querySelector('#menuDiv-dropdown');

void check(Event e) {
  if ( !checkParent(e.target, menu) ) {
    //click NOT on the menu
    if ( checkParent(e.target, toggle) ) {
      // click on the link
      // http://stackoverflow.com/questions/13789879/getting-elements-global-document-coordinates-in-dart-aka-some-offset
      menu.style.left = toggle.getBoundingClientRect().left.toString() + 'px';
      // http://stackoverflow.com/questions/17756044/how-do-i-toggle-a-css-class-based-on-a-boolean-with-dart
      menu.classes.toggle('invisible');
    } else {
      // click both outside link and outside menu, hide menu
      menu.classes.add('invisible');
    }
  }
}

bool checkParent(target, elm) {
  while(target.parent != null) {
    if( target == elm ) { return true; }
    target = target.parent;
  }
  return false;
}

void main() {
  document.onClick.listen(check);
}

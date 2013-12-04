import 'dart:html';

void showSidebar(Event e) {
  query(".toggle-sidebar-left-arrow").style.display = "none";
  query(".toggle-sidebar-right-arrow").style.display = "";
  query(".right-helf").style.display = "";
  query(".left-helf").style.width = "50%";
}

void hideSidebar(Event e) {
  query(".toggle-sidebar-right-arrow").style.display = "none";
  query(".toggle-sidebar-left-arrow").style.display = "";
  query(".right-helf").style.display = "none";
  query(".left-helf").style.width = "100%";
}

void main() {
  query(".toggle-sidebar-left-arrow")
    ..style.display = "none"
    ..onClick.listen(showSidebar);

  query(".toggle-sidebar-right-arrow").onClick.listen(hideSidebar);
}

var dropdownMenuDiv = document.getElementById("menuDiv-dropdown");
var dropdownMenu = document.getElementById("menu-dropdown");
// hide popup div when clicking outside the div
// http://www.webdeveloper.com/forum/showthread.php?t=98973
document.onclick = check;
// Event accessing
// http://www.quirksmode.org/js/events_access.html
// Event properties
// http://www.quirksmode.org/js/events_properties.html
function check(e){
  var target = (e && e.target) || (event && event.srcElement);

  if (!checkParent(target, dropdownMenuDiv)) {
    // click NOT on the menu
    if (checkParent(target, dropdownMenu)) {
      // click on the link
      if (dropdownMenuDiv.classList.contains("invisible")) {
        // Dynamically retrieve Html element (X,Y) position with JavaScript
        // http://stackoverflow.com/questions/442404/retrieve-the-position-x-y-of-an-html-element
        dropdownMenuDiv.style.left = dropdownMenu.getBoundingClientRect().left + 'px';
        dropdownMenuDiv.classList.remove("invisible");
      } else {dropdownMenuDiv.classList.add("invisible");}
    } else {
      // click both outside link and outside menu, hide menu
      dropdownMenuDiv.classList.add("invisible");
    }
  }
}

function checkParent(t, elm) {
  while(t.parentNode) {
    if( t == elm ) {return true;}
    t = t.parentNode;
  }
  return false;
}

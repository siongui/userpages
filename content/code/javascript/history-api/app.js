"use strict";

var info = document.getElementById("info");
var persons = document.querySelectorAll(".person");

var i;
for (i=0; i<persons.length; i++) {
  persons[i].addEventListener("click", function(event) {
    event.preventDefault();
    var person = event.target;
    var href = person.getAttribute("href");
    var content = person.dataset.content;
    window.history.pushState(content, null, href);
    info.innerHTML = content;
  }, false);
}

window.addEventListener("popstate",function(event) {
  if (event.state === null) {
    info.innerHTML = "Entry Page";
  } else {
    info.innerHTML = event.state;
  }
}, false);

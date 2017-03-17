package main

const indexHtml = `<!DOCTYPE html>
<html>
<head><title>Go JSONP Server</title></head>
<body>
<button id="btn">Click to get HTTP header via JSONP</button>
<pre id="result"></pre>
<script>
'use strict';

var btn = document.getElementById("btn");
var result = document.getElementById("result");

function myCallback(acptlang) {
  result.innerHTML = JSON.stringify(acptlang, null, 2);
}

function jsonp() {
  result.innerHTML = "Loading ...";
  var tag = document.createElement("script");
  tag.src = "/jsonp?callback=myCallback";
  document.querySelector("head").appendChild(tag);
}

btn.addEventListener("click", jsonp);
</script>
</body>
</html>`

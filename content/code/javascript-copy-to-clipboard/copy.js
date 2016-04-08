var textareaElm = document.getElementById("text");
var copyBtnElm = document.getElementById("copy");

copyBtnElm.onclick = function(event) {
  textareaElm.select();
  var isSuccessful = document.execCommand('copy');
  if (isSuccessful) {
    textareaElm.value = "Copy OK";
  } else {
    textareaElm.value = "Copy Fail";
  }
}

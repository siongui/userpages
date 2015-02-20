/* send data to GAE Python server */
function jsonp() {
  var input1value = document.getElementById('input1').value;
  var input2value = document.getElementById('input2').value;
  var url = '/jsonp?callback=mycallback' +
           '&input1=' + encodeURIComponent(input1value) +
           '&input2=' + encodeURIComponent(input2value);
  var ext = document.createElement('script');
  ext.setAttribute('src', url);
  document.getElementsByTagName("head")[0].appendChild(ext);
}

/* receive JSON data from GAE Python server */
function mycallback(JSONdata) {
  /* In order to parse data, we have to know the structure of data from server in advance */

  /* show returned data */
  var infoElm = document.getElementById('info');
  infoElm.innerHTML = 'input1: ' + JSONdata[0]['input1'] + '<br />';
  infoElm.innerHTML += 'input2: ' + JSONdata[1]['input2'] + '<br />';
  infoElm.innerHTML += JSONdata[2] + '<br />';
  infoElm.innerHTML += JSONdata[3] + '<br />';
}

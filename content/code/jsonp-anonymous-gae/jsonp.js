/* send data to GAE Python server */
function jsonp() {
  var input1value = document.getElementById('input1').value;
  var input2value = document.getElementById('input2').value;

  // anonymous callback function to process JSON data returned from server.
  var callback = function(JSONdata) {
    /* In order to parse data, we have to know the structure of data from server in advance */
    /* show data returned from server */
    var infoElem = document.getElementById('info');
    infoElem.innerHTML  = 'input1: ' + JSONdata[0]['input1'] + '<br />';
    infoElem.innerHTML += 'input2: ' + JSONdata[1]['input2'] + '<br />';
    infoElem.innerHTML += JSONdata[2] + '<br />';
    infoElem.innerHTML += JSONdata[3] + '<br />';
  };

  var url = '/jsonp?callback=' + encodeURIComponent(callback.toString()) +
           '&input1=' + encodeURIComponent(input1value) +
           '&input2=' + encodeURIComponent(input2value);

  var ext = document.createElement('script');
  ext.setAttribute('src', url);
  document.getElementsByTagName("head")[0].appendChild(ext);
}

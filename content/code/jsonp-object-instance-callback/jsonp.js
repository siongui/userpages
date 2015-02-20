// object contructor method
demoObj = function(formId) {
  var form = document.getElementById(formId);
  form.action = "javascript:void(0);";
  form.onsubmit = this.jsonp.bind(this);

  window['demoObjInGlobal'] = this;
};

demoObj.prototype.jsonp = function() {
  /* send data to Google App Engine Python server */
  var input1value = document.getElementById('input1').value;
  var input2value = document.getElementById('input2').value;

  var url = '/jsonp?callback=' +
        encodeURIComponent('demoObjInGlobal["callback"]') +
        '&input1=' + encodeURIComponent(input1value) +
        '&input2=' + encodeURIComponent(input2value);

  var ext = document.createElement('script');
  ext.setAttribute('src', url);
  document.getElementsByTagName("head")[0].appendChild(ext);
};

demoObj.prototype.callback = function(JSONdata) {
  /* In order to parse data, we have to know the structure of data from server in advance */
  /* show data returned from server */
  var infoElm = document.getElementById('info');
  infoElm.innerHTML = 'input1: ' + JSONdata[0]['input1'] + '<br />';
  infoElm.innerHTML += 'input2: ' + JSONdata[1]['input2'] + '<br />';
  infoElm.innerHTML += JSONdata[2] + '<br />';
  infoElm.innerHTML += JSONdata[3] + '<br />';
};

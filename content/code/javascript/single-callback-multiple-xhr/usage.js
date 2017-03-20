/**
 * callback function after all AJAX requests ar completed successfully.
 * @param {object} data The JavaScript object which contains url-responseText
 *                      pairs.
 */
callbackMulti = function(data) {
  // write your own callbackMulti here.
  for (var index in data)
    alert(index + ':\n' + data[index]);
};

/**
 * callback function if one of AJAX requests fails.
 * @param {string} url The url of AJAX request that fails.
 */
failCallbackMulti = function(url) {
  // write your own failCallbackMulti here.
  alert(url + ' failed');
};

var urls = ['https://golden-operator-130720.appspot.com/ekadesa.json',
            'https://golden-operator-130720.appspot.com/caturassa.json',
            'https://golden-operator-130720.appspot.com/buddho.json'];

document.getElementById('bt').onclick = function() {
  AjaxRequestsMulti(urls, callbackMulti, failCallbackMulti);
};

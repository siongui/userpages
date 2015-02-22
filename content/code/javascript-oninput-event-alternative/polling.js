/**
 * event handler to be fired if the content of input element have changed.
 */
function eventHandler() {
  // write your own event handler here.
  document.getElementById('info').innerHTML = 'input content changed to ' +
    document.getElementById('input').value;
}

/**
 * Variable which keeps track of the content of input element
 * @type {string}
 */
INPUT_CONTENT = '';

/**
 * Check input element periodically.
 */
polling = function() {
  if (INPUT_CONTENT != document.getElementById('input').value)
    eventHandler();

  // update INPUT_CONTENT
  INPUT_CONTENT = document.getElementById('input').value;

  // polling interval: 500ms.
  // you can change the polling interval to fit your need.
  setTimeout(polling, 500);
};

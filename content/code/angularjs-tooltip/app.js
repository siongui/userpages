'use strict';


angular.module('demoTooltip', ['tooltip']).
  run(['tooltip',
  function(tooltip) {
    // when user's mouse hovers over words, delay a period of time before look up.
    var DELAY_INTERVAL = 1000; // ms

    var isMouseInWord = false;

    function onWordMouseOver(e) {
      isMouseInWord = true;
      this.style.color = 'red';

      setTimeout(angular.bind(this, function() {
        // 'this' keyword here refers to raw dom element
        if (this.style.color === 'red') {
          // mouse is still on word
          var word = this.innerHTML.toLowerCase();
          tooltip.setContent(
            word + ' ' + word + '<br>' +
            '<span>' + word + '</span>' + ' ' + word
          );
          tooltip.setPosition({
            'left': (this.getBoundingClientRect().left + 'px'),
            'top': (this.getBoundingClientRect().top + this.offsetHeight + 'px'),
          });
          tooltip.show();
        }
      }), DELAY_INTERVAL);
    }

    function onWordMouseOut(e) {
      isMouseInWord = false;
      this.style.color = '';

      setTimeout(angular.bind(this, function() {
        if (!isMouseInWord) {
          tooltip.hide();
        }
      }), DELAY_INTERVAL);
    }

    var spans = document.getElementById("container").querySelectorAll("span");
    for (var i=0; i < spans.length; i++) {
      spans[i].onmouseover = onWordMouseOver;
      spans[i].onmouseout = onWordMouseOut;
    }

  }]);

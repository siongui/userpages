'use strict';


angular.module('demoTooltip', ['tooltip']).
  run(['AddTooltipToElement',
  function(AddTooltipToElement) {

    var spans = document.getElementById("container").querySelectorAll("span");
    for (var i=0; i < spans.length; i++) {
      AddTooltipToElement.add(spans[i]);
    }

  }]);

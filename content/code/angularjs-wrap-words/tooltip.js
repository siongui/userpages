'use strict';

/* Services */


angular.module('tooltip', []).
  factory('tooltip', ['$rootScope', '$compile', function($rootScope, $compile) {
    // init tooltip
    var scope = $rootScope.$new(true);
    var isMouseInTooltip = false;
    scope.onmouseenter = function() {
      // mouse enters tooltip
      isMouseInTooltip = true;
    };
    scope.onmouseleave = function() {
      // mouse leaves tooltip
      isMouseInTooltip = false;
      hide();
    };
    var _left = 0;
    var _top = 0;

    var tooltip = $compile('<div class="tooltip" ng-mouseenter="onmouseenter()" ng-mouseleave="onmouseleave()"></div>')(scope);
    tooltip.css('max-width', viewWidth() + 'px');

    // append tooltip to the end of body element
    angular.element(document.getElementsByTagName('body')[0]).append(tooltip);


    function viewWidth() {
      // FIXME: why -32 here?
      return (window.innerWidth || document.documentElement.clientWidth) - 32;
    }

    function setContent(content) {
      tooltip.children().remove();
      if (angular.isUndefined(content)) {
        throw 'In tooltip: content undefined!';
      } else if (angular.isString(content)) {
        tooltip.html(content);
      } else {
        tooltip.append(content);
      }
    }

    function setPosition(position) {
      _left = parseInt(position.left.replace('px', ''));
      _top = parseInt(position.top.replace('px', ''));
    }

    function show() {
      // move tooltip to the right
      // (don't cross the right side of browser inner window)
      var _right = _left + tooltip.prop('offsetWidth');
      if ( _right > viewWidth() )
        _left -= (_right - viewWidth());

      tooltip.css('left', _left + 'px');
      tooltip.css('top', _top + 'px');
    }

    function hide() {
      if (!isMouseInTooltip)
        tooltip.css('left', '-9999px');
    }

    var serviceInstance = {
      viewWidth: viewWidth,
      isHidden: function() {return (tooltip.css('left') === '-9999px');},
      getLeftSpace: function() {return _left;},
      getRightSpace: function() {return viewWidth() - _left - tooltip.prop('offsetWidth');},
      setContent: setContent,
      setPosition: setPosition,
      show: show,
      hide: hide
    };
    return serviceInstance;
  }]).

  factory('AddTooltipToElement', ['tooltip', function(tooltip) {
    // when user's mouse hovers over words,
    // delay a period of time before look up.
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

    function add(elm) {
      elm.onmouseover = onWordMouseOver;
      elm.onmouseout = onWordMouseOut;
    }

    var serviceInstance = { add: add };
    return serviceInstance;
  }]);

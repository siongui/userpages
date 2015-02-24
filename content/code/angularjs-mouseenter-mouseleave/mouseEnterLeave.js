'use strict';

/* Directives */


angular.module('pali.mouseEnterLeave', []).
  directive('mouseenter', [function() {
    return {
      restrict: 'A',
      link: function(scope, elm, attrs) {
        /**
         * check if element is targer or the parent of target
         * @param {DOM Element} target
         * @param {DOM Element} element
         * @return {boolean} Return true if element is target or the parent of target
         *                   else return false.
         */
        function checkParent(target, element) {
          if (angular.isUndefined(target)) return false;
          // Chrome and Firefox use parentNode, while Opera use offsetParent
          while(target.parentNode) {
            if( target == element ) return true;
            target = target.parentNode;
          }
          while(target.offsetParent) {
            if( target == element ) return true;
            target = target.offsetParent;
          }
          return false;
        };

        elm.bind('mouseover', function(e) {
          var evt = e || window.event;
          var targetElement = evt.target || evt.srcElement;

          // check if mouse moves inside the element, if yes, return.
          var relTarg = evt.relatedTarget || evt.fromElement;
          if (checkParent(relTarg, elm[0])) return;

          // https://gist.github.com/siongui/a8d9a9003772315e2cba
          if (scope.$root.$$phase)
            scope.$eval(attrs['mouseenter']);
          else
            scope.$apply(attrs['mouseenter']);
        });
      }
    };
  }]).

  directive('mouseleave', [function() {
    return {
      restrict: 'A',
      link: function(scope, elm, attrs) {
        /**
         * check if element is targer or the parent of target
         * @param {DOM Element} target
         * @param {DOM Element} element
         * @return {boolean} Return true if element is target or the parent of target
         *                   else return false.
         */
        function checkParent(target, element) {
          if (angular.isUndefined(target)) return false;
          // Chrome and Firefox use parentNode, while Opera use offsetParent
          while(target.parentNode) {
            if( target == element ) return true;
            target = target.parentNode;
          }
          while(target.offsetParent) {
            if( target == element ) return true;
            target = target.offsetParent;
          }
          return false;
        };

        elm.bind('mouseout', function(e) {
          var evt = e || window.event;
          var targetElement = evt.target || evt.srcElement;

          // check if mouse moves inside the element, if yes, return.
          var relTarg = evt.relatedTarget || evt.toElement;
          if (checkParent(relTarg, elm[0])) return;

          // https://gist.github.com/siongui/a8d9a9003772315e2cba
          if (scope.$root.$$phase)
            scope.$eval(attrs['mouseleave']);
          else
            scope.$apply(attrs['mouseleave']);
        });
      }
    };
  }]);

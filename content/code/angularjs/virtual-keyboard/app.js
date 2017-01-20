'use strict';

angular.module('virtualKeypad', []).
  run(['$rootScope', function($rootScope) {
    $rootScope.addLtr = function(letter) {
      if (angular.isUndefined($rootScope.paliWord)) {
        $rootScope.paliWord = letter;
      } else {
        $rootScope.paliWord += letter;
      }
    };
}]);

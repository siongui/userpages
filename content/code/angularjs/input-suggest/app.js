'use strict';

angular.module('inputSuggestApp', [])
  .controller('suggestController', ['$scope', 'suggestService', function($scope, ss) {
    $scope.word = "";
    $scope.suggestedWords = [];
    $scope.inputChange = function() {
      // empty $scope.suggestedWords
      while ($scope.suggestedWords.length > 0) {
        $scope.suggestedWords.pop();
      }
      $scope.suggestedWords = ss.GetSuggestedWords($scope.word);
    };
    $scope.isShowSuggestMenu = function() {
      return $scope.suggestedWords.length > 0;
    };

    var inputElm = angular.element(document.querySelector('input'));
    $scope.suggestMenuStyle = {
      left: inputElm.prop('offsetLeft') + 'px',
      minWidth: inputElm.prop('offsetWidth') + 'px',
    };
  }])
  .factory('suggestService', [function() {
    // implement your suggest function here.
    // return random strings for demo purpose.

    function getRandomInt(min, max) {
      return Math.floor(Math.random() * (max - min + 1)) + min;
    }

    function RandomString(strlen) {
      const chars = "abcdefghijklmnopqrstuvwxyz0123456789";
      var result = "";
      for (var i=0; i<strlen; i++) {
        result += chars[getRandomInt(0,35)];
      }
      return result;
    }

    function GetSuggestedWords(word) {
      // Remove leading and trailing whitespaces
      var wordTrimed = word.replace(/(^\s+)|(\s+$)/g, "");

      var suggestedWords = [];
      if (wordTrimed.length > 0) {
        for(var i=0; i<10; i++) {
          suggestedWords.push(RandomString(wordTrimed.length));
        }
      }

      return suggestedWords;
    }

    return { GetSuggestedWords: GetSuggestedWords };
  }]);

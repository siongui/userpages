// tested on 1.0.4 and 1.3.11
angular.module('event', [])
  .controller('parentCtrl', ['$scope', function($scope) {
    $scope.parentProperty = 0;
    $scope.$on('updateParentScopeEvent', function(event, data) {
      $scope.parentProperty = data;
    });
  }])
  .controller('childCtrl', ['$scope', function($scope) {
    $scope.update = function() {
      $scope.$emit('updateParentScopeEvent', 1)
    };
  }]);

// tested on AngularJS 1.3.11 and 1.0.4
angular.module('demo', [])
  .directive('datepicker', ['$parse', function($parse) {
    var directiveDefinitionObject = {
      restrict: 'A',
      link: function postLink(scope, iElement, iAttrs) {
        iElement.datepicker({
          dateFormat: 'yy-mm-dd',
          onSelect: function(dateText, inst) {
            scope.$apply(function(scope){
              $parse(iAttrs.ngModel).assign(scope, dateText);
            });
          }
        });
      }
    };
    return directiveDefinitionObject;
  }])
  .controller('myCtrl', ['$scope', function($scope) {
    $scope.date = "1212-12-12";
  }]);

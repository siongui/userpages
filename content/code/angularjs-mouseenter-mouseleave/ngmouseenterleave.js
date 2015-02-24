'use strict';

angular.module('demo', []).
  run(['$rootScope', '$compile', function($rootScope, $compile) {
    $rootScope.timesEntered = 0;
    $rootScope.timesLeft = 0;
    var scope = $rootScope.$new(true);
    scope.ngmouseenter = function() {
      $rootScope.timesEntered += 1;
    };
    scope.ngmouseleave = function() {
      $rootScope.timesLeft += 1;
    };

    var divNg = $compile('<div style="border: red solid 1px; width: 300px; height: 300px" ng-mouseenter="ngmouseenter()" ng-mouseleave="ngmouseleave()">I am div using ng-mouseenter and ng-mouseleave.</div>')(scope);

    var elm = $compile('<div style="background-color: yellow;">Please mouse in me and then move to the aqua-color region<br /><br /><br /><br /></div>')(scope);
    var elm2 = $compile('<div style="background-color: aqua;">You will see ng-mouseenter and ng-mouseleave called again (incorrect) if mouse moved in me from yellow region<br /><br /><br /><br /></div>')(scope);
    divNg.append(elm);
    divNg.append(elm2);

    var body = angular.element(document.getElementsByTagName('body')[0]);
    body.append(divNg);
  }]);

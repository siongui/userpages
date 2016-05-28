'use strict';


angular.module('demoI18nDirective', ['i18nDirective', "gettextData"]).
  run(['$rootScope', function($rootScope) {
    // default locale is English
    $rootScope.i18nLocale = "en";
  }]);

'use strict';


angular.module('demoGettext', ['gettext', "gettextData"]).
  run(['$rootScope', function($rootScope) {
    // default locale is English
    $rootScope.i18nLocale = "en";
  }]);

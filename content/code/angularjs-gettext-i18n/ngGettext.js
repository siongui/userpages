'use strict';


angular.module('gettext', ['gettextData']).
  run(['$rootScope', 'i18nserv', function($rootScope, i18nserv) {
  // initialization code (similar to main)
    /**
     * usage: {{_("i18n_string")}}
     */
    $rootScope._ = function(str) {
      return i18nserv.gettext(str, $rootScope.i18nLocale);
    };
  }]).

  factory('i18nserv', ['gettextData', function(gettextData) {
  // service: for translating texts according to locale

    var i18nStr = gettextData.all;

    function gettext(value, locale) {
      if (i18nStr.hasOwnProperty(locale)) {
        if (i18nStr[locale].hasOwnProperty(value)) {
          if (i18nStr[locale][value] !== '' &&
              i18nStr[locale][value] !== null)
            return i18nStr[locale][value];
        }
      }
      return value;
    }

    return { gettext: gettext };
  }]);

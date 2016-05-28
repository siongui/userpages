'use strict';


angular.module('i18nDirective', ['gettextData']).
  directive('i18n', ['i18nserv', '$rootScope', function(i18nserv, $rootScope) {
    /**
     * wrap the string to be translated in ELEMENT
     * with attribute 'i18n', 'str', and 'locale'(optional)
     * example: <ELEMENT i18n str='Home'>Home</ELEMENT>
     *      or  <ELEMENT i18n locale={{locale}} str='Home'>Home</ELEMENT>
     */
    return {
      restrict: 'A',
      link: function(scope, elm, attrs) {
        // if "locale" attribute exists, use it
        attrs.$observe('locale', function() {
          var trText = i18nserv.gettext(attrs.str, attrs.locale);
          elm.html(trText);
        });

        // if there is no "locale" attribute, use $rootScope.i18nLocale
        if (angular.isUndefined(attrs.locale)) {
          $rootScope.$watch('i18nLocale', function() {
            var trText = i18nserv.gettext(attrs.str, $rootScope.i18nLocale);
            elm.html(trText);
          });
        }
      }
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

'use strict';

/* Directives */


angular.module('treeview', []).
  directive('treeview', ['$compile', 'treeviewData', function($compile, treeviewData) {
    return {
      restrict: 'A',
      link: function(scope, elm, attrs) {

        // show items
        for (var i=0; i< treeviewData.all['child'].length; i++) {
          var node = traverseTreeviewData( treeviewData.all['child'][i]);
          elm.append(node);
        }

        function traverseTreeviewData(node) {
          var text = node['text'];
          if (node['child']) {
            // not leaf node, keys: 'text', 'child'
            var element = angular.element('<div class="item"></div>');
            var sign = angular.element('<span>+</span>');
            var textElm = $compile(
                '<span class="treeNode">'+ text + '</span>'
                )(scope);

            element.append(sign);
            element.append(textElm);

            var childrenContainer = $compile('<div class="childrenContainer"></div>')(scope);
            for (var i=0; i<node['child'].length; i++) {
              var child = node['child'][i];
              childrenContainer.append(traverseTreeviewData(child));
            }
            childrenContainer.css('display', 'none');

            textElm.bind('click', function() {
              if (childrenContainer.css('display') === 'none') {
                childrenContainer.css('display', '');
                sign.html('-');
              } else {
                childrenContainer.css('display', 'none');
                sign.html('+');
              }
            });

            var all = angular.element('<div></div>');
            all.append(element);
            all.append(childrenContainer);
            return all;
          } else {
            // leaf node, keys: 'text'
            var element = $compile(
                '<div class="item"><span class="treeNode">' + text + '</span></div>'
                )(scope);

            return element;
          }
        }
      }
    };
  }]);

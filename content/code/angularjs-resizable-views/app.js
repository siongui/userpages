'use strict';


angular.module('demoResizableViews', ['resizableViews']).
  run(['resizableViews',
  function(resizableViews) {
    // initialize resizable views
    resizableViews.initViews('allContainer', 'leftview', 'viewwrapper', 'viewarrow', 'viewseparator', 'rightview');
  }]);

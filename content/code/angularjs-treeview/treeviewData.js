angular.module('treeviewData', []).
  factory('treeviewData', [function() {
    var data = {"child": [
      { "text": "item 1",
        "child": [
          { "text": "item 1-1" },
          { "text": "item 1-2",
            "child": [
              { "text": "item 1-2-1" },
              { "text": "item 1-2-2" }
            ]},
          { "text": "item 1-3" }
        ]},
      { "text": "item 2",
        "child": [
          { "text": "item 2-1" },
          { "text": "item 2-2" }
        ]},
      { "text": "item 3" }
    ]}
    var serviceInstance = { all: data };
    return serviceInstance;
  }]);

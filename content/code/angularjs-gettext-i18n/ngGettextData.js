angular.module('gettextData', []).
  factory('gettextData', [function() {
    var data = {
      "zh_TW": {
        "Home": "首頁",
        "Canon": "經典"
      },
      "vi_VN": {
        "Home": "Trang chính",
        "Canon": "Kinh điển"
      }
    }
    var serviceInstance = { all: data };
    return serviceInstance;
  }]);
